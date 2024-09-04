import { useState, useEffect, useMemo, useRef, useCallback } from 'react'

import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { constructStringifiedBridgeSelections } from './useBridgeValidations'

export const useConfirmNewBridgePrice = () => {
  const quoteRef = useRef<any>(null)

  const [hasQuoteOutputChanged, setHasQuoteOutputChanged] =
    useState<boolean>(false)
  const [hasUserConfirmedChange, setHasUserConfirmedChange] =
    useState<boolean>(false)

  const { bridgeQuote, previousBridgeQuote } = useBridgeQuoteState()

  const createBridgeSelections = useCallback(
    (quote) =>
      constructStringifiedBridgeSelections(
        quote?.inputAmountForQuote,
        quote?.originChainId,
        quote?.originTokenForQuote,
        quote?.destChainId,
        quote?.destTokenForQuote
      ),
    []
  )

  const currentBridgeQuoteSelections = useMemo(
    () => createBridgeSelections(bridgeQuote),
    [bridgeQuote, createBridgeSelections]
  )

  const previousBridgeQuoteSelections = useMemo(
    () => createBridgeSelections(previousBridgeQuote),
    [previousBridgeQuote, createBridgeSelections]
  )

  const calculateOutputRelativeDifference = useCallback((quoteA, quoteB) => {
    if (!quoteA?.outputAmountString || !quoteB?.outputAmountString) return null

    const outputA = parseFloat(quoteA.outputAmountString)
    const outputB = parseFloat(quoteB.outputAmountString)

    return Math.abs(outputA - outputB) / outputB
  }, [])

  const handleRequestUserConfirmChange = (previousQuote) => {
    if (!hasQuoteOutputChanged && !hasUserConfirmedChange) {
      quoteRef.current = previousQuote
      setHasQuoteOutputChanged(true)
    }
    setHasUserConfirmedChange(false)
  }

  const handleUserAcceptChange = () => {
    quoteRef.current = null
    setHasUserConfirmedChange(true)
  }

  const handleReset = () => {
    if (hasUserConfirmedChange) {
      quoteRef.current = null
      setHasQuoteOutputChanged(false)
      setHasUserConfirmedChange(false)
    }
  }

  useEffect(() => {
    const validQuotes =
      bridgeQuote?.outputAmount && previousBridgeQuote?.outputAmount
    const selectionsMatch =
      currentBridgeQuoteSelections === previousBridgeQuoteSelections

    const outputAmountDiffMoreThan1bps = validQuotes
      ? calculateOutputRelativeDifference(
          bridgeQuote,
          quoteRef.current ?? previousBridgeQuote
        ) > 0.0001
      : false

    console.log('quoteRef.current:', quoteRef.current?.outputAmountString)
    console.log(
      'bridgeQuote?.outputAmountString: ',
      bridgeQuote?.outputAmountString
    )
    console.log(
      'previousBridgeQuote?.outputAmountString:',
      previousBridgeQuote?.outputAmountString
    )
    console.log(
      'relative difference: ',
      calculateOutputRelativeDifference(
        bridgeQuote,
        quoteRef.current ?? previousBridgeQuote
      )
    )
    console.log('outputAmountDiffMoreThan1bps: ', outputAmountDiffMoreThan1bps)

    if (validQuotes && selectionsMatch && outputAmountDiffMoreThan1bps) {
      handleRequestUserConfirmChange(previousBridgeQuote)
    } else {
      handleReset()
    }
  }, [
    bridgeQuote,
    previousBridgeQuote,
    currentBridgeQuoteSelections,
    previousBridgeQuoteSelections,
    calculateOutputRelativeDifference,
  ])

  return {
    hasQuoteOutputChanged,
    hasUserConfirmedChange,
    handleUserAcceptChange,
  }
}
