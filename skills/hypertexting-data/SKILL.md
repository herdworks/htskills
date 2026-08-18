---
name: hypertexting-data
description: HyperTexting data management skill.
compatibility: Requires hyperctl.
metadata:
  version: "2026-08-18T19:16:00Z"
---

# HyperTexting data

## Overview

HyperTexting is a data-driven content management system and static site generator. It builds template data from website and page files plus namespaced data files on disk. HyperTemplates renders HTML from that template data.

- Add custom properties, including arbitrary data structures, to website and page objects when their names do not conflict with schema-defined fields.

## Workflow

1. Choose the data scope that owns the value.
2. Select a supported source format or a computed namespace.
3. Validate the resulting website and page object contracts.

## Instructions

### Website data

- Create a `site.yaml`, `site.yml`, or `site.json` configuration file.
- Treat the website as a `website` object; consult `hyperctl agent schema describe website` as needed.
- Define the content families, collections, and browse paths before configuring their feeds. Consult [information-architecture.md](references/information-architecture.md) when a website's page, feed, tag, archive, or navigation-data model is not yet clear.
- Define named Atom feeds in `site.feeds`. Each feed selects pages with `include` and `exclude` selectors, and may set ordered `sort` rules and an `html_url` alternate.
- Consult [feeds.md](references/feeds.md) before defining a feed identifier or selector.

### Page data

- Create pages as directories containing a page data file.
- Treat each page data file as a `page` object; consult `hyperctl agent schema describe page` as needed.
- Name page data files `index.md`, `index.yaml`, `index.json`, `page.md`, `page.yaml`, or `page.json`.
- Use `index.html` or `page.html` instead of a page data file for a pre-rendered static page.
- Prefer YAML or JSON for structured content, such as landing pages; prefer Markdown with YAML front matter for prose, such as blog posts.
- Put page-specific structured data, such as `hero` or `features` blocks, directly in the corresponding page data file.
- Set `visibility` to control whether the page appears on the website, in feeds, and in builds. Set `spec.visibility` on a content type when the content type owns that visibility choice.
- Use `feeds.<alias>` to expose a configured site feed or an automatic tag feed to this page's template. Omit `source` and `tag` to use the alias as the configured site-feed name. Consult [feeds.md](references/feeds.md) and [visibility.md](references/visibility.md) for the authored-data contracts.

**Example:** For a complex home page that is not primarily prose, use `content/index.yaml` or `content/index.json` with custom properties. Do not split its data between `content/index.md` and a separate `data/homepage.yaml` namespace.

### Namespace data

- Place unstructured namespace data files in the global data directory configured by `site.config.data_dir`.
- Use YAML (`.yaml` or `.yml`), JSON (`.json`), OPML (`.opml`), or CSV (`.csv`) for static namespace data.
- Derive each namespace from the file's relative path; consult [namespaces.md](references/namespaces.md) as needed for nested paths and collision rules.
- Define a computed namespace when source data already exists and multiple layouts or pages need another reusable shape.
- Use a template-value plugin instead when one layout needs only a deterministic formatted value.
- Use `hypertexting-plugins` to implement the computed-data plugin after defining its namespace and data contract.

**Example:** For an archive page that groups `site.pages` by year and month, create `data/archive.js` to return an archive map instead of maintaining a duplicate static `archive.yaml` file.

### Theme data

- Place unstructured theme data in the directory configured by `theme.config.data_dir`.
- Use YAML (`.yaml` or `.yml`), JSON (`.json`), OPML (`.opml`), or CSV (`.csv`) for theme data files.
- Put theme-provided static data, such as placeholder or default values, in that directory.

## Tools

- `hyperctl agent schema describe website` describes the `website` object contract.
- `hyperctl agent schema describe page` describes the `page` object contract.
- `hyperctl website init` creates a website in the current directory.
- `hyperctl website page new` creates a page.

## References

- Consult [information-architecture.md](references/information-architecture.md) as needed to model page families, content collections, feed projections, navigation data, and visibility as authored data contracts.
- Consult [namespaces.md](references/namespaces.md) as needed for namespace paths, nested files, and collisions.
- Consult [feeds.md](references/feeds.md) as needed for website feeds and page feed references.
- Consult [visibility.md](references/visibility.md) as needed for page and content-type visibility.
- [templating](https://hypertemplates.net/docs/#introduction-to-templating)
- [data](https://hypertemplates.net/docs/reference/core/data/)
- [data sources](https://hypertemplates.net/docs/reference/core/data/#template-data-sources)
- [namespaces](https://hypertemplates.net/docs/reference/cms/namespaces/)
- [plugins](https://hypertemplates.net/docs/reference/core/plugins/)
- [computed namespaces](https://hypertemplates.net/docs/reference/cms/namespaces/#computed-namespace-plugins)
- [CMS data](https://hypertemplates.net/docs/reference/cli/commands/cms/data/)
- [CMS pages](https://hypertemplates.net/docs/reference/cli/commands/cms/page/)
- `hyperctl agent skill describe hypertexting-plugins` for computed-data plugin implementation.
