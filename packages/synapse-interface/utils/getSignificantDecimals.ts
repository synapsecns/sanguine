export const getSignificantDecimals = (
  numberString: string,
  defaultDecimals: number = 2
): number => {
  const parts = numberString.split('.')
  const decimalPart = parts[1]

  if (!decimalPart) return 0

  if (/^0*$/.test(decimalPart)) {
    return defaultDecimals
  }

  let significantDecimals = 0

  for (let i = 0; i < decimalPart.length; i++) {
    if (decimalPart[i] !== '0') {
      significantDecimals = i + 1
    }
  }

  return significantDecimals
}
