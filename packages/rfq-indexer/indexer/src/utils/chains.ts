export const chainIdToName: { [key: number]: string } = {
  1: 'ethereum',
  10: 'optimism',
  42161: 'arbitrum',
  8453: 'base',
  81457: 'blast',
  534352: 'scroll',
  59144: 'linea',
  56: 'bnb',
  480: 'worldchain',
  130: 'unichain',
  80094: 'berachain',
}

export const getChainName = (chainId: number): string => {
  return chainIdToName[chainId] || 'unknown'
}
