# ht-each

## Contract

`ht-each` is the per-resolved-value variant of `ht-with`. Both directives bind
values and make them available while processing the target and its descendants,
including included fragments. `ht-each` binds and renders once per resolved
value from an explicit [template-variable](layouts-data.md) source. Arrays and
slices produce one rendered element per item; a map, string, or scalar produces
one rendered element; `nil` renders no element.

## Syntax

```html
<li ht-each='item in ${ page.items }'>${ item }</li>
```

- The expression is exactly `<binding> in ${ variable }`.
- `each.index`, `each.count`, and `each.total` are available for each rendered element.
- `ht-offset` skips leading values; `ht-limit` caps the number of rendered values. Both take integer attribute values.
- Do not combine `ht-each` with `ht-with` on one element, and do not repeat `ht-each`, `ht-offset`, or `ht-limit`.

## Examples

```html
<article ht-each='post in ${ page.posts }' ht-offset="1" ht-limit="3">
	<h2>${ post.title }</h2>
	<p>${ post.author.name }</p>
</article>
```

## Related references

Consult [layouts-ht-with.md](layouts-ht-with.md) for the single-pass variant, which binds and
renders exactly once.

## Specification

Consult the [iteration directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-each/) only if you encounter a question this bundled reference does not answer.
