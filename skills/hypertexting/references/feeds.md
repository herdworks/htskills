# Feeds

## Overview

Define site-wide Atom feeds in `site.feeds` and select a feed for a page through `page.feeds.<alias>`.

## Everything feed

An empty feed configuration selects every feed-visible page.

```yaml
feeds:
  main:
```

It includes public and status pages, uses newest-first ordering, and uses the site URL as its HTML alternate.

## Site feeds

Each key in `site.feeds` is a feed identifier.
An identifier must be non-empty, contain no slash, and must not equal the configured tag namespace (`site.config.tag_path`).

```yaml
feeds:
  blog:
    metadata:
      title: Blog
      description: Recent writing
      author: Example Author
      email: author@example.com
      favicon: /favicon.png
    include:
      - path: blog/**
    exclude:
      - tags: [archived]
    sort:
      - created_at: desc
    html_url: /blog/
```

- Every selector needs at least one of `path` or `tags`. Within one selector,
  `path` and `tags` are alternatives: either match selects the page.
- Any matching `include` admits a page. Omit `include` to admit all pages. Any
  matching `exclude` rejects a page.
- `sort` is an ordered list of rules. Each rule contains exactly one field whose direction is `asc` or `desc`.
- `html_url` is the feed's HTML alternate URL. When omitted, the site URL is the alternate.
- `metadata` overrides Atom feed identity. Its `title`, `description`, `author`, `email`, and `favicon` values inherit the corresponding site title, description, and byline values when omitted. The favicon is emitted on the Atom feed author.
- A configured site feed must not duplicate an automatic tag feed;
  a feed that includes exactly one tag with no excludes, sort, or `html_url` is a validation error.
- User-feed Atom files publish as `/feeds/<id>.xml` by default.
  Automatic tag feeds publish as `/feeds/tags/<tag>.xml` and use `/tags/<tag>/` as their HTML alternate by default.
- HyperTexting generates a tag landing page at `/tags/<tag>/` for each tag feed unless an authored page already occupies that path.
  Generated tag pages use `site.config.tag_layout` as their layout (default `default`).
  On an authored tag page, add a `tag` feed reference when its template needs the feed items;
  for example: `ht-each='page in ${ page.feeds.tag.items }'`.

### Path expressions

Use a slash-separated path expression to match canonical page paths.

- `/blog/**` is root-anchored and matches every page below `/blog/`; `/blog` matches that exact page path.
- `*`, `?`, and character classes such as `[a-z]*` match within one path segment.
- `**` matches any number of complete path segments.
- A bare name matches that name at any depth. A trailing `/` matches every page below that directory.
- Use the feed's `include` and `exclude` lists to express inclusion and exclusion rules.

## Page feed references

Use a page-local alias to choose the feed data that its template renders.

```yaml
feeds:
  main:
  latest:
    source: blog
    limit: 5
  news:
    tag: news
```

- An alias may set `source` or `tag`, but not both. When neither is set, the
  alias is used as the configured `site.feeds` identifier; for example,
  `main:` selects `site.feeds.main`.
- `source`, whether explicit or inferred from the alias, names a configured
  `site.feeds` identifier. An unknown source is a validation error.
- `tag` selects the automatic feed for that tag.
  Unknown tags are treated as empty feeds.
- `limit` is optional and must be non-negative.
  A positive limit caps the exposed items.
- The template receives each selected feed as `page.feeds.<alias>.url` and `page.feeds.<alias>.items`.

## Specification

Consult the website and page schemas with `hyperctl website schema describe website` and `hyperctl website schema describe page` when this reference does not answer a field-level question.
