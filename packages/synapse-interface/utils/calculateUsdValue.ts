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

  // Parse amount to number (remove commas first)
  const numericAmount = parseFloat(amount.replace(/,/g, ''))

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
 * Formats USD difference for inline display next to USD values
 *
 * @param diff - USD difference between output and input values (null = unavailable)
 * @returns Formatted string: " (+$5.00)", " (-$5.00)", or "" if null/under 1 cent
 * @example
 * formatInlineUsdDifference(5.25)   // " (+$5.25)"
 * formatInlineUsdDifference(-3.50)  // " (-$3.50)"
 * formatInlineUsdDifference(0.005)  // "" (don't show if under 1 cent)
 * formatInlineUsdDifference(null)   // ""
 */
export const formatInlineUsdDifference = (diff: number | null): string => {
  // Return empty string if difference is unavailable
  if (diff === null) {
    return ''
  }

  // Get absolute value
  const absValue = Math.abs(diff)

  // Don't show if under 1 cent
  if (absValue < 0.01) {
    return ''
  }

  // Format with sign and 2 decimal places
  const sign = diff >= 0 ? '+' : '-'
  return ` (${sign}$${absValue.toFixed(2)})`
}
