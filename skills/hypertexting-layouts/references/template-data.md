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

- `ht-apply` renders only the target element's attributes and direct text-node children; it does not traverse descendants or create a local binding.
- `ht-each` and `ht-with` bind values and make them available while processing the target and its descendants. `ht-each` binds and renders once per resolved value, including `each.index`, `each.count`, and `each.total`; `ht-with` binds and renders exactly once. Their bindings are available in included fragments.
- Repeated `ht-with` bindings are source-ordered and may shadow an outer value; they do not escape their directive scope.
- `ht-block` provides its placeholder attributes as `block` and child HTML as `block.content`; see [ht-block.md](ht-block.md).

## Fragment contracts

Choose the directive that matches where the fragment consumes its input:

- Use `ht-with` on an `ht-include` when one included fragment needs a bound value anywhere in its markup.
- Use `ht-each` on an `ht-include` when it should render one included fragment per resolved value.
- Name the loop binding after the fragment contract, such as `heading` or `stat`, and add a brief comment at the top of a parameterized fragment naming that binding.

```html
<!-- `label` is available throughout fragments/status-pill.html -->
<status-pill ht-include="fragments/status-pill" ht-with='${ page.hero.pill } as label'></status-pill>

<!-- `stat` is available throughout fragments/stat-item.html -->
<stat-item ht-each='stat in ${ page.hero.stats }' ht-include="fragments/stat-item"></stat-item>
```

## Related references

- Consult [ht-include.md](ht-include.md), [ht-apply.md](ht-apply.md), [ht-with.md](ht-with.md), and [ht-each.md](ht-each.md) as needed for directive-specific syntax.
- Consult the [template-data sources documentation](https://hypertemplates.net/docs/reference/core/data/#template-data-sources) only if this bundled reference does not answer the question.
