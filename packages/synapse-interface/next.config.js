/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: false, //RESET TO TRUE LATER, doing tests and dont want the double render
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
