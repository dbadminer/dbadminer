<script lang="ts">
	import { site } from '#lib';
</script>

<svelte:head>
	<title>Install · {site.name}</title>
</svelte:head>

<h1>Install</h1>
<p>
	GitHub Releases publish <code>dbadminer</code> for Linux, macOS, and Windows on amd64 and arm64.
</p>

<ol>
	<li>
		Download the archive for your system from
		<a href={site.releases} rel="noopener noreferrer" target="_blank">Releases</a>.
	</li>
	<li>Unpack it and put the binary on your <code>PATH</code>.</li>
	<li>Point it at a database and open the UI.</li>
</ol>

<pre><code>dbadminer --dsn 'postgres://user:pass@127.0.0.1:5432/app?sslmode=disable'</code></pre>

<p>
	Or leave <code>--dsn</code> empty and export <code>DATABASE_URL</code> (or another supported
	variable). Then open <a href="http://127.0.0.1:8080">http://127.0.0.1:8080</a>. The binary listens on
	<code>127.0.0.1:8080</code>.
</p>

<pre><code>dbadminer --version</code></pre>

<h2>From a checkout</h2>
<pre><code>go run ./cmd/dbadminer</code></pre>

<h2>Next</h2>
<ul>
	<li><a href="/docs/cli/">CLI flags and DSN resolution</a></li>
	<li><a href="/docs/docker/">Docker image</a></li>
	<li><a href="/docs/configuration/">Access mode and limits</a></li>
</ul>
