// Package cli composes the reusable HyperTexting Agent Skills command-line interface.
package cli

import (
	"github.com/herdworks/clic"
	"github.com/herdworks/htskills/readme"
)

// ReadmeCommand writes the bundled agent onboarding README.
func ReadmeCommand() clic.Command {
	return clic.Command{
		Meta: clic.CommandMeta{
			Name:        "readme",
			Description: "Show onboarding context for a new agent.",
			Usage: []string{
				"{{ .Command }}",
			},
			Order: 1,
		},
		Spec: clic.CommandSpec{
			Key:       "readme",
			Help:      clic.DefaultHelp,
			Validator: expectedPositionalArguments(0),
			Handler: func(execution *clic.Execution) error {
				return writeReadme(execution)
			},
		},
	}
}

func writeReadme(execution *clic.Execution) error {
	template, err := readme.README()
	if err != nil {
		return err
	}
	return template.Execute(execution.Stdout, map[string]string{"Command": commandPrefix(execution)})
}
