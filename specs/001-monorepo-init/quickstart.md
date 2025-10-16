# Quickstart Guide: Monorepo Initialization

**Feature**: 001-monorepo-init  
**Date**: 2025-10-16  
**Audience**: Developers setting up the web scraper project for the first time

---

## Prerequisites

Before starting, ensure you have the following installed:

- **Bun**: 1.0 or higher ([install instructions](https://bun.sh/))
- **Go**: 1.21 or higher ([install instructions](https://go.dev/doc/install))
- **Git**: Version control system
- **Terminal**: bash, zsh, or WSL2 on Windows
- **Code Editor**: VS Code recommended (with Svelte and Go extensions)

**Verify installations**:

```bash
bun --version   # Should output 1.0.0+
go version      # Should output go1.21+
git --version   # Any recent version
```

---

## Step 1: Initialize Frontend (SvelteKit + Bun)

### 1.1 Create SvelteKit Project

```bash
# From repository root
npx sv create frontend

# When prompted, select:
# - Which Svelte app template? → SvelteKit demo app (or minimal if you prefer)
# - Add type checking with TypeScript? → Yes, using TypeScript syntax
# - Select additional options → ESLint, Prettier (use spacebar to select)
# - Which package manager? → bun

# Navigate to frontend directory
cd frontend

# Install dependencies if not done during creation
bun install
```

**Expected output**: SvelteKit project structure created with `package.json` containing `"type": "module"`, `svelte.config.js`, `vite.config.ts`, and standard SvelteKit folder structure (`src/routes/`, `src/lib/`, etc.)

### 1.2 Install Tailwind CSS

```bash
# Still in frontend/ directory
npx sv add tailwindcss

# This command automatically:
# - Installs tailwindcss, postcss, autoprefixer
# - Creates tailwind.config.js and postcss.config.js
# - Adds Tailwind directives to app.css
# - Updates +layout.svelte to import app.css
```

**Verify Tailwind setup** - Check that `tailwind.config.js` exists:

```javascript
/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,js,svelte,ts}"],
  theme: {
    extend: {},
  },
  plugins: [],
};
```

**Verify Tailwind directives** - Check `src/app.css`:

```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

**Verify layout imports CSS** - Check `src/routes/+layout.svelte`:

```svelte
<script>
  import '../app.css';
</script>

<slot />
```

### 1.3 Install Lucide Svelte (Icons)

```bash
bun add lucide-svelte
```

### 1.4 Install shadcn-svelte (UI Components)

```bash
bunx shadcn-svelte@latest init

# When prompted, configure:
# - Which base color would you like to use? → Slate
# - Where is your global CSS file? → src/app.css
# - Configure the import alias for lib: → $lib (default)
# - Configure the import alias for components: → $lib/components (default)
# - Configure the import alias for utils: → $lib/utils (default)
# - Configure the import alias for hooks: → $lib/hooks (default)
# - Configure the import alias for ui: → $lib/components/ui (default)
```

This creates a `components.json` configuration file.

**Verify shadcn-svelte** - Install a test component:

```bash
bunx shadcn-svelte@latest add button
```

This creates `src/lib/components/ui/button/` with the Button component.

### 1.5 Configure TypeScript Checking

Edit `package.json` to ensure the `check` script exists:

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

### 1.6 Verify Frontend Setup

```bash
# Run TypeScript checks
bun check

# Start development server
bun dev
```

**Expected output**:

- TypeScript checks pass with no errors
- Dev server starts on `http://localhost:5173`
- Browser opens showing SvelteKit welcome page

**Troubleshooting**:

- If port 5173 is in use, kill the process: `lsof -ti:5173 | xargs kill -9`
- If Bun is not found, restart terminal after installation

---

## Step 2: Initialize Backend (Go + Gin)

### 2.1 Create Backend Directory Structure

```bash
# From repository root
mkdir -p backend/{handlers,middleware,tests}
cd backend
```

### 2.2 Initialize Go Module

```bash
go mod init github.com/yourusername/web-scraper-backend
# Replace yourusername with your actual GitHub username
# Example: go mod init github.com/Michael-Obele/web-scraper-backend
```

### 2.3 Install Go Dependencies

```bash
# Install all dependencies with -u flag for latest versions
go get -u github.com/gin-gonic/gin
go get -u github.com/gin-contrib/cors
go get -u github.com/gocolly/colly/v2
go get -u github.com/chromedp/chromedp
go get -u github.com/PuerkitoBio/goquery

# Clean up and verify dependencies
go mod tidy
```

**Expected output**: 
- Dependencies added to `go.mod` with version numbers
- `go.sum` created with dependency checksums
- `go mod tidy` removes unused dependencies and downloads missing ones

### 2.4 Create Main Entry Point

Create `backend/main.go`:

```go
package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/web-scraper-backend/handlers"
)

func main() {
	router := gin.Default()

	// CORS middleware
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	// Health check endpoint
	router.GET("/health", handlers.HealthCheck)

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
```

### 2.5 Create Health Check Handler

Create `backend/handlers/health.go`:

```go
package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":    "healthy",
		"service":   "web-scraper-backend",
		"timestamp": time.Now().Unix(),
	})
}
```

### 2.6 Create Smoke Test

Create `backend/tests/health_test.go`:

```go
package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/web-scraper-backend/handlers"
)

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/health", handlers.HealthCheck)

	req, _ := http.NewRequest("GET", "/health", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &body)

	if body["status"] != "healthy" {
		t.Errorf("Expected status 'healthy', got '%v'", body["status"])
	}
}
```

### 2.7 Verify Backend Setup

```bash
# Run tests
go test ./... -v

# Start backend server
go run main.go
```

**Expected output**:

- Tests pass: `PASS: TestHealthCheck`
- Server starts: `Starting server on :8080`
- Navigate to `http://localhost:8080/health` in browser → see JSON response

**Troubleshooting**:

- If port 8080 is in use: `lsof -ti:8080 | xargs kill -9`
- If imports fail, run `go mod tidy` to sync dependencies

---

## Step 3: Create Frontend Health Check Proxy

### 3.1 Create Server Route

Create `frontend/src/routes/health/+page.server.js`:

```javascript
/** @type {import('./$types').PageServerLoad} */
export async function load({ fetch }) {
  try {
    const response = await fetch("http://localhost:8080/health");
    const data = await response.json();
    return { health: data };
  } catch (error) {
    return {
      health: null,
      error: "Backend service is not available",
    };
  }
}
```

### 3.2 Create Health Check Page

Create `frontend/src/routes/health/+page.svelte`:

```svelte
<script>
	import { Check, X } from 'lucide-svelte';

	let { data } = $props();
</script>

<div class="container mx-auto p-8">
	<h1 class="text-3xl font-bold mb-6">Backend Health Check</h1>

	{#if data.health}
		<div class="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded flex items-center gap-2">
			<Check class="w-5 h-5" />
			<div>
				<p class="font-bold">Backend is healthy</p>
				<p class="text-sm">Service: {data.health.service}</p>
				<p class="text-sm">Timestamp: {new Date(data.health.timestamp * 1000).toLocaleString()}</p>
			</div>
		</div>
	{:else}
		<div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded flex items-center gap-2">
			<X class="w-5 h-5" />
			<div>
				<p class="font-bold">Backend is unavailable</p>
				<p class="text-sm">{data.error}</p>
			</div>
		</div>
	{/if}
</div>
```

### 3.3 Update Home Page

Edit `frontend/src/routes/+page.svelte`:

```svelte
<script>
	import { Button } from '$lib/components/ui/button';
</script>

<div class="container mx-auto p-8">
	<h1 class="text-4xl font-bold mb-4">Web Scraper</h1>
	<p class="text-gray-600 mb-6">Full-stack web scraping application built with SvelteKit and Go</p>

	<a href="/health">
		<Button>Check Backend Health</Button>
	</a>
</div>
```

---

## Step 4: Test End-to-End Integration

### 4.1 Start Both Services

**Terminal 1 - Backend**:

```bash
cd backend
go run main.go
```

**Terminal 2 - Frontend**:

```bash
cd frontend
bun dev
```

### 4.2 Verify Integration

1. **Open browser**: Navigate to `http://localhost:5173`
2. **Click "Check Backend Health" button**
3. **Verify**: Green success message shows backend is healthy
4. **Check browser console**: No CORS errors

### 4.3 Test Backend Failure Scenario

1. **Stop backend**: Kill the `go run` process in Terminal 1
2. **Refresh `/health` page**: Should show red error message
3. **Restart backend**: Start `go run main.go` again
4. **Refresh**: Should return to green success message

---

## Step 5: Run Validation Checks

### 5.1 Frontend Validation

```bash
cd frontend
bun check  # TypeScript validation
```

**Expected**: No errors

### 5.2 Backend Validation

```bash
cd backend
go test ./... -v  # Smoke tests
```

**Expected**: All tests pass

### 5.3 Manual Checklist

- [ ] Both services start without errors
- [ ] Frontend accessible at `http://localhost:5173`
- [ ] Backend accessible at `http://localhost:8080/health`
- [ ] Health check page shows backend status
- [ ] No CORS errors in browser console
- [ ] Tailwind CSS styles render correctly
- [ ] Lucide icons display on health page
- [ ] Button component from shadcn-svelte works

---

## Common Issues & Solutions

### Issue: `bunx: command not found`

**Solution**: Update Bun to latest version: `bun upgrade`

### Issue: `go: cannot find module providing package`

**Solution**: Run `go mod tidy` to sync dependencies

### Issue: CORS errors persist

**Solution**: Verify CORS middleware origin matches frontend URL exactly (including port)

### Issue: TypeScript errors in Svelte files

**Solution**: Run `bun check` and install missing type definitions: `bun add -D @types/node`

### Issue: Port already in use

**Solution**:

```bash
# Kill process on port 5173
lsof -ti:5173 | xargs kill -9

# Kill process on port 8080
lsof -ti:8080 | xargs kill -9
```

### Issue: shadcn-svelte components not rendering

**Solution**: Verify `app.css` is imported in `+layout.svelte` and Tailwind content paths include all Svelte files

---

## Next Steps

After successful initialization:

1. **Commit changes**: `git add . && git commit -m "feat: initialize monorepo structure"`
2. **Create feature branch**: Ready to implement scraping features
3. **Review constitution**: Validate setup against principles in `.specify/memory/constitution.md`
4. **Run `/speckit.tasks`**: Generate implementation tasks for your next feature

---

## Project Structure Overview

```
web-scraper/
├── backend/
│   ├── main.go                    # Entry point
│   ├── go.mod                     # Go dependencies
│   ├── handlers/
│   │   └── health.go              # Health check handler
│   ├── middleware/                # CORS (future)
│   └── tests/
│       └── health_test.go         # Smoke test
├── frontend/
│   ├── package.json               # Bun dependencies
│   ├── bun.lockb                  # Lock file
│   ├── src/
│   │   ├── routes/
│   │   │   ├── +page.svelte       # Home page
│   │   │   ├── +layout.svelte     # Layout with CSS import
│   │   │   └── health/
│   │   │       ├── +page.svelte   # Health check UI
│   │   │       └── +page.server.js # Proxy to backend
│   │   ├── lib/
│   │   │   └── components/ui/     # shadcn-svelte components
│   │   └── app.css                # Tailwind directives
│   ├── svelte.config.js           # SvelteKit config
│   ├── vite.config.ts             # Vite config
│   └── tailwind.config.js         # Tailwind config
└── specs/
    └── 001-monorepo-init/         # This feature's docs
```

---

**Quickstart Status**: Complete. Follow steps sequentially for successful setup. Refer to `research.md` for architectural decisions.
