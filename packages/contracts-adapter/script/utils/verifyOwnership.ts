import {
  createPublicClient,
  http,
  getAddress,
  type PublicClient,
  defineChain,
  type Address,
} from 'viem'
import * as fs from 'node:fs'
import * as path from 'node:path'

// ANSI color codes for console output
const colors = {
  reset: '\x1b[0m',
  red: '\x1b[31m',
  green: '\x1b[32m',
  yellow: '\x1b[33m',
  cyan: '\x1b[36m',
  bold: '\x1b[1m',
  dim: '\x1b[2m',
}

const synapseBridgeAdapterAbi = [
  {
    name: 'owner',
    type: 'function',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'address' }],
  },
  {
    name: 'endpoint',
    type: 'function',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'address' }],
  },
] as const

const layerZeroEndpointAbi = [
  {
    name: 'delegates',
    type: 'function',
    stateMutability: 'view',
    inputs: [{ name: '_oapp', type: 'address' }],
    outputs: [{ name: '', type: 'address' }],
  },
] as const

// Multicall3 address (universal deployment)
const MULTICALL3_ADDRESS = '0xcA11bde05977b3631167028862bE2a173976CA11'

type DeploymentInfo = {
  address: Address
}

type MultisigConfig = {
  [chain: string]: Address
}

type VerificationIssue = {
  chain: string
  category: 'owner' | 'delegate' | 'error'
  severity: 'error' | 'warning' | 'info'
  message: string
}

type ChainVerificationResult = {
  chain: string
  chainId: number
  deployment: Address
  actualOwner: Address | null
  actualDelegate: Address | null
  expectedOwner: Address | null
  issues: VerificationIssue[]
}

type ResultStats = {
  totalChains: number
  chainsWithIssues: number
  errors: number
  warnings: number
  totalIssues: number
}

type AddressCallResult = {
  status: 'success' | 'failure'
  result?: Address
}

const categoryLabels: Record<VerificationIssue['category'], string> = {
  owner: 'Ownership',
  delegate: 'Delegate',
  error: 'Errors',
}

const severityIcons: Record<VerificationIssue['severity'], string> = {
  error: '✗',
  warning: '⚠',
  info: '✓',
}

const severityColors: Record<VerificationIssue['severity'], string> = {
  error: colors.red,
  warning: colors.yellow,
  info: colors.green,
}

function hasActionableIssue(issue: VerificationIssue): boolean {
  return issue.severity === 'error' || issue.severity === 'warning'
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

function groupIssuesByCategory(
  issues: VerificationIssue[]
): Record<string, VerificationIssue[]> {
  return issues.reduce((grouped, issue) => {
    if (!grouped[issue.category]) {
      grouped[issue.category] = []
    }
    grouped[issue.category].push(issue)
    return grouped
  }, {} as Record<string, VerificationIssue[]>)
}

function calculateResultStats(results: ChainVerificationResult[]): ResultStats {
  const totalChains = results.length
  const chainsWithIssues = results.filter((result) =>
    result.issues.some(hasActionableIssue)
  ).length
  const errors = results.reduce(
    (sum, result) =>
      sum + result.issues.filter((issue) => issue.severity === 'error').length,
    0
  )
  const warnings = results.reduce(
    (sum, result) =>
      sum +
      result.issues.filter((issue) => issue.severity === 'warning').length,
    0
  )

  return {
    totalChains,
    chainsWithIssues,
    errors,
    warnings,
    totalIssues: errors + warnings,
  }
}

function printOverview(stats: ResultStats) {
  console.log(`Total chains verified: ${stats.totalChains}`)
  console.log(
    `${colors.green}✓ Chains without issues: ${
      stats.totalChains - stats.chainsWithIssues
    }${colors.reset}`
  )

  if (stats.chainsWithIssues === 0) {
    return
  }

  console.log(
    `${colors.yellow}⚠ Chains with issues: ${stats.chainsWithIssues}${colors.reset}`
  )
  console.log(`  Total issues: ${stats.totalIssues}`)

  if (stats.errors > 0) {
    console.log(`  ${colors.red}✗ Errors: ${stats.errors}${colors.reset}`)
  }

  if (stats.warnings > 0) {
    console.log(
      `  ${colors.yellow}⚠ Warnings: ${stats.warnings}${colors.reset}`
    )
  }
}

function printChainResult(result: ChainVerificationResult) {
  console.log(
    `\n${colors.bold}${colors.cyan}${result.chain} (Chain ID: ${result.chainId})${colors.reset}`
  )
  console.log(`  Deployment: ${result.deployment}`)
  console.log(
    `  Expected multisig: ${
      result.expectedOwner || colors.red + 'Missing' + colors.reset
    }`
  )
  console.log(
    `  Actual owner: ${
      result.actualOwner || colors.red + 'Unknown' + colors.reset
    }`
  )
  console.log(
    `  Actual delegate: ${
      result.actualDelegate || colors.red + 'Unknown' + colors.reset
    }`
  )

  const errorWarningIssues = result.issues.filter(hasActionableIssue)
  if (errorWarningIssues.length === 0) {
    console.log(`  ${colors.green}✓ All checks passed${colors.reset}`)
  } else {
    console.log(
      `  ${colors.yellow}Issues found: ${errorWarningIssues.length}${colors.reset}`
    )
  }

  const issuesByCategory = groupIssuesByCategory(result.issues)
  for (const [category, issues] of Object.entries(issuesByCategory)) {
    const categoryLabel =
      categoryLabels[category as VerificationIssue['category']] || category

    console.log(`\n    ${colors.bold}${categoryLabel}:${colors.reset}`)
    for (const issue of issues) {
      const icon = severityIcons[issue.severity]
      const color = severityColors[issue.severity]
      console.log(`      ${color}${icon} ${issue.message}${colors.reset}`)
    }
  }
}

function printSummary(stats: ResultStats) {
  console.log(`\n${colors.bold}${colors.cyan}Summary${colors.reset}`)

  if (stats.totalIssues === 0) {
    console.log(
      `${colors.green}✓ All SynapseBridgeAdapter owners and delegates match multisig config!${colors.reset}\n`
    )
    return
  }

  console.log(
    `${colors.yellow}⚠ Found ${stats.totalIssues} issue(s) across ${stats.chainsWithIssues} chain(s)${colors.reset}`
  )

  if (stats.errors > 0) {
    console.log(
      `${colors.red}Please fix ${stats.errors} error(s) before proceeding${colors.reset}\n`
    )
    return
  }

  console.log(
    `${colors.yellow}Please review ${stats.warnings} warning(s)${colors.reset}\n`
  )
}

// Function to load chain IDs from deployment directories
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

// Function to load deployment info for a chain
function loadDeployment(chain: string): DeploymentInfo | null {
  const deploymentPath = path.join(
    __dirname,
    '../../deployments',
    chain,
    'SynapseBridgeAdapter.json'
  )

  if (!fs.existsSync(deploymentPath)) {
    return null
  }

  try {
    const deployment = JSON.parse(fs.readFileSync(deploymentPath, 'utf8'))
    return {
      address: getAddress(deployment.address),
    }
  } catch (error) {
    console.warn(
      `${
        colors.yellow
      }Warning: Failed to read deployment for ${chain}: ${formatError(error)}${
        colors.reset
      }`
    )
    return null
  }
}

function loadMultisigConfig(): MultisigConfig {
  const multisigConfigPath = path.join(
    __dirname,
    '../../configs/global/multisig.json'
  )
  const rawConfig = JSON.parse(
    fs.readFileSync(multisigConfigPath, 'utf8')
  ) as Record<string, string>

  return Object.fromEntries(
    Object.entries(rawConfig).map(([chain, owner]) => [
      chain,
      getAddress(owner),
    ])
  )
}

// Function to get RPC URL for a chain
function getRpcUrl(chainId: number): string {
  const baseUrl = process.env.RPC_BASE_URL

  if (!baseUrl) {
    throw new Error(
      'RPC_BASE_URL environment variable is not set. Please set it to your RPC provider base URL'
    )
  }

  // Special case for Moonbeam (chainId 1284)
  if (chainId === 1284) {
    return 'https://moonbeam.api.pocket.network'
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
    contracts: {
      multicall3: {
        address: MULTICALL3_ADDRESS,
      },
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

function createIssue(
  chain: string,
  category: VerificationIssue['category'],
  severity: VerificationIssue['severity'],
  message: string
): VerificationIssue {
  return { chain, category, severity, message }
}

function addIssue(
  result: ChainVerificationResult,
  category: VerificationIssue['category'],
  severity: VerificationIssue['severity'],
  message: string
) {
  result.issues.push(createIssue(result.chain, category, severity, message))
}

function assignAddressFromCall(
  result: ChainVerificationResult,
  callResult: AddressCallResult,
  failureMessage: string,
  assignAddress: (address: Address) => void
): Address | null {
  if (callResult.status !== 'success' || !callResult.result) {
    addIssue(result, 'error', 'error', failureMessage)
    return null
  }

  assignAddress(callResult.result)
  return callResult.result
}

function addMissingExpectedOwnerIssues(result: ChainVerificationResult) {
  result.issues.push(
    createIssue(
      result.chain,
      'owner',
      'error',
      'Missing expected owner in multisig config'
    ),
    createIssue(
      result.chain,
      'delegate',
      'error',
      'Missing expected delegate in multisig config'
    )
  )
}

function addAddressMatchIssue(
  result: ChainVerificationResult,
  category: 'owner' | 'delegate',
  expectedAddress: Address,
  actualAddress: Address | null,
  label: 'Owner' | 'Delegate'
) {
  if (actualAddress === null) {
    return
  }

  const addressesMatch =
    actualAddress.toLowerCase() === expectedAddress.toLowerCase()

  addIssue(
    result,
    category,
    addressesMatch ? 'info' : 'error',
    addressesMatch
      ? `${label} matches multisig config`
      : `${label} mismatch: expected ${expectedAddress}, got ${actualAddress}`
  )
}

async function fetchDelegateCallResult(
  client: PublicClient,
  endpointAddress: Address,
  deploymentAddress: Address
): Promise<AddressCallResult> {
  const [delegateResult] = await client.multicall({
    contracts: [
      {
        address: endpointAddress,
        abi: layerZeroEndpointAbi,
        functionName: 'delegates',
        args: [deploymentAddress],
      },
    ],
    allowFailure: true,
  })

  return delegateResult
}

async function verifyChain(
  chain: string,
  chainId: number,
  deployment: DeploymentInfo,
  multisigConfig: MultisigConfig
): Promise<ChainVerificationResult> {
  const result: ChainVerificationResult = {
    chain,
    chainId,
    deployment: deployment.address,
    actualOwner: null,
    actualDelegate: null,
    expectedOwner: multisigConfig[chain] || null,
    issues: [],
  }

  try {
    const client = createChainClient(chainId)
    const [ownerResult, endpointResult] = await client.multicall({
      contracts: [
        {
          address: deployment.address,
          abi: synapseBridgeAdapterAbi,
          functionName: 'owner',
        },
        {
          address: deployment.address,
          abi: synapseBridgeAdapterAbi,
          functionName: 'endpoint',
        },
      ],
      allowFailure: true,
    })

    assignAddressFromCall(
      result,
      ownerResult,
      'Failed to fetch current owner',
      (ownerAddress) => {
        result.actualOwner = ownerAddress
      }
    )

    const endpointAddress = assignAddressFromCall(
      result,
      endpointResult,
      'Failed to fetch app endpoint',
      () => undefined
    )

    if (endpointAddress !== null) {
      const delegateResult = await fetchDelegateCallResult(
        client,
        endpointAddress,
        deployment.address
      )

      assignAddressFromCall(
        result,
        delegateResult,
        'Failed to fetch current delegate',
        (delegateAddress) => {
          result.actualDelegate = delegateAddress
        }
      )
    }

    if (result.expectedOwner === null) {
      addMissingExpectedOwnerIssues(result)
      return result
    }

    addAddressMatchIssue(
      result,
      'owner',
      result.expectedOwner,
      result.actualOwner,
      'Owner'
    )
    addAddressMatchIssue(
      result,
      'delegate',
      result.expectedOwner,
      result.actualDelegate,
      'Delegate'
    )
  } catch (error) {
    addIssue(result, 'error', 'error', `RPC error: ${formatError(error)}`)
  }

  return result
}

function printResults(results: ChainVerificationResult[]) {
  console.log(
    `\n${colors.bold}${colors.cyan}SynapseBridgeAdapter Ownership & Delegate Verification Results${colors.reset}\n`
  )

  const stats = calculateResultStats(results)
  printOverview(stats)

  for (const result of results) {
    printChainResult(result)
  }

  printSummary(stats)
}

async function main() {
  console.log(
    `${colors.bold}Starting SynapseBridgeAdapter ownership verification...${colors.reset}\n`
  )

  const chainIds = loadChainIds()
  console.log(
    `Loaded ${Object.keys(chainIds).length} chain IDs from deployments`
  )

  const multisigConfig = loadMultisigConfig()
  console.log(
    `Loaded multisig config for ${Object.keys(multisigConfig).length} chains\n`
  )

  const deployedChains = Object.keys(chainIds).filter(
    (chain) => loadDeployment(chain) !== null
  )

  const requestedChains = process.argv.slice(2)
  const chainsToVerify =
    requestedChains.length > 0
      ? deployedChains.filter((chain) => requestedChains.includes(chain))
      : deployedChains

  if (requestedChains.length > 0 && chainsToVerify.length === 0) {
    console.error(
      `${
        colors.red
      }Error: None of the requested chains found: ${requestedChains.join(
        ', '
      )}${colors.reset}`
    )
    console.error(`Available chains: ${deployedChains.join(', ')}`)
    process.exit(1)
  }

  console.log(`Verifying ${chainsToVerify.length} deployed chains...\n`)

  const results: ChainVerificationResult[] = []

  for (const chain of chainsToVerify) {
    const chainId = chainIds[chain]
    const deployment = loadDeployment(chain)

    if (!deployment) {
      console.log(
        `${colors.yellow}Skipping ${chain}: deployment not found${colors.reset}`
      )
      continue
    }

    console.log(`${colors.dim}Verifying ${chain}...${colors.reset}`)
    const result = await verifyChain(chain, chainId, deployment, multisigConfig)
    results.push(result)
  }

  printResults(results)

  const hasErrors = results.some((r) =>
    r.issues.some((i) => i.severity === 'error')
  )
  if (hasErrors) {
    process.exit(1)
  }

  const hasWarnings = results.some((r) =>
    r.issues.some((i) => i.severity === 'warning')
  )
  if (hasWarnings) {
    process.exit(78)
  }
}

void (async () => {
  try {
    await main()
  } catch (error) {
    console.error(`${colors.red}Error: ${formatError(error)}${colors.reset}`)
    process.exit(1)
  }
})()
