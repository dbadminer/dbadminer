<script lang="ts">
	import { site } from '#lib';
</script>

<svelte:head>
	<title>Dialects · {site.name}</title>
</svelte:head>

<h1>Dialects</h1>
<p>
	DBAdminer ships with PostgreSQL, MySQL / MariaDB, and SQLite support. Dialect detection runs from the
	driver name and connection metadata. Custom databases can implement the <code>Dialect</code>
	interface.
</p>

<h2>Built-in</h2>
<table>
	<thead>
		<tr>
			<th>Dialect</th>
			<th>Notes</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>PostgreSQL</td>
			<td>
				Schema-aware listing, <code>search_path</code> when <code>--schema</code> is set, read-only
				transactions in read-only mode.
			</td>
		</tr>
		<tr>
			<td>MySQL / MariaDB</td>
			<td>Schema maps to database name; read-only transactions in read-only mode.</td>
		</tr>
		<tr>
			<td>SQLite</td>
			<td>
				File path or <code>sqlite://</code> DSN. Opens with <code>mode=ro</code> while the explorer is
				read-only.
			</td>
		</tr>
	</tbody>
</table>

<h2>Extending</h2>
<p>
	Implement <code>Dialect</code>: <code>Name</code>, <code>Quote</code>, <code>TablesQuery</code>,
	<code>ColumnsQuery</code>, <code>ColumnsArgs</code>, <code>Qualify</code>,
	<code>SupportsReadOnlyTx</code>. Planned additions for richer schema views include index and foreign
	key queries, <code>DDL</code>, dialect-specific pagination, and placeholders for filters.
</p>

<p>
	Tested versions and contribution notes are in the
	<a href={site.github} rel="noopener noreferrer" target="_blank">repository</a>. Features should work
	on all three dialects or be clearly marked as dialect-specific.
</p>
