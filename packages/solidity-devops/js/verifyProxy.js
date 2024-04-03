#!/usr/bin/env node
const { getSavedDeployment } = require('./utils/deployments.js')
const { loadEnv } = require('./utils/env.js')
const {
  initiateVerifyProxy,
  getRequestStatus,
} = require('./utils/etherscan.js')
const { logError, logSuccess } = require('./utils/logger.js')
const { parseCommandLineArgs } = require('./utils/options.js')

loadEnv()

const { positionalArgs } = parseCommandLineArgs({
  requiredArgsCount: 2,
  usage: 'Usage: "yarn vp <chain-name> <contract-alias>"',
})
const [chainName, contractAlias] = positionalArgs
const deployment = getSavedDeployment(chainName, contractAlias)
if (!deployment) {
  process.exit(0)
}
const { address } = deployment
if (!address) {
  logError(`Missing address in deployment file for ${contractAlias}`)
  process.exit(1)
}

const guid = initiateVerifyProxy(chainName, address)
if (!guid) {
  process.exit(1)
}
}
const response = getRequestStatus(chainName, 'checkproxyverification', guid)
if (!response) {
  process.exit(1)
}
if (response.status === '0') {
  logError(`Verification failed: ${response.result}`)
} else {
  logSuccess(response.result)
}
