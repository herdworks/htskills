# Semantic HTML and custom elements

## Overview

Use native HTML for meaning and custom elements for a reusable presentation or behavior layer. This keeps a HyperTexting page useful as pre-rendered static HTML, while allowing CSS and JavaScript to progressively enhance it. Author the HTML first: HyperTexting pre-renders static HTML at build time through HyperTemplates, so let markup drive theme authoring.

## Instructions

### Use semantic HTML

- Choose the most specific native element for its intended content: prefer `header`, `footer`, `main`, `nav`, `section`, `article`, `aside`, `address`, `time`, `menu`, `dialog`, and `details` over generic containers.
- When a fragment represents a document region, use the appropriate native structural element as its root: `header`, `main`, `section`, `article`, or `footer`. Custom elements may provide component styling within that root; do not use a custom element in place of a page-region root solely for styling or selector identity.
- When a semantic region needs a stable styling hook but has no genuine reusable component boundary, use a narrowly scoped class on the native root rather than replacing that root with a custom element.
- Keep pre-rendered markup useful and understandable without CSS or JavaScript.
- Use a `<div>` only as a minimal positional container without its own semantic or custom styling identity.
- Use native and ARIA attributes for native semantics and accessibility; do not add arbitrary styling attributes to native elements.

### Use custom elements as component boundaries

- Name every custom element with a hyphen, such as `published-date` or `site-menu`.
- Use custom elements as reusable CSS component boundaries, even when their initial purpose is styling rather than browser behavior. Put component variants and state on the custom element as custom attributes; style them with element and attribute selectors.
- Keep a native semantic element inside the custom element when one represents the content or control, such as an `<a>` inside `<link-button>`.
- Avoid classes when a custom element or custom attribute expresses the component identity or variant.

### Progressively enhance custom elements

- Start with meaningful HTML, then add custom-element CSS, then a Web Component only when interaction requires it.
- Do not require JavaScript for essential content, navigation, or actions.
- Use [javascript.md](javascript.md) for browser-time behavior.

## Tools

- Use `hyperctl theme render --theme <theme.json> --data <fixture>.json --fragment <fragment>` with fixture JSON to verify rootless component markup.

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
