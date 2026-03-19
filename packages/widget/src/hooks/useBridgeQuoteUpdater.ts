import { isNull, isNumber } from 'lodash'
import { useEffect, useRef } from 'react'

import { type BridgeQuote } from '@/state/slices/bridgeQuote/reducer'
import { calculateTimeBetween } from '@/utils/calculateTimeBetween'
import { useIntervalTimer } from '@/hooks/useIntervalTimer'

/**
 * Refreshes quotes based on selected stale timeout duration.
 * Will refresh quote when browser is active and wallet prompt is not pending.
 */
export const useBridgeQuoteUpdater = (
  quote: BridgeQuote,
  refreshQuoteCallback: () => Promise<void>,
  isQuoteLoading: boolean,
  isWalletPending: boolean,
  staleTimeout: number = 15000 // 15_000ms or 15s
) => {
  const quoteTime = quote?.timestamp
  const isValidQuote = isNumber(quoteTime) && !isNull(quoteTime)
  const currentTime = useIntervalTimer(staleTimeout, !isValidQuote)
  const eventListenerRef = useRef<null | (() => void)>(null)
  const refreshQuoteCallbackRef = useRef(refreshQuoteCallback)

  useEffect(() => {
    refreshQuoteCallbackRef.current = refreshQuoteCallback
  }, [refreshQuoteCallback])

  useEffect(() => {
    if (!isValidQuote || isQuoteLoading || isWalletPending) {
      if (eventListenerRef.current) {
        document.removeEventListener('mousemove', eventListenerRef.current)
        eventListenerRef.current = null
      }

      return
    }

    const timeDifference = calculateTimeBetween(currentTime, quoteTime)
    const isStaleQuote = timeDifference >= staleTimeout

    if (!isStaleQuote) {
      return
    }

    if (eventListenerRef.current) {
      document.removeEventListener('mousemove', eventListenerRef.current)
    }

    const newEventListener = () => {
      refreshQuoteCallbackRef.current()
      eventListenerRef.current = null
    }

    document.addEventListener('mousemove', newEventListener, {
      once: true,
    })

    eventListenerRef.current = newEventListener

    return () => {
      if (eventListenerRef.current) {
        document.removeEventListener('mousemove', eventListenerRef.current)
        eventListenerRef.current = null
      }
    }
  }, [
    currentTime,
    quoteTime,
    isValidQuote,
    isQuoteLoading,
    isWalletPending,
    staleTimeout,
  ])
}
