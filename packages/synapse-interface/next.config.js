/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'export',
  reactStrictMode: true,
  webpack: (config, { isServer }) => {
    config.resolve.extensions = ['.js', '.jsx', '.ts', '.tsx', '.json']
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
