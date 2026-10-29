package cli

import (
	"fmt"

	"github.com/herdworks/clic"
	"github.com/herdworks/htskills"
)

// SkillDescribeCommand writes a bundled SKILL.md file.
func SkillDescribeCommand() clic.Command {
	return clic.Command{
		Meta: clic.CommandMeta{
			Name:        "describe",
			Description: "Show one SKILL.md file.",
			Usage: []string{
				"{{ .Command }} <name>",
			},
			Order: 3,
		},
		Spec: clic.CommandSpec{
			Key:       "describe",
			Help:      clic.DefaultHelp,
			Validator: expectedPositionalArguments(1),
			Handler: func(execution *clic.Execution) error {
				content, err := htskills.RenderSkill(execution.Command.State.Arguments[0], map[string]string{"Command": commandPrefix(execution)})
				if err != nil {
					return err
				}
				_, err = fmt.Fprint(execution.Stdout, string(content))
				return err
			},
		},
	}
}
