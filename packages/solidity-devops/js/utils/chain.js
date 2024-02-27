const fs = require('fs')

const { getChainId } = require('./cast.js')
const { readConfigValue, tryReadConfigValue } = require('./config.js')
const { readEnv } = require('./env.js')
const { createDir } = require('./utils.js')

/**
 * Reads the URL of the chain's RPC from the environment variables.
 *
 * @param {string} chainName - The name of the chain
 * @returns {string} The URL of the chain's RPC
 */
const readChainRPC = (chainName) => {
  return readEnv(chainName, 'RPC')
}

/**
 * Reads chain specific options from the devops configuration.
 * If no options are found, returns an empty string.
 *
 * @param {string} chainName - The name of the chain
 * @returns {string} The chain specific options
 */
const readChainSpecificOptions = (chainName) => {
  const options = tryReadConfigValue('chains', chainName)
  return options || ''
}

/**
 * Creates and initializes the deployment directories for the chain, if they don't exist.
 *
 * @param {string} chainName - The name of the chain
 */
const createDeploymentDirs = (chainName) => {
  const freshDeployments = readConfigValue('freshDeployments')
  createDir(freshDeployments, chainName)
  const deployments = readConfigValue('deployments')
  createDir(deployments, chainName)
  createChainIdFile(deployments, chainName)
}

const createChainIdFile = (deployments, chainName) => {
  const chainIdFile = `${deployments}/${chainName}/.chainId`
  // Exit if the chain ID file already exists
  if (fs.existsSync(chainIdFile)) {
    return
  }
  const chainRpc = readChainRPC(chainName)
  const chainId = getChainId(chainRpc)
  fs.writeFileSync(chainIdFile, chainId)
}

module.exports = {
  readChainRPC,
  readChainSpecificOptions,
  createDeploymentDirs,
}
