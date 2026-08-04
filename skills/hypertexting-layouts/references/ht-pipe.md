# ht-pipe

## Contract

`ht-pipe` collects matching source elements into the destination that owns the directive. Use it when a layout needs to assemble fragment-owned elements or text in one location. It runs once on the final root document, after `ht-block` and before `ht-base`; it does not run while rendering fragments. See [expressions.md](expressions.md) for grammar terminology.

## Syntax

```html
<style id='components' ht-pipe='from "style.component" as css'></style>
```

- Put `ht-pipe` on the destination. The selector is a required quoted literal: `from "<source selector>"`.
- Mark source elements so they match the selector. `style.component` with `<style class='component'>` is the standard component-style convention, but destinations may use a different or broader source selector when needed.
- Omit `as <type>` to append matching source elements. `text`, `css`, `js`, and `javascript` append their text instead.
- Template variables in the selector remain literal text; selectors are never rendered.
- Repeated directives on one destination are processed in source order. Matching sources are collected in document order and removed from their original parents.
- `ht-pipe` excludes its destination from the source set and deduplicates matching source elements by exact rendered outer-HTML bytes before appending elements or text.

## Examples

```html
<!-- layout -->
<style id='components' ht-pipe='from "style.component" as css'></style>

<!-- fragment -->
<style class='component'>
  status-pill { display: inline-block; }
</style>
```

## Specification

Consult the [pipe directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-pipe/) only if you encounter a question this bundled reference does not answer.
