# Directive expressions

## Contract

A directive expression is the value of an `ht-*` attribute. Each directive defines which tokens and token order it accepts; do not infer one directive's grammar from another.

## Terminology

- A **variable** is one complete `${ ... }` token. See [variables.md](variables.md).
- An **operator** is a directive-defined keyword or symbol, such as `in`, `on`, `to`, `as`, or `==`. Operators are exact and case-sensitive.
- A **literal value** is a non-variable operand. A quoted literal preserves spaces; whether quoted literals are allowed depends on the directive.
- A **delimiter** separates expression segments rather than acting on operands. `OR` is the `ht-if` delimiter.

## Rules

- Whitespace separates tokens except inside a variable or quoted literal.
- Multi-word operators, such as `not in`, allow whitespace between their words.
- Follow the target directive's syntax exactly, including whether it accepts variables, literals, operators, or delimiters.

## Examples

```html
<li ht-each='item in ${ page.items }'>${ item }</li>
<a ht-if='${ page.kind } == "article"'>Read more</a>
```
