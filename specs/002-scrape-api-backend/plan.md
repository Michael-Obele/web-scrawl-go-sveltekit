# Implementation Plan: [FEATURE]

**Branch**: `[###-feature-name]` | **Date**: [DATE] | **Spec**: [link]
**Input**: Feature specification from `/specs/[###-feature-name]/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/commands/plan.md` for the execution workflow.

## Summary

Implement a small Go backend endpoint GET /scrape that accepts a URL and depth parameter and returns a ScrapeResult JSON (title, markdown, links, warnings, fetchedAt). The backend will use Colly for crawling and link discovery, Chromedp for rendering JS-heavy pages, and GoQuery for extraction/parsing. The SvelteKit frontend will proxy requests through a server-side handler (+page.server.ts) and provide a minimal UI (shadcn-svelte form, progress indicator, results area). Ethical defaults (per-host delay, user-agent rotation, respect for robots.txt) and sensible timeouts are included. A smoke test against https://example.com will validate the endpoint returns valid ScrapeResult JSON.

## Technical Context

**Language/Version**: Go 1.21+ (backend); SvelteKit (Svelte 5) + TypeScript (frontend)
**Primary Dependencies**: Gin (HTTP server), Colly (crawler), Chromedp (headless rendering), GoQuery (parsing), shadcn-svelte (frontend UI), Tailwind CSS v4
**Storage**: N/A for MVP — in-memory response only; no persistent storage
**Testing**: go test for backend smoke/integration tests; Bun check + TypeScript validation and visual/manual verification for SvelteKit frontend
**Target Platform**: Linux server environment (local dev and CI runners), browser for frontend
**Project Type**: Web application (monorepo with `backend/` and `frontend/` directories)
**Performance Goals**: User-facing scrapes for small/stable pages complete within 10s (95th percentile) for smoke-test targets
**Constraints**: Respect robots.txt by default (configurable), per-host delay default 2s, Chromedp render timeout default 10s, overall request timeout 30s
**Scale/Scope**: Low-volume manual/QA scraping for preview and testing; NOT intended for large-scale crawling in this phase

## Constitution Check

GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.

Using the repository constitution (`.specify/memory/constitution.md`):

- Stack constraint: MUST use Go backend and SvelteKit frontend — OK (spec requests Go + SvelteKit)
- Declared libraries: Colly, Chromedp, GoQuery — OK (explicitly requested in spec)
- Ethical scraping: Rate limiting/delays and user-agent rotation required — OK (FR-004)
- Testing gates: go test smoke test and bun check for frontend — will be added to tasks

No constitution violations detected; proceed with Phase 0 research.

## Project Structure

### Documentation (this feature)

```
specs/[###-feature]/
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command)
├── data-model.md        # Phase 1 output (/speckit.plan command)
├── quickstart.md        # Phase 1 output (/speckit.plan command)
├── contracts/           # Phase 1 output (/speckit.plan command)
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)
```

### Source Code (repository root)

<!--
  ACTION REQUIRED: Replace the placeholder tree below with the concrete layout
  for this feature. Delete unused options and expand the chosen structure with
  real paths (e.g., apps/admin, packages/something). The delivered plan must
  not include Option labels.
-->

```
# [REMOVE IF UNUSED] Option 1: Single project (DEFAULT)
src/
├── models/
├── services/
├── cli/
└── lib/

tests/
├── contract/
├── integration/
└── unit/

# [REMOVE IF UNUSED] Option 2: Web application (when "frontend" + "backend" detected)
backend/
├── src/
│   ├── models/
│   ├── services/
│   └── api/
└── tests/

frontend/
├── src/
│   ├── components/
│   ├── pages/
│   └── services/
└── tests/

# [REMOVE IF UNUSED] Option 3: Mobile + API (when "iOS/Android" detected)
api/
└── [same as backend above]

ios/ or android/
└── [platform-specific structure: feature modules, UI flows, platform tests]
```

**Structure Decision**: Use the existing monorepo layout detected in repository root: `backend/` for Go service and `frontend/` for SvelteKit UI. This matches constitution and repo conventions.

## Complexity Tracking

_Fill ONLY if Constitution Check has violations that must be justified_

| Violation                  | Why Needed         | Simpler Alternative Rejected Because |
| -------------------------- | ------------------ | ------------------------------------ |
| [e.g., 4th project]        | [current need]     | [why 3 projects insufficient]        |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient]  |
