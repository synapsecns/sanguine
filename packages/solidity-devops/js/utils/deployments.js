const fs = require('fs')

const { getChainId, hasCode } = require('./chain.js')
const { readConfigValue } = require('./config.js')
const { logSuccess, logError, logWarning } = require('./logger.js')
const { createDir } = require('./utils.js')
const { assertCondition } = require('./utils.js')

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
  const chainId = getChainId(chainName)
  fs.writeFileSync(chainIdFile, chainId)
}

const saveDeployment = (chainName, contractAlias) => {
  const freshDeployments = readConfigValue('freshDeployments')
  const deployments = readConfigValue('deployments')
  const freshDeploymentFN = `${freshDeployments}/${chainName}/${contractAlias}.json`
  // Silent exit if the fresh deployment file does not exist
  if (!fs.existsSync(freshDeploymentFN)) {
    logError(`No fresh deployment file found for ${contractAlias}`)
    return
  }
  const artifact = JSON.parse(fs.readFileSync(freshDeploymentFN))
  assertCondition(
    'address' in artifact,
    `Missing address in fresh deployment file ${freshDeploymentFN}`
  )
  const address = artifact.address
  // Silent exit if the contract is not deployed
  if (!hasCode(chainName, address)) {
    logWarning(`${contractAlias} is NOT deployed at ${address} on ${chainName}`)
    return
  }
  logSuccess(`${contractAlias} is deployed at ${address} on ${chainName}`)
  // Move the file to the deployments directory
  const deploymentFN = `${deployments}/${chainName}/${contractAlias}.json`
  fs.renameSync(freshDeploymentFN, deploymentFN)
}

const getNewDeployments = (chainName, timestamp) => {
  const freshDeployments = readConfigValue('freshDeployments')
  const chainDir = `${freshDeployments}/${chainName}`
  // Looks for files created after the given timestamp
  // Then remove the extension and return the list of file names
  const files = fs.readdirSync(chainDir)
  return files
    .filter((file) => {
      const stats = fs.statSync(`${chainDir}/${file}`)
      return file.endsWith('.json') && stats.mtimeMs > timestamp
    })
    .map((file) => file.slice(0, -5))
}

module.exports = {
  createDeploymentDirs,
  saveDeployment,
  getNewDeployments,
}
