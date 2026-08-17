# Experience architecture

## Overview

Turn the documented information architecture into experience architecture: the
user experience across page presentations, shared regions, and narrative
patterns. Use `hypertexting-layouts` to implement layouts, fragments, and
directives.

## Instructions

- Define one page presentation for each distinct page family in the page-and-navigation matrix.
- Use a shared header, primary navigation, footer, and document head wherever their structure and behavior are common.
- Add section-specific navigation only when a section contains enough depth to require it.
- Select section patterns that help visitors understand the page's purpose before decorating it.
- Define the shared regions and reusable patterns the theme needs; defer all layout implementation details and best practices to `hypertexting-layouts`.

### Select narrative patterns

- Use a hero to orient visitors to a page's purpose and next step.
- Use a feature pattern when a page must explain capabilities, benefits, or differentiators.
- Use product screenshots, imagery, examples, or outputs as evidence when they make the claim more understandable.
- Use proof patterns for testimonials, recognizable users, traction, awards, or compatibility when evidence exists.
- Use explanatory or how-it-works patterns when the subject is otherwise difficult to understand.
- Use an editorial pattern for reading, metadata, related content, and continuation when the detail page is prose-led.
- Keep a primary CTA distinct from any secondary action.

### Ask about narrative direction

- Ask or revisit which claim, action, or evidence a page should emphasize only when the mockup, supplied context, or feedback leaves that choice materially unclear.
- Ask or revisit whether proof, screenshots, examples, or testimonials are available when they would determine the selected pattern.
- Do not ask the user which page presentations to build; derive them from the information architecture.
- Revisit these questions as needed when feedback changes a page's narrative direction or available evidence; otherwise retain the documented decisions.

### Refine from feedback

Ask these questions only when user feedback makes the answer material:

- Does the feedback change the visitor journey or only this page's presentation?
- Does the requested pattern need to work across a page family, or is it a one-page-presentation exception?
- Does the feedback reveal a system-level issue with the responsive grid, color scheme, typography, or component behavior?

## Examples

| Page family | Typical shared regions | Candidate patterns |
| --- | --- | --- |
| home | header and footer | hero, proof, feature, CTA |
| landing | header, section navigation, footer | introduction, feature, index preview |
| index | header, footer | filters when required, preview collection, pagination |
| detail | header, footer | metadata, body, related content, CTA |
| standard | header and footer | title, body, optional callout |
| status | minimal header or footer | explanation, recovery link |

## Specification

Consult `hypertexting-layouts` for HTML document composition, fragments, directives, and rendered verification.
