export const generateEntryId = (chainId: number, dbNonce: number) => {
  return `${chainId}-${Number(dbNonce)}`
}
