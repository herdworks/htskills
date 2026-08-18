# Component styles

## Overview

Style the semantic markup and custom elements defined by the HTML and layout
skills. Give each component a baseline treatment and only the states its
contract supports.

## Instructions

- Style custom elements and their attributes directly instead of creating class-based component identities.
- Establish baseline styling for links, actions, lists, figures, tables, forms, feedback messages, and navigation when the website uses them.
- Add `:focus-visible` styling to every interactive component.
- Add hover, disabled, selected, expanded, error, empty, and loading styles only when the component exposes that state.
- Keep component defaults and responsive rules in the fragment that owns the component markup.
- Keep layout-specific placement and spacing in the layout style block after collected component styles.
- Use `hypertexting-html` for semantics, `hypertexting-layouts` for reusable markup, and `hypertexting-javascript` only for necessary browser behavior.

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

Consult `hypertexting-layouts` for the component-style collector and layout-style source-order contract.
