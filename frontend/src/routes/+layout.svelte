<script lang="ts">
  import "../app.css";
  import favicon from "$lib/assets/favicon.svg";
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
    <div class="container mx-auto px-4 py-4 flex justify-between items-center">
      <a
        href="/"
        class="text-2xl font-bold hover:text-primary transition-colors"
      >
        Web Scraper
      </a>
      <ThemeToggle />
    </div>
  </header>

  {@render children?.()}
</div>
