<script lang="ts">
  import { Button } from "$lib/components/ui/button";
  import * as Card from "$lib/components/ui/card";
</script>

<div class="container mx-auto px-4 py-12 max-w-6xl">
  <div class="text-center mb-12">
    <h1 class="text-5xl font-bold tracking-tight mb-4">Web Scraper</h1>
    <p class="text-xl text-muted-foreground max-w-2xl mx-auto">
      A modern web scraping application built with SvelteKit frontend and Go
      backend. Extract data from websites efficiently and ethically.
    </p>
  </div>

  <div class="grid md:grid-cols-2 gap-6 mb-12">
    <!-- Scraper Card -->
    <Card.Root>
      <Card.Header>
        <Card.Title>Web Scraper</Card.Title>
        <Card.Description>
          Extract content and links from any website with customizable crawl depth
        </Card.Description>
      </Card.Header>
      <Card.Content>
        <ul class="space-y-2 text-sm text-muted-foreground">
          <li>✓ Markdown content extraction</li>
          <li>✓ Link discovery and mapping</li>
          <li>✓ Configurable crawl depth (1-3 levels)</li>
          <li>✓ Real-time progress tracking</li>
        </ul>
      </Card.Content>
      <Card.Footer>
        <Button href="/scrape" class="w-full">Start Scraping</Button>
      </Card.Footer>
    </Card.Root>

    <!-- Health Check Card -->
    <Card.Root>
      <Card.Header>
        <Card.Title>System Health</Card.Title>
        <Card.Description>
          Monitor the backend API status and performance
        </Card.Description>
      </Card.Header>
      <Card.Content>
        <ul class="space-y-2 text-sm text-muted-foreground">
          <li>✓ Backend API status</li>
          <li>✓ Service availability check</li>
          <li>✓ Real-time health monitoring</li>
          <li>✓ System diagnostics</li>
        </ul>
      </Card.Content>
      <Card.Footer>
        <Button href="/health" variant="outline" class="w-full">Check Health</Button>
      </Card.Footer>
    </Card.Root>
  </div>

  <!-- Features Section -->
  <div class="bg-muted rounded-lg p-8">
    <h2 class="text-2xl font-semibold mb-6 text-center">Features</h2>
    <div class="grid md:grid-cols-3 gap-6">
      <div class="text-center">
        <div class="text-3xl mb-2">🚀</div>
        <h3 class="font-semibold mb-2">Fast & Efficient</h3>
        <p class="text-sm text-muted-foreground">
          Built with Go for high-performance web scraping
        </p>
      </div>
      <div class="text-center">
        <div class="text-3xl mb-2">🎨</div>
        <h3 class="font-semibold mb-2">Modern UI</h3>
        <p class="text-sm text-muted-foreground">
          Clean interface powered by SvelteKit and shadcn-svelte
        </p>
      </div>
      <div class="text-center">
        <div class="text-3xl mb-2">🔒</div>
        <h3 class="font-semibold mb-2">Ethical Scraping</h3>
        <p class="text-sm text-muted-foreground">
          Respects robots.txt and implements rate limiting
        </p>
      </div>
    </div>
  </div>
</div>
