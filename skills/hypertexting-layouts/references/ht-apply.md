# ht-apply

## Contract

`ht-apply` renders [template variables](variables.md) in an element's own attributes and direct text-node children. It does not traverse or mutate descendant elements; add `ht-apply` to a descendant that needs rendering.


## Syntax

```html
<a ht-apply href="${ page.href }">${ page.title }</a>
```

- Repeating `ht-apply` has no additional effect. Use `data-ht-apply` when strict HTML attribute naming is required.

## Examples

```html
<a ht-apply href="${ page.author.href }">${ page.author.name }</a>
```

## Related references

Consult [ht-with.md](ht-with.md) and [ht-each.md](ht-each.md) for the scoped-binding variants that traverse descendants.

## Specification

Consult the [apply directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-apply/) only if you encounter a question this bundled reference does not answer.
