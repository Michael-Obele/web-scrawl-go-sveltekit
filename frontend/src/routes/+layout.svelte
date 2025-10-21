<script lang="ts">
    import "../app.css";
    import favicon from "$lib/assets/web-scrawl-logo.svg";
    import { ModeWatcher } from "mode-watcher";
    import ThemeToggle from "$lib/components/theme-toggle.svelte";

    let { children } = $props();
</script>

<svelte:head>
    <link rel="icon" href={favicon} />
</svelte:head>

<ModeWatcher />

<div class="min-h-screen bg-background">
    <header class="border-b">
        <div
            class="container mx-auto px-4 py-4 flex justify-between items-center"
        >
            <a
                href="/"
                class="text-2xl flex font-bold hover:text-primary transition-colors"
            >
                <img src={favicon} alt="logo" class="w-8 h-8 mr-2" />
                Web Scraper
            </a>
            <ThemeToggle />
        </div>
    </header>

    {@render children?.()}
</div>
