# Color schemes

## Overview

A theme's color system is a fixed set of semantic roles, not merely a count of
brand colors. Define those roles before styling page presentations and components so a
theme can remain cohesive while supporting its intended modes and states.

## Instructions

- Name palette values by role: canvas, surface, text, muted text, link, accent, border, and feedback or state colors.
- Choose the smallest palette that expresses the supplied direction and supports required contrast and interaction states.
- Define hover, focus, selected, disabled, error, warning, and success colors only for states the theme supports.
- Support light and dark schemes for a light-colored theme as an accessibility feature unless the user or supplied direction explicitly indicates otherwise; a dark-only theme is appropriate when the direction calls for it.
- Support the operating system's color-scheme preference when the theme supports both schemes, even when it has no manual scheme switch.
- Define every semantic role in both schemes when the theme supports both.
- Do not use color as the only indicator of a link, control, error, or state.
- Verify readable foreground/background contrast for body text, controls, and focus indicators in every supported scheme.

## Examples

| Role | Use |
| --- | --- |
| canvas | page background |
| surface | cards, callouts, and raised regions |
| text | primary readable text |
| muted text | supporting metadata |
| accent | primary action and emphasis |
| border | separation without competing with content |

## Specification

Consult `hypertexting-css` to implement the resulting color system with static CSS foundations and custom properties.
