import * as fs from 'fs'
import * as path from 'path'
import {
  createPublicClient,
  defineChain,
  getAddress,
  http,
  type Address,
  type PublicClient,
} from 'viem'

type DeploymentArtifact = {
  address?: string
}

type MultisigConfig = Record<string, Address>
type WalletType =
  | 'gnosis-safe'
  | 'legacy-multisig'
  | 'legacy-multisig-with-daily-limit'

type MultisigDetails = {
  signers: readonly Address[]
  threshold: bigint
  version?: string
  walletType: WalletType
}

const SYNAPSE_CONTRACTS_REPO = 'synapsecns/synapse-contracts'
const SYNAPSE_CONTRACTS_REF = process.env.SYNAPSE_CONTRACTS_REF || 'master'
const RAW_BASE_URL = `https://raw.githubusercontent.com/${SYNAPSE_CONTRACTS_REPO}/${SYNAPSE_CONTRACTS_REF}`

const gnosisSafeAbi = [
  {
    name: 'VERSION',
    type: 'function',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ type: 'string' }],
  },
  {
    name: 'getOwners',
    type: 'function',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ type: 'address[]' }],
  },
  {
    name: 'getThreshold',
    type: 'function',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ type: 'uint256' }],
  },
] as const

const legacyMultisigAbi = [
  {
    name: 'getOwners',
    type: 'function',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ type: 'address[]' }],
  },
  {
    name: 'required',
    type: 'function',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ type: 'uint256' }],
  },
  {
    name: 'dailyLimit',
    type: 'function',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ type: 'uint256' }],
  },
] as const

// Some deployment directories in synapse-contracts use legacy or non-matching names.
const upstreamChainAliases: Record<string, string[]> = {
  ethereum: ['mainnet'],
  bnb: ['bsc'],
  kaia: ['klatyn', 'klaytn'],
}

function getAdapterDeploymentChains(): string[] {
  const deploymentsPath = path.join(__dirname, '../../deployments')

  if (!fs.existsSync(deploymentsPath)) {
    throw new Error(`Deployments directory not found: ${deploymentsPath}`)
  }

  return fs
    .readdirSync(deploymentsPath)
    .filter((chain) => {
      const chainPath = path.join(deploymentsPath, chain)
      const adapterPath = path.join(chainPath, 'SynapseBridgeAdapter.json')
      return fs.statSync(chainPath).isDirectory() && fs.existsSync(adapterPath)
    })
    .sort()
}

function loadChainIds(): Record<string, number> {
  const deploymentsPath = path.join(__dirname, '../../deployments')
  const chainIds: Record<string, number> = {}

  if (!fs.existsSync(deploymentsPath)) {
    throw new Error(`Deployments directory not found: ${deploymentsPath}`)
  }

  const directories = fs.readdirSync(deploymentsPath)

  for (const dir of directories) {
    const chainIdPath = path.join(deploymentsPath, dir, '.chainId')

    if (!fs.existsSync(chainIdPath)) {
      continue
    }

    const chainId = parseInt(fs.readFileSync(chainIdPath, 'utf8').trim(), 10)
    if (!isNaN(chainId)) {
      chainIds[dir] = chainId
    }
  }

  return chainIds
}

function getUpstreamChainCandidates(chain: string): string[] {
  return [...new Set([...(upstreamChainAliases[chain] || []), chain])]
}

function getRpcUrl(chainId: number): string {
  const baseUrl = process.env.RPC_BASE_URL

  if (!baseUrl) {
    throw new Error(
      'RPC_BASE_URL environment variable is not set. Please set it to your RPC provider base URL in the .env file'
    )
  }

  // Special case for Moonbeam to match the existing verification utility pattern.
  if (chainId === 1284) {
    return 'https://moonbeam.api.pocket.network'
  }

  return `${baseUrl}/${chainId}`
}

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

async function fetchMultisigAddress(
  chain: string
): Promise<{ address: Address; upstreamChain: string }> {
  const candidates = getUpstreamChainCandidates(chain)
  const failures: string[] = []

  for (const upstreamChain of candidates) {
    const url = `${RAW_BASE_URL}/deployments/${upstreamChain}/DevMultisig.json`
    const response = await fetch(url)

    if (response.status === 404) {
      failures.push(`${upstreamChain}:404`)
      continue
    }

    if (!response.ok) {
      throw new Error(
        `Failed to fetch DevMultisig for ${chain} from ${upstreamChain}: ${response.status} ${response.statusText}`
      )
    }

    const deployment = (await response.json()) as DeploymentArtifact

    if (!deployment.address) {
      throw new Error(
        `DevMultisig deployment for ${chain} from ${upstreamChain} is missing an address`
      )
    }

    return {
      address: getAddress(deployment.address),
      upstreamChain,
    }
  }

  throw new Error(
    `Could not find DevMultisig deployment for ${chain}. Tried: ${failures.join(', ')}`
  )
}

function formatWalletType(details: MultisigDetails): string {
  if (details.walletType === 'gnosis-safe') {
    return details.version
      ? `gnosis-safe (version ${details.version})`
      : 'gnosis-safe'
  }

  if (details.walletType === 'legacy-multisig-with-daily-limit') {
    return 'legacy-multisig-with-daily-limit'
  }

  return 'legacy-multisig'
}

function formatError(error: unknown): string {
  if (error instanceof Error) {
    return error.message
  }
  return String(error)
}

async function fetchMultisigDetails(
  client: PublicClient,
  address: Address
): Promise<MultisigDetails> {
  try {
    const [version, signers, threshold] = await Promise.all([
      client.readContract({
        address,
        abi: gnosisSafeAbi,
        functionName: 'VERSION',
      }),
      client.readContract({
        address,
        abi: gnosisSafeAbi,
        functionName: 'getOwners',
      }),
      client.readContract({
        address,
        abi: gnosisSafeAbi,
        functionName: 'getThreshold',
      }),
    ])

    return {
      signers,
      threshold,
      version,
      walletType: 'gnosis-safe',
    }
  } catch (safeError) {
    try {
      const [signers, threshold] = await Promise.all([
        client.readContract({
          address,
          abi: legacyMultisigAbi,
          functionName: 'getOwners',
        }),
        client.readContract({
          address,
          abi: legacyMultisigAbi,
          functionName: 'required',
        }),
      ])

      let walletType: WalletType = 'legacy-multisig'
      try {
        await client.readContract({
          address,
          abi: legacyMultisigAbi,
          functionName: 'dailyLimit',
        })
        walletType = 'legacy-multisig-with-daily-limit'
      } catch {
        // No-op: plain legacy multisig wallets do not expose dailyLimit().
      }

      return {
        signers,
        threshold,
        walletType,
      }
    } catch (legacyError) {
      throw new Error(
        `Could not classify multisig at ${address}. Safe probe failed: ${formatError(
          safeError
        )}. Legacy probe failed: ${formatError(legacyError)}`
      )
    }
  }
}

function sortMultisigConfig(config: MultisigConfig): MultisigConfig {
  const sorted: MultisigConfig = {}

  for (const chain of Object.keys(config).sort()) {
    sorted[chain] = config[chain]
  }

  return sorted
}

async function main() {
  const chains = getAdapterDeploymentChains()
  const chainIds = loadChainIds()

  console.log(`Found ${chains.length} chains with SynapseBridgeAdapter deployments`)

  const entries: Array<readonly [string, Address]> = []

  for (const chain of chains) {
    const chainId = chainIds[chain]
    if (chainId === undefined) {
      throw new Error(`Missing .chainId for ${chain}`)
    }

    const { address, upstreamChain } = await fetchMultisigAddress(chain)
    const client = createChainClient(chainId)
    const details = await fetchMultisigDetails(client, address)

    console.log(`Resolved ${chain} -> ${upstreamChain}: ${address}`)
    console.log(`  Chain ID: ${chainId}`)
    console.log(`  Wallet type: ${formatWalletType(details)}`)
    console.log(`  Threshold: ${details.threshold.toString()}`)
    console.log(`  Owners (${details.signers.length}):`)
    for (const signer of details.signers) {
      console.log(`    - ${signer}`)
    }

    entries.push([chain, address] as const)
  }

  const multisigConfig = sortMultisigConfig(Object.fromEntries(entries))
  const outputPath = path.join(__dirname, '../../configs/global/multisig.json')

  fs.mkdirSync(path.dirname(outputPath), { recursive: true })
  fs.writeFileSync(outputPath, JSON.stringify(multisigConfig, null, 2) + '\n')

  console.log(`Successfully wrote multisig config to ${outputPath}`)
}

main().catch((error) => {
  console.error('Error:', error)
  process.exit(1)
})
