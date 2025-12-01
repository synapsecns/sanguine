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
const _formatUsd = (value: number): string => {
  const [integerPart, decimalPart = '00'] = value.toFixed(2).split('.')
  return `${commify(integerPart)}.${decimalPart}`
}

/**
 * Formats a USD value as a display string
 *
 * @param usdValue - USD value to format (null = unavailable)
 * @returns Formatted USD string: "$123.45", "<$0.01", or ""
 */
export const formatUsdValue = (usdValue: number | null): string => {
  if (!usdValue) {
    return USD_VALUE_PLACEHOLDER
  }

  // Handle very small amounts (less than 1 cent)
  if (usdValue < 0.01) {
    return '<$0.01'
  }

  // Format with dynamic decimal places and thousand separators
  return `$${_formatUsd(usdValue)}`
}

/**
 * Calculates the USD value for a token amount
 *
 * @param amount - Token amount as string (e.g., "100.5")
 * @param price - USD price per token (undefined = loading, null = unavailable)
 * @returns Numeric USD value, or null if calculation not possible
 */
export const calculateUsdValue = (
  amount: string | null | undefined,
  price: number | null | undefined
): number | null => {
  if (!Number.isFinite(price)) {
    return null
  }

  const numericAmount = Number.parseFloat(amount?.replaceAll(',', '') ?? '')
  if (!Number.isFinite(numericAmount)) {
    return null
  }

  return numericAmount * price
}

/**
 * Calculates total USD value for a token amount plus optional native component
 *
 * Universal utility for:
 * - Input value: token amount + native fee
 * - Output value: token amount + gas airdrop
 *
 * @param tokenAmount - Token amount as string
 * @param tokenPrice - Token USD price
 * @param nativeAmount - Native amount as string (fee or airdrop)
 * @param nativePrice - Native token USD price
 * @returns Total USD value, or null if token value unavailable
 */
export const calculateTotalUsdValue = (
  tokenAmount: string | null | undefined,
  tokenPrice: number | null | undefined,
  nativeAmount: string | null | undefined,
  nativePrice: number | null | undefined
): {
  tokenValue: number | null
  nativeValue: number | null
  totalValue: number | null
} => {
  const tokenValue = calculateUsdValue(tokenAmount, tokenPrice)
  const nativeValue = calculateUsdValue(nativeAmount, nativePrice)
  const totalValue = Number.isFinite(tokenValue)
    ? tokenValue + (nativeValue ?? 0)
    : null
  return { tokenValue, nativeValue, totalValue }
}

/**
 * Formats a tooltip showing USD breakdown (token + native component)
 *
 * Rounds native first (consistent across quotes), derives token as remainder.
 *
 * @param nativeValue - Native USD value (fee or airdrop)
 * @param totalValue - Total USD value - what's displayed to user
 * @param nativeLabel - Label for native component (e.g., "gas fee", "gas airdrop")
 * @returns Tooltip string: "$100.00 + $0.50 (gas fee)" or null if no breakdown needed
 */
export const formatUsdBreakdownTooltip = (
  nativeValue: number | null,
  totalValue: number | null,
  nativeLabel: string
): string | null => {
  if (!nativeValue || !totalValue) {
    return null
  }

  // Work in cents (integers) to avoid floating-point issues
  // Round native first (consistent across quotes), derive token as remainder
  const totalCents = Math.round(totalValue * 100)
  const nativeCents = Math.round(nativeValue * 100)
  const tokenCents = totalCents - nativeCents

  // Don't show breakdown if native < 1 cent
  if (nativeCents < 1) {
    return null
  }

  return `${formatUsdValue(tokenCents / 100)} + ${formatUsdValue(nativeCents / 100)} (${nativeLabel})`
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
  return ` (${sign}$${_formatUsd(absValue)})`
}
