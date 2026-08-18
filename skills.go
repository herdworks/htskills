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

// FS contains the complete portable HyperTexting Agent Skills bundle.
//
//go:embed all:skills
var FS embed.FS

// Skill is the supported YAML frontmatter schema for a bundled skill.
type Skill struct {
	Name          string        `yaml:"name"`
	Description   string        `yaml:"description"`
	Compatibility string        `yaml:"compatibility"`
	Metadata      SkillMetadata `yaml:"metadata"`
}

// SkillMetadata is the supported metadata block in a skill's YAML frontmatter.
type SkillMetadata struct {
	Version  string `yaml:"version"`
	Replaces string `yaml:"replaces"`
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
	skill.Compatibility = strings.TrimSpace(skill.Compatibility)
	skill.Metadata.Version = strings.TrimSpace(skill.Metadata.Version)
	skill.Metadata.Replaces = strings.TrimSpace(skill.Metadata.Replaces)
	if skill.Name == "" {
		return skill, fmt.Errorf("skill frontmatter is missing name")
	}
	if skill.Description == "" {
		return skill, fmt.Errorf("skill frontmatter is missing description")
	}
	aliases := make(map[string]struct{})
	for _, alias := range skill.ReplacementAliases() {
		if !validSkillName(alias) {
			return skill, fmt.Errorf("skill %q replaces invalid skill name %q", skill.Name, alias)
		}
		if alias == skill.Name {
			return skill, fmt.Errorf("skill %q replaces itself", skill.Name)
		}
		if _, ok := aliases[alias]; ok {
			return skill, fmt.Errorf("skill %q replaces %q more than once", skill.Name, alias)
		}
		aliases[alias] = struct{}{}
	}
	return skill, nil
}

// ReplacementAliases returns the trimmed replacement names declared in metadata.
func (skill Skill) ReplacementAliases() []string {
	if skill.Metadata.Replaces == "" {
		return nil
	}
	aliases := strings.Split(skill.Metadata.Replaces, ",")
	for index, alias := range aliases {
		aliases[index] = strings.TrimSpace(alias)
	}
	return aliases
}

func validSkillName(name string) bool {
	return name != "." && fs.ValidPath(name) && path.Base(name) == name
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
		if skill.Metadata.Version == "" {
			return nil, fmt.Errorf("bundled skill %q is unversioned", skill.Name)
		}
		if skill.Compatibility == "" {
			return nil, fmt.Errorf("bundled skill %q has no compatibility declaration", skill.Name)
		}
		skills = append(skills, skill)
	}
	if err = validateCatalog(skills); err != nil {
		return nil, err
	}
	return skills, nil
}

func validateCatalog(skills Skills) error {
	canonical := make(map[string]struct{}, len(skills))
	for _, skill := range skills {
		canonical[skill.Name] = struct{}{}
	}
	replacedBy := make(map[string]string)
	for _, skill := range skills {
		for _, alias := range skill.ReplacementAliases() {
			if _, ok := canonical[alias]; ok {
				return fmt.Errorf("skill %q replaces bundled skill %q", skill.Name, alias)
			}
			if owner, ok := replacedBy[alias]; ok {
				return fmt.Errorf("skills %q and %q both replace %q", owner, skill.Name, alias)
			}
			replacedBy[alias] = skill.Name
		}
	}
	return nil
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
