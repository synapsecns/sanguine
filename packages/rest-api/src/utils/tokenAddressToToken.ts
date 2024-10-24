import { BRIDGE_MAP } from '../constants/bridgeMap'

export const tokenAddressToToken = (chain: string, tokenAddress: string) => {
  const chainData = BRIDGE_MAP[chain]
  if (!chainData) {
    return null
  }

  const tokenInfo = chainData[tokenAddress]

  if (!tokenInfo) {
    return null
  }

  return {
    address: tokenAddress,
    symbol: tokenInfo.symbol,
    decimals: tokenInfo.decimals,
    swappable: tokenInfo.swappable,
  }
}
