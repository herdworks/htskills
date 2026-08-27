# Changelog

All notable changes to the HyperTexting Agent Skills bundle are documented in
this file.

## Unreleased

- **Added:** Documented `${ markdown(...) }` inline/block behavior and new
  options object.

## v0.9.0 (2026-08-26)

- **Breaking:** Reframed `ht-with` as the single-pass scoped-binding variant of
  `ht-each`. `ht-with` binds and renders exactly once, while making its values
  available as the target, descendants, and included fragments are processed;
  `ht-each` binds and renders once per resolved value. `ht-apply` remains
  target-only rendering and does not traverse descendants.
- **Added:** Documented `site.feeds.<id>.metadata` overrides for Atom titles,
  descriptions, authors, emails, and favicons, including inheritance from the
  website values.
- **Breaking:** Documented canonical Atom feed URLs ending in `.xml` instead
  of `.atom.xml`, for both user feeds and automatic tag feeds.

## v0.8.0 (2026-08-18)

- **Added:** A `compatibility` declaration to bundled skill frontmatter; the
  catalog now requires it and every bundled skill declares `Requires hyperctl.`
- **Added:** Data-oriented information-architecture guidance for page families,
  collections, feed projections, navigation data, and visibility.
- **Added:** CSS references for color schemes, responsive grids, typography,
  and component styles.
- **Added:** Reusable `htskills/cli` Clic command tree and standalone `htskills`
  executable with `readme`, `ls`, `describe`, `install`, `update`, and `up`
  commands.
- **Changed:** Reworked `hypertexting` into a focused-skill router with a
  mandatory build and rendered-output validation gate, plus conditional
  component rendering and browser preview.
- **Changed:** Updated theme packaging guidance to verify generated ZIP archives
  with `unzip -t`.
- **Changed:** `up [skills-directory]` synchronizes bundled skills into every
  discovered project-local skill root, rejects global Claude and Codex skill
  directories, and reports globally installed HyperTexting Skills as strongly
  discouraged before writing local skills.
- **Breaking:** Removed the bundled `hypertexting-theme-development` workflow;
  its concerns are now owned by the focused skills.
- **Breaking:** Moved the onboarding README from `htskills.FS` and
  `htskills.Readme` to the templatable `htskills/readme` package's `README()`
  function. `htskills.FS` now contains only bundled skills.

## v0.7.1 (2026-08-17)

- **Added:** Design-asset management that preserves customer-provided inputs
  separately from web-delivered static assets, with a documented website-root
  design directory.
- **Changed:** Routed theme bootstrap and review through the design-asset
  workflow when users provide new inputs.
- **Changed:** Aligned skill descriptions so HyperTexting-specific skills
  identify their product or website scope, while HTML, CSS, and JavaScript
  remain general authoring guidance.

## v0.7.0 (2026-08-17)

- **Added:** `hypertexting-theme-development`, a visual-direction-to-theme
  workflow for complete, content-agnostic themes: requirements, design
  systems, author customization, implementation, and iterative review.
- **Changed:** Routed new-theme and redesign work from the HyperTexting
  orchestration skill to `hypertexting-theme-development`; narrowed `hypertexting-themes`
  to theme organization and lifecycle operations.
- **Changed:** Updated every bundled skill revision to `2026-08-17`.

## v0.6.0 (2026-08-15)

- **Added:** Documented `hyperctl theme check` as the standard rendered-output
  validation step in the HyperTexting debugging skill, including `--tidy`
  directive artifact checks.
- **Changed:** Documented implicit `page.feeds` sources: an alias without
  `source` or `tag` selects the same-named configured site feed.

## v0.5.0 (2026-08-13)

- **Changed:** Updated the bundled skills and agent onboarding to use
  Hyperctl's `website` and `theme` command tree, and set every skill revision
  to `2026-08-13`.

## v0.4.0 (2026-08-12)

- **Breaking:** Renamed `Skills()` to `Catalog()` and introduced the `Skills`
  collection type with `ColumnValues` for catalog listings.
- **Changed:** Updated every bundled skill revision to `2026-08-12`.

## v0.3.0 (2026-08-11)

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
- **Added:** Documented rootless component rendering with `httt render` and fixture JSON data.
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
