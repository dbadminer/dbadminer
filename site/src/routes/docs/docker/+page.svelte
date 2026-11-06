<script lang="ts">
	import { site } from '#lib';
</script>

<svelte:head>
	<title>Docker · {site.name}</title>
</svelte:head>

<h1>Docker</h1>
<p>
	Version tags (<code>v*</code>) publish
	<a href={site.docker} rel="noopener noreferrer" target="_blank">dbadminer/dbadminer</a> for that
	version and for <code>latest</code>. Publish the port only on the host loopback.
</p>

<pre><code>docker pull dbadminer/dbadminer:latest
docker run --rm -p 127.0.0.1:8080:8080 \
  -e DATABASE_URL='postgres://postgres:5432/app?sslmode=disable' \
  -e DB_USER=ada \
  -e DB_PASSWORD='p@ss' \
  -e DB_SCHEMA=public \
  dbadminer/dbadminer:latest</code></pre>

<p>Open <a href="http://127.0.0.1:8080">http://127.0.0.1:8080</a>. The image starts read-only.</p>

<pre><code>docker run --rm -p 127.0.0.1:8080:8080 \
  -e DBADMINER_ACCESS=full \
  -e DATABASE_URL='postgres://...' \
  dbadminer/dbadminer:latest</code></pre>

<h2>Multiple databases</h2>
<p>
	Set <code>DBADMINER_DATABASES</code> to a comma-separated list of names. Each name is a prefix:
	<code>SHOP_DATABASE_URL</code> is read as <code>DATABASE_URL</code>, and
	<code>SHOP_DB_PASSWORD</code> is read as <code>DB_PASSWORD</code>.
</p>

<pre><code>services:
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
      APP_DB_SCHEMA: public</code></pre>

<h2>Demo stack</h2>
<pre><code>docker compose -f examples/demo/compose.yaml up --build</code></pre>
<p>
	Open <code>http://127.0.0.1:8090</code> for the demo app and <code>http://127.0.0.1:8080</code> for
	DBAdminer.
</p>
