# Theme bootstrap

## Overview

Themes should be developed alongside a HyperTexting website.
Use an existing website when available, or create a demo website with
`hyperctl website init`. Use `hyperctl theme init` to initialize the new theme.
See `hypertexting-themes` for theme configuration and organization.

## Instructions

### Resolve theme prerequisites

- Inspect the current directory for an existing website and selected theme before creating either one.
- Use an existing website when it has representative content; otherwise create a demo website or ask for a destination directory.
- Initialize or identify the theme root inside that website; consult `hypertexting-themes` before choosing its configuration or organization.
- Ask whether the user has content to use when the website is empty or unrepresentative.
- Offer to generate sample content and fixture data when the user has no content to provide; use it to exercise the theme, not as claimed production content.

## Examples

```text
Existing website and representative content
  -> initialize or select its theme
  -> develop the theme against the existing site
```

```text
No website or no representative content
  -> ask for content or offer generated fixtures
  -> hyperctl website init <website-dir> ...
  -> hyperctl theme init ...
  -> develop the theme against the fixture website
```

## Specification

Consult `hyperctl website init --help` and `hyperctl theme init --help` for
current bootstrap options. Consult `hypertexting-themes` for the theme
configuration and directory contract.
