---
name: hypertexting
description: Create, deploy, update, and debug HyperTexting websites and themes. 
compatibility: Requires hyperctl.
metadata:
  version: "2026-09-11T00:40:00Z"
  replaces: hypertexting-assets, hypertexting-css, hypertexting-data, hypertexting-debugging, hypertexting-html, hypertexting-javascript, hypertexting-layouts, hypertexting-plugins, hypertexting-themes
---

# HyperTexting

## Overview

HyperTexting is a data-driven website management system that combines file-based
content management with a static site generator (SSG) to create, deploy, and
update websites.

It uses a pure-HTML templating and theme-management system called HyperTemplates 
to generate websites from a website configuration file, a theme, and a 
repository of data files.

## Workflow

1. From the website workspace root, run `hyperctl skill up` before starting a HyperTexting task to refresh project-local skills.
2. Define content and template contracts, then choose the narrowest owning layer.
3. Render focused layouts or fragments while developing.
4. Build with an explicit identifier and run a tidy output check before handoff.
5. Preview visual or interaction work in a browser and give the user the local URL.

## Instructions

### Choose the implementation layer

- Use HTML and layouts for content, document structure, and deterministic composition.
- Use CSS for presentation, responsive rules, and visual-only transformations.
- Use a build-time plugin for one deterministic rendered value or reusable derived build data.
- Use browser JavaScript for isolated interaction in an already useful component.
- Use an API only for runtime state, mutation, personalization, or an external service.
- Do not turn an entire page into a client application.

### Verify website work

- Render a rootless fragment with `hyperctl theme render --theme <theme.json> --data <fixture>.json --fragment <fragment>` when its contract does not depend on CMS data.
- Reproduce composition and CMS-data behavior in a site context.
- Build with `hyperctl website build --id <build-id>`; pass the predictable identifier directly instead of wrapping the command to capture an automatic identifier.
- Run `hyperctl theme check --tidy <build-output>` after every website build before handoff.
- Inspect rendered output and build logs when bindings or directives behave unexpectedly.

### Preview website work

- Keep a preview server running for the duration of a HyperTexting task when browser preview is useful.
- Before starting a new preview server, check whether `site.config.preview` is configured; if a server is running at `http://<preview.bind or localhost>:<preview.port>` with a `/.well-known/hypertexting/build` that matches `preview.id`, use it.
- Prefer adding `site.config.preview` with a stable ID, `localhost` bind, and an unused project-specific port. Use that configuration as the source of truth for preview-server management.
- Otherwise start `hyperctl website server --config <site-config>` and keep it running for the duration of the session. Do not stop a process that occupies the configured port but does not identify as this website.
- Inspect visual and interaction changes in a browser.
- At handoff, provide the relevant HTTP page URL(s), such as `http://localhost:<port>/example`.

## Tools

- `hyperctl skill up [--force] [skills-directory]` synchronizes project-local skills; `--force` replaces installed skills regardless of their versions.
- `hyperctl skill describe hypertexting` shows this skill after installation.
- `hyperctl skill --help` and `hyperctl skill <command> --help` show skill-command usage.

## References

- Consult [assets.md](references/assets.md) for design assets and static asset scope.
- Consult [data.md](references/data.md) for website, page, namespace, and visibility data; consult [feeds.md](references/feeds.md) for feed contracts.
- Consult [layouts.md](references/layouts.md) for composition and directive selection; consult [layouts-data.md](references/layouts-data.md) for template data, variables, and expressions.
- Consult [layouts-ht-apply.md](references/layouts-ht-apply.md), [layouts-ht-attrs.md](references/layouts-ht-attrs.md), [layouts-ht-base.md](references/layouts-ht-base.md), [layouts-ht-block.md](references/layouts-ht-block.md), [layouts-ht-debug.md](references/layouts-ht-debug.md), [layouts-ht-each.md](references/layouts-ht-each.md), [layouts-ht-if.md](references/layouts-ht-if.md), [layouts-ht-include.md](references/layouts-ht-include.md), [layouts-ht-pipe.md](references/layouts-ht-pipe.md), or [layouts-ht-with.md](references/layouts-ht-with.md) for directive-specific syntax.
- Consult [html.md](references/html.md), [css.md](references/css.md), or [javascript.md](references/javascript.md) for standard website implementation guidance.
- Consult [plugins.md](references/plugins.md) for deterministic build-time data and value transformation.
- Consult [themes.md](references/themes.md) for theme organization, configuration, installation, packaging, and migration.
- Consult [debugging.md](references/debugging.md) for rendered-output and template diagnostics.
