import { NativeGasAddress, ZeroAddress } from '../constants'
import { BRIDGE_MAP } from '../constants/bridgeMap'

export const tokenAddressToToken = (chain: string, tokenAddress: string) => {
  const chainData = BRIDGE_MAP[chain]
  if (!chainData) {
    return null
  }

  const address = tokenAddress === ZeroAddress ? NativeGasAddress : tokenAddress

  const tokenInfo = chainData[address]
  if (!tokenInfo) {
    return null
  }
  return {
    address: tokenAddress,
    symbol: tokenInfo.symbol,
    decimals: tokenInfo.decimals,
  }
}
