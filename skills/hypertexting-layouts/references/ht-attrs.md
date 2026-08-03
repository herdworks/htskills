# ht-attrs

## Contract

`ht-attrs` applies a string-keyed map as HTML attributes on its element. Prefer `ht-apply` with template variables for ordinary, known attributes; use `ht-attrs` only when the attribute map is genuinely dynamic.

## Syntax

```html
<a ht-attrs='${ page.link.attributes }'>Read more</a>
```

- Each value must be exactly one [template variable](variables.md) resolving to a string-keyed map.
- Repeat the directive to layer maps in source order; later maps win conflicting attribute names.
- Nil map values do not produce an attribute. Directive attribute names are ignored.

## Examples

```html
<a ht-attrs='${ page.link.attributes }' ht-apply href="${ page.link.href }">Read more</a>
```

## Specification

Consult the [attribute-map directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-attrs/) only if you encounter a question this bundled reference does not answer.
