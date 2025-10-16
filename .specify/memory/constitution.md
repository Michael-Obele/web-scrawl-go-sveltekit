<!--
Sync Impact Report
==================
Version Change: INITIAL → 1.0.0
Modified Principles: N/A (initial constitution)
Added Sections:
  - Core Principles (5 principles)
  - Technology Stack Requirements
  - Development Workflow
  - Governance
Templates Status:
  ✅ plan-template.md - Verified compatible (Constitution Check section aligns)
  ✅ spec-template.md - Verified compatible (Requirements alignment maintained)
  ✅ tasks-template.md - Verified compatible (Task categorization supports principles)
Follow-up TODOs: None
-->

# Web Scraper Constitution

## Core Principles

### I. Velocity-First Development

**MUST** prioritize shipping speed over perfection while maintaining architectural integrity:

- Direct AI code generation for 80% of implementation
- Human review and fixes focused on quality and architecture alignment
- End-to-end ownership model: spec → deploy → self-verification
- Decision framework: Build vs. buy evaluated for time-to-value (favor existing libraries when quality matches)

**Rationale**: Fast iteration enables rapid user feedback and product-market fit discovery in the competitive web scraping market. AI assistance accelerates implementation without sacrificing maintainability when properly reviewed.

### II. Code Quality Standards

**MUST** maintain modular, maintainable code across the full stack:

- **Go Backend**: Concurrent patterns with goroutines; clear separation of concerns (scraping, parsing, API layers); idiomatic Go error handling
- **SvelteKit Frontend**: Reactive state management with Svelte stores; component composition following Svelte 5 runes patterns; TypeScript for type safety
- **No Over-Engineering**: Defer complexity until proven necessary; avoid premature abstractions

**Rationale**: Quality code compounds productivity over time. Modularity enables parallel development and independent testing. Simplicity reduces cognitive load and maintenance burden.

### III. Lightweight Verification (NON-NEGOTIABLE)

**MUST** validate code without blocking velocity:

- **SvelteKit**: Run `bun check` after generation to validate dependencies and TypeScript integrity
- **Go**: Generate and execute quick smoke tests (`go test ./... -v`) on core paths (e.g., scrape endpoints, parsing logic)
- **AI-Assisted Assertions**: AI validates basic output contracts (e.g., valid JSON response, expected field structure) without full test suites
- **Manual Checks Allowed**: Svelte component validation via visual inspection when automated checks would stall workflow

**Rationale**: Testing provides safety without TDD overhead. Smoke tests catch critical regressions; full suites deferred to maturity phase. Manual verification acceptable for UI where visual correctness matters more than unit coverage.

### IV. Product-Driven Development

**MUST** maintain product mindset throughout development:

- **User Value First**: Every feature justified by user outcome (e.g., "reliable markdown extraction for AI training" vs. "implement parser X")
- **Subtle Bug Awareness**: Proactively identify edge cases (e.g., JavaScript-rendered content not loading, encoding issues, timeout failures)
- **UX Excellence**: Responsive design mandatory; intuitive workflows using shadcn-svelte defaults; accessibility via ARIA patterns
- **Real-World Focus**: Target actual use cases (e.g., e-commerce scraping for tourism apps, blog content extraction for LLMs)

**Rationale**: Technical excellence without user value is waste. Product thinking ensures features solve real problems. UX quality differentiates from commodity scrapers.

### V. Ethical Scraping & Governance Alignment

**MUST** operate ethically and align AI output with stack constraints:

- **Ethical Practices**: Rate limiting and delays to avoid server overload; user-agent rotation for transparency; proxy rotation when necessary; scrape only permitted sites (user responsibility to verify)
- **Ignores robots.txt**: Intentional design for use cases where user has permission (e.g., own sites, licensed data)
- **Stack Constraints**: AI-generated code MUST use Go backend (not JavaScript), SvelteKit frontend (not React/Vue), and declared libraries (Colly, Chromedp, GoQuery, Tailwind, shadcn-svelte)
- **Constitution Compliance**: Every development phase validated against this constitution

**Rationale**: Ethical scraping protects from legal/reputational risk. Stack adherence prevents technical debt from AI hallucinations. Governance ensures consistent decision-making.

## Technology Stack Requirements

### Mandatory Technologies

**Backend (Go)**:

- **Framework**: Gin for API routing and middleware
- **Scraping**: Colly for HTML crawling, Chromedp for JavaScript rendering, GoQuery for DOM parsing
- **Concurrency**: Goroutines with channels/sync primitives for parallel scraping

**Frontend (SvelteKit + TypeScript)**:

- **Framework**: SvelteKit with Vite and Bun runtime
- **Language**: TypeScript for type safety
- **Styling**: Tailwind CSS v4 (using modern `@import` syntax, not v3 `@tailwind` directives)
- **Components**: shadcn-svelte (https://shadcn-svelte.com/) for accessible UI primitives (forms, tables, progress indicators, etc.)
- **Icons**: Lucide Svelte for consistent iconography

**Architecture**:

- **Monorepo**: Side-by-side services (separate runtimes: `bun dev` on :5173, `go run main.go` on :8080)
- **Data Flow**: SvelteKit server routes (`+page.server.js`) fetch from Go backend via HTTP; no direct client-to-Go communication
- **CORS**: Handled server-side via SvelteKit proxying
- **Storage**: In-memory for MVP (no shared database initially)

### Deployment

- **Frontend**: Vercel (SvelteKit SSR/static)
- **Backend**: Render (Go containerized service)
- **Environment**: Separate production/staging environments with distinct API endpoints

## Development Workflow

### Code Generation Flow

1. **Specification**: Create feature spec in `.specify/specs/[###-feature]/spec.md` with user stories and acceptance criteria
2. **Planning**: Run `/speckit.plan` to generate implementation plan with constitution checks
3. **AI Implementation**: Use Claude Code/Cursor for 80% code generation targeting the planned structure
4. **Validation**:
   - Run `bun check` for SvelteKit code integrity
   - Generate and run `go test ./...` for Go smoke tests
   - Manual visual checks for UI components when needed
5. **Review**: Human review for architecture alignment, edge case handling, and constitution compliance
6. **Fix & Iterate**: Address issues found in validation/review
7. **Deploy**: Push to staging → quick self-verification → production

### Quality Gates

**Pre-Merge Requirements**:

- [ ] Constitution compliance verified (stack constraints, ethical practices)
- [ ] SvelteKit: `bun check` passes without errors
- [ ] Go: Core paths have smoke tests executing successfully
- [ ] Architecture review confirms modularity and no over-engineering
- [ ] User value clearly articulated in commit/PR description

**Post-Deploy Verification**:

- [ ] Quick smoke test of deployed feature (manual or automated)
- [ ] No critical bugs in primary user flows
- [ ] Performance acceptable for target use case (e.g., scrape completes within reasonable time)

### Decision Framework: Build vs. Buy

When encountering a feature need, apply this decision tree:

1. **Existing Library Exists?** → Use it if quality matches (favor npm/Go ecosystem)
2. **Custom Required?** → Implement minimal viable version
3. **Uncertain?** → Prototype both, choose based on time-to-value and maintenance burden

## Governance

### Amendment Process

1. **Proposal**: Document proposed change with rationale and impact analysis
2. **Validation**: Verify alignment with product goals and technical constraints
3. **Version Bump**: Apply semantic versioning (MAJOR for breaking changes, MINOR for new principles, PATCH for clarifications)
4. **Propagation**: Update affected templates (plan, spec, tasks) and documentation
5. **Approval**: Human approval required before merge

### Compliance Enforcement

- **Every Feature**: Constitution Check section in `plan.md` must pass before implementation
- **AI Output**: Review all AI-generated code for stack constraint violations (no JS backend, no non-declared libraries)
- **Pull Requests**: Include constitution compliance statement in PR description
- **Retrospectives**: Quarterly review of principle effectiveness and adherence

### Complexity Justification

Any deviation from simplicity principles (e.g., introducing new framework, complex abstraction) MUST include:

- **Justification**: Why simpler approach insufficient
- **Value**: Quantified benefit (time saved, performance gain, critical capability unlocked)
- **Maintenance Plan**: How complexity will be managed long-term

**Version**: 1.0.0 | **Ratified**: 2025-10-16 | **Last Amended**: 2025-10-16
