# Components

## Overview

Components are a core theme system alongside the responsive grid, color scheme,
and typography system. Define a cohesive visual and state contract for each
component that the theme uses or explicitly supports. Use semantic HTML as the
starting point; delegate markup, CSS, fragments, and browser behavior to their
own focused skills.

## Instructions

- Start with the native semantic element or pattern that expresses the component's purpose.
- Define a baseline visual treatment for links, buttons, lists, figures and captions, tables, forms and native controls, feedback/status messages, and navigation.
- Define focus-visible styling for every interactive component.
- Define hover, disabled, selected, expanded, error, empty, and loading states only where the component supports them.
- Add breadcrumbs, pagination, badges, preview cards, and metadata patterns when the documented information architecture needs them.
- Add accordions, dialogs, modals, popovers, tabs, carousels, and toasts only when the documented experience architecture exposes them.
- Keep components visually consistent with the theme grid, color, and typography systems.
- Do not create a component solely because it appears in a generic component checklist.

## Examples

| Category | Semantic starting points | Include when |
| --- | --- | --- |
| Actions | `a`, `button`, `menu` | every theme for navigation and actions |
| Content | `ul`, `ol`, `figure`, `table` | content or page presentations need them |
| Forms | `form`, `label`, `input`, `select`, `textarea` | a search, contact, signup, or other form exists |
| Disclosure | `details`, `dialog` | content or interaction requires disclosure or overlay |
| Collections | `article`, `time`, `nav` | indexes or related-content regions exist |
| Feedback | status text, error message | a user-facing action has a result |

## Specification

Consult `hypertexting-html` for semantic markup, `hypertexting-layouts` for reusable fragments, `hypertexting-css` for styling, and `hypertexting-javascript` only for necessary progressive enhancement.
