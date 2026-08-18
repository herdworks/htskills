# htskills

`github.com/herdworks/htskills` is the versioned source package for the
portable HyperTexting Agent Skills bundle.

The package exposes the complete skill bundle through `htskills.FS`, rooted at
`htskills.Root`. The `htskills/readme` package exposes the templatable agent
onboarding README, and `htskills/cli` exposes the reusable Clic command tree.
Consumers can embed the files in an offline tool or copy an individual skill
directory for installation.

`htskills.Skill` defines the supported `SKILL.md` frontmatter schema. Use
`htskills.Catalog()` to enumerate validated bundled metadata, `htskills.Read`
to read one bundled `SKILL.md`, and `htskills.Parse` to validate frontmatter
from an installed skill.

Skill authors should edit the content under `skills/`; agent onboarding content
lives at `readme/README.md`. Hyperctl remounts the `htskills/cli` tree for its
agent-skill commands.
