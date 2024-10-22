import { ethers } from 'ethers'

import { tokenAddressToToken } from '../utils/tokenAddressToToken'
import { getTokenDecimals } from '../utils/getTokenDecimals'
import { BridgeTransactionInfo } from '../types'

export const serializeBridgeInfo = (info: BridgeTransactionInfo) => {
  const { tokenAddress, value, chainID, ...restInfo } = info

  const tokenInfo = tokenAddressToToken(chainID.toString(), tokenAddress)
  const tokenDecimals = getTokenDecimals(chainID, tokenAddress)
  const formattedValue = ethers.utils.formatUnits(value, tokenDecimals)

  return {
    chainID,
    ...restInfo,
    tokenSymbol: tokenInfo ? tokenInfo?.symbol : null,
    formattedValue: `${formattedValue}`,
  }
}
