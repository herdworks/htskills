# ht-with

## Contract

`ht-with` resolves a complete [template variable](variables.md) and gives it a local binding name. Its syntax is exactly `${ variable } as binding`.

`ht-with` is the single-pass variant of `ht-each`. Both directives bind values
and make them available while processing the target and its descendants,
including included fragments. `ht-with` binds and renders exactly once: arrays,
maps, scalars, and `nil` each bind as one value.

Repeat `ht-with` attributes to create source-ordered bindings. Later bindings may use earlier bindings and may shadow outer values, including with `nil`. Bindings never escape their directive scope.

## Examples

Use `ht-with` when one target and its descendants consume the value:

```html
<a ht-with='${ page.hero.link } as link' href='${ link.href }'>
  ${ link.label }
  <span ht-if='${ link.external }'>↗️</span>
</a>
```

Use `ht-with` on an `ht-include` to align source data with the included fragment’s expected bindings:

```html
<stat-item
  ht-include='fragments/stat-item'
  ht-with='${ data.stats.summary } as stat'></stat-item>
```

Use `ht-each` on an `ht-include` when the target must repeat for every item:

```html
<stat-item
  ht-include='fragments/stat-item'
  ht-each='stat in ${ data.stats.details }'></stat-item>
```

## Related references

- `ht-with` implicitly applies template variables in its root's attributes and direct text-node children; it may also include an explicit `ht-apply`.
- `ht-with` cannot be combined with `ht-each` on one element.
- Consult [ht-each.md](ht-each.md) for the per-resolved-value variant, loop metadata, and iteration syntax.
- Consult [template-data.md](template-data.md) when choosing between target-only application and scoped binding that traverses descendants.

## Specification

Consult the [apply directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-apply/) or [iteration directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-each/) only if these bundled references do not answer the question.
