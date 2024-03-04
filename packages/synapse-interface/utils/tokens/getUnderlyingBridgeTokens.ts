import { zeroAddress } from 'viem'

import type { Token } from '@/utils/types'
import { BRIDGE_MAP } from '@/constants/bridgeMap'
import { ETHEREUM_ADDRESS } from '@/constants'

export const getUnderlyingBridgeTokens = (token: Token, chainId: number) => {
  let tokenAddress
  if (token?.addresses[chainId] === zeroAddress) {
    tokenAddress = ETHEREUM_ADDRESS
  } else {
    tokenAddress = token?.addresses[chainId]
  }

  return BRIDGE_MAP[chainId]?.[tokenAddress]?.origin
}
