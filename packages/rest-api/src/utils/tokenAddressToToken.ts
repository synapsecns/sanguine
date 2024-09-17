import { getAddress } from '@ethersproject/address'

import { BRIDGE_MAP } from '../constants/bridgeMap'

export const tokenAddressToToken = (chain: string, tokenAddress: string) => {
  const address = getAddress(tokenAddress)
  const chainData = BRIDGE_MAP[chain]
  if (!chainData) {
    return null
  }
  const tokenInfo = chainData[address]
  if (!tokenInfo) {
    return null
  }
  return {
    address,
    symbol: tokenInfo.symbol,
    decimals: tokenInfo.decimals,
  }
}
