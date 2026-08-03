# Semantic HTML and custom elements

## Overview

Use native HTML for meaning and custom elements for a reusable presentation or behavior layer. This keeps a HyperTexting page useful as pre-rendered static HTML, while allowing CSS and JavaScript to progressively enhance it.

## Instructions

- Name every custom element with a hyphen, such as `published-date` or `site-menu`.
- Keep a native semantic element inside the custom element when one represents the content or control.
- Put component variants and state on the custom element as custom attributes; style them with element and attribute selectors.
- Avoid classes when a custom element or custom attribute expresses the component identity or variant.

## Examples

```html
<published-date>
  <time datetime='2026-07-28'>July 28, 2026</time>
</published-date>

<site-menu expanded>
  <menu>
    <li><a href='/'>Home</a></li>
  </menu>
</site-menu>
```

```css
published-date {
  display: block;
}

site-menu[expanded] menu {
  display: block;
}
```

## Specification

Consult the [custom elements specification](https://html.spec.whatwg.org/multipage/custom-elements.html) only if this bundled reference does not answer the question.
