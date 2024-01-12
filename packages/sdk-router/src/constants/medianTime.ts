import { SupportedChainId } from './chainIds'

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
  [SupportedChainId.ETH]: 30,
  [SupportedChainId.OPTIMISM]: 30,
  [SupportedChainId.ARBITRUM]: 30,
}
