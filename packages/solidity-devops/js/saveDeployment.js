#!/usr/bin/env node
const {
  createDeploymentDirs,
  getAllDeploymentReceipts,
  saveNewDeployment,
} = require('./utils/deployments.js')
const { loadEnv } = require('./utils/env.js')
const { parseCommandLineArgs } = require('./utils/options.js')

loadEnv()

const { positionalArgs } = parseCommandLineArgs({
  requiredArgsCount: 2,
  usage: 'Usage: "npx sd <chain-name> <contract-alias>"',
})
const [chainName, contractAlias] = positionalArgs
createDeploymentDirs(chainName)
const allReceipts = getAllDeploymentReceipts(chainName)
saveNewDeployment(chainName, contractAlias, allReceipts)
