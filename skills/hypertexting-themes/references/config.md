# Theme configuration

## Overview

A HyperTexting theme root is the directory containing `theme.json`. It configures the reusable presentation bundle for a website. Theme configuration files are always JSON; do not create `theme.yaml` or `theme.yml` files. Use `hyperctl agent schema describe theme` for the machine-readable companion schema provided by HyperTemplates.

## Instructions

- Set `site.config.theme` to a website-relative path ending in `theme.json`; it defaults to `theme.json` in the website root.
- Use `site.config.themes_dir` for discovered themes; it defaults to `themes` and discovery recursively recognizes files named `theme.json`.
- Use a top-level `data` object for theme-provided template data.
- Set `name`, `description`, and `version` to identify a distributable theme. An omitted name or version resolves to `Untitled Theme` and `0.0.0`.
- Use `config.url` for the theme URL when needed.
- Configure only safe relative directory names in `config`; absolute paths, parent traversal, and backslash-separated paths are invalid.
- Omit a directory field to use its default name.

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
