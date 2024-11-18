export const validateDecimals = (
  amount: string,
  expectedDecimals: number
): boolean => {
  const parts = amount.split('.')
  if (parts.length === 2) {
    const decimals = parts[1].length
    return decimals <= expectedDecimals
  }
  return true
}
