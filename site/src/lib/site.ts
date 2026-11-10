export const site = {
	name: 'DBAdminer',
	tagline: 'A database explorer for Go.',
	description:
		'Start read-only. Download a binary, embed it next to a *sql.DB, or pull the Docker image.',
	url: 'https://dbadminer.github.io',
	github: 'https://github.com/dbadminer/dbadminer',
	docker: 'https://hub.docker.com/r/dbadminer/dbadminer',
	releases: 'https://github.com/dbadminer/dbadminer/releases',
	devto: 'https://dev.to/rp01',
	ecohub: 'https://ecohub.mariadb.org/database-management/dbadminer'
} as const;

/** Newest first. Add posts here as you publish. */
export const posts = [
	{
		title: "I built a database explorer that can't ruin your Friday",
		href: 'https://dev.to/rp01/i-built-a-database-explorer-that-cant-ruin-your-friday-hnh',
		date: '2026'
	}
] as const;

export const docsNav = [
	{ href: '/docs/', label: 'Quick start' },
	{ href: '/docs/install/', label: 'Install' },
	{ href: '/docs/library/', label: 'Library' },
	{ href: '/docs/cli/', label: 'CLI' },
	{ href: '/docs/docker/', label: 'Docker' },
	{ href: '/docs/configuration/', label: 'Configuration' },
	{ href: '/docs/security/', label: 'Security' },
	{ href: '/docs/http-api/', label: 'HTTP API' },
	{ href: '/docs/dialects/', label: 'Dialects' }
] as const;
