# Data namespaces

## Overview

HyperTexting derives a data namespace from each file's path relative to `site.config.data_dir`. Strip the file extension and convert subdirectories into nested keypath segments.

## Instructions

- Use a filename for the final namespace segment and directories for parent segments.
- Use only letters, digits, and underscores in every filename and directory segment.
- Do not use dots, hyphens, whitespace, or empty segments in a namespace path.
- Do not define two files with the same extensionless relative path, including one static file and one computed `.js` file.
- Do not define both a file and a directory at the same namespace segment.

## Examples

| Relative path under `site.config.data_dir` | Namespace |
| --- | --- |
| `links.yaml` | `links` |
| `social/twitter.yaml` | `social.twitter` |
| `archive.js` | `archive` |
| `social/twitter.js` | `social.twitter` |

For example, `social/twitter.yaml` makes `${ social.twitter }` available to a template. A computed `social/twitter.js` would claim that same namespace, so use one or the other.

## Specification

Consult the [CMS namespaces reference](https://hypertemplates.net/docs/reference/cms/namespaces/) only if this bundled reference does not answer the question.
