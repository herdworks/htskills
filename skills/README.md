# Embedded skill style guide

Use this guide when authoring or updating the embedded HyperTexting skills.

## Required shape

- Use YAML frontmatter with the skill `name`, a concise agent-facing `description`, and `version: 1`.
- Start with a level-one title.
- Include `## Overview`, `## Workflow`, `## Instructions`, `## Tools`, and `## References` in every `SKILL.md`.
- Add `## Known issues` only when a documented, acceptable behavior can otherwise cause needless investigation.

## Instruction style

- Write instructions as short, imperative bullets.
- Use numbered lists only for ordered workflows.
- Group related instruction bullets under descriptive level-three headings.
- Keep prose for a brief overview or rationale that cannot be expressed as an action.
- Put one action or constraint in each bullet; do not hide several rules in a long paragraph.
- Use active voice and name the preferred command or directive directly.

## References and examples

- Keep `SKILL.md` decision-oriented; move detailed syntax, contracts, and extended examples into `references/` files.
- Link each reference directly from its owning `SKILL.md` and tell agents to consult it as needed.
- Use external documentation only as a fallback when bundled references do not answer the question.
- Use Markdown for reference files without YAML frontmatter.
- Shape reference files as `# Title`, then `## Overview`, `## Instructions`, and `## Examples`; add `## Specification` only for a fallback link.
- Order `ht-*` attributes before ordinary HTML attributes in every example.
