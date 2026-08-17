---
name: hypertexting-theme-development
description: Develop new HyperTexting themes from visual design assets.
version: "2026-08-17"
---

# HyperTexting theme

## Overview

Use this skill to extend supplied visual direction—such as Figma, Adobe design assets,
screenshots, or an AI/code mockup—into a complete HyperTexting website theme.
A complete theme provides layouts, shared regions, visual systems, components,
and a documented customization contract. It does not provide the website's real
copy, records, editorial images, or content strategy.

Establish theme requirements before designing it. The requirements decide what
the theme must cover; the design creates a cohesive user experience across
every required surface. Document the reconciled decisions in the theme README
before authoring layouts or styles.

## Workflow

1. Use [bootstrap.md](references/prerequisites/bootstrap.md) to resolve theme prerequisites.
2. Inspect the supplied visual direction and website context.
3. Create or update theme requirements using [information-architecture.md](references/requirements/information-architecture.md), then [experience.md](references/requirements/experience.md).
4. Create or update the theme design using [responsive-grid.md](references/design/responsive-grid.md), [color-schemes.md](references/design/color-schemes.md), [typography.md](references/design/typography.md), [components.md](references/design/components.md), and [customization.md](references/design/customization.md); extend the supplied direction conservatively and avoid unnecessary follow-up questions.
5. Reconcile the theme requirements and design decisions, then document them in the theme's `README.md`.
6. Implement the theme using the focused HyperTexting skills, rendering each required surface with existing content or fixture data.
7. Use [design-review.md](references/review/design-review.md) to review the working theme with the user and gather feedback.
8. Repeat steps 3–7 as needed based on user feedback.

## Instructions

### Establish theme requirements

- Read [bootstrap.md](references/prerequisites/bootstrap.md) to resolve theme prerequisites before establishing requirements.
- Read [information-architecture.md](references/requirements/information-architecture.md) before choosing experience patterns or components.
- Infer the information architecture from the supplied mockups, existing website data, and stated goals.
- Recommend the documented baseline information architecture when the supplied context is insufficient.
- Treat the documented matrix as the checklist for the complete theme, not as a requirement to create production content.
- Use [experience.md](references/requirements/experience.md) to turn the documented matrix into experience architecture: page presentations, shared regions, and reusable section patterns.

### Design the theme systems

- Use the responsive grid, color scheme, typography, and components references as one cohesive system.
- Treat components as a core theme building block alongside the grid, color scheme, and typography system.
- Start every component with an appropriate semantic HTML pattern; define its visual and state contract before delegating its mechanics.
- Define only the components and states required by the documented experience architecture or explicitly supported by the theme.
- Use [customization.md](references/design/customization.md) to design theme defaults and website-author inputs before writing layouts that consume them.

### Route implementation

- Use `hypertexting-data` for namespace files, supported formats, and computed data.
- Use `hypertexting-layouts` for layouts, fragments, directives, and fragment data contracts.
- Use `hypertexting-html` for semantic markup and accessibility semantics.
- Use `hypertexting-css` for static CSS foundations, component styles, and cascade mechanics.
- Use `hypertexting-assets` for asset management, placement, and precedence.
- Use `hypertexting-javascript` only for progressive enhancement that semantic HTML and CSS cannot provide.
- Use `hypertexting-themes` to organize, package, install, switch, or migrate a completed theme.

## Tools

- Use `hyperctl theme init --config <site.yaml> --name <theme>` to create a named theme in an existing website's configured `themes/` directory.
- Use `hyperctl theme render --theme <theme.json> --data <fixture>.json --binding page --layout <layout>` to render required layouts against fixture data.
- Use `hyperctl website build --config <site.yaml> --id <build-id>` to build a complete demo website deterministically.

## References

- Consult [bootstrap.md](references/prerequisites/bootstrap.md) first to resolve theme prerequisites.
- Consult [information-architecture.md](references/requirements/information-architecture.md) to decide the theme's coverage.
- Consult [experience.md](references/requirements/experience.md) to define experience architecture and section patterns.
- Consult [responsive-grid.md](references/design/responsive-grid.md), [color-schemes.md](references/design/color-schemes.md), [typography.md](references/design/typography.md), and [components.md](references/design/components.md) to establish the design system.
- Consult [customization.md](references/design/customization.md) to define theme defaults and copyable website data inputs.
- Consult [design-review.md](references/review/design-review.md) to prepare and lead a working-theme review.
- [WordPress Theme Handbook](https://developer.wordpress.org/themes/getting-started/reading-this-handbook/) for theme structure, customization, accessibility, and testing precedents.
- [Ghost themes](https://docs.ghost.org/themes) for template/data separation, custom settings, and validation precedents.
- [Squarespace Guide to Website Design](https://static1.squarespace.com/static/5134cbefe4b0c6fb04df8065/t/622a6beb7503a36b116a4a18/1646947329644/guide-to-website-design.pdf) for audience-first site structure, calls to action, and responsive design.
- [thoughtbot: The Disciplines of Web Design](https://thoughtbot.com/blog/designing-for-the-web) for web-design principles and visual hierarchy.
- [Evil Martians: 100 developer-tool landing pages](https://evilmartians.com/chronicles/we-studied-100-devtool-landing-pages-here-is-what-actually-works-in-2025) for feature, proof, and CTA pattern evidence.
