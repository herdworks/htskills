---
name: hypertexting-css
description: General CSS authoring guidelines.
compatibility: Requires hyperctl.
metadata:
  version: "2026-08-26T00:00:00Z"
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
- Use light DOM and static CSS in this order: `reset.css`, `variables.css`, optional `grid.css`, then `colors.css` for every color scheme the theme supports.
- Implement the theme's light and dark schemes with the operating system's `prefers-color-scheme` preference unless the theme requirements explicitly indicate otherwise.
- Use no more than three or four responsive breakpoints.
- Give CSS custom-property reads a fallback, such as `padding: var(--example, 15px)`.
- Use CSS for visual-only transformations, such as `text-transform: uppercase`; do not create a build-time plugin for them.

### Colocate component-specific CSS

- Include these two style elements in every layout `<head>`, with `style#components` before `style#layout`:

  ```html
  <style id='components' ht-pipe='from "style.component" as css'></style>
  <style id='layout'>
    /* layout-specific styles go here */
  </style>
  ```

- Add any layout-specific component styles to `style#layout`. `ht-pipe` will collect all matching component styles in `style#components`, allowing `style#layout` declarations with otherwise equal cascade weights to benefit from source order precedence (the last matching declaration wins). This predictable source ordering often reduces the need for advanced `@layer` rules or brute-force `!important` declarations.
- Treat `from "style.component"` as the default source-selector convention. Change or expand the `ht-pipe` selector when the layout needs to collect a different set of component styles.
- Style custom elements and their custom attributes directly. Prefer `article-card`, `article-card[featured]`, and `site-menu[collapsed]` over class-based selectors.
- Create only the few CSS classes a website genuinely needs; most HyperTexting websites need few or none.
- Put `<style class='component'></style>` in the fragment that owns the component markup.
- Keep the component's default and trailing `@media` rules in that fragment.
- Do not move fragment-owned component styles into a generic static `style.css` during a layout refactor.
- Reserve static CSS files for shared foundations: reset, variables, grid, colors, and genuinely site-wide rules.

### Add layout-specific component styles

- Put layout-specific component styles, including placement, spacing, and contextual styling, in the layout's `<style id='layout'></style>` block, after `style#components`.
- Do not create a fragment variant solely for layout-specific styling.
- Create a fragment variant only when its markup, data contract, or reusable behavior differs.

### Cache-bust styles

- Cache-bust static CSS with template variables: `<link ht-apply rel='stylesheet' href='/css/reset.css?version=${ site.version, build.id }'>`.

## Tools

- Use `hyperctl theme render --theme <theme.json> --data <fixture>.json --fragment <fragment>` with fixture JSON to verify component markup and styles during development; use layout mode to verify the final cascade and responsive behavior.

## References

- Consult [color-schemes.md](references/color-schemes.md) as needed to implement semantic colors, supported schemes, and accessible states.
- Consult [responsive-grid.md](references/responsive-grid.md) as needed to implement containers, layout geometry, breakpoints, and narrow-screen behavior.
- Consult [typography.md](references/typography.md) as needed to implement type roles, readable measure, and resilient text styles.
- Consult [component-styles.md](references/component-styles.md) as needed to style component baselines and supported states.
- Consult `hyperctl agent skill describe hypertexting-layouts` for the `ht-pipe` contract and examples.
- Consult the [ht-pipe reference](https://hypertemplates.net/docs/reference/core/directives/ht-pipe/) only if the bundled layout skill does not answer the question.
- [CSS Web Components for Marketing Sites](https://hawkticehurst.com/2024/11/css-web-components-for-marketing-sites/)
