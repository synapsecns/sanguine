const TsconfigPathsPlugin = require('tsconfig-paths-webpack-plugin');

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  eslint: {
    // TODO: enable
    ignoreDuringBuilds: true,
  },
  webpack: (config, { isServer }) => {
    config.resolve.plugins = [
      new TsconfigPathsPlugin({ configFileName: './tsconfig.json' }),
    ]
    return config
  },
}

module.exports = {
  ...nextConfig,
}
