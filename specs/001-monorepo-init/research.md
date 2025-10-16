# Research: Monorepo Initialization

**Feature**: 001-monorepo-init  
**Date**: 2025-10-16  
**Status**: Complete

## Research Tasks

This document consolidates research findings for unknowns identified during planning. Each decision includes rationale, alternatives considered, and selected approach.

---

## 1. SvelteKit + Bun Initialization Best Practices

**Task**: Determine the optimal way to initialize a SvelteKit project with Bun runtime, ensuring TypeScript, Vite, and Tailwind CSS are properly configured.

### Decision

Use **official SvelteKit CLI (`sv create`)** with Bun package manager:

```bash
npx sv create frontend
# Select: SvelteKit demo app (or minimal), TypeScript, ESLint, Prettier, Bun
cd frontend
bun install  # If not done during creation
npx sv add tailwindcss  # Automated Tailwind setup
```

### Rationale

1. **Official Support**: `sv create` is the official Svelte CLI maintained by the Svelte team with proper TypeScript configuration
2. **Bun Integration**: The CLI supports Bun as a first-class package manager option during project creation
3. **Vite Integration**: SvelteKit uses Vite by default; no additional configuration needed
4. **Automated Tailwind**: `sv add tailwindcss` handles all Tailwind setup automatically (install, config, directives)
5. **Modern Best Practices**: Uses latest Svelte 5 conventions with runes and proper project structure

### Alternatives Considered

- **`bun create svelte@latest`**: Deprecated in favor of `sv create` according to official Svelte documentation
- **Manual Setup**: Rejected due to time cost and potential for misconfiguration (violates velocity-first principle)
- **Custom Boilerplate**: Rejected as it adds maintenance burden and doesn't leverage community-maintained templates
- **Vite + Svelte (not SvelteKit)**: Rejected because we need server-side routes for proxying to Go backend

### Implementation Notes

- Run `bun check` immediately after initialization to validate TypeScript integrity
- Verify `package.json` includes `"type": "module"` for ES modules support
- The `sv add tailwindcss` command automatically:
  - Installs tailwindcss, postcss, autoprefixer
  - Creates tailwind.config.js with proper content paths
  - Adds Tailwind directives to app.css
  - Ensures +layout.svelte imports app.css

---

## 2. shadcn-svelte Installation Order & Configuration

**Task**: Determine when and how to install shadcn-svelte components to avoid conflicts with Tailwind CSS setup.

### Decision

Use **`bunx shadcn-svelte@latest init`** AFTER Tailwind CSS is fully configured:

```bash
# After Tailwind is configured
bunx shadcn-svelte@latest init
# Follow prompts: Select Slate base color, confirm aliases

# Add specific components
bunx shadcn-svelte@latest add button
```

### Rationale

1. **Official Installation Method**: According to [shadcn-svelte.com](https://shadcn-svelte.com/docs/installation/sveltekit), `shadcn-svelte@latest init` is the official CLI command
2. **Dependency Requirement**: shadcn-svelte expects Tailwind CSS to be already configured and will modify `tailwind.config.js`
3. **Configuration Merging**: The shadcn-svelte init script adds its own theme tokens to Tailwind config; running it first causes conflicts
4. **Consistent with Official Docs**: The official documentation shows using `bunx` (or `pnpm dlx`/`npx`) with the full package name

### Alternatives Considered

- **`sv add shadcn-svelte`**: Does not exist - shadcn-svelte is not integrated into the Svelte CLI
- **Simultaneous Installation**: Rejected due to race conditions and config file conflicts
- **shadcn-svelte First**: Rejected because the CLI fails if Tailwind is not detected
- **Manual Component Setup**: Rejected as it bypasses accessibility defaults and theme configuration

### Implementation Notes

- Creates `components.json` configuration file in project root
- Components will be placed in `src/lib/components/ui/` by default
- The CLI automatically handles TypeScript types and Svelte 5 runes compatibility
- Use `bunx shadcn-svelte@latest add <component-name>` to add individual components
- Configure import aliases during init to match project structure ($lib, $lib/components, etc.)

---

## 3. Go Gin Server Structure for Monorepo

**Task**: Design the Go backend directory structure that supports modular development, clear separation of concerns, and easy testing.

### Decision

Use **standard Go project layout** with domain-focused packages:

```
backend/
├── main.go           # Entry point
├── handlers/         # HTTP handlers (one file per domain)
├── middleware/       # Reusable middleware
├── models/           # Data structures (future)
├── services/         # Business logic (future)
└── tests/            # Test files
```

### Rationale

1. **Community Standard**: Aligns with golang-standards/project-layout and Gin's own documentation
2. **Modularity**: Clear separation enables parallel development and independent testing (constitutional principle II)
3. **Future Scaling**: Structure accommodates future scraping services without refactoring
4. **Simplicity**: Avoids over-engineering; no complex DI frameworks or unnecessary abstractions

### Alternatives Considered

- **Flat Structure** (all code in main.go): Rejected as it violates modularity principle and becomes unmaintainable beyond 200 lines
- **Domain-Driven Design** (bounded contexts): Rejected as premature for MVP; defer until scraping features are built
- **Hexagonal Architecture**: Rejected as over-engineered for a simple API server with no complex business logic yet

### Implementation Notes

- Keep `main.go` under 50 lines: setup router, register middleware, start server
- Use `handlers/health.go` pattern: one handler function per file for now
- CORS middleware in `middleware/cors.go` with explicit origin allowlist for development

---

## 4. CORS Middleware Configuration for Development

**Task**: Configure CORS to allow SvelteKit server routes to communicate with Go backend during development without errors.

### Decision

Use **Gin's CORS middleware** with explicit development origin:

```go
// middleware/cors.go
import "github.com/gin-contrib/cors"

config := cors.Config{
    AllowOrigins:     []string{"http://localhost:5173"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Content-Type", "Authorization"},
    AllowCredentials: true,
}
router.Use(cors.New(config))
```

### Rationale

1. **Security by Default**: Explicit origin allowlist prevents accidental exposure to all origins (\*)
2. **SvelteKit Pattern**: Server routes make requests from Node.js context, not browser, so we need permissive CORS
3. **Development Convenience**: Allows hot-reload without CORS errors
4. **Production Ready**: Easy to switch origins to production URLs (Vercel domain) in environment variables

### Alternatives Considered

- **Allow All Origins (\*)**: Rejected due to security implications, even in development
- **No CORS Middleware**: Rejected because preflight OPTIONS requests would fail
- **Manual CORS Headers**: Rejected as error-prone and less maintainable than using established library

### Implementation Notes

- Use environment variable `ALLOWED_ORIGINS` to switch between development and production
- For production: `ALLOWED_ORIGINS=https://your-app.vercel.app`
- Consider adding request logging in CORS middleware for debugging

---

## 5. SvelteKit Server Route Proxy Pattern

**Task**: Design the pattern for SvelteKit server routes to proxy requests to the Go backend, ensuring type safety and error handling.

### Decision

Use **SvelteKit load functions** in `+page.server.js` with native fetch:

```javascript
// src/routes/health/+page.server.js
export async function load({ fetch }) {
  const response = await fetch("http://localhost:8080/health");
  const data = await response.json();
  return { health: data };
}
```

### Rationale

1. **Built-in Feature**: SvelteKit's server-side `load` functions run on server, avoiding CORS entirely from client perspective
2. **Type Safety**: Can add TypeScript interfaces for response data
3. **Error Handling**: SvelteKit automatically handles errors and can show error pages
4. **SSR Compatible**: Pattern works for both server-side rendering and client-side navigation

### Alternatives Considered

- **SvelteKit API Routes** (+server.js): Rejected as unnecessary indirection; load functions are simpler for GET requests
- **Direct Client Fetch**: Rejected because it requires CORS preflight and exposes backend URL to client
- **Proxy Middleware**: Rejected as it adds complexity (Vite proxy config) when load functions suffice

### Implementation Notes

- Use `http://localhost:8080` for development; switch to `BACKEND_URL` environment variable for production
- Add error boundary in `+error.svelte` to handle backend failures gracefully
- For POST requests (future features), use form actions in `+page.server.js`

---

## 6. Health Check Endpoint Design

**Task**: Design a simple health check endpoint that validates backend connectivity and follows REST conventions.

### Decision

**GET /health** endpoint returning JSON:

```go
// handlers/health.go
func HealthCheck(c *gin.Context) {
    c.JSON(200, gin.H{
        "status": "healthy",
        "service": "web-scraper-backend",
        "timestamp": time.Now().Unix(),
    })
}
```

### Rationale

1. **REST Convention**: `/health` is standard path for health checks (used by Docker, Kubernetes, monitoring tools)
2. **Simple Response**: JSON format with explicit status field enables easy validation
3. **Extensibility**: Timestamp field allows future monitoring of uptime; can add version, dependencies later
4. **HTTP 200**: Always returns success when server is running (failures are 5xx errors from network/server down)

### Alternatives Considered

- **GET /api/health**: Rejected as `/health` is more universal and not part of application API
- **Complex Health Checks** (DB, external services): Rejected as premature; no dependencies to check yet
- **HEAD Request**: Rejected as less debuggable; JSON response body provides more information

### Implementation Notes

- Register route in `main.go`: `router.GET("/health", handlers.HealthCheck)`
- Write smoke test: assert response status 200 and `status` field equals "healthy"
- Frontend displays health data on `/health` page for developer visibility

---

## 7. Dependency Version Selection

**Task**: Select specific versions of key dependencies that are stable, well-documented, and compatible with each other.

### Decision

**Frontend:**

- SvelteKit: Latest stable (via `@latest` in CLI)
- Bun: 1.0+ (whatever developer has installed)
- TypeScript: 5.x (bundled with SvelteKit template)
- Tailwind CSS: 3.x (latest stable)
- Lucide Svelte: Latest (via `bun add lucide-svelte`)
- shadcn-svelte: Latest via CLI tool

**Backend:**

- Go: 1.21+ (minimum for modern features; developer's installed version)
- Gin: `github.com/gin-gonic/gin` latest (via `go get`)
- Colly: `github.com/gocolly/colly/v2` (v2 for Go modules support)
- Chromedp: `github.com/chromedp/chromedp` latest
- GoQuery: `github.com/PuerkitoBio/goquery` latest

### Rationale

1. **Latest Stable**: Use latest versions to get bug fixes and security patches (velocity principle)
2. **SemVer Trust**: All selected dependencies follow semantic versioning; patch updates are safe
3. **Community Adoption**: All dependencies have 1k+ GitHub stars and active maintenance
4. **Breaking Changes**: Lock files (`bun.lockb`, `go.sum`) prevent surprise updates

### Alternatives Considered

- **Pin Exact Versions**: Rejected as it creates maintenance burden for security updates
- **Use Alpha/Beta**: Rejected for stability reasons (production deployment requirement)
- **Vendoring**: Rejected as unnecessary for MVP; revisit if supply chain becomes a concern

### Implementation Notes

- Run `bun update` quarterly to get latest patches
- Monitor dependency security advisories via GitHub Dependabot
- Document breaking changes in `CHANGELOG.md` when major version updates occur

---

## 8. Development Workflow: Run Scripts & Commands

**Task**: Define convenient commands for developers to start both services, run tests, and build for production.

### Decision

**Frontend Scripts** (package.json):

```json
{
  "scripts": {
    "dev": "vite dev --port 5173",
    "build": "vite build",
    "preview": "vite preview",
    "check": "svelte-kit sync && svelte-check --tsconfig ./tsconfig.json"
  }
}
```

**Backend Commands** (documented in README):

```bash
# Development
cd backend && go run main.go

# Tests
cd backend && go test ./... -v

# Build
cd backend && go build -o bin/server main.go
```

### Rationale

1. **Standard Conventions**: `dev`, `build`, `check` scripts are SvelteKit conventions; developers expect them
2. **Explicit Ports**: Hard-code port 5173 in script to ensure consistency (avoids dynamic port allocation)
3. **Go Simplicity**: Go doesn't need a task runner; `go run` and `go test` are sufficient
4. **Discoverability**: Commands documented in README for onboarding

### Alternatives Considered

- **Monorepo Task Runner** (Turborepo, Nx): Rejected as over-engineered for 2 services; adds complexity
- **Makefile**: Rejected as less cross-platform (Windows compatibility issues without WSL)
- **Docker Compose**: Rejected for development (adds overhead); defer to production deployment

### Implementation Notes

- Add `"type": "module"` to `package.json` for ES modules support
- Document commands in `specs/001-monorepo-init/quickstart.md`
- Consider adding `concurrently` package in root for running both services simultaneously (optional)

---

## Summary of Decisions

| Decision Area   | Selected Approach                    | Key Reason                         |
| --------------- | ------------------------------------ | ---------------------------------- |
| SvelteKit Init  | Official CLI with Bun                | Velocity + official support        |
| shadcn-svelte   | Install after Tailwind               | Dependency order requirement       |
| Go Structure    | Standard layout with handlers/       | Modularity + simplicity            |
| CORS Config     | Gin middleware with explicit origins | Security + development convenience |
| Proxy Pattern   | SvelteKit load functions             | Built-in SSR support               |
| Health Endpoint | GET /health with JSON                | REST convention + extensibility    |
| Versions        | Latest stable via SemVer             | Security + modern features         |
| Run Scripts     | Standard npm scripts + Go commands   | Developer conventions              |

All research findings align with constitutional principles (velocity-first, code quality, no over-engineering). Proceed to Phase 1: Design & Contracts.
