import { useState, useEffect, useMemo, useRef } from 'react'

import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { constructStringifiedBridgeSelections } from './useBridgeValidations'

export const useConfirmNewBridgePrice = () => {
  const [hasQuoteOutputChanged, setHasQuoteOutputChanged] =
    useState<boolean>(false)
  const [hasUserConfirmedChange, setHasUserConfirmedChange] =
    useState<boolean>(false)

  const triggeredQuoteRef = useRef<any>(null)

  const { bridgeQuote, previousBridgeQuote } = useBridgeQuoteState()

  const currentBridgeQuoteSelections = useMemo(
    () =>
      constructStringifiedBridgeSelections(
        bridgeQuote?.inputAmountForQuote,
        bridgeQuote?.originChainId,
        bridgeQuote?.originTokenForQuote,
        bridgeQuote?.destChainId,
        bridgeQuote?.destTokenForQuote
      ),
    [bridgeQuote]
  )

  const previousBridgeQuoteSelections = useMemo(
    () =>
      constructStringifiedBridgeSelections(
        previousBridgeQuote?.inputAmountForQuote,
        previousBridgeQuote?.originChainId,
        previousBridgeQuote?.originTokenForQuote,
        previousBridgeQuote?.destChainId,
        previousBridgeQuote?.destTokenForQuote
      ),
    [previousBridgeQuote]
  )

  useEffect(() => {
    const validQuotes =
      bridgeQuote?.outputAmount && previousBridgeQuote?.outputAmount

    const selectionsMatch =
      currentBridgeQuoteSelections === previousBridgeQuoteSelections

    const outputAmountChanged =
      bridgeQuote?.outputAmount !== previousBridgeQuote?.outputAmount

    if (selectionsMatch && validQuotes && outputAmountChanged) {
      // Ref quote that triggered the change
      triggeredQuoteRef.current = bridgeQuote
      setHasQuoteOutputChanged(true)
      setHasUserConfirmedChange(false)
    } else if (
      selectionsMatch &&
      bridgeQuote?.outputAmount === triggeredQuoteRef?.current?.outputAmount
    ) {
      // Maintain status until User confirms ref quote update
      setHasQuoteOutputChanged(true)
    } else {
      setHasQuoteOutputChanged(false)
    }
  }, [
    bridgeQuote,
    previousBridgeQuote,
    currentBridgeQuoteSelections,
    previousBridgeQuoteSelections,
  ])

  return {
    hasQuoteOutputChanged,
    hasUserConfirmedChange,
    setHasUserConfirmedChange,
  }
}
