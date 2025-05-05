export const isSameAddress = (addr1?: string, addr2?: string): boolean => {
  return !!addr1 && !!addr2 && addr1.toLowerCase() === addr2.toLowerCase()
}
