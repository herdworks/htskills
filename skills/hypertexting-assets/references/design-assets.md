# Design assets

## Overview

Design assets are customer-provided or approved inputs that inform a website or theme.
They are distinct from static assets delivered by a built website.

## Instructions

- Keep design assets in a dedicated website-root directory (e.g. `design/`, `design-assets/`, or `assets/`).
  NOTE: the design asset directory should always be distinct from the configured static directory.
- Preserve supplied files as received. Do not silently rename, overwrite, optimize, or convert an original.
- Maintain a `README.md` file in the design assets directory to document asset provenance, licensing, and precedence information (e.g. if multiple revisions of an asset have been provided).
- Use an `originals/` subdirectory for files that may become static assets (e.g. logos, fonts, images, videos, graphics, and other artwork).
- Use a `references/` subdirectory for reference materials that should not be promoted to static assets (e.g. mockups, brand guides, etc).
- Use a `modified/` subdirectory for renamed or modified originals (e.g. optimized images with `_WxH` filenames).
- Always copy design assets when promoting them to static assets (e.g. theme, website, or page static assets). Never move design assets.

## Example

```text
design/
  README.md
  originals/
    logos/
    fonts/
    images/
  references/
    mockups/
    brand-guides/
```

```markdown
This directory preserves customer-provided design assets for the website.
Live production assets belong in the configured static directory; this
directory is not published as part of the website.
```

## Specification

Use the configured `theme.config.static_dir`, `site.config.static_dir`, or
page static-asset scope for files delivered by the website.
