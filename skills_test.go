package htskills

import (
	"strings"
	"testing"
)

func TestCatalogExposesBundledFrontmatter(t *testing.T) {
	skills, err := Catalog()
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"hypertexting":            "HyperTexting website orchestration skill.",
		"hypertexting-assets":     "HyperTexting asset management skill.",
		"hypertexting-css":        "General CSS authoring guidelines.",
		"hypertexting-data":       "HyperTexting data management skill.",
		"hypertexting-debugging":  "HyperTexting debugging skill.",
		"hypertexting-html":       "General HTML authoring guidelines.",
		"hypertexting-javascript": "General JavaScript authoring guidelines.",
		"hypertexting-layouts":    "HyperTexting HTML templating skill.",
		"hypertexting-plugins":    "HyperTexting plugins skill.",
		"hypertexting-themes":     "HyperTexting theme management skill.",
	}
	if len(skills) != len(want) {
		t.Fatalf("Catalog returned %d skills, want %d", len(skills), len(want))
	}
	for _, skill := range skills {
		if skill.Description != want[skill.Name] {
			t.Fatalf("%s description = %q, want %q", skill.Name, skill.Description, want[skill.Name])
		}
		if skill.Metadata.Version == "" {
			t.Fatalf("%s is unversioned", skill.Name)
		}
		if skill.Compatibility != "Requires hyperctl." {
			t.Fatalf("%s compatibility = %q, want Requires hyperctl.", skill.Name, skill.Compatibility)
		}
		content, err := Read(skill.Name)
		if err != nil {
			t.Fatalf("Read(%q): %v", skill.Name, err)
		}
		if !strings.HasPrefix(string(content), "---\nname: "+skill.Name+"\n") {
			t.Fatalf("Read(%q) returned invalid frontmatter", skill.Name)
		}
	}
}

func TestSkillsColumnValues(t *testing.T) {
	skills := Skills{
		{Name: "one", Description: "First skill."},
		{Name: "two", Description: "Second skill."},
	}
	if got, want := strings.Join(skills.ColumnValues("Name"), ","), "one,two"; got != want {
		t.Fatalf("ColumnValues(Name) = %q, want %q", got, want)
	}
	if got, want := strings.Join(skills.ColumnValues("Description"), ","), "First skill.,Second skill."; got != want {
		t.Fatalf("ColumnValues(Description) = %q, want %q", got, want)
	}
	if got := skills.ColumnValues("Version"); got != nil {
		t.Fatalf("ColumnValues(Version) = %v, want nil", got)
	}
}

func TestParseValidatesSupportedFrontmatter(t *testing.T) {
	content := []byte("---\nname: example\ndescription: |\n  Example skill.\ncompatibility: Requires hyperctl.\nmetadata:\n  version: 2\n  replaces: old-example, another-old-example\n---\n# Example\n")
	skill, err := Parse(content)
	if err != nil {
		t.Fatal(err)
	}
	if skill.Name != "example" || skill.Description != "Example skill." || skill.Compatibility != "Requires hyperctl." || skill.Metadata.Version != "2" {
		t.Fatalf("Parse() = %#v", skill)
	}
	if got, want := strings.Join(skill.ReplacementAliases(), ","), "old-example,another-old-example"; got != want {
		t.Fatalf("Parse() replacement aliases = %q, want %q", got, want)
	}
	for _, content := range [][]byte{
		[]byte("# No frontmatter\n"),
		[]byte("---\nname: example\nmetadata:\n  version: 1\n---\n"),
		[]byte("---\nname: example\ndescription: Example\nmetadata:\n  version: 1\n  replaces: example\n---\n"),
		[]byte("---\nname: example\ndescription: Example\nmetadata:\n  version: 1\n  replaces: ../old\n---\n"),
		[]byte("---\nname: example\ndescription: Example\nmetadata:\n  version: 1\n  replaces: old,,older\n---\n"),
		[]byte("---\nname: example\ndescription: Example\nmetadata:\n  version: 1\n  replaces: [old]\n---\n"),
	} {
		if _, err := Parse(content); err == nil {
			t.Fatalf("Parse(%q) succeeded", content)
		}
	}
}

func TestParseAllowsUnversionedInstalledSkills(t *testing.T) {
	skill, err := Parse([]byte("---\nname: example\ndescription: Example skill.\n---\n# Example\n"))
	if err != nil {
		t.Fatal(err)
	}
	if skill.Metadata.Version != "" {
		t.Fatalf("unversioned skill version = %q, want empty", skill.Metadata.Version)
	}
}

func TestParseIgnoresLegacyTopLevelVersion(t *testing.T) {
	skill, err := Parse([]byte("---\nname: example\ndescription: Example skill.\nversion: 1\n---\n# Example\n"))
	if err != nil {
		t.Fatal(err)
	}
	if skill.Metadata.Version != "" {
		t.Fatalf("legacy version = %q, want empty metadata version", skill.Metadata.Version)
	}
}

func TestValidateCatalogRejectsReplacementCollisions(t *testing.T) {
	for _, skills := range []Skills{
		{{Name: "one", Metadata: SkillMetadata{Replaces: "two"}}, {Name: "two"}},
		{{Name: "one", Metadata: SkillMetadata{Replaces: "old"}}, {Name: "two", Metadata: SkillMetadata{Replaces: "old"}}},
	} {
		if err := validateCatalog(skills); err == nil {
			t.Fatalf("validateCatalog(%#v) succeeded", skills)
		}
	}
}
