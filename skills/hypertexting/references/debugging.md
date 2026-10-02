# HyperTexting debugging

## Overview

Use HyperTexting’s declarative commands and directives before writing diagnostic
scripts.

## Instructions

- Reproduce an isolated component problem with `hyperctl theme render --theme <theme.json> --data <fixture>.json --fragment <fragment>`, or reproduce a composition problem in a site context.
- Use `ht-debug` to inspect the affected element or complete template variable; consult [layouts-ht-debug.md](layouts-ht-debug.md) for its sole syntax reference.
- Render again, inspect the output and debug log, then run `hyperctl theme check <path>`.
- Run `hyperctl theme check --tidy <path>` for final-output checks that also reject remaining HyperTemplates directive attributes.
- Pass one HTML file or a build-output directory; directories are scanned recursively and every selected file reports `ok`, `errors`, or `error`.
- `pre` and `code` attributes remain checked, but the checker skips their child subtrees, including syntax-highlighted descendants and their attributes.
- Check the owning reference’s known issues before investigating expected behavior.
- Write a custom script only when no command or directive exposes the required information.

## Examples

Use `hyperctl theme render` for rootless fragments with fixture data, then use a
site context when the problem depends on composition or CMS data.

## Specification

Consult `hyperctl theme check --help` only if this reference does not answer a
tool-contract question.
