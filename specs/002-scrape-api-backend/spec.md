```markdown
# Feature Specification: Build /scrape API: Go backend endpoint

**Feature Branch**: `002-scrape-api-backend`
**Created**: 2025-10-16
**Status**: Draft
**Input**: User description: "Build /scrape API: Go backend endpoint (GET /scrape?url=...&depth=1) using Colly for crawl, Chromedp for JS render, GoQuery for extraction—output ScrapeResult JSON (markdown, title, links). Add ethical delays/user-agent rotation. Proxy via SvelteKit +page.server.ts to a basic +page.svelte UI (shadcn form for URL input, progress indicator, results table/pre). Include quick smoke test: Assert valid JSON on example.com."

## User Scenarios & Testing _(mandatory)_

### User Story 1 - Run a single-site scrape (Priority: P1)

A developer or operator wants to fetch and extract the visible content of a single URL (including JS-rendered content) and receive a structured JSON response suitable for downstream processing or display.

**Why this priority**: This is the core feature: returning a reliable, structured ScrapeResult for a single URL enables all higher-level uses (preview, indexing, transformation).

**Independent Test**: Call the backend endpoint GET /scrape?url=https://example.com&depth=1 and verify the response is valid ScrapeResult JSON containing a title, markdown content, and an array of discovered links.

**Acceptance Scenarios**:

1. **Given** a reachable HTTP(S) URL, **When** the client requests GET /scrape?url={url}&depth=1, **Then** the service returns 200 with a JSON body matching the ScrapeResult schema (title, markdown, links).
2. **Given** a URL that requires JS rendering to reveal content, **When** requested, **Then** the service uses headless rendering (Chromedp) to include rendered DOM text in the markdown output.

---

### User Story 2 - UI trigger & results (Priority: P2)

An end-user (product evaluator) uses the SvelteKit frontend to submit a URL and watch progress; when complete, they can view and copy the scraped markdown and see links in a table.

**Why this priority**: Provides quick verification and demo capability for the backend API; useful for manual QA and stakeholders.

**Independent Test**: Open the frontend page, submit https://example.com, observe progress indicator, and verify result table and markdown output render correctly.

**Acceptance Scenarios**:

1. **Given** the frontend page, **When** the user submits a valid URL, **Then** the UI shows an in-progress state and then displays the ScrapeResult JSON rendered as markdown and a table of links when complete.

---

### User Story 3 - Ethical scraping controls (Priority: P3)

An operator wants the scraper to behave politely: use delays between requests, rotate user-agents, and respect robots.txt where applicable.

**Why this priority**: Prevents misuse or accidental overload of target sites and aligns with good-citizen scraping practices.

**Independent Test**: Run the scraper against a test site and assert that the request headers vary (user-agent rotation) and that per-host crawl delays are applied (observed timestamps between requests).

**Acceptance Scenarios**:

1. **Given** a scrape job that touches multiple pages on the same host, **When** the crawler is running, **Then** the crawler inserts a configurable delay between requests to the same host.
2. **Given** a scrape job, **When** new requests are made, **Then** the outgoing requests cycle through a configurable list of user-agent strings.

---

### Edge Cases

- If the target URL is unreachable or returns a non-2xx status, return a 4xx/5xx response with a clear error message and no ScrapeResult body.
- If the page contains very large resources (huge DOM or infinite scroll), apply a time and depth limit and return a partial result plus a warning in the response.
- If robots.txt disallows crawling, respect it and return 403 with a reason unless an operator override is explicitly set (assumption: default is to respect robots.txt).
- If Chromedp fails for a page (timeouts, missing headless), fall back to Colly's static fetch and include a warning flag in the response.

## Requirements _(mandatory)_

### Functional Requirements

- **FR-001**: Backend MUST expose GET /scrape?url={url}&depth={n} that accepts a URL and optional depth (default 1).
- **FR-002**: Backend MUST perform crawling using Colly for link discovery up to the requested depth and use Chromedp for pages that require JS rendering.
- **FR-003**: Backend MUST extract the page title, main readable content converted to Markdown, and outbound links (absolute URLs) and return them as a ScrapeResult JSON object.
- **FR-004**: Backend MUST include ethical controls: per-host delay (configurable), user-agent rotation (configurable list), and respect for robots.txt by default.
- **FR-005**: Backend MUST provide clear error responses when the target is unreachable, disallowed, or when scraping fails.
- **FR-006**: Frontend (SvelteKit) MUST proxy requests to the backend via a server-side handler (+page.server.ts) to avoid exposing backend endpoints directly to client origin policies.
- **FR-007**: Frontend MUST provide a minimal UI with a URL input form (shadcn-svelte), a progress indicator, and a results area showing markdown and links in a table.
- **FR-008**: Include a quick smoke test script (or test) that calls GET /scrape?url=https://example.com&depth=1 and asserts the response is valid JSON matching the ScrapeResult schema.

_FR-009 (non-functional)_: The service should have sensible defaults: depth=1, per-host delay=2s, user-agent rotation list of at least 3 common browser strings, Chromedp timeout=10s, overall request timeout=30s.

### Key Entities _(include if feature involves data)_

- **ScrapeResult**: Represents output from a scrape job
  - title: string
  - markdown: string (main page content converted to Markdown)
  - links: array of { href: string, text?: string }
  - warnings?: array of string (optional)
  - fetchedAt: ISO-8601 timestamp

## Success Criteria _(mandatory)_

### Measurable Outcomes

- **SC-001**: 95% of successful scrapes for stable, small pages (e.g., example.com) return a ScrapeResult within 10 seconds.
- **SC-002**: The smoke test for https://example.com passes: the endpoint returns valid JSON containing a non-empty title and markdown and at least one link.
- **SC-003**: Frontend UI shows progress and displays results correctly for at least the primary flow in manual testing by a QA user.
- **SC-004**: The scraper respects robots.txt for at least 90% of tested domains by default (where robots.txt is present and parseable).

## Assumptions

- The repository will host the Go backend under `/backend` and the SvelteKit frontend under `/frontend` (consistent with current monorepo layout).
- Chromedp is available in the deployment environment or CI for smoke tests; if not, tests will use a non-JS-rendered fallback.
- The microservice will be used for low-volume, manual/QA scraping; it is not intended for large-scale crawling without additional rate-limiting and infrastructure.
- Operators can configure delays, user-agents, and timeouts via environment variables or a simple config file.

## Security & Privacy Considerations

- The scraper will not persist scraped content by default; by default, results are returned in-memory in the response.
- Operators must ensure any use of the scraper complies with target site terms-of-service and legal restrictions; the feature includes ethical defaults (robots.txt respect, delays) but does not enforce legal compliance.
- Avoid executing arbitrary user-provided JavaScript. Chromedp is used to render DOM but the service will not execute user-supplied scripts.

## Testing & QA

- Smoke test: HTTP request to GET /scrape?url=https://example.com&depth=1 must return HTTP 200 with valid ScrapeResult JSON (title, markdown, links array).
- Unit tests: small integration tests for the Colly/Chromedp pipeline can be included under `/backend/tests`.
- Manual QA: Use the frontend page to submit URLs that require JS rendering (simple SPA pages) and verify markdown output.

## Out of Scope

- High-volume distributed crawling, persistent storage, indexing, or scheduling of crawl jobs.
- Advanced content extraction heuristics (ML-based article extraction) beyond converting visible DOM text into markdown.

## Next Steps

1. Implement backend handler in `/backend` exposing GET /scrape.
2. Implement minimal SvelteKit page and server proxy to call backend and render results.
3. Add smoke test to backend tests and CI.
4. Iterate on extraction heuristics (optional follow-up).

---

SUCCESS: Spec ready for planning
```

# Feature Specification: [FEATURE NAME]

**Feature Branch**: `[###-feature-name]`  
**Created**: [DATE]  
**Status**: Draft  
**Input**: User description: "$ARGUMENTS"

## User Scenarios & Testing _(mandatory)_

<!--
  IMPORTANT: User stories should be PRIORITIZED as user journeys ordered by importance.
  Each user story/journey must be INDEPENDENTLY TESTABLE - meaning if you implement just ONE of them,
  you should still have a viable MVP (Minimum Viable Product) that delivers value.

  Assign priorities (P1, P2, P3, etc.) to each story, where P1 is the most critical.
  Think of each story as a standalone slice of functionality that can be:
  - Developed independently
  - Tested independently
  - Deployed independently
  - Demonstrated to users independently
-->

### User Story 1 - [Brief Title] (Priority: P1)

[Describe this user journey in plain language]

**Why this priority**: [Explain the value and why it has this priority level]

**Independent Test**: [Describe how this can be tested independently - e.g., "Can be fully tested by [specific action] and delivers [specific value]"]

**Acceptance Scenarios**:

1. **Given** [initial state], **When** [action], **Then** [expected outcome]
2. **Given** [initial state], **When** [action], **Then** [expected outcome]

---

### User Story 2 - [Brief Title] (Priority: P2)

[Describe this user journey in plain language]

**Why this priority**: [Explain the value and why it has this priority level]

**Independent Test**: [Describe how this can be tested independently]

**Acceptance Scenarios**:

1. **Given** [initial state], **When** [action], **Then** [expected outcome]

---

### User Story 3 - [Brief Title] (Priority: P3)

[Describe this user journey in plain language]

**Why this priority**: [Explain the value and why it has this priority level]

**Independent Test**: [Describe how this can be tested independently]

**Acceptance Scenarios**:

1. **Given** [initial state], **When** [action], **Then** [expected outcome]

---

[Add more user stories as needed, each with an assigned priority]

### Edge Cases

<!--
  ACTION REQUIRED: The content in this section represents placeholders.
  Fill them out with the right edge cases.
-->

- What happens when [boundary condition]?
- How does system handle [error scenario]?

## Requirements _(mandatory)_

<!--
  ACTION REQUIRED: The content in this section represents placeholders.
  Fill them out with the right functional requirements.
-->

### Functional Requirements

- **FR-001**: System MUST [specific capability, e.g., "allow users to create accounts"]
- **FR-002**: System MUST [specific capability, e.g., "validate email addresses"]
- **FR-003**: Users MUST be able to [key interaction, e.g., "reset their password"]
- **FR-004**: System MUST [data requirement, e.g., "persist user preferences"]
- **FR-005**: System MUST [behavior, e.g., "log all security events"]

_Example of marking unclear requirements:_

- **FR-006**: System MUST authenticate users via [NEEDS CLARIFICATION: auth method not specified - email/password, SSO, OAuth?]
- **FR-007**: System MUST retain user data for [NEEDS CLARIFICATION: retention period not specified]

### Key Entities _(include if feature involves data)_

- **[Entity 1]**: [What it represents, key attributes without implementation]
- **[Entity 2]**: [What it represents, relationships to other entities]

## Success Criteria _(mandatory)_

<!--
  ACTION REQUIRED: Define measurable success criteria.
  These must be technology-agnostic and measurable.
-->

### Measurable Outcomes

- **SC-001**: [Measurable metric, e.g., "Users can complete account creation in under 2 minutes"]
- **SC-002**: [Measurable metric, e.g., "System handles 1000 concurrent users without degradation"]
- **SC-003**: [User satisfaction metric, e.g., "90% of users successfully complete primary task on first attempt"]
- **SC-004**: [Business metric, e.g., "Reduce support tickets related to [X] by 50%"]
