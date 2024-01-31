export const formatBigIntToString = (
  bi: bigint | undefined,
  nativePrecision: number | undefined,
  decimalPlaces?: number
) => {
  if (typeof bi !== 'bigint' && !bi) {
    return
  }

  if (!nativePrecision) {
    return
  }

  try {
    // Check if input is zero
    if (bi === 0n) {
      return '0.0'
    }

    // Check if the input is negative
    const isNegative = bi < 0n
    if (isNegative) {
      // Convert to positive for the calculation
      bi = -bi
    }
    // Convert to string and add padding zeros if necessary
    let str = bi.toString().padStart(nativePrecision, '0')

    // Insert decimal point
    const idx = str.length - nativePrecision
    str = `${str.slice(0, idx)}.${str.slice(idx)}`

    // Handle values below zero by adding a '0' before the decimal point
    if (str.startsWith('.')) {
      str = '0' + str
    }

    // Trim to desired number of decimal places
    if (decimalPlaces !== undefined) {
      const decimalIdx = str.indexOf('.')
      str = str.slice(0, decimalIdx + decimalPlaces + 1)
    }

    // Add the negative sign back if necessary
    if (isNegative) {
      str = '-' + str
    }

    return str
  } catch (error) {
    console.log(`error`, error)
  }
}
