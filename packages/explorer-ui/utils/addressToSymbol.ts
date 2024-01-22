import { TOKEN_HASH_MAP } from 'synapse-constants/dist'

export function addressToSymbol({ tokenAddress, chainId }) {
  if (
    tokenAddress === '0x53f7c5869a859F0AeC3D334ee8B4Cf01E3492f21' &&
    chainId === 43114
  ) {
    return 'AVWETH'
  }

  const symbol =
    tokenAddress && chainId && TOKEN_HASH_MAP[chainId][tokenAddress]?.symbol

  return symbol
}
