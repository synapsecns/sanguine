/**
 * Mapping of chain IDs to DefiLlama chain name identifiers
 * Based on https://api.llama.fi/chains
 *
 * Used for constructing DefiLlama price API requests:
 * https://coins.llama.fi/prices/current/{chainName}:{tokenAddress}
 */
export const CHAIN_ID_TO_DEFILLAMA_NAME: Record<number, string> = {
  1: 'ethereum',
  10: 'optimism',
  25: 'cronos',
  56: 'binance',
  130: 'unichain',
  137: 'polygon',
  250: 'fantom',
  288: 'boba',
  480: 'wc',
  999: 'hyperliquid',
  1088: 'metis',
  1284: 'moonbeam',
  1285: 'moonriver',
  2000: 'dogechain',
  7700: 'canto',
  8217: 'klaytn',
  8453: 'base',
  42161: 'arbitrum',
  43114: 'avax',
  53935: 'dfk',
  59144: 'linea',
  534352: 'scroll',
  80094: 'berachain',
  81457: 'blast',
  1313161554: 'aurora',
  1666600000: 'harmony',
}
