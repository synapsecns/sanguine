/* eslint-disable prefer-arrow/prefer-arrow-functions */
import * as path from 'path'

import { themes as prismThemes } from 'prism-react-renderer'
import type { Config } from '@docusaurus/types'
import type * as Preset from '@docusaurus/preset-classic'
import { codecovWebpackPlugin } from '@codecov/webpack-plugin'
import searchPlugin from '@easyops-cn/docusaurus-search-local'

const options = {
  id: 'api', // plugin id
  docsPluginId: 'classic', // id of plugin-content-docs or preset for rendering docs
  config: {
    rfqapi: {
      // the <id> referenced when running CLI commands
      specPath: '../../services/rfq/api/docs/swagger.yaml', // path to OpenAPI spec, URLs supported
      baseUrl: 'https://rfq-api.omnirpc.io/',
      outputDir: 'docs/rfq/API', // output directory for generated files
      sidebarOptions: {
        // optional, instructs plugin to generate sidebar.js
        groupPathsBy: 'tag', // group sidebar items by operation "tag"
      },
    },
  },
}

const config: Config = {
  title: 'Synapse Docs',
  tagline: 'The future is cross-chain.',
  favicon: 'brand-assets/synapse-mark.svg',

  // Set the production url of your site here
  url: 'https://docs.bridge.synapseprotocol.com',
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: '/',

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'synapsecns', // Usually your GitHub org/user name.
  projectName: 'sanguine', // Usually your repo name.

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'throw',

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.ts',
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/synapsecns/sanguine/edit/master/docs/bridge/blog-posts/',
          docRootComponent: '@theme/DocRoot',
          docItemComponent: '@theme/ApiItem', // derived from docusaurus-theme-openapi-docs
          // docItemComponent: '@theme/ApiItem', // derived from docusaurus-theme-openapi-docs
          routeBasePath: 'docs',
          path: 'docs',
        },
        blog: {
          showReadingTime: true,
          editUrl:
            'https://github.com/synapsecns/sanguine/tree/master/docs/bridge/',
          blogSidebarTitle: 'All posts',
          blogSidebarCount: 'ALL',
          path: 'blog-posts',
          routeBasePath: 'blog',
          // authorsMapPath: 'blog-posts/authors.yml',
        },
        theme: {
          customCss: './src/css/custom.css',
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    colorMode: {
      respectPrefersColorScheme: true,
    },
    // Replace with your project's social card
    image: 'img/docusaurus-social-card.jpg',
    // announcementBar: {
    // id: 'announcementBar-v3.2', // Increment on change
    // content: `⭐️ If you like Docusaurus, give it a star on <a target="_blank" rel="noopener noreferrer" href="https://github.com/facebook/docusaurus">GitHub</a> and follow us on <a target="_blank" rel="noopener noreferrer" href="https://twitter.com/docusaurus">Twitter ${TwitterSvg}</a>`,
    // },
    navbar: {
      title: 'Synapse Docs',
      logo: {
        alt: 'Synapse logo mark',
        src: 'brand-assets/synapse-mark.svg',
      },
      items: [
        {
          type: 'doc',
          docId: 'About/index',
          position: 'left',
          label: 'Docs',
        },
        { to: '/blog', label: 'Blog', position: 'left' },
        {
          href: 'https://github.com/synapsecns/sanguine',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      // style: 'dark',
      links: [
        {
          title: 'Community',
          items: [
            {
              label: 'Telegram',
              href: 'https://t.me/@synapseprotocol',
            },
            {
              label: 'Discord',
              href: 'https://discord.com/invite/synapseprotocol',
            },
            {
              label: 'Twitter',
              href: 'https://twitter.com/SynapseProtocol',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'GitHub',
              href: 'https://github.com/synapsecns/sanguine',
            },
          ],
        },
      ],
      // copyright: `Copyright © ${new Date().getFullYear()} My Project, Inc. Built with Docusaurus.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
  } satisfies Preset.ThemeConfig,
  themes: ['docusaurus-theme-openapi-docs'], // export theme components
  plugins: [
    ['docusaurus-plugin-openapi-docs', options],
    [
      searchPlugin,
      {
        hashed: true,
        language: ['en'],
      },
    ],
    // please see: https://github.com/facebook/docusaurus/issues/8091#issuecomment-1269112001 for an explanation.
    function () {
      return {
        name: 'resolve-react',
        configureWebpack() {
          return {
            resolve: {
              alias: {
                // assuming root node_modules is up from "./packages/<your-docusaurus>
                react: path.resolve('../../node_modules/react'),
              },
            },
          }
        },
      }
    },
    function () {
      return {
        name: 'bundle-analyzer',
        configureWebpack() {
          return {
            plugins: [
              codecovWebpackPlugin({
                enableBundleAnalysis: process.env.CODECOV_TOKEN !== undefined,
                bundleName: 'docs-bridge',
                uploadToken: process.env.CODECOV_TOKEN,
              }),
            ],
          }
        },
      }
    },
  ],
}

export default config satisfies Config
