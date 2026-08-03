# ht-with

## Contract

`ht-with` resolves a complete [template variable](variables.md) and gives it a local binding name. Its syntax is exactly `${ variable } as binding`.

`ht-with` changes scope according to the directive it modifies:

- On `ht-apply`, the binding is available only to the target element's attributes and direct text-node children.
- On `ht-each`, the binding resolves after the loop item and `each.*` metadata, and is available throughout that iteration's clone, including an included fragment.
- Without either directive, `ht-with` causes an implicit `ht-apply`, so it has `ht-apply`'s shallow scope.

Repeat `ht-with` attributes to create source-ordered bindings. Later bindings may use earlier bindings and may shadow outer values, including with `nil`. Bindings never escape their directive scope.

## Examples

Use `ht-with` when the root element consumes the value:

```html
<status-pill
  ht-include="fragments/status-pill"
  ht-with='${ page.hero.pill } as label'>${ label }</status-pill>
```

Use `ht-each` when the binding must be available throughout an included fragment:

```html
<stat-item
  ht-each='stat in ${ page.hero.stats }'
  ht-include="fragments/stat-item"></stat-item>
```

## Related references

- Consult [ht-apply.md](ht-apply.md) for shallow rendering and root-only bindings.
- Consult [ht-each.md](ht-each.md) for clone scope, loop metadata, and iteration syntax.
- Consult [template-data.md](template-data.md) when choosing between a root-only binding and fragment-wide iteration.

## Specification

Consult the [apply directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-apply/) or [iteration directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-each/) only if these bundled references do not answer the question.
