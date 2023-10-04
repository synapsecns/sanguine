export const isTransactionHash = (s: string): boolean => {
  const valid = /^0x([A-Fa-f0-9]{64})$/.test(s)
  return valid
}
