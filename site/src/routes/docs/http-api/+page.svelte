<script lang="ts">
	import { site } from '#lib';
</script>

<svelte:head>
	<title>HTTP API · {site.name}</title>
</svelte:head>

<h1>HTTP API</h1>
<p>
	With one database the unprefixed routes are the whole API. With several, unprefixed routes follow
	the first name in <code>DBADMINER_DATABASES</code>. The same paths exist under
	<code>/api/databases/&#123;name&#125;/…</code>.
</p>

<h2>Read routes</h2>
<ul>
	<li><code>GET /api/databases</code></li>
	<li><code>GET /api/databases/&#123;name&#125;/info</code></li>
	<li><code>GET /api/databases/&#123;name&#125;/tables</code></li>
	<li><code>GET /api/databases/&#123;name&#125;/tables/&#123;schema&#125;/&#123;table&#125;</code></li>
	<li>
		<code>GET /api/databases/&#123;name&#125;/tables/&#123;schema&#125;/&#123;table&#125;/schema</code>
		(columns, indexes, foreign keys, DDL)
	</li>
	<li>
		<code
			>GET /api/databases/&#123;name&#125;/tables/&#123;schema&#125;/&#123;table&#125;/rows?limit&amp;offset&amp;order&amp;dir&amp;count=1</code
		>
	</li>
	<li>
		<code>POST /api/databases/&#123;name&#125;/query</code> with
		<code>&#123;"sql":"select 1"&#125;</code> and <code>Content-Type: application/json</code>
	</li>
	<li><code>GET /api/info</code></li>
	<li><code>GET /api/tables</code></li>
	<li><code>GET /api/tables/&#123;schema&#125;/&#123;table&#125;</code></li>
	<li><code>GET /api/tables/&#123;schema&#125;/&#123;table&#125;/schema</code></li>
	<li><code>GET /api/tables/&#123;schema&#125;/&#123;table&#125;/rows?…</code></li>
	<li><code>POST /api/query</code></li>
</ul>

<p>
	<code>count=1</code> opts into <code>COUNT(*)</code>. The page size is capped by
	<code>MaxRows</code>.
</p>

<h2>Write routes</h2>
<p>
	Write routes require full access (<code>DBADMINER_ACCESS=full</code>). They return 403 while
	read-only.
</p>
<ul>
	<li><code>POST /api/tables</code> create table</li>
	<li><code>POST /api/views</code> create view</li>
	<li><code>PATCH /api/tables/&#123;schema&#125;/&#123;table&#125;</code> alter columns</li>
	<li><code>DELETE /api/tables/&#123;schema&#125;/&#123;table&#125;?type=table|view</code></li>
	<li><code>POST|DELETE …/indexes[…]</code></li>
	<li><code>POST|DELETE …/foreign-keys[…]</code></li>
	<li><code>POST|PATCH|DELETE …/rows</code> insert / update / delete by primary key</li>
</ul>

<p>
	<code>GET /api/databases</code> lists <code>name</code>, <code>dialect</code>,
	<code>schema</code>, <code>access</code>, and <code>readOnly</code>. It does not include the host,
	user, or password. Row edit APIs require a primary key; tables without one stay SQL-console only.
</p>
