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

// LayerZero config types
const CONFIG_TYPE_ULN = 2

// Verification constants
const TARGET_CONFIRMATION_TIME_SECONDS = 3600 // 1 hour
const CONFIRMATION_TIME_TOLERANCE_PERCENT = 0.05 // ±5%
const MS_TO_SECONDS = 1000

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
  confirmationTimeSeconds: number
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

// Multicall3 address (universal deployment)
const MULTICALL3_ADDRESS = '0xcA11bde05977b3631167028862bE2a173976CA11'

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
function sortAddresses(addresses: Address[]): Address[] {
  return addresses
    .map((addr) => ({ addr, num: BigInt(addr) }))
    .sort((a, b) => (a.num < b.num ? -1 : a.num > b.num ? 1 : 0))
    .map((item) => item.addr)
}

// Utility helper functions
function pluralize(count: number, word: string): string {
  return `${count} ${word}${count === 1 ? '' : 's'}`
}

function blockTimeToSeconds(blockTimeMs: number): number {
  return blockTimeMs / MS_TO_SECONDS
}

// Validation helper functions
function validateConfirmationTime(
  confirmations: bigint,
  blockTimeMs: number
): { isValid: boolean; actualSeconds: number } {
  const actualTimeSeconds =
    Number(confirmations) * blockTimeToSeconds(blockTimeMs)
  const tolerance =
    TARGET_CONFIRMATION_TIME_SECONDS * CONFIRMATION_TIME_TOLERANCE_PERCENT
  const isValid =
    Math.abs(actualTimeSeconds - TARGET_CONFIRMATION_TIME_SECONDS) <= tolerance
  return { isValid, actualSeconds: actualTimeSeconds }
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
  blockTimeMs: number,
  expectedDVNs: Address[],
  configDirection: 'send' | 'receive'
): VerificationIssue[] {
  const issues: VerificationIssue[] = []

  try {
    const decoded = decodeUlnConfig(configBytes)

    // Validate confirmation time
    const { isValid, actualSeconds } = validateConfirmationTime(
      decoded.confirmations,
      blockTimeMs
    )
    if (!isValid) {
      issues.push({
        chain,
        category: 'confirmations',
        severity: 'warning',
        message: `${
          configDirection === 'send' ? 'Send' : 'Receive'
        } confirmations for ${remoteChain}: ${
          decoded.confirmations
        } blocks × ${blockTimeToSeconds(blockTimeMs).toFixed(
          1
        )}s = ${actualSeconds.toFixed(
          0
        )}s (expected ~${TARGET_CONFIRMATION_TIME_SECONDS}s ±${
          CONFIRMATION_TIME_TOLERANCE_PERCENT * 100
        }%)`,
      })
    }

    // Validate DVNs
    if (!areDVNsEqual(decoded.requiredDVNs, expectedDVNs)) {
      const sortedActual = sortAddresses(decoded.requiredDVNs)
      const sortedExpected = sortAddresses(expectedDVNs)
      issues.push({
        chain,
        category: 'dvn',
        severity: 'error',
        message: `${
          configDirection === 'send' ? 'Send' : 'Receive'
        } DVNs for ${remoteChain} incorrect: expected [${sortedExpected.join(
          ', '
        )}], got [${sortedActual.join(', ')}]`,
      })
    }
  } catch (error) {
    issues.push({
      chain,
      category: 'error',
      severity: 'error',
      message: `Failed to decode ${configDirection} config for ${remoteChain}: ${
        error instanceof Error ? error.message : String(error)
      }`,
    })
  }

  return issues
}

// Function to verify a single chain
async function verifyChain(
  chain: string,
  chainId: number,
  deployment: DeploymentInfo,
  chainsConfig: ChainsConfig,
  tokensConfig: TokenConfig,
  dvnsConfig: DVNConfig,
  securityConfig: SecurityConfig,
  allChains: string[]
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
    const chainInfo = chainsConfig[chain]

    // Prepare multicall for basic info
    const basicInfoCalls = [
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
    ] as const

    const basicResults = await client.multicall({
      contracts: basicInfoCalls,
      allowFailure: true,
    })

    if (basicResults[0].status === 'success' && basicResults[0].result) {
      result.owner = basicResults[0].result
    }

    if (basicResults[1].status === 'success' && basicResults[1].result) {
      result.bridge = basicResults[1].result
      if (result.bridge === '0x0000000000000000000000000000000000000000') {
        result.issues.push({
          chain,
          category: 'bridge',
          severity: 'error',
          message: 'Bridge address not set',
        })
      } else if (
        result.bridge.toLowerCase() !== chainInfo.synapseBridge.toLowerCase()
      ) {
        result.issues.push({
          chain,
          category: 'bridge',
          severity: 'warning',
          message: `Bridge address mismatch: expected ${chainInfo.synapseBridge}, got ${result.bridge}`,
        })
      }
    } else {
      result.issues.push({
        chain,
        category: 'error',
        severity: 'error',
        message: 'Failed to fetch bridge address',
      })
    }

    // Verify LayerZero endpoint
    if (basicResults[2].status === 'success' && basicResults[2].result) {
      const endpoint = basicResults[2].result
      if (endpoint.toLowerCase() !== chainInfo.endpointV2.toLowerCase()) {
        result.issues.push({
          chain,
          category: 'error',
          severity: 'error',
          message: `Endpoint mismatch: expected ${chainInfo.endpointV2}, got ${endpoint}`,
        })
      }
    }

    // Verify peers for all other chains
    const peerChains: string[] = []
    const peerEids: number[] = []
    const peerExpected: string[] = []
    const peerCalls = []

    for (const remoteChain of allChains) {
      if (remoteChain === chain) continue
      const remoteDeployment = loadDeployment(remoteChain)
      if (!remoteDeployment) continue
      const remoteEid = chainsConfig[remoteChain].eid
      const expectedPeer = `0x${remoteDeployment.address
        .slice(2)
        .padStart(64, '0')}`

      peerChains.push(remoteChain)
      peerEids.push(remoteEid)
      peerExpected.push(expectedPeer)
      peerCalls.push({
        address: deployment.address,
        abi: synapseBridgeAdapterAbi,
        functionName: 'peers' as const,
        args: [remoteEid] as const,
      })
    }

    const peerResults = await client.multicall({
      contracts: peerCalls,
      allowFailure: true,
    })

    let peerSuccessCount = 0
    peerResults.forEach((peer, i) => {
      const remoteChain = peerChains[i]
      const remoteEid = peerEids[i]
      const expectedPeer = peerExpected[i]

      if (peer.status === 'success' && peer.result) {
        if (peer.result.toLowerCase() !== expectedPeer.toLowerCase()) {
          if (peer.result === '0x' + '0'.repeat(64)) {
            result.issues.push({
              chain,
              category: 'peer',
              severity: 'error',
              message: `Peer not set for ${remoteChain} (eid: ${remoteEid})`,
            })
          } else {
            result.issues.push({
              chain,
              category: 'peer',
              severity: 'error',
              message: `Peer mismatch for ${remoteChain}: expected ${expectedPeer}, got ${peer.result}`,
            })
          }
        } else {
          peerSuccessCount++
        }
      } else {
        result.issues.push({
          chain,
          category: 'error',
          severity: 'error',
          message: `Failed to fetch peer for ${remoteChain}`,
        })
      }
    })

    if (peerSuccessCount > 0) {
      result.issues.push({
        chain,
        category: 'peer',
        severity: 'info',
        message: `Verified ${pluralize(peerSuccessCount, 'peer')}`,
      })
    }

    // Prepare expected DVNs (sorted)
    const expectedDVNs = sortAddresses(
      securityConfig.DVNs.map((dvnName) => dvnsConfig[chain][dvnName])
    )

    // Build metadata arrays for library and config verification
    const libChains: string[] = []
    const libEids: number[] = []

    for (const remoteChain of allChains) {
      if (remoteChain === chain) continue
      if (!loadDeployment(remoteChain)) continue
      const remoteEid = chainsConfig[remoteChain].eid

      libChains.push(remoteChain)
      libEids.push(remoteEid)
    }

    // Verify send libraries
    const sendLibraryCalls = libEids.map((remoteEid) => ({
      address: chainInfo.endpointV2,
      abi: layerZeroEndpointAbi,
      functionName: 'getSendLibrary' as const,
      args: [deployment.address, remoteEid] as const,
    }))

    const sendLibraryResults = await client.multicall({
      contracts: sendLibraryCalls,
      allowFailure: true,
    })

    let sendLibSuccessCount = 0
    sendLibraryResults.forEach((sendLib, i) => {
      const remoteChain = libChains[i]
      if (sendLib.status === 'success' && sendLib.result) {
        if (
          sendLib.result.toLowerCase() !== chainInfo.sendUln302.toLowerCase()
        ) {
          result.issues.push({
            chain,
            category: 'library',
            severity: 'error',
            message: `Send library for ${remoteChain} incorrect: expected ${chainInfo.sendUln302}, got ${sendLib}`,
          })
        } else {
          sendLibSuccessCount++
        }
      }
    })

    if (sendLibSuccessCount > 0) {
      result.issues.push({
        chain,
        category: 'library',
        severity: 'info',
        message: `Send library verified for ${pluralize(
          sendLibSuccessCount,
          'chain'
        )}`,
      })
    }

    // Verify receive libraries
    const receiveLibraryCalls = libEids.map((remoteEid) => ({
      address: chainInfo.endpointV2,
      abi: layerZeroEndpointAbi,
      functionName: 'getReceiveLibrary' as const,
      args: [deployment.address, remoteEid] as const,
    }))

    const receiveLibraryResults = await client.multicall({
      contracts: receiveLibraryCalls,
      allowFailure: true,
    })

    let receiveLibSuccessCount = 0
    receiveLibraryResults.forEach((receiveLib, i) => {
      const remoteChain = libChains[i]
      if (receiveLib.status === 'success' && receiveLib.result) {
        const [receiveLibAddr, isDefault] = receiveLib.result
        let hasError = false
        if (
          receiveLibAddr.toLowerCase() !== chainInfo.receiveUln302.toLowerCase()
        ) {
          result.issues.push({
            chain,
            category: 'library',
            severity: 'error',
            message: `Receive library for ${remoteChain} incorrect: expected ${chainInfo.receiveUln302}, got ${receiveLibAddr}`,
          })
          hasError = true
        }
        if (isDefault) {
          result.issues.push({
            chain,
            category: 'library',
            severity: 'error',
            message: `Receive library for ${remoteChain} is using default library, should be custom`,
          })
          hasError = true
        }
        if (!hasError) {
          receiveLibSuccessCount++
        }
      }
    })

    if (receiveLibSuccessCount > 0) {
      result.issues.push({
        chain,
        category: 'library',
        severity: 'info',
        message: `Receive library verified for ${pluralize(
          receiveLibSuccessCount,
          'chain'
        )}`,
      })
    }

    // Verify send configs
    const sendConfigCalls = libEids.map((remoteEid) => ({
      address: chainInfo.endpointV2,
      abi: layerZeroEndpointAbi,
      functionName: 'getConfig' as const,
      args: [
        deployment.address,
        chainInfo.sendUln302,
        remoteEid,
        CONFIG_TYPE_ULN,
      ] as const,
    }))

    const sendConfigResults = await client.multicall({
      contracts: sendConfigCalls,
      allowFailure: true,
    })

    let sendConfigSuccessCount = 0
    sendConfigResults.forEach((sendConfig, i) => {
      const remoteChain = libChains[i]
      if (sendConfig.status === 'success' && sendConfig.result) {
        const configIssues = verifyUlnConfig(
          chain,
          remoteChain,
          sendConfig.result,
          chainInfo.blockTime,
          expectedDVNs,
          'send'
        )
        if (configIssues.length === 0) {
          sendConfigSuccessCount++
        } else {
          result.issues.push(...configIssues)
        }
      }
    })

    if (sendConfigSuccessCount > 0) {
      result.issues.push({
        chain,
        category: 'confirmations',
        severity: 'info',
        message: `Send config verified for ${pluralize(
          sendConfigSuccessCount,
          'chain'
        )} (DVNs & confirmations)`,
      })
    }

    // Verify receive configs
    const receiveConfigCalls = libEids.map((remoteEid) => ({
      address: chainInfo.endpointV2,
      abi: layerZeroEndpointAbi,
      functionName: 'getConfig' as const,
      args: [
        deployment.address,
        chainInfo.receiveUln302,
        remoteEid,
        CONFIG_TYPE_ULN,
      ] as const,
    }))

    const receiveConfigResults = await client.multicall({
      contracts: receiveConfigCalls,
      allowFailure: true,
    })

    let receiveConfigSuccessCount = 0
    receiveConfigResults.forEach((receiveConfig, i) => {
      const remoteChain = libChains[i]
      const remoteBlockTime = chainsConfig[remoteChain].blockTime
      if (receiveConfig.status === 'success' && receiveConfig.result) {
        const configIssues = verifyUlnConfig(
          chain,
          remoteChain,
          receiveConfig.result,
          remoteBlockTime,
          expectedDVNs,
          'receive'
        )
        if (configIssues.length === 0) {
          receiveConfigSuccessCount++
        } else {
          result.issues.push(...configIssues)
        }
      }
    })

    if (receiveConfigSuccessCount > 0) {
      result.issues.push({
        chain,
        category: 'confirmations',
        severity: 'info',
        message: `Receive config verified for ${pluralize(
          receiveConfigSuccessCount,
          'chain'
        )} (DVNs & confirmations)`,
      })
    }

    // Verify tokens
    const tokensOnChain = Object.entries(tokensConfig).filter(
      ([_, chains]) => chains[chain] !== undefined
    )

    // Build metadata arrays for token type verification
    const tokenTypeSymbols: string[] = []
    const tokenTypeAddrs: Address[] = []
    const tokenTypeCalls = []

    // Build metadata arrays for remote address verification
    const remoteAddrSymbols: string[] = []
    const remoteAddrChains: string[] = []
    const remoteAddrExpected: Address[] = []
    const remoteAddrCalls = []

    // Track skipped tokens
    let skippedGMX = false
    const skippedSingleChainTokens: string[] = []

    for (const [tokenSymbol, chains] of tokensOnChain) {
      // Skip GMX token (not supported by new adapter)
      if (tokenSymbol === 'GMX') {
        skippedGMX = true
        continue
      }

      // Skip tokens that only exist on a single chain (no remote addresses to verify)
      const chainCount = Object.keys(chains).length
      if (chainCount <= 1) {
        skippedSingleChainTokens.push(tokenSymbol)
        continue
      }

      const tokenAddr = chains[chain].tokenAddress

      // Add token type check
      tokenTypeSymbols.push(tokenSymbol)
      tokenTypeAddrs.push(tokenAddr)
      tokenTypeCalls.push({
        address: deployment.address,
        abi: synapseBridgeAdapterAbi,
        functionName: 'getTokenType' as const,
        args: [tokenAddr] as const,
      })

      // Add remote address checks for all chains where this token exists
      for (const remoteChain of allChains) {
        if (remoteChain === chain) continue
        if (!chains[remoteChain]) continue
        if (!loadDeployment(remoteChain)) continue

        const remoteEid = chainsConfig[remoteChain].eid
        remoteAddrSymbols.push(tokenSymbol)
        remoteAddrChains.push(remoteChain)
        remoteAddrExpected.push(chains[remoteChain].tokenAddress)
        remoteAddrCalls.push({
          address: deployment.address,
          abi: synapseBridgeAdapterAbi,
          functionName: 'getRemoteAddress' as const,
          args: [remoteEid, tokenAddr] as const,
        })
      }
    }

    // Verify token types
    const tokenTypeResults = await client.multicall({
      contracts: tokenTypeCalls,
      allowFailure: true,
    })

    let tokenTypeSuccessCount = 0
    tokenTypeResults.forEach((tokenType, i) => {
      const tokenSymbol = tokenTypeSymbols[i]
      const tokenAddr = tokenTypeAddrs[i]

      if (tokenType.status === 'success' && tokenType.result !== undefined) {
        if (tokenType.result === 0) {
          result.issues.push({
            chain,
            category: 'token',
            severity: 'error',
            message: `Token ${tokenSymbol} (${tokenAddr}) not added`,
          })
        } else {
          tokenTypeSuccessCount++
        }
      } else {
        result.issues.push({
          chain,
          category: 'token',
          severity: 'error',
          message: `Failed to check token type for ${tokenSymbol}`,
        })
      }
    })

    // Verify remote addresses
    const remoteAddrResults = await client.multicall({
      contracts: remoteAddrCalls,
      allowFailure: true,
    })

    let remoteAddrSuccessCount = 0
    remoteAddrResults.forEach((remoteAddr, i) => {
      const tokenSymbol = remoteAddrSymbols[i]
      const remoteChain = remoteAddrChains[i]
      const expectedAddr = remoteAddrExpected[i]

      if (remoteAddr.status === 'success' && remoteAddr.result) {
        if (
          remoteAddr.result === '0x0000000000000000000000000000000000000000'
        ) {
          result.issues.push({
            chain,
            category: 'token',
            severity: 'error',
            message: `Token ${tokenSymbol} remote address for ${remoteChain} not set`,
          })
        } else if (
          remoteAddr.result.toLowerCase() !== expectedAddr.toLowerCase()
        ) {
          result.issues.push({
            chain,
            category: 'token',
            severity: 'error',
            message: `Token ${tokenSymbol} remote address for ${remoteChain} mismatch: expected ${expectedAddr}, got ${remoteAddr.result}`,
          })
        } else {
          remoteAddrSuccessCount++
        }
      } else {
        result.issues.push({
          chain,
          category: 'token',
          severity: 'error',
          message: `Failed to check token ${tokenSymbol} remote address for ${remoteChain}`,
        })
      }
    })

    if (tokenTypeSuccessCount > 0 || remoteAddrSuccessCount > 0) {
      result.issues.push({
        chain,
        category: 'token',
        severity: 'info',
        message: `Verified ${pluralize(
          tokenTypeSuccessCount,
          'token'
        )} with ${remoteAddrSuccessCount} remote address${
          remoteAddrSuccessCount === 1 ? '' : 'es'
        }`,
      })
    }

    // Add info logs for skipped tokens
    if (skippedGMX) {
      result.issues.push({
        chain,
        category: 'token',
        severity: 'info',
        message: `Skipped GMX (not supported)`,
      })
    }

    if (skippedSingleChainTokens.length > 0) {
      result.issues.push({
        chain,
        category: 'token',
        severity: 'info',
        message: `Skipped ${pluralize(
          skippedSingleChainTokens.length,
          'single-chain token'
        )}: ${skippedSingleChainTokens.join(', ')}`,
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

// Function to print results
function printResults(results: ChainVerificationResult[]) {
  // Maps for cleaner output formatting
  const categoryLabels: Record<string, string> = {
    peer: 'Peer Configuration',
    library: 'Library Configuration',
    dvn: 'DVN Configuration',
    confirmations: 'Block Confirmations',
    token: 'Token Configuration',
    bridge: 'Bridge Configuration',
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
    `\n${colors.bold}${colors.cyan}SynapseBridgeAdapter Verification Results${colors.reset}\n`
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

  // Print detailed results for each chain
  for (const result of results) {
    console.log(
      `\n${colors.bold}${colors.cyan}${result.chain} (Chain ID: ${result.chainId})${colors.reset}`
    )
    console.log(
      `  Owner: ${result.owner || colors.red + 'Unknown' + colors.reset}`
    )
    console.log(
      `  Bridge: ${result.bridge || colors.red + 'Not set' + colors.reset}`
    )

    const errorWarningIssues = result.issues.filter(
      (i) => i.severity === 'error' || i.severity === 'warning'
    )
    const infoIssues = result.issues.filter((i) => i.severity === 'info')

    if (errorWarningIssues.length === 0) {
      console.log(`  ${colors.green}✓ All checks passed${colors.reset}`)
    } else {
      console.log(
        `  ${colors.yellow}Issues found: ${errorWarningIssues.length}${colors.reset}`
      )
    }

    // Group all issues by category (including info)
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

  // Print summary
  console.log(`\n${colors.bold}${colors.cyan}Summary${colors.reset}`)
  if (totalIssues === 0) {
    console.log(
      `${colors.green}✓ All SynapseBridgeAdapter contracts are correctly configured!${colors.reset}\n`
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

  console.log(
    `Loaded configuration for ${Object.keys(chainsConfig).length} chains`
  )
  console.log(
    `Security config: ${
      securityConfig.confirmationTimeSeconds
    }s confirmation time, DVNs: ${securityConfig.DVNs.join(', ')}\n`
  )

  // Get all chains with deployments
  const allChains = Object.keys(chainIds).filter(
    (chain) => chainsConfig[chain] !== undefined
  )

  // Filter by command line arguments if provided
  const requestedChains = process.argv.slice(2)
  const chainsToVerify =
    requestedChains.length > 0
      ? allChains.filter((chain) => requestedChains.includes(chain))
      : allChains

  if (requestedChains.length > 0 && chainsToVerify.length === 0) {
    console.error(
      `${
        colors.red
      }Error: None of the requested chains found: ${requestedChains.join(
        ', '
      )}${colors.reset}`
    )
    console.error(`Available chains: ${allChains.join(', ')}`)
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

    const result = await verifyChain(
      chain,
      chainId,
      deployment,
      chainsConfig,
      tokensConfig,
      dvnsConfig,
      securityConfig,
      allChains
    )

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

// Run the script
main().catch((error) => {
  console.error(`${colors.red}Error: ${error}${colors.reset}`)
  process.exit(1)
})
