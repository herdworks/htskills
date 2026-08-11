# Visibility

## Overview

Set `page.visibility` directly in page data, or set `spec.visibility` in a content type when the content type owns that choice. Visibility independently controls ordinary website/Atom-alternate behavior, feed eligibility, and build output.

## Visibility values

| Visibility values | Ordinary website | Feeds | Build | Preview build |
| --- | --- | --- | --- | --- |
| omitted, `default`, `public` | Yes | Yes | Yes | Yes |
| `unlisted`, `website` | Yes | No | Yes | Yes |
| `status`, `feed`, `story`, `chat`, `message`, `response`, `reply` | No | Yes | Yes | Yes |
| `draft`, `preview` | No | No | No | Yes |
| `none`, `hidden` | No | No | No | No |

`site.pages` contains the build-visible catalog: public, unlisted, status, and
generated tag pages; preview builds also include drafts.

## Status pages

Status pages are built and included in `site.pages`. Their layout is selected
by `site.config.status_layout`, not page `layout`. The default status layout
redirects to `page.redirect`, then `site.config.status_redirect`, then `/`.
A custom status layout may render ordinary page content instead.

## Redirects

Set `page.redirect` to immediately redirect any rendered page. Status pages
default it from `site.config.status_redirect`.

## Specification

Consult the page schema with `hyperctl agent schema describe page` when this reference does not answer a field-level question.
