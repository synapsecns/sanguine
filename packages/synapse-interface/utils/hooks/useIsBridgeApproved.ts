import { useEffect, useState } from 'react'
import { zeroAddress } from 'viem'

import { stringToBigInt } from '@/utils/bigint/format'
import { useBridgeState } from '@/slices/bridge/hooks'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'

export const useIsBridgeApproved = () => {
  const { debouncedFromValue, fromChainId, fromToken } = useBridgeState()
  const { bridgeQuote } = useBridgeQuoteState()

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
