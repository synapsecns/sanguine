/**
  This regex matches a string that either:
  - consists only of one or more '0' characters
  - consists of a single '0' followed by a '.' and zero or more '0' characters
  - consists solely of a '.' followed by one or more '0' characters
**/
export const isOnlyZeroes = (input: string): boolean => {
  const regex = /^(0+|0?\.0+)$/

  return regex.test(input.trim())
}
