import { SupportedChainId } from '../constants'

export const SBA_MIN_GAS_LIMIT = 200_000
export const SBA_ESTIMATED_TIME_CACHE_TTL = 60 * 60
export const SBA_EXECUTION_BUFFER_SECONDS = 60

export type SynapseBridgeAdapterChainMetadata = {
  adapterAddress: string
  lzEid: number
  originBlockConfirmations: number
}

// Metadata snapshot sourced from:
// - packages/contracts-adapter/deployments/*/SynapseBridgeAdapter.json
// - packages/contracts-adapter/configs/global/chains.json
// - packages/contracts-adapter/configs/global/security.json
// Name normalization required by sdk-router:
// - ethereum -> SupportedChainId.ETH
// - bnb -> SupportedChainId.BSC
// - kaia -> SupportedChainId.KLAYTN
// Current phase-1 deployments use the same SBA address on all supported chains.
const SBA_DEPLOYMENT_ADDRESS = '0x5Ba000Bb06230E0582e111F08e1f2F2F200005BA'

export const SBA_CHAIN_METADATA: Partial<
  Record<SupportedChainId, SynapseBridgeAdapterChainMetadata>
> = {
  [SupportedChainId.ETH]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30101,
    originBlockConfirmations: 64,
  },
  [SupportedChainId.OPTIMISM]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30111,
    originBlockConfirmations: 100,
  },
  [SupportedChainId.BSC]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30102,
    originBlockConfirmations: 100,
  },
  [SupportedChainId.POLYGON]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30109,
    originBlockConfirmations: 200,
  },
  [SupportedChainId.FANTOM]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30112,
    originBlockConfirmations: 100,
  },
  [SupportedChainId.METIS]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30151,
    originBlockConfirmations: 100,
  },
  [SupportedChainId.CANTO]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30159,
    originBlockConfirmations: 100,
  },
  [SupportedChainId.KLAYTN]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30150,
    originBlockConfirmations: 100,
  },
  [SupportedChainId.BASE]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30184,
    originBlockConfirmations: 100,
  },
  [SupportedChainId.ARBITRUM]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30110,
    originBlockConfirmations: 100,
  },
  [SupportedChainId.AVALANCHE]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30106,
    originBlockConfirmations: 100,
  },
  [SupportedChainId.DFK]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30115,
    originBlockConfirmations: 100,
  },
  [SupportedChainId.HARMONY]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30116,
    originBlockConfirmations: 100,
  },
  [SupportedChainId.BLAST]: {
    adapterAddress: SBA_DEPLOYMENT_ADDRESS,
    lzEid: 30243,
    originBlockConfirmations: 100,
  },
}

export const getSbaChainMetadata = (
  chainId: number
): SynapseBridgeAdapterChainMetadata | undefined => {
  return SBA_CHAIN_METADATA[chainId as SupportedChainId]
}
