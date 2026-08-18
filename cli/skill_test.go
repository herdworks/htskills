package cli

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/herdworks/clic"
	"github.com/herdworks/htskills"
)

func runCLI(args ...string) (bytes.Buffer, clic.Result) {
	return runCLIAt(".", "", args...)
}

func runCLIAt(workspace, home string, args ...string) (bytes.Buffer, clic.Result) {
	var output bytes.Buffer
	if home == "" {
		home = filepath.Join(workspace, ".test-home")
	}
	lookupEnv := func(key string) (string, bool) { return home, key == "HOME" }
	result := CLI.Run(clic.Execution{
		Context:          context.Background(),
		Stdout:           &output,
		WorkingDirectory: workspace,
		LookupEnv:        lookupEnv,
	}, args)
	return output, result
}

func TestSkillCommandRegistersCompleteStandaloneTree(t *testing.T) {
	command := SkillCommand()
	if command.Name() != "htskills" {
		t.Fatalf("command name = %q, want htskills", command.Name())
	}
	for _, name := range []string{"readme", "ls", "describe", "install", "update", "up"} {
		if _, ok := command.SubCommand(name); !ok {
			t.Fatalf("command does not register %q", name)
		}
	}
}

func TestSkillUpRequiresSkillsDirectory(t *testing.T) {
	output, result := runCLIAt(t.TempDir(), "", "up")
	if result.ExitCode != 1 || result.Error == nil {
		t.Fatalf("htskills up = %#v, want validation error", result)
	}
	if output.Len() != 0 {
		t.Fatalf("htskills up wrote stdout %q", output.String())
	}
	if got, want := result.Error.Error(), `htskills up requires a skills directory; retry with "htskills up <skills-directory>" using your agent's project-local or global skills directory`; got != want {
		t.Fatalf("htskills up error = %q, want %q", got, want)
	}
}

func TestSkillUpSynchronizesSkillsAndReportsResults(t *testing.T) {
	directory := t.TempDir()
	name := "hypertexting-assets"
	content, err := htskills.Read(name)
	if err != nil {
		t.Fatal(err)
	}
	metadata, err := htskills.Parse(content)
	if err != nil {
		t.Fatal(err)
	}
	installed := strings.Replace(string(content), `version: "`+metadata.Metadata.Version+"\"", `version: "2026-01"`, 1)
	installedPath := filepath.Join(directory, name, "SKILL.md")
	if err = os.MkdirAll(filepath.Dir(installedPath), 0755); err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(installedPath, []byte(installed), 0644); err != nil {
		t.Fatal(err)
	}

	output, result := runCLIAt(filepath.Dir(directory), "", "up", directory)
	if result.Error != nil {
		t.Fatalf("htskills up = %#v", result)
	}
	if !strings.Contains(output.String(), `skill "`+installedPath+`" updated to "`+metadata.Metadata.Version+`"`) {
		t.Fatalf("htskills up output does not report updated skill: %q", output.String())
	}
	if strings.Contains(output.String(), "HyperTexting Agent Skills") {
		t.Fatalf("htskills up printed the onboarding README: %q", output.String())
	}
	for _, skill := range bundledSkills(t) {
		installed, err := os.ReadFile(filepath.Join(directory, skill.Name, "SKILL.md"))
		if err != nil {
			t.Fatalf("read installed %s: %v", skill.Name, err)
		}
		embedded, err := htskills.Read(skill.Name)
		if err != nil {
			t.Fatal(err)
		}
		if string(installed) != string(embedded) {
			t.Fatalf("installed %s differs from bundle", skill.Name)
		}
	}
}

func TestSkillUpReportsCurrentSkills(t *testing.T) {
	directory := t.TempDir()
	if _, result := runCLIAt(filepath.Dir(directory), "", "up", directory); result.Error != nil {
		t.Fatalf("initial htskills up = %#v", result)
	}

	output, result := runCLIAt(filepath.Dir(directory), "", "up", directory)
	if result.Error != nil {
		t.Fatalf("second htskills up = %#v", result)
	}
	for _, skill := range bundledSkills(t) {
		path := filepath.Join(directory, skill.Name, "SKILL.md")
		if !strings.Contains(output.String(), `skill "`+path+`" current`) {
			t.Fatalf("htskills up output does not report current skill %q: %q", path, output.String())
		}
	}
}

func TestSkillUpUpdatesUnversionedSkill(t *testing.T) {
	directory := t.TempDir()
	name := "hypertexting-assets"
	content, err := htskills.Read(name)
	if err != nil {
		t.Fatal(err)
	}
	metadata, err := htskills.Parse(content)
	if err != nil {
		t.Fatal(err)
	}
	unversioned := strings.Replace(string(content), "metadata:\n  version: \""+metadata.Metadata.Version+"\"\n", "", 1)
	installedPath := filepath.Join(directory, name, "SKILL.md")
	if err = os.MkdirAll(filepath.Dir(installedPath), 0755); err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(installedPath, []byte(unversioned), 0644); err != nil {
		t.Fatal(err)
	}

	output, result := runCLIAt(filepath.Dir(directory), "", "up", directory)
	if result.Error != nil {
		t.Fatalf("htskills up = %#v", result)
	}
	if !strings.Contains(output.String(), `skill "`+installedPath+`" updated to "`+metadata.Metadata.Version+`"`) {
		t.Fatalf("htskills up output does not report unversioned skill update: %q", output.String())
	}
}

func TestSkillUpFailsWithoutWritingOutput(t *testing.T) {
	directory := t.TempDir()
	name := "hypertexting-assets"
	content, err := htskills.Read(name)
	if err != nil {
		t.Fatal(err)
	}
	metadata, err := htskills.Parse(content)
	if err != nil {
		t.Fatal(err)
	}
	newer := strings.Replace(string(content), `version: "`+metadata.Metadata.Version+"\"", `version: "9999-01"`, 1)
	path := filepath.Join(directory, name)
	if err = os.MkdirAll(path, 0755); err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(filepath.Join(path, "SKILL.md"), []byte(newer), 0644); err != nil {
		t.Fatal(err)
	}

	output, result := runCLIAt(filepath.Dir(directory), "", "up", directory)
	if result.ExitCode != 2 || result.Error == nil {
		t.Fatalf("htskills up = %#v, want handler error", result)
	}
	if output.Len() != 0 {
		t.Fatalf("failed htskills up wrote stdout %q", output.String())
	}
	if !strings.Contains(result.Error.Error(), "newer installed version") {
		t.Fatalf("htskills up error = %q", result.Error)
	}
}

func TestSkillUpdateIsIdempotentForEqualVersions(t *testing.T) {
	name := "hypertexting-assets"
	destination := filepath.Join(t.TempDir(), name)
	if _, err := installSkill(name, destination); err != nil {
		t.Fatal(err)
	}
	written, err := updateSkill(name, destination)
	if err != nil {
		t.Fatal(err)
	}
	if len(written) != 0 {
		t.Fatalf("equal-version update wrote %v", written)
	}
}

func TestSkillUpRejectsNonDirectory(t *testing.T) {
	path := filepath.Join(t.TempDir(), "skills")
	if err := os.WriteFile(path, []byte("not a directory"), 0644); err != nil {
		t.Fatal(err)
	}
	output, result := runCLIAt(t.TempDir(), "", "up", path)
	if result.ExitCode != 2 || result.Error == nil {
		t.Fatalf("htskills up = %#v, want handler error", result)
	}
	if output.Len() != 0 || !strings.Contains(result.Error.Error(), "not a directory") {
		t.Fatalf("htskills up = output %q, error %v", output.String(), result.Error)
	}
}

func TestSkillUpDiscoversEverySkillRootAndSuppliedDirectory(t *testing.T) {
	workspace := t.TempDir()
	roots := []string{
		filepath.Join(workspace, ".agents", "skills"),
		filepath.Join(workspace, ".claude", "skills"),
		filepath.Join(workspace, ".codex", "skills"),
		t.TempDir(),
	}
	for _, root := range roots[:3] {
		writeSkillMarker(t, root, "custom")
	}
	writeSkillMarker(t, roots[0], "custom-two")

	output, result := runCLIAt(workspace, "", "up", roots[3])
	if result.Error != nil {
		t.Fatalf("htskills up = %#v", result)
	}
	for _, root := range roots {
		if !strings.Contains(output.String(), `skills "`+root+`" discovered`) {
			t.Fatalf("htskills up output does not report discovered root %q: %q", root, output.String())
		}
		path := filepath.Join(root, "hypertexting", "SKILL.md")
		if !strings.Contains(output.String(), `skill "`+path+`" installed`) {
			t.Fatalf("htskills up output does not report installed skill %q: %q", path, output.String())
		}
	}
	for _, root := range roots {
		for _, skill := range bundledSkills(t) {
			if _, err := os.Stat(filepath.Join(root, skill.Name, "SKILL.md")); err != nil {
				t.Fatalf("%s was not synchronized into %s: %v", skill.Name, root, err)
			}
		}
	}
}

func TestDiscoverSkillRootsDeduplicatesSkillDirectories(t *testing.T) {
	workspace := t.TempDir()
	first := filepath.Join(workspace, ".agents", "skills")
	second := filepath.Join(workspace, ".claude", "skills")
	writeSkillMarker(t, first, "one")
	writeSkillMarker(t, first, "two")
	writeSkillMarker(t, second, "three")

	roots, err := discoverSkillRoots(workspace)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := roots, []string{first, second}; !slices.Equal(got, want) {
		t.Fatalf("discovered roots = %v, want %v", got, want)
	}
}

func TestSkillUpRejectsGlobalDirectoriesAndInstalledGlobalSkills(t *testing.T) {
	workspace := t.TempDir()
	home := t.TempDir()
	global := filepath.Join(home, ".codex", "skills")
	local := filepath.Join(workspace, ".agents", "skills")
	writeSkillMarker(t, local, "custom")

	output, result := runCLIAt(workspace, home, "up", global)
	if result.ExitCode != 1 || result.Error == nil || !strings.Contains(result.Error.Error(), "strongly discouraged") {
		t.Fatalf("htskills up global directory = %#v", result)
	}
	if output.Len() != 0 {
		t.Fatalf("htskills up global directory wrote stdout %q", output.String())
	}
	if _, err := os.Stat(filepath.Join(local, "hypertexting", "SKILL.md")); !os.IsNotExist(err) {
		t.Fatalf("local root changed after rejected global argument: %v", err)
	}

	if err := os.MkdirAll(filepath.Join(global, "hypertexting"), 0755); err != nil {
		t.Fatal(err)
	}
	content, err := htskills.Read("hypertexting")
	if err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(filepath.Join(global, "hypertexting", "SKILL.md"), content, 0644); err != nil {
		t.Fatal(err)
	}
	output, result = runCLIAt(workspace, home, "up")
	if result.ExitCode != 1 || result.Error == nil || !strings.Contains(result.Error.Error(), "strongly discouraged") || !strings.Contains(result.Error.Error(), filepath.Join(global, "hypertexting", "SKILL.md")) {
		t.Fatalf("htskills up installed global skill = %#v", result)
	}
	if output.Len() != 0 {
		t.Fatalf("htskills up installed global skill wrote stdout %q", output.String())
	}
}

func TestInstallAndUpdateAcceptGlobalSkillPaths(t *testing.T) {
	workspace := t.TempDir()
	home := t.TempDir()
	destination := filepath.Join(home, ".claude", "skills", "hypertexting-assets")
	if err := os.MkdirAll(filepath.Dir(destination), 0755); err != nil {
		t.Fatal(err)
	}
	_, result := runCLIAt(workspace, home, "install", "hypertexting-assets", destination)
	if result.Error != nil {
		t.Fatalf("htskills install global path = %#v", result)
	}
	content, err := htskills.Read("hypertexting-assets")
	if err != nil {
		t.Fatal(err)
	}
	metadata, err := htskills.Parse(content)
	if err != nil {
		t.Fatal(err)
	}
	older := strings.Replace(string(content), `version: "`+metadata.Metadata.Version+`"`, `version: "2026-01"`, 1)
	if err = os.WriteFile(filepath.Join(destination, "SKILL.md"), []byte(older), 0644); err != nil {
		t.Fatal(err)
	}
	_, result = runCLIAt(workspace, home, "update", "hypertexting-assets", destination)
	if result.Error != nil {
		t.Fatalf("htskills update global path = %#v", result)
	}
}

func TestSkillUpIgnoresUnavailableGlobalRoots(t *testing.T) {
	workspace := t.TempDir()
	home := t.TempDir()
	if err := os.MkdirAll(filepath.Join(home, ".claude"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(home, ".claude", "skills"), []byte("not a directory"), 0644); err != nil {
		t.Fatal(err)
	}
	root := filepath.Join(workspace, ".agents", "skills")
	writeSkillMarker(t, root, "custom")

	_, result := runCLIAt(workspace, home, "up")
	if result.Error != nil {
		t.Fatalf("htskills up with unavailable global root = %#v", result)
	}
}

func writeSkillMarker(t *testing.T, root, name string) {
	t.Helper()
	path := filepath.Join(root, name, "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("marker"), 0644); err != nil {
		t.Fatal(err)
	}
}

func bundledSkills(t *testing.T) []htskills.Skill {
	t.Helper()
	skills, err := htskills.Catalog()
	if err != nil {
		t.Fatal(err)
	}
	return skills
}
