# Responsive grid

## Overview

Implement only the shared geometry a website needs: readable containers,
responsive regions, gutters, spacing, and component layouts. Keep the grid in
static foundations and component-specific behavior with its owning fragment.

## Instructions

- Define a readable content container and add wider-media or full-bleed regions only when layouts need them.
- Use a small number of flexible columns; do not introduce a generic framework for unused geometry.
- Define shared gutter and vertical-spacing tokens for narrow and wide viewports.
- Use no more than three or four breakpoints unless the supplied direction requires more.
- Stack, simplify, scroll, or collapse a region at narrow widths without changing its meaningful reading order.
- Keep navigation usable at narrow widths without making essential destinations JavaScript-only.
- Test long labels, text expansion, wide tables, and oversized media at every supported width.

## Examples

```css
:root {
  --content-width: 72rem;
  --gutter: clamp(1rem, 4vw, 3rem);
}

main > section {
  margin-inline: auto;
  max-width: var(--content-width);
  padding-inline: var(--gutter);
}

article-list {
  display: grid;
  gap: 1.5rem;
  grid-template-columns: repeat(auto-fit, minmax(min(100%, 18rem), 1fr));
}
```

## Specification

Consult `hypertexting-layouts` when a responsive change requires different markup or a distinct fragment contract.
