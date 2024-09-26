import { BridgeableToken } from '../types'
import * as bridgeableTokens from '../constants/bridgeable'

export const isTokenSupportedOnChain = (
  tokenAddress: string,
  chainId: string
): boolean => {
  const normalizedAddress = tokenAddress.toLowerCase()
  const chainIdNumber = parseInt(chainId, 10)

  return Object.values(bridgeableTokens).some(
    (token: BridgeableToken) =>
      token.addresses[chainIdNumber] !== undefined &&
      token.addresses[chainIdNumber].toLowerCase() === normalizedAddress
  )
}
