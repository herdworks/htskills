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
		"hypertexting":                   "HyperTexting website orchestration skill.",
		"hypertexting-assets":            "HyperTexting asset management skill.",
		"hypertexting-css":               "General CSS authoring guidelines.",
		"hypertexting-data":              "HyperTexting data management skill.",
		"hypertexting-debugging":         "HyperTexting debugging skill.",
		"hypertexting-html":              "General HTML authoring guidelines.",
		"hypertexting-javascript":        "General JavaScript authoring guidelines.",
		"hypertexting-layouts":           "HyperTexting HTML templating skill.",
		"hypertexting-plugins":           "HyperTexting plugins skill.",
		"hypertexting-theme-development": "Develop new HyperTexting themes from visual design assets.",
		"hypertexting-themes":            "HyperTexting theme management skill.",
	}
	if len(skills) != len(want) {
		t.Fatalf("Catalog returned %d skills, want %d", len(skills), len(want))
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

func TestReadmeReturnsBundledOnboardingContent(t *testing.T) {
	content, err := Readme()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), "HyperTexting is a content management system (CMS) and static site generator (SSG)") {
		t.Fatalf("Readme() returned incomplete onboarding content: %q", content)
	}
	if !strings.HasSuffix(string(content), "\n") {
		t.Fatal("Readme() does not end with a newline")
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
