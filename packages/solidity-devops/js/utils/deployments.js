const fs = require('fs')
const path = require('path')

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

const saveDeploymentArtifact = (chainName, contractAlias, artifact) => {
  const deployments = readConfigValue('deployments')
  const deploymentFN = `${deployments}/${chainName}/${contractAlias}.json`
  fs.writeFileSync(
    deploymentFN,
    JSON.stringify(refactorArtifact(artifact), null, 2)
  )
}

const getConfirmedFreshDeployment = (chainName, contractAlias) => {
  const freshDeployments = readConfigValue('freshDeployments')
  const freshDeploymentFN = `${freshDeployments}/${chainName}/${contractAlias}.json`
  // Silent exit if the fresh deployment file does not exist
  if (!fs.existsSync(freshDeploymentFN)) {
    logError(`No fresh deployment file found for ${contractAlias}`)
    return null
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
    return null
  }
  logSuccess(`${contractAlias} is deployed at ${address} on ${chainName}`)
  return artifact
}

// Refactor the artifact in a way that its abi property is printed last in the JSON
const refactorArtifact = (artifact) => {
  if (!artifact.abi) {
    return artifact
  }
  const { abi, ...rest } = artifact
  return { ...rest, abi }
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

const getNewDeploymentReceipts = (chainName, scriptFN) => {
  const chainId = getChainId(chainName)
  const scriptBaseName = path.basename(scriptFN)
  const broadcastFN = `broadcast/${scriptBaseName}/${chainId}/run-latest.json`
  // Silent exit if the broadcast file does not exist
  if (!fs.existsSync(broadcastFN)) {
    return []
  }
  const broadcast = JSON.parse(fs.readFileSync(broadcastFN))
  const transactions = broadcast.transactions || []
  const receipts = broadcast.receipts || []
  // Direct deployment transactions have transactionType = CREATE
  // Save hash and contractAddress
  const directDeployments = transactions
    .filter((tx) => tx.transactionType === 'CREATE')
    .map((tx) => ({
      hash: tx.hash,
      address: tx.contractAddress,
    }))
  // Indirect deployments are in additionalContracts field
  // Save tx.hash and tx.additionalContracts[i].address
  const indirectDeployments = transactions
    .filter((tx) => tx.additionalContracts)
    .map((tx) =>
      tx.additionalContracts.map((c) => ({
        hash: tx.hash,
        address: c.address,
      }))
    )
    .flat()
  // Merge two arrays
  const allDeployments = directDeployments.concat(indirectDeployments)
  // Find the block number for each deployment in the receipts
  // Discard the deployment if the receipt is not found
  return allDeployments
    .map((d) => {
      const receipt = receipts.find((r) => r.transactionHash === d.hash)
      if (!receipt) {
        return null
      }
      return {
        address: d.address,
        blockNumber: receipt.blockNumber.startsWith('0x')
          ? parseInt(receipt.blockNumber, 16)
          : receipt.blockNumber,
        hash: d.hash,
      }
    })
    .filter((d) => d)
}

module.exports = {
  createDeploymentDirs,
  saveDeploymentArtifact,
  getConfirmedFreshDeployment,
  getNewDeployments,
  getNewDeploymentReceipts,
}
