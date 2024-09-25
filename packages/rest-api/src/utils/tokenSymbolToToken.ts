import { BRIDGE_MAP } from '../constants/bridgeMap'

export const tokenSymbolToToken = (chain: string, tokenSymbol: string) => {
  const chainData = BRIDGE_MAP[chain]
  if (!chainData) {
    return null
  }
  for (const tokenAddress in chainData) {
    if (chainData[tokenAddress].symbol === tokenSymbol) {
      return {
        address: tokenAddress,
        decimals: chainData[tokenAddress].decimals,
      }
    }
  }
  return null
}
