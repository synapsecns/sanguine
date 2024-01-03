import { TOKEN_HASH_MAP } from '../constants/tokens/index'

export function addressToSymbol({ tokenAddress, chainId }) {
  const symbol =
    tokenAddress && chainId && TOKEN_HASH_MAP[chainId][tokenAddress]?.symbol

  return symbol
}
