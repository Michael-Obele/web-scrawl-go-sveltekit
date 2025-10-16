```markdown
# Research: 002-scrape-api-backend

## Overview

This document captures decisive choices for implementing the `/scrape` feature. The feature had no unresolved [NEEDS CLARIFICATION] markers: the user explicitly requested Colly, Chromedp and GoQuery and mandated Go + SvelteKit per repository constitution. Research therefore confirms and documents those choices and alternatives.

## Decisions

- Decision: Use Go backend (Gin) with Colly, Chromedp and GoQuery for crawling/rendering/extraction.

  - Rationale: These libraries are established in Go ecosystem for web scraping and are listed in project constitution and user request.
  - Alternatives considered: headless browsers via external services (Puppeteer remote), or JS-first scraper; rejected due to constitution (Go backend) and extra infra complexity.

- Decision: Respect robots.txt by default but allow operator override via environment variable.

  - Rationale: Ethical default reduces accidental misuse; override allows authorized uses (e.g., scraping own sites).

- Decision: Default timeouts and delays

  - per-host delay: 2s (configurable via SCRAPER_DELAY_S)
  - Chromedp timeout: 10s (configurable via CHROMEDP_TIMEOUT_S)
  - Overall scrape timeout: 30s (configurable via SCRAPER_TIMEOUT_S)

- Decision: User-agent rotation using a short list of common browser strings (configurable via SCRAPER_USER_AGENTS env var)

  - Rationale: Reduces simple fingerprinting and provides politeness; not intended for evasion.

- Decision: Output schema (ScrapeResult) includes title, markdown, links[], warnings[], fetchedAt
  - Rationale: Minimal useful set for preview and downstream processing.

## Next steps (Phase 1 inputs)

- Implement API contract (OpenAPI) for GET /scrape
- Implement ScrapeResult type in Go and basic handler wiring (Gin)
- Add smoke test against https://example.com
```
