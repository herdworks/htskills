# ht-include

## Contract

The `ht-include` [directive](layouts.md) replaces an eligible element with an HTML fragment's root from the configured theme.

## Syntax

```html
<element ht-include="fragments/header"></element>
```

- `ht-include` expressions are one theme-relative fragment path.
- The `.html` extension is optional: `fragments/header` and `fragments/header.html` are equivalent.
- Use `ht-include` when strict HTML attribute naming is not required.
- Use `data-ht-include` when strict HTML attribute naming is required.
- Do not mix `ht-include` and `data-ht-include` across layouts and fragments within a given theme.
- Do not use `ht-include` on `<html>`, `<head>`, or `<body>` elements.
- The fragment root determines the final document structure. Use a matching placeholder to show the expected contract, but changing only the placeholder does not change the included fragment's root.
- Fragment filenames describe their roles, not their root elements. When an included fragment represents a document region, its root should be the appropriate native structural element.
- HTML attributes are forwarded to the included fragment's first element node;
  _NOTE: `ht-include` and `ht-block` attributes are excluded from attribute forwarding_.

## Examples

```html
<!-- Layout snippet with example include directive -->
<header ht-include="fragments/header" class="site-header"></header>
```

```html
<!-- Example fragments/header.html fragment -->
<header><a href="/">Example</a></header>
```

```html
<!-- Rendered result with forwarded attribute -->
<header class="site-header"><a href="/">Example</a></header>
```

```html
<!-- A document-region fragment uses its semantic root even when its filename differs. -->
<!-- Layout -->
<section ht-include="fragments/home-section"></section>

<!-- fragments/home-section.html -->
<section>
  <home-section variant="feature">...</home-section>
</section>
```

## Specification

Consult the `ht-include` reference documentation only if you encounter a question this bundled reference does not answer:

https://hypertemplates.net/docs/reference/core/directives/ht-include/
