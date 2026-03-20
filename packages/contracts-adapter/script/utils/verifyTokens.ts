import {
  createPublicClient,
  http,
  getAddress,
  type PublicClient,
  defineChain,
} from 'viem'
import * as fs from 'node:fs'
import * as path from 'node:path'

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

type TokenDeploymentTarget = {
  tokenSymbol: string
  chain: string
  address: string
}

function formatError(error: unknown): string {
  if (error instanceof Error) {
    return error.message
  }

  if (typeof error === 'object' && error !== null) {
    try {
      return JSON.stringify(error)
    } catch {
      return 'Unknown object error'
    }
  }

  return String(error)
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
        const chainId = Number.parseInt(
          fs.readFileSync(chainIdPath, 'utf8').trim(),
          10
        )
        if (!Number.isNaN(chainId)) {
          chainIds[dir] = chainId
        }
      } catch (error) {
        console.warn(
          `${
            colors.yellow
          }Warning: Failed to read chain ID for ${dir}: ${formatError(error)}${
            colors.reset
          }`
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
    return 'https://rpc.api.moonbeam.network'
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
    result.error = `RPC error: ${formatError(error)}`
  }

  return result
}

function getOrCreateSet<K, V>(map: Map<K, Set<V>>, key: K): Set<V> {
  let existing = map.get(key)
  if (!existing) {
    existing = new Set<V>()
    map.set(key, existing)
  }

  return existing
}

function addObservedValues(
  symbolsByToken: Map<string, Set<string>>,
  decimalsByToken: Map<string, Set<number>>,
  result: TokenVerificationResult
) {
  const symbols = getOrCreateSet(symbolsByToken, result.tokenSymbol)
  if (result.onChainSymbol) {
    symbols.add(result.onChainSymbol)
  }

  const decimals = getOrCreateSet(decimalsByToken, result.tokenSymbol)
  if (result.onChainDecimals !== null) {
    decimals.add(result.onChainDecimals)
  }
}

function collectMultiValueEntries<T>(
  valuesByToken: Map<string, Set<T>>
): Map<string, Set<T>> {
  const mismatches = new Map<string, Set<T>>()

  for (const [token, values] of valuesByToken.entries()) {
    if (values.size > 1) {
      mismatches.set(token, values)
    }
  }

  return mismatches
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

    addObservedValues(symbolsByToken, decimalsByToken, result)
  }

  return {
    symbolMismatches: collectMultiValueEntries(symbolsByToken),
    decimalMismatches: collectMultiValueEntries(decimalsByToken),
  }
}

function groupResultsByToken(
  results: TokenVerificationResult[]
): Map<string, TokenVerificationResult[]> {
  const grouped = new Map<string, TokenVerificationResult[]>()

  for (const result of results) {
    const tokenResults = grouped.get(result.tokenSymbol)
    if (tokenResults) {
      tokenResults.push(result)
      continue
    }

    grouped.set(result.tokenSymbol, [result])
  }

  return grouped
}

function printMismatchSection<T>(
  title: string,
  description: string,
  mismatches: Map<string, Set<T>>,
  resultsByToken: Map<string, TokenVerificationResult[]>,
  valueSelector: (result: TokenVerificationResult) => T | null
) {
  if (mismatches.size === 0) {
    return
  }

  console.log(`\n${colors.bold}${colors.red}${title}:${colors.reset}`)
  for (const token of mismatches.keys()) {
    console.log(`\n  ${colors.yellow}${token}${colors.reset} ${description}`)

    for (const result of resultsByToken.get(token) || []) {
      const value = valueSelector(result)
      if (value === null) {
        continue
      }

      console.log(
        `    ${result.chain} (${result.chainId}): ${colors.red}${value}${colors.reset}`
      )
    }
  }
}

function printErrors(errors: TokenVerificationResult[]) {
  if (errors.length === 0) {
    return
  }

  console.log(`\n${colors.bold}${colors.yellow}Errors:${colors.reset}`)
  for (const result of errors) {
    console.log(
      `  ${result.tokenSymbol} on ${result.chain} (${result.chainId}): ${colors.red}${result.error}${colors.reset}`
    )
  }
}

function printVerificationSummary(
  symbolMismatches: Map<string, Set<string>>,
  decimalMismatches: Map<string, Set<number>>,
  errors: TokenVerificationResult[]
) {
  if (
    symbolMismatches.size === 0 &&
    decimalMismatches.size === 0 &&
    errors.length === 0
  ) {
    console.log(
      `\n${colors.bold}${colors.green}✓ All tokens verified successfully with consistent symbols and decimals!${colors.reset}\n`
    )
    return
  }

  console.log(
    `\n${colors.bold}${colors.yellow}⚠ Verification completed with warnings${colors.reset}\n`
  )
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

  const resultsByToken = groupResultsByToken(results)
  printMismatchSection(
    'Symbol Mismatches Found',
    'has different symbols:',
    symbolMismatches,
    resultsByToken,
    (result) => result.onChainSymbol
  )
  printMismatchSection(
    'Decimal Mismatches Found',
    'has different decimals:',
    decimalMismatches,
    resultsByToken,
    (result) => result.onChainDecimals
  )

  const errors = results.filter((r) => r.error)
  printErrors(errors)
  printVerificationSummary(symbolMismatches, decimalMismatches, errors)
}

function loadTokensConfig(): TokenConfig {
  const tokensConfigPath = path.join(
    __dirname,
    '../../configs/global/tokens.json'
  )

  if (!fs.existsSync(tokensConfigPath)) {
    throw new Error(`Tokens config not found: ${tokensConfigPath}`)
  }

  return JSON.parse(fs.readFileSync(tokensConfigPath, 'utf8')) as TokenConfig
}

function collectDeploymentsToVerify(
  tokensConfig: TokenConfig,
  chainIds: Record<string, number>
): { deploymentsToVerify: TokenDeploymentTarget[]; skippedCount: number } {
  const deploymentsToVerify: TokenDeploymentTarget[] = []
  let skippedCount = 0

  for (const [tokenSymbol, chains] of Object.entries(tokensConfig)) {
    for (const [chain, config] of Object.entries(chains)) {
      if (!chainIds[chain]) {
        continue
      }

      if (shouldSkipVerification(tokenSymbol, chain)) {
        skippedCount++
        console.log(
          `${colors.cyan}Skipping ${tokenSymbol} on ${chain} (in skip list)${colors.reset}`
        )
        continue
      }

      deploymentsToVerify.push({
        tokenSymbol,
        chain,
        address: config.tokenAddress,
      })
    }
  }

  return { deploymentsToVerify, skippedCount }
}

async function verifyDeployments(
  deploymentsToVerify: TokenDeploymentTarget[],
  chainIds: Record<string, number>
): Promise<TokenVerificationResult[]> {
  const results: TokenVerificationResult[] = []
  let processed = 0

  for (const deployment of deploymentsToVerify) {
    const chainId = chainIds[deployment.chain]
    results.push(
      await verifyTokenDeployment(
        deployment.tokenSymbol,
        deployment.chain,
        chainId,
        deployment.address
      )
    )

    processed++
    if (processed % 10 === 0) {
      console.log(`Progress: ${processed}/${deploymentsToVerify.length}`)
    }
  }

  return results
}

function groupChainsBySymbol(
  results: TokenVerificationResult[],
  tokenSymbol: string
): Record<string, string[]> {
  const symbolToChains: Record<string, string[]> = {}

  results
    .filter((result) => result.tokenSymbol === tokenSymbol)
    .forEach((result) => {
      if (result.onChainSymbol === null) {
        return
      }

      const symbol = result.onChainSymbol
      if (!symbolToChains[symbol]) {
        symbolToChains[symbol] = []
      }
      symbolToChains[symbol].push(result.chain)
    })

  return symbolToChains
}

function groupChainsByDecimals(
  results: TokenVerificationResult[],
  tokenSymbol: string
): Record<number, string[]> {
  const decimalsToChains: Record<number, string[]> = {}

  results
    .filter((result) => result.tokenSymbol === tokenSymbol)
    .forEach((result) => {
      if (result.onChainDecimals === null) {
        return
      }

      const decimals = result.onChainDecimals
      if (!decimalsToChains[decimals]) {
        decimalsToChains[decimals] = []
      }
      decimalsToChains[decimals].push(result.chain)
    })

  return decimalsToChains
}

function buildSummary(
  tokenSymbols: string[],
  results: TokenVerificationResult[],
  symbolMismatches: Map<string, Set<string>>,
  decimalMismatches: Map<string, Set<number>>
): VerificationSummary {
  return {
    timestamp: new Date().toISOString(),
    totalTokens: tokenSymbols.length,
    totalDeployments: results.length,
    successfulVerifications: results.filter((result) => !result.error).length,
    failedVerifications: results.filter((result) => result.error).length,
    mismatches: {
      symbolMismatches: Array.from(symbolMismatches.keys()).map(
        (tokenSymbol) => ({
          tokenSymbol,
          symbols: groupChainsBySymbol(results, tokenSymbol),
        })
      ),
      decimalMismatches: Array.from(decimalMismatches.keys()).map(
        (tokenSymbol) => ({
          tokenSymbol,
          decimals: groupChainsByDecimals(results, tokenSymbol),
        })
      ),
    },
  }
}

function hasVerificationWarnings(
  results: TokenVerificationResult[],
  symbolMismatches: Map<string, Set<string>>,
  decimalMismatches: Map<string, Set<number>>
): boolean {
  return (
    symbolMismatches.size > 0 ||
    decimalMismatches.size > 0 ||
    results.some((result) => result.error)
  )
}

async function main() {
  console.log(`${colors.bold}Starting token verification...${colors.reset}\n`)

  // Load chain IDs
  const chainIds = loadChainIds()
  console.log(
    `Loaded ${Object.keys(chainIds).length} chain IDs from deployments`
  )

  const tokensConfig = loadTokensConfig()
  const tokenSymbols = Object.keys(tokensConfig)
  console.log(`Found ${tokenSymbols.length} tokens in config`)

  const { deploymentsToVerify, skippedCount } = collectDeploymentsToVerify(
    tokensConfig,
    chainIds
  )

  if (skippedCount > 0) {
    console.log(`\nSkipped ${skippedCount} deployment(s) from verification`)
  }
  console.log(
    `Verifying ${deploymentsToVerify.length} token deployments across ${
      Object.keys(chainIds).length
    } chains...\n`
  )

  const results = await verifyDeployments(deploymentsToVerify, chainIds)
  console.log(`\nCompleted verification of ${results.length} deployments`)

  const { symbolMismatches, decimalMismatches } = detectMismatches(results)
  printResults(results, symbolMismatches, decimalMismatches)

  const summary = buildSummary(
    tokenSymbols,
    results,
    symbolMismatches,
    decimalMismatches
  )

  const outputPath = path.join(
    __dirname,
    '../../configs/global/token-verification.json'
  )
  fs.writeFileSync(outputPath, JSON.stringify(summary, null, 2) + '\n')
  console.log(`Verification results saved to ${outputPath}`)

  if (hasVerificationWarnings(results, symbolMismatches, decimalMismatches)) {
    process.exit(78)
  }
}

// Run the script
void (async () => {
  try {
    await main()
  } catch (error) {
    console.error(`${colors.red}Error: ${formatError(error)}${colors.reset}`)
    process.exit(1)
  }
})()
