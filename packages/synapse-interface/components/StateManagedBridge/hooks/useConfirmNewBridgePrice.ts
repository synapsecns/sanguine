import { useState, useEffect, useMemo, useRef } from 'react'

import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { constructStringifiedBridgeSelections } from './useBridgeValidations'

export const useConfirmNewBridgePrice = () => {
  const [hasQuoteOutputChanged, setHasQuoteOutputChanged] =
    useState<boolean>(false)
  const [hasUserConfirmedChange, setHasUserConfirmedChange] =
    useState<boolean>(false)

  const quoteRef = useRef<any>(null)

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

    const outputAmountDiffMoreThan1bps =
      validQuotes && quoteRef?.current?.outputAmountString
        ? Math.abs(
            parseFloat(bridgeQuote?.outputAmountString) -
              parseFloat(quoteRef?.current?.outputAmountString)
          ) /
            parseFloat(quoteRef?.current?.outputAmountString) >
          0.0001
        : validQuotes
        ? Math.abs(
            parseFloat(bridgeQuote?.outputAmountString) -
              parseFloat(previousBridgeQuote?.outputAmountString)
          ) /
            parseFloat(previousBridgeQuote?.outputAmountString) >
          0.0001
        : false

    // console.log('outputAmountDiffMoreThan1bps:', outputAmountDiffMoreThan1bps)
    // console.log(
    //   'bridgeQuote?.outputAmountString: ',
    //   bridgeQuote?.outputAmountString
    // )
    // console.log(
    //   'previousBridgeQuote?.outputAmountString: ',
    //   previousBridgeQuote?.outputAmountString
    // )

    if (
      validQuotes &&
      selectionsMatch &&
      outputAmountChanged &&
      outputAmountDiffMoreThan1bps
    ) {
      quoteRef.current = bridgeQuote
      setHasQuoteOutputChanged(true)
      setHasUserConfirmedChange(false)
    } else if (
      selectionsMatch &&
      bridgeQuote?.outputAmount === quoteRef?.current?.outputAmount
    ) {
      // Maintain status until User confirms ref quote update
      setHasQuoteOutputChanged(true)
    } else {
      quoteRef.current = null
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
