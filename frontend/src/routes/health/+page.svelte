<script lang="ts">
    import { Check, X, RefreshCw } from "lucide-svelte";
    import { invalidateAll } from "$app/navigation";

    let { data } = $props();

    let refreshing = $state(false);
    async function refresh() {
        refreshing = true;
        await invalidateAll();
        refreshing = false;
    }
</script>

<div class="container mx-auto p-8">
    <div class="flex justify-between items-center mb-6">
        <h1 class="text-3xl font-bold">Health Check</h1>
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

    {#if data.health}
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
    {:else}
        <div class="bg-red-50 border border-red-200 rounded-lg p-6">
            <div class="flex items-start gap-4">
                <X class="w-6 h-6 text-red-600 flex-shrink-0 mt-1" />
                <div>
                    <h2 class="text-xl font-semibold text-red-900 mb-2">
                        Backend Service Unavailable
                    </h2>
                    <p class="text-red-800">{data.error}</p>
                    <p class="text-sm text-red-700 mt-2">
                        API Response Time: {data.duration}ms
                    </p>
                    <p class="text-red-700 text-sm mt-2">
                        Make sure the backend server is running on port 8080.
                    </p>
                </div>
            </div>
        </div>
    {/if}

    <div class="mt-8">
        <a href="/" class="text-primary hover:underline">← Back to Scraper</a>
    </div>
</div>
