# Contributing

Run the checks before opening a pull request:

```bash
make check
```

`make check` formats the tree, runs `go vet`, `golangci-lint`, the race detector, and `govulncheck`. Those tools are pinned in `tools/go.mod`.

Integration tests need a database. `compose.yaml` starts PostgreSQL 16, MySQL 8, and MariaDB 11:

```bash
docker compose up -d
make integration
```

`make integration` sets `DBADMINER_INTEGRATION=1` and points `DATABASE_URL` at the Compose Postgres service. Override `DATABASE_URL` with a `mysql://` URL for MySQL or MariaDB. The integration test creates `dbadminer_*` tables and drops them when it finishes.

Format Go files with `gofmt`. The library module must stay free of database drivers. Drivers belong in `cmd/dbadminer`, which has its own `go.mod`. Helpers that are not part of the public API belong in `internal/`.

## UI

The page in `ui/` is the build output of the Svelte app in `web/`. Go embeds that directory, so a checkout can be built without Node. Change the UI from `web/`, then rebuild and commit `ui/`.

```bash
cd web
npm ci
npm run dev
```

`npm run dev` serves the page with Vite and proxies `api` to `http://127.0.0.1:8080`. Start DBAdminer on that address first. `make ui-check` runs `svelte-check` and the production build. `make ui` only rebuilds `ui/`.

## Marketing site

The marketing and docs site lives in `site/` (SvelteKit, static). Edits belong there; do not hand-edit [`dbadminer/dbadminer.github.io`](https://github.com/dbadminer/dbadminer.github.io).

```bash
cd site
npm ci
npm run dev
```

`npm run build` writes `site/build`. On pushes to `main` that touch `site/**` (and on `v*` tags), the Pages workflow builds that output and force-pushes it to `dbadminer.github.io` as `https://dbadminer.github.io`.

### One-time GitHub setup

1. Create a fine-grained PAT (or deploy key) with **Contents: Read and write** on `dbadminer/dbadminer.github.io` only.
2. Add it as the Actions secret `PAGES_DEPLOY_TOKEN` on `dbadminer/dbadminer`.
3. In `dbadminer.github.io` → Settings → Pages: source **Deploy from a branch**, branch `main`, folder `/` (root).
4. Keep a short README in that repo stating it is CI-managed from `dbadminer/site`.

## Releases

Pushing a `v*` tag runs the Release workflow: GoReleaser publishes GitHub Release binaries, and the Docker job pushes [`dbadminer/dbadminer`](https://hub.docker.com/r/dbadminer/dbadminer) for that version and `latest`. Login uses the `DOCKERHUB_USERNAME` and `DOCKERHUB_TOKEN` Actions secrets. The same tag also redeploys the docs site when the Pages workflow runs.
