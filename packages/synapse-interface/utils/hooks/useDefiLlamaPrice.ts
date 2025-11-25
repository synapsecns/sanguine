import { useSwr } from '@hooks/useSwr'
import { Token } from '@utils/types'
import { CHAIN_ID_TO_DEFILLAMA_NAME } from '@constants/defiLlama'

/**
 * Hook to fetch token price from DefiLlama API
 *
 * Returns undefined while loading, null if price unavailable, or the price number.
 *
 * @param token - Token object with addresses
 * @returns Price in USD, undefined (loading), or null (not available)
 */
export const useDefiLlamaPrice = (
  token: Pick<Token, 'addresses' | 'priceOverride'> | null
): number | null | undefined => {
  // Use first chain's address for consistent pricing across all chains for multi-chain tokens
  const chainId = token ? Object.keys(token.addresses)[0] : null

  // Get chain name from mapping
  const chainName = chainId ? CHAIN_ID_TO_DEFILLAMA_NAME[chainId] : null

  // Warn if chain ID is not in mapping
  if (chainId && !chainName) {
    console.warn(
      `DefiLlama chain mapping missing for chainId: ${chainId}. Add to CHAIN_ID_TO_DEFILLAMA_NAME in constants/defiLlama.ts`
    )
  }

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

  const { data, error } = useSwr(apiUrl, {
    refreshInterval: 60000,
    dedupingInterval: 10000,
  })

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
