# ht-each

## Contract

`ht-each` clones an element once for each value from an explicit [template-variable](variables.md) source. Arrays and slices produce one clone per item; a map, string, or scalar produces one clone.

## Syntax

```html
<li ht-each='item in ${ page.items }'>${ item }</li>
```

- The expression is exactly `<binding> in ${ variable }`.
- `each.index`, `each.count`, and `each.total` are available inside each clone.
- `ht-offset` skips leading values; `ht-limit` caps the number of rendered values. Both take integer attribute values.
- Repeat `ht-with='${ ... } as binding'` to make source-ordered local bindings. These resolve after the loop item and `each.*` metadata exist. Do not repeat `ht-each`, `ht-offset`, or `ht-limit` on one element.

## Examples

```html
<article ht-each='post in ${ page.posts }' ht-offset="1" ht-limit="3" ht-with='${ post.author } as author'>
    <h2>${ post.title }</h2>
    <p>${ author.name }</p>
</article>
```

## Related references

Consult [ht-with.md](ht-with.md) for binding syntax and the difference between `ht-each` clone scope and `ht-apply`'s shallow scope.

## Specification

Consult the [iteration directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-each/) only if you encounter a question this bundled reference does not answer.
