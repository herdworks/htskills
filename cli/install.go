package cli

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/herdworks/clic"
	"github.com/herdworks/htskills"
)

// SkillInstallCommand copies a bundled skill to a new destination.
func SkillInstallCommand() clic.Command {
	return clic.Command{
		Meta: clic.CommandMeta{
			Name:        "install",
			Description: "Copy a complete skill directory to a new destination.",
			Usage: []string{
				"{{ .Command }} <name> <path>",
			},
			Order: 4,
		},
		Spec: clic.CommandSpec{
			Key:       "install",
			Help:      clic.DefaultHelp,
			Validator: expectedPositionalArguments(2),
			Handler: func(execution *clic.Execution) error {
				written, err := installSkill(execution.Command.State.Arguments[0], execution.Command.State.Arguments[1], commandPrefix(execution))
				if err != nil {
					return err
				}
				return writePaths(execution, written)
			},
		},
	}
}

func writePaths(execution *clic.Execution, paths []string) error {
	for _, name := range paths {
		if _, err := fmt.Fprintln(execution.Stdout, name); err != nil {
			return err
		}
	}
	return nil
}

func installSkill(name, destination, command string) ([]string, error) {
	if _, err := htskills.Read(name); err != nil {
		return nil, err
	}
	if _, err := os.Lstat(destination); err == nil {
		return nil, fmt.Errorf("destination already exists: %s", destination)
	} else if !os.IsNotExist(err) {
		return nil, err
	}
	if err := os.Mkdir(destination, 0755); err != nil {
		return nil, err
	}
	return copyEmbeddedSkill(name, destination, command)
}

func copyEmbeddedSkill(name, destination, command string) ([]string, error) {
	root := htskills.Root + "/" + name
	written := []string{}
	err := fs.WalkDir(htskills.FS, root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relative := strings.TrimPrefix(path, root)
		if relative == "" {
			return nil
		}
		output := filepath.Join(destination, filepath.FromSlash(relative))
		if entry.IsDir() {
			return os.MkdirAll(output, 0755)
		}
		content, err := htskills.FS.ReadFile(path)
		if err != nil {
			return err
		}
		filename := strings.TrimPrefix(relative, "/")
		if filename == "SKILL.md" || strings.HasPrefix(filename, "references/") {
			content, err = htskills.RenderSkillFile(name, filename, map[string]string{"Command": command})
			if err != nil {
				return err
			}
		}
		if err = os.WriteFile(output, content, 0644); err != nil {
			return err
		}
		written = append(written, output)
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(written)
	return written, nil
}
