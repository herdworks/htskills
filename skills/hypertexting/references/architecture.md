# Static-first architecture

## Overview

HyperTexting builds static websites: HyperTemplates pre-renders HTML and static assets at build time, and the resulting files can be served without request-time page rendering. This follows a static-first, decoupled approach aligned with Jamstack principles.

## Instructions

- Build complete, useful HTML and static assets before adding browser-time behavior or backend dependencies.
- Use semantic HTML and CSS as the baseline experience; progressively enhance only where a component needs behavior.
- Treat each interactive custom element as an isolated, islands-style enhancement. Do not turn an entire page into a client application.
- Use the Web Component API only for the individual custom elements that need browser-time interaction.
- Use serverless functions or other backend APIs only when runtime state, personalized data, mutations, or an external service requires them.
- Keep initial content and navigation independent of JavaScript and backend APIs whenever possible.

## Examples

```text
Static page: pre-rendered HTML, CSS, and static assets
Interactive menu: a progressively enhanced <site-menu> custom element
Account action: an isolated API request from the component that needs it
```

## Specification

Consult the [Jamstack architecture overview](https://jamstack.org/what-is-jamstack/) or [MDN static site generator glossary](https://developer.mozilla.org/en-US/docs/Glossary/SSG) only if this bundled reference does not answer the question.
