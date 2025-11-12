import { useDefiLlamaPrice } from '@hooks/useDefiLlamaPrice'
import { Token } from '@utils/types'
import { formatBigIntToString } from '@utils/bigint/format'

const SLIPPAGE_WARNING_THRESHOLD = -2.5

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
  const hasAllParams =
    originToken &&
    destToken &&
    originChainId !== null &&
    destChainId !== null &&
    inputAmount &&
    outputAmount &&
    inputAmount > 0n &&
    outputAmount > 0n

  // Calculate slippage (no useMemo needed - calculation is lightweight)
  let slippage: number | null = null

  if (hasAllParams) {
    // Still loading (undefined means SWR is fetching)
    if (originPrice !== undefined && destPrice !== undefined) {
      // Price not available (null means not found in DefiLlama)
      if (originPrice !== null && destPrice !== null) {
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
            formatBigIntToString(inputAmount, originDecimals, originDecimals)
          )
          const outputAmountDecimal = parseFloat(
            formatBigIntToString(outputAmount, destDecimals, destDecimals)
          )

          // Calculate USD values
          const valueIn = inputAmountDecimal * originPrice
          const valueOut = outputAmountDecimal * destPrice

          // Calculate slippage: ((valueOut - valueIn) / valueIn) * 100
          slippage = ((valueOut - valueIn) / valueIn) * 100
        } catch (err) {
          console.error('Error calculating USD slippage:', err)
          slippage = null
        }
      }
    }
  }

  // Determine loading state (prices are undefined = SWR loading)
  const isLoading =
    hasAllParams && (originPrice === undefined || destPrice === undefined)

  // Determine error state (prices are null = not found)
  const error =
    hasAllParams &&
    originPrice !== undefined &&
    destPrice !== undefined &&
    (originPrice === null || destPrice === null)
      ? 'Price data unavailable'
      : null

  // Determine text color based on slippage value
  const textColor =
    slippage === null
      ? ''
      : slippage >= 0
      ? 'text-green-500'
      : slippage > SLIPPAGE_WARNING_THRESHOLD
      ? 'text-amber-500'
      : 'text-red-500'

  return { slippage, isLoading, error, textColor }
}
