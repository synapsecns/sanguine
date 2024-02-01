import { zeroAddress } from 'viem'

import { Token } from './types'
import { BRIDGE_MAP } from '@/constants/bridgeMap'

export const getUnderlyingBridgeTokens = (token: Token, chainId: number) => {
  let tokenAddress
  if (token?.addresses[chainId] === zeroAddress) {
    tokenAddress = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'
  } else {
    tokenAddress = token?.addresses[chainId]
  }

  return BRIDGE_MAP[chainId]?.[tokenAddress]?.origin
}
