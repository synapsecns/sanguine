const fs = require('fs')
const path = require('path')

const { getChainId, hasCode } = require('./chain.js')
const { readConfigValue } = require('./config.js')
const { logSuccess, logInfo, logError, logWarning } = require('./logger.js')
const { areEqualAddresses, createDir } = require('./utils.js')
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

const saveNewDeployment = (chainName, contractAlias, potentialReceipts) => {
  // No-op if artifact has already been saved
  const deploymentFN = getDeploymentFN(chainName, contractAlias)
  if (fs.existsSync(deploymentFN)) {
    logInfo(`Deployment artifact already exists for ${contractAlias}`)
    return
  }
  const artifact = getConfirmedFreshDeployment(chainName, contractAlias)
  if (!artifact) {
    return
  }
  // Find the matching receipt
  const receipt = potentialReceipts.find((r) =>
    areEqualAddresses(r.address, artifact.address)
  )
  if (!receipt) {
    logInfo(`No receipt found for ${contractAlias} at ${artifact.address}`)
    return
  }
  // Add receipt.hash and receipt.blockNumber to the artifact, but don't add receipt.address
  artifact.receipt = {
    hash: receipt.hash,
    blockNumber: receipt.blockNumber,
  }
  // Add ABI from the build artifact
  const buildArtifact = getBuildArtifact(contractAlias)
  if (!buildArtifact || !buildArtifact.abi) {
    logInfo(`No ABI found for ${contractAlias}`)
  } else {
    artifact.abi = buildArtifact.abi
  }
  // Save the artifact
  saveDeploymentArtifact(chainName, contractAlias, artifact)
}

const saveDeploymentArtifact = (chainName, contractAlias, artifact) => {
  const deploymentFN = getDeploymentFN(chainName, contractAlias)
  fs.writeFileSync(deploymentFN, JSON.stringify(artifact, null, 2))
}

/**
 * Extracts the contract name from the contract alias: everything before the first dot.
 * Example: 'LinkedPool' from 'LinkedPool.USDC'
 *
 * @param {string} contractAlias - The contract alias
 * @returns {string} The contract name
 */
const getContractName = (contractAlias) => {
  return contractAlias.split('.')[0]
}

const getDeploymentFN = (chainName, contractAlias) => {
  const deployments = readConfigValue('deployments')
  return `${deployments}/${chainName}/${contractAlias}.json`
}

const getFreshDeploymentFN = (chainName, contractAlias) => {
  const freshDeployments = readConfigValue('freshDeployments')
  return `${freshDeployments}/${chainName}/${contractAlias}.json`
}

const getBuildArtifactFN = (contractAlias) => {
  const contractName = getContractName(contractAlias)
  const forgeArtifacts = readConfigValue('forgeArtifacts')
  return `${forgeArtifacts}/${contractName}.sol/${contractName}.json`
}

const getBuildArtifact = (contractAlias) => {
  const artifactFN = getBuildArtifactFN(contractAlias)
  // Silent exit if the artifact file does not exist
  if (!fs.existsSync(artifactFN)) {
    logError(`No artifact file found for ${contractAlias} at ${artifactFN}`)
    return null
  }
  return JSON.parse(fs.readFileSync(artifactFN))
}

const getConfirmedFreshDeployment = (chainName, contractAlias) => {
  const freshDeploymentFN = getFreshDeploymentFN(chainName, contractAlias)
  // Silent exit if the fresh deployment file does not exist
  if (!fs.existsSync(freshDeploymentFN)) {
    logError(
      `No fresh deployment file found for ${contractAlias} at ${freshDeploymentFN}`
    )
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

const getSavedDeployment = (chainName, contractAlias) => {
  const deploymentFN = getDeploymentFN(chainName, contractAlias)
  // Silent exit if the deployment file does not exist
  if (!fs.existsSync(deploymentFN)) {
    logError(`No deployment file found for ${contractAlias} at ${deploymentFN}`)
    return null
  }
  return JSON.parse(fs.readFileSync(deploymentFN))
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

const getAllDeploymentReceipts = (chainName) => {
  // Get the list of script-related directories in the 'broadcast' directory, excluding files
  const scriptDirs = fs
    .readdirSync('broadcast', { withFileTypes: true })
    .filter((d) => d.isDirectory())
    .map((d) => d.name)
  const chainId = getChainId(chainName)
  // We are interested in all 'broadcast/scriptDir/chainId/*.json' files for every scriptDir in the list
  return scriptDirs.flatMap((scriptDir) => {
    const scriptDirPath = path.join('broadcast', scriptDir, chainId)
    if (!fs.existsSync(scriptDirPath)) {
      return []
    }
    return fs
      .readdirSync(scriptDirPath)
      .filter((f) => f.endsWith('.json'))
      .flatMap((broadcastFN) =>
        extractDeploymentReceipts(path.join(scriptDirPath, broadcastFN))
      )
  })
}

const getNewDeploymentReceipts = (chainName, scriptFN, timestamp) => {
  const chainId = getChainId(chainName)
  const scriptBaseName = path.basename(scriptFN)
  const broadcastDir = path.join('broadcast', scriptBaseName, chainId)
  // Silent exit if the broadcast directory does not exist
  if (!fs.existsSync(broadcastDir)) {
    logError(
      `No broadcast directory found for ${scriptBaseName} at ${broadcastDir}`
    )
    return []
  }
  // Look for "*-latest.json" files created after the given timestamp.
  // These are named after the script entry function, which is usually "run", but could be different.
  // In practice there should be only one file, but we implement a generic logic just in case.
  return fs
    .readdirSync(broadcastDir)
    .filter((fn) => {
      if (!fn.endsWith('-latest.json')) {
        return false
      }
      const stats = fs.statSync(path.join(broadcastDir, fn))
      return stats.mtimeMs > timestamp
    })
    .flatMap((fn) => extractDeploymentReceipts(path.join(broadcastDir, fn)))
}

const extractDeploymentReceipts = (broadcastFN) => {
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
  saveNewDeployment,
  saveDeploymentArtifact,
  getContractName,
  getBuildArtifact,
  getConfirmedFreshDeployment,
  getSavedDeployment,
  getNewDeployments,
  getNewDeploymentReceipts,
  getAllDeploymentReceipts,
}
