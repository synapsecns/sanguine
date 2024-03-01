#!/usr/bin/env node
const fs = require('fs')

const { readChainSpecificOptions, logWallet } = require('./utils/chain.js')
const {
  createDeploymentDirs,
  getConfirmedFreshDeployment,
  getNewDeployments,
  getNewDeploymentReceipts,
  saveDeploymentArtifact,
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
    'Usage: "yarn fsr <path-to-script> <chain-name> <wallet-name> [<options>]"',
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

const newDeployments = getNewDeployments(chainName, currentTimestamp)
const newReceipts = getNewDeploymentReceipts(chainName, scriptFN)
newDeployments.forEach((contractAlias) => {
  const artifact = getConfirmedFreshDeployment(chainName, contractAlias)
  if (!artifact) {
    return
  }
  // Find the matching receipt
  const receipt = newReceipts.find((r) => r.address === artifact.address)
  if (!receipt) {
    logInfo(`No receipt found for ${contractAlias} at ${artifact.address}`)
    return
  }
  // Add receipt.hash and receipt.blockNumber to the artifact, but don't add receipt.address
  artifact.receipt = {
    hash: receipt.hash,
    blockNumber: receipt.blockNumber,
  }
  saveDeploymentArtifact(chainName, contractAlias, artifact)
})
