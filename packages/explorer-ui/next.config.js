const { codecovWebpackPlugin } = require('@codecov/webpack-plugin')

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  // open source project, will not affect performance (one extra comment ot parse)
  productionBrowserSourceMaps: true,
  webpack: (config, { isServer }) => {
    if (!isServer) {
      config.output.publicPath = '/assets'
    }
    config.resolve.extensions = ['.js', '.jsx', '.ts', '.tsx', '.json']
    config.stats = {
      warnings: true,
      errors: true,
      errorDetails: true,
      modules: true,
    }
    config.module.rules.push({
      test: /\.tsx?$/,
      use: 'ts-loader',
    })
    config.plugins.push(
      codecovWebpackPlugin({
        enableBundleAnalysis: process.env.CODECOV_TOKEN !== undefined,
        bundleName: 'explorer-ui',
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
}

module.exports = {
  ...nextConfig,
}
