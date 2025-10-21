<script lang="ts">
    import { Check, X, RefreshCw } from "lucide-svelte";
    import { invalidateAll } from "$app/navigation";
    import * as Alert from "$lib/components/ui/alert/index.js";
    import CheckCircle2Icon from "@lucide/svelte/icons/check-circle-2";
    import AlertCircleIcon from "@lucide/svelte/icons/alert-circle";
    import { Badge } from "$lib/components/ui/badge/index.js";
    import type { PageProps } from "./$types";

    let { data }: PageProps = $props();

    let refreshing = $state(false);
    let waking = $state(false);

    async function refresh() {
        refreshing = true;
        await invalidateAll();
        refreshing = false;
    }

    async function checkHealth() {
        const baseUrl = import.meta.env.VITE_API_URL || "http://localhost:8080";
        try {
            const response = await fetch(`${baseUrl}/health`);
            return response.ok;
        } catch {
            return false;
        }
    }

    async function wakeBackend() {
        waking = true;
        const baseUrl = import.meta.env.VITE_API_URL || "http://localhost:8080";
        // Open backend URL in new tab to wake it up
        window.open(baseUrl, "_blank");

        // Poll health every 5 seconds
        const poll = async () => {
            if (await checkHealth()) {
                waking = false;
                await refresh(); // Refresh the page to show healthy status
                return;
            }
            setTimeout(poll, 5000);
        };
        poll();
    }
</script>

<div class="container mx-auto p-8">
    <div class="flex justify-between items-center mb-6">
        <h2 class="text-3xl font-bold">Health Check</h2>

        <div class="flex items-center gap-3">
            <!-- Environment badge using shadcn Badge component -->
            {#if data?.env}
                <Badge
                    variant={data.env === "development"
                        ? "secondary"
                        : "outline"}
                >
                    Mode: {data.env}
                </Badge>
            {/if}

            <button
                onclick={refresh}
                disabled={refreshing}
                class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2"
            >
                <RefreshCw
                    class="w-4 h-4 mr-2 {refreshing ? 'animate-spin' : ''}"
                />
                Refresh
            </button>
        </div>
    </div>

    <!-- Show wake instructions (if present) -->
    {#if data?.wakeInstructions}
        <div
            class="mb-6 rounded-md bg-amber-50 border border-amber-200 p-4 text-amber-900"
        >
            <strong class="block font-semibold mb-1">Wake Instructions</strong>
            <p class="text-sm">{data.wakeInstructions}</p>
        </div>
    {/if}

    {#if data?.health}
        <div class="bg-green-50 border border-green-200 rounded-lg p-6">
            <div class="flex items-start gap-4">
                <Check class="w-6 h-6 text-green-600 flex-shrink-0 mt-1" />
                <div>
                    <h2 class="text-xl font-semibold text-green-900 mb-2">
                        Backend Service is Healthy
                    </h2>
                    <dl
                        class="grid grid-cols-1 sm:grid-cols-2 gap-x-4 gap-y-2 text-green-800"
                    >
                        <div>
                            <dt class="inline font-medium">Status:</dt>
                            <dd
                                class="inline ml-2 bg-green-100 text-green-900 px-2 py-0.5 rounded-full text-xs"
                            >
                                {data.health.status}
                            </dd>
                        </div>
                        <div>
                            <dt class="inline font-medium">Service:</dt>
                            <dd class="inline ml-2">{data.health.service}</dd>
                        </div>
                        <div>
                            <dt class="inline font-medium">API Response:</dt>
                            <dd class="inline ml-2">{data.duration}ms</dd>
                        </div>
                        <div>
                            <dt class="inline font-medium">Timestamp:</dt>
                            <dd class="inline ml-2">
                                {new Date(
                                    data.health.timestamp * 1000,
                                ).toLocaleString()}
                            </dd>
                        </div>
                    </dl>
                </div>
            </div>
        </div>

        <Alert.Root class="mt-6">
            <CheckCircle2Icon class="h-4 w-4" />
            <Alert.Title>Backend Connection Successful</Alert.Title>
            <Alert.Description>
                Your web scraper backend is running and responding correctly.
                All scraping functionality should work as expected.
            </Alert.Description>
        </Alert.Root>
    {:else}
        <div class="bg-red-50 border border-red-200 rounded-lg p-6">
            <div class="flex items-start gap-4">
                <X class="w-6 h-6 text-red-600 flex-shrink-0 mt-1" />
                <div>
                    <h2 class="text-xl font-semibold text-red-900 mb-2">
                        Backend Service Unavailable
                    </h2>
                    <p class="text-red-800">{data?.error}</p>
                    <p class="text-sm text-red-700 mt-2">
                        API Response Time: {data?.duration}ms
                    </p>

                    <p class="text-red-700 text-sm mt-2">
                        The backend server may be sleeping. Click the button
                        below to wake it up.
                    </p>

                    <button
                        onclick={wakeBackend}
                        disabled={waking}
                        class="mt-4 inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2"
                    >
                        {waking ? "Waking up..." : "Wake Backend Server"}
                    </button>
                </div>
            </div>
        </div>

        <Alert.Root variant="destructive" class="mt-6">
            <AlertCircleIcon class="h-4 w-4" />
            <Alert.Title>Backend Server Not Running</Alert.Title>
            <Alert.Description>
                <p class="mb-2">
                    The backend server is not currently running. If you're
                    running this locally, start the Go backend server:
                </p>
                <div class="bg-muted p-3 rounded-md font-mono text-sm">
                    cd backend && go run main.go
                </div>
                <p class="mt-2">
                    This will start the API server on <code
                        class="bg-muted px-1 py-0.5 rounded text-xs"
                        >http://localhost:8080</code
                    >.
                </p>
            </Alert.Description>
        </Alert.Root>
    {/if}

    <div class="mt-8">
        <a href="/" class="text-primary hover:underline">← Back to Scraper</a>
    </div>
</div>
