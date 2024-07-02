import { codecovRollupPlugin } from '@codecov/rollup-plugin'

module.exports = {
  rollup(config, options) {
    config.plugins.push(
      codecovRollupPlugin({
        enableBundleAnalysis: process.env.CODECOV_TOKEN !== undefined,
        bundleName: 'sdk-router',
        uploadToken: process.env.CODECOV_TOKEN,
      })
    )
    return config
  },
}
