package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/herdworks/clic"
	"github.com/herdworks/htskills"
)

// SkillUpCommand synchronizes bundled HyperTexting skills and reports each result.
func SkillUpCommand() clic.Command {
	return clic.Command{
		Meta: clic.CommandMeta{
			Name:        "up",
			Description: "Install and update bundled HyperTexting skills.",
			Usage: []string{
				"{{ .Command }} [--force] [skills-directory]",
			},
			Order: 6,
		},
		Spec: clic.CommandSpec{
			Key:  "up",
			Help: clic.DefaultHelp,
			Options: clic.Options{
				"force": {
					Meta: clic.OptionMeta{
						Name:        "force",
						Description: "Replace installed skills regardless of their versions.",
					},
					Spec: clic.OptionSpec{Key: "force", Kind: clic.BoolOption},
				},
			},
			Validator: skillUpArguments,
			Handler: func(execution *clic.Execution) error {
				var destination string
				if execution.Command.PositionalArguments() == 1 {
					destination = execution.Command.State.Arguments[0]
				}
				force := execution.Options.BoolValue("force")
				results, err := upSkills(execution.WorkingDirectory, execution.LookupEnv, execution.CommandPathString(), commandPrefix(execution), destination, force)
				if err != nil {
					return err
				}
				return writeSkillUpResults(execution.Stdout, results)
			},
		},
	}
}

func skillUpArguments(execution *clic.Execution) error {
	if execution.Command.PositionalArguments() > 1 {
		return positionalArgumentsError(execution, 1)
	}
	var destination string
	if execution.Command.PositionalArguments() == 1 {
		destination = execution.Command.State.Arguments[0]
	}
	_, err := skillUpRoots(execution.WorkingDirectory, execution.LookupEnv, execution.CommandPathString(), destination)
	return err
}

func upSkills(workspace string, lookupEnv func(string) (string, bool), commandPath, command, destination string, force bool) ([]skillUpResult, error) {
	roots, err := skillUpRoots(workspace, lookupEnv, commandPath, destination)
	if err != nil {
		return nil, err
	}
	var results []skillUpResult
	for _, root := range roots {
		synced, err := syncSkills(root, command, force)
		if err != nil {
			return nil, err
		}
		results = append(results, synced...)
	}
	return results, nil
}

func skillUpRoots(workspace string, lookupEnv func(string) (string, bool), commandPath, destination string) ([]string, error) {
	globalRoots, err := globalSkillRoots(lookupEnv)
	if err != nil {
		return nil, err
	}
	if destination != "" && isGlobalSkillRoot(destination, globalRoots) {
		return nil, globalSkillsError([]string{destination})
	}
	installed, err := installedGlobalSkills(globalRoots)
	if err != nil {
		return nil, err
	}
	if len(installed) > 0 {
		return nil, globalSkillsError(installed)
	}
	roots, err := discoverSkillRoots(workspace)
	if err != nil {
		return nil, err
	}
	if destination != "" {
		roots = append(roots, destination)
	}
	roots = uniquePaths(roots)
	if len(roots) == 0 {
		return nil, fmt.Errorf("%s requires a skills directory; retry with %q using your agent's project-local or global skills directory", commandPath, commandPath+" <skills-directory>")
	}
	return roots, nil
}

func syncSkills(destination, command string, force bool) ([]skillUpResult, error) {
	info, err := os.Stat(destination)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		if err = os.MkdirAll(destination, 0755); err != nil {
			return nil, err
		}
	} else if !info.IsDir() {
		return nil, fmt.Errorf("skills directory is not a directory: %s", destination)
	}
	skills, err := htskills.Catalog()
	if err != nil {
		return nil, err
	}
	results := []skillUpResult{{Action: skillUpDiscovered, Path: destination}}
	for _, skill := range skills {
		synced, err := syncSkill(destination, skill, command, force)
		if err != nil {
			return nil, err
		}
		results = append(results, synced...)
	}
	for _, skill := range skills {
		replaced, err := removeReplacementAliases(destination, skill)
		if err != nil {
			return nil, err
		}
		results = append(results, replaced...)
	}
	return results, nil
}

type installedSkill struct {
	Path     string
	Metadata htskills.Skill
}

type skillUpAction string

const (
	skillUpDiscovered skillUpAction = "discovered"
	skillUpInstalled  skillUpAction = "installed"
	skillUpCurrent    skillUpAction = "current"
	skillUpUpdated    skillUpAction = "updated"
	skillUpReplaced   skillUpAction = "replaced"
)

type skillUpResult struct {
	Action  skillUpAction
	Path    string
	Version string
}

func writeSkillUpResults(output io.Writer, results []skillUpResult) error {
	for _, result := range results {
		switch result.Action {
		case skillUpDiscovered:
			if _, err := fmt.Fprintf(output, "skills %q discovered\n", result.Path); err != nil {
				return err
			}
		case skillUpInstalled:
			if _, err := fmt.Fprintf(output, "skill %q installed\n", result.Path); err != nil {
				return err
			}
		case skillUpCurrent:
			if _, err := fmt.Fprintf(output, "skill %q current\n", result.Path); err != nil {
				return err
			}
		case skillUpUpdated:
			if _, err := fmt.Fprintf(output, "skill %q updated to %q\n", result.Path, result.Version); err != nil {
				return err
			}
		case skillUpReplaced:
			if _, err := fmt.Fprintf(output, "skill %q removed\n", result.Path); err != nil {
				return err
			}
		}
	}
	return nil
}

func syncSkill(destination string, skill htskills.Skill, command string, force bool) ([]skillUpResult, error) {
	content, err := htskills.Read(skill.Name)
	if err != nil {
		return nil, err
	}
	embedded, err := htskills.Parse(content)
	if err != nil {
		return nil, fmt.Errorf("bundled skill %q: %w", skill.Name, err)
	}
	path := filepath.Join(destination, skill.Name)
	result := skillUpResult{Path: filepath.Join(path, "SKILL.md"), Version: embedded.Metadata.Version}
	if _, err = os.Lstat(path); os.IsNotExist(err) {
		if _, err = installSkill(skill.Name, path, command); err != nil {
			return nil, fmt.Errorf("install %q: %w", skill.Name, err)
		}
		result.Action = skillUpInstalled
	} else if err != nil {
		return nil, err
	} else if written, err := updateSkillWithForce(skill.Name, path, command, force); err != nil {
		return nil, fmt.Errorf("update %q: %w", skill.Name, err)
	} else if len(written) == 0 {
		result.Action = skillUpCurrent
	} else {
		result.Action = skillUpUpdated
	}
	return []skillUpResult{result}, nil
}

func removeReplacementAliases(destination string, skill htskills.Skill) ([]skillUpResult, error) {
	aliases, err := installedReplacementAliases(destination, skill)
	if err != nil {
		return nil, err
	}
	results := make([]skillUpResult, 0, len(aliases))
	for _, alias := range aliases {
		if err = os.RemoveAll(alias.Path); err != nil {
			return nil, fmt.Errorf("remove replaced skill %q at %s: %w", alias.Metadata.Name, alias.Path, err)
		}
		results = append(results, skillUpResult{
			Action: skillUpReplaced,
			Path:   filepath.Join(alias.Path, "SKILL.md"),
		})
	}
	return results, nil
}

func installedReplacementAliases(destination string, skill htskills.Skill) ([]installedSkill, error) {
	replacementAliases := skill.ReplacementAliases()
	aliases := make(map[string]struct{}, len(replacementAliases))
	for _, alias := range replacementAliases {
		aliases[alias] = struct{}{}
	}
	installed, err := installedSkills(destination)
	if err != nil {
		return nil, err
	}
	matching := make([]installedSkill, 0, len(installed))
	for _, candidate := range installed {
		if _, ok := aliases[candidate.Metadata.Name]; ok {
			matching = append(matching, candidate)
		}
	}
	return matching, nil
}

func installedSkills(root string) ([]installedSkill, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	installed := make([]installedSkill, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		path := filepath.Join(root, entry.Name())
		content, err := os.ReadFile(filepath.Join(path, "SKILL.md"))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		metadata, err := htskills.Parse(content)
		if err != nil {
			continue
		}
		installed = append(installed, installedSkill{Path: path, Metadata: metadata})
	}
	return installed, nil
}

func discoverSkillRoots(workspace string) ([]string, error) {
	if workspace == "" {
		workspace = "."
	}
	var roots []string
	err := filepath.WalkDir(workspace, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() && entry.Name() == "SKILL.md" {
			roots = append(roots, filepath.Dir(filepath.Dir(path)))
		}
		return nil
	})
	if os.IsNotExist(err) {
		return nil, nil
	}
	return uniquePaths(roots), err
}

func globalSkillRoots(lookupEnv func(string) (string, bool)) ([]string, error) {
	if lookupEnv == nil {
		lookupEnv = os.LookupEnv
	}
	home, ok := lookupEnv("HOME")
	if !ok || home == "" {
		var err error
		home, err = os.UserHomeDir()
		if err != nil {
			return nil, err
		}
	}
	return []string{
		filepath.Join(home, ".claude", "skills"),
		filepath.Join(home, ".codex", "skills"),
	}, nil
}

func installedGlobalSkills(roots []string) ([]string, error) {
	skills, err := htskills.Catalog()
	if err != nil {
		return nil, err
	}
	known := make(map[string]struct{}, len(skills))
	aliases := make(map[string]struct{})
	for _, skill := range skills {
		known[skill.Name] = struct{}{}
		for _, alias := range skill.ReplacementAliases() {
			known[alias] = struct{}{}
			aliases[alias] = struct{}{}
		}
	}
	var installed []string
	for _, root := range roots {
		if _, err := os.ReadDir(root); err != nil {
			continue
		}
		for name := range known {
			path := filepath.Join(root, name, "SKILL.md")
			if _, err := os.Stat(path); err == nil {
				installed = append(installed, path)
			} else if !os.IsNotExist(err) {
				return nil, err
			}
		}
		entries, err := installedSkills(root)
		if err != nil {
			return nil, err
		}
		for _, skill := range entries {
			if _, ok := aliases[skill.Metadata.Name]; ok {
				installed = append(installed, filepath.Join(skill.Path, "SKILL.md"))
			}
		}
	}
	return uniquePaths(installed), nil
}

func isGlobalSkillRoot(destination string, globalRoots []string) bool {
	for _, root := range globalRoots {
		if samePath(destination, root) {
			return true
		}
	}
	return false
}

func samePath(left, right string) bool {
	left, err := filepath.Abs(left)
	if err != nil {
		return false
	}
	right, err = filepath.Abs(right)
	if err != nil {
		return false
	}
	return filepath.Clean(left) == filepath.Clean(right)
}

func uniquePaths(paths []string) []string {
	unique := make(map[string]string, len(paths))
	for _, path := range paths {
		absolute, err := filepath.Abs(path)
		if err != nil {
			absolute = filepath.Clean(path)
		}
		unique[absolute] = path
	}
	paths = paths[:0]
	for _, path := range unique {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	return paths
}

func globalSkillsError(paths []string) error {
	return fmt.Errorf("globally installed HyperTexting Skills are strongly discouraged: %v; uninstall them and retry with project-local skills", paths)
}
