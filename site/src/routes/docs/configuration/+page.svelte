<script lang="ts">
	import { site } from '#lib';
</script>

<svelte:head>
	<title>Configuration · {site.name}</title>
</svelte:head>

<h1>Configuration</h1>

<h2>Access mode</h2>
<p>
	Read-only mode allows one <code>SELECT</code>, <code>WITH</code>, <code>EXPLAIN</code>,
	<code>VALUES</code>, <code>TABLE</code>, or <code>SHOW</code>. Write statements are rejected before
	they reach the database. PostgreSQL, MySQL, and MariaDB also run those statements in a read-only
	transaction. SQLite opens with <code>mode=ro</code>.
</p>
<p>
	Full access runs one statement with the database user's privileges, including DML and DDL.
</p>

<pre><code>dbadminer --access=full
dbadminer --readonly=false
DBADMINER_ACCESS=full</code></pre>

<p>
	<code>--access</code> takes <code>readonly</code> or <code>full</code>. Set only one of
	<code>--access</code> and <code>--readonly</code>. Write UIs stay hidden in read-only mode and their
	APIs return 403.
</p>

<h2>Library and CLI flags</h2>
<p>
	Flags accepted by <code>Flags</code> and <code>Run</code>: <code>--addr</code>,
	<code>--access</code>, <code>--readonly</code>, <code>--user</code>, <code>--password</code>,
	<code>--max-rows</code>, <code>--query-timeout</code>, <code>--max-body</code>,
	<code>--max-response</code>.
</p>
<p>
	<code>cmd/dbadminer</code> also accepts <code>--dsn</code>, <code>--db-host</code>,
	<code>--db-port</code>, <code>--db-user</code>, <code>--db-password</code>,
	<code>--db-name</code>, and <code>--schema</code>. <code>DBADMINER_DATABASES</code> names more than
	one database; each name prefixes those variables.
</p>
<p>
	<code>--user</code> and <code>--password</code> are basic auth for the page. They are never taken
	from the database password. <code>Config.Schema</code> hides every other schema for a pool you
	already opened.
</p>

<h2>Limits</h2>
<p>
	Row page size is capped by <code>MaxRows</code>. Request bodies, query time, and encoded response
	size are also capped. Prefer tighter limits when the explorer is shared across a team VPN.
</p>
