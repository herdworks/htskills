---
name: hypertexting-assets
description: Website asset management skill.
version: "2026-08-12"
---

# HyperTexting assets

## Overview

Organize static assets by the narrowest scope that owns and reuses them.

## Workflow

1. Classify the asset by its reuse scope.
2. Place it in that scope's static directory.
3. Reference it with a stable path and deploy-specific cache busting when needed.

## Instructions

### Place assets by scope

- Place default theme assets, including placeholders, in `theme.config.static_dir`.
- Place assets shared by multiple pages in `site.config.static_dir`.
- Place an asset used by only one page alongside that page’s data file.
- Rely on HyperTemplates resolution order—theme static directory, website static directory, then page assets—when later scopes should override earlier ones.

### Reference assets

- Keep filenames stable.
- Use template variables for deploy-specific cache-busting values.

## Tools

- `hyperctl cms asset add` adds a static asset to a website.
- `hyperctl cms asset ls` lists static assets.
- `hyperctl cms asset lookup` finds a static asset by path.

## References

- [CMS assets](https://hypertemplates.net/docs/reference/cms/assets/)
- [theme static_dir](https://hypertemplates.net/docs/reference/core/themes/#theme-config-static_dir)
- [site static_dir](https://hypertemplates.net/docs/reference/cms/website/#site-config-static_dir)
- [theme commands](https://hypertemplates.net/docs/reference/cli/commands/theme/)
- [asset commands](https://hypertemplates.net/docs/reference/cli/commands/cms/asset/)
