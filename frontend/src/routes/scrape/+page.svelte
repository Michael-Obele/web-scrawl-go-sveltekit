<script lang="ts">
    import { scrapeWebsite } from "$lib/remote/scraper.remote";
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input";
    import { Label } from "$lib/components/ui/label";
    import * as Table from "$lib/components/ui/table";
    import * as Card from "$lib/components/ui/card";
    import * as Tabs from "$lib/components/ui/tabs";
    import { Scale } from "lucide-svelte";
</script>

<div class="container mx-auto max-w-6xl px-4 py-8">
    <div class="mb-12 text-center">
        <h1 class="text-4xl font-bold tracking-tight mb-4">Web Scraper</h1>
        <p class="text-xl text-muted-foreground">
            Extract content, links, and raw HTML from any website with ease.
        </p>
    </div>

    {#if scrapeWebsite.result?.final}
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div>
                <Card.Root>
                    <Card.Header>
                        <Card.Title>Scrape Configuration</Card.Title>
                        <Card.Description>
                            Enter the URL and crawl depth to extract content.
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
                                    <p class="text-sm text-destructive">
                                        {issue.message}
                                    </p>
                                {/each}
                            </div>

                            <div class="space-y-2">
                                <Label for="scrape-depth">Crawl Depth</Label>
                                <Input
                                    id="scrape-depth"
                                    {...scrapeWebsite.fields.depth.as("text")}
                                    class="w-24"
                                />
                                {#each scrapeWebsite.fields.depth.issues() as issue, i (i)}
                                    <p class="text-sm text-destructive">
                                        {issue.message}
                                    </p>
                                {/each}
                                <p class="text-sm text-muted-foreground">
                                    How many levels deep to crawl (1-3)
                                </p>
                            </div>

                            <Button
                                type="submit"
                                disabled={!!scrapeWebsite.pending}
                                class="w-full"
                            >
                                {scrapeWebsite.pending
                                    ? "Scraping..."
                                    : "Start Scraping"}
                            </Button>
                        </form>
                    </Card.Content>
                </Card.Root>
            </div>

            <div>
                {#if scrapeWebsite.pending}
                    <Card.Root>
                        <Card.Content class="pt-6">
                            <div class="space-y-2">
                                <div
                                    class="h-2 bg-primary/20 rounded-full overflow-hidden"
                                >
                                    <div
                                        class="h-full bg-primary w-1/3 animate-pulse"
                                    ></div>
                                </div>
                                <p
                                    class="text-sm text-center text-muted-foreground"
                                >
                                    Fetching and processing content...
                                </p>
                            </div>
                        </Card.Content>
                    </Card.Root>
                {/if}

                {#if scrapeWebsite.result?.final}
                    <Card.Root>
                        <Card.Header>
                            <Card.Title
                                >{scrapeWebsite.result.final.title}</Card.Title
                            >
                            <Card.Description>
                                Scraped at: {new Date(
                                    scrapeWebsite.result.final.fetchedAt,
                                ).toLocaleString()}
                            </Card.Description>
                        </Card.Header>
                        <Card.Content>
                            <Tabs.Tabs value="markdown">
                                <Tabs.TabsList>
                                    <Tabs.TabsTrigger value="markdown"
                                        >Markdown</Tabs.TabsTrigger
                                    >
                                    <Tabs.TabsTrigger value="html"
                                        >Raw HTML</Tabs.TabsTrigger
                                    >
                                    <Tabs.TabsTrigger value="links"
                                        >Links</Tabs.TabsTrigger
                                    >
                                </Tabs.TabsList>
                                <Tabs.TabsContent value="markdown">
                                    <div
                                        class="markdown-content bg-muted rounded-md p-4 overflow-x-auto max-h-96 mt-4"
                                    >
                                        <pre
                                            class="whitespace-pre-wrap break-words text-sm font-mono leading-relaxed">{scrapeWebsite
                                                .result.final.markdown}</pre>
                                    </div>
                                </Tabs.TabsContent>
                                <Tabs.TabsContent value="html">
                                    <div
                                        class="markdown-content bg-muted rounded-md p-4 overflow-x-auto max-h-96 mt-4"
                                    >
                                        <pre
                                            class="whitespace-pre-wrap break-words text-sm font-mono leading-relaxed">{scrapeWebsite
                                                .result.final.rawHtml ||
                                                "Raw HTML not available."}</pre>
                                    </div>
                                </Tabs.TabsContent>
                                <Tabs.TabsContent value="links">
                                    <div
                                        class="rounded-md border overflow-hidden mt-4"
                                    >
                                        <Table.Root>
                                            <Table.Header>
                                                <Table.Row>
                                                    <Table.Head class="w-1/3"
                                                        >Link Text</Table.Head
                                                    >
                                                    <Table.Head>URL</Table.Head>
                                                </Table.Row>
                                            </Table.Header>
                                            <Table.Body>
                                                {#each scrapeWebsite.result.final.links.slice(0, 50) as link, i (link.href + i)}
                                                    <Table.Row>
                                                        <Table.Cell
                                                            class="font-medium"
                                                        >
                                                            {link.text ||
                                                                "(no text)"}
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
                                        <p
                                            class="text-sm text-muted-foreground mt-3"
                                        >
                                            Showing first 50 of {scrapeWebsite
                                                .result.final.links.length} links
                                        </p>
                                    {/if}
                                </Tabs.TabsContent>
                            </Tabs.Tabs>

                            {#if scrapeWebsite.result.final.warnings?.length}
                                <div
                                    class="border-l-4 border-yellow-400 bg-yellow-50 p-4 mt-6"
                                >
                                    <div class="flex">
                                        <div class="flex-shrink-0">
                                            <Scale
                                                class="h-5 w-5 text-yellow-400"
                                            />
                                        </div>
                                        <div class="ml-3">
                                            <p class="text-sm text-yellow-700">
                                                Warnings detected during
                                                scraping:
                                            </p>
                                            <ul
                                                class="list-disc list-inside text-sm text-yellow-700 mt-2"
                                            >
                                                {#each scrapeWebsite.result.final.warnings as warning, i (i)}
                                                    <li>{warning}</li>
                                                {/each}
                                            </ul>
                                        </div>
                                    </div>
                                </div>
                            {/if}
                        </Card.Content>
                    </Card.Root>
                {/if}
            </div>
        </div>
    {:else}
        <div class="max-w-xl mx-auto">
            <Card.Root>
                <Card.Header>
                    <Card.Title>Scrape Configuration</Card.Title>
                    <Card.Description>
                        Enter the URL and crawl depth to extract content.
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
                                <p class="text-sm text-destructive">
                                    {issue.message}
                                </p>
                            {/each}
                        </div>

                        <div class="space-y-2">
                            <Label for="scrape-depth">Crawl Depth</Label>
                            <Input
                                id="scrape-depth"
                                {...scrapeWebsite.fields.depth.as("text")}
                                class="w-24"
                            />
                            {#each scrapeWebsite.fields.depth.issues() as issue, i (i)}
                                <p class="text-sm text-destructive">
                                    {issue.message}
                                </p>
                            {/each}
                            <p class="text-sm text-muted-foreground">
                                How many levels deep to crawl (1-3)
                            </p>
                        </div>

                        <Button
                            type="submit"
                            disabled={!!scrapeWebsite.pending}
                            class="w-full"
                        >
                            {scrapeWebsite.pending
                                ? "Scraping..."
                                : "Start Scraping"}
                        </Button>
                    </form>
                </Card.Content>
            </Card.Root>

            {#if scrapeWebsite.pending}
                <div class="mt-8">
                    <Card.Root>
                        <Card.Content class="pt-6">
                            <div class="space-y-2">
                                <div
                                    class="h-2 bg-primary/20 rounded-full overflow-hidden"
                                >
                                    <div
                                        class="h-full bg-primary w-1/3 animate-pulse"
                                    ></div>
                                </div>
                                <p
                                    class="text-sm text-center text-muted-foreground"
                                >
                                    Fetching and processing content...
                                </p>
                            </div>
                        </Card.Content>
                    </Card.Root>
                </div>
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
