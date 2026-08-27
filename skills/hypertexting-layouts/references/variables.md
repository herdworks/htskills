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

## Markdown values

Use `markdown(...)` to render a template-data value as Markdown:

```html
<h1 ht-apply>${ markdown(page.title, site.title, "Untitled") }</h1>
```

By default, `markdown(...)` automatically unwraps its output when the rendered Markdown is exactly one paragraph. This makes short Markdown values suitable inside an element that already supplies the block structure:

```html
<h1 ht-apply>${ markdown(page.title) }</h1>
```

For `page.title: Hello **World**`, the rendered content is `Hello <strong>World</strong>`, not a nested `<p>` element. Markdown that produces multiple blocks keeps its block markup.

Pass one final strict-JSON options map with `"display": "block"` to preserve a sole paragraph wrapper:

```html
<div ht-apply>${ markdown(page.description, { "display": "block" }) }</div>
```

Use `display: "block"` only when the surrounding markup expects Markdown block output. `display: "inline"` is the default and does not need to be specified.

## Specification

Consult the [template-variable reference](https://hypertemplates.net/docs/reference/core/variables/) only if you encounter a question this bundled reference does not answer.
