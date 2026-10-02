# Choosing a plugin

## Overview

Choose the smallest layer that satisfies the requirement. Use build-time plugins for deterministic data or value transformation, not for visual styling, semantic markup, or browser interaction.

Plugins are JavaScript modules whose default export is a function. HyperTemplates runs them while building template data or pre-rendering a value.

## Instructions

| Requirement | Use |
| --- | --- |
| Visual-only presentation, such as uppercase text | CSS |
| Semantic element or static document structure | HTML and layouts |
| One deterministic rendered value | Template-value plugin |
| Reusable derived structured data | Computed-data plugin |
| Browser interaction, visitor state, or API request | Web Component and browser JavaScript |

- Use a template-value plugin when a layout needs one value in a specific pre-rendered form.
- Use a computed-data plugin when multiple layouts or pages need a derived object or namespace.
- Keep the semantic element in the layout and return only the value it needs.
- Do not use a plugin when the output varies by visitor locale, current time, interaction, browser state, or API response.

### Plugin contracts

- Implement a template-value plugin only when a value must take a specific deterministic form in pre-rendered output and HTML, CSS, directives, or source data cannot express it.
- Implement a computed-data plugin only after [data.md](data.md) defines a reusable derived namespace and its data contract.
- Derive a computed namespace from the `.js` file’s relative path under `site.config.data_dir`: `social/twitter.js` becomes `social.twitter`.
- Return a rendered value from a template-value plugin or structured data from a computed-data plugin.
- Keep plugin inputs and outputs narrow, explicit, and serializable.
- Keep plugin ownership with the website or theme that owns its contract.
- Keep semantic HTML and markup ownership in layouts; a plugin returns data or a value rather than becoming an alternate HTML authoring system.

## Tools

- `hyperctl website build` verifies build-time plugin output in the full site context.

## Examples

```html
<time ht-apply datetime='${ duration_iso(page.runtime_seconds) }'>
  ${ duration_label(page.runtime_seconds) }
</time>
```

Use template-value plugins for `duration_iso` and `duration_label` when the source duration must appear in a specific static form. If multiple pages need a duration object with `seconds`, `iso`, and `label`, return that object from a computed-data plugin instead.

The plugin filename supplies the template function name:

```js
// plugins/duration_iso.js
export default function durationISO(seconds) {
  const value = Math.max(0, Math.round(Number(seconds) || 0))
  const hours = Math.floor(value / 3600)
  const minutes = Math.floor((value % 3600) / 60)
  const remainingSeconds = value % 60
  const parts = []

  if (hours) parts.push(`${hours}H`)
  if (minutes) parts.push(`${minutes}M`)
  if (remainingSeconds || parts.length === 0) parts.push(`${remainingSeconds}S`)

  return `PT${parts.join("")}`
}
```

```js
// plugins/duration_label.js
export default function durationLabel(seconds) {
  const value = Math.max(0, Math.round(Number(seconds) || 0))
  const minutes = Math.round(value / 60)

  if (minutes < 1) return "Less than a minute"
  if (minutes === 1) return "1 minute"
  return `${minutes} minutes`
}
```

Use CSS, not a plugin, for uppercase text:

```css
article-card[featured] h2 {
  text-transform: uppercase;
}
```

## Specification

Consult the [template-variable plugins](https://hypertemplates.net/docs/reference/core/variables/#template-variable-plugins) or [computed namespaces](https://hypertemplates.net/docs/reference/cms/namespaces/#computed-namespace-plugins) documentation only if this bundled reference does not answer the question.
