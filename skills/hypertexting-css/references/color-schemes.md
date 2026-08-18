# Color schemes

## Overview

Implement color as a small semantic token system. Keep palette values in static
CSS foundations and let components consume roles rather than hard-coded brand
colors.

## Instructions

- Define roles for canvas, surface, text, muted text, link, accent, border, and supported feedback states.
- Define every supported role in each supported color scheme.
- Use `prefers-color-scheme` for light and dark schemes unless the theme contract requires another selection mechanism.
- Define hover, focus, selected, disabled, error, warning, and success colors only for states the component supports.
- Preserve a non-color indicator for links, controls, errors, and state changes.
- Verify readable foreground/background contrast and focus indicators in every supported scheme.

## Examples

```css
:root {
  --color-canvas: #fff;
  --color-text: #1d1d1f;
  --color-accent: #0b6e69;
  --color-focus: #0067d8;
}

@media (prefers-color-scheme: dark) {
  :root {
    --color-canvas: #161617;
    --color-text: #f5f5f7;
    --color-accent: #6ee7dc;
    --color-focus: #8dc7ff;
  }
}
```

## Specification

Use [WCAG contrast guidance](https://www.w3.org/WAI/WCAG21/Understanding/contrast-minimum.html) when a supplied design does not establish an accessible color treatment.
