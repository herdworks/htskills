# Customization

## Overview

Define the theme's author-facing contract before layouts consume it. A theme may
provide placeholder and default values in `theme.*`, read semantic website and
page content from `site.*` and `page.*`, and request website-provided namespace
data through `data.*`. These are distinct namespaces, not an automatic
precedence chain.

## Instructions

- Put theme-provided placeholders and defaults in theme data so layouts can read them from `theme.*`.
- Read website identity and site-level semantic values from `site.*`; read page-specific semantic values from `page.*`.
- Define every expected website namespace under `data.*` in the theme README, including key paths, value shapes, required/optional status, and fallback behavior.
- Keep copyable website-data examples outside the theme's own data directory, such as `examples/data/nav.yaml` and `demo/data/nav.yaml`, because those files populate `data.*` in a website rather than `theme.*` in a theme.
- Maintain a theme-root `README.md` that documents the theme requirements, design decisions, supported components, and all customization inputs as they settle.
- Provide a demo website `/getting-started/` page when the theme includes a demo website and that page will help users install, copy sample data, or customize it.
- Use `--ht-var()` only for explicitly supported CSS customization values and always provide a fallback value.
- Do not describe `theme.*`, `site.*`, `page.*`, or `data.*` as an implicit override chain.

## Examples

```yaml
# examples/data/nav.yaml, copied to a website data directory
main:
  - label: Journal
    href: /journal/
footer:
  - label: Privacy
    href: /privacy/
```

```css
site-cta a {
  background: --ht-var("theme.accent", #0b6e69);
}
```

## Specification

Consult `hypertexting-data` for data directories, supported formats, and namespace resolution. Consult `hypertexting-themes` for the theme data directory and packaging contract.
