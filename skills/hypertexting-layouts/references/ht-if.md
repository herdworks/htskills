# ht-if

## Contract

`ht-if` keeps an element when its conditional expression is true. A false condition hides the element, or removes it when tidy mode is enabled. Read [variables.md](variables.md) and [expressions.md](expressions.md) for shared terminology.

## Syntax

```html
<article ht-if='${ page.published }'>...</article>
```

### Supported conditions

- **Presence:** `${ page.published }`
- **Negated presence / unary negation:** `!${ page.draft }`
- **Binary equality and inequality:** `${ page.kind } == "article"` or `${ page.kind } != "draft"`
- **Binary regular-expression match:** `${ page.path } =~ "^/blog/"` or `${ page.path } !~ "^/admin/"`
- **Membership:** `${ page.kind } in ${ site.allowed_kinds }`
- **Negated membership:** `${ page.kind } not in ${ site.blocked_kinds }`

Unquoted comparison operands are literal values; use `${ ... }` for data lookups. Use uppercase `OR` between alternative statements. Evaluation is lazy, but every statement must be syntactically valid. The runtime also accepts `=!`, `~=`, and `~!` as alternate spellings of `!=`, `=~`, and `!~`; use the canonical spellings above in new templates.

### Repeat the directive for AND

Repeated `ht-if` attributes on one element are AND conditions: every attribute must match.

## Examples

```html
<article ht-if='${ page.published } OR ${ site.preview }'>...</article>
<a ht-if='${ page.kind } == "article"' ht-apply href="${ page.href }">Read more</a>
<a ht-if='${ page.published }' ht-if='${ page.kind } == "article"' ht-apply href="${ page.href }">Read more</a>
```

## Specification

Consult the [conditional directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-if/) only if you encounter a question this bundled reference does not answer.
