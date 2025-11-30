interface FormatOptions {
  fullAmount?: boolean
  standardDigits?: number
  useCompactNotation?: boolean
  compactDigits?: number
  minimumAmount?: number
  roundingMode?: string
}

export const formatAmount = (
  amount: string,
  options?: FormatOptions
): string => {
  if (amount === '') return ''

  const floatAmount = parseFloat(amount)

  try {
    if (!Number.isFinite(floatAmount)) {
      throw new TypeError(`"${amount}" is not a finite number`)
    }
  } catch ({ name, message }) {
    // console.error(name, message)
    return amount
  }

  const fullAmount = options?.fullAmount ?? false
  const standardDigits = options?.standardDigits ?? 4
  const useCompactNotation = options?.useCompactNotation ?? true
  const compactDigits = options?.compactDigits ?? (useCompactNotation ? 2 : 0)
  const minimumAmount = options?.minimumAmount ?? 0.0001

  const locales = 'en-UK'

  if (!floatAmount) return '0.0'

  if (fullAmount) return Intl.NumberFormat(locales).format(floatAmount)

  if (floatAmount < minimumAmount) return `< ${minimumAmount}`

  const absAmount = Math.abs(floatAmount)

  if (absAmount < 0.0001)
    return Intl.NumberFormat(locales, {
      maximumSignificantDigits: 1,
    }).format(floatAmount)

  if (absAmount < 1)
    return Intl.NumberFormat(locales, {
      minimumFractionDigits: standardDigits,
    }).format(floatAmount)

  if (absAmount < 1000)
    return Intl.NumberFormat(locales, {
      minimumSignificantDigits: standardDigits,
      maximumSignificantDigits: standardDigits,
    }).format(floatAmount)

  if (absAmount < 1000000)
    return Intl.NumberFormat(locales, {
      maximumFractionDigits: 0,
    }).format(floatAmount)

  return Intl.NumberFormat(locales, {
    minimumFractionDigits: compactDigits,
    maximumFractionDigits: compactDigits,
    notation: useCompactNotation ? 'compact' : 'standard',
  }).format(floatAmount)
}

/**
 * Format token amount based on price so each decimal digit represents at least $0.01 USD.
 *
 * Algorithm: decimals = clamp(floor(log10(price * 100)), 0, 18)
 *
 * Examples:
 * - ETH at $3,500: floor(log10(350,000)) = 5 decimals → "0.12345"
 * - USDC at $1.00: floor(log10(100)) = 2 decimals → "1,234.56"
 * - BTC at $100,000: floor(log10(10,000,000)) = 7 decimals → "0.0012345"
 *
 * @param amount - Token amount as string
 * @param price - Token price in USD (null/undefined if unavailable)
 * @returns Formatted amount string with thousand separators
 */
export const formatAmountByPrice = (
  amount: string,
  price: number | null | undefined
): string => {
  if (amount === '') return ''

  // Fallback to existing formatter if no price available
  if (price === null || price === undefined || price <= 0) {
    return formatAmount(amount)
  }

  const floatAmount = Number.parseFloat(amount.replaceAll(',', ''))

  if (!Number.isFinite(floatAmount)) {
    return amount
  }

  if (floatAmount === 0) {
    return '0'
  }

  // Calculate decimals where 1 unit at that position = $0.01
  const decimals = Math.floor(Math.log10(price * 100))
  const clampedDecimals = Math.max(0, Math.min(18, decimals))

  return new Intl.NumberFormat('en-UK', {
    minimumFractionDigits: 0,
    maximumFractionDigits: clampedDecimals,
  }).format(floatAmount)
}

/**
 * Returns tooltip text if showValue differs numerically from fullValue.
 * Handles cases like "20.0" vs "20" (equal) and "1,234" vs "1234" (equal).
 */
export const getTooltipValue = (
  showValue: string,
  fullValue: string,
  symbol?: string
): string | undefined => {
  if (!showValue || !fullValue) return undefined

  const showNum = Number.parseFloat(showValue.replaceAll(',', ''))
  const fullNum = Number.parseFloat(fullValue.replaceAll(',', ''))
  if (Math.abs(showNum - fullNum) < 1e-8) return undefined

  return symbol ? `${fullValue} ${symbol}` : fullValue
}
