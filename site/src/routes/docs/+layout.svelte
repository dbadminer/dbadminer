<script lang="ts">
	import DocsSidebar from '#lib/components/DocsSidebar.svelte';
	import type { LayoutProps } from './$types';

	let { children }: LayoutProps = $props();
</script>

<div class="docs">
	<DocsSidebar />
	<article class="prose prose-neutral max-w-none prose-headings:tracking-tight prose-a:text-teal-700 prose-code:before:content-none prose-code:after:content-none">
		{@render children()}
	</article>
</div>

<style>
	.docs {
		max-width: 72rem;
		margin: 0 auto;
		padding: 1.25rem 1.5rem 3rem;
		display: grid;
		gap: 1.75rem;
		animation: rise 0.55s ease both;
	}

	@media (min-width: 860px) {
		.docs {
			grid-template-columns: 13.5rem minmax(0, 1fr);
			gap: 2.5rem;
			align-items: start;
		}
	}

	article {
		min-width: 0;
		padding: 1.25rem 1.35rem 1.5rem;
		border: 1px solid rgb(24 24 31 / 8%);
		border-radius: 0.75rem;
		background: rgb(255 255 255 / 72%);
		backdrop-filter: blur(8px);
		box-shadow: 0 10px 30px rgb(24 24 31 / 5%);
	}

	article :global(pre) {
		overflow-x: auto;
		border-radius: 0.55rem;
		background: #18181f !important;
		color: #e8eeeb;
		padding: 1rem 1.1rem;
		font-size: 0.86rem;
		line-height: 1.5;
	}

	article :global(code) {
		font-size: 0.9em;
	}

	article :global(:not(pre) > code) {
		padding: 0.1em 0.35em;
		border-radius: 0.3rem;
		background: rgb(24 24 31 / 6%);
	}

	article :global(table) {
		display: block;
		overflow-x: auto;
		width: 100%;
	}
</style>
