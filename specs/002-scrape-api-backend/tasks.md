```markdown
# Tasks: 002-scrape-api-backend

**Feature**: Build /scrape API: Go backend endpoint
**Spec**: /home/node/Documents/GitHub/web-scrawl-go-sveltekit/specs/002-scrape-api-backend/spec.md
**Plan**: /home/node/Documents/GitHub/web-scrawl-go-sveltekit/specs/002-scrape-api-backend/plan.md

---

Phase 1 — Setup

- [x] T001 Initialize backend module and dependencies (no story) - `backend/go.mod`
- [x] T002 Add Gin, Colly, Chromedp, GoQuery to backend `go.mod` (no story) - `backend/go.mod`
- [x] T003 Create backend folder structure `backend/src/api` and `backend/src/models` (no story) - `backend/src/`
- [x] T004 Add environment config loader for scraper settings (no story) - `backend/src/config/config.go`

Phase 2 — Foundational (blocking prerequisites)

- [x] T005 Create ScrapeResult model and Link type [US1] - `backend/src/models/scraper.go`
- [x] T006 Implement basic HTTP server bootstrap using Gin (listens :8080) (no story) - `backend/main.go`
- [x] T007 Implement configuration values with defaults: SCRAPER_DELAY_S, SCRAPER_USER_AGENTS, CHROMEDP_TIMEOUT_S, SCRAPER_TIMEOUT_S (no story) - `backend/src/config/config.go`
- [x] T008 Add basic logging and error types for scraper responses (no story) - `backend/src/api/errors.go`

Phase 3 — User Story 1 (P1) — Run a single-site scrape

- [x] T009 [US1] Implement core scraping service skeleton that accepts URL+depth and returns ScrapeResult (Colly+GoQuery skeleton; Chromedp NOP hook) - `backend/src/services/scraper_service.go`
- [x] T010 [US1] Implement endpoint handler GET /scrape?url={url}&depth={n} and wire to scraper service - `backend/src/api/scrape_handler.go`
- [x] T011 [US1] Normalize and validate URL input, return 400 for invalid URLs - `backend/src/api/scrape_handler.go`
- [x] T012 [US1] Implement Colly-based crawl for depth=1 that collects page body and outbound links and populates ScrapeResult fields (title, links) - `backend/src/services/scraper_service.go`
- [x] T013 [US1] Implement Markdown conversion from extracted HTML/main content and populate `markdown` in ScrapeResult - `backend/src/services/scraper_service.go`
- [x] T014 [US1] Add fetchedAt timestamp generation and include in ScrapeResult - `backend/src/services/scraper_service.go`
- [x] T015 [US1] Add unit/integration smoke test: call GET /scrape?url=https://example.com&depth=1 and assert valid JSON and required fields - `backend/tests/smoke_test.go`

Phase 4 — User Story 2 (P2) — UI trigger & results

- [ ] T016 [US2] Add SvelteKit route `frontend/src/routes/scrape/+page.server.ts` that proxies POST/GET to backend `/scrape` on server-side - `frontend/src/routes/scrape/+page.server.ts`
- [ ] T017 [US2] Add minimal UI page `frontend/src/routes/scrape/+page.svelte` with shadcn-svelte form for URL input and a submit button - `frontend/src/routes/scrape/+page.svelte`
- [ ] T018 [US2] Add progress indicator component and wiring to show in-progress state during scrape - `frontend/src/lib/components/Progress.svelte`
- [ ] T019 [US2] Render ScrapeResult: markdown area and links table (use shadcn-svelte table) - `frontend/src/routes/scrape/+page.svelte`
- [ ] T020 [US2] Ensure TypeScript types for ScrapeResult in `frontend/src/lib/types/scraper.ts` (matching backend contract) - `frontend/src/lib/types/scraper.ts`

Phase 5 — User Story 3 (P3) — Ethical scraping controls

- [ ] T021 [US3] Implement per-host delay enforcement in scraper service using configuration value (SCRAPER_DELAY_S) - `backend/src/services/scraper_service.go`
- [ ] T022 [US3] Implement user-agent rotation using configured list SCRAPER_USER_AGENTS - `backend/src/services/scraper_service.go`
- [ ] T023 [US3] Add respect-for-robots.txt check with configurable override (env var SCRAPER_IGNORE_ROBOTS=false by default) - `backend/src/services/scraper_service.go`
- [ ] T024 [US3] Add warning flags to ScrapeResult when enforcement (robots blocked, rendering fallback) occurs - `backend/src/models/scraper.go`

Phase 6 — Polish & Cross-cutting concerns

- [ ] T025 Add README / quickstart usage to `specs/002-scrape-api-backend/quickstart.md` (update with observed commands) - `specs/002-scrape-api-backend/quickstart.md`
- [ ] T026 Add logging, graceful shutdown, and context cancellation to server and scraper (no story) - `backend/main.go`, `backend/src/services/scraper_service.go`
- [ ] T027 Add documentation comment blocks for public functions and types in backend (no story) - `backend/src/**.go`
- [ ] T028 Ensure frontend `bun check` passes and run Svelte MCP autofixer if issues (no story) - `frontend/` (CI step)

Dependencies (story completion order)

1. Setup (T001..T004)
2. Foundational (T005..T008)
3. US1 (T009..T015)
4. US2 (T016..T020) — depends on US1 for backend
5. US3 (T021..T024) — can be implemented in parallel after Foundational tasks
6. Polish (T025..T028)

Parallel execution opportunities

- T005 (model) and T006 (server bootstrap) can be implemented in parallel (different files) — mark as [P]
- T021..T023 (ethical controls) can be implemented in parallel with US1 implementation after the core service skeleton exists — mark as [P]

Implementation strategy

- MVP scope: Only User Story 1 (Run a single-site scrape) implemented and smoke-tested. That is the smallest independently testable slice that delivers core value.
- Deliverables per sprint:
  - Sprint 1: Setup + Foundational + US1 core (T001..T015) — smoke test green
  - Sprint 2: UI integration (T016..T020)
  - Sprint 3: Ethical controls + polish (T021..T028)

Task counts & summary

- Total tasks: 28
- Tasks per story:
  - US1: 7 (T009..T015)
  - US2: 5 (T016..T020)
  - US3: 4 (T021..T024)
  - Setup/Foundational/Polish: 12 (T001..T008, T025..T028)

All tasks follow the checklist format with IDs and file paths. Tasks marked [P] are parallelizable. The independent test for US1 is the smoke test T015 against `https://example.com`.
```
