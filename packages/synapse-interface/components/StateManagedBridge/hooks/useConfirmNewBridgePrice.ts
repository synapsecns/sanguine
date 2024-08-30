import { useState, useEffect, useMemo } from 'react'

import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { constructStringifiedBridgeSelections } from './useBridgeValidations'

export const useConfirmNewBridgePrice = () => {
  const [hasQuoteOutputChanged, setHasQuoteOutputChanged] =
    useState<boolean>(false)
  const [hasUserConfirmedChange, setHasUserConfirmedChange] =
    useState<boolean>(false)

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
    const selectionsMatch =
      currentBridgeQuoteSelections === previousBridgeQuoteSelections

    const outputAmountChanged =
      bridgeQuote?.outputAmount !== previousBridgeQuote?.outputAmount

    setHasQuoteOutputChanged(selectionsMatch && outputAmountChanged)

    if (!selectionsMatch || !outputAmountChanged) {
      setHasUserConfirmedChange(false)
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
