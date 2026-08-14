// Package htskills provides the versioned HyperTexting Agent Skills bundle.
package htskills

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

// Root is the directory containing the bundled skills in FS.
const Root = "skills"

// ReadmePath is the path to the bundled agent onboarding README in FS.
const ReadmePath = "readme/README.md"

// FS contains the complete portable HyperTexting Agent Skills bundle and agent
// onboarding README.
//
//go:embed all:skills all:readme
var FS embed.FS

// Skill is the supported YAML frontmatter schema for a bundled skill.
type Skill struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
}

// Skills is a collection of bundled skill metadata.
type Skills []Skill

// ColumnValues returns the values for a supported catalog column.
func (skills Skills) ColumnValues(field string) (values []string) {
	for _, skill := range skills {
		switch field {
		case "Name":
			values = append(values, skill.Name)
		case "Description":
			values = append(values, skill.Description)
		}
	}
	return values
}

// Parse reads and validates a skill's YAML frontmatter.
func Parse(content []byte) (skill Skill, err error) {
	const opening = "---\n"
	if !bytes.HasPrefix(content, []byte(opening)) {
		return skill, fmt.Errorf("skill is missing YAML frontmatter")
	}
	end := bytes.Index(content[len(opening):], []byte("\n---\n"))
	if end < 0 {
		return skill, fmt.Errorf("skill frontmatter is not terminated")
	}
	if err = yaml.Unmarshal(content[len(opening):len(opening)+end], &skill); err != nil {
		return skill, fmt.Errorf("invalid skill frontmatter: %w", err)
	}
	skill.Description = strings.TrimSpace(skill.Description)
	if skill.Name == "" {
		return skill, fmt.Errorf("skill frontmatter is missing name")
	}
	if skill.Description == "" {
		return skill, fmt.Errorf("skill frontmatter is missing description")
	}
	if skill.Version == "" {
		return skill, fmt.Errorf("skill %q is unversioned", skill.Name)
	}
	return skill, nil
}

// Catalog returns the metadata for every bundled skill in deterministic order.
func Catalog() (Skills, error) {
	entries, err := fs.ReadDir(FS, Root)
	if err != nil {
		return nil, err
	}
	skills := make(Skills, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		content, err := FS.ReadFile(path.Join(Root, entry.Name(), "SKILL.md"))
		if err != nil {
			return nil, fmt.Errorf("read bundled skill %q: %w", entry.Name(), err)
		}
		skill, err := Parse(content)
		if err != nil {
			return nil, fmt.Errorf("bundled skill %q: %w", entry.Name(), err)
		}
		if skill.Name != entry.Name() {
			return nil, fmt.Errorf("bundled skill directory %q does not match frontmatter name %q", entry.Name(), skill.Name)
		}
		skills = append(skills, skill)
	}
	return skills, nil
}

// Read returns a bundled skill's raw SKILL.md content by name.
func Read(name string) ([]byte, error) {
	if !fs.ValidPath(name) {
		return nil, fmt.Errorf("unknown skill %q", name)
	}
	skills, err := Catalog()
	if err != nil {
		return nil, err
	}
	for _, skill := range skills {
		if skill.Name == name {
			return FS.ReadFile(path.Join(Root, name, "SKILL.md"))
		}
	}
	return nil, fmt.Errorf("unknown skill %q", name)
}

// Readme returns the bundled agent onboarding README.
func Readme() ([]byte, error) {
	return FS.ReadFile(ReadmePath)
}
