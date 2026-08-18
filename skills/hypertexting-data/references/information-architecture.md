# Information architecture as data

## Overview

Model information architecture as the website's authored content contract:
page families, content collections, browse paths, feed projections, taxonomy,
navigation data, and visibility. Keep visual hierarchy, component placement,
and navigation styling in the owning layout, HTML, and CSS skills.

## Instructions

### Define page families and collections

- Identify each content family the website needs to author and browse.
- Decide whether a family needs a landing page, an index, a reusable detail page, or more than one of those surfaces.
- Define a named `site.feeds` collection for every reusable, ordered group of pages that more than one page or layout needs to consume.
- Define tag feeds, archives, search, and additional detail types only when the content model and expected browsing behavior need them.
- Keep page-specific structured content in that page's data file; use a namespace only when multiple pages or layouts need the same derived or unstructured data.

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

Consult [feeds.md](feeds.md) for `site.feeds` and `page.feeds` syntax, and
[visibility.md](visibility.md) for publishing behavior. Use
`hyperctl agent schema describe website` and `hyperctl agent schema describe
page` when the references do not answer a field-level question.
