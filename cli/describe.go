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
			Description: "Write the raw embedded SKILL.md bytes.",
			Usage: []string{
				"{{ .Command }} describe <name>",
			},
			Order: 3,
		},
		Spec: clic.CommandSpec{
			Key:       "describe",
			Help:      clic.DefaultHelp,
			Validator: expectedPositionalArguments(1),
			Handler: func(execution *clic.Execution) error {
				content, err := htskills.Read(execution.Command.State.Arguments[0])
				if err != nil {
					return err
				}
				_, err = fmt.Fprint(execution.Stdout, string(content))
				return err
			},
		},
	}
}
