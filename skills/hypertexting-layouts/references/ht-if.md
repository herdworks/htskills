# ht-if

## Contract

`ht-if` keeps an element when its conditional expression is true. A false condition hides the element, or removes it when tidy mode is enabled. Read [variables.md](variables.md) and [expressions.md](expressions.md) for shared terminology.

## Syntax

```html
<article ht-if='${ page.title }'>...</article>
```

### Supported conditions

- **Presence:** `${ page.title }`
- **Negated presence / unary negation:** `!${ page.more }`
- **Binary equality and inequality:** `${ page.visibility } == "draft"` or `${ page.visibility } != "draft"`
- **Binary regular-expression match:** `${ page.path } =~ "^/blog/"` or `${ page.path } !~ "^/admin/"`
- **Membership:** `${ page.category } in ${ site.featured_categories }`
- **Negated membership:** `${ page.category } not in ${ site.hidden_categories }`

Unquoted comparison operands are literal values; use `${ ... }` for data lookups. Use uppercase `OR` between alternative statements. Evaluation is lazy, but every statement must be syntactically valid. The runtime also accepts `=!`, `~=`, and `~!` as alternate spellings of `!=`, `=~`, and `!~`; use the canonical spellings above in new templates.

### Repeat the directive for AND

Repeated `ht-if` attributes on one element are AND conditions: every attribute must match.

## Examples

```html
<article ht-if='${ page.visibility } == "public" OR ${ page.visibility } == "unlisted"'>...</article>
<a ht-apply ht-if='${ page.more }' href="${ page.path }">Read more</a>
<a ht-apply ht-if='${ page.more }' ht-if='${ page.visibility } == "public"' href="${ page.path }">Read more</a>
```

## Specification

Consult the [conditional directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-if/) only if you encounter a question this bundled reference does not answer.
