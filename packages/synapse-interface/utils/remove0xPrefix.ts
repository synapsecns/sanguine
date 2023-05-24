export const remove0xPrefix = (str: string): string => {
  if (str.startsWith('0x')) {
    return str.slice(2)
  }
  return str
}
