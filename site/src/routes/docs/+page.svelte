<script lang="ts">
	import { site } from '#lib';
</script>

<svelte:head>
	<title>Quick start · {site.name}</title>
</svelte:head>

<h1>Quick start</h1>
<p>
	{site.name} is a database explorer for Go. It starts read-only. Use a binary, embed the library next
	to a <code>*sql.DB</code> you already have, or run the Docker image.
</p>

<h2>Pick a path</h2>
<ul>
	<li><a href="/docs/install/">Install a release binary</a> and open <code>http://127.0.0.1:8080</code></li>
	<li><a href="/docs/library/">Embed the library</a> in an existing Go app</li>
	<li><a href="/docs/docker/">Run the Docker image</a> on loopback</li>
	<li><a href="/docs/cli/">Wire DSN and environment variables</a> for the CLI</li>
</ul>

<h2>Five-minute library embed</h2>
<pre><code>cfg := dbadminer.DefaultConfig()
ex, err := dbadminer.New(db, cfg)
if err != nil &#123;
	log.Fatal(err)
&#125;
if err := ex.ListenAndServe(ctx); err != nil &#123;
	log.Fatal(err)
&#125;</code></pre>
<p>
	Import <code>github.com/dbadminer/dbadminer</code>. That import does not add database drivers to your
	module. <code>DefaultConfig</code> is read-only and listens on <code>127.0.0.1:8080</code>.
</p>

<h2>Safety first</h2>
<p>
	Read-only mode is best effort. Prefer a database user that only has <code>SELECT</code>, and open
	SQLite with <code>mode=ro</code>. See <a href="/docs/security/">Security</a> before exposing the UI
	beyond your machine.
</p>
