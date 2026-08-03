---
name: hypertexting-javascript
description: General JavaScript authoring guidelines.
version: 1
---

# HyperTexting JavaScript

## Overview

Use browser JavaScript to progressively enhance a useful pre-rendered HTML component. Browser JavaScript is distinct from build-time JavaScript plugins: browser scripts run in the visitor's browser, while `hypertexting-plugins` transforms data or template values during a HyperTexting build.

## Workflow

1. Start with useful HTML and the custom element defined by `hypertexting-html`.
2. Add CSS for the component's baseline appearance and state.
3. Implement the smallest browser-side behavior for that one component.
4. Call an API only when runtime state, a mutation, personalized data, or an external service requires it.
5. Verify both the no-JavaScript baseline and the enhanced interaction.

## Instructions

### Keep browser behavior isolated

- Do not use browser JavaScript when HTML, CSS, template variables, or HyperTemplates directives express the requirement.
- Do not use a Web Component to format a deterministic static value; use a template-value plugin when build-time formatting is necessary.
- Enhance one custom element at a time; do not turn the entire page into a client application.
- Keep initial content, navigation, and essential actions useful without JavaScript.
- Keep selectors local to the component and avoid shadow DOM.
- Use `hypertexting-plugins` for build-time JavaScript and `hypertexting-assets` to choose the scope for a browser script asset.

### Implement Web Components

- Prefer an HTML Web Component when a component needs browser-time interaction.
- Register context-preserving listeners with `document.addEventListener("<event>", this)`.
- Implement an async `handleEvent(event)` router and named `on<EventType>` handlers, including `onDOMContentLoaded(event)` when the component initializes after parsing.
- Preserve the component's semantic HTML when JavaScript fails or is unavailable.

### Use APIs as optional islands

- Call a serverless function or other backend API only from the component that needs it.
- Treat API results, loading states, and failures as enhancements; do not replace the component's baseline content or navigation.
- Use `ht-pipe` with `as js` or `as javascript` only when fragment-owned inline code must compose into a known layout destination. Otherwise, use a static browser script asset.

## Known issues

- `ht-pipe` does not deduplicate piped JavaScript. Repeated fragment inclusion can emit repeated script text; treat that output as expected behavior.

## Tools

- `hyperctl dev render` verifies the pre-rendered HTML baseline in a site context.
- Use browser developer tools to verify the enhanced interaction and the no-JavaScript baseline.

## References

- Consult [runtime-boundaries.md](references/runtime-boundaries.md) as needed to choose HTML, CSS, browser JavaScript, a build-time plugin, or an API.
- Consult [web-components.md](references/web-components.md) as needed for the event-router pattern and component rules.
- Consult `hyperctl agent skill describe hypertexting-html`, `hyperctl agent skill describe hypertexting-assets`, or `hyperctl agent skill describe hypertexting-plugins` as needed for the owning contract.
- Consult the [HTML Web Components pattern](https://adactio.com/journal/20618) or [ht-pipe reference](https://hypertemplates.net/docs/reference/core/directives/ht-pipe/) only if the bundled references do not answer the question.
