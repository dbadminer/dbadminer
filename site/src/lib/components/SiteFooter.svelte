<script lang="ts">
	import { posts, site } from '#lib';
</script>

<footer class="site-footer">
	<p>
		<a href="/">{site.name}</a>
		·
		<a href={site.github} rel="noopener noreferrer" target="_blank">GitHub</a>
		·
		<a href={site.docker} rel="noopener noreferrer" target="_blank">Docker Hub</a>
		·
		<a href={site.ecohub} rel="noopener noreferrer" target="_blank">MariaDB Ecohub</a>
		·
		<a href={posts[0].href} rel="noopener noreferrer" target="_blank">Blog</a>
	</p>
	<p class="muted">MIT · Built for local databases, not public networks.</p>
</footer>

<style>
	.site-footer {
		max-width: 72rem;
		margin: 0 auto;
		padding: 2.5rem 1.5rem 3rem;
		border-top: 1px solid rgb(24 24 31 / 8%);
		display: grid;
		gap: 0.45rem;
		font-size: 0.85rem;
	}

	.site-footer a {
		text-decoration: none;
	}

	.site-footer a:hover {
		text-decoration: underline;
	}

	.muted {
		color: var(--color-muted);
		margin: 0;
	}

	p {
		margin: 0;
	}
</style>
