<script lang="ts">
	import { Check, X } from 'lucide-svelte';
	
	let { data } = $props();
</script>

<div class="container mx-auto p-8">
	<h1 class="text-3xl font-bold mb-6">Health Check</h1>
	
	{#if data.health}
		<div class="bg-green-50 border border-green-200 rounded-lg p-6 flex items-start gap-4">
			<Check class="w-6 h-6 text-green-600 flex-shrink-0 mt-1" />
			<div>
				<h2 class="text-xl font-semibold text-green-900 mb-2">Backend Service is Healthy</h2>
				<dl class="space-y-2 text-green-800">
					<div>
						<dt class="inline font-medium">Status:</dt>
						<dd class="inline ml-2">{data.health.status}</dd>
					</div>
					<div>
						<dt class="inline font-medium">Service:</dt>
						<dd class="inline ml-2">{data.health.service}</dd>
					</div>
					<div>
						<dt class="inline font-medium">Timestamp:</dt>
						<dd class="inline ml-2">{new Date(data.health.timestamp * 1000).toLocaleString()}</dd>
					</div>
				</dl>
			</div>
		</div>
	{:else}
		<div class="bg-red-50 border border-red-200 rounded-lg p-6 flex items-start gap-4">
			<X class="w-6 h-6 text-red-600 flex-shrink-0 mt-1" />
			<div>
				<h2 class="text-xl font-semibold text-red-900 mb-2">Backend Service Unavailable</h2>
				<p class="text-red-800">{data.error}</p>
				<p class="text-red-700 text-sm mt-2">
					Make sure the backend server is running on port 8080.
				</p>
			</div>
		</div>
	{/if}
	
	<div class="mt-6">
		<a href="/" class="text-blue-600 hover:text-blue-800 underline">← Back to home</a>
	</div>
</div>
