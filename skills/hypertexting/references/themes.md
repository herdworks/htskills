# Theme configuration

## Overview

A HyperTexting theme root is the directory containing `theme.json`. It configures the reusable presentation bundle for a website. Theme configuration files are always JSON; do not create `theme.yaml` or `theme.yml` files. Use `hyperctl website schema describe theme` for the machine-readable companion schema provided by HyperTemplates.

A theme is the reusable presentation bundle for a website: layouts, fragments, styles, static assets, plugins, and theme data. Keep website-specific content outside the theme.

## Instructions

- Set `site.config.theme` to a website-relative path ending in `theme.json`; it defaults to `theme.json` in the website root.
- Use `site.config.themes_dir` for discovered themes; it defaults to `themes` and discovery recursively recognizes files named `theme.json`.
- Use a top-level `data` object for theme-provided template data.
- Set `name`, `description`, and `version` to identify a distributable theme. An omitted name or version resolves to `Untitled Theme` and `0.0.0`.
- Use `config.url` for the theme URL when needed.
- Configure only safe relative directory names in `config`; absolute paths, parent traversal, and backslash-separated paths are invalid.
- Omit a directory field to use its default name.

### Choose and organize a theme

- Use a root `theme.json` only for a very small website or while prototyping.
- Use `themes/` for most websites; `site.config.themes_dir` defaults to `themes`.
- Create a custom theme at `themes/<theme>/theme.json` when the website is its primary consumer and maintainer.
- Install an external theme at `themes/<theme>/<version>/theme.json` when the website will not heavily maintain or customize it.
- Set `site.config.theme` to the selected website-relative `theme.json` path; do not escape the website root.
- Keep reusable static assets in the theme and website-specific static assets outside it.
- Rely on embedded builtins only as a fallback for an omitted theme resource; add a local resource when the theme must own or customize it.

### Package and migrate themes

- Package `theme.json` and configured theme-resource directories with `hyperctl theme package`.
- Run `unzip -t <package>.zip` after packaging to verify archive integrity.
- Use `hyperctl theme install` to install a package archive and `hyperctl theme migrate` to migrate current directive syntax.

## Tools

- `hyperctl website schema describe theme` describes the theme configuration contract.
- `hyperctl theme install` installs a theme package archive.
- `hyperctl theme package` packages a theme as a ZIP archive.
- `hyperctl theme migrate` migrates a theme to current directive syntax.

| Field | Default directory |
| --- | --- |
| `config.contenttypes_dir` | `types` |
| `config.blocks_dir` | `blocks` |
| `config.layouts_dir` | `layouts` |
| `config.fragments_dir` | `fragments` |
| `config.plugins_dir` | `plugins` |
| `config.data_dir` | `data` |
| `config.static_dir` | `static` |

## Examples

```json
{
  "name": "Website Custom Theme",
  "version": "1.0.0",
  "data": {
    "accent": "teal"
  },
  "config": {
    "layouts_dir": "layouts",
    "fragments_dir": "fragments",
    "static_dir": "static"
  }
}
```

```yaml
# site.yaml
config:
  themes_dir: themes
  theme: themes/custom/theme.json
```

## Specification

Consult the [theme configuration documentation](https://hypertemplates.net/docs/reference/core/themes/) only if this bundled reference does not answer the question.
