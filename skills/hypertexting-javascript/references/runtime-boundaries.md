# Browser JavaScript boundaries

## Overview

Keep HyperTexting static-first. Choose the smallest layer that satisfies the requirement, and do not confuse browser scripts with build-time plugins.

## Instructions

| Requirement | Owner | Use |
| --- | --- | --- |
| Content, document structure, or native behavior | `hypertexting-html` and `hypertexting-layouts` | Pre-rendered HTML and directives |
| Presentation, responsive rules, or component state styling | `hypertexting-css` | Light-DOM CSS |
| One deterministic pre-rendered value | `hypertexting-plugins` | Template-value plugin |
| Reusable derived build data | `hypertexting-plugins` | Computed-data plugin |
| Isolated interaction in an already useful component | `hypertexting-javascript` | Browser JavaScript or Web Component |
| Runtime state, mutation, personalized data, or external service | Component plus API | As-needed serverless function or backend API |

- Keep browser JavaScript in a static asset at the narrowest theme, website, or page scope that owns it.
- Do not use browser APIs such as `window`, `document`, or `fetch` in a build-time plugin.
- Do not use a build-time plugin to implement visitor interaction.
- Keep an API call inside the browser component that owns the interaction.

## Examples

```text
Archive grouping from site.pages: computed-data plugin
ISO duration or static editorial date: template-value plugin
Expandable site menu: browser behavior for <site-menu>
Relative time that updates or localizes per visitor: browser behavior for <relative-time>
Newsletter signup submission: isolated API request from the signup component
```

## Specification

Consult the [Jamstack architecture overview](https://jamstack.org/what-is-jamstack/) only if this bundled reference does not answer the question.
