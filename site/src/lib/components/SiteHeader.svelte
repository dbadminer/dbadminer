<script lang="ts">
	import { page } from '$app/state';
	import { site } from '#lib';

	const links = [
		{ href: '/docs/', label: 'Docs' },
		{ href: '/docs/install/', label: 'Install' },
		{ href: '/#writing', label: 'Writing' },
		{ href: site.github, label: 'GitHub', external: true }
	] as const;

	function isActive(href: string) {
		if (href === '/docs/') return page.url.pathname.startsWith('/docs');
		return page.url.pathname === href || page.url.pathname === href.slice(0, -1);
	}
</script>

<header class="site-header">
	<a class="brand" href="/">
		<img src="/logo.svg" alt="" width="36" height="36" />
		<span>{site.name}</span>
	</a>
	<nav aria-label="Primary">
		{#each links as link (link.href)}
			{#if 'external' in link && link.external}
				<a href={link.href} rel="noopener noreferrer" target="_blank">{link.label}</a>
			{:else}
				<a href={link.href} aria-current={isActive(link.href) ? 'page' : undefined}>{link.label}</a>
			{/if}
		{/each}
	</nav>
</header>

<style>
	.site-header {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 1rem;
		padding: 1.1rem 1.5rem;
		max-width: 72rem;
		margin: 0 auto;
		animation: fade 0.5s ease both;
	}

	.brand {
		display: inline-flex;
		align-items: center;
		gap: 0.7rem;
		font-size: 1.15rem;
		font-weight: 650;
		letter-spacing: -0.03em;
		text-decoration: none;
	}

	.brand img {
		border-radius: 0.45rem;
		box-shadow: 0 1px 0 rgb(24 24 31 / 8%);
	}

	nav {
		display: flex;
		flex-wrap: wrap;
		gap: 0.35rem 1.1rem;
		font-size: 0.92rem;
	}

	nav a {
		text-decoration: none;
		color: var(--color-muted);
		transition: color 160ms ease;
	}

	nav a:hover,
	nav a[aria-current='page'] {
		color: var(--color-ink);
	}

	nav a[aria-current='page'] {
		box-shadow: inset 0 -2px 0 var(--color-teal-glow);
	}
</style>
