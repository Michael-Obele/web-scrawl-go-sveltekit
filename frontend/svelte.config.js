import adapter from "@sveltejs/adapter-auto";
import adapterNode from "@sveltejs/adapter-node";
import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

// Conditionally use adapter-node for Render deployments
const useNodeAdapter =
  Boolean(process.env.FORCE_ADAPTER_NODE) ||
  process.env.SVELTEKIT_ADAPTER === "node";
const selectedAdapter = useNodeAdapter ? adapterNode() : adapter();

/** @type {import('@sveltejs/kit').Config} */
const config = {
  // Consult https://svelte.dev/docs/kit/integrations
  // for more information about preprocessors
  preprocess: vitePreprocess(),

  kit: {
    // adapter-auto automatically chooses the best adapter for your deployment environment
    // If your environment is not supported, or you settled on a specific environment, switch out the adapter.
    // See https://svelte.dev/docs/kit/adapters for more information about adapters.
    adapter: selectedAdapter,
    experimental: {
      remoteFunctions: true,
    },
  },
  vitePlugin: {
    inspector: {
      toggleKeyCombo: "alt-x",
      showToggleButton: "always",
      toggleButtonPos: "bottom-right",
    },
  },

  compilerOptions: {
    experimental: {
      async: true,
    },
  },
};

export default config;
