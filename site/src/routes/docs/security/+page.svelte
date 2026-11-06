<script lang="ts">
	import { site } from '#lib';
</script>

<svelte:head>
	<title>Security · {site.name}</title>
</svelte:head>

<h1>Security</h1>
<p>
	DBAdminer is a local window onto a database you already have. It is not a sandbox and it is not safe
	to expose on a public network.
</p>

<h2>What the library does</h2>
<ul>
	<li>
		Listens on a loopback address by default. <code>0.0.0.0</code> and <code>::</code> are accepted so
		a container can receive connections that Docker forwards. A specific public address is refused.
	</li>
	<li>
		Rejects requests whose <code>Host</code> is not loopback. A request <code>Origin</code>, when
		present, must match that host.
	</li>
	<li>Defaults to read-only access; write statements are rejected before they reach the driver.</li>
	<li>
		Accepts schema, table, and order names only when they match a live catalog entry, then quotes
		them. Query values are bound parameters.
	</li>
	<li>Compares basic-auth credentials in constant time and slows repeated failures.</li>
	<li>
		Sets <code>Content-Security-Policy</code>, <code>X-Content-Type-Options</code>, and
		<code>Referrer-Policy</code>.
	</li>
	<li>Caps request bodies, row counts, query time, and encoded response size.</li>
	<li>Renders values in the UI with <code>textContent</code>.</li>
</ul>

<h2>What you should do</h2>
<ul>
	<li>Create a database user with <code>SELECT</code> only when you want read-only browsing.</li>
	<li>Open SQLite with <code>mode=ro</code> (the CLI does this in read-only mode).</li>
	<li>
		Publish the Docker port on host loopback only:
		<code>-p 127.0.0.1:8080:8080</code>.
	</li>
	<li>
		Set <code>--user</code> and <code>--password</code> if the page itself should require a login.
		Database passwords are never used as page credentials.
	</li>
</ul>

<p>
	Report vulnerabilities privately to the repository maintainers. Full notes live in
	<a href={`${site.github}/blob/main/SECURITY.md`} rel="noopener noreferrer" target="_blank"
		>SECURITY.md</a
	>.
</p>
