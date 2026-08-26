# Directives

## Contract

HyperTemplates uses a fixed set of HTML attributes as templating directives. All templating directives are prefixed with `ht-` (e.g. `ht-include`, `ht-if`, `ht-apply`, `ht-with`, and `ht-each`).

NOTE: HyperTemplates supports `data-` prefixed aliases for all templating directives (e.g. for strict HTML5 compliance). See MDN's [How to use data attributes](https://developer.mozilla.org/en-US/docs/Web/HTML/How_to/Use_data_attributes) guide for more information.

## Syntax

```html
<element directive="expression" directive="expression"></element>
```

* A HyperTemplates directive is an HTML attribute name and value
* The directive attribute value is called an "expression"
* Each directive defines its own expression syntax
* Each directive defines its supported elements
