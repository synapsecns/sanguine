import { useDefiLlamaPrice } from '@hooks/useDefiLlamaPrice'
import { Token } from '@utils/types'
import { formatBigIntToString } from '@utils/bigint/format'
import { AbsoluteThreshold, PercentageThreshold } from '@constants/slippage'

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

const DEFAULT_RESULT: UseUsdSlippageResult = {
  slippage: null,
  usdDifference: null,
  isLoading: false,
  error: null,
  textColor: '',
}

/**
 * Extracts token decimals for a given chain
 */
const getTokenDecimals = (
  token: Token,
  chainId: number
): number | undefined => {
  return typeof token.decimals === 'number'
    ? token.decimals
    : token.decimals[chainId]
}

/**
 * Determines text color based on slippage percentage and USD difference
 */
const calculateSlippageColor = (
  slippage: number,
  usdDifference: number
): string => {
  // We show white if either of values are in the neutral range
  const isSlippageNeutral = Math.abs(slippage) <= PercentageThreshold.NEUTRAL
  const isDiffNeutral = Math.abs(usdDifference) <= AbsoluteThreshold.NEUTRAL
  if (isSlippageNeutral || isDiffNeutral) return 'text-white'

  // For positive slippage (gain), we show green if it's not neutral
  if (slippage > 0) return 'text-green-500'

  // Show amber if either of values are in the warning range
  const isSlippageWarning = Math.abs(slippage) <= PercentageThreshold.WARNING
  const isDiffWarning = Math.abs(usdDifference) <= AbsoluteThreshold.WARNING
  if (isSlippageWarning || isDiffWarning) return 'text-amber-500'

  // Show red if both are out of warning range
  return 'text-red-500'
}

/**
 * Calculates USD-based slippage between origin and destination tokens
 *
 * Formula: ((valueOut - valueIn) / valueIn) * 100
 *
 * @returns slippage percentage (positive = gain, negative = loss), USD difference, loading state, and error
 */
export const useUsdSlippage = ({
  originToken,
  destToken,
  originChainId,
  destChainId,
  inputAmount,
  outputAmount,
}: UseUsdSlippageParams): UseUsdSlippageResult => {
  // Fetch prices for both tokens
  const originPrice = useDefiLlamaPrice(originToken)
  const destPrice = useDefiLlamaPrice(destToken)

  // Validate all required parameters are present
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
  if (!hasAllParams) return DEFAULT_RESULT

  // Handle loading state (prices not yet fetched)
  if (originPrice === undefined || destPrice === undefined) {
    return { ...DEFAULT_RESULT, isLoading: true }
  }

  // Handle unavailable prices
  if (originPrice === null || destPrice === null) {
    return { ...DEFAULT_RESULT, error: 'Price data unavailable' }
  }

  // Extract decimals for both tokens
  const originDecimals = getTokenDecimals(originToken, originChainId)
  const destDecimals = getTokenDecimals(destToken, destChainId)

  if (originDecimals === undefined) {
    console.error(
      'Missing decimals for origin token',
      originToken.symbol,
      'on chain',
      originChainId
    )
    return { ...DEFAULT_RESULT, error: 'Missing token decimals' }
  }

  if (destDecimals === undefined) {
    console.error(
      'Missing decimals for destination token',
      destToken.symbol,
      'on chain',
      destChainId
    )
    return { ...DEFAULT_RESULT, error: 'Missing token decimals' }
  }

  // Calculate slippage
  try {
    const inputAmountDecimal = parseFloat(
      formatBigIntToString(inputAmount, originDecimals)
    )
    const outputAmountDecimal = parseFloat(
      formatBigIntToString(outputAmount, destDecimals)
    )

    const valueIn = inputAmountDecimal * originPrice
    const valueOut = outputAmountDecimal * destPrice

    // Guard against division by zero
    if (valueIn === 0 || valueOut === 0) {
      return DEFAULT_RESULT
    }

    const usdDifference = valueOut - valueIn
    const slippage = (usdDifference / valueIn) * 100

    return {
      slippage,
      usdDifference,
      isLoading: false,
      error: null,
      textColor: calculateSlippageColor(slippage, usdDifference),
    }
  } catch (err) {
    console.error('Error calculating USD slippage:', err)
    return DEFAULT_RESULT
  }
}
