package cli

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/herdworks/clic"
)

func TestSkillHelpUsage(t *testing.T) {
	mounted := SkillCommand()
	mounted.Meta.Name = "skill"
	mounted.Spec.Key = "skill"
	hyperctl := clic.CLI{Command: clic.Command{
		Spec:     clic.CommandSpec{Key: "hyperctl", Help: clic.DefaultHelp},
		Commands: clic.Commands{"skill": &mounted},
		Links:    clic.Links{{From: "readme", To: []string{"skill", "readme"}}},
	}}
	commands := []struct {
		name   string
		suffix string
	}{
		{"", " [command]"},
		{"readme", ""},
		{"ls", ""},
		{"describe", " <name>"},
		{"install", " <name> <path>"},
		{"update", " <name> <path>"},
		{"up", " [--force] [skills-directory]"},
	}
	for _, entrypoint := range []struct {
		cli    clic.CLI
		prefix string
		args   []string
	}{
		{CLI, "htskills", nil},
		{hyperctl, "hyperctl skill", []string{"skill"}},
	} {
		for _, command := range commands {
			path := strings.TrimSpace(entrypoint.prefix + " " + command.name)
			t.Run(path, func(t *testing.T) {
				args := append([]string{}, entrypoint.args...)
				if command.name != "" {
					args = append(args, command.name)
				}
				assertHelpUsage(t, entrypoint.cli, append(args, "--help"), path+command.suffix)
			})
		}
	}
	t.Run("hyperctl readme link", func(t *testing.T) {
		assertHelpUsage(t, hyperctl, []string{"readme", "--help"}, "hyperctl skill readme")
	})
}

func assertHelpUsage(t *testing.T, cli clic.CLI, args []string, want string) {
	t.Helper()
	var output bytes.Buffer
	result := cli.Run(clic.Execution{Context: context.Background(), Stdout: &output}, args)
	if result.Error != nil || result.ExitCode != 0 {
		t.Fatalf("help %v = %#v", args, result)
	}
	_, usage, ok := strings.Cut(output.String(), "USAGE\n")
	if !ok {
		t.Fatalf("help has no usage section: %q", output.String())
	}
	usage, _, _ = strings.Cut(usage, "\n\n")
	if got := strings.TrimSpace(usage); got != want {
		t.Fatalf("usage = %q, want %q", got, want)
	}
}
