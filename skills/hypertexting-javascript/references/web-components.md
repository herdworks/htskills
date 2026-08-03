# Browser Web Components

## Overview

Use a component to add behavior to an existing custom element, not to replace its pre-rendered content. Keep its selectors, events, and optional API calls local to that element.

## Instructions

- Define behavior for the custom element authored by `hypertexting-html`.
- Keep semantic child HTML in place; render only the interactive enhancement or state changes.
- Register document-level events with the component instance so `this` remains the event handler.
- Route events through `handleEvent(event)` and implement named asynchronous handlers such as `onDOMContentLoaded(event)`.
- Avoid shadow DOM and broad document selectors.

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
