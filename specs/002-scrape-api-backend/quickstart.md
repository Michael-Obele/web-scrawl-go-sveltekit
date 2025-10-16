````markdown
# Quickstart: Scrape API (developer)

1. Start backend service

   ```bash
   cd /home/node/Documents/GitHub/web-scrawl-go-sveltekit/backend
   go run main.go
   ```
````

2. Start frontend (optional)

   ```bash
   cd /home/node/Documents/GitHub/web-scrawl-go-sveltekit/frontend
   bun dev
   ```

3. Smoke test (validate example.com)

   ```bash
   # From repo root
   curl -s "http://localhost:8080/scrape?url=https://example.com&depth=1" | jq .
   ```

Expected: JSON with `title`, `markdown`, `links` and `fetchedAt` fields.

Configuration:

- SCRAPER_DELAY_S: per-host delay (default 2)
- SCRAPER_USER_AGENTS: comma-separated list of user agents
- CHROMEDP_TIMEOUT_S: Chromedp render timeout (default 10)
- SCRAPER_TIMEOUT_S: overall scrape timeout (default 30)

```

```
