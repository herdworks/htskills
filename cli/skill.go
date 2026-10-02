package cli

import (
	"fmt"
	"strings"

	"github.com/herdworks/clic"
)

// CLI is the standalone htskills command-line interface.
var CLI = clic.CLI{Command: SkillCommand()}

// SkillCommand returns the complete reusable HyperTexting Agent Skills command tree.
func SkillCommand() clic.Command {
	return clic.Command{
		Meta: clic.CommandMeta{
			Description: "Read, install, update, and synchronize bundled portable Agent Skills.",
			Usage: []string{
				"{{ .Command }} [command]",
			},
		},
		Spec: clic.CommandSpec{
			Key:  "htskills",
			Help: clic.DefaultHelp,
		},
		Commands: clic.MustCommandsFrom(
			ReadmeCommand,
			SkillListCommand,
			SkillDescribeCommand,
			SkillInstallCommand,
			SkillUpdateCommand,
			SkillUpCommand,
		),
	}
}

func expectedPositionalArguments(count int) func(*clic.Execution) error {
	return func(execution *clic.Execution) error {
		if execution.Command.PositionalArguments() == count {
			return nil
		}
		return positionalArgumentsError(execution, count)
	}
}

func positionalArgumentsError(execution *clic.Execution, expected int) error {
	return fmt.Errorf(
		"%s expected %d argument(s), got %d (%s)",
		execution.CommandPathString(),
		expected,
		execution.Command.PositionalArguments(),
		strings.Join(execution.Command.State.Arguments, ", "),
	)
}

func commandPrefix(execution *clic.Execution) string {
	if len(execution.CommandPath) > 1 {
		return strings.Join(execution.CommandPath[:len(execution.CommandPath)-1], " ")
	}
	return execution.RootCommandName()
}
