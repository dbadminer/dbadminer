export const site = {
	name: 'DBAdminer',
	tagline: 'A database explorer for Go.',
	description:
		'Start read-only. Download a binary, embed it next to a *sql.DB, or pull the Docker image.',
	url: 'https://dbadminer.github.io',
	github: 'https://github.com/dbadminer/dbadminer',
	docker: 'https://hub.docker.com/r/dbadminer/dbadminer',
	releases: 'https://github.com/dbadminer/dbadminer/releases'
} as const;

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
