# Information architecture

## Overview

Information architecture decides what the theme must present, not the website's
real content. Establish it before defining page presentations, components, or navigation
styling. The result is a page-and-navigation matrix that gives the theme a
finite completeness checklist.

## Instructions

- Inspect supplied mockups, website configuration, existing pages, and stated goals for evidence of sections, page families, navigation, and primary actions.
- Treat a repeated visual treatment or repeated content relationship as evidence for a page family or shared region.
- Distinguish a section landing page, a browsable index, and a reusable detail page; one section may need one, two, or all three.
- Put destinations visitors need to move between frequently in primary navigation.
- Put lower-frequency utility, legal, privacy, and secondary destinations in the footer.
- Make one action visually primary when the website has a conversion, subscription, contact, download, or other next step.

### Ask targeted requirements questions

- Ask or revisit which visitor outcome should be primary only when the supplied direction or feedback does not make an action clear.
- Ask or revisit which section or content family must exist only when the supplied context or feedback cannot distinguish it from the baseline architecture.
- Do not ask the user to enumerate page presentations; derive them from the information architecture.
- Revisit these questions as needed when feedback changes the information architecture; otherwise retain the documented decisions.
- Otherwise recommend the baseline and extend the supplied direction conservatively.

### Start from the baseline

- Recommend this baseline only when the supplied context does not establish a more specific architecture: home, primary-section landing, browse/index, reusable detail, standard page, and status/utility surfaces.
- Start primary navigation with Home, up to three audience-relevant primary sections, an About or trust destination, and one clear action.
- Start footer navigation with the needed primary destinations plus legal, privacy, and utility links.
- Add search only when visitors need to find a meaningful body of content.
- Add tags, feeds, archives, and extra detail types only when the website's content model or user context calls for them.

### Produce the page-and-navigation matrix

- Give every section a purpose, audience, primary navigation status, and footer status.
- Name each required page family and its relationship to a section: landing, index, detail, standard, or status/utility.
- Identify the shared header, primary navigation, footer, and any shared secondary navigation.
- Treat this matrix as the theme's coverage contract; do not add real page copy or records to satisfy it.

### Refine from feedback

Ask these questions only when user feedback makes the answer material:

- Does the feedback introduce a new audience, primary task, section, or content family the theme must support?
- Should a requested destination become a primary-navigation item, a footer or secondary item, or remain contextual?
- Does the feedback change the information architecture, or only the presentation of an existing surface?

## Examples

### Personal blog

A personal blog is an editorial website. Its primary repeated relationship is a
reader browsing a journal and opening individual entries, so the journal index
and entry detail are required theme surfaces. Search, tags, and archives become
useful only when the expected body of writing makes them useful to readers.

| Section or surface | Purpose | Page family | Navigation |
| --- | --- | --- | --- |
| Home | Orient readers and feature current or representative writing | home | primary |
| Journal | Browse all writing | index | primary |
| Journal entry | Read one post with metadata and related writing | detail | linked from index |
| About | Explain the author and the publication | standard | primary or footer |
| Tags or archive | Browse a large body of writing by topic or time | index | secondary or footer when needed |
| Subscribe or contact | Give readers a clear next step | standard or action | primary action or footer |
| Privacy and not found | State policy and recover from invalid URLs | standard and status | footer and none |

### Professional services company

A professional services company needs to help a prospective client understand
its offer, establish trust, and make contact. Service landing pages and work or
case-study details are required only when the company presents more than one
service or has evidence to show; a journal is optional rather than assumed.

| Section or surface | Purpose | Page family | Navigation |
| --- | --- | --- | --- |
| Home | State the company proposition, evidence, and primary contact action | home | primary |
| Services | Introduce the service offering | landing | primary |
| Service | Explain one distinct service when offerings need separate depth | detail or standard | linked from services |
| Work or case studies | Demonstrate outcomes and relevant experience | index | primary when evidence exists |
| Case study | Explain one engagement, outcome, and supporting evidence | detail | linked from work |
| About | Establish the company, team, approach, and trust signals | standard | primary or footer |
| Contact | Give prospective clients a clear next step | standard or action | primary action |
| Privacy and not found | State policy and recover from invalid URLs | standard and status | footer and none |

## Specification

Use [HyperTexting data](../../../hypertexting-data/SKILL.md) for the existing website and page data contracts. This reference decides how those contracts should be presented by a theme.
