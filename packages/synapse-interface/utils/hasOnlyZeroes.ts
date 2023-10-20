export const hasOnlyZeroes = (input: string): boolean => {
  return /^0+(\.0+)?$/.test(input)
}
