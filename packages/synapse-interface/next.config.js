const path = require('path')

const { codecovWebpackPlugin } = require('@codecov/webpack-plugin')

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  // open source project, will not affect performance (one extra comment ot parse)
  productionBrowserSourceMaps: true,
  webpack: (config, { isServer }) => {
    config.module.rules.push({
      test: /\.nvmrc$/,
      use: [
        {
          loader: path.resolve(__dirname, './nvmrc-loader.js'),
        },
      ],
    })
    config.resolve.extensions = [
      '.js',
      '.jsx',
      '.ts',
      '.tsx',
      '.json',
      '.nvmrc',
    ]
    config.plugins.push(
      codecovWebpackPlugin({
        enableBundleAnalysis: process.env.CODECOV_TOKEN !== undefined,
        bundleName: 'synapse-interface',
        uploadToken: process.env.CODECOV_TOKEN,
        uploadOverrides: {
          sha: process.env.GH_COMMIT_SHA,
        },
      })
    )
    return config
  },
  eslint: {
    // TODO: enable
    ignoreDuringBuilds: true,
  },
  typescript: {
    tsconfigPath: './tsconfig.json',
  },
  i18n: {
    locales: ['en-US', 'fr', 'lorem-ipsum', 'jp'],
    defaultLocale: 'en-US',
  },
}

module.exports = {
  ...nextConfig,
}
