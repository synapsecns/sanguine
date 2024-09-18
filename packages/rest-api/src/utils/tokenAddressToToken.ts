import { getAddress } from '@ethersproject/address'

import { BRIDGE_MAP } from '../constants/bridgeMap'
import { ZeroAddress } from '../constants'

export const tokenAddressToToken = (chain: string, tokenAddress: string) => {
  let address
  if (tokenAddress === ZeroAddress) {
    address = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'
  } else {
    address = getAddress(tokenAddress)
  }
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
