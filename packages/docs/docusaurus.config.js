// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github')
const darkCodeTheme = require('prism-react-renderer/themes/dracula')

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Synapse Docs',
  tagline: 'Building for a multi-chain world',
  favicon: 'img/favicon.ico',

  // Set the production url of your site here
  url: 'https://to-do.com',
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: '/',

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'synapsecns', // Usually your GitHub org/user name.
  projectName: 'sanguine', // Usually your repo name.

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },
  themes: ['@docusaurus/theme-live-codeblock'],
  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/',
        },
        blog: false,
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      // Replace with your project's social card
      image: 'img/docusaurus-social-card.jpg',
      navbar: {
        title: 'Synapse SDK Docs',
        logo: {
          alt: 'Syn Logo',
          src: 'img/logo.svg',
        },
        items: [
          {
            type: 'doc',
            docId: 'sdk/index',
            position: 'left',
            label: 'Tutorial',
            items: ['sdk/usage', 'sdk/examples'],
          },
          {
            href: 'https://github.com/synapsecns/sanguine',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Docs',
            items: [
              {
                label: 'Tutorial',
                to: '/',
              },
            ],
          },
          {
            title: 'Community',
            items: [
              {
                label: 'Stack Overflow',
                href: 'https://stackoverflow.com/questions/tagged/synapse-protocol',
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
                label: 'Blog',
                to: 'https://synapse.mirror.xyz/',
              },
              {
                label: 'GitHub',
                href: 'https://github.com/synapsecns',
              },
            ],
          },
        ],
        copyright: `Built with Docusaurus.`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
}

module.exports = config
