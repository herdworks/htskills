package cli

import (
	"github.com/herdworks/clic"
	"github.com/herdworks/htskills"
)

// SkillListCommand lists bundled skills.
func SkillListCommand() clic.Command {
	return clic.Command{
		Meta: clic.CommandMeta{
			Name:        "ls",
			Description: "List bundled offline skills.",
			Usage: []string{
				"{{ .Command }}",
			},
			Order: 2,
		},
		Spec: clic.CommandSpec{
			Key:  "ls",
			Help: clic.DefaultHelp,
			Handler: func(execution *clic.Execution) error {
				skills, err := htskills.Catalog()
				if err != nil {
					return err
				}
				return clic.WriteList(execution.Stdout, clic.ColumnarFormat, skills,
					clic.Field{Key: "Name", Header: "Name"},
					clic.Field{Key: "Description", Header: "Description"},
				)
			},
		},
	}
}
