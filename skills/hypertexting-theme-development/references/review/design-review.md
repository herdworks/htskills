# Design review

## Overview

Use a working preview to show how the completed or in-progress theme extends
the supplied direction across its required surfaces. Review the theme, not
unfinished production content.

## Instructions

- Render the page families in the theme README with representative existing content or fixture data before the review.
- Walk through the homepage, every distinct page presentation, responsive navigation, and the component states that materially affect the user experience.
- When the theme supports both schemes, show the user a representative sampling in light and dark mode, including navigation and at least one interactive component state; do not require a separate walkthrough of every surface in both schemes.
- Relate each surface to the documented information architecture, visual systems, and customization contract.
- Call out consequential extensions to the supplied direction as decisions, explain their purpose, and invite feedback on them.
- Ask for feedback on the work shown; do not ask the user to design missing surfaces from scratch.
- Preserve and classify any new user-provided design assets with `hypertexting-assets` before using them in the next implementation pass.
- Record accepted review feedback in the theme README before making the next implementation pass.

### Preview on a local network

To review from another device on the same local network, start the preview
server from the website with:

```sh
hyperctl website server --config <site.yaml> --bind 0.0.0.0
```

- Supply `--theme <theme.json>` when the website configuration does not already select the theme under review.
- Tell the user the server's local-network URL and port after it starts.
- Stop the preview server when the review is complete.

## Specification

Use `hyperctl website build` and `hyperctl theme render` to verify required
surfaces before the review. Use `hyperctl website server` only for local review;
it is not a deployment command.
