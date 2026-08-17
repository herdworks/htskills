# Responsive grid

## Overview

The responsive grid is the theme-wide geometry system. Define a flexible
system to guide implementation of page presentations, so layouts, fragments,
blocks, and components reuse its geometry rather than creating redundant grid
rules.

## Instructions

- Define a readable content container and any wider media or full-bleed regions.
- Choose a small, flexible column system that supports the documented experience architecture without forcing every section into identical geometry.
- Do not impose a generic one-size-fits-all framework; define only the geometry the theme's layouts, fragments, blocks, and components need.
- Define gutter and vertical-spacing scales that work at narrow and wide widths.
- Use no more than three or four breakpoints unless the supplied visual direction proves another is necessary.
- Define each page presentation's mobile transformation: stack, reorder only when meaning remains clear, scroll, collapse, or simplify.
- Keep primary navigation usable at narrow widths without hiding essential destinations behind JavaScript-only behavior.
- Test text expansion, long navigation labels, wide tables, and oversized media against the grid.

## Examples

| Region | Wide behavior | Narrow behavior |
| --- | --- | --- |
| Page container | centered maximum width | fluid width with side gutters |
| Feature pair | text and evidence in columns | stacked with the claim first |
| Preview collection | multiple columns | fewer columns or one column |
| Primary navigation | inline destinations | disclosed menu with the same destinations |

## Specification

Consult `hypertexting-css` for CSS foundation files, media queries, and component-style mechanics.
