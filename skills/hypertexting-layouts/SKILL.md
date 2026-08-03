---
name: hypertexting-layouts
description: HTML templating skill.
version: 1
---

# HyperTexting layouts

## Overview

Compose complete HTML documents from reusable fragments. Define each fragment's data contract before choosing the directive that renders it.

## Workflow

1. Model the document structure and extract reusable fragments.
2. Define the data each fragment needs and choose its scope.
3. Select directives that match the fragment and data contracts.
4. Render the layout in a site context and verify the output.

## Instructions

### Compose layouts

- Model websites as distinct page types, with one layout per type.
- Use `ht-include` for common page sections and inline only page-specific content.
- Start every layout with a `<header>` element that includes the header fragment, and end it with a `<footer>` element that includes the footer fragment.
- Extract a fragment after a temporary design proves reusable.

### Select directives

- Use `ht-include` for fragments, `ht-if` for conditions, `ht-each` for iteration, `ht-apply` with template variables for content and attributes, and `ht-block` for composable regions.
- Define the fragment's expected data before selecting `ht-with` or `ht-each`; consult [template-data.md](references/template-data.md) as needed.
- Use `ht-debug`, `ht-pipe`, `ht-attrs`, and `ht-base` only when the simpler directives cannot express the layout.
- Do not reach for JavaScript components until the HTML model is exhausted.
- Use `ht-apply` with template variables instead of legacy `ht-content`, `ht-attr`, `ht-attrs`, `ht-param`, and `ht-query` patterns.

### Structure layout documents

- Organize `<body>` into top-level `<header>`, `<main>`, and `<footer>`.
- Build page regions as `<section>` elements and use `<article>` for article content.
- Make each substantive section its own fragment; do not accumulate a whole homepage in one layout file.
- Inline a section only when it is page-specific, simple, and primarily composes other fragments.
- When two or more top-level section fragments reuse a component, make that component a fragment too.
- Compare extracted section fragments for repeated HTML blocks and component styles; extract those into smaller shared fragments.
- Use `hypertexting-html` to choose semantic elements, custom elements, component attributes, and styling hooks.
- Colocate a fragment’s component styles in that fragment’s HTML file.

### Verify rendered output

- Test fragment composition inside a site directory with a throwaway page.
- Validate rendered output as well as build logs, especially when debugging bindings.

## Known issues

- `ht-pipe` does not deduplicate piped content. Repeated inclusion can emit repeated CSS or text; treat that output as expected behavior.

## Tools

- `hyperctl dev render` renders within a site and theme context; it is not a standalone fragment sandbox and rejects absolute layout paths.

## References

- Consult [template-data.md](references/template-data.md) as needed when deciding what template data is visible or how a fragment receives it.
- Consult [directives.md](references/directives.md), [variables.md](references/variables.md), or [expressions.md](references/expressions.md) as needed for directive terminology, template variables, and expression tokens.
- Consult [ht-include.md](references/ht-include.md), [ht-if.md](references/ht-if.md), [ht-each.md](references/ht-each.md), [ht-apply.md](references/ht-apply.md), [ht-with.md](references/ht-with.md), or [ht-block.md](references/ht-block.md) as needed for the core layout directives.
- Consult [ht-debug.md](references/ht-debug.md), [ht-attrs.md](references/ht-attrs.md), [ht-base.md](references/ht-base.md), or [ht-pipe.md](references/ht-pipe.md) as needed for specialized layout directives.
- Consult `hyperctl agent skill describe hypertexting-html` as needed for semantic element, custom-element, attribute, and class decisions.
- Consult the [templating](https://hypertemplates.net/docs/#introduction-to-templating), [layouts](https://hypertemplates.net/docs/reference/core/layouts/), [fragments](https://hypertemplates.net/docs/reference/core/fragments/), or [directives](https://hypertemplates.net/docs/reference/core/directives/) documentation only if the bundled references do not answer the question.
