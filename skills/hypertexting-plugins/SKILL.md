---
name: hypertexting-plugins
description: HyperTexting plugins skill.
version: "2026-08-17"
---

# HyperTexting plugins

## Overview

Use a plugin only when static HTML, CSS, template variables, and source data cannot express deterministic build-time data or value transformation. Plugins are JavaScript modules whose default export is a function; HyperTemplates runs them while building template data or pre-rendering a value.

## Workflow

1. Define the data or template contract before writing code.
2. Choose a computed-data plugin for reusable derived data or a template-value plugin for one rendered value.
3. Keep the plugin with the website or theme that owns its contract.
4. Render the affected page and verify its pre-rendered output.

## Instructions

### Choose the plugin contract

- Use CSS for a visual-only transformation, such as uppercase text; do not create a plugin for `text-transform: uppercase`.
- Implement a template-value plugin only when a value must take a specific deterministic form in the pre-rendered output and HTML, CSS, directives, or source data cannot express it.
- Implement a computed-data plugin only after `hypertexting-data` has defined a reusable derived namespace and its data contract.
- Derive a computed namespace from the `.js` file's relative path under `site.config.data_dir`: `social/twitter.js` becomes `social.twitter`.
- Return a rendered value from a template-value plugin or structured data from a computed-data plugin.
- Keep plugin inputs and outputs narrow, explicit, and serializable.

### Preserve the markup-first architecture

- Do not use a plugin for browser-time interaction; use `hypertexting-javascript` to progressively enhance semantic HTML instead.
- Do not use a plugin when a static data file, template variable, or directive expresses the requirement.
- Keep semantic HTML and markup ownership in the layout. A plugin should return data or a value, not become an alternate HTML authoring system.
- Keep pre-rendered HTML useful without JavaScript.

## Tools

- `hyperctl website build` verifies build-time plugin output in the full site context.

## References

- Consult [choosing-plugins.md](references/choosing-plugins.md) as needed to distinguish CSS, template-value plugins, computed-data plugins, and Web Components.
- Consult `hyperctl agent skill describe hypertexting-data` as needed before implementing a computed-data plugin.
- Consult the [plugins](https://hypertemplates.net/docs/reference/core/plugins/), [template-variable plugins](https://hypertemplates.net/docs/reference/core/variables/#template-variable-plugins), or [computed namespaces](https://hypertemplates.net/docs/reference/cms/namespaces/#computed-namespace-plugins) documentation only if the bundled skill does not answer the question.
