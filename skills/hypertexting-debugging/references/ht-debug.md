# ht-debug

## Overview

`ht-debug` writes diagnostic information during rendering and removes itself from the rendered HTML. Use it to inspect an element or a template value without creating a separate script.

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

## Specification

Consult the [debug directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-debug/) only if this bundled reference does not answer the question.
