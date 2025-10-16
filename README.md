# Web Scraper - SvelteKit & Go

<!-- release-badge-start -->

![Release](https://img.shields.io/badge/release-Unreleased-yellow)

<!-- release-badge-end -->
<!-- Example release badge (uncomment when you create release tags):
![Release](https://img.shields.io/badge/release-v0.1.0-yellow)
-->

A modern web scraping application built with SvelteKit frontend and Go backend in a monorepo structure.

## Project Overview

This project provides a full-stack web scraping solution:

- **Frontend**: SvelteKit + TypeScript + Tailwind CSS + shadcn-svelte
- **Backend**: Go + Gin + Colly + Chromedp + GoQuery
- **Runtime**: Bun (frontend), Go (backend)

## Monorepo Structure

```
├── frontend/          # SvelteKit application (port 5173)
├── backend/           # Go Gin API server (port 8080)
└── specs/            # Feature specifications and documentation
```

## Prerequisites

- **Bun**: 1.0 or higher ([install instructions](https://bun.sh/))
- **Go**: 1.21 or higher ([install instructions](https://go.dev/doc/install))
- **Git**: For version control

## Quick Start

For detailed setup instructions, see [specs/001-monorepo-init/quickstart.md](specs/001-monorepo-init/quickstart.md)

### Start Frontend (Development)

```bash
cd frontend
bun install  # First time only
bun dev      # Starts on http://localhost:5173
```

### Start Backend (Development)

```bash
cd backend
go mod download  # First time only
go run main.go   # Starts on http://localhost:8080
```

## Development

Both services can run concurrently for full-stack development:

```bash
# Terminal 1
cd frontend && bun dev

# Terminal 2
cd backend && go run main.go
```

## Architecture

- **Frontend**: SvelteKit server routes proxy requests to Go backend (no direct client-to-backend communication)
- **Backend**: RESTful API with Gin framework, CORS configured for frontend origin
- **Communication**: HTTP/JSON between services

## Features

- ✅ Monorepo initialization with proper dependency management
- ✅ Development environment with hot-reloading
- ✅ Health check endpoint validating frontend-to-backend communication
- 🚧 Web scraping capabilities (coming soon)

## Documentation

- [Quick Start Guide](specs/001-monorepo-init/quickstart.md)
- [Implementation Plan](specs/001-monorepo-init/plan.md)
- [API Contracts](specs/001-monorepo-init/contracts/)

## Contributing

This project follows a specification-driven development workflow:

1. Specifications are defined in `specs/`
2. Implementation follows the task breakdown
3. All changes are validated against constitutional principles

## License

[Add your license here]
