---
name: hypertexting-css
description: General CSS authoring guidelines.
version: 1
---

# HyperTexting CSS

## Overview

Build light-DOM styles from shared foundations and fragment-owned component CSS.

## Workflow

1. Establish shared CSS foundations in static files.
2. Colocate component CSS with the fragment that owns its markup.
3. Render the layout and verify the resulting cascade and responsive behavior.

## Instructions

### Establish CSS foundations

- Avoid shadow DOM.
- Use light DOM and static CSS in this order: `reset.css`, `variables.css`, optional `grid.css`, then `colors.css` with at least light and dark schemes.
- Use no more than three or four responsive breakpoints.
- Give CSS custom-property reads a fallback, such as `padding: var(--example, 15px)`.
- Use CSS for visual-only transformations, such as `text-transform: uppercase`; do not create a build-time plugin for them.

### Colocate component-specific CSS

- Always add `<style id='components'></style>` before `<style id='layout'></style>` in each layout head so layout rules can override component defaults.
- Style custom elements and their custom attributes directly. Prefer `article-card`, `article-card[featured]`, and `site-menu[collapsed]` over class-based selectors.
- Create only the few CSS classes a website genuinely needs; most HyperTexting websites need few or none.
- Put `<style ht-pipe='to "style#components" as css'></style>` in the fragment that owns the component markup.
- Keep the component's default and trailing `@media` rules in that fragment.
- Do not move fragment-owned component styles into a generic static `style.css` during a layout refactor.
- Reserve static CSS files for shared foundations: reset, variables, grid, colors, and genuinely site-wide rules.

### Apply layout-specific overrides

- Put layout-specific placement, spacing, and contextual overrides in the layout's `<style id='layout'></style>` block, after `style#components`.
- Do not create a fragment variant solely for layout-specific styling.
- Create a fragment variant only when its markup, data contract, or reusable behavior differs.

### Cache-bust styles

- Cache-bust static CSS with template variables: `<link ht-apply rel='stylesheet' href='/css/reset.css?version=${ site.version, build.id }'>`.

## Known issues

- `ht-pipe` does not deduplicate piped CSS. Repeated inclusion can emit repeated component styles; treat that output as expected behavior.

## Tools

- `hyperctl dev render` renders the layout for cascade and responsive-behavior checks.

## References

- [ht-pipe](https://hypertemplates.net/docs/reference/core/directives/ht-pipe/)
- [CSS Web Components for Marketing Sites](https://hawkticehurst.com/2024/11/css-web-components-for-marketing-sites/)
