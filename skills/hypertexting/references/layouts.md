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

## Instructions

### Compose layouts

- Model websites as distinct page types, with one layout per type.
- Let the source design choose the entrypoint: begin with a monolithic layout for a full-page mockup or fragments for a component-level design. Converge on reusable fragments with explicit data contracts, composed with `ht-include`.
- Do not split a layout merely to anticipate reuse.
- As a default, immediate children of `<body>` are native landmark or structural elements; immediate children of `<main>` are usually `<section>` or `<article>`. A custom element at either level needs a component-specific reason beyond acting as a page-region styling wrapper.
- Use `<main>` when the document has a distinct primary-content region; a top-level `<section>` or `<article>` may be appropriate when it directly represents the document's content.
- Fragment filenames should describe the fragment's role. Choose a fragment's root element for its use case: a fragment that represents a document region uses the appropriate native structural root, while a reusable presentation component may use a custom-element root. A custom root and filename may naturally align; a semantic root and filename often do not.
- Extract a section once it has a clear data contract or reuse opportunity. Keep it inline while its design or contract is evolving, or when it remains simple and page-specific.
- When top-level section fragments reuse a component, make that component a fragment too.
- Colocate fragment-owned component styles with fragment markup. Include `style#components` then `style#layout` in every layout head; consult [css.md](css.md) for the collector and cascade contract.

### Select directives

- Use `ht-include` for fragments, `ht-if` for conditions, `ht-apply` / `ht-with` / `ht-each` for attribute and content templating, and `ht-block` for content-addressable templating.
- Use `ht-apply` for a target’s attributes and direct text nodes. Use `ht-with` or `ht-each` when bindings must be available throughout a target and its descendants.
- Use `ht-debug`, `ht-pipe`, `ht-attrs`, and `ht-base` only when the simpler directives cannot express the layout.
- Do not reach for JavaScript components until the HTML model is exhausted.
- Use `ht-apply` with template variables instead of legacy `ht-content`, `ht-attr`, `ht-attrs`, `ht-param`, and `ht-query` patterns.

### Verify rendered output

- Use `hyperctl theme render --theme <theme.json> --data <fixture>.json --fragment <fragment>` to develop a rootless component against fixture data. It renders only fragment nodes and does not require page data or inject built-ins.
- Test fragment composition inside a site directory with a throwaway page.
- Validate rendered output as well as build logs, especially when debugging bindings.
