---
name: hypertexting
description: HyperTexting website orchestration skill.
version: "2026-08-17"
---

# HyperTexting

## Overview

- Select and combine focused HyperTexting skills.
- Keep each concern in its owning skill instead of duplicating detailed instructions here.
- Build static, pre-rendered websites first; add isolated browser behavior and backend APIs only as needed.

## Workflow

1. Identify the affected concerns and read each corresponding focused skill.
2. Apply the focused workflows in dependency order: data and themes, layouts and HTML, then plugins, assets, CSS, and JavaScript.
3. Validate the rendered site or the relevant focused result.

## Instructions

### Compose at three levels

- Use `hypertexting-html` to author semantic elements and markup-driven components.
- Use `hypertexting-layouts` to compose HTML into reusable layouts and fragments with template data and directives.
- Use `hypertexting-theme-development` to turn visual direction and website context into a complete, content-agnostic theme.
- Use `hypertexting-themes` to organize reusable layouts and related presentation resources into a website theme.

### Route focused work

- Use `hypertexting-data` to choose data scope, source files, and computed-namespace contracts.
- Use `hypertexting-theme-development` to create or redesign a theme before its layouts and implementation details are authored.
- Use `hypertexting-themes` for reusable themes and theme packaging or migration.
- Use `hypertexting-layouts` for semantic layouts, fragments, and directives.
- Use `hypertexting-assets` for design assets and theme, website, and page static assets.
- Use `hypertexting-html` for semantic markup, custom elements, and markup-first progressive enhancement.
- Use `hypertexting-plugins` for reusable template-variable or computed-data plugins after their contract is defined.
- Use `hypertexting-css` for light-DOM and component CSS.
- Use `hypertexting-javascript` for progressive enhancement and Web Components.
- Use `hypertexting-debugging` for unexpected rendered output, directive behavior, or template-data scope.

### Deterministic builds

- Pass a predictable build identifier directly: `hyperctl website build --id <build-id>`.
- Do not wrap `hyperctl website build` in a script merely to capture its auto-generated ID.

## Tools

- `hyperctl agent skill install hypertexting <path>` installs this orchestration skill.
- `hyperctl agent skill update hypertexting <path>` updates this orchestration skill.
- `hyperctl agent skill install hypertexting-data <path>` installs the data skill.
- `hyperctl agent skill update hypertexting-data <path>` updates the data skill.
- `hyperctl agent skill install hypertexting-theme-development <path>` installs the theme development skill.
- `hyperctl agent skill update hypertexting-theme-development <path>` updates the theme development skill.
- `hyperctl agent skill install hypertexting-themes <path>` installs the themes skill.
- `hyperctl agent skill update hypertexting-themes <path>` updates the themes skill.
- `hyperctl agent skill install hypertexting-layouts <path>` installs the layouts skill.
- `hyperctl agent skill update hypertexting-layouts <path>` updates the layouts skill.
- `hyperctl agent skill install hypertexting-assets <path>` installs the assets skill.
- `hyperctl agent skill update hypertexting-assets <path>` updates the assets skill.
- `hyperctl agent skill install hypertexting-html <path>` installs the HTML skill.
- `hyperctl agent skill update hypertexting-html <path>` updates the HTML skill.
- `hyperctl agent skill install hypertexting-plugins <path>` installs the plugins skill.
- `hyperctl agent skill update hypertexting-plugins <path>` updates the plugins skill.
- `hyperctl agent skill install hypertexting-css <path>` installs the CSS skill.
- `hyperctl agent skill update hypertexting-css <path>` updates the CSS skill.
- `hyperctl agent skill install hypertexting-javascript <path>` installs the JavaScript skill.
- `hyperctl agent skill update hypertexting-javascript <path>` updates the JavaScript skill.
- `hyperctl agent skill install hypertexting-debugging <path>` installs the debugging skill.
- `hyperctl agent skill update hypertexting-debugging <path>` updates the debugging skill.

## References

- Consult [architecture.md](references/architecture.md) as needed for HyperTexting's static-first architecture and progressive-enhancement model.
- `hyperctl agent skill describe hypertexting-data`
- `hyperctl agent skill describe hypertexting-theme-development`
- `hyperctl agent skill describe hypertexting-themes`
- `hyperctl agent skill describe hypertexting-layouts`
- `hyperctl agent skill describe hypertexting-debugging`
- `hyperctl agent skill describe hypertexting-assets`
- `hyperctl agent skill describe hypertexting-html`
- `hyperctl agent skill describe hypertexting-plugins`
- `hyperctl agent skill describe hypertexting-css`
- `hyperctl agent skill describe hypertexting-javascript`
