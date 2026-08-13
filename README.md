# htskills

`github.com/herdworks/htskills` is the versioned source package for the
portable HyperTexting Agent Skills bundle.

The package exposes the complete bundle through `htskills.FS`, rooted at
`htskills.Root`. Consumers can embed the files in an offline tool or copy an
individual skill directory for installation.

`htskills.Skill` defines the supported `SKILL.md` frontmatter schema. Use
`htskills.Catalog()` to enumerate validated bundled metadata, `htskills.Read`
to read one bundled `SKILL.md`, and `htskills.Parse` to validate frontmatter
from an installed skill.

Skill authors should edit the content under `skills/`. Hyperctl consumes this
package to ship the same bundle offline.
