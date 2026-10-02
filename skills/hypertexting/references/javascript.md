# Browser Web Components

## Overview

Use a component to add behavior to an existing custom element, not to replace its pre-rendered content. Keep its selectors, events, and optional API calls local to that element. Browser JavaScript progressively enhances useful pre-rendered HTML; it is distinct from build-time JavaScript plugins, which transform data or template values during a HyperTexting build.

## Instructions

- Define behavior for the custom element authored by [html.md](html.md).
- Keep semantic child HTML in place; render only the interactive enhancement or state changes.
- Register document-level events with the component instance so `this` remains the event handler.
- Route events through `handleEvent(event)` and implement named asynchronous handlers such as `onDOMContentLoaded(event)`.
- Avoid shadow DOM and broad document selectors.

### Keep browser behavior isolated

- Do not use browser JavaScript when HTML, CSS, template variables, or directives express the requirement.
- Do not use a Web Component to format a deterministic static value; use a template-value plugin when build-time formatting is necessary.
- Enhance one custom element at a time; do not turn the entire page into a client application.
- Keep initial content, navigation, and essential actions useful without JavaScript.
- Keep selectors local to the component and browser-script assets in the narrowest owning theme, website, or page static scope.
- Do not use browser APIs such as `window`, `document`, or `fetch` in a build-time plugin.

### Use APIs as optional islands

- Call a serverless function or other backend API only from the component that needs it.
- Treat API results, loading states, and failures as enhancements; do not replace the component’s baseline content or navigation.
- Use destination-side `ht-pipe='from "script.component" as js'` only when fragment-owned inline code must compose into a known layout destination; otherwise use a static browser script asset.

## Tools

- Use `hyperctl theme render --theme <theme.json> --data <fixture>.json --fragment <fragment>` to verify a component’s pre-rendered HTML baseline, then use browser developer tools to verify enhanced interaction and the no-JavaScript baseline.

## Examples

```js
class SiteMenu extends HTMLElement {
  connectedCallback() {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", this, { once: true })
      return
    }
    void this.onDOMContentLoaded(new Event("DOMContentLoaded"))
  }

  async handleEvent(event) {
    const handler = this[`on${event.type}`]
    if (handler) await handler.call(this, event)
  }

  async onDOMContentLoaded(event) {
    this.render()
  }

  render() {
    this.setAttribute("ready", "")
  }
}

customElements.define("site-menu", SiteMenu)
```

## Specification

Consult the [HTML Web Components pattern](https://adactio.com/journal/20618) only if this bundled reference does not answer the question.
