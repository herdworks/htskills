# Information architecture as data

## Overview

Model information architecture as the website's authored content contract:
page families, content collections, browse paths, feed projections, taxonomy,
navigation data, and visibility. Keep visual hierarchy, component placement,
and navigation styling in the owning layout, HTML, and CSS skills.

HyperTexting builds template data from website and page files plus namespaced
data files on disk. Add custom properties, including arbitrary data structures,
to website and page objects when their names do not conflict with schema-defined
fields.

## Instructions

### Define page families and collections

- Identify each content family the website needs to author and browse.
- Decide whether a family needs a landing page, an index, a reusable detail page, or more than one of those surfaces.
- Define a named `site.feeds` collection for every reusable, ordered group of pages that more than one page or layout needs to consume.
- Define tag feeds, archives, search, and additional detail types only when the content model and expected browsing behavior need them.
- Keep page-specific structured content in that page's data file; use a namespace only when multiple pages or layouts need the same derived or unstructured data.

### Author website and page data

- Create a `site.yaml`, `site.yml`, or `site.json` configuration file and treat it as a `website` object; use `hyperctl website schema describe website` for field-level questions.
- Create pages as directories containing `index.md`, `index.yaml`, `index.json`, `page.md`, `page.yaml`, or `page.json`; treat each as a `page` object.
- Use `index.html` or `page.html` instead of a page data file for a pre-rendered static page.
- Prefer YAML or JSON for structured content and Markdown with YAML front matter for prose.
- Put page-specific structured data, such as `hero` or `features`, directly in its page data file.
- Define named Atom feeds in `site.feeds` only after defining the content families, collections, and browse paths they represent. Consult [feeds.md](feeds.md) for feed identifiers, selectors, and page feed references.

### Use namespace and theme data

- Place unstructured namespace data under `site.config.data_dir` in YAML, JSON, OPML, or CSV files.
- Derive a namespace from the extensionless relative path: `social/twitter.yaml` becomes `social.twitter`.
- Use only letters, digits, and underscores in namespace path segments; do not collide a file with a directory or a static file with a computed `.js` file at the same extensionless path.
- Define a computed namespace only when existing source data needs a reusable derived shape; use a template-value plugin when one layout needs one deterministic formatted value.
- Implement computed-data plugins only after defining the namespace and data contract. Keep inputs and outputs narrow, explicit, and serializable.
- Place theme-provided static data in `theme.config.data_dir` using the same supported static formats.

#### Page visibility

- Set `page.visibility` directly, or `spec.visibility` when a content type owns the choice. Feed eligibility and build visibility are separate choices.

| Visibility | Ordinary website | Feeds | Build | Preview build |
| --- | --- | --- | --- | --- |
| omitted, `default`, `public` | Yes | Yes | Yes | Yes |
| `unlisted`, `website` | Yes | No | Yes | Yes |
| `status`, `feed`, `story`, `chat`, `message`, `response`, `reply` | No | Yes | Yes | Yes |
| `draft`, `preview` | No | No | No | Yes |
| `none`, `hidden` | No | No | No | No |

- `site.pages` contains public, unlisted, status, and generated tag pages; preview builds also include drafts.
- Status pages use `site.config.status_layout`, not page `layout`. The default redirects to `page.redirect`, then `site.config.status_redirect`, then `/`; a custom status layout may render ordinary content.
- Set `page.redirect` to immediately redirect a rendered page.

### Project collections into pages

- Use `page.feeds.<alias>` to expose a complete collection, a limited subset, or a tag feed to the page that renders it.
- Set a feed's `include`, `exclude`, and `sort` rules according to the collection's content contract, not the current number of entries.
- Set `html_url` when the collection's canonical browse page differs from the site root.
- Use an authored tag page only when its layout needs custom content or presentation; otherwise use the generated tag landing page.

### Keep navigation and publishing data explicit

- Store navigation destinations as website or page data when a layout needs to render them; let the layout decide their placement and styling.
- Set `page.visibility` or `spec.visibility` deliberately for public, unlisted, feed-only, draft, preview, and hidden content.
- Treat feed eligibility and build visibility as separate choices.

## Examples

### Editorial website

A journal needs post detail pages and a browse page. Define `site.feeds.blog`
to select and order posts, then expose it to `/blog/` through
`page.feeds.latest`. Add tag feeds when readers need topic-based browsing and
an archive namespace only when the archive requires a reusable grouped shape.

### Services website

Model services as ordinary pages or a distinct page family when each service
needs its own depth. Add a case-study collection only when the website has
case studies to browse; use a feed-backed index page when that collection is
present. Keep the contact destination as page data and let the layout decide
whether it appears in primary or footer navigation.

## Specification

Consult [feeds.md](feeds.md) for `site.feeds` and `page.feeds` syntax. Use
`hyperctl website schema describe website` and `hyperctl website schema describe
page` when the references do not answer a field-level question.
