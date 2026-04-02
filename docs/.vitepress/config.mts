import { defineConfig } from 'vitepress'

export default defineConfig({
	base: '/go/',
	title: 'go',
	description: 'Go standard library extension, adding the missing parts used in the foomo ecosystem to keep dry.',
	themeConfig: {
		logo: '/logo.png',
		outline: [2, 4],
		sidebar: [
			{
				text: 'Overview',
				items: [
					{ text: 'Introduction', link: '/' },
				],
			},
			{
				text: 'Packages',
				items: [
					{ text: 'fmt', link: '/fmt' },
					{ text: 'net', link: '/net' },
					{ text: 'option', link: '/option' },
					{ text: 'os', link: '/os' },
					{ text: 'runtime', link: '/runtime' },
					{ text: 'sec', link: '/sec' },
					{ text: 'slices', link: '/slices' },
					{ text: 'strings', link: '/strings' },
					{ text: 'testing', link: '/testing' },
				],
			},
			{
				text: 'Contributing',
				collapsed: true,
				items: [
					{
						text: "Guideline",
						link: '/CONTRIBUTING.md',
					},
					{
						text: "Code of conduct",
						link: '/CODE_OF_CONDUCT.md',
					},
					{
						text: "Security guidelines",
						link: '/SECURITY.md',
					},
				],
			},
		],
		editLink: {
			pattern: 'https://github.com/foomo/go/edit/main/docs/:path',
			text: 'Suggest changes to this page',
		},
		search: {
			provider: 'local',
		},
		footer: {
			message: 'Made with ♥ <a href="https://www.foomo.org">foomo</a> by <a href="https://www.bestbytes.com">bestbytes</a>',
		},
		socialLinks: [
			{
				icon: 'github',
				link: 'https://github.com/foomo/go',
			},
		],
	},
	head: [
		['meta', { name: 'theme-color', content: '#ffffff' }],
		['link', { rel: 'icon', href: '/logo.png' }],
		['meta', { name: 'author', content: 'foomo by bestbytes' }],
		['meta', { property: 'og:title', content: 'foomo/go' }],
		[
			'meta',
			{
				property: 'og:image',
				content: 'https://github.com/foomo/go/blob/main/docs/public/banner.png?raw=true',
			},
		],
		[
			'meta',
			{
				property: 'og:description',
				content: 'Stop using `go func`, start using `go`',
			},
		],
		['meta', { name: 'twitter:card', content: 'summary_large_image' }],
		[
			'meta',
			{
				name: 'twitter:image',
				content: 'https://github.com/foomo/go/blob/main/docs/public/banner.png?raw=true',
			},
		],
		[
			'meta',
			{
				name: 'viewport',
				content: 'width=device-width, initial-scale=1.0, viewport-fit=cover',
			},
		],
	],
	markdown: {
		theme: {
			dark: 'github-dark',
			light: 'github-light',
		}
	},
	sitemap: {
		hostname: 'https://foomo.github.io/go',
	},
	ignoreDeadLinks: true,
})
