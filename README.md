web-scrawl-go-sveltekit/README.md
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

## Quick Start

Get up and running in minutes:

1. **Clone and install dependencies:**
   ```bash
   git clone <repository-url>
   cd web-scrawl-go-sveltekit
   cd frontend && bun install
   cd ../backend && go mod download
   ```

2. **Start the backend:**
   ```bash
   cd backend
   go run main.go
   ```
   Server runs on http://localhost:8080

3. **Start the frontend (in another terminal):**
   ```bash
   cd frontend
   bun dev
   ```
   App runs on http://localhost:5173

4. **Test the scraper:**
   - Visit http://localhost:5173/scrape
   - Enter `https://example.com` and depth `1`
   - Click "Start Scraping" to see results

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

### ✅ Core Infrastructure
- Monorepo initialization with proper dependency management
- Development environment with hot-reloading
- Health check endpoint validating frontend-to-backend communication

### ✅ Web Scraping API
- **Single-site scraping**: Extract page title, content, and outbound links from any website
- **Markdown conversion**: Convert HTML content to clean, readable markdown format
- **Link extraction**: Collect and normalize all outbound links with anchor text
- **URL validation**: Comprehensive input validation and error handling
- **Configurable timeouts**: Customizable scraping (30s) and rendering (10s) timeouts
- **Fallback rendering**: Automatic fallback from headless Chrome to static HTML when needed
- **Warning system**: Clear warnings for rendering fallbacks and other issues

### ✅ User Interface
- **Scrape form**: Input URL and depth parameters with validation
- **Progress indicators**: Visual feedback during scraping operations
- **Results display**: Tabbed interface showing markdown, raw HTML, and links table
- **Responsive design**: Mobile-friendly interface with shadcn-svelte components
- **Progressive enhancement**: Works without JavaScript (server-side rendering)

### 🚧 Planned Features
- **Ethical scraping controls**: Per-host rate limiting, full user-agent rotation, robots.txt compliance
- **Multi-depth crawling**: Support for crawling beyond depth=1 (currently depth=1 only)
- **Advanced content extraction**: Better content parsing and filtering algorithms
- **Export functionality**: Download scraped data in various formats
- **Graceful shutdown**: Proper server shutdown handling and context cancellation

## Development Status

### ✅ Completed (MVP)
- **Core scraping functionality**: Single-site scraping with title, content, and link extraction
- **API backend**: RESTful Go service with Gin framework, Colly crawler, and Chromedp rendering
- **Frontend UI**: SvelteKit interface with form validation, progress indicators, and results display
- **Testing**: Backend smoke tests and frontend type checking
- **Build system**: Both frontend and backend build successfully

### 🔄 In Progress
- **Ethical controls**: Basic delay configuration implemented; robots.txt checking and user-agent rotation pending
- **Content processing**: Markdown conversion working; HTML parsing could be enhanced

### 🎯 Next Priorities
1. **Implement robots.txt compliance** - Check and respect robots.txt files by default with override option
2. **Add user-agent rotation** - Cycle through configured user agents for each request
3. **Enhance per-host rate limiting** - Ensure proper delays between requests to the same domain
4. **Add multi-depth crawling** - Support crawling beyond depth=1 with proper link following
5. **Polish and documentation** - Add comprehensive logging, error handling, and API documentation

## API Endpoints

### GET /health
Health check endpoint.

**Response:**
```json
{
  "status": "ok",
  "timestamp": "2024-01-01T00:00:00Z"
}
```

### GET /scrape
Scrape a website and return structured data.

**Parameters:**
- `url` (required): The URL to scrape
- `depth` (optional): Crawl depth (currently supports depth=1 only)

**Response:**
```json
{
  "title": "Example Domain",
  "rawHtml": "<!doctype html>\n<html>\n<head>\n    <title>Example Domain</title>...",
  "markdown": "# Example Domain\n\nThis domain is for use in illustrative examples...",
  "links": [
    {
      "href": "https://www.iana.org/domains/example",
      "text": "More information..."
    }
  ],
  "fetchedAt": "2024-01-01T00:00:00Z",
  "warnings": []
}
```

## Testing

### Backend Testing

Run unit and integration tests:
```bash
cd backend
go test ./tests -v
```

Run tests with coverage:
```bash
cd backend
go test ./tests -v -cover
```

### Frontend Testing

Run TypeScript type checking:
```bash
cd frontend
bun check
```

Build validation:
```bash
cd frontend
bun run build
```

### Integration Testing

Test full-stack functionality:

1. Start backend server:
```bash
cd backend && go run main.go
```

2. In another terminal, test the API:
```bash
curl -s "http://localhost:8080/scrape?url=https://example.com&depth=1" | jq .
```

3. Start frontend and test UI:
```bash
cd frontend && bun dev
```
Visit `http://localhost:5173/scrape` and test scraping example.com

### CI Validation

Before committing, ensure:
- ✅ Backend tests pass: `cd backend && go test ./tests`
- ✅ Frontend type check passes: `cd frontend && bun check`
- ✅ Frontend builds successfully: `cd frontend && bun run build`
- ✅ Backend builds successfully: `cd backend && go build`

## Documentation

- [Monorepo Setup Guide](specs/001-monorepo-init/quickstart.md)
- [Scrape API Quickstart](specs/002-scrape-api-backend/quickstart.md)
- [Implementation Plans](specs/002-scrape-api-backend/plan.md)
- [API Contracts](specs/002-scrape-api-backend/contracts/)

## Contributing

This project follows a specification-driven development workflow:

1. Specifications are defined in `specs/`
2. Implementation follows the task breakdown
3. All changes are validated against constitutional principles

## License

[Add your license here]