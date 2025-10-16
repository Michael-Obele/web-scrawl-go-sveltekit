---
description: "Task list for monorepo initialization feature"
---

# Tasks: Monorepo Initialization

**Input**: Design documents from `/specs/001-monorepo-init/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: Tests are OPTIONAL - only included if explicitly requested. This feature uses lightweight verification per constitutional principle III (bun check + go test smoke tests).

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- **Web app**: `backend/` and `frontend/` at repository root
- Paths shown below follow the monorepo structure from plan.md

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and prerequisite checks

- [X] T001 Verify Bun runtime installed (version 1.0+) via `bun --version`
- [X] T002 Verify Go installed (version 1.21+) via `go version`
- [X] T003 Verify ports 5173 and 8080 are available (check with `lsof` or `netstat`)
- [X] T004 Create root-level README.md with project overview and setup instructions

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T005 Create `backend/` directory at repository root
- [X] T006 Create `frontend/` directory at repository root
- [X] T007 Create `.gitignore` with entries for `node_modules/`, `bun.lockb`, `go.sum`, `.DS_Store`, `dist/`, `.svelte-kit/`

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Project Structure Creation (Priority: P1) 🎯 MVP

**Goal**: Establish monorepo structure with frontend (SvelteKit + Bun + Tailwind + shadcn-svelte) and backend (Go + Gin) directories, complete with all dependencies and configuration files.

**Independent Test**: Verify `frontend/` and `backend/` directories exist with complete configuration files. Check `frontend/package.json` includes all required dependencies (SvelteKit, Tailwind, Lucide Svelte, shadcn-svelte). Check `backend/go.mod` includes Gin, Colly, Chromedp, GoQuery. Run `bun check` in frontend (should pass) and attempt `go build` in backend (should compile).

#### Frontend Setup (T008-T020)

- [X] T008 [P] [US1] Run `npx sv create frontend` selecting: SvelteKit demo app, TypeScript syntax, ESLint, Prettier, and use Bun as package manager
- [X] T009 [US1] Run `bun install` in frontend/ directory to install SvelteKit dependencies (if not done during project creation)
- [X] T010 [US1] Verify frontend/package.json created with SvelteKit, Vite, and TypeScript dependencies and `"type": "module"`
- [X] T011 [P] [US1] Install Tailwind CSS v4: Run `cd frontend && bun add -D tailwindcss postcss autoprefixer` (Tailwind v4 uses new import syntax)
- [X] T012 [US1] Verify Tailwind files created: frontend/tailwind.config.js, frontend/postcss.config.js with v4 configuration
- [X] T013 [US1] Confirm Tailwind content paths in frontend/tailwind.config.js scan `./src/**/*.{html,js,svelte,ts}`
- [X] T014 [P] [US1] Verify frontend/src/app.css contains Tailwind v4 imports: `@import 'tailwindcss';` (v4 uses @import, not @tailwind directives)
- [X] T015 [US1] Verify frontend/src/routes/+layout.svelte imports '../app.css' and includes `<slot />`
- [X] T016 [P] [US1] Install Lucide Svelte icons: `bun add lucide-svelte` in frontend/
- [X] T017 [US1] Initialize shadcn-svelte: Run `bunx shadcn-svelte@latest init` in frontend/, selecting Slate base color, confirming src/app.css path
- [X] T018 [US1] Install test Button component: Run `bunx shadcn-svelte@latest add button` in frontend/ to verify shadcn-svelte setup
- [X] T019 [US1] Update frontend/package.json scripts to include: `"dev": "vite dev --port 5173"`, `"build": "vite build"`, `"check": "svelte-kit sync && svelte-check --tsconfig ./tsconfig.json"`
- [X] T020 [US1] Run `bun check` in frontend/ to validate TypeScript configuration (should pass with no errors)
- [ ] T009 [US1] Run `bun install` in frontend/ directory to install SvelteKit dependencies (if not done during project creation)
- [ ] T010 [US1] Verify frontend/package.json created with SvelteKit, Vite, and TypeScript dependencies and `"type": "module"`
- [X] T011 [P] [US1] Install Tailwind CSS v4: Run `cd frontend && bun add -D tailwindcss postcss autoprefixer` (Tailwind v4 uses new import syntax)
- [X] T012 [US1] Verify Tailwind files created: frontend/tailwind.config.js, frontend/postcss.config.js with v4 configuration
- [X] T013 [US1] Confirm Tailwind content paths in frontend/tailwind.config.js scan `./src/**/*.{html,js,svelte,ts}`
- [X] T014 [P] [US1] Verify frontend/src/app.css contains Tailwind v4 imports: `@import 'tailwindcss';` (v4 uses @import, not @tailwind directives)
- [ ] T015 [US1] Verify frontend/src/routes/+layout.svelte imports '../app.css' and includes `<slot />`
- [ ] T016 [P] [US1] Install Lucide Svelte icons: `bun add lucide-svelte` in frontend/
- [ ] T017 [US1] Initialize shadcn-svelte: Run `bunx shadcn-svelte@latest init` in frontend/, selecting Slate base color, confirming src/app.css path
- [ ] T018 [US1] Install test Button component: Run `bunx shadcn-svelte@latest add button` in frontend/ to verify shadcn-svelte setup
- [ ] T019 [US1] Update frontend/package.json scripts to include: `"dev": "vite dev --port 5173"`, `"build": "vite build"`, `"check": "svelte-kit sync && svelte-check --tsconfig ./tsconfig.json"`
- [ ] T020 [US1] Run `bun check` in frontend/ to validate TypeScript configuration (should pass with no errors)

### Backend Setup (Go + Gin)

- [X] T021 [P] [US1] Create backend/ subdirectories: `mkdir -p backend/{handlers,middleware,tests}` from repo root
- [X] T022 [US1] Initialize Go module: `go mod init github.com/yourusername/web-scraper-backend` in backend/ (replace yourusername with actual GitHub username)
- [X] T023 [P] [US1] Install Gin framework: `go get -u github.com/gin-gonic/gin` in backend/
- [X] T024 [P] [US1] Install CORS middleware: `go get -u github.com/gin-contrib/cors` in backend/
- [X] T025 [P] [US1] Install Colly v2 scraping library: `go get -u github.com/gocolly/colly/v2` in backend/
- [X] T026 [P] [US1] Install Chromedp for browser automation: `go get -u github.com/chromedp/chromedp` in backend/
- [X] T027 [P] [US1] Install GoQuery for HTML parsing: `go get -u github.com/PuerkitoBio/goquery` in backend/
- [X] T028 [US1] Verify backend/go.mod includes all dependencies: Gin, CORS, Colly v2, Chromedp, GoQuery with proper versioning
- [X] T029 [US1] Run `go mod tidy` in backend/ to clean up dependencies and verify backend/go.sum generated with checksums

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently. Both directories exist with all dependencies installed.

---

## Phase 4: User Story 2 - Development Environment Setup (Priority: P2)

**Goal**: Enable developers to start frontend (SvelteKit on :5173) and backend (Gin on :8080) services independently with hot-reloading and proper logging.

**Independent Test**: Run `bun dev` from frontend/ directory - verify dev server starts on port 5173 with HMR enabled. Run `go run main.go` from backend/ directory - verify Gin server starts on :8080 with startup logs. Confirm both services can run simultaneously without port conflicts. Make a change to a frontend file and verify browser auto-reloads.

### Backend Server Implementation

- [X] T030 [P] [US2] Create backend/main.go with package main, import statements for Gin and CORS
- [X] T031 [US2] Implement main() function in backend/main.go: initialize Gin router with `gin.Default()`
- [X] T032 [US2] Add CORS middleware config in backend/main.go: allow origin `http://localhost:5173`, methods GET/POST/PUT/DELETE, headers Content-Type/Authorization
- [X] T033 [US2] Apply CORS middleware to router using `router.Use(cors.New(config))` in backend/main.go
- [X] T034 [US2] Add startup logging in backend/main.go: `log.Println("Starting server on :8080")`
- [X] T035 [US2] Add server start command in backend/main.go: `router.Run(":8080")` with error handling
- [X] T036 [US2] Test backend startup: run `go run main.go` from backend/ and verify server starts without errors

### Frontend Development Server

- [X] T037 [P] [US2] Verify frontend/vite.config.ts exists and configures SvelteKit plugin
- [X] T038 [US2] Test frontend startup: run `bun dev` from frontend/ and verify dev server starts on port 5173
- [X] T039 [US2] Verify hot module replacement: make a change to frontend/src/routes/+page.svelte and confirm browser auto-reloads
- [X] T040 [US2] Update frontend/src/routes/+page.svelte with minimal content: heading "Web Scraper", paragraph with project description

### Concurrent Execution Validation

- [X] T041 [US2] Test concurrent execution: start backend with `go run main.go` in one terminal, start frontend with `bun dev` in another terminal
- [X] T042 [US2] Verify both services accessible: `http://localhost:5173` (frontend) and `http://localhost:8080` (backend logs confirm running)
- [X] T043 [US2] Document run commands in README.md: separate sections for starting frontend and backend

**Checkpoint**: Both services now start successfully and can run concurrently. Development workflow enabled.

---

## Phase 5: User Story 3 - Health Check Endpoint & Proxy (Priority: P3)

**Goal**: Implement `/health` endpoint in Go backend, create SvelteKit server route proxy, and build UI page to display health status - validating complete frontend-to-backend communication pattern.

**Independent Test**: Start both services. Access `http://localhost:8080/health` directly - verify returns JSON `{status: "healthy", service: "web-scraper-backend", timestamp: <number>}`. Access `http://localhost:5173/health` through frontend - verify page displays health status from backend without CORS errors. Stop backend, refresh frontend health page - verify error handling shows appropriate message.

### Backend Health Endpoint

- [X] T044 [P] [US3] Create backend/handlers/health.go with package handlers
- [X] T045 [US3] Implement HealthCheck handler in backend/handlers/health.go: return JSON with status="healthy", service="web-scraper-backend", timestamp=Unix time
- [X] T046 [US3] Register health route in backend/main.go: `router.GET("/health", handlers.HealthCheck)` before server start
- [X] T047 [US3] Test health endpoint directly: run `curl http://localhost:8080/health` and verify JSON response with correct fields

### Backend Health Test (Smoke Test)

- [X] T048 [P] [US3] Create backend/tests/health_test.go with package tests
- [X] T049 [US3] Implement TestHealthCheck in backend/tests/health_test.go: use httptest to call /health handler, assert status 200, assert response body contains status="healthy"
- [X] T050 [US3] Run backend tests: `go test ./... -v` from backend/ directory and verify TestHealthCheck passes

### Frontend Health Proxy & UI

- [X] T051 [P] [US3] Create frontend/src/routes/health/ directory
- [X] T052 [US3] Create frontend/src/routes/health/+page.server.js with load function that fetches from `http://localhost:8080/health`
- [X] T053 [US3] Add error handling in +page.server.js load function: catch fetch failures, return `{health: null, error: "Backend service is not available"}`
- [X] T054 [P] [US3] Create frontend/src/routes/health/+page.svelte with script importing Check and X icons from lucide-svelte
- [X] T055 [US3] Add props destructuring in +page.svelte: `let { data } = $props();`
- [X] T056 [US3] Implement success UI in +page.svelte: green alert with Check icon showing backend health data (status, service, timestamp)
- [X] T057 [US3] Implement error UI in +page.svelte: red alert with X icon showing error message when backend unavailable
- [X] T058 [US3] Update frontend/src/routes/+page.svelte to import Button from '$lib/components/ui/button' and add link to /health route

### End-to-End Integration Testing

- [X] T059 [US3] Test health proxy with backend running: access `http://localhost:5173/health` and verify green success message displays
- [X] T060 [US3] Test error handling: stop backend, refresh frontend /health page, verify red error message displays
- [X] T061 [US3] Verify no CORS errors: check browser console while accessing /health route - should show no CORS-related errors
- [X] T062 [US3] Test timestamp display: verify health page shows formatted timestamp (convert Unix epoch to readable date)

**Checkpoint**: Complete frontend-to-backend communication validated. Health check proves architectural pattern for all future API calls.

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Documentation, final validation, and project cleanup

- [X] T063 [P] Create backend/README.md with: Go version requirement, dependency installation, run commands, test commands
- [X] T064 [P] Create frontend/README.md with: Bun version requirement, dependency installation, run commands, available scripts
- [X] T065 Update root README.md with: project overview, monorepo structure explanation, quick start guide referencing specs/001-monorepo-init/quickstart.md
- [X] T066 Run final validation: `bun check` in frontend/ (TypeScript validation)
- [X] T067 Run final validation: `go test ./... -v` in backend/ (smoke tests)
- [X] T068 Verify .gitignore excludes: node_modules/, bun.lockb, .svelte-kit/, dist/, go.sum (check with `git status`)
- [X] T069 Create backend/.env.example with PORT=8080 and ALLOWED_ORIGINS=http://localhost:5173
- [X] T070 Create frontend/.env.example with PUBLIC_BACKEND_URL=http://localhost:8080
- [ ] T071 Test complete workflow from scratch: follow quickstart.md steps and verify all commands execute successfully
- [ ] T072 Commit all changes: `git add . && git commit -m "feat: initialize monorepo with SvelteKit and Go"`

---

## Dependencies & Execution Order

### User Story Completion Order

1. **User Story 1 (P1)**: MUST complete first - establishes directory structure and dependencies
2. **User Story 2 (P2)**: Can start after US1 complete - enables development workflow
3. **User Story 3 (P3)**: Can start after US2 complete - validates integration pattern

**Dependency Graph**:

```
Setup (T001-T004)
  ↓
Foundational (T005-T007)
  ↓
US1: Structure (T008-T029) ← Foundation for everything
  ↓
US2: Dev Environment (T030-T043) ← Needs US1 files
  ↓
US3: Health & Proxy (T044-T062) ← Needs US2 servers running
  ↓
Polish (T063-T072)
```

### Parallel Execution Opportunities

**Within User Story 1**:

- Frontend tasks (T008-T020) and Backend tasks (T021-T029) can run in parallel
- Dependency installation (T011, T016, T023-T027) can run in parallel after project init

**Within User Story 2**:

- Backend implementation (T030-T036) and Frontend updates (T037-T040) can run in parallel

**Within User Story 3**:

- Backend handler (T044-T047) and Backend test (T048-T050) can run in parallel
- Frontend proxy (T051-T053) and Frontend UI (T054-T058) can run in parallel after initial file creation

**In Polish Phase**:

- Documentation tasks (T063-T065) can run in parallel
- Validation tasks (T066-T067) can run in parallel

---

## Implementation Strategy

### MVP Definition (Minimum Viable Product)

**Suggested MVP Scope**: User Story 1 ONLY

- Delivers: Complete monorepo structure with all dependencies installed
- Value: Enables any developer to start building features immediately
- Independent Test: Run `bun check` and attempt `go build` - both should succeed
- Time Estimate: 2-3 hours with AI assistance (80% code generation)

### Incremental Delivery Plan

1. **Sprint 1 (MVP)**: User Story 1 - Project Structure

   - Deliverable: Both directories with working dependency management
   - Validation: `bun check` passes, `go build` compiles

2. **Sprint 2**: User Story 2 - Development Environment

   - Deliverable: Both services start and run concurrently
   - Validation: Access both ports, confirm HMR works

3. **Sprint 3**: User Story 3 - Health Check Integration

   - Deliverable: Complete frontend-to-backend communication
   - Validation: Health page displays backend status, no CORS errors

4. **Sprint 4**: Polish & Documentation
   - Deliverable: Production-ready monorepo with complete docs
   - Validation: New developer can follow quickstart.md successfully

---

## Task Summary

**Total Tasks**: 72

- **Setup Phase**: 4 tasks (T001-T004)
- **Foundational Phase**: 3 tasks (T005-T007)
- **User Story 1 (P1)**: 22 tasks (T008-T029)
  - Frontend: 13 tasks
  - Backend: 9 tasks
- **User Story 2 (P2)**: 14 tasks (T030-T043)
  - Backend: 7 tasks
  - Frontend: 4 tasks
  - Integration: 3 tasks
- **User Story 3 (P3)**: 19 tasks (T044-T062)
  - Backend: 7 tasks
  - Frontend: 8 tasks
  - Integration: 4 tasks
- **Polish Phase**: 10 tasks (T063-T072)

**Parallel Opportunities Identified**: 21 tasks marked with [P] can run in parallel

**Independent Test Criteria**:

- ✅ US1: Directories exist with complete config files, `bun check` passes, `go build` compiles
- ✅ US2: Both services start on correct ports, HMR works, concurrent execution successful
- ✅ US3: Health endpoint returns JSON, frontend displays status, no CORS errors, error handling works

**Suggested MVP**: User Story 1 (22 tasks) - Establishes foundation for all future development

---

## Format Validation

✅ **All tasks follow required checklist format**:

- Checkbox: `- [ ]` prefix on every task
- Task ID: Sequential T001-T072
- [P] marker: 21 tasks identified as parallelizable
- [Story] label: All user story tasks labeled [US1], [US2], or [US3]
- File paths: Included in all implementation tasks

✅ **Organization**: Tasks grouped by user story for independent implementation

✅ **Dependencies**: Clear completion order documented with dependency graph

✅ **Tests**: Smoke tests included per constitutional lightweight verification principle

---

**Tasks Status**: ✅ Ready for implementation. Execute sequentially by phase, leverage parallel opportunities within each user story.
