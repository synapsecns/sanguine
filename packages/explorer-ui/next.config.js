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
