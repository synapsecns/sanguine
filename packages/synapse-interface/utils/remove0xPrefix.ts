export const remove0xPrefix = (str?: string): string => {
  if (!str) {
    return str
  }
  if (str.startsWith('0x')) {
    return str.slice(2)
  }
  return str
}
