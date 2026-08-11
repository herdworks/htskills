# Template data

## Overview

Template data is the render context available to a layout or fragment. The `hypertexting-data` skill defines where that data is authored; this reference defines how layouts read it and how directives scope it.

## Context

- Use `${ ... }` template variables to read `site`, `page`, `theme`, and named data namespaces.
- Use the current page's custom properties through `page`, such as `${ page.hero.title }`; use `site` for website-wide values.
- Use a template-variable fallback when more than one context can supply a value: `${ page.title, site.title, "Untitled" }`.
- Use `page.feeds.<alias>.url` for a selected feed's Atom URL and `page.feeds.<alias>.items` for its page-item views. The page author selects each alias with `page.feeds`.
- Use `site.pages` for the complete build-visible page catalog. It remains complete during an incremental build, so layouts and computed data can derive navigation, archives, and indexes from one stable catalog.
- Status pages are present in `site.pages`. The default status layout resolves
  its redirect from `page.redirect`, then `site.config.status_redirect`, then
  `/`; a custom status layout may render ordinary content.
- See [variables.md](variables.md) for variable syntax and [hypertexting-data](../../hypertexting-data/SKILL.md) when creating or restructuring source data.

## Directive scope

- `ht-apply` and a standalone `ht-with` render only the target element's attributes and direct text-node children. A standalone `ht-with` causes an implicit `ht-apply`; its bindings do not reach descendant elements.
- `ht-each` creates a scoped clone. Its loop binding, `each.index`, `each.count`, `each.total`, and any `ht-with` bindings are available throughout that clone, including an included fragment.
- Bindings are source-ordered and may shadow an outer value; they do not escape their directive scope.
- `ht-block` provides its placeholder attributes as `block` and child HTML as `block.content`; see [ht-block.md](ht-block.md).

## Fragment contracts

Choose the directive that matches where the fragment consumes its input:

- Use `ht-with` on an `ht-include` only when the included fragment's root element consumes one value in its own attributes or direct text. This is root-shallow rendering, not a way to pass data to descendant elements.
- Use `ht-each` on an `ht-include` when the included fragment needs a binding anywhere in its markup. A list renders one included fragment per item; a plain object renders one fragment with that object as the binding.
- Name the loop binding after the fragment contract, such as `heading` or `stat`, and add a brief comment at the top of a parameterized fragment naming that binding.

```html
<!-- Root-only label in fragments/status-pill.html -->
<status-pill ht-include="fragments/status-pill" ht-with='${ page.hero.pill } as label'></status-pill>

<!-- `stat` is available throughout fragments/stat-item.html -->
<stat-item ht-each='stat in ${ page.hero.stats }' ht-include="fragments/stat-item"></stat-item>
```

## Related references

- Consult [ht-include.md](ht-include.md), [ht-each.md](ht-each.md), and [ht-apply.md](ht-apply.md) as needed for directive-specific syntax.
- Consult the [template-data sources documentation](https://hypertemplates.net/docs/reference/core/data/#template-data-sources) only if this bundled reference does not answer the question.
