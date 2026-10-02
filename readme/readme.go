// Package readme provides the templatable HyperTexting Agent Skills onboarding README.
package readme

import (
	_ "embed"
	"text/template"

	"github.com/herdworks/htskills"
)

// content is the source for the templatable onboarding README.
//
//go:embed README.md
var content string

// README returns a fresh parsed onboarding README template.
func README() (*template.Template, error) {
	return htskills.ParseTemplate("README.md", []byte(content))
}
