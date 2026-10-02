# Template data

## Overview

Template data is the render context available to a layout or fragment. [data.md](data.md) defines where that data is authored; this reference defines how layouts read it and how directives scope it.

## Context

- Use `${ ... }` template variables to read `site`, `page`, `theme`, and named data namespaces.
- Use the current page's custom properties through `page`, such as `${ page.hero.title }`; use `site` for website-wide values.
- Use a template-variable fallback when more than one context can supply a value: `${ page.title, site.title, "Untitled" }`.
- Use `page.feeds.<alias>.url` for a selected feed's Atom URL and `page.feeds.<alias>.items` for its page-item views. The page author selects each alias with `page.feeds`.
- Use `site.pages` for the complete build-visible page catalog. It remains complete during an incremental build, so layouts and computed data can derive navigation, archives, and indexes from one stable catalog.
- Status pages are present in `site.pages`. The default status layout resolves
  its redirect from `page.redirect`, then `site.config.status_redirect`, then
  `/`; a custom status layout may render ordinary content.
- See the template-variable rules below and [data.md](data.md) when creating or restructuring source data.

## Directive scope

- `ht-apply` renders only the target element's attributes and direct text-node children; it does not traverse descendants or create a local binding.
- `ht-each` and `ht-with` bind values and make them available while processing the target and its descendants. `ht-each` binds and renders once per resolved value, including `each.index`, `each.count`, and `each.total`; `ht-with` binds and renders exactly once. Their bindings are available in included fragments.
- Repeated `ht-with` bindings are source-ordered and may shadow an outer value; they do not escape their directive scope.
- `ht-block` provides its placeholder attributes as `block` and child HTML as `block.content`; see [layouts-ht-block.md](layouts-ht-block.md).

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

- Consult [layouts-ht-include.md](layouts-ht-include.md), [layouts-ht-apply.md](layouts-ht-apply.md), [layouts-ht-with.md](layouts-ht-with.md), and [layouts-ht-each.md](layouts-ht-each.md) as needed for directive-specific syntax.
- Consult the [template-data sources documentation](https://hypertemplates.net/docs/reference/core/data/#template-data-sources) only if this bundled reference does not answer the question.

## Template variables

A template variable is a complete `${ ... }` expression. Directives that require a variable require the complete form; a bare keypath such as `page.title` is a literal value, not a data lookup.

- Use keypaths such as `${ page.title }` to read template data.
- Separate fallback candidates with commas: `${ page.title, site.title, "Untitled" }`.
- Unquoted arguments are keypaths; quoted strings, booleans, and finite numbers are literal values.
- A variable resolves to its native value where a directive expects a value, or renders as text where `ht-apply`, `ht-with`, or `ht-each` processes it.
- Use `markdown(...)` to render a template-data value as Markdown. It unwraps a sole paragraph by default; pass a final strict-JSON `{ "display": "block" }` options map only when surrounding markup expects block output.

```html
<a ht-apply href="${ page.href, site.url }">${ page.title, "Untitled" }</a>
```

## Directive expressions

A directive expression is the value of an `ht-*` attribute. Each directive defines which tokens and order it accepts; do not infer one directive’s grammar from another.

- A **variable** is one complete `${ ... }` token.
- An **operator** is a directive-defined keyword or symbol, such as `in`, `on`, `to`, `as`, or `==`; operators are exact and case-sensitive.
- A **literal value** is a non-variable operand. A quoted literal preserves spaces when that directive permits it.
- A **delimiter** separates expression segments rather than acting on operands; `OR` is the `ht-if` delimiter.
- Whitespace separates tokens except inside a variable or quoted literal. Follow the target directive’s syntax exactly.
