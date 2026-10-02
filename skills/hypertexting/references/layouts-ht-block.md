# ht-block

## Contract

`ht-block` replaces an element with the theme block named after that element. A `<callout ht-block>` placeholder resolves the theme's `blocks/callout.html` resource.

## Syntax

```html
<callout ht-block class="featured">Fallback content</callout>
```

- Placeholder attributes are forwarded to the block fragment's first element, except block and include directive attributes.
- The block fragment receives placeholder attributes as `block` and the placeholder's child HTML as `block.content`.
- Blocks run only in the top-level document pass, not inside included fragments.

## Examples

```html
<!-- Layout -->
<callout ht-block class="featured">Welcome</callout>

<!-- blocks/callout.html -->
<aside ht-apply class="${ block.class }">${ block.content }</aside>
```

## Specification

Consult the [block directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-block/) only if you encounter a question this bundled reference does not answer.
