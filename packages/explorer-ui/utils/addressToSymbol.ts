import { TOKEN_HASH_MAP } from 'synapse-constants/dist'

export function addressToSymbol({ tokenAddress, chainId }) {
  const symbol =
    tokenAddress && chainId && TOKEN_HASH_MAP[chainId][tokenAddress]?.symbol

  return symbol
}
