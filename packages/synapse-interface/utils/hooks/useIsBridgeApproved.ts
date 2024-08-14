import { useEffect, useState } from 'react'
import { zeroAddress } from 'viem'

import { stringToBigInt } from '@/utils/bigint/format'
import { BridgeQuote, Token } from '@/utils/types'

export const useIsBridgeApproved = (
  fromToken: Token | null,
  fromChainId: number,
  bridgeQuote: BridgeQuote | null,
  debouncedFromValue: string
) => {
  const [isApproved, setIsApproved] = useState<boolean>(false)

  useEffect(() => {
    if (fromToken && fromToken.addresses[fromChainId] === zeroAddress) {
      setIsApproved(true)
    } else if (
      fromToken &&
      bridgeQuote?.allowance &&
      stringToBigInt(debouncedFromValue, fromToken.decimals[fromChainId]) <=
        bridgeQuote.allowance
    ) {
      setIsApproved(true)
    } else {
      setIsApproved(false)
    }
  }, [bridgeQuote, fromToken, debouncedFromValue, fromChainId])

  return isApproved
}
