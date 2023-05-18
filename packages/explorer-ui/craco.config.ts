// craco.config.js
// @ts-expect-error TS(1208): 'craco.config.ts' cannot be compiled under '--isol... Remove this comment to see the full error message
const CracoAlias = require('craco-alias')

module.exports = {
  webpack: {
    configure: {
      resolve: {
        fallback: {},
      },
      plugins: [],
      ignoreWarnings: [/Failed to parse source map/],
    },
  },
  plugins: [
    {
      plugin: CracoAlias,
      options: {
        source: 'jsconfig',
        baseUrl: './src',
      },
    },
  ],
}
