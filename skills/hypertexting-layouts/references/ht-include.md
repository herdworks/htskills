# ht-include

## Contract

The `ht-include` [directive](directives.md) replaces an eligible element with an HTML fragment from the configured theme.

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

## Specification

Consult the `ht-include` reference documentation only if you encounter a question this bundled reference does not answer:

https://hypertemplates.net/docs/reference/core/directives/ht-include/
