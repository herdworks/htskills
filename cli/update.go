package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/herdworks/clic"
	"github.com/herdworks/htskills"
)

// SkillUpdateCommand updates a bundled skill when needed.
func SkillUpdateCommand() clic.Command {
	return clic.Command{
		Meta: clic.CommandMeta{
			Name:        "update",
			Description: "Update an installed skill when the bundled version is newer.",
			Usage: []string{
				"{{ .Command }} update <name> <path>",
			},
			Order: 5,
		},
		Spec: clic.CommandSpec{
			Key:       "update",
			Help:      clic.DefaultHelp,
			Validator: expectedPositionalArguments(2),
			Handler: func(execution *clic.Execution) error {
				written, err := updateSkill(execution.Command.State.Arguments[0], execution.Command.State.Arguments[1], commandPrefix(execution))
				if err != nil {
					return err
				}
				return writePaths(execution, written)
			},
		},
	}
}

func updateSkill(name, destination, command string) ([]string, error) {
	return updateSkillWithForce(name, destination, command, false)
}

func updateSkillWithForce(name, destination, command string, force bool) ([]string, error) {
	if _, err := htskills.Read(name); err != nil {
		return nil, err
	}
	info, err := os.Stat(destination)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("skill is not installed: %s", destination)
		}
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("skill installation is not a directory: %s", destination)
	}
	installed, err := os.ReadFile(filepath.Join(destination, "SKILL.md"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("skill is not installed: %s", destination)
		}
		return nil, err
	}
	installedMetadata, err := htskills.Parse(installed)
	if err != nil {
		return nil, fmt.Errorf("installed skill %q: %w", destination, err)
	}
	if installedMetadata.Name != name {
		return nil, fmt.Errorf("installed skill name %q does not match requested skill %q", installedMetadata.Name, name)
	}
	embedded, err := htskills.Read(name)
	if err != nil {
		return nil, err
	}
	embeddedMetadata, err := htskills.Parse(embedded)
	if err != nil {
		return nil, fmt.Errorf("bundled skill %q: %w", name, err)
	}
	if !force && installedMetadata.Metadata.Version == embeddedMetadata.Metadata.Version {
		return nil, nil
	}
	if !force && installedMetadata.Metadata.Version > embeddedMetadata.Metadata.Version {
		return nil, fmt.Errorf("skill %q has newer installed version %s (bundled version %s)", name, installedMetadata.Metadata.Version, embeddedMetadata.Metadata.Version)
	}
	if err = os.RemoveAll(filepath.Join(destination, "references")); err != nil {
		return nil, fmt.Errorf("remove installed skill references: %w", err)
	}
	return copyEmbeddedSkill(name, destination, command)
}
