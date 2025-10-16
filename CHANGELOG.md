# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

- chore: add release changelog workflow and update README with release badge
- chore(ci): add README release badge and changelog dry-run script
- chore: add changelog verification workflow to ensure CHANGELOG.md updates
- refactor: update package.json and app.css for improved Tailwind CSS integration; enhance task list formatting in tasks.md
- Merge PR #2: enforce TypeScript & Svelte MCP
- Merging feature branch 001-monorepo-init which enforces TypeScript across frontend and adds Svelte MCP validation policies.
- refactor: enforce TypeScript across frontend and add Svelte MCP validation
- - Convert server routes from .js to .ts (+page.server.ts)
- - Add lang=ts to all .svelte files
- - Update constitution to mandate TypeScript usage
- - Add Svelte MCP documentation and autofixer requirements
- - Update copilot instructions with Svelte validation workflow
- - Update all documentation (README, quickstart) to reflect TypeScript
- TypeScript validation: bun check passes (0 errors)
- Svelte validation: svelte-autofixer passes (0 issues)
- docs: mark all 72 tasks as complete
- feat: initialize monorepo with SvelteKit and Go
- - Setup SvelteKit frontend with Svelte 5 runes, TypeScript, Tailwind v4
- - Setup Go backend with Gin, CORS, and web scraping dependencies
- - Implement health check endpoint with frontend-to-backend integration
- - Add comprehensive documentation (README files, quickstart guide)
- - Configure environment with .env.example files
- - Establish development workflow and testing infrastructure
- All 72 tasks complete per specs/001-monorepo-init/tasks.md
- feat: update quickstart and research documentation for SvelteKit and Go setup; add task list for monorepo initialization
- feat: initialize monorepo structure with frontend and backend services
- - Added implementation plan for monorepo initialization (plan.md)
- - Created quickstart guide for setting up the web scraper project (quickstart.md)
- - Documented research findings and decisions for project setup (research.md)
- Merge pull request #1 from Michael-Obele/master
- Add initial constitution and specification quality checklist for mono…
- Add initial constitution and specification quality checklist for monorepo setup
- Initial commit from Specify template

- Initial monorepo initialization (2025-10-16)
  - Frontend: SvelteKit (Svelte 5) + TypeScript, Tailwind CSS v4, shadcn-svelte, Lucide Svelte
  - Backend: Go + Gin; health endpoint and basic server/tests
  - Docs: README, frontend/backend READMEs, specs/001-monorepo-init
  - CI: Added changelog-check workflow to require CHANGELOG.md updates on PRs to `master`

<!-- Keep this file lightweight. Contributors should add short notes under "Unreleased" for their PRs. -->
