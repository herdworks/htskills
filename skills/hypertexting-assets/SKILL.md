---
name: hypertexting-assets
description: HyperTexting asset management skill.
version: "2026-08-17"
---

# HyperTexting assets

## Overview

Manage durable design assets separately from static assets delivered by a
HyperTexting website. Design assets are customer-provided or approved inputs
that inform the website or theme; static assets are the web-delivered files in
configured theme, website, or page static directories.

## Workflow

1. Classify the input as a design asset or a static asset.
2. Preserve design assets at the website root and document them.
3. Place static assets in the narrowest scope that owns and reuses them.
4. Reference static assets with a stable path and deploy-specific cache busting when needed.

## Instructions

### Preserve design assets

- Keep customer-provided or approved design assets in a dedicated website-root directory outside every configured static directory. Prefer `design/`; use `design-assets/` or `assets/` when appropriate.
- Never configure the design-asset directory as a static directory; copy a deliberate production version to the appropriate static scope instead.
- Create a `README.md` in that design-asset directory. At minimum, state that the directory preserves design assets and that live production assets belong in the configured static directory.
- Preserve originals and references, such as mockups, logos, fonts, imagery, video, artwork, and brand guides, without treating them as publishable by default.
- Add notes to the design-asset `README.md` as needed for provenance, intended use, licensing, or derived static assets.
- Consult [design-assets.md](references/design-assets.md) when organizing or converting design material.

### Place static assets by scope

- Place default theme static assets, including placeholders, in `theme.config.static_dir`.
- Place static assets shared by multiple pages in `site.config.static_dir`.
- Place a static asset used by only one page alongside that page’s data file.
- Rely on HyperTemplates resolution order—theme static directory, website static directory, then page static assets—when later scopes should override earlier ones.
- Copy or derive a static asset from a design asset deliberately; do not move, optimize, or overwrite the design asset in place.

### Reference assets

- Keep filenames stable.
- Use template variables for deploy-specific cache-busting values.

## Tools

- `hyperctl website asset add` adds a static asset to a website.
- `hyperctl website asset ls` lists static assets.
- `hyperctl website asset lookup` finds a static asset by path.

## References

- [CMS assets](https://hypertemplates.net/docs/reference/cms/assets/)
- Consult [design-assets.md](references/design-assets.md) as needed to distinguish durable design inputs from live static assets.
- [theme static_dir](https://hypertemplates.net/docs/reference/core/themes/#theme-config-static_dir)
- [site static_dir](https://hypertemplates.net/docs/reference/cms/website/#site-config-static_dir)
- [theme commands](https://hypertemplates.net/docs/reference/cli/commands/theme/)
- [asset commands](https://hypertemplates.net/docs/reference/cli/commands/cms/asset/)
