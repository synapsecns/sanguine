import { BridgeableToken } from '../types'
import * as bridgeableTokens from '../constants/bridgeable'

export const isTokenAddress = (address: string): boolean => {
  const normalizedAddress = address.toLowerCase()

  return Object.values(bridgeableTokens).some((token: BridgeableToken) =>
    Object.values(token.addresses).some(
      (tokenAddress: string) => tokenAddress.toLowerCase() === normalizedAddress
    )
  )
}
