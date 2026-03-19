import {
  createPublicClient,
  http,
  getAddress,
  type PublicClient,
  defineChain,
  type Address,
  decodeAbiParameters,
  parseAbiParameters,
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

// LayerZero config types
const CONFIG_TYPE_ULN = 2

// ABI definitions
const synapseBridgeAdapterAbi = [
  {
    name: 'owner',
    type: 'function',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'address' }],
  },
  {
    name: 'bridge',
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
  {
    name: 'peers',
    type: 'function',
    stateMutability: 'view',
    inputs: [{ name: 'eid', type: 'uint32' }],
    outputs: [{ name: 'peer', type: 'bytes32' }],
  },
  {
    name: 'getTokenType',
    type: 'function',
    stateMutability: 'view',
    inputs: [{ name: 'localAddr', type: 'address' }],
    outputs: [{ name: 'tokenType', type: 'uint8' }],
  },
  {
    name: 'getRemoteAddress',
    type: 'function',
    stateMutability: 'view',
    inputs: [
      { name: 'eid', type: 'uint32' },
      { name: 'localAddr', type: 'address' },
    ],
    outputs: [{ name: 'remoteAddr', type: 'address' }],
  },
] as const

const layerZeroEndpointAbi = [
  {
    name: 'getSendLibrary',
    type: 'function',
    stateMutability: 'view',
    inputs: [
      { name: '_sender', type: 'address' },
      { name: '_dstEid', type: 'uint32' },
    ],
    outputs: [{ name: 'sendLibrary', type: 'address' }],
  },
  {
    name: 'isDefaultSendLibrary',
    type: 'function',
    stateMutability: 'view',
    inputs: [
      { name: '_sender', type: 'address' },
      { name: '_dstEid', type: 'uint32' },
    ],
    outputs: [{ name: 'isDefault', type: 'bool' }],
  },
  {
    name: 'getReceiveLibrary',
    type: 'function',
    stateMutability: 'view',
    inputs: [
      { name: '_receiver', type: 'address' },
      { name: '_srcEid', type: 'uint32' },
    ],
    outputs: [
      { name: 'receiveLib', type: 'address' },
      { name: 'isDefault', type: 'bool' },
    ],
  },
  {
    name: 'getConfig',
    type: 'function',
    stateMutability: 'view',
    inputs: [
      { name: '_oapp', type: 'address' },
      { name: '_lib', type: 'address' },
      { name: '_eid', type: 'uint32' },
      { name: '_configType', type: 'uint32' },
    ],
    outputs: [{ name: 'config', type: 'bytes' }],
  },
] as const

// Type definitions
type ChainInfo = {
  blockTime: number
  eid: number
  endpointV2: Address
  receiveUln302: Address
  sendUln302: Address
  synapseBridge: Address
}

type ChainsConfig = {
  [chain: string]: ChainInfo
}

type TokenConfig = {
  [tokenSymbol: string]: {
    [chain: string]: {
      tokenAddress: Address
      isUnderlying: boolean
    }
  }
}

type DVNConfig = {
  [chain: string]: {
    [dvnName: string]: Address
  }
}

type SecurityConfig = {
  blockConfirmations: Record<string, number>
  DVNs: string[]
}

type DeploymentInfo = {
  address: Address
}

type UlnConfig = {
  confirmations: bigint
  requiredDVNCount: number
  optionalDVNCount: number
  optionalDVNThreshold: number
  requiredDVNs: Address[]
  optionalDVNs: Address[]
}

type VerificationIssue = {
  chain: string
  category:
    | 'peer'
    | 'library'
    | 'dvn'
    | 'confirmations'
    | 'token'
    | 'bridge'
    | 'error'
  severity: 'error' | 'warning' | 'info'
  message: string
}

type ChainVerificationResult = {
  chain: string
  chainId: number
  owner: Address | null
  bridge: Address | null
  issues: VerificationIssue[]
}

type VerificationContext = {
  chainsConfig: ChainsConfig
  tokensConfig: TokenConfig
  dvnsConfig: DVNConfig
  securityConfig: SecurityConfig
  allChains: string[]
}

type RemoteChainDeployment = {
  chain: string
  eid: number
  deployment: DeploymentInfo
}

type BasicInfoVerification = {
  owner: Address | null
  bridge: Address | null
  issues: VerificationIssue[]
}

type TokenTypeRequest = {
  tokenSymbol: string
  tokenAddr: Address
}

type RemoteAddressRequest = {
  tokenSymbol: string
  remoteChain: string
  remoteEid: number
  expectedAddr: Address
  tokenAddr: Address
}

type TokenVerificationPlan = {
  tokenTypeRequests: TokenTypeRequest[]
  remoteAddressRequests: RemoteAddressRequest[]
  skippedGMX: boolean
  skippedSingleChainTokens: string[]
}

type CountedIssues = {
  successCount: number
  issues: VerificationIssue[]
}

type ResultStats = {
  totalChains: number
  chainsWithIssues: number
  errors: number
  warnings: number
  totalIssues: number
}

// Multicall3 address (universal deployment)
const MULTICALL3_ADDRESS = '0xcA11bde05977b3631167028862bE2a173976CA11'
const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000'
const ZERO_PEER = `0x${'0'.repeat(64)}`

const categoryLabels: Record<VerificationIssue['category'], string> = {
  peer: 'Peer Configuration',
  library: 'Library Configuration',
  dvn: 'DVN Configuration',
  confirmations: 'Block Confirmations',
  token: 'Token Configuration',
  bridge: 'Bridge Configuration',
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

function createIssue(
  chain: string,
  category: VerificationIssue['category'],
  severity: VerificationIssue['severity'],
  message: string
): VerificationIssue {
  return { chain, category, severity, message }
}

function hasActionableIssue(issue: VerificationIssue): boolean {
  return issue.severity === 'error' || issue.severity === 'warning'
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

// Function to decode ULN config from bytes
function decodeUlnConfig(configBytes: string): UlnConfig {
  if (configBytes === '0x' || configBytes === '0x0') {
    // Return default/empty config
    return {
      confirmations: 0n,
      requiredDVNCount: 0,
      optionalDVNCount: 0,
      optionalDVNThreshold: 0,
      requiredDVNs: [],
      optionalDVNs: [],
    }
  }

  // The config is returned wrapped in bytes encoding, which adds a 32-byte offset pointer
  // Skip the first 32 bytes (0x + 64 hex chars) to get to the actual struct data
  const structData = ('0x' + configBytes.slice(66)) as `0x${string}`

  // Now decode the actual UlnConfig struct
  const decoded = decodeAbiParameters(
    parseAbiParameters(
      'uint64 confirmations, uint8 requiredDVNCount, uint8 optionalDVNCount, uint8 optionalDVNThreshold, address[] requiredDVNs, address[] optionalDVNs'
    ),
    structData
  )

  return {
    confirmations: decoded[0],
    requiredDVNCount: decoded[1],
    optionalDVNCount: decoded[2],
    optionalDVNThreshold: decoded[3],
    requiredDVNs: decoded[4].map((addr) => getAddress(addr)),
    optionalDVNs: decoded[5].map((addr) => getAddress(addr)),
  }
}

// Function to sort addresses
function compareBigInts(left: bigint, right: bigint): number {
  if (left === right) {
    return 0
  }

  return left < right ? -1 : 1
}

function sortAddresses(addresses: Address[]): Address[] {
  return addresses
    .map((addr) => ({ addr, num: BigInt(addr) }))
    .sort((left, right) => compareBigInts(left.num, right.num))
    .map((item) => item.addr)
}

// Utility helper functions
function pluralize(count: number, word: string): string {
  return `${count} ${word}${count === 1 ? '' : 's'}`
}

function getExpectedConfirmations(
  securityConfig: SecurityConfig,
  chain: string
): number {
  const confirmations = securityConfig.blockConfirmations[chain]
  if (typeof confirmations !== 'number') {
    throw new TypeError(`Missing block confirmations for chain: ${chain}`)
  }
  return confirmations
}

function areDVNsEqual(actual: Address[], expected: Address[]): boolean {
  if (actual.length !== expected.length) return false
  const sortedActual = sortAddresses(actual)
  const sortedExpected = sortAddresses(expected)
  return sortedActual.every(
    (dvn, idx) => dvn.toLowerCase() === sortedExpected[idx].toLowerCase()
  )
}

// Function to verify a single ULN config (send or receive)
function verifyUlnConfig(
  chain: string,
  remoteChain: string,
  configBytes: string,
  expectedConfirmations: number,
  expectedDVNs: Address[],
  configDirection: 'send' | 'receive'
): VerificationIssue[] {
  const issues: VerificationIssue[] = []
  const directionLabel = configDirection === 'send' ? 'Send' : 'Receive'

  try {
    const decoded = decodeUlnConfig(configBytes)

    // Validate block confirmations
    if (decoded.confirmations !== BigInt(expectedConfirmations)) {
      issues.push(
        createIssue(
          chain,
          'confirmations',
          'warning',
          `${directionLabel} confirmations for ${remoteChain} incorrect: expected ${expectedConfirmations}, got ${decoded.confirmations.toString()}`
        )
      )
    }

    // Validate DVNs
    if (!areDVNsEqual(decoded.requiredDVNs, expectedDVNs)) {
      const sortedActual = sortAddresses(decoded.requiredDVNs)
      const sortedExpected = sortAddresses(expectedDVNs)
      issues.push(
        createIssue(
          chain,
          'dvn',
          'error',
          `${directionLabel} DVNs for ${remoteChain} incorrect: expected [${sortedExpected.join(
            ', '
          )}], got [${sortedActual.join(', ')}]`
        )
      )
    }
  } catch (error) {
    issues.push(
      createIssue(
        chain,
        'error',
        'error',
        `Failed to decode ${configDirection} config for ${remoteChain}: ${formatError(
          error
        )}`
      )
    )
  }

  return issues
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
  console.log(
    `  Owner: ${result.owner || colors.red + 'Unknown' + colors.reset}`
  )
  console.log(
    `  Bridge: ${result.bridge || colors.red + 'Not set' + colors.reset}`
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
      `${colors.green}✓ All SynapseBridgeAdapter contracts are correctly configured!${colors.reset}\n`
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

function getRemoteDeployments(
  chain: string,
  context: VerificationContext
): RemoteChainDeployment[] {
  const remoteDeployments: RemoteChainDeployment[] = []

  for (const remoteChain of context.allChains) {
    if (remoteChain === chain) {
      continue
    }

    const deployment = loadDeployment(remoteChain)
    if (!deployment) {
      continue
    }

    remoteDeployments.push({
      chain: remoteChain,
      eid: context.chainsConfig[remoteChain].eid,
      deployment,
    })
  }

  return remoteDeployments
}

function formatPeerAddress(address: Address): string {
  return `0x${address.slice(2).padStart(64, '0')}`
}

async function verifyBasicInfo(
  chain: string,
  client: PublicClient,
  deployment: DeploymentInfo,
  chainInfo: ChainInfo
): Promise<BasicInfoVerification> {
  const issues: VerificationIssue[] = []
  let owner: Address | null = null
  let bridge: Address | null = null

  const basicResults = await client.multicall({
    contracts: [
      {
        address: deployment.address,
        abi: synapseBridgeAdapterAbi,
        functionName: 'owner',
      },
      {
        address: deployment.address,
        abi: synapseBridgeAdapterAbi,
        functionName: 'bridge',
      },
      {
        address: deployment.address,
        abi: synapseBridgeAdapterAbi,
        functionName: 'endpoint',
      },
    ],
    allowFailure: true,
  })

  if (basicResults[0].status === 'success' && basicResults[0].result) {
    owner = basicResults[0].result
  }

  if (basicResults[1].status === 'success' && basicResults[1].result) {
    bridge = basicResults[1].result
    if (bridge === ZERO_ADDRESS) {
      issues.push(
        createIssue(chain, 'bridge', 'error', 'Bridge address not set')
      )
    } else if (bridge.toLowerCase() !== chainInfo.synapseBridge.toLowerCase()) {
      issues.push(
        createIssue(
          chain,
          'bridge',
          'warning',
          `Bridge address mismatch: expected ${chainInfo.synapseBridge}, got ${bridge}`
        )
      )
    }
  } else {
    issues.push(
      createIssue(chain, 'error', 'error', 'Failed to fetch bridge address')
    )
  }

  if (basicResults[2].status === 'success' && basicResults[2].result) {
    const endpoint = basicResults[2].result
    if (endpoint.toLowerCase() !== chainInfo.endpointV2.toLowerCase()) {
      issues.push(
        createIssue(
          chain,
          'error',
          'error',
          `Endpoint mismatch: expected ${chainInfo.endpointV2}, got ${endpoint}`
        )
      )
    }
  }

  return { owner, bridge, issues }
}

async function verifyPeerConfiguration(
  chain: string,
  client: PublicClient,
  deployment: DeploymentInfo,
  remoteDeployments: RemoteChainDeployment[]
): Promise<VerificationIssue[]> {
  if (remoteDeployments.length === 0) {
    return []
  }

  const peerResults = await client.multicall({
    contracts: remoteDeployments.map(({ eid }) => ({
      address: deployment.address,
      abi: synapseBridgeAdapterAbi,
      functionName: 'peers' as const,
      args: [eid] as const,
    })),
    allowFailure: true,
  })

  const issues: VerificationIssue[] = []
  let successCount = 0

  peerResults.forEach((peer, index) => {
    const remoteDeployment = remoteDeployments[index]
    const expectedPeer = formatPeerAddress(remoteDeployment.deployment.address)

    if (peer.status !== 'success' || !peer.result) {
      issues.push(
        createIssue(
          chain,
          'error',
          'error',
          `Failed to fetch peer for ${remoteDeployment.chain}`
        )
      )
      return
    }

    if (peer.result.toLowerCase() === expectedPeer.toLowerCase()) {
      successCount++
      return
    }

    if (peer.result === ZERO_PEER) {
      issues.push(
        createIssue(
          chain,
          'peer',
          'error',
          `Peer not set for ${remoteDeployment.chain} (eid: ${remoteDeployment.eid})`
        )
      )
      return
    }

    issues.push(
      createIssue(
        chain,
        'peer',
        'error',
        `Peer mismatch for ${remoteDeployment.chain}: expected ${expectedPeer}, got ${peer.result}`
      )
    )
  })

  if (successCount > 0) {
    issues.push(
      createIssue(
        chain,
        'peer',
        'info',
        `Verified ${pluralize(successCount, 'peer')}`
      )
    )
  }

  return issues
}

function getExpectedDVNs(
  chain: string,
  context: VerificationContext
): Address[] {
  return sortAddresses(
    context.securityConfig.DVNs.map(
      (dvnName) => context.dvnsConfig[chain][dvnName]
    )
  )
}

async function verifySendLibraries(
  chain: string,
  client: PublicClient,
  deployment: DeploymentInfo,
  chainInfo: ChainInfo,
  remoteDeployments: RemoteChainDeployment[]
): Promise<VerificationIssue[]> {
  if (remoteDeployments.length === 0) {
    return []
  }

  const sendLibraryResults = await client.multicall({
    contracts: remoteDeployments.map(({ eid }) => ({
      address: chainInfo.endpointV2,
      abi: layerZeroEndpointAbi,
      functionName: 'getSendLibrary' as const,
      args: [deployment.address, eid] as const,
    })),
    allowFailure: true,
  })

  const defaultLibraryResults = await client.multicall({
    contracts: remoteDeployments.map(({ eid }) => ({
      address: chainInfo.endpointV2,
      abi: layerZeroEndpointAbi,
      functionName: 'isDefaultSendLibrary' as const,
      args: [deployment.address, eid] as const,
    })),
    allowFailure: true,
  })

  const issues: VerificationIssue[] = []
  let successCount = 0

  sendLibraryResults.forEach((sendLibrary, index) => {
    if (sendLibrary.status !== 'success' || !sendLibrary.result) {
      return
    }

    const remoteChain = remoteDeployments[index].chain
    let hasError = false

    if (
      sendLibrary.result.toLowerCase() !== chainInfo.sendUln302.toLowerCase()
    ) {
      issues.push(
        createIssue(
          chain,
          'library',
          'error',
          `Send library for ${remoteChain} incorrect: expected ${chainInfo.sendUln302}, got ${sendLibrary.result}`
        )
      )
      hasError = true
    }

    const defaultLibrary = defaultLibraryResults[index]
    if (defaultLibrary.status === 'success' && defaultLibrary.result === true) {
      issues.push(
        createIssue(
          chain,
          'library',
          'error',
          `Send library for ${remoteChain} is using default library, should be custom`
        )
      )
      hasError = true
    }

    if (!hasError) {
      successCount++
    }
  })

  if (successCount > 0) {
    issues.push(
      createIssue(
        chain,
        'library',
        'info',
        `Send library verified for ${pluralize(successCount, 'chain')}`
      )
    )
  }

  return issues
}

async function verifyReceiveLibraries(
  chain: string,
  client: PublicClient,
  deployment: DeploymentInfo,
  chainInfo: ChainInfo,
  remoteDeployments: RemoteChainDeployment[]
): Promise<VerificationIssue[]> {
  if (remoteDeployments.length === 0) {
    return []
  }

  const receiveLibraryResults = await client.multicall({
    contracts: remoteDeployments.map(({ eid }) => ({
      address: chainInfo.endpointV2,
      abi: layerZeroEndpointAbi,
      functionName: 'getReceiveLibrary' as const,
      args: [deployment.address, eid] as const,
    })),
    allowFailure: true,
  })

  const issues: VerificationIssue[] = []
  let successCount = 0

  receiveLibraryResults.forEach((receiveLibrary, index) => {
    if (receiveLibrary.status !== 'success' || !receiveLibrary.result) {
      return
    }

    const remoteChain = remoteDeployments[index].chain
    const [receiveLibraryAddress, isDefault] = receiveLibrary.result
    let hasError = false

    if (
      receiveLibraryAddress.toLowerCase() !==
      chainInfo.receiveUln302.toLowerCase()
    ) {
      issues.push(
        createIssue(
          chain,
          'library',
          'error',
          `Receive library for ${remoteChain} incorrect: expected ${chainInfo.receiveUln302}, got ${receiveLibraryAddress}`
        )
      )
      hasError = true
    }

    if (isDefault) {
      issues.push(
        createIssue(
          chain,
          'library',
          'error',
          `Receive library for ${remoteChain} is using default library, should be custom`
        )
      )
      hasError = true
    }

    if (!hasError) {
      successCount++
    }
  })

  if (successCount > 0) {
    issues.push(
      createIssue(
        chain,
        'library',
        'info',
        `Receive library verified for ${pluralize(successCount, 'chain')}`
      )
    )
  }

  return issues
}

async function verifyUlnConfigs(
  chain: string,
  client: PublicClient,
  deployment: DeploymentInfo,
  chainInfo: ChainInfo,
  remoteDeployments: RemoteChainDeployment[],
  securityConfig: SecurityConfig,
  expectedDVNs: Address[],
  direction: 'send' | 'receive'
): Promise<VerificationIssue[]> {
  if (remoteDeployments.length === 0) {
    return []
  }

  const libraryAddress =
    direction === 'send' ? chainInfo.sendUln302 : chainInfo.receiveUln302
  const localConfirmations =
    direction === 'send'
      ? getExpectedConfirmations(securityConfig, chain)
      : null
  const directionLabel = direction === 'send' ? 'Send' : 'Receive'

  const configResults = await client.multicall({
    contracts: remoteDeployments.map(({ eid }) => ({
      address: chainInfo.endpointV2,
      abi: layerZeroEndpointAbi,
      functionName: 'getConfig' as const,
      args: [deployment.address, libraryAddress, eid, CONFIG_TYPE_ULN] as const,
    })),
    allowFailure: true,
  })

  const issues: VerificationIssue[] = []
  let successCount = 0

  configResults.forEach((configResult, index) => {
    if (configResult.status !== 'success' || !configResult.result) {
      return
    }

    const remoteChain = remoteDeployments[index].chain
    const expectedConfirmations =
      localConfirmations ??
      getExpectedConfirmations(securityConfig, remoteChain)
    const configIssues = verifyUlnConfig(
      chain,
      remoteChain,
      configResult.result,
      expectedConfirmations,
      expectedDVNs,
      direction
    )

    if (configIssues.length === 0) {
      successCount++
      return
    }

    issues.push(...configIssues)
  })

  if (successCount > 0) {
    issues.push(
      createIssue(
        chain,
        'confirmations',
        'info',
        `${directionLabel} config verified for ${pluralize(
          successCount,
          'chain'
        )} (DVNs & confirmations)`
      )
    )
  }

  return issues
}

function buildTokenVerificationPlan(
  chain: string,
  context: VerificationContext,
  remoteChainNames: Set<string>
): TokenVerificationPlan {
  const tokenTypeRequests: TokenTypeRequest[] = []
  const remoteAddressRequests: RemoteAddressRequest[] = []
  let skippedGMX = false
  const skippedSingleChainTokens: string[] = []

  for (const [tokenSymbol, chains] of Object.entries(context.tokensConfig)) {
    if (chains[chain] === undefined) {
      continue
    }

    if (tokenSymbol === 'GMX') {
      skippedGMX = true
      continue
    }

    if (Object.keys(chains).length <= 1) {
      skippedSingleChainTokens.push(tokenSymbol)
      continue
    }

    const tokenAddr = chains[chain].tokenAddress
    tokenTypeRequests.push({ tokenSymbol, tokenAddr })

    for (const remoteChain of context.allChains) {
      if (remoteChain === chain) {
        continue
      }

      if (!chains[remoteChain] || !remoteChainNames.has(remoteChain)) {
        continue
      }

      remoteAddressRequests.push({
        tokenSymbol,
        remoteChain,
        remoteEid: context.chainsConfig[remoteChain].eid,
        expectedAddr: chains[remoteChain].tokenAddress,
        tokenAddr,
      })
    }
  }

  return {
    tokenTypeRequests,
    remoteAddressRequests,
    skippedGMX,
    skippedSingleChainTokens,
  }
}

async function verifyTokenTypes(
  chain: string,
  client: PublicClient,
  deployment: DeploymentInfo,
  requests: TokenTypeRequest[]
): Promise<CountedIssues> {
  if (requests.length === 0) {
    return { successCount: 0, issues: [] }
  }

  const tokenTypeResults = await client.multicall({
    contracts: requests.map(({ tokenAddr }) => ({
      address: deployment.address,
      abi: synapseBridgeAdapterAbi,
      functionName: 'getTokenType' as const,
      args: [tokenAddr] as const,
    })),
    allowFailure: true,
  })

  const issues: VerificationIssue[] = []
  let successCount = 0

  tokenTypeResults.forEach((tokenType, index) => {
    const request = requests[index]

    if (tokenType.status !== 'success' || tokenType.result === undefined) {
      issues.push(
        createIssue(
          chain,
          'token',
          'error',
          `Failed to check token type for ${request.tokenSymbol}`
        )
      )
      return
    }

    if (tokenType.result === 0) {
      issues.push(
        createIssue(
          chain,
          'token',
          'error',
          `Token ${request.tokenSymbol} (${request.tokenAddr}) not added`
        )
      )
      return
    }

    successCount++
  })

  return { successCount, issues }
}

async function verifyRemoteAddresses(
  chain: string,
  client: PublicClient,
  deployment: DeploymentInfo,
  requests: RemoteAddressRequest[]
): Promise<CountedIssues> {
  if (requests.length === 0) {
    return { successCount: 0, issues: [] }
  }

  const remoteAddressResults = await client.multicall({
    contracts: requests.map(({ remoteEid, tokenAddr }) => ({
      address: deployment.address,
      abi: synapseBridgeAdapterAbi,
      functionName: 'getRemoteAddress' as const,
      args: [remoteEid, tokenAddr] as const,
    })),
    allowFailure: true,
  })

  const issues: VerificationIssue[] = []
  let successCount = 0

  remoteAddressResults.forEach((remoteAddress, index) => {
    const request = requests[index]

    if (remoteAddress.status !== 'success' || !remoteAddress.result) {
      issues.push(
        createIssue(
          chain,
          'token',
          'error',
          `Failed to check token ${request.tokenSymbol} remote address for ${request.remoteChain}`
        )
      )
      return
    }

    if (
      remoteAddress.result.toLowerCase() === request.expectedAddr.toLowerCase()
    ) {
      successCount++
      return
    }

    if (remoteAddress.result === ZERO_ADDRESS) {
      issues.push(
        createIssue(
          chain,
          'token',
          'error',
          `Token ${request.tokenSymbol} remote address for ${request.remoteChain} not set`
        )
      )
      return
    }

    issues.push(
      createIssue(
        chain,
        'token',
        'error',
        `Token ${request.tokenSymbol} remote address for ${request.remoteChain} mismatch: expected ${request.expectedAddr}, got ${remoteAddress.result}`
      )
    )
  })

  return { successCount, issues }
}

function buildTokenInfoIssues(
  chain: string,
  plan: TokenVerificationPlan,
  tokenTypeSuccessCount: number,
  remoteAddressSuccessCount: number
): VerificationIssue[] {
  const issues: VerificationIssue[] = []

  if (tokenTypeSuccessCount > 0 || remoteAddressSuccessCount > 0) {
    issues.push(
      createIssue(
        chain,
        'token',
        'info',
        `Verified ${pluralize(
          tokenTypeSuccessCount,
          'token'
        )} with ${remoteAddressSuccessCount} remote address${
          remoteAddressSuccessCount === 1 ? '' : 'es'
        }`
      )
    )
  }

  if (plan.skippedGMX) {
    issues.push(
      createIssue(chain, 'token', 'info', 'Skipped GMX (not supported)')
    )
  }

  if (plan.skippedSingleChainTokens.length > 0) {
    issues.push(
      createIssue(
        chain,
        'token',
        'info',
        `Skipped ${pluralize(
          plan.skippedSingleChainTokens.length,
          'single-chain token'
        )}: ${plan.skippedSingleChainTokens.join(', ')}`
      )
    )
  }

  return issues
}

async function verifyTokensForChain(
  chain: string,
  client: PublicClient,
  deployment: DeploymentInfo,
  context: VerificationContext,
  remoteDeployments: RemoteChainDeployment[]
): Promise<VerificationIssue[]> {
  const plan = buildTokenVerificationPlan(
    chain,
    context,
    new Set(remoteDeployments.map((remoteDeployment) => remoteDeployment.chain))
  )
  const tokenTypes = await verifyTokenTypes(
    chain,
    client,
    deployment,
    plan.tokenTypeRequests
  )
  const remoteAddresses = await verifyRemoteAddresses(
    chain,
    client,
    deployment,
    plan.remoteAddressRequests
  )

  return [
    ...tokenTypes.issues,
    ...remoteAddresses.issues,
    ...buildTokenInfoIssues(
      chain,
      plan,
      tokenTypes.successCount,
      remoteAddresses.successCount
    ),
  ]
}

// Function to verify a single chain
async function verifyChain(
  chain: string,
  chainId: number,
  deployment: DeploymentInfo,
  context: VerificationContext
): Promise<ChainVerificationResult> {
  const result: ChainVerificationResult = {
    chain,
    chainId,
    owner: null,
    bridge: null,
    issues: [],
  }

  try {
    const client = createChainClient(chainId)
    const chainInfo = context.chainsConfig[chain]
    const remoteDeployments = getRemoteDeployments(chain, context)
    const basicInfo = await verifyBasicInfo(
      chain,
      client,
      deployment,
      chainInfo
    )
    const expectedDVNs = getExpectedDVNs(chain, context)

    result.owner = basicInfo.owner
    result.bridge = basicInfo.bridge
    result.issues.push(...basicInfo.issues)
    result.issues.push(
      ...(await verifyPeerConfiguration(
        chain,
        client,
        deployment,
        remoteDeployments
      ))
    )
    result.issues.push(
      ...(await verifySendLibraries(
        chain,
        client,
        deployment,
        chainInfo,
        remoteDeployments
      ))
    )
    result.issues.push(
      ...(await verifyReceiveLibraries(
        chain,
        client,
        deployment,
        chainInfo,
        remoteDeployments
      ))
    )
    result.issues.push(
      ...(await verifyUlnConfigs(
        chain,
        client,
        deployment,
        chainInfo,
        remoteDeployments,
        context.securityConfig,
        expectedDVNs,
        'send'
      ))
    )
    result.issues.push(
      ...(await verifyUlnConfigs(
        chain,
        client,
        deployment,
        chainInfo,
        remoteDeployments,
        context.securityConfig,
        expectedDVNs,
        'receive'
      ))
    )
    result.issues.push(
      ...(await verifyTokensForChain(
        chain,
        client,
        deployment,
        context,
        remoteDeployments
      ))
    )
  } catch (error) {
    result.issues.push(
      createIssue(chain, 'error', 'error', `RPC error: ${formatError(error)}`)
    )
  }

  return result
}

// Function to print results
function printResults(results: ChainVerificationResult[]) {
  console.log(
    `\n${colors.bold}${colors.cyan}SynapseBridgeAdapter Verification Results${colors.reset}\n`
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
    `${colors.bold}Starting SynapseBridgeAdapter verification...${colors.reset}\n`
  )

  // Load configurations
  const chainIds = loadChainIds()
  console.log(
    `Loaded ${Object.keys(chainIds).length} chain IDs from deployments`
  )

  const chainsConfigPath = path.join(
    __dirname,
    '../../configs/global/chains.json'
  )
  const chainsConfig: ChainsConfig = JSON.parse(
    fs.readFileSync(chainsConfigPath, 'utf8')
  )

  const tokensConfigPath = path.join(
    __dirname,
    '../../configs/global/tokens.json'
  )
  const tokensConfig: TokenConfig = JSON.parse(
    fs.readFileSync(tokensConfigPath, 'utf8')
  )

  const dvnsConfigPath = path.join(__dirname, '../../configs/global/dvns.json')
  const dvnsConfig: DVNConfig = JSON.parse(
    fs.readFileSync(dvnsConfigPath, 'utf8')
  )

  const securityConfigPath = path.join(
    __dirname,
    '../../configs/global/security.json'
  )
  const securityConfig: SecurityConfig = JSON.parse(
    fs.readFileSync(securityConfigPath, 'utf8')
  )
  const context: VerificationContext = {
    chainsConfig,
    tokensConfig,
    dvnsConfig,
    securityConfig,
    allChains: Object.keys(chainIds).filter(
      (chain) => chainsConfig[chain] !== undefined
    ),
  }

  console.log(
    `Loaded configuration for ${Object.keys(chainsConfig).length} chains`
  )
  console.log(
    `Security config: block confirmations for ${
      Object.keys(securityConfig.blockConfirmations).length
    } chains, DVNs: ${securityConfig.DVNs.join(', ')}\n`
  )

  // Filter by command line arguments if provided
  const requestedChains = process.argv.slice(2)
  const chainsToVerify =
    requestedChains.length > 0
      ? context.allChains.filter((chain) => requestedChains.includes(chain))
      : context.allChains

  if (requestedChains.length > 0 && chainsToVerify.length === 0) {
    console.error(
      `${
        colors.red
      }Error: None of the requested chains found: ${requestedChains.join(
        ', '
      )}${colors.reset}`
    )
    console.error(`Available chains: ${context.allChains.join(', ')}`)
    process.exit(1)
  }

  console.log(`Verifying ${chainsToVerify.length} deployed chains...\n`)

  // Verify each chain
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

    const result = await verifyChain(chain, chainId, deployment, context)

    results.push(result)
  }

  // Print results
  printResults(results)

  // Exit with appropriate code
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
