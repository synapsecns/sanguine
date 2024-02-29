export const trimTrailingZeroesAfterDecimal = (input: string): string => {
  const parts = input.split('.')

  if (parts.length === 2) {
    const integerPart = parts[0]
    let fractionalPart = parts[1]

    // Remove trailing '0's from the fractional part
    fractionalPart = fractionalPart.replace(/0+$/, '')

    // Reconstruct the trimmed number
    if (fractionalPart.length > 0) {
      return `${integerPart}.${fractionalPart}`
    } else {
      return integerPart
    }
  }

  return input
}
