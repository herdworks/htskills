---
name: hypertexting-layouts
description: HyperTexting HTML templating skill.
compatibility: Requires hyperctl.
metadata:
  version: "2026-08-26T00:00:00Z"
---

# HyperTexting layouts

## Overview

Compose complete HTML documents from reusable fragments. Let the source design determine whether authoring begins with a page layout or component fragment; both paths must converge on layouts composed from fragments with clear data contracts.

## Workflow

1. Let the source design dictate the authoring workflow.
2. Build the page or component until its data contract and reuse opportunities are clear.
3. Extract reusable fragments and define their data contracts.
4. Select directives that match the resulting layout and fragment contracts.
5. Render the layout in a site context and verify the output.

## Instructions

### Compose layouts

- Model websites as distinct page types, with one layout per type.
- Use the source design to choose an entrypoint. For full-page mockups, begin with a monolithic layout; for component-level designs, begin with fragments.
- Follow the chosen entrypoint until the data contract and reuse opportunities are clear, then converge on the required outcome: reusable fragments with explicit data contracts, composed into layouts with `ht-include`.
- Do not split a layout merely to anticipate reuse.

### Select directives

- Use `ht-include` for fragments, `ht-if` for conditions, `ht-apply` / `ht-with` / `ht-each` for attribute and content templating, and `ht-block` for content-addressable templating.
- Use `ht-apply` to render variables in a target element's attributes and direct text nodes. Use `ht-with` or `ht-each` to bind values and make them available throughout the target and its descendants as they are processed. Define the fragment's expected data before selecting a directive; consult [template-data.md](references/template-data.md) as needed.
- Use `ht-debug`, `ht-pipe`, `ht-attrs`, and `ht-base` only when the simpler directives cannot express the layout.
- Do not reach for JavaScript components until the HTML model is exhausted.
- Use `ht-apply` with template variables instead of legacy `ht-content`, `ht-attr`, `ht-attrs`, `ht-param`, and `ht-query` patterns.

### Structure layout documents

- Organize `<body>` into top-level `<header>`, `<main>`, and `<footer>`.
- Build page regions as `<section>` elements and use `<article>` for article content.
- Once a section has a clear data contract or reuse opportunity, make it a fragment; do not retain established reusable sections in one growing layout file.
- Keep a section inline while its design or data contract is still evolving, or when it remains page-specific and simple.
- When two or more top-level section fragments reuse a component, make that component a fragment too.
- Compare extracted section fragments for repeated HTML blocks and component styles; extract those into smaller shared fragments.
- Use `hypertexting-html` to choose semantic elements, custom elements, component attributes, and styling hooks.
- Colocate a fragment’s component styles in that fragment’s HTML file.
- Include the required `style#components` collector and later `style#layout` block in every layout head; consult `hyperctl agent skill describe hypertexting-css` for their contract and cascade guidance.

### Verify rendered output

- Use `hyperctl theme render --theme <theme.json> --data <fixture>.json --fragment <fragment>` to develop a rootless component against its fixture data. It renders only the fragment nodes and does not require page data or inject built-ins.
- Test fragment composition inside a site directory with a throwaway page.
- Validate rendered output as well as build logs, especially when debugging bindings.

## Tools

- `hyperctl theme render` renders either a configured theme layout (`--layout`) or rootless theme fragment (`--fragment`) against fixture data in any supported template-data format. It does not construct HyperTexting site or page data.

## References

- Consult [template-data.md](references/template-data.md) as needed when deciding what template data is visible or how a fragment receives it.
- Consult [directives.md](references/directives.md), [variables.md](references/variables.md), or [expressions.md](references/expressions.md) as needed for directive terminology, template variables, and expression tokens.
- Consult [ht-include.md](references/ht-include.md), [ht-if.md](references/ht-if.md), [ht-apply.md](references/ht-apply.md), [ht-with.md](references/ht-with.md), [ht-each.md](references/ht-each.md), or [ht-block.md](references/ht-block.md) as needed for the core layout directives.
- Consult [ht-debug.md](references/ht-debug.md), [ht-attrs.md](references/ht-attrs.md), [ht-base.md](references/ht-base.md), or [ht-pipe.md](references/ht-pipe.md) as needed for specialized layout directives.
- Consult `hyperctl agent skill describe hypertexting-html` as needed for semantic element, custom-element, attribute, and class decisions.
- Consult the [templating](https://hypertemplates.net/docs/#introduction-to-templating), [layouts](https://hypertemplates.net/docs/reference/core/layouts/), [fragments](https://hypertemplates.net/docs/reference/core/fragments/), or [directives](https://hypertemplates.net/docs/reference/core/directives/) documentation only if the bundled references do not answer the question.
