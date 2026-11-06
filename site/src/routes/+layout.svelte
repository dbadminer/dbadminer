<script lang="ts">
	import './layout.css';
	import SiteFooter from '#lib/components/SiteFooter.svelte';
	import SiteHeader from '#lib/components/SiteHeader.svelte';
	import { site } from '#lib';
	import type { LayoutProps } from './$types';

	let { children }: LayoutProps = $props();
</script>

<svelte:head>
	<link rel="icon" href="/favicon.svg" />
	<meta name="description" content={site.description} />
	<meta property="og:title" content={site.name} />
	<meta property="og:description" content={site.description} />
	<meta property="og:image" content="{site.url}/screenshot.png" />
	<meta property="og:url" content={site.url} />
	<meta name="twitter:card" content="summary_large_image" />
</svelte:head>

<div class="shell">
	<SiteHeader />
	<main>
		{@render children()}
	</main>
	<SiteFooter />
</div>

<style>
	.shell {
		min-height: 100vh;
		display: flex;
		flex-direction: column;
	}

	main {
		flex: 1;
	}
</style>
