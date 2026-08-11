package htskills

import (
	"strings"
	"testing"
)

func TestSkillsExposeBundledFrontmatter(t *testing.T) {
	skills, err := Skills()
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"hypertexting":            "HyperTexting orchestration skill.",
		"hypertexting-assets":     "Website asset management skill.",
		"hypertexting-css":        "General CSS authoring guidelines.",
		"hypertexting-data":       "Website data management skill.",
		"hypertexting-debugging":  "Website and template debugging skill.",
		"hypertexting-html":       "General HTML authoring guidelines.",
		"hypertexting-javascript": "General JavaScript authoring guidelines.",
		"hypertexting-layouts":    "HTML templating skill.",
		"hypertexting-plugins":    "Website plugins skill.",
		"hypertexting-themes":     "Website theme management skill.",
	}
	if len(skills) != len(want) {
		t.Fatalf("Skills returned %d skills, want %d", len(skills), len(want))
	}
	for _, skill := range skills {
		if skill.Description != want[skill.Name] {
			t.Fatalf("%s description = %q, want %q", skill.Name, skill.Description, want[skill.Name])
		}
		if skill.Version == "" {
			t.Fatalf("%s is unversioned", skill.Name)
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

func TestParseValidatesSupportedFrontmatter(t *testing.T) {
	content := []byte("---\nname: example\ndescription: |\n  Example skill.\nversion: 2\n---\n# Example\n")
	skill, err := Parse(content)
	if err != nil {
		t.Fatal(err)
	}
	if skill.Name != "example" || skill.Description != "Example skill." || skill.Version != "2" {
		t.Fatalf("Parse() = %#v", skill)
	}
	for _, content := range [][]byte{
		[]byte("# No frontmatter\n"),
		[]byte("---\nname: example\nversion: 1\n---\n"),
		[]byte("---\nname: example\ndescription: Example\nversion: \"\"\n---\n"),
	} {
		if _, err := Parse(content); err == nil {
			t.Fatalf("Parse(%q) succeeded", content)
		}
	}
}
