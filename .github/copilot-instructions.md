# web-scrawl-go-sveltekit Development Guidelines

Auto-generated from all feature plans. Last updated: 2025-10-16

## Active Technologies
- Go 1.21+ (backend); SvelteKit (Svelte 5) + TypeScript (frontend) + Gin (HTTP server), Colly (crawler), Chromedp (headless rendering), GoQuery (parsing), shadcn-svelte (frontend UI), Tailwind CSS v4 (002-scrape-api-backend)
- N/A for MVP — in-memory response only; no persistent storage (002-scrape-api-backend)

- SvelteKit (latest) with Svelte 5 runes
- Bun 1.0+ (package manager and runtime)
- TypeScript 5.x
- Tailwind CSS v4 (use `@import 'tailwindcss';` syntax, NOT v3 `@tailwind` directives)
- shadcn-svelte for UI components
- Lucide Svelte for icons
- Go 1.21+ with Gin framework
- Colly v2, Chromedp, GoQuery for web scraping

## Project Structure

```
frontend/          # SvelteKit app (port 5173)
backend/           # Go Gin API (port 8080)
specs/             # Feature specifications
```

## Commands

### Frontend

- `cd frontend && bun dev` - Start development server
- `bun check` - TypeScript validation

### Backend

- `cd backend && go run main.go` - Start API server
- `go test ./... -v` - Run tests

## Code Style

- **Frontend**: Svelte 5 runes ($state, $props, $derived), TypeScript MANDATORY (all `.svelte` files use `lang="ts"`, server routes use `.ts` extension), Tailwind CSS v4
- **Backend**: Idiomatic Go, clear error handling, modular structure
- **CSS**: Tailwind v4 uses `@import 'tailwindcss';` in app.css (not v3's `@tailwind` directives)

## Svelte Development Workflow

1. **Use Svelte MCP for Documentation**: Always use `mcp_svelte_*` tools to fetch official Svelte 5 and SvelteKit documentation
2. **Validate Svelte Code**: ALWAYS run `mcp_svelte_svelte-autofixer` after writing/editing Svelte components to check for:
   - Svelte 5 compatibility issues
   - Incorrect runes usage
   - Migration suggestions from Svelte 4 patterns
3. **TypeScript Validation**: Run `bun check` after Svelte fixes to ensure TypeScript integrity

## Recent Changes
- 002-scrape-api-backend: Added Go 1.21+ (backend); SvelteKit (Svelte 5) + TypeScript (frontend) + Gin (HTTP server), Colly (crawler), Chromedp (headless rendering), GoQuery (parsing), shadcn-svelte (frontend UI), Tailwind CSS v4

- 001-monorepo-init: Added

<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->
