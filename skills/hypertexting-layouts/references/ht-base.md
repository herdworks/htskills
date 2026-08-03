# ht-base

## Contract

`ht-base` resolves relative values of a target attribute against a rendered base URL. It applies to the directive element and eligible descendants. See [expressions.md](expressions.md) for grammar terminology.

## Syntax

```html
<main ht-base='href on ${ page.canonical_url }'>...</main>
```

- Each expression is exactly `<attribute> on ${ variable }`.
- Repeat the directive for distinct target attributes, such as `href` and `src`; duplicate targets are invalid.
- Absolute URLs and `data:` URLs remain unchanged.
- Legacy `ht-base` forms are not migrated automatically; use the current grammar.

## Examples

```html
<main ht-base='href on ${ page.canonical_url }' ht-base='src on ${ page.canonical_url }'>
    <a href="about">About</a>
    <img src="images/logo.svg" alt="Example">
</main>
```

## Specification

Consult the [base directive reference](https://hypertemplates.net/docs/reference/core/directives/ht-base/) only if you encounter a question this bundled reference does not answer.
