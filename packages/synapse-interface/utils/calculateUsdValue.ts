import { commify } from '@/utils/bigint/format'

/**
 * Formats a USD value with thousand separators and dynamic decimal places
 *
 * @param value - The numeric USD value to format
 * @returns Formatted USD string with proper separators and decimals
 *
 * Formatting rules:
 * - >= $10,000: 0 decimals (e.g., "12,345")
 *   Rationale: Cents are insignificant at this scale, cleaner display
 * - >= $1,000: 1 decimal (e.g., "5,678.9")
 *   Rationale: Balances precision with readability
 * - < $1,000: 2 decimals (e.g., "123.45", "0.10")
 *   Rationale: Standard currency precision, important for smaller amounts
 * - All values get thousand separators via commify()
 */
export const formatUsdValue = (value: number): string => {
  const absValue = Math.abs(value)

  let decimals: number
  if (absValue >= 10000) {
    decimals = 0
  } else if (absValue >= 1000) {
    decimals = 1
  } else {
    decimals = 2
  }

  const formatted = value.toFixed(decimals)

  // Special handling for 2-decimal values to preserve trailing zeros
  // toFixed() returns "0.10" but commify() would strip to "0.1"
  // We split and rejoin to maintain precision (e.g., "$0.10" not "$0.1")
  if (decimals === 2) {
    const [integerPart, decimalPart] = formatted.split('.')
    return `${commify(integerPart)}.${decimalPart}`
  }

  return commify(formatted)
}

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

  // Format with dynamic decimal places and thousand separators
  return `$${formatUsdValue(usdValue)}`
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

  // Format with sign, dynamic decimal places, and thousand separators
  const sign = diff >= 0 ? '+' : '-'
  return ` (${sign}$${formatUsdValue(absValue)})`
}
