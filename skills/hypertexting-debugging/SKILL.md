---
name: hypertexting-debugging
description: Website and template debugging skill.
version: "2026-08-15"
---

# HyperTexting debugging

## Overview

Use HyperTexting's declarative commands and directives before writing diagnostic scripts.

## Workflow

1. Reproduce an isolated component problem with `hyperctl theme render --theme <theme.json> --data <fixture>.json --fragment <fragment>`, or reproduce a composition problem in a site context.
2. Use `ht-debug` to inspect the affected element or complete template variable.
3. Render again, inspect the output and debug log, then run `hyperctl theme check <path>`.
4. Check the owning skill's known issues before investigating expected behavior.
5. Write a custom script only when no command or directive exposes the required information.

## Instructions

### Inspect rendered templates

- Add temporary `ht-debug` attributes to the affected element.
- Use `element` to print the rendered element.
- Use a complete `${ ... }` expression to print a template value.
- Remove debugging attributes after confirming the result.
- Do not write a custom script to determine whether a fragment rendered or what template data is in scope; use `ht-debug` and `hyperctl theme render` first.

### Check build output

- Run `hyperctl theme check <file-or-directory>` after a build to find unrendered `${`
  template artifacts in HTML text and attribute values.
- Pass a single HTML file or a build-output directory; directories are scanned
  recursively for `.html` files.
- The checker reports each completed file as `ok`, `errors`, or `error` and
  continues scanning every selected file after finding template artifacts.
- Use `hyperctl theme check --tidy <path>` for final-output checks that also
  reject remaining HyperTemplates directive attributes.
- Literal examples in direct `pre` and `code` text nodes are intentionally
  ignored by the variable check; their attributes are still checked.

## Tools

- `hyperctl theme render` renders configured theme layouts or rootless fragments with fixture data in any supported template-data format; use a site context when the problem depends on composition or CMS data.
- `hyperctl theme check` validates rendered HTML files and directories for template
  artifacts; `--tidy` also validates that directive attributes were removed.

## References

- Consult [ht-debug.md](references/ht-debug.md) as needed for syntax and examples.
- Consult the owning focused skill's `## Known issues` section as needed before troubleshooting behavior that may be expected.
