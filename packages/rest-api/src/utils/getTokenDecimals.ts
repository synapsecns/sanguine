import * as tokensList from '../constants/bridgeable'

export const getTokenDecimals = (
  chainId: number,
  tokenAddress: string
): number => {
  for (const [, token] of Object.entries(tokensList)) {
    if (
      token.addresses[chainId]?.toLowerCase() === tokenAddress.toLowerCase()
    ) {
      const decimals =
        typeof token.decimals === 'object'
          ? token.decimals[chainId]
          : token.decimals
      return decimals
    }
  }
  return 18
}
