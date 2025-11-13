import { useDefiLlamaPrice } from '@hooks/useDefiLlamaPrice'
import { Token } from '@utils/types'
import { formatBigIntToString } from '@utils/bigint/format'
import {
  SLIPPAGE_WARNING_THRESHOLD,
  USD_SLIPPAGE_WARNING_THRESHOLD,
} from '@constants/slippage'

interface UseUsdSlippageParams {
  originToken: Token | null
  destToken: Token | null
  originChainId: number | null
  destChainId: number | null
  inputAmount: bigint | null
  outputAmount: bigint | null
}

interface UseUsdSlippageResult {
  slippage: number | null
  usdDifference: number | null
  isLoading: boolean
  error: string | null
  textColor: string
}

/**
 * Hook to calculate USD-based slippage between origin and destination tokens
 *
 * Slippage formula: ((valueOut - valueIn) / valueIn) * 100
 *
 * @returns slippage percentage (positive = gain, negative = loss), loading state, and error
 */
export const useUsdSlippage = ({
  originToken,
  destToken,
  originChainId,
  destChainId,
  inputAmount,
  outputAmount,
}: UseUsdSlippageParams): UseUsdSlippageResult => {
  // Fetch prices using SWR hooks
  const originPrice = useDefiLlamaPrice(originToken, originChainId)
  const destPrice = useDefiLlamaPrice(destToken, destChainId)

  // Check if we have all required parameters
  const hasAllParams = Boolean(
    originToken &&
      destToken &&
      originChainId !== null &&
      destChainId !== null &&
      inputAmount &&
      outputAmount &&
      inputAmount > 0n &&
      outputAmount > 0n
  )

  // Check if prices are still loading (undefined = SWR fetching)
  const arePricesLoading = originPrice === undefined || destPrice === undefined

  // Check if prices are unavailable (null = not found in DefiLlama)
  const arePricesUnavailable = originPrice === null || destPrice === null

  // Calculate slippage (no useMemo needed - calculation is lightweight)
  let slippage: number | null = null
  let usdDifference: number | null = null

  if (hasAllParams && !arePricesLoading && !arePricesUnavailable) {
    try {
      // Get decimals for both tokens
      const originDecimals =
        typeof originToken.decimals === 'number'
          ? originToken.decimals
          : originToken.decimals[originChainId] ?? 18

      const destDecimals =
        typeof destToken.decimals === 'number'
          ? destToken.decimals
          : destToken.decimals[destChainId] ?? 18

      // Convert amounts to decimal numbers
      const inputAmountDecimal = parseFloat(
        formatBigIntToString(inputAmount, originDecimals)
      )
      const outputAmountDecimal = parseFloat(
        formatBigIntToString(outputAmount, destDecimals)
      )

      // Calculate USD values
      const valueIn = inputAmountDecimal * originPrice!
      const valueOut = outputAmountDecimal * destPrice!

      // Guard against division by zero
      if (valueIn === 0 || valueOut === 0) {
        slippage = null
        usdDifference = null
      } else {
        // Calculate USD difference and slippage
        usdDifference = valueOut - valueIn
        slippage = (usdDifference / valueIn) * 100
      }
    } catch (err) {
      console.error('Error calculating USD slippage:', err)
      slippage = null
    }
  }

  // Determine loading state
  const isLoading = hasAllParams && arePricesLoading

  // Determine error state (prices were fetched but not found)
  const error =
    hasAllParams && !arePricesLoading && arePricesUnavailable
      ? 'Price data unavailable'
      : null

  // Determine text color based on slippage value and USD difference
  // Amber if percentage loss <= 2.5% OR USD loss <= $1
  const textColor =
    slippage === null || usdDifference === null
      ? ''
      : slippage >= 0
      ? 'text-green-500'
      : slippage > SLIPPAGE_WARNING_THRESHOLD ||
        usdDifference > USD_SLIPPAGE_WARNING_THRESHOLD
      ? 'text-amber-500'
      : 'text-red-500'

  return { slippage, usdDifference, isLoading, error, textColor }
}
