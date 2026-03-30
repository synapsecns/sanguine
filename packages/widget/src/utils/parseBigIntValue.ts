const DECIMAL_BIGINT_PATTERN = /^-?\d+$/

export const parseBigIntValue = (value: unknown): bigint | null => {
  if (typeof value === 'bigint') {
    return value
  }

  if (typeof value === 'string') {
    const normalizedValue = value.trim()

    if (DECIMAL_BIGINT_PATTERN.test(normalizedValue)) {
      return BigInt(normalizedValue)
    }
  }

  return null
}
