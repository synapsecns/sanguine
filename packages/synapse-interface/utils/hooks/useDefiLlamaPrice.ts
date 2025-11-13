import { useSwr } from '@hooks/useSwr'
import { Token } from '@utils/types'
import { CHAIN_ID_TO_DEFILLAMA_NAME } from '@constants/defiLlama'

/**
 * Hook to fetch token price from DefiLlama API
 *
 * Follows the SWR pattern used in useCoingeckoPrice for consistency.
 * Returns undefined while loading, null if price unavailable, or the price number.
 *
 * @param token - Token object with addresses
 * @param chainId - Chain ID to fetch price for
 * @returns Price in USD, undefined (loading), or null (not available)
 */
export const useDefiLlamaPrice = (
  token: Token | null,
  chainId: number | null
): number | null | undefined => {
  // Get chain name from mapping
  const chainName = chainId ? CHAIN_ID_TO_DEFILLAMA_NAME[chainId] : null

  // Get token address
  const address =
    token && chainId ? token.addresses[chainId]?.toLowerCase() : null

  // Build coin key (chain:address format)
  const coinKey = chainName && address ? `${chainName}:${address}` : null

  // Build API URL - return null if priceOverride exists (SWR won't fetch)
  const apiUrl =
    token?.priceOverride !== undefined
      ? null
      : coinKey
      ? `https://coins.llama.fi/prices/current/${coinKey}`
      : null

  const { data, error } = useSwr(apiUrl)

  // Return override price if defined (checked after hook call)
  if (token?.priceOverride !== undefined) {
    return token.priceOverride
  }

  // Log error for debugging but don't throw (graceful degradation)
  if (error) {
    console.warn('DefiLlama price fetch failed:', error)
    return null
  }

  if (data && coinKey) {
    // Return price if available, null if not found
    return data.coins?.[coinKey]?.price ?? null
  }

  // Return undefined while loading or if no valid URL
  return undefined
}
