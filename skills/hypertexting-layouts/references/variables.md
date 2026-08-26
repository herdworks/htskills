# Template variables

## Contract

A template variable is a complete `${ ... }` expression. Directives that require a variable require the complete form; a bare keypath such as `page.title` is a literal value, not a data lookup.

## Syntax

```html
<h1 ht-apply>${ page.title }</h1>
```

- Use keypaths to read template data, such as `${ page.title }`.
- Separate fallback candidates with commas: `${ page.title, site.title, "Untitled" }`.
- Unquoted arguments are keypaths; quoted strings, booleans, and finite numbers are literal values.
- A template variable can resolve to its native value where a directive expects a value, or render as text where `ht-apply`, `ht-with`, or `ht-each` processes it.

## Examples

```html
<a ht-apply href="${ page.href, site.url }">${ page.title, "Untitled" }</a>
```

## Specification

Consult the [template-variable reference](https://hypertemplates.net/docs/reference/core/variables/) only if you encounter a question this bundled reference does not answer.
