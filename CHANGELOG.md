# Changelog

All notable changes to the HyperTexting Agent Skills bundle are documented in
this file.

## Unreleased

- **Changed:** Made `Skill.Version` a required opaque string, enabling
  lexicographically ordered ISO-style skill revisions such as `2026-08`.
- **Added:** Documented declarative site and page feeds, automatic tag feeds
  and landing pages, visibility groups, status-page layouts, and redirects.
- **Changed:** Updated conditional-template examples to use supported page
  data, including `page.visibility`.

## v0.2.0 (2026-08-03)

- **Changed:** Documented pull-based `ht-pipe` destinations, root-document collection, source removal, and rendered-element deduplication.
- **Changed:** Required ordered component and layout style blocks so layout-specific CSS overrides component defaults through the normal cascade.
- **Changed:** Allowed full-page and component-first layout development workflows to converge on the same reusable fragment structure.
- **Added:** Documented rootless component rendering with `hyperctl dev render` and fixture JSON data.
- **Changed:** Aligned HTML, CSS, JavaScript, debugging, and plugin verification guidance with component rendering and full-site build boundaries.

## v0.1.0 (2026-08-03)

Initial implementation of the `htskills` package.

- Added `hypertexting` HyperTexting orchestration skill
- Added `hypertexting-data` website data management skill
- Added `hypertexting-themes` website theme management skill
- Added `hypertexting-layouts` HTML templating skill
- Added `hypertexting-debugging` website and template debugging skill
- Added `hypertexting-assets` website asset management skill
- Added `hypertexting-plugins` website plugins skill
- Added `hypertexting-html` general HTML authoring guidelines
- Added `hypertexting-css` general CSS authoring guidelines
- Added `hypertexting-javascript` general JavaScript authoring guidelines
- Added `htskills.FS` for embedding skills in `hyperctl`
- Added `htskills.Skill` type and helper methods for skill frontmatter parsing
