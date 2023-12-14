import { TOKEN_HASH_MAP } from '@constants-new/tokens/index'
// import { TOKEN_HASH_MAP } from '@constants/tokens/basic'
export function addressToSymbol({ tokenAddress, chainId }) {
  const symbol =
    tokenAddress && chainId && TOKEN_HASH_MAP[chainId][tokenAddress]?.symbol

  return symbol
}
