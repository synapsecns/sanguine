import * as fs from 'fs'
import * as path from 'path'
import { type Address, createPublicClient, getAddress, http } from 'viem'
import * as dotenv from 'dotenv'

// Load environment variables
dotenv.config()

// Type definitions for LayerZero metadata
interface EndpointV2 {
  address: string
}

interface Deployment {
  version: number
  eid?: string
  endpointV2?: EndpointV2
  sendUln302?: { address: string }
  receiveUln302?: { address: string }
}

interface ChainMetadata {
  deployments?: Deployment[]
}

interface MetadataResponse {
  [key: string]: ChainMetadata
}

interface ChainInfo {
  blockTime: number
  eid: number
  endpointV2: Address
  receiveUln302: Address
  sendUln302: Address
  synapseBridge: Address
}

interface ChainsConfig {
  [chain: string]: ChainInfo
}

// Type guard to validate metadata structure
function isMetadataResponse(data: unknown): data is MetadataResponse {
  return typeof data === 'object' && data !== null
}

// Fetch SynapseBridge address from GitHub
async function fetchSynapseBridgeAddress(
  chainName: string
): Promise<string | null> {
  try {
    // Chain name mappings for synapse-contracts repo (only non-equal names)
    const synapseNameMap: Record<string, string> = {
      ethereum: 'mainnet',
      bnb: 'bsc',
      kaia: 'klatyn',
    }

    // Use mapped name if exists, otherwise use the chain name as-is
    const synapseChainName = synapseNameMap[chainName] || chainName

    const url = `https://raw.githubusercontent.com/synapsecns/synapse-contracts/master/deployments/${synapseChainName}/SynapseBridge.json`
    const response = await fetch(url)

    if (!response.ok) {
      console.log(
        `Failed to fetch SynapseBridge for ${chainName} (${synapseChainName}): ${response.status}`
      )
      return null
    }

    const data = (await response.json()) as { address: string }
    return data.address || null
  } catch (error) {
    console.log(`Error fetching SynapseBridge for ${chainName}: ${error}`)
    return null
  }
}

// Fetch average block time for a chain
async function fetchAverageBlockTime(
  chainName: string
): Promise<number | null> {
  try {
    // Get RPC URL from environment
    const rpcUrl = process.env[`${chainName.toUpperCase()}_RPC`]
    if (!rpcUrl) {
      console.log(`No RPC URL found for ${chainName}`)
      return null
    }

    // Create public client without chain parameter
    const client = createPublicClient({
      transport: http(rpcUrl),
    })

    // Fetch latest block
    const latestBlock = await client.getBlock({ blockTag: 'latest' })
    const latestBlockNumber = latestBlock.number
    const latestTimestamp = latestBlock.timestamp

    // Fetch block 100000 blocks before
    const earlierBlockNumber = latestBlockNumber - 100000n
    if (earlierBlockNumber <= 0n) {
      console.log(
        `Chain ${chainName} doesn't have enough blocks for calculation`
      )
      return null
    }

    const earlierBlock = await client.getBlock({
      blockNumber: earlierBlockNumber,
    })
    const earlierTimestamp = earlierBlock.timestamp
    if (earlierTimestamp >= latestTimestamp) {
      console.log(`Chain ${chainName} has incorrect block time`)
      return null
    }

    // Approximate block number that was a month ago
    const monthInSeconds = 60n * 60n * 24n * 30n
    const monthAgoBlockNumber =
      latestBlockNumber -
      ((latestBlockNumber - earlierBlockNumber) * monthInSeconds) /
        (latestTimestamp - earlierTimestamp)
    if (monthAgoBlockNumber <= 0n) {
      console.log(
        `Chain ${chainName} doesn't have enough blocks for calculation`
      )
      return null
    }
    const monthAgoBlock = await client.getBlock({
      blockNumber: monthAgoBlockNumber,
    })
    const monthAgoTimestamp = monthAgoBlock.timestamp
    if (monthAgoTimestamp >= latestTimestamp) {
      console.log(`Chain ${chainName} has incorrect block time`)
      return null
    }

    // Calculate average block time in milliseconds based on last month
    const timeDiff = latestTimestamp - monthAgoTimestamp
    const blockDiff = latestBlockNumber - monthAgoBlockNumber
    const avgBlockTimeMs = Number((timeDiff * 1000n) / blockDiff)

    // Round to nearest 50ms for consistency
    return Math.round(avgBlockTimeMs / 50) * 50
  } catch (error) {
    console.log(`Error fetching block time for ${chainName}: ${error}`)
    return null
  }
}

async function extractChainInfo(): Promise<void> {
  const parentDir = path.join(__dirname, '../..')
  const deploymentsDir = path.join(parentDir, 'deployments')
  const outputPath = path.join(parentDir, 'configs', 'global', 'chains.json')

  // Folder name mappings for LayerZero
  const lzNameMap: Record<string, string> = {
    kaia: 'klaytn',
    cronos: 'cronosevm',
    bnb: 'bsc',
  }

  // Fetch metadata
  const response = await fetch(
    'https://metadata.layerzero-api.com/v1/metadata/deployments'
  )
  const metadata = await response.json()

  if (!isMetadataResponse(metadata)) {
    throw new Error('Invalid metadata response format')
  }

  // Process chains
  const allChains = fs
    .readdirSync(deploymentsDir)
    .filter((f) => fs.statSync(path.join(deploymentsDir, f)).isDirectory())

  const chains: ChainsConfig = {}
  const failed: string[] = []

  // Process chains with LayerZero data and SynapseBridge addresses
  for (const chain of allChains) {
    const key = `${lzNameMap[chain] || chain}-mainnet`
    const v2 = metadata[key]?.deployments?.find((d) => d.version === 2)
    const bridgeAddress = await fetchSynapseBridgeAddress(chain)
    const blockTime = await fetchAverageBlockTime(chain)

    if (
      v2?.eid &&
      v2?.endpointV2?.address &&
      v2?.receiveUln302?.address &&
      v2?.sendUln302?.address &&
      bridgeAddress &&
      blockTime !== null
    ) {
      chains[chain] = {
        blockTime: blockTime,
        eid: parseInt(v2.eid),
        endpointV2: getAddress(v2.endpointV2.address),
        receiveUln302: getAddress(v2.receiveUln302.address),
        sendUln302: getAddress(v2.sendUln302.address),
        synapseBridge: getAddress(bridgeAddress),
      }
    } else {
      failed.push(chain)
    }
  }

  // Save result
  fs.mkdirSync(path.dirname(outputPath), { recursive: true })
  fs.writeFileSync(outputPath, JSON.stringify(chains, null, 2))

  console.log(`Saved ${Object.keys(chains).length} chains to ${outputPath}`)
  if (failed.length > 0) {
    console.log(`Failed to extract data for: ${failed.join(', ')}`)
  }
}

extractChainInfo().catch(console.error)
