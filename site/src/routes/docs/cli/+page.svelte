<script lang="ts">
	import { site } from '#lib';
</script>

<svelte:head>
	<title>CLI · {site.name}</title>
</svelte:head>

<h1>CLI</h1>
<p>
	The standalone binary resolves a DSN from flags and environment variables, then serves the explorer
	UI.
</p>

<h2>Resolution order</h2>
<ol>
	<li><code>--dsn</code></li>
	<li><code>DATABASE_URL</code>, otherwise <code>DATABASE_PRIVATE_URL</code></li>
	<li>
		One of <code>POSTGRES_URL</code>, <code>POSTGRESQL_URL</code>, <code>MYSQL_URL</code>,
		<code>MYSQL_DSN</code>, <code>SQLITE_URL</code>
	</li>
	<li>
		libpq <code>PG*</code> when <code>PGDATABASE</code> is set (<code>PGHOST</code>,
		<code>PGPORT</code>, <code>PGUSER</code>, <code>PGPASSWORD</code>, <code>PGSSLMODE</code>)
	</li>
	<li>Docker Postgres <code>POSTGRES_*</code> when <code>POSTGRES_DB</code> is set</li>
	<li>
		<code>MYSQL_*</code> when <code>MYSQL_DATABASE</code> is set
	</li>
	<li>
		Generic <code>DB_*</code> when <code>DB_NAME</code> or <code>DB_DATABASE</code> is set
	</li>
	<li><code>SQLITE_PATH</code>, otherwise <code>SQLITE_DATABASE</code></li>
</ol>

<p>
	SQLite opens with <code>mode=ro</code> while the explorer is read-only. A missing host or port
	becomes <code>localhost</code> and <code>5432</code> or <code>3306</code>.
</p>

<h2>Field overrides</h2>
<p>
	These fields override the chosen Postgres or MySQL URL, including <code>--dsn</code> and
	<code>DATABASE_URL</code>. A flag wins over the matching variable. SQLite ignores host, port, user,
	password, and database name.
</p>

<table>
	<thead>
		<tr>
			<th>Field</th>
			<th>Flag</th>
			<th>Variable</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>Host</td>
			<td><code>--db-host</code></td>
			<td><code>DB_HOST</code></td>
		</tr>
		<tr>
			<td>Port</td>
			<td><code>--db-port</code></td>
			<td><code>DB_PORT</code></td>
		</tr>
		<tr>
			<td>User</td>
			<td><code>--db-user</code></td>
			<td><code>DB_USER</code>, otherwise <code>DB_USERNAME</code></td>
		</tr>
		<tr>
			<td>Password</td>
			<td><code>--db-password</code></td>
			<td><code>DB_PASSWORD</code></td>
		</tr>
		<tr>
			<td>Database</td>
			<td><code>--db-name</code></td>
			<td><code>DB_NAME</code>, otherwise <code>DB_DATABASE</code></td>
		</tr>
		<tr>
			<td>Schema</td>
			<td><code>--schema</code></td>
			<td><code>DB_SCHEMA</code></td>
		</tr>
	</tbody>
</table>

<pre><code>dbadminer \
  --dsn 'postgres://postgres:5432/app?sslmode=disable' \
  --db-user ada \
  --db-password 'p@ss' \
  --schema public</code></pre>

<p>
	<code>--schema</code> limits the table list to that schema. On PostgreSQL it also sets
	<code>search_path</code>. <code>--user</code> and <code>--password</code> are basic auth for the
	page, not the database.
</p>

<p>More detail lives in the <a href={site.github} rel="noopener noreferrer" target="_blank">README</a>.</p>
