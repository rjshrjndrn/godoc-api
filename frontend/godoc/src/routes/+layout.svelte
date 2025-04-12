<script lang="ts">
	import '../app.css';

	let { children } = $props();
	const domainName = 'MedClinic';
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
