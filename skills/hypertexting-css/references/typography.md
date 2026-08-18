# Typography

## Overview

Implement a small type system that preserves semantic document hierarchy and
keeps long-form content, controls, metadata, and headings readable across
supported viewport sizes.

## Instructions

- Define only the font roles the website needs, such as display, body, UI, and code.
- Place every licensed custom webfont the theme ships and uses in the owning theme or website static directory; use `hypertexting-assets` to choose that scope.
- Declare each used custom font family, style, and weight resource with `@font-face`; do not declare unused faces from a supplied font package.
- Prefer `woff2` files and declare a variable font's supported `font-weight` or `font-style` range when available.
- Set `font-display: swap` and provide a sensible system fallback in every `font-family` stack.
- Use a consistent scale for headings, body text, supporting text, captions, and metadata.
- Set readable line lengths and line heights for sustained reading.
- Let semantic heading levels define hierarchy; use CSS only for visual treatment.
- Scale type conservatively at narrow widths so text remains readable and controls remain operable.
- Test long titles, labels, URLs, and non-English text before treating a type rule as complete.

## Examples

```css
@font-face {
  font-family: "Example Sans";
  src: url("/fonts/example-sans.woff2") format("woff2");
  font-display: swap;
  font-style: normal;
  font-weight: 100 900;
}

:root {
  --font-body: "Example Sans", system-ui, sans-serif;
  --font-display: Georgia, serif;
  --text-body: 1rem;
  --leading-body: 1.6;
}

article {
  font: var(--text-body) / var(--leading-body) var(--font-body);
  max-width: 70ch;
}

h1, h2, h3 {
  font-family: var(--font-display);
  line-height: 1.1;
}
```

## Specification

Consult `hypertexting-html` when the markup does not express a meaningful heading or text relationship.
