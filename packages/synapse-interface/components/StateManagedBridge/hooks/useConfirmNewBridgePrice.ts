import { useState, useEffect, useMemo, useRef } from 'react'

import { useBridgeState } from '@/slices/bridge/hooks'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { constructStringifiedBridgeSelections } from './useBridgeValidations'
import { BridgeQuote } from '@/utils/types'

export const useConfirmNewBridgePrice = () => {
  const quoteRef = useRef<any>(null)

  const [hasQuoteOutputChanged, setHasQuoteOutputChanged] =
    useState<boolean>(false)
  const [hasUserConfirmedChange, setHasUserConfirmedChange] =
    useState<boolean>(false)

  const { bridgeQuote, previousBridgeQuote } = useBridgeQuoteState()
  const { debouncedFromValue, fromToken, toToken, fromChainId, toChainId } =
    useBridgeState()

  const currentBridgeQuoteSelections = useMemo(
    () =>
      constructStringifiedBridgeSelections(
        debouncedFromValue,
        fromChainId,
        fromToken,
        toChainId,
        toToken
      ),
    [debouncedFromValue, fromChainId, fromToken, toChainId, toToken]
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

  const hasSameSelectionsAsPreviousQuote = useMemo(
    () => currentBridgeQuoteSelections === previousBridgeQuoteSelections,
    [currentBridgeQuoteSelections, previousBridgeQuoteSelections]
  )

  useEffect(() => {
    const validQuotes =
      bridgeQuote?.outputAmount && previousBridgeQuote?.outputAmount

    const outputAmountDiffMoreThan1bps = validQuotes
      ? calculateOutputRelativeDifference(
          bridgeQuote,
          quoteRef.current ?? previousBridgeQuote
        ) > 0.0001
      : false

    if (
      validQuotes &&
      outputAmountDiffMoreThan1bps &&
      hasSameSelectionsAsPreviousQuote
    ) {
      requestUserConfirmChange(previousBridgeQuote)
    } else {
      resetConfirm()
    }
  }, [bridgeQuote, previousBridgeQuote, hasSameSelectionsAsPreviousQuote])

  const requestUserConfirmChange = (previousQuote: BridgeQuote) => {
    if (!hasQuoteOutputChanged && !hasUserConfirmedChange) {
      quoteRef.current = previousQuote
      setHasQuoteOutputChanged(true)
    }
    setHasUserConfirmedChange(false)
  }

  const resetConfirm = () => {
    if (hasUserConfirmedChange) {
      quoteRef.current = null
      setHasQuoteOutputChanged(false)
      setHasUserConfirmedChange(false)
    }
  }

  const onUserAcceptChange = () => {
    quoteRef.current = null
    setHasUserConfirmedChange(true)
  }

  return {
    hasSameSelectionsAsPreviousQuote,
    hasQuoteOutputChanged,
    hasUserConfirmedChange,
    onUserAcceptChange,
  }
}

const calculateOutputRelativeDifference = (
  quoteA?: BridgeQuote,
  quoteB?: BridgeQuote
) => {
  if (!quoteA?.outputAmountString || !quoteB?.outputAmountString) return null

  const outputA = parseFloat(quoteA.outputAmountString)
  const outputB = parseFloat(quoteB.outputAmountString)

  return Math.abs(outputA - outputB) / outputB
}
