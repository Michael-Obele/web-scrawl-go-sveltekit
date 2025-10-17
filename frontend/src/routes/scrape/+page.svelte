<script lang="ts">
  import { enhance } from "$app/forms";
  import { scrapeWebsite } from "$lib/remote/scraper.remote";
  import { Button } from "$lib/components/ui/button";
  import { Input } from "$lib/components/ui/input";
  import { Label } from "$lib/components/ui/label";
  import * as Table from "$lib/components/ui/table";
  import * as Card from "$lib/components/ui/card";

  // Default values for the form
  let initialUrl = "https://dexie.org/docs/Tutorial/Svelte";
  let initialDepth = 1;
</script>

<div class="container mx-auto max-w-4xl px-4 py-8">
  <div class="mb-8 text-center">
    <h1 class="text-3xl font-bold tracking-tight mb-2">Web Scraper</h1>
    <p class="text-muted-foreground">
      Extract content and links from any website
    </p>
  </div>

  <Card.Root class="mb-6">
    <Card.Header>
      <Card.Title>Scrape Configuration</Card.Title>
      <Card.Description>
        Enter the URL and crawl depth to extract content
      </Card.Description>
    </Card.Header>
    <Card.Content class="space-y-6">
      <form {...scrapeWebsite} class="space-y-6">
        <div class="space-y-2">
          <Label for="scrape-url">Website URL</Label>
          <Input
            id="scrape-url"
            {...scrapeWebsite.fields.url.as("url")}
            placeholder="https://example.com"
            class="w-full"
          />
          {#each scrapeWebsite.fields.url.issues() as issue, i (i)}
            <p class="text-sm text-destructive">{issue.message}</p>
          {/each}
          <p class="text-sm text-muted-foreground">
            Enter the URL you want to scrape
          </p>
        </div>

        <div class="space-y-2">
          <Label for="scrape-depth">Crawl Depth</Label>
          <Input
            id="scrape-depth"
            {...scrapeWebsite.fields.depth.as("text")}
            class="w-24"
          />
          {#each scrapeWebsite.fields.depth.issues() as issue, i (i)}
            <p class="text-sm text-destructive">{issue.message}</p>
          {/each}
          <p class="text-sm text-muted-foreground">
            How many levels deep to crawl (1-3)
          </p>
        </div>

        <Button type="submit" disabled={!!scrapeWebsite.pending} class="w-full">
          {scrapeWebsite.pending ? "Scraping..." : "Start Scraping"}
        </Button>
      </form>
    </Card.Content>
  </Card.Root>

  <!-- Loading State -->
  {#if scrapeWebsite.pending}
    <Card.Root class="mb-6">
      <Card.Content class="pt-6">
        <div class="space-y-2">
          <div class="h-2 bg-primary/20 rounded-full overflow-hidden">
            <div class="h-full bg-primary w-1/3 animate-pulse"></div>
          </div>
          <p class="text-sm text-center text-muted-foreground">
            Fetching and processing content...
          </p>
        </div>
      </Card.Content>
    </Card.Root>
  {/if}

  <!-- Results Display -->
  {#if scrapeWebsite.result?.final}
    <div class="space-y-6">
      <!-- Title Section -->
      <Card.Root>
        <Card.Header>
          <Card.Title>{scrapeWebsite.result.final.title}</Card.Title>
          <Card.Description>
            Scraped at: {new Date(
              scrapeWebsite.result.final.fetchedAt
            ).toLocaleString()}
          </Card.Description>
        </Card.Header>
      </Card.Root>

      <!-- Markdown Content -->
      <Card.Root>
        <Card.Header>
          <Card.Title>Page Content</Card.Title>
          <Card.Description>
            Extracted markdown content from the page
          </Card.Description>
        </Card.Header>
        <Card.Content>
          <div
            class="markdown-content bg-muted rounded-md p-4 overflow-x-auto max-h-96"
          >
            <pre
              class="whitespace-pre-wrap break-words text-sm font-mono leading-relaxed">{scrapeWebsite
                .result.final.markdown}</pre>
          </div>
        </Card.Content>
      </Card.Root>

      <!-- Links Table -->
      <Card.Root>
        <Card.Header>
          <Card.Title>Discovered Links</Card.Title>
          <Card.Description>
            Found {scrapeWebsite.result.final.links.length} link{scrapeWebsite
              .result.final.links.length !== 1
              ? "s"
              : ""} on the page
          </Card.Description>
        </Card.Header>
        <Card.Content>
          <div class="rounded-md border overflow-hidden">
            <Table.Root>
              <Table.Header>
                <Table.Row>
                  <Table.Head class="w-1/3">Link Text</Table.Head>
                  <Table.Head>URL</Table.Head>
                </Table.Row>
              </Table.Header>
              <Table.Body>
                {#each scrapeWebsite.result.final.links.slice(0, 50) as link, i (link.href + i)}
                  <Table.Row>
                    <Table.Cell class="font-medium">
                      {link.text || "(no text)"}
                    </Table.Cell>
                    <Table.Cell>
                      <a
                        href={link.href}
                        target="_blank"
                        rel="noopener noreferrer"
                        class="text-primary hover:underline text-sm break-all"
                      >
                        {link.href}
                      </a>
                    </Table.Cell>
                  </Table.Row>
                {/each}
              </Table.Body>
            </Table.Root>
          </div>
          {#if scrapeWebsite.result.final.links.length > 50}
            <p class="text-sm text-muted-foreground mt-3">
              Showing first 50 of {scrapeWebsite.result.final.links.length} links
            </p>
          {/if}
        </Card.Content>
      </Card.Root>

      <!-- Warnings Section -->
      {#if scrapeWebsite.result.final.warnings?.length}
        <Card.Root class="border-yellow-200 dark:border-yellow-800">
          <Card.Header>
            <Card.Title class="text-yellow-800 dark:text-yellow-200">
              Warnings
            </Card.Title>
            <Card.Description class="text-yellow-700 dark:text-yellow-300">
              {scrapeWebsite.result.final.warnings.length} warning{scrapeWebsite
                .result.final.warnings.length !== 1
                ? "s"
                : ""} detected during scraping
            </Card.Description>
          </Card.Header>
          <Card.Content>
            <ul class="text-sm text-yellow-700 dark:text-yellow-300 space-y-1">
              {#each scrapeWebsite.result.final.warnings as warning, i (i)}
                <li>• {warning}</li>
              {/each}
            </ul>
          </Card.Content>
        </Card.Root>
      {/if}
    </div>
  {/if}
</div>

<style>
  .markdown-content {
    word-wrap: break-word;
    overflow-wrap: break-word;
    word-break: break-word;
  }

  .markdown-content pre {
    white-space: pre-wrap;
    word-wrap: break-word;
    overflow-wrap: break-word;
  }
</style>
