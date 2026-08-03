# ht-pipe

## Contract

`ht-pipe` moves an element or appends its text to the first matching destination in the document. Use it when a layout cannot place content at its destination directly. See [expressions.md](expressions.md) for grammar terminology.

## Syntax

```html
<style ht-pipe='to "style#components" as css'>...</style>
```

- The selector is a required quoted literal: `to "<selector>"`.
- Omit `as <type>` to move the element. `text`, `css`, `js`, and `javascript` append text instead.
- Template variables in the selector remain literal text; selectors are never rendered.
- Repeated directives are source-ordered fallbacks: the first one with a matching destination wins.

## Examples

```html
<style ht-pipe='to "style#components" as css'>
  status-pill { display: inline-block; }
</style>
```

## Known issues

`ht-pipe` does not deduplicate piped content. Repeated inclusion can emit repeated CSS or text; treat that output as expected behavior.

## Specification

Consult the [pipe directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-pipe/) only if you encounter a question this bundled reference does not answer.
