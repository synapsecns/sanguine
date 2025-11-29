import { commify } from '@/utils/bigint/format'
import {
  USD_DIFFERENCE_PLACEHOLDER,
  USD_VALUE_PLACEHOLDER,
} from '@/constants/placeholders'

/**
 * Formats a USD value with thousand separators and 2 decimal places
 *
 * @param value - The numeric USD value to format
 * @returns Formatted USD string with proper separators and 2 decimals
 *
 * Examples: "12,345.60", "5,678.90", "123.45", "0.10"
 */
export const formatUsdValue = (value: number): string => {
  const formatted = value.toFixed(2)

  // Split to preserve trailing zeros (toFixed returns "0.10" but commify strips to "0.1")
  const [integerPart, decimalPart = '00'] = formatted.split('.')
  return `${commify(integerPart)}.${decimalPart}`
}

/**
 * Calculates and formats USD value for a token amount
 *
 * @param amount - Token amount as string (e.g., "100.5")
 * @param price - USD price per token (undefined = loading, null = unavailable)
 * @returns Formatted USD string: "$123.45", "<$0.01", or ""
 */
export const calculateUsdValue = (
  amount: string | null | undefined,
  price: number | null | undefined
): string => {
  // Return placeholder if price is loading (undefined) or unavailable (null)
  if (price === undefined || price === null) {
    return USD_VALUE_PLACEHOLDER
  }

  // Return placeholder if amount is empty or invalid
  if (!amount || amount === '') {
    return USD_VALUE_PLACEHOLDER
  }

  // Parse amount to number (remove commas first)
  const numericAmount = Number.parseFloat(amount.replaceAll(',', ''))

  // Return placeholder if amount is invalid or zero
  if (Number.isNaN(numericAmount) || numericAmount === 0) {
    return USD_VALUE_PLACEHOLDER
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
    return USD_DIFFERENCE_PLACEHOLDER
  }

  // Get absolute value
  const absValue = Math.abs(diff)

  // Don't show if under 1 cent
  if (absValue < 0.01) {
    return USD_DIFFERENCE_PLACEHOLDER
  }

  // Format with sign, dynamic decimal places, and thousand separators
  const sign = diff >= 0 ? '+' : '−'
  return ` (${sign}$${formatUsdValue(absValue)})`
}
