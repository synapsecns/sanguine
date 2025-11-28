export type ChainHash = {
  chainId: number
  hash: string
}

export const marshallChainHash = (chainHash: ChainHash): string => {
  return `${chainHash.hash}:${chainHash.chainId}`
}

export const unmarshallChainHash = (chainHashStr: string): ChainHash => {
  const items = chainHashStr.split(':')
  if (items.length !== 2) {
    throw new Error('Invalid chain hash format')
  }
  return {
    chainId: Number.parseInt(items[1], 10),
    hash: items[0],
  }
}
