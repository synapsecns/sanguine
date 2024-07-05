const { codecovRollupPlugin } = require('@codecov/rollup-plugin')

module.exports = {
  rollup: (config, options) => {
    config.plugins.push(
      codecovRollupPlugin({
        enableBundleAnalysis: process.env.CODECOV_TOKEN !== undefined,
        bundleName: 'sdk-router',
        uploadToken: process.env.CODECOV_TOKEN,
        uploadOverrides: {
          sha: process.env.GH_COMMIT_SHA,
        },
      })
    )
    return config
  },
}
