/**
 * Calculates and formats USD value for a token amount
 *
 * @param amount - Token amount as string (e.g., "100.5")
 * @param price - USD price per token (undefined = loading, null = unavailable)
 * @returns Formatted USD string: "$123.45", "<$0.01", or "—"
 */
export const calculateUsdValue = (
  amount: string | null | undefined,
  price: number | null | undefined
): string => {
  // Return placeholder if price is loading (undefined) or unavailable (null)
  if (price === undefined || price === null) {
    return '—'
  }

  // Return placeholder if amount is empty or invalid
  if (!amount || amount === '') {
    return '—'
  }

  // Parse amount to number
  const numericAmount = parseFloat(amount)

  // Return placeholder if amount is invalid or zero
  if (isNaN(numericAmount) || numericAmount === 0) {
    return '—'
  }

  // Calculate USD value
  const usdValue = numericAmount * price

  // Handle very small amounts (less than 1 cent)
  if (usdValue < 0.01) {
    return '<$0.01'
  }

  // Format with 2 decimal places
  return `$${usdValue.toFixed(2)}`
}

/**
 * Formats USD difference for display in slippage information
 *
 * @param diff - USD difference between output and input values (null = unavailable)
 * @returns Formatted string: " ($5.00)", " (<$0.01)", or "" if null
 * @example
 * formatUsdDifference(5.25)   // " ($5.25)"
 * formatUsdDifference(0.005)  // " (<$0.01)"
 * formatUsdDifference(null)   // ""
 */
export const formatUsdDifference = (diff: number | null): string => {
  // Return empty string if difference is unavailable
  if (diff === null) {
    return ''
  }

  // Get absolute value for display
  const absValue = Math.abs(diff)

  // Handle very small amounts (less than 1 cent)
  if (absValue < 0.01) {
    return ' (<$0.01)'
  }

  // Format with 2 decimal places, always show absolute value
  return ` ($${absValue.toFixed(2)})`
}
