# ht-apply

## Contract

`ht-apply` renders [template variables](variables.md) in an element's own attributes and direct text-node children. It does not traverse or mutate descendant elements; add `ht-apply` to a descendant that needs rendering.

`ht-with` bindings have the same shallow scope. A binding on an element is available only to that element's attributes and direct text-node children; it is not available to descendant directives. Template variables in descendant text remain literal until that descendant has its own rendering directive and scope.

## Syntax

```html
<a ht-apply href="${ page.href }">${ page.title }</a>
```

- Repeat `ht-with='${ ... } as binding'` to create source-ordered local bindings. Later bindings may use earlier ones and may shadow outer values, including with `nil`. Bindings never escape the directive scope. Repeating `ht-apply` has no additional effect.
- A standalone `ht-with` is an alias for an implicit `ht-apply`.
- Use `data-ht-apply` and `data-ht-with` when strict HTML attribute naming is required.

## Examples

```html
<a ht-apply ht-with='${ page.author } as author' href="${ author.href }">${ author.name }</a>
<p ht-with='${ page.summary } as summary'>${ summary }</p>
```

## Related references

Consult [ht-with.md](ht-with.md) for binding syntax, source order, and the distinction between `ht-apply` and `ht-each` scope.

## Specification

Consult the [apply directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-apply/) only if you encounter a question this bundled reference does not answer.
