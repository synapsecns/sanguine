import {
  createPublicClient,
  http,
  getAddress,
  type PublicClient,
  defineChain,
} from 'viem'
import * as fs from 'fs'
import * as path from 'path'

// Multicall3 address (universal deployment across chains)
const MULTICALL3_ADDRESS = '0xcA11bde05977b3631167028862bE2a173976CA11'

// Tokens to skip verification
const SKIP_VERIFICATION: Array<{ token: string; chain?: string }> = [
  { token: 'GMX' },
]

// ERC20 ABI for symbol and decimals
const erc20ABI = [
  {
    name: 'symbol',
    type: 'function',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'string' }],
  },
  {
    name: 'decimals',
    type: 'function',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint8' }],
  },
] as const

// ANSI color codes for console output
const colors = {
  reset: '\x1b[0m',
  red: '\x1b[31m',
  green: '\x1b[32m',
  yellow: '\x1b[33m',
  cyan: '\x1b[36m',
  bold: '\x1b[1m',
}

type TokenConfig = {
  [tokenSymbol: string]: {
    [chain: string]: {
      tokenAddress: string
      isUnderlying: boolean
    }
  }
}

type TokenVerificationResult = {
  tokenSymbol: string
  chain: string
  chainId: number
  address: string
  onChainSymbol: string | null
  onChainDecimals: number | null
  error: string | null
}

type VerificationSummary = {
  timestamp: string
  totalTokens: number
  totalDeployments: number
  successfulVerifications: number
  failedVerifications: number
  mismatches: {
    symbolMismatches: Array<{
      tokenSymbol: string
      symbols: Record<string, string[]> // symbol -> chains[]
    }>
    decimalMismatches: Array<{
      tokenSymbol: string
      decimals: Record<number, string[]> // decimals -> chains[]
    }>
  }
}

// Function to check if a token should be skipped
function shouldSkipVerification(tokenSymbol: string, chain: string): boolean {
  return SKIP_VERIFICATION.some(
    (skip) =>
      skip.token === tokenSymbol &&
      (skip.chain === undefined || skip.chain === chain)
  )
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

  // Define chain with multicall3 contract
  const chain = defineChain({
    id: chainId,
    name: `Chain ${chainId}`,
    nativeCurrency: { name: 'ETH', symbol: 'ETH', decimals: 18 },
    rpcUrls: {
      default: { http: [rpcUrl] },
    },
    contracts: {
      multicall3: {
        address: MULTICALL3_ADDRESS,
      },
    },
  })

  return createPublicClient({
    chain,
    transport: http(rpcUrl, {
      timeout: 30_000, // 30 second timeout
      retryCount: 3,
    }),
  })
}

// Function to verify a single token deployment
async function verifyTokenDeployment(
  tokenSymbol: string,
  chain: string,
  chainId: number,
  address: string
): Promise<TokenVerificationResult> {
  const result: TokenVerificationResult = {
    tokenSymbol,
    chain,
    chainId,
    address,
    onChainSymbol: null,
    onChainDecimals: null,
    error: null,
  }

  try {
    const client = createChainClient(chainId)
    const tokenAddress = getAddress(address)

    // Fetch symbol and decimals using multicall
    const results = await client.multicall({
      contracts: [
        {
          address: tokenAddress,
          abi: erc20ABI,
          functionName: 'symbol',
        },
        {
          address: tokenAddress,
          abi: erc20ABI,
          functionName: 'decimals',
        },
      ],
      allowFailure: true,
    })

    if (results[0].status === 'success') {
      result.onChainSymbol = results[0].result
    } else {
      result.error = `Failed to fetch symbol: ${
        results[0].error?.message || 'Unknown error'
      }`
    }

    if (results[1].status === 'success') {
      result.onChainDecimals = results[1].result
    } else {
      result.error = result.error
        ? `${result.error}; Failed to fetch decimals`
        : `Failed to fetch decimals: ${
            results[1].error?.message || 'Unknown error'
          }`
    }
  } catch (error) {
    result.error = `RPC error: ${
      error instanceof Error ? error.message : String(error)
    }`
  }

  return result
}

// Function to detect mismatches
function detectMismatches(results: TokenVerificationResult[]): {
  symbolMismatches: Map<string, Set<string>>
  decimalMismatches: Map<string, Set<number>>
} {
  const symbolsByToken = new Map<string, Set<string>>()
  const decimalsByToken = new Map<string, Set<number>>()

  for (const result of results) {
    if (result.error) continue

    if (!symbolsByToken.has(result.tokenSymbol)) {
      symbolsByToken.set(result.tokenSymbol, new Set())
    }
    if (result.onChainSymbol) {
      symbolsByToken.get(result.tokenSymbol)!.add(result.onChainSymbol)
    }

    if (!decimalsByToken.has(result.tokenSymbol)) {
      decimalsByToken.set(result.tokenSymbol, new Set())
    }
    if (result.onChainDecimals !== null) {
      decimalsByToken.get(result.tokenSymbol)!.add(result.onChainDecimals)
    }
  }

  const symbolMismatches = new Map<string, Set<string>>()
  const decimalMismatches = new Map<string, Set<number>>()

  for (const [token, symbols] of symbolsByToken.entries()) {
    if (symbols.size > 1) {
      symbolMismatches.set(token, symbols)
    }
  }

  for (const [token, decimals] of decimalsByToken.entries()) {
    if (decimals.size > 1) {
      decimalMismatches.set(token, decimals)
    }
  }

  return { symbolMismatches, decimalMismatches }
}

// Function to print results to console
function printResults(
  results: TokenVerificationResult[],
  symbolMismatches: Map<string, Set<string>>,
  decimalMismatches: Map<string, Set<number>>
) {
  console.log(
    `\n${colors.bold}${colors.cyan}Token Verification Results${colors.reset}\n`
  )

  const successCount = results.filter((r) => !r.error).length
  const failCount = results.filter((r) => r.error).length

  console.log(`Total deployments verified: ${results.length}`)
  console.log(`${colors.green}✓ Successful: ${successCount}${colors.reset}`)
  if (failCount > 0) {
    console.log(`${colors.red}✗ Failed: ${failCount}${colors.reset}`)
  }

  // Group results by token
  const resultsByToken = new Map<string, TokenVerificationResult[]>()
  for (const result of results) {
    if (!resultsByToken.has(result.tokenSymbol)) {
      resultsByToken.set(result.tokenSymbol, [])
    }
    resultsByToken.get(result.tokenSymbol)!.push(result)
  }

  // Print symbol mismatches
  if (symbolMismatches.size > 0) {
    console.log(
      `\n${colors.bold}${colors.red}Symbol Mismatches Found:${colors.reset}`
    )
    for (const token of symbolMismatches.keys()) {
      console.log(
        `\n  ${colors.yellow}${token}${colors.reset} has different symbols:`
      )
      const tokenResults = resultsByToken.get(token)!
      for (const result of tokenResults) {
        if (result.onChainSymbol) {
          console.log(
            `    ${result.chain} (${result.chainId}): ${colors.red}${result.onChainSymbol}${colors.reset}`
          )
        }
      }
    }
  }

  // Print decimal mismatches
  if (decimalMismatches.size > 0) {
    console.log(
      `\n${colors.bold}${colors.red}Decimal Mismatches Found:${colors.reset}`
    )
    for (const token of decimalMismatches.keys()) {
      console.log(
        `\n  ${colors.yellow}${token}${colors.reset} has different decimals:`
      )
      const tokenResults = resultsByToken.get(token)!
      for (const result of tokenResults) {
        if (result.onChainDecimals !== null) {
          console.log(
            `    ${result.chain} (${result.chainId}): ${colors.red}${result.onChainDecimals}${colors.reset}`
          )
        }
      }
    }
  }

  // Print errors if any
  const errors = results.filter((r) => r.error)
  if (errors.length > 0) {
    console.log(`\n${colors.bold}${colors.yellow}Errors:${colors.reset}`)
    for (const result of errors) {
      console.log(
        `  ${result.tokenSymbol} on ${result.chain} (${result.chainId}): ${colors.red}${result.error}${colors.reset}`
      )
    }
  }

  // Print summary
  if (
    symbolMismatches.size === 0 &&
    decimalMismatches.size === 0 &&
    errors.length === 0
  ) {
    console.log(
      `\n${colors.bold}${colors.green}✓ All tokens verified successfully with consistent symbols and decimals!${colors.reset}\n`
    )
  } else {
    console.log(
      `\n${colors.bold}${colors.yellow}⚠ Verification completed with warnings${colors.reset}\n`
    )
  }
}

async function main() {
  console.log(`${colors.bold}Starting token verification...${colors.reset}\n`)

  // Load chain IDs
  const chainIds = loadChainIds()
  console.log(
    `Loaded ${Object.keys(chainIds).length} chain IDs from deployments`
  )

  // Load tokens config
  const tokensConfigPath = path.join(
    __dirname,
    '../../configs/global/tokens.json'
  )

  if (!fs.existsSync(tokensConfigPath)) {
    throw new Error(`Tokens config not found: ${tokensConfigPath}`)
  }

  const tokensConfig: TokenConfig = JSON.parse(
    fs.readFileSync(tokensConfigPath, 'utf8')
  )

  const tokenSymbols = Object.keys(tokensConfig)
  console.log(`Found ${tokenSymbols.length} tokens in config`)

  // Build list of all deployments to verify
  const deploymentsToVerify: Array<{
    tokenSymbol: string
    chain: string
    address: string
  }> = []
  let skippedCount = 0

  for (const [tokenSymbol, chains] of Object.entries(tokensConfig)) {
    for (const [chain, config] of Object.entries(chains)) {
      if (chainIds[chain]) {
        if (shouldSkipVerification(tokenSymbol, chain)) {
          skippedCount++
          console.log(
            `${colors.cyan}Skipping ${tokenSymbol} on ${chain} (in skip list)${colors.reset}`
          )
        } else {
          deploymentsToVerify.push({
            tokenSymbol,
            chain,
            address: config.tokenAddress,
          })
        }
      }
    }
  }

  if (skippedCount > 0) {
    console.log(`\nSkipped ${skippedCount} deployment(s) from verification`)
  }
  console.log(
    `Verifying ${deploymentsToVerify.length} token deployments across ${
      Object.keys(chainIds).length
    } chains...\n`
  )

  // Verify each deployment
  const results: TokenVerificationResult[] = []
  let processed = 0

  for (const deployment of deploymentsToVerify) {
    const chainId = chainIds[deployment.chain]
    const result = await verifyTokenDeployment(
      deployment.tokenSymbol,
      deployment.chain,
      chainId,
      deployment.address
    )
    results.push(result)

    processed++
    if (processed % 10 === 0) {
      console.log(`Progress: ${processed}/${deploymentsToVerify.length}`)
    }
  }

  console.log(`\nCompleted verification of ${results.length} deployments`)

  // Detect mismatches
  const { symbolMismatches, decimalMismatches } = detectMismatches(results)

  // Print results to console
  printResults(results, symbolMismatches, decimalMismatches)

  // Prepare summary for JSON output
  const summary: VerificationSummary = {
    timestamp: new Date().toISOString(),
    totalTokens: tokenSymbols.length,
    totalDeployments: results.length,
    successfulVerifications: results.filter((r) => !r.error).length,
    failedVerifications: results.filter((r) => r.error).length,
    mismatches: {
      symbolMismatches: Array.from(symbolMismatches.keys()).map(
        (tokenSymbol) => {
          // Group chains by their on-chain symbol
          const symbolToChains: Record<string, string[]> = {}
          results
            .filter(
              (r) => r.tokenSymbol === tokenSymbol && r.onChainSymbol !== null
            )
            .forEach((r) => {
              const symbol = r.onChainSymbol!
              if (!symbolToChains[symbol]) {
                symbolToChains[symbol] = []
              }
              symbolToChains[symbol].push(r.chain)
            })
          return {
            tokenSymbol,
            symbols: symbolToChains,
          }
        }
      ),
      decimalMismatches: Array.from(decimalMismatches.keys()).map(
        (tokenSymbol) => {
          // Group chains by their on-chain decimals
          const decimalsToChains: Record<number, string[]> = {}
          results
            .filter(
              (r) => r.tokenSymbol === tokenSymbol && r.onChainDecimals !== null
            )
            .forEach((r) => {
              const decimals = r.onChainDecimals!
              if (!decimalsToChains[decimals]) {
                decimalsToChains[decimals] = []
              }
              decimalsToChains[decimals].push(r.chain)
            })
          return {
            tokenSymbol,
            decimals: decimalsToChains,
          }
        }
      ),
    },
  }

  // Save to JSON file
  const outputPath = path.join(
    __dirname,
    '../../configs/global/token-verification.json'
  )
  fs.writeFileSync(outputPath, JSON.stringify(summary, null, 2) + '\n')
  console.log(`Verification results saved to ${outputPath}`)

  // Exit with warning code if mismatches found
  if (
    symbolMismatches.size > 0 ||
    decimalMismatches.size > 0 ||
    results.some((r) => r.error)
  ) {
    process.exit(78)
  }
}

// Run the script
main().catch((error) => {
  console.error(`${colors.red}Error: ${error}${colors.reset}`)
  process.exit(1)
})
