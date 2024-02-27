#!/usr/bin/env node
const fs = require('fs')

const {
  readChainSpecificOptions,
  createDeploymentDirs,
} = require('./utils/chain.js')
const { loadEnv } = require('./utils/env.js')
const { forgeScript } = require('./utils/forge.js')
const { logWallet } = require('./utils/logger.js')
const {
  parseCommandLineArgs,
  isBroadcasted,
  addVerifyOptions,
  addOptions,
} = require('./utils/options.js')
const { assertCondition } = require('./utils/utils.js')
const { readWalletOptions } = require('./utils/wallet.js')

loadEnv()

const { positionalArgs, options } = parseCommandLineArgs()
assertCondition(
  positionalArgs.length === 3,
  'Usage: "yarn sol:run <path-to-script> <chain-name> <wallet-name> [-- <options>]"'
)

const [scriptFN, chainName, walletName] = positionalArgs
assertCondition(
  fs.existsSync(scriptFN),
  `Script file ${scriptFN} does not exist`
)
let forgeOptions = addOptions(`-f ${chainName}`, options)
// Check if this script is broadcasted
const isBroadcast = isBroadcasted(forgeOptions)
if (isBroadcast) {
  forgeOptions = addVerifyOptions(forgeOptions)
}

createDeploymentDirs(chainName)
logWallet(chainName, walletName)
forgeOptions = addOptions(forgeOptions, readWalletOptions(walletName))
forgeOptions = addOptions(forgeOptions, readChainSpecificOptions(chainName))

forgeScript(scriptFN, forgeOptions)
// TODO: handle saved fresh deployments
