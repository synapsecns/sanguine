import { useState, useEffect, useMemo, useRef } from 'react'

import { useBridgeState } from '@/slices/bridge/hooks'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { constructStringifiedBridgeSelections } from './useBridgeValidations'
import { BridgeQuote } from '@/utils/types'

export const useConfirmNewBridgePrice = () => {
  const triggerQuoteRef = useRef<any>(null)
  const bpsThreshold = 0.0001 // 1bps

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

  const isPendingConfirmChange =
    hasQuoteOutputChanged &&
    hasSameSelectionsAsPreviousQuote &&
    !hasUserConfirmedChange

  useEffect(() => {
    const validQuotes =
      bridgeQuote?.outputAmount && previousBridgeQuote?.outputAmount

    const hasBridgeModuleChanged =
      bridgeQuote?.bridgeModuleName !==
      (triggerQuoteRef.current?.bridgeModuleName ??
        previousBridgeQuote?.bridgeModuleName)

    const outputAmountDiffMoreThanThreshold = validQuotes
      ? calculateOutputRelativeDifference(
          bridgeQuote,
          triggerQuoteRef.current ?? previousBridgeQuote
        ) > bpsThreshold
      : false

    if (
      validQuotes &&
      hasSameSelectionsAsPreviousQuote &&
      hasBridgeModuleChanged
    ) {
      requestUserConfirmChange(previousBridgeQuote)
    } else if (
      validQuotes &&
      hasSameSelectionsAsPreviousQuote &&
      outputAmountDiffMoreThanThreshold
    ) {
      requestUserConfirmChange(previousBridgeQuote)
    } else {
      resetConfirm()
    }
  }, [bridgeQuote, previousBridgeQuote, hasSameSelectionsAsPreviousQuote])

  const requestUserConfirmChange = (previousQuote: BridgeQuote) => {
    if (!hasQuoteOutputChanged && !hasUserConfirmedChange) {
      triggerQuoteRef.current = previousQuote
      setHasQuoteOutputChanged(true)
    }
    setHasUserConfirmedChange(false)
  }

  const resetConfirm = () => {
    if (hasUserConfirmedChange) {
      triggerQuoteRef.current = null
      setHasQuoteOutputChanged(false)
      setHasUserConfirmedChange(false)
    }
  }

  const onUserAcceptChange = () => {
    triggerQuoteRef.current = null
    setHasUserConfirmedChange(true)
  }

  return {
    isPendingConfirmChange,
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
