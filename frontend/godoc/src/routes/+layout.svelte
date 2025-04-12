<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';

	let { children } = $props();
	let domainName = $state('MedClinic');
	onMount(() => {
		// Extract domain name from URL when component mounts
		domainName = window.location.hostname.split('.')[0];
		// Make first letter uppercase for better display
		domainName = domainName.charAt(0).toUpperCase() + domainName.slice(1);
	});
</script>

<div class="min-h-screen bg-gray-50">
	<header class="bg-white shadow">
		<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
			<a href="/" class="text-2xl font-bold text-blue-600 hover:text-blue-800 transition-colors">
				{domainName}
			</a>
		</div>
	</header>
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
		{@render children()}
	</div>
</div>
