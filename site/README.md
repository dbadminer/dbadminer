# DBAdminer site

Marketing and documentation site for [DBAdminer](https://github.com/dbadminer/dbadminer), published to [dbadminer.github.io](https://dbadminer.github.io).

## Develop

```bash
npm ci
npm run dev
```

## Build

```bash
npm run build
npm run preview
```

Output lands in `build/` and is deployed by [`.github/workflows/pages.yml`](../.github/workflows/pages.yml) to `dbadminer/dbadminer.github.io`. See [CONTRIBUTING.md](../CONTRIBUTING.md) for the one-time GitHub Pages token and branch setup.

Do not commit hand-edited HTML into the Pages repo; change sources here instead.

`static/llms.txt` is published at `/llms.txt` so coding agents can find the curated docs index ([llms.txt](https://llmstxt.org/)).
