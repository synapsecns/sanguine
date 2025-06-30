import {
  createPublicClient,
  http,
  getAddress,
  type Address,
  zeroAddress,
} from 'viem'
import { mainnet } from 'viem/chains'
import * as fs from 'fs'
import * as path from 'path'

// BridgeConfig contract address
const BRIDGE_CONFIG_ADDRESS = '0x5217c83ca75559B1f8a8803824E5b7ac233A12a1'

// BridgeConfig contract ABI (only the methods we need)
const bridgeConfigABI = [
  {
    name: 'getAllTokenIDs',
    type: 'function',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: 'result', type: 'string[]' }],
  },
  {
    name: 'getToken',
    type: 'function',
    stateMutability: 'view',
    inputs: [
      { name: 'tokenID', type: 'string' },
      { name: 'chainId', type: 'uint256' },
    ],
    outputs: [
      {
        name: 'result',
        type: 'tuple',
        components: [
          { name: 'chainId', type: 'uint256' },
          { name: 'tokenAddress', type: 'string' },
          { name: 'tokenDecimals', type: 'uint8' },
          { name: 'maxSwap', type: 'uint256' },
          { name: 'minSwap', type: 'uint256' },
          { name: 'swapFee', type: 'uint256' },
          { name: 'maxSwapFee', type: 'uint256' },
          { name: 'minSwapFee', type: 'uint256' },
          { name: 'hasUnderlying', type: 'bool' },
          { name: 'isUnderlying', type: 'bool' },
        ],
      },
    ],
  },
] as const

type BridgeTokenData = {
  [tokenID: string]: {
    [chain: string]: {
      tokenAddress: Address
      isUnderlying: boolean
    }
  }
}

// Function to sort BridgeTokenData keys
function sortBridgeTokenData(data: BridgeTokenData): BridgeTokenData {
  const sortedData: BridgeTokenData = {}

  // Sort token IDs
  const sortedTokenIds = Object.keys(data).sort()

  for (const tokenId of sortedTokenIds) {
    sortedData[tokenId] = {}

    // Sort chain names for each token
    const sortedChains = Object.keys(data[tokenId]).sort()

    for (const chain of sortedChains) {
      sortedData[tokenId][chain] = data[tokenId][chain]
    }
  }

  return sortedData
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
        console.warn(`Failed to read chain ID for ${dir}:`, error)
      }
    }
  }

  return chainIds
}

async function main() {
  // Load chain IDs dynamically
  const chainIds = loadChainIds()
  console.log(
    `Loaded ${Object.keys(chainIds).length} chain IDs from deployments`
  )

  // Create a public client for mainnet (where BridgeConfig is deployed)
  const client = createPublicClient({
    chain: mainnet,
    transport: http(),
  })

  // Read chains config
  const chainsConfigPath = path.join(
    __dirname,
    '../../configs/global/chains.json'
  )
  const chainsConfig = JSON.parse(fs.readFileSync(chainsConfigPath, 'utf8'))
  const chains = Object.keys(chainsConfig)

  // Filter to only include chains that have chain IDs
  const validChains = chains.filter((chain) => chainIds[chain] !== undefined)
  const allChainIds = validChains.map((chain) => chainIds[chain])

  console.log(`Found ${validChains.length} chains with chain IDs`)

  // Step 1: Get all token IDs
  console.log('Fetching all token IDs...')
  const tokenIDs = await client.readContract({
    address: BRIDGE_CONFIG_ADDRESS,
    abi: bridgeConfigABI,
    functionName: 'getAllTokenIDs',
  })

  console.log(`Found ${tokenIDs.length} tokens`)

  // Step 2: Build multicall for all token/chain combinations
  console.log('Building multicall for token data...')
  const calls: Array<{
    address: Address
    abi: typeof bridgeConfigABI
    functionName: 'getToken'
    args: readonly [string, bigint]
  }> = []

  for (const tokenID of tokenIDs) {
    for (const chainId of allChainIds) {
      calls.push({
        address: BRIDGE_CONFIG_ADDRESS,
        abi: bridgeConfigABI,
        functionName: 'getToken',
        args: [tokenID, BigInt(chainId)] as const,
      })
    }
  }

  // Step 3: Execute multicall
  console.log(`Executing ${calls.length} calls...`)
  const results = await client.multicall({ contracts: calls })

  // Step 4: Process results
  console.log('Processing results...')
  const bridgeTokens: BridgeTokenData = {}

  for (let t = 0; t < tokenIDs.length; t++) {
    const tokenID = tokenIDs[t]
    bridgeTokens[tokenID] = {}

    for (let c = 0; c < validChains.length; c++) {
      const index = t * validChains.length + c
      const result = results[index]

      if (result.status === 'failure') {
        console.error(
          `Failed to get token ${tokenID} on chain ${validChains[c]}`
        )
        process.exit(1)
      }

      if (!result.result) {
        continue
      }

      const token = result.result

      // Skip unsupported chains without token address
      if (
        token.tokenAddress.length === 0 ||
        token.tokenAddress === zeroAddress
      ) {
        continue
      }
      // sanity check chainId
      if (token.chainId !== BigInt(allChainIds[c])) {
        console.error(
          `Chain ID mismatch for token ${tokenID} on chain ${validChains[c]}`
        )
        process.exit(1)
      }

      // Parse the token address and checksum it
      try {
        bridgeTokens[tokenID][validChains[c]] = {
          isUnderlying: token.isUnderlying,
          tokenAddress: getAddress(token.tokenAddress),
        }
      } catch (error) {
        console.error(
          `Failed to parse address for ${tokenID} on ${validChains[c]}: ${token.tokenAddress}`
        )
        throw error
      }
    }
  }

  // Step 5: Sort the output and write
  const sortedTokens = sortBridgeTokenData(bridgeTokens)
  const outputPath = path.join(__dirname, '../../configs/global/tokens.json')
  fs.writeFileSync(outputPath, JSON.stringify(sortedTokens, null, 2) + '\n')

  console.log(`Successfully wrote token data to ${outputPath}`)
}

// Run the script
main().catch((error) => {
  console.error('Error:', error)
  process.exit(1)
})
