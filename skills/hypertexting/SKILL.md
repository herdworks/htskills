---
name: hypertexting
description: HyperTexting website orchestration skill.
compatibility: Requires hyperctl.
metadata:
  version: "2026-08-18T19:07:00Z"
---

# HyperTexting

## Overview

The HyperTexting skill routes website work using a collection of focused discipline skills for theme, data, and asset management, HTML, CSS, and JavaScript best practices, and a debugging skill for troubleshooting.

## Workflow

1. Always follow the instructions at `hyperctl agent readme` before beginning a new HyperTexting task
2. Use the corresponding hypertexting-* skill(s) for a given task
3. For multi-disciplinary work, proceed through themes, data, and assets; then layouts, HTML, CSS, JavaScript, and plugins.
4. Render isolated layouts or fragments with `hyperctl theme render` for direct verification.
5. Always build websites with `hyperctl website build --id <build-id>` and run `hyperctl theme check --tidy <build-output>` before handoff.
6. For visual or interaction changes, run `hyperctl website server` and inspect the results in a browser.
7. Use `hypertexting-debugging` for unexpected rendered output, directive behavior, or template-data scope.

## Instructions

### Route by concern

- Use `hypertexting-themes` to organize, install, package, switch, or migrate a theme.
- Use `hypertexting-data` for website, page, theme, and namespace data.
- Use `hypertexting-assets` for durable design inputs and static asset scope.
- Use `hypertexting-layouts` for layouts, fragments, and HyperTemplates directives.
- Use `hypertexting-html` for semantic markup and custom-element contracts.
- Use `hypertexting-css` for shared foundations and component styles.
- Use `hypertexting-javascript` for isolated browser-time progressive enhancement.
- Use `hypertexting-plugins` for deterministic build-time data or value transforms.
- Use `hypertexting-debugging` for unexpected rendered output, directive behavior, or template-data scope.

### Preserve the static-first boundary

- Define content and template contracts before implementing a plugin or browser behavior.
- Prefer HTML, CSS, static data, and template directives before JavaScript or an API.
- Pass a predictable build identifier directly: `hyperctl website build --id <build-id>`.
- Do not wrap `hyperctl website build` in a script merely to capture its auto-generated ID.
- Use `hyperctl website server` to preview the website after a meaningful visual or interaction change.

## Tools

- `hyperctl agent skill up <skills directory>` synchronizes project-local skills
- `hyperctl agent skill describe <skill>` shows a named skill's instructions
- `hyperctl` lists all available subcommands; use `-h` or `--help` with any command for help/usage information

## References

- Consult [architecture.md](references/architecture.md) as needed for HyperTexting's static-first architecture and progressive-enhancement model.
