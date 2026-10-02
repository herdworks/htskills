# ht-debug

## Overview

`ht-debug` writes diagnostic information during rendering and removes itself from the rendered HTML. Use it to inspect an element or a template value without creating a separate script.

Within an `ht-each` or `ht-with` scope, it can inspect local bindings. Treat its
output as a view of data available at that directive stage, then validate the
rendered HTML as well.

## Instructions

- Give each `ht-debug` attribute exactly one value: `element` or one complete `${ ... }` template variable.
- Repeated `ht-debug` attributes inspect more than one target on the same element.
- Use `data-ht-debug` when strict HTML attribute naming is required.
- Do not use bare keypaths, comma-separated values, or other modes.

## Examples

```html
<article ht-debug="element" ht-debug='${ page.hero }'>
  <h1 ht-apply>${ page.hero.title }</h1>
</article>
```

The debug log contains the element HTML and the resolved `${ page.hero }` value. The rendered HTML does not retain the debug attributes.

```html
<status-pill
  ht-include="fragments/status-pill"
  ht-with='${ page.hero.pill } as label'
  ht-debug='${ label }'>${ label }</status-pill>
```

Use `ht-debug="element"` to inspect the element at the debug stage.

## Related references

Consult [debugging.md](debugging.md) for the diagnostic workflow and output checks.

## Specification

Consult the [debug directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-debug/) only if this bundled reference does not answer the question.
