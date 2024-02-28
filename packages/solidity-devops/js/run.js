#!/usr/bin/env node
const fs = require('fs')

const { readChainSpecificOptions, logWallet } = require('./utils/chain.js')
const {
  createDeploymentDirs,
  getNewDeployments,
  saveDeployment,
} = require('./utils/deployments.js')
const { loadEnv } = require('./utils/env.js')
const { forgeScript } = require('./utils/forge.js')
const {
  parseCommandLineArgs,
  isBroadcasted,
  addVerifyOptions,
  addOptions,
} = require('./utils/options.js')
const { assertCondition } = require('./utils/utils.js')
const { readWalletOptions } = require('./utils/wallet.js')

loadEnv()

const { positionalArgs, options } = parseCommandLineArgs({
  requiredArgsCount: 3,
  usage:
    'Usage: "yarn sol:run <path-to-script> <chain-name> <wallet-name> [-- <options>]"',
})
const [scriptFN, chainName, walletName] = positionalArgs
assertCondition(
  fs.existsSync(scriptFN),
  `Script file ${scriptFN} does not exist`
)
// Check if this script is being broadcasted
const isBroadcast = isBroadcasted(options)
createDeploymentDirs(chainName)
logWallet(chainName, walletName)
let forgeOptions = addOptions(
  `-f ${chainName}`,
  readWalletOptions(walletName, isBroadcast)
)
forgeOptions = addOptions(forgeOptions, readChainSpecificOptions(chainName))
forgeOptions = addOptions(forgeOptions, options)
if (isBroadcast) {
  forgeOptions = addVerifyOptions(forgeOptions)
}

const currentTimestamp = Date.now()
forgeScript(scriptFN, forgeOptions)
getNewDeployments(chainName, currentTimestamp).forEach((contractAlias) => {
  saveDeployment(chainName, contractAlias)
})
