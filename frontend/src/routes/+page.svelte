<script lang="ts">
  import { Button } from "$lib/components/ui/button";
  import * as Card from "$lib/components/ui/card";
  import { Check } from "lucide-svelte";

  const features = [
    {
      emoji: "🚀",
      title: "Fast & Efficient",
      description: "Built with Go for high-performance web scraping",
    },
    {
      emoji: "🎨",
      title: "Modern UI",
      description: "Clean interface powered by SvelteKit and shadcn-svelte",
    },
    {
      emoji: "🔒",
      title: "Ethical Scraping",
      description: "Respects robots.txt and implements rate limiting",
    },
  ];

  const scraperFeatures = [
    "Markdown content extraction",
    "Link discovery and mapping",
    "Configurable crawl depth (1-3 levels)",
    "Real-time progress tracking",
  ];

  const healthFeatures = [
    "Backend API status",
    "Service availability check",
    "Real-time health monitoring",
    "System diagnostics",
  ];
</script>

<div class="container mx-auto px-4 py-12 max-w-6xl">
  <div class="text-center mb-12">
    <h2 class="text-5xl font-bold tracking-tight mb-4">Web Scraper</h2>
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
          Extract content and links from any website with customizable crawl
          depth
        </Card.Description>
      </Card.Header>
      <Card.Content>
        <ul class="space-y-2 text-sm text-muted-foreground">
          {#each scraperFeatures as feature}
            <li class="flex items-center gap-2">
              <Check class="h-4 w-4 text-green-600 flex-shrink-0" />
              {feature}
            </li>
          {/each}
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
          {#each healthFeatures as feature}
            <li class="flex items-center gap-2">
              <Check class="h-4 w-4 text-green-600 flex-shrink-0" />
              {feature}
            </li>
          {/each}
        </ul>
      </Card.Content>
      <Card.Footer>
        <Button href="/health" variant="outline" class="w-full"
          >Check Health</Button
        >
      </Card.Footer>
    </Card.Root>
  </div>

  <!-- Features Section -->
  <div class="bg-muted rounded-lg p-8">
    <h2 class="text-2xl font-semibold mb-6 text-center">Features</h2>
    <div class="grid md:grid-cols-3 gap-6">
      {#each features as feature}
        <div class="text-center">
          <div class="text-3xl mb-2">{feature.emoji}</div>
          <h3 class="font-semibold mb-2">{feature.title}</h3>
          <p class="text-sm text-muted-foreground">
            {feature.description}
          </p>
        </div>
      {/each}
    </div>
  </div>
</div>
