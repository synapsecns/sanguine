/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  resolve: {
    extensions: ['.js', '.jsx', '.ts', '.tsx', '.json'],
  },
}

module.exports = {
  ...nextConfig,
}
