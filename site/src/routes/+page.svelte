<script lang="ts">
	import { site } from '#lib';
</script>

<svelte:head>
	<title>{site.name} — database explorer for Go</title>
</svelte:head>

<section class="hero">
	<div class="copy">
		<p class="brand">{site.name}</p>
		<h1>Browse PostgreSQL, MySQL, and SQLite without leaving your stack.</h1>
		<p class="lede">
			Starts read-only. Ship a binary, embed next to a <code>*sql.DB</code>, or pull the Docker
			image. The library imports only the Go standard library.
		</p>
		<div class="ctas">
			<a class="primary" href="/docs/install/">Install</a>
			<a class="secondary" href="/docs/">Docs</a>
			<a class="ghost" href={site.github} rel="noopener noreferrer" target="_blank">GitHub</a>
		</div>
	</div>
</section>

<section class="visual" aria-label="Product screenshot">
	<div class="frame">
		<img
			src="/screenshot.png"
			alt="DBAdminer browsing MySQL shop.customers with tables and routines in the explorer"
			width="1600"
			height="1000"
		/>
	</div>
</section>

<style>
	.hero {
		max-width: 72rem;
		margin: 0 auto;
		padding: 2.5rem 1.5rem 1.25rem;
		animation: rise 0.7s ease both;
	}

	.copy {
		max-width: 42rem;
	}

	.brand {
		margin: 0 0 0.85rem;
		font-size: clamp(2.4rem, 6vw, 4.4rem);
		font-weight: 700;
		letter-spacing: -0.05em;
		line-height: 0.95;
		color: var(--color-ink);
	}

	h1 {
		margin: 0;
		font-size: clamp(1.15rem, 2.4vw, 1.55rem);
		font-weight: 500;
		line-height: 1.35;
		letter-spacing: -0.02em;
		color: var(--color-ink-soft);
	}

	.lede {
		margin: 1rem 0 0;
		max-width: 38rem;
		font-size: 1rem;
		line-height: 1.55;
		color: var(--color-muted);
		animation: rise 0.7s ease 0.08s both;
	}

	.lede code {
		font-size: 0.92em;
		padding: 0.05em 0.3em;
		border-radius: 0.3rem;
		background: rgb(24 24 31 / 6%);
	}

	.ctas {
		display: flex;
		flex-wrap: wrap;
		gap: 0.65rem;
		margin-top: 1.6rem;
		animation: rise 0.7s ease 0.14s both;
	}

	.ctas a {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		min-height: 2.55rem;
		padding: 0 1.05rem;
		border-radius: 0.45rem;
		font-size: 0.95rem;
		font-weight: 600;
		text-decoration: none;
		transition:
			transform 160ms ease,
			background 160ms ease,
			box-shadow 160ms ease;
	}

	.ctas a:hover {
		transform: translateY(-1px);
	}

	.primary {
		background: var(--color-ink);
		color: #f7faf8;
		box-shadow: 0 8px 20px rgb(24 24 31 / 18%);
	}

	.primary:hover {
		background: #101016;
	}

	.secondary {
		background: rgb(20 184 166 / 16%);
		color: var(--color-teal);
	}

	.ghost {
		color: var(--color-muted);
		border: 1px solid rgb(24 24 31 / 12%);
		background: rgb(255 255 255 / 45%);
	}

	.visual {
		max-width: 72rem;
		margin: 0 auto;
		padding: 0.5rem 0 3rem;
		animation: rise 0.9s ease 0.18s both;
	}

	.frame {
		position: relative;
		overflow: hidden;
		border-block: 1px solid rgb(24 24 31 / 12%);
		background:
			linear-gradient(90deg, transparent 0, rgb(45 212 191 / 8%) 50%, transparent 100%),
			var(--color-ink);
		box-shadow: 0 24px 60px rgb(24 24 31 / 18%);
	}

	.frame::before {
		content: '';
		position: absolute;
		inset: 0;
		pointer-events: none;
		opacity: 0.35;
		background-image: linear-gradient(rgb(255 255 255 / 4%) 1px, transparent 1px),
			linear-gradient(90deg, rgb(255 255 255 / 4%) 1px, transparent 1px);
		background-size: 40px 40px;
		animation: drift 18s linear infinite;
	}

	.frame img {
		display: block;
		width: 100%;
		height: auto;
		position: relative;
		z-index: 1;
	}

	@media (min-width: 900px) {
		.visual {
			padding-inline: 0;
		}
	}
</style>
