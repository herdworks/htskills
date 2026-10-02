# Component styles

## Overview

Style the semantic markup and custom elements defined by the HTML and layout
skills. Give each component a baseline treatment and only the states its
contract supports.

Build light-DOM styles from shared foundations and fragment-owned component CSS.

## Instructions

- Style custom elements and their attributes directly instead of creating class-based component identities. This component-style boundary does not replace the native root of a fragment that represents a document region.
- Establish baseline styling for links, actions, lists, figures, tables, forms, feedback messages, and navigation when the website uses them.
- Add `:focus-visible` styling to every interactive component.
- Add hover, disabled, selected, expanded, error, empty, and loading styles only when the component exposes that state.
- Keep component defaults and responsive rules in the fragment that owns the component markup.
- Keep layout-specific placement and spacing in the layout style block after collected component styles.
- Use [html.md](html.md) for semantics, [layouts.md](layouts.md) for reusable markup, and [javascript.md](javascript.md) only for necessary browser behavior.

### Establish CSS foundations

- Avoid shadow DOM.
- Use light DOM and static CSS in this order: `reset.css`, `variables.css`, optional `grid.css`, then `colors.css` for every color scheme the theme supports.
- Implement light and dark schemes with `prefers-color-scheme` unless the theme requirements explicitly indicate otherwise.
- Give CSS custom-property reads a fallback, such as `padding: var(--example, 15px)`.
- Use CSS for visual-only transformations, such as `text-transform: uppercase`; do not create a build-time plugin for them.
- Reserve static CSS files for shared foundations and genuinely site-wide rules.

### Colocate component-specific CSS

- Include these two style elements in every layout `<head>`, with `style#components` before `style#layout`:

  ```html
  <style id='components' ht-pipe='from "style.component" as css'></style>
  <style id='layout'>
    /* layout-specific styles go here */
  </style>
  ```

- Put `<style class='component'></style>` and its default and trailing `@media` rules in the fragment that owns the component markup.
- Keep layout-specific placement, spacing, and contextual styling in `style#layout`, after collected component styles.
- Do not create a fragment variant solely for layout-specific styling; create one only when markup, data contract, or reusable behavior differs.
- Consult [layouts-ht-pipe.md](layouts-ht-pipe.md) for the collector contract.
- Cache-bust static CSS with template variables: `<link ht-apply rel='stylesheet' href='/css/reset.css?version=${ site.version, build.id }'>`.

### Color, grid, and type systems

- Define semantic roles for canvas, surface, text, muted text, link, accent, border, focus, and only the feedback states a component supports. Define every supported role in every supported scheme and preserve a non-color indicator for links, controls, errors, and state changes.
- Define a readable content container, flexible columns, shared gutter and vertical-spacing tokens, and only the shared geometry a website needs.
- Use no more than three or four responsive breakpoints unless supplied direction requires more. Preserve meaningful reading order and usable navigation at narrow widths.
- Test long labels, text expansion, wide tables, oversized media, URLs, and non-English text at every supported width.
- Define only needed font roles; use licensed `woff2` resources from the owning static scope, `font-display: swap`, sensible fallbacks, a consistent scale, readable line lengths, and semantic heading hierarchy.

## Tools

- Use `hyperctl theme render --theme <theme.json> --data <fixture>.json --fragment <fragment>` with fixture JSON to verify component markup and styles; use layout mode to verify the final cascade and responsive behavior.

## Examples

```css
site-menu a {
  color: var(--color-text);
}

site-menu a:focus-visible {
  outline: 0.2rem solid var(--color-focus);
  outline-offset: 0.2rem;
}

article-card[featured] {
  background: var(--color-surface);
}
```

## Specification

Consult [layouts-ht-pipe.md](layouts-ht-pipe.md) for the component-style collector and layout-style source-order contract.
