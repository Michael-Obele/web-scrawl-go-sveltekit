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
│   │   └── health/          # Health check page
│   │       ├── +page.server.ts  # Server-side proxy (TypeScript)
│   │       └── +page.svelte     # Health UI (lang="ts")
│   └── lib/                 # Reusable components & utilities
│       └── components/
│           └── ui/          # shadcn-svelte components
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

- `/` - Home page
- `/health` - Backend health check status

## Configuration

The frontend connects to the backend via SvelteKit's server routes (+page.server.ts):

```typescript
// Proxy pattern example (TypeScript)
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
  const response = await fetch("http://localhost:8080/health");
  const data = await response.json();
  return { health: data };
};
```

## Components

UI components from shadcn-svelte are located in `src/lib/components/ui/`:

```bash
# Add new components
bunx shadcn-svelte@latest add button
bunx shadcn-svelte@latest add card
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

## Backend Integration

The frontend communicates with the Go backend (port 8080) via:

1. **SvelteKit Server Routes**: +page.server.ts files (TypeScript) proxy API calls
2. **CORS**: Backend configured to allow localhost:5173
3. **Error Handling**: All API calls include error handling and user-friendly messages
4. **TypeScript**: All `.svelte` files use `lang="ts"` and server routes use `.ts` extension

## License

[Add your license here]
