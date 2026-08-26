# ht-debug

## Contract

`ht-debug` writes diagnostic information during rendering and removes itself from the rendered HTML. Give each `ht-debug` attribute exactly one value: `element` or one complete `${ ... }` template variable. Repeat the attribute to inspect more than one target on an element.

Within an `ht-each` or `ht-with` scope, `ht-debug` can inspect the local bindings. Treat its output as a view of the data available at its directive stage, then validate the rendered HTML as well.

## Examples

```html
<status-pill
  ht-include="fragments/status-pill"
  ht-with='${ page.hero.pill } as label'
  ht-debug='${ label }'>${ label }</status-pill>
```

The debug output includes the resolved `${ label }` value. Use `ht-debug="element"` to inspect the element at the debug stage.

## Syntax

- Use `data-ht-debug` when strict HTML attribute naming is required.
- Do not use bare keypaths, comma-separated values, or other modes.

## Related references

Consult the [debugging skill's ht-debug reference](../../hypertexting-debugging/references/ht-debug.md) for the general debugging workflow and output examples.

## Specification

Consult the [debug directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-debug/) only if these bundled references do not answer the question.
