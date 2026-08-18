---
name: hypertexting-themes
description: HyperTexting theme management skill.
compatibility: Requires hyperctl.
metadata:
  version: "2026-08-18T20:28:00Z"
---

# HyperTexting themes

## Overview

A HyperTexting theme is the reusable presentation bundle for a website: at least one layout and its associated fragments, styles, static assets, plugins, and theme data. Theme data supplies default and sample values that layouts can use as fallbacks. Every HyperTexting website uses a theme—whether configured as a local `theme.json` at the website root or organized under a subdirectory. Keep website-specific content outside the theme. HyperTemplates provides the underlying theme organization and page-rendering engine.

Use this skill to organize, install, package, switch, or migrate themes.

## Workflow

1. Choose a local, custom, or versioned theme pattern.
2. Point `site.config.theme` at that pattern's `theme.json` root.
3. Organize reusable layouts, fragments, styles, static assets, plugins, and theme data in the theme.
4. Validate, package, or migrate the theme with the appropriate command.

## Instructions

### Choose a theme pattern

- Use a local theme at `theme.json` in the website root only for a very small website or while prototyping a new design.
- Use the `themes/` directory for most websites; `site.config.themes_dir` defaults to `themes`.
- Create a custom theme at `themes/<theme>/theme.json` when the website is the primary consumer and maintainer. Name it `custom`, `default`, or after the website.
- Install an external theme at `themes/<theme>/<version>/theme.json` when the website will not heavily maintain or customize it. Keep the version directory to simplify upgrades and coexistence.
- Set `site.config.theme` to the selected, website-relative `theme.json` path. Do not use paths that escape the website root.

```text
site.yaml
themes/
  <theme>/
    theme.json
    <version>/
      theme.json
```

### Organize a theme root

- Treat the theme as the reusable presentation layer of every HyperTexting website.
- Keep website-specific content and static assets outside the theme; keep reusable theme static assets in it.
- Put only broadly reusable defaults in the theme.
- Place package resources in the configured theme directories: `types`, `blocks`, `layouts`, `fragments`, `plugins`, `data`, and `static` by default.
- Configure alternate directory names only as safe relative paths in `theme.json`.
- Rely on embedded builtins only as a fallback for an omitted theme resource; add a local resource when the theme must own or customize it.

### Package themes

- Package `theme.json` and files from configured theme resource directories with `hyperctl theme package`.
- Run `unzip -t <package>.zip` after packaging to verify archive integrity.
- Keep unrelated files outside those directories or configure their directory before packaging.

## Tools

- `hyperctl agent schema describe theme` describes the theme configuration contract.
- `hyperctl theme install` installs a theme package archive.
- `hyperctl theme package` packages a theme as a ZIP archive.
- `hyperctl theme migrate` migrates a theme to current directive syntax.

## References

- Consult [config.md](references/config.md) as needed for the `theme.json` contract and default resource directories.
- Consult the [themes](https://hypertemplates.net/docs/reference/core/themes/) or [theme commands](https://hypertemplates.net/docs/reference/cli/commands/theme/) documentation only if the bundled reference does not answer the question.
