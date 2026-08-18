---
name: hypertexting-html
description: General HTML authoring guidelines.
compatibility: Requires hyperctl.
metadata:
  version: "2026-08-18T19:07:00Z"
---

# HyperTexting HTML

## Overview

Author the HTML first. HyperTexting pre-renders static HTML at build time through HyperTemplates, so let markup drive theme authoring; add CSS and JavaScript only as progressive enhancements. Most markup is ordinary, standards-based HTML rather than HyperTexting-specific syntax.

## Workflow

1. Choose the native HTML elements that express the content and document structure.
2. Add a custom element when the markup needs a reusable custom style or behavior.
3. Style the custom element and its attributes with CSS.
4. Enhance the custom element with the Web Component API only when browser-time behavior needs it.

## Instructions

### Use semantic HTML

- Choose the most specific native HTML element for its intended content: use elements such as `header`, `footer`, `main`, `nav`, `section`, `article`, `aside`, `address`, `time`, `menu`, `dialog`, and `details` instead of generic containers.
- Keep the pre-rendered markup useful and understandable without CSS or JavaScript.
- Avoid div soup. Use a `<div>` only as a minimal positional container without its own semantic or custom styling identity.

### Create markup-driven components

- Prefer a hyphenated custom element over a `<div>` with a custom class when markup needs custom styles or reusable behavior.
- Keep the relevant semantic native element inside a custom element when a custom variant is needed: wrap `<time>` in a date element and `<menu>` in a menu element.
- Give custom elements non-`data-*` attributes for component variants or states that CSS or JavaScript owns.
- Use native attributes and ARIA attributes for native semantics and accessibility; do not add arbitrary styling attributes to native elements.
- Prefer custom elements and their attributes over CSS classes for component identity, variants, and styling hooks. Create only the few classes a website genuinely needs; most HyperTexting websites need few or none.

### Progressively enhance custom elements

- Start with meaningful HTML, then add custom-element CSS, then add a Web Component only if interaction requires it.
- Do not require JavaScript for essential content, navigation, or actions.
- Use `hypertexting-javascript` when implementing browser-time behavior.

## Tools

- Use `hyperctl theme render --theme <theme.json> --data <fixture>.json --fragment <fragment>` with fixture JSON to verify rootless component markup without creating a page layout first.

## References

- Consult [semantic-html.md](references/semantic-html.md) as needed for custom-element patterns and examples.
- Consult the [HTML elements reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements) only if the bundled reference does not answer the question.
