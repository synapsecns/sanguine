#!/usr/bin/env node
const { getChainId, readChainVerificationOptions } = require('./utils/chain.js')
const {
  getSavedDeployment,
  getContractName,
} = require('./utils/deployments.js')
const { loadEnv } = require('./utils/env.js')
const { forgeVerify } = require('./utils/forge.js')
const { logError } = require('./utils/logger.js')
const { addOptions, parseCommandLineArgs } = require('./utils/options.js')

loadEnv()

const { positionalArgs, options } = parseCommandLineArgs({
  requiredArgsCount: 2,
  usage: 'Usage: "yarn fvc <chain-name> <contract-alias> [<options>]"',
})
const [chainName, contractAlias] = positionalArgs
const contractName = getContractName(contractAlias)
const chainId = getChainId(chainName)
const deployment = getSavedDeployment(chainName, contractAlias)
if (!deployment) {
  process.exit(0)
}
const chainVerificationOptions = readChainVerificationOptions(chainName)
if (!chainVerificationOptions) {
  process.exit(0)
}
const { address, constructorArgs } = deployment
if (!address) {
  logError(`Missing address in deployment file for ${contractAlias}`)
  process.exit(0)
}
if (!constructorArgs) {
  constructorArgs = '0x'
}

let forgeOptions = `${address} ${contractName}`
forgeOptions = addOptions(forgeOptions, `--chain ${chainId} --watch`)
forgeOptions = addOptions(forgeOptions, chainVerificationOptions)
forgeOptions = addOptions(forgeOptions, options)
forgeOptions = addOptions(forgeOptions, `--constructor-args ${constructorArgs}`)

forgeVerify(forgeOptions)
