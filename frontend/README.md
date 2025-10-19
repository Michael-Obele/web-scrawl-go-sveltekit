# Web Scraper Frontend

Modern SvelteKit frontend application built with Svelte 5, TypeScript, and Tailwind CSS v4.

## Requirements

- **Bun**: 1.0 or higher (package manager and runtime)
- **Node.js**: 18+ (alternative runtime, if not using Bun)

## Installation

```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
bun install
```

## Development

```bash
# Start development server (http://localhost:5173)
bun dev

# Start with network access
bun dev --host
```

## Build & Production

```bash
# Build for production
bun run build

# Preview production build
bun run preview

# Type check
bun check

# Type check in watch mode
bun check --watch
```

## Project Structure

```
frontend/
├── src/
│   ├── app.css              # Global styles with Tailwind v4
│   ├── app.html             # HTML template
│   ├── routes/              # File-based routing
│   │   ├── +layout.svelte   # Root layout (lang="ts")
│   │   ├── +page.svelte     # Home page (lang="ts")
│   │   ├── health/          # Health check page
│   │   │   ├── +page.server.ts  # Server-side proxy (TypeScript)
│   │   │   └── +page.svelte     # Health UI (lang="ts")
│   │   └── scrape/          # Web scraping interface
│   │       └── +page.svelte     # Scrape form and results (lang="ts")
│   └── lib/                 # Reusable components & utilities
│       ├── components/
│       │   └── ui/          # shadcn-svelte components
│       ├── remote/          # Remote functions for API calls
│       │   ├── index.ts     # Remote function exports
│       │   └── scraper.remote.ts  # Scraper API integration
│       └── types/           # TypeScript type definitions
│           └── scraper.ts   # Scraper data types
├── static/                  # Static assets
├── tailwind.config.js       # Tailwind v4 configuration
├── postcss.config.js        # PostCSS configuration
├── vite.config.ts           # Vite configuration
├── svelte.config.js         # SvelteKit configuration
├── tsconfig.json            # TypeScript configuration
└── package.json             # Dependencies and scripts
```

## Key Technologies

- **SvelteKit**: 2.x (full-stack framework)
- **Svelte**: 5.x with runes (`$state`, `$props`, `$derived`, `$effect`)
- **TypeScript**: 5.x with strict checking
- **Tailwind CSS**: v4.x (modern `@import` syntax)
- **shadcn-svelte**: UI component library
- **Lucide Svelte**: Icon library
- **Vite**: 7.x (build tool)
- **Valibot**: Form validation library

## Tailwind CSS v4

This project uses **Tailwind CSS v4** with the modern syntax:

```css
/* src/app.css */
@import "tailwindcss";
```

**Note**: Do NOT use Tailwind v3 directives (`@tailwind base;`, `@tailwind components;`, `@tailwind utilities;`).

## Svelte 5 Runes

This project uses Svelte 5 with runes for reactivity:

```svelte
<script>
  // State
  let count = $state(0);

  // Props
  let { title } = $props();

  // Derived values
  let doubled = $derived(count * 2);

  // Side effects
  $effect(() => {
    console.log('Count changed:', count);
  });
</script>
```

## Available Routes

- `/` - Home page with project overview
- `/health` - Backend health check status
- `/scrape` - Web scraping interface with form and results

## Remote Functions Architecture

The frontend uses SvelteKit's remote functions for type-safe API communication:

```typescript
// src/lib/remote/scraper.remote.ts
import { form } from "$app/server";
import * as v from "valibot";

const ScrapeInputSchema = v.object({
  url: v.pipe(
    v.string(),
    v.nonEmpty("URL is required"),
    v.url("Please enter a valid URL")
  ),
  depth: v.pipe(
    v.string(),
    v.transform(Number),
    v.number(),
    v.minValue(1, "Depth must be at least 1"),
    v.maxValue(3, "Depth cannot exceed 3")
  ),
});

export const scrapeWebsite = form(ScrapeInputSchema, async (data) => {
  const response = await fetch(`http://localhost:8080/scrape?url=${encodeURIComponent(data.url)}&depth=${data.depth}`);
  const result = await response.json();
  return { success: true, final: result, ...data };
});
```

## Components

UI components from shadcn-svelte are located in `src/lib/components/ui/`:

```bash
# Add new components
bunx shadcn-svelte@latest add button
bunx shadcn-svelte@latest add card
bunx shadcn-svelte@latest add table
bunx shadcn-svelte@latest add tabs
```

## Web Scraping Features

The `/scrape` route provides:

- **URL Input**: Form validation for website URLs
- **Depth Control**: Configurable crawl depth (1-3)
- **Progress Indicators**: Visual feedback during scraping
- **Results Display**: Tabbed interface showing:
  - Markdown content
  - Raw HTML
  - Links table
- **Error Handling**: User-friendly error messages
- **Responsive Design**: Mobile-friendly interface

## Configuration

The frontend connects to the backend via remote functions:

```typescript
// Progressive enhancement - works without JavaScript
// Server-side rendering with form actions
// Type-safe API calls with validation
```

## Linting & Formatting

```bash
# Type checking
bun check

# Format with Prettier (if configured)
bun format
```

## Troubleshooting

### Port 5173 already in use

```bash
# Kill process using port 5173
kill $(lsof -ti:5173)

# Or specify different port
bun dev --port 3000
```

### Bun installation issues

```bash
# Reinstall dependencies
rm -rf node_modules bun.lockb
bun install
```

### TypeScript errors

```bash
# Run type checking
bun check

# Clear SvelteKit cache
rm -rf .svelte-kit
bun dev
```

### Scraper not working

```bash
# Ensure backend is running on port 8080
cd ../backend && go run main.go

# Check backend health
curl http://localhost:8080/health

# Test scraper API
curl "http://localhost:8080/scrape?url=https://example.com&depth=1"
```

## Backend Integration

The frontend communicates with the Go backend (port 8080) via:

1. **Remote Functions**: Type-safe server-side API calls
2. **Progressive Enhancement**: Works without JavaScript
3. **CORS**: Backend configured to allow localhost:5173
4. **Error Handling**: Comprehensive error handling and validation
5. **TypeScript**: Strict typing throughout the application

## License

[Add your license here]