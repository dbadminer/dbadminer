<script lang="ts">
	import { posts, site } from '#lib';

	const composeSnippet = `services:
  dbadminer:
    image: dbadminer/dbadminer:latest
    ports:
      - "127.0.0.1:8080:8080"
    environment:
      DBADMINER_DATABASES: shop,app
      SHOP_DATABASE_URL: mysql://mysql:3306/shop
      SHOP_DB_USER: ada
      SHOP_DB_PASSWORD: secret
      APP_DATABASE_URL: postgres://postgres:5432/app?sslmode=disable
      APP_DB_USER: ada
      APP_DB_PASSWORD: secret
      APP_DB_SCHEMA: public`;
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

<section class="compose">
	<div class="compose-copy">
		<h2>Drop it into Compose</h2>
		<p>
			Publish only on loopback, point env vars at your services, and open
			<code>http://127.0.0.1:8080</code>. Multi-database names become a switcher in the UI.
		</p>
		<p class="more">
			<a href="/docs/docker/">Full Docker guide</a>
			·
			<a href={`${site.github}/tree/main/examples/demo`} rel="noopener noreferrer" target="_blank"
				>Demo stack</a
			>
		</p>
	</div>
	<pre><code>{composeSnippet}</code></pre>
</section>

<section id="writing" class="writing">
	<h2>Writing</h2>
	<ul>
		{#each posts as post (post.href)}
			<li>
				<a href={post.href} rel="noopener noreferrer" target="_blank">{post.title}</a>
				<span>{post.date}</span>
			</li>
		{/each}
	</ul>
	<p class="more">
		More on
		<a href={site.devto} rel="noopener noreferrer" target="_blank">dev.to/rp01</a>.
	</p>
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
		padding: 0.5rem 0 1.5rem;
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

	.compose,
	.writing {
		max-width: 72rem;
		margin: 0 auto;
		padding: 2rem 1.5rem 2.5rem;
		animation: rise 0.7s ease both;
	}

	.compose {
		display: grid;
		gap: 1.25rem;
	}

	.compose-copy h2,
	.writing h2 {
		margin: 0 0 0.55rem;
		font-size: 1.35rem;
		letter-spacing: -0.03em;
	}

	.compose-copy p,
	.writing p {
		margin: 0;
		max-width: 36rem;
		color: var(--color-muted);
		line-height: 1.55;
	}

	.compose-copy code {
		font-size: 0.92em;
		padding: 0.05em 0.3em;
		border-radius: 0.3rem;
		background: rgb(24 24 31 / 6%);
	}

	.more {
		margin-top: 0.75rem !important;
		font-size: 0.92rem;
	}

	.more a {
		color: var(--color-teal);
		text-decoration: none;
		font-weight: 600;
	}

	.more a:hover {
		text-decoration: underline;
	}

	.compose pre {
		margin: 0;
		overflow-x: auto;
		border-radius: 0.65rem;
		background: #18181f;
		color: #e8eeeb;
		padding: 1.1rem 1.2rem;
		font-size: 0.84rem;
		line-height: 1.5;
		box-shadow: 0 16px 40px rgb(24 24 31 / 14%);
	}

	.compose pre code {
		font-family: var(--font-mono);
	}

	.writing ul {
		list-style: none;
		margin: 0.85rem 0 0;
		padding: 0;
		display: grid;
		gap: 0.55rem;
	}

	.writing li {
		display: flex;
		flex-wrap: wrap;
		align-items: baseline;
		justify-content: space-between;
		gap: 0.35rem 1rem;
		padding: 0.7rem 0;
		border-bottom: 1px solid rgb(24 24 31 / 8%);
	}

	.writing li a {
		text-decoration: none;
		font-weight: 600;
		letter-spacing: -0.02em;
	}

	.writing li a:hover {
		color: var(--color-teal);
	}

	.writing li span {
		color: var(--color-muted);
		font-size: 0.85rem;
	}

	@media (min-width: 900px) {
		.visual {
			padding-inline: 0;
		}

		.compose {
			grid-template-columns: minmax(14rem, 0.9fr) minmax(0, 1.4fr);
			align-items: start;
			gap: 2rem;
		}
	}
</style>
