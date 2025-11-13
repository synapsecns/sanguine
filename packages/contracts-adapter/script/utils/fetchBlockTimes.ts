import { createPublicClient, http, defineChain, type PublicClient } from 'viem'
import * as fs from 'fs'
import * as path from 'path'

// ANSI color codes for console output
const colors = {
  reset: '\x1b[0m',
  red: '\x1b[31m',
  green: '\x1b[32m',
  yellow: '\x1b[33m',
  cyan: '\x1b[36m',
  bold: '\x1b[1m',
}

type ChainInfo = {
  blockTime: number
  eid: number
  endpointV2: string
  receiveUln302: string
  sendUln302: string
  synapseBridge: string
}

type ChainsConfig = {
  [chain: string]: ChainInfo
}

// Function to read chain IDs from deployment directories
function loadChainIds(): Record<string, number> {
  const deploymentsPath = path.join(__dirname, '../../deployments')
  const chainIds: Record<string, number> = {}

  if (!fs.existsSync(deploymentsPath)) {
    throw new Error(`Deployments directory not found: ${deploymentsPath}`)
  }

  const directories = fs.readdirSync(deploymentsPath)

  for (const dir of directories) {
    const chainIdPath = path.join(deploymentsPath, dir, '.chainId')

    if (fs.existsSync(chainIdPath)) {
      try {
        const chainId = parseInt(
          fs.readFileSync(chainIdPath, 'utf8').trim(),
          10
        )
        if (!isNaN(chainId)) {
          chainIds[dir] = chainId
        }
      } catch (error) {
        console.warn(
          `${colors.yellow}Warning: Failed to read chain ID for ${dir}${colors.reset}`
        )
      }
    }
  }

  return chainIds
}

// Function to get RPC URL for a chain
function getRpcUrl(chainId: number): string {
  const baseUrl = process.env.RPC_BASE_URL

  if (!baseUrl) {
    throw new Error(
      'RPC_BASE_URL environment variable is not set. Please set it to your RPC provider base URL in the .env file'
    )
  }

  // Special case for Moonbeam (chainId 1284)
  if (chainId === 1284) {
    return 'https://moonbeam.unitedbloc.com'
  }

  return `${baseUrl}/${chainId}`
}

// Function to create a public client for a chain
function createChainClient(chainId: number): PublicClient {
  const rpcUrl = getRpcUrl(chainId)

  const chain = defineChain({
    id: chainId,
    name: `Chain ${chainId}`,
    nativeCurrency: { name: 'ETH', symbol: 'ETH', decimals: 18 },
    rpcUrls: {
      default: { http: [rpcUrl] },
    },
  })

  return createPublicClient({
    chain,
    transport: http(rpcUrl, {
      timeout: 30_000,
      retryCount: 3,
    }),
  })
}

// Fetch average block time for a chain (based on extractChainInfo.ts)
async function fetchAverageBlockTime(
  chainName: string,
  chainId: number
): Promise<number | null> {
  try {
    console.log(`  ${colors.cyan}Fetching block time for ${chainName}...${colors.reset}`)

    const client = createChainClient(chainId)

    // Fetch latest block
    const latestBlock = await client.getBlock({ blockTag: 'latest' })
    const latestBlockNumber = latestBlock.number
    const latestTimestamp = latestBlock.timestamp

    // Fetch block 100000 blocks before
    const earlierBlockNumber = latestBlockNumber - 100000n
    if (earlierBlockNumber <= 0n) {
      console.log(
        `    ${colors.red}Chain ${chainName} doesn't have enough blocks for calculation${colors.reset}`
      )
      return null
    }

    const earlierBlock = await client.getBlock({
      blockNumber: earlierBlockNumber,
    })
    const earlierTimestamp = earlierBlock.timestamp
    if (earlierTimestamp >= latestTimestamp) {
      console.log(
        `    ${colors.red}Chain ${chainName} has incorrect block time${colors.reset}`
      )
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
        `    ${colors.red}Chain ${chainName} doesn't have enough blocks for calculation${colors.reset}`
      )
      return null
    }
    const monthAgoBlock = await client.getBlock({
      blockNumber: monthAgoBlockNumber,
    })
    const monthAgoTimestamp = monthAgoBlock.timestamp
    if (monthAgoTimestamp >= latestTimestamp) {
      console.log(
        `    ${colors.red}Chain ${chainName} has incorrect block time${colors.reset}`
      )
      return null
    }

    // Calculate average block time in milliseconds based on last month
    const timeDiff = latestTimestamp - monthAgoTimestamp
    const blockDiff = latestBlockNumber - monthAgoBlockNumber
    const avgBlockTimeMs = Number((timeDiff * 1000n) / blockDiff)

    // Round to nearest 50ms for consistency
    const roundedBlockTime = Math.round(avgBlockTimeMs / 50) * 50

    console.log(
      `    ${colors.green}✓ Block time: ${roundedBlockTime}ms${colors.reset}`
    )

    return roundedBlockTime
  } catch (error) {
    console.log(
      `    ${colors.red}Error fetching block time for ${chainName}: ${error instanceof Error ? error.message : String(error)}${colors.reset}`
    )
    return null
  }
}

async function main() {
  console.log(`${colors.bold}Fetching average block times...${colors.reset}\n`)

  // Load chain IDs
  const chainIds = loadChainIds()
  console.log(
    `Loaded ${Object.keys(chainIds).length} chain IDs from deployments\n`
  )

  // Load chains config
  const chainsConfigPath = path.join(
    __dirname,
    '../../configs/global/chains.json'
  )

  if (!fs.existsSync(chainsConfigPath)) {
    throw new Error(`Chains config not found: ${chainsConfigPath}`)
  }

  const chainsConfig: ChainsConfig = JSON.parse(
    fs.readFileSync(chainsConfigPath, 'utf8')
  )

  console.log(`Found ${Object.keys(chainsConfig).length} chains in config\n`)

  // Build list of chains to update
  const chainsToUpdate = Object.keys(chainsConfig).filter(
    (chain) => chainIds[chain] !== undefined
  )

  console.log(
    `Updating block times for ${chainsToUpdate.length} chains...\n`
  )

  // Fetch block times
  const results: Record<
    string,
    { oldBlockTime: number; newBlockTime: number | null }
  > = {}

  for (const chainName of chainsToUpdate) {
    const chainId = chainIds[chainName]
    const oldBlockTime = chainsConfig[chainName].blockTime
    const newBlockTime = await fetchAverageBlockTime(chainName, chainId)

    results[chainName] = { oldBlockTime, newBlockTime }

    if (newBlockTime !== null) {
      chainsConfig[chainName].blockTime = newBlockTime
    }
  }

  // Print summary
  console.log(`\n${colors.bold}${colors.cyan}Summary${colors.reset}\n`)

  const updated = Object.entries(results).filter(
    ([_, r]) => r.newBlockTime !== null && r.newBlockTime !== r.oldBlockTime
  )
  const unchanged = Object.entries(results).filter(
    ([_, r]) => r.newBlockTime !== null && r.newBlockTime === r.oldBlockTime
  )
  const failed = Object.entries(results).filter(([_, r]) => r.newBlockTime === null)

  if (updated.length > 0) {
    console.log(`${colors.bold}Updated chains:${colors.reset}`)
    for (const [chainName, { oldBlockTime, newBlockTime }] of updated) {
      console.log(
        `  ${chainName}: ${colors.red}${oldBlockTime}ms${colors.reset} → ${colors.green}${newBlockTime}ms${colors.reset}`
      )
    }
    console.log()
  }

  if (unchanged.length > 0) {
    console.log(`${colors.bold}Unchanged chains:${colors.reset}`)
    for (const [chainName, { newBlockTime }] of unchanged) {
      console.log(`  ${chainName}: ${newBlockTime}ms`)
    }
    console.log()
  }

  if (failed.length > 0) {
    console.log(`${colors.bold}${colors.yellow}Failed chains:${colors.reset}`)
    for (const [chainName, { oldBlockTime }] of failed) {
      console.log(
        `  ${chainName}: ${colors.red}Failed to fetch (kept old value: ${oldBlockTime}ms)${colors.reset}`
      )
    }
    console.log()
  }

  // Save updated chains config
  fs.writeFileSync(
    chainsConfigPath,
    JSON.stringify(chainsConfig, null, 2) + '\n'
  )

  console.log(
    `${colors.green}✓ Saved updated chains config to ${chainsConfigPath}${colors.reset}`
  )

  if (failed.length > 0) {
    console.log(
      `\n${colors.yellow}Warning: ${failed.length} chain(s) failed to update${colors.reset}`
    )
    process.exit(78)
  }
}

// Run the script
main().catch((error) => {
  console.error(`${colors.red}Error: ${error}${colors.reset}`)
  process.exit(1)
})
