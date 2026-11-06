<script lang="ts">
	import { page } from '$app/state';
	import { docsNav } from '#lib';

	function isActive(href: string) {
		return page.url.pathname === href || page.url.pathname + '/' === href;
	}
</script>

<aside class="sidebar">
	<p class="label">Documentation</p>
	<nav aria-label="Docs">
		{#each docsNav as item (item.href)}
			<a href={item.href} aria-current={isActive(item.href) ? 'page' : undefined}>{item.label}</a>
		{/each}
	</nav>
</aside>

<style>
	.sidebar {
		position: sticky;
		top: 1rem;
		align-self: start;
	}

	.label {
		margin: 0 0 0.75rem;
		font-size: 0.72rem;
		letter-spacing: 0.08em;
		text-transform: uppercase;
		color: var(--color-muted);
	}

	nav {
		display: grid;
		gap: 0.2rem;
	}

	a {
		display: block;
		padding: 0.4rem 0.55rem;
		border-radius: 0.4rem;
		text-decoration: none;
		color: var(--color-muted);
		font-size: 0.92rem;
		transition:
			background 140ms ease,
			color 140ms ease;
	}

	a:hover {
		color: var(--color-ink);
		background: rgb(24 24 31 / 4%);
	}

	a[aria-current='page'] {
		color: var(--color-ink);
		background: rgb(20 184 166 / 14%);
		font-weight: 600;
	}
</style>
