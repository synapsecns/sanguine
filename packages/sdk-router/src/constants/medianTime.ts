import { SupportedChainId } from './chainIds'

/**
 * Median time (in seconds) for a chain's block to be produced.
 * Note: chains are listed in alphabetical order to make it easier to add new chains.
 */
export const MEDIAN_TIME_BLOCK: Record<SupportedChainId, number> = {
  [SupportedChainId.ARBITRUM]: 0.25,
  [SupportedChainId.AURORA]: 1,
  [SupportedChainId.AVALANCHE]: 2,
  [SupportedChainId.BASE]: 2,
  [SupportedChainId.BERACHAIN]: 2,
  [SupportedChainId.BLAST]: 2,
  [SupportedChainId.BOBA]: 2,
  [SupportedChainId.BSC]: 3,
  [SupportedChainId.CANTO]: 6,
  [SupportedChainId.CRONOS]: 6,
  [SupportedChainId.DFK]: 2,
  [SupportedChainId.DOGECHAIN]: 2,
  [SupportedChainId.ETH]: 12,
  [SupportedChainId.FANTOM]: 2,
  [SupportedChainId.HARMONY]: 2,
  [SupportedChainId.HYPEREVM]: 0.2,
  [SupportedChainId.KLAYTN]: 1,
  [SupportedChainId.LINEA]: 2,
  [SupportedChainId.METIS]: 2,
  [SupportedChainId.MOONBEAM]: 6,
  [SupportedChainId.MOONRIVER]: 6,
  [SupportedChainId.OPTIMISM]: 2,
  [SupportedChainId.POLYGON]: 2,
  [SupportedChainId.SCROLL]: 3,
  [SupportedChainId.UNICHAIN]: 1,
  [SupportedChainId.WORLDCHAIN]: 2,
}

/**
 * Median time (in seconds) for a SynapseBridge transaction to be completed,
 * when the transaction is sent from a given chain.
 */
export const MEDIAN_TIME_BRIDGE = {
  [SupportedChainId.ETH]: 420,
  [SupportedChainId.OPTIMISM]: 1290,
  [SupportedChainId.CRONOS]: 30,
  [SupportedChainId.BSC]: 60,
  [SupportedChainId.POLYGON]: 300,
  [SupportedChainId.FANTOM]: 90,
  [SupportedChainId.BOBA]: 180,
  [SupportedChainId.METIS]: 60,
  [SupportedChainId.MOONBEAM]: 90,
  [SupportedChainId.MOONRIVER]: 45,
  [SupportedChainId.DOGECHAIN]: 15,
  [SupportedChainId.CANTO]: 150,
  [SupportedChainId.KLAYTN]: 15,
  [SupportedChainId.BASE]: 1230,
  [SupportedChainId.ARBITRUM]: 30,
  [SupportedChainId.AVALANCHE]: 30,
  [SupportedChainId.DFK]: 30,
  [SupportedChainId.BLAST]: 1230,
  [SupportedChainId.AURORA]: 30,
  [SupportedChainId.HARMONY]: 30,
}

/**
 * Median time (in seconds) for a SynapseCCTP transaction to be completed,
 * when the transaction is sent from a given chain.
 */
export const MEDIAN_TIME_CCTP = {
  [SupportedChainId.ETH]: 1020,
  [SupportedChainId.OPTIMISM]: 1170,
  [SupportedChainId.BASE]: 1170,
  [SupportedChainId.ARBITRUM]: 1170,
  [SupportedChainId.AVALANCHE]: 30,
  [SupportedChainId.POLYGON]: 480,
}

/**
 * Median time (in seconds) for a SynapseRFQ transaction to be completed,
 * when the transaction is sent from a given chain.
 * TODO: Update this value once we have a better estimate.
 */
export const MEDIAN_TIME_RFQ = {
  [SupportedChainId.ETH]: 25,
  [SupportedChainId.OPTIMISM]: 15,
  [SupportedChainId.BSC]: 15,
  [SupportedChainId.BASE]: 15,
  [SupportedChainId.ARBITRUM]: 15,
  [SupportedChainId.LINEA]: 15,
  [SupportedChainId.BLAST]: 15,
  [SupportedChainId.SCROLL]: 15,
  [SupportedChainId.WORLDCHAIN]: 15,
  [SupportedChainId.UNICHAIN]: 15,
  [SupportedChainId.BERACHAIN]: 15,
  [SupportedChainId.HYPEREVM]: 15,
}
