import { Token } from '@utils/types'
import { useDefiLlamaPrice } from '@hooks/useDefiLlamaPrice'
import { calculateUsdValue } from '@utils/calculateUsdValue'

/**
 * Custom hook to fetch token price and calculate USD value display
 * Combines useDefiLlamaPrice and calculateUsdValue for DRY code
 *
 * @param token - The token to fetch price for
 * @param amount - The token amount as a string (e.g., "100.5")
 * @param chainId - The chain ID where the token exists
 * @returns Formatted USD string: "$123.45", "<$0.01", or "—"
 *
 * @example
 * ```tsx
 * const usdValue = useUsdDisplay(fromToken, localInputValue, fromChainId)
 * // Returns: "$1,234.56" or "—" if price unavailable
 * ```
 */
export const useUsdDisplay = (
  token: Token | null,
  amount: string | null | undefined,
  chainId: number | null
): string => {
  const price = useDefiLlamaPrice(token, chainId)
  return calculateUsdValue(amount, price)
}
