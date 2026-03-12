import {
  createPublicClient,
  http,
  getAddress,
  type PublicClient,
  defineChain,
  type Address,
} from 'viem'
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
  category: 'owner' | 'error'
  severity: 'error' | 'warning' | 'info'
  message: string
}

type ChainVerificationResult = {
  chain: string
  chainId: number
  deployment: Address
  actualOwner: Address | null
  expectedOwner: Address | null
  issues: VerificationIssue[]
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
      `${colors.yellow}Warning: Failed to read deployment for ${chain}${colors.reset}`
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
    Object.entries(rawConfig).map(([chain, owner]) => [chain, getAddress(owner)])
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

function pluralize(count: number, word: string): string {
  return `${count} ${word}${count === 1 ? '' : 's'}`
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
    expectedOwner: multisigConfig[chain] || null,
    issues: [],
  }

  try {
    const client = createChainClient(chainId)
    const ownerResult = await client.multicall({
      contracts: [
        {
          address: deployment.address,
          abi: synapseBridgeAdapterAbi,
          functionName: 'owner',
        },
      ],
      allowFailure: true,
    })

    if (ownerResult[0].status === 'success' && ownerResult[0].result) {
      result.actualOwner = ownerResult[0].result
    } else {
      result.issues.push({
        chain,
        category: 'error',
        severity: 'error',
        message: 'Failed to fetch current owner',
      })
      return result
    }

    if (!result.expectedOwner) {
      result.issues.push({
        chain,
        category: 'owner',
        severity: 'error',
        message: 'Missing expected owner in multisig config',
      })
      return result
    }

    if (
      result.actualOwner.toLowerCase() !== result.expectedOwner.toLowerCase()
    ) {
      result.issues.push({
        chain,
        category: 'owner',
        severity: 'error',
        message: `Owner mismatch: expected ${result.expectedOwner}, got ${result.actualOwner}`,
      })
    } else {
      result.issues.push({
        chain,
        category: 'owner',
        severity: 'info',
        message: 'Owner matches multisig config',
      })
    }
  } catch (error) {
    result.issues.push({
      chain,
      category: 'error',
      severity: 'error',
      message: `RPC error: ${
        error instanceof Error ? error.message : String(error)
      }`,
    })
  }

  return result
}

function printResults(results: ChainVerificationResult[]) {
  const categoryLabels: Record<string, string> = {
    owner: 'Ownership',
    error: 'Errors',
  }

  const severityIcons: Record<string, string> = {
    error: '✗',
    warning: '⚠',
    info: '✓',
  }

  const severityColors: Record<string, string> = {
    error: colors.red,
    warning: colors.yellow,
    info: colors.green,
  }

  console.log(
    `\n${colors.bold}${colors.cyan}SynapseBridgeAdapter Ownership Verification Results${colors.reset}\n`
  )

  const totalChains = results.length
  const chainsWithIssues = results.filter((r) =>
    r.issues.some((i) => i.severity === 'error' || i.severity === 'warning')
  ).length
  const errors = results.reduce(
    (sum, r) => sum + r.issues.filter((i) => i.severity === 'error').length,
    0
  )
  const warnings = results.reduce(
    (sum, r) => sum + r.issues.filter((i) => i.severity === 'warning').length,
    0
  )
  const totalIssues = errors + warnings

  console.log(`Total chains verified: ${totalChains}`)
  console.log(
    `${colors.green}✓ Chains without issues: ${totalChains - chainsWithIssues}${
      colors.reset
    }`
  )
  if (chainsWithIssues > 0) {
    console.log(
      `${colors.yellow}⚠ Chains with issues: ${chainsWithIssues}${colors.reset}`
    )
    console.log(`  Total issues: ${totalIssues}`)
    if (errors > 0) {
      console.log(`  ${colors.red}✗ Errors: ${errors}${colors.reset}`)
    }
    if (warnings > 0) {
      console.log(`  ${colors.yellow}⚠ Warnings: ${warnings}${colors.reset}`)
    }
  }

  for (const result of results) {
    console.log(
      `\n${colors.bold}${colors.cyan}${result.chain} (Chain ID: ${result.chainId})${colors.reset}`
    )
    console.log(`  Deployment: ${result.deployment}`)
    console.log(
      `  Expected owner: ${
        result.expectedOwner || colors.red + 'Missing' + colors.reset
      }`
    )
    console.log(
      `  Actual owner: ${
        result.actualOwner || colors.red + 'Unknown' + colors.reset
      }`
    )

    const errorWarningIssues = result.issues.filter(
      (i) => i.severity === 'error' || i.severity === 'warning'
    )

    if (errorWarningIssues.length === 0) {
      console.log(`  ${colors.green}✓ All checks passed${colors.reset}`)
    } else {
      console.log(
        `  ${colors.yellow}Issues found: ${errorWarningIssues.length}${colors.reset}`
      )
    }

    const issuesByCategory = result.issues.reduce((acc, issue) => {
      if (!acc[issue.category]) {
        acc[issue.category] = []
      }
      acc[issue.category].push(issue)
      return acc
    }, {} as Record<string, VerificationIssue[]>)

    for (const [category, issues] of Object.entries(issuesByCategory)) {
      const categoryLabel = categoryLabels[category] || category

      console.log(`\n    ${colors.bold}${categoryLabel}:${colors.reset}`)
      for (const issue of issues) {
        const icon = severityIcons[issue.severity] || '?'
        const color = severityColors[issue.severity] || colors.reset
        console.log(`      ${color}${icon} ${issue.message}${colors.reset}`)
      }
    }
  }

  console.log(`\n${colors.bold}${colors.cyan}Summary${colors.reset}`)
  if (totalIssues === 0) {
    console.log(
      `${colors.green}✓ All SynapseBridgeAdapter owners match multisig config!${colors.reset}\n`
    )
  } else {
    console.log(
      `${colors.yellow}⚠ Found ${totalIssues} issue(s) across ${chainsWithIssues} chain(s)${colors.reset}`
    )
    if (errors > 0) {
      console.log(
        `${colors.red}Please fix ${errors} error(s) before proceeding${colors.reset}\n`
      )
    } else {
      console.log(
        `${colors.yellow}Please review ${warnings} warning(s)${colors.reset}\n`
      )
    }
  }
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
    const result = await verifyChain(
      chain,
      chainId,
      deployment,
      multisigConfig
    )
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

main().catch((error) => {
  console.error(`${colors.red}Error: ${error}${colors.reset}`)
  process.exit(1)
})
