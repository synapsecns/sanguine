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
  if (amount === '') {
    return ''
  }

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

  if (!floatAmount) {
    return '0.0'
  }

  if (fullAmount) {
    return Intl.NumberFormat(locales).format(floatAmount)
  }

  if (floatAmount < minimumAmount) {
    return `< ${minimumAmount}`
  }

  const absAmount = Math.abs(floatAmount)

  if (absAmount < 0.0001) {
    return Intl.NumberFormat(locales, {
      maximumSignificantDigits: 1,
    }).format(floatAmount)
  }

  if (absAmount < 1) {
    return Intl.NumberFormat(locales, {
      minimumFractionDigits: standardDigits,
    }).format(floatAmount)
  }

  if (absAmount < 1000) {
    return Intl.NumberFormat(locales, {
      minimumSignificantDigits: standardDigits,
      maximumSignificantDigits: standardDigits,
    }).format(floatAmount)
  }

  if (absAmount < 1000000) {
    return Intl.NumberFormat(locales, {
      maximumFractionDigits: 0,
    }).format(floatAmount)
  }

  return Intl.NumberFormat(locales, {
    minimumFractionDigits: compactDigits,
    maximumFractionDigits: compactDigits,
    notation: useCompactNotation ? 'compact' : 'standard',
  }).format(floatAmount)
}
