import { isNull, isNumber } from 'lodash'
import { useEffect, useRef } from 'react'

import { type BridgeQuote } from '@/state/slices/bridgeQuote/reducer'

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
  const activeQuoteRequestIdRef = useRef<number | null>(null)
  const eventListenerRef = useRef<null | (() => void)>(null)
  const refreshQuoteCallbackRef = useRef(refreshQuoteCallback)
  const staleCycleRef = useRef(false)
  const staleCycleClosedRef = useRef(false)
  const staleDeadlineRef = useRef<number | null>(null)
  const staleTimeoutRef = useRef<number | null>(null)

  const clearEventListener = () => {
    if (eventListenerRef.current) {
      document.removeEventListener('mousemove', eventListenerRef.current)
      eventListenerRef.current = null
    }
  }

  const clearStaleTimeout = () => {
    if (staleTimeoutRef.current !== null) {
      window.clearTimeout(staleTimeoutRef.current)
      staleTimeoutRef.current = null
    }
  }

  const resetStaleCycle = () => {
    staleCycleRef.current = false
    clearStaleTimeout()
    clearEventListener()
  }

  const armStaleListener = (activeQuoteRequestId: number) => {
    if (staleCycleClosedRef.current) {
      return
    }

    staleCycleRef.current = true

    const newEventListener = () => {
      if (
        !staleCycleRef.current ||
        activeQuoteRequestIdRef.current !== activeQuoteRequestId
      ) {
        return
      }

      staleCycleRef.current = false
      staleCycleClosedRef.current = true
      clearEventListener()
      void refreshQuoteCallbackRef.current()
    }

    eventListenerRef.current = newEventListener
    document.addEventListener('mousemove', newEventListener, {
      once: true,
    })
  }

  useEffect(() => {
    refreshQuoteCallbackRef.current = refreshQuoteCallback
  }, [refreshQuoteCallback])

  useEffect(() => {
    return () => {
      activeQuoteRequestIdRef.current = null
      staleCycleClosedRef.current = false
      staleDeadlineRef.current = null
      resetStaleCycle()
    }
  }, [])

  useEffect(() => {
    const activeQuoteRequestId = quote?.requestId
    const isValidQuote =
      isNumber(activeQuoteRequestId) && !isNull(activeQuoteRequestId)

    if (!isValidQuote) {
      activeQuoteRequestIdRef.current = null
      staleCycleClosedRef.current = false
      staleDeadlineRef.current = null
      resetStaleCycle()
      return
    }

    if (activeQuoteRequestIdRef.current !== activeQuoteRequestId) {
      activeQuoteRequestIdRef.current = activeQuoteRequestId
      staleCycleClosedRef.current = false
      staleDeadlineRef.current = Date.now() + staleTimeout
      resetStaleCycle()
    }

    if (isQuoteLoading || isWalletPending) {
      resetStaleCycle()
      return
    }

    if (
      staleCycleClosedRef.current ||
      staleDeadlineRef.current === null ||
      staleCycleRef.current ||
      staleTimeoutRef.current !== null
    ) {
      return
    }

    const remainingStaleTime = staleDeadlineRef.current - Date.now()

    if (remainingStaleTime <= 0) {
      armStaleListener(activeQuoteRequestId)
      return
    }

    staleTimeoutRef.current = window.setTimeout(() => {
      if (activeQuoteRequestIdRef.current !== activeQuoteRequestId) {
        return
      }

      staleTimeoutRef.current = null

      if (staleCycleClosedRef.current) {
        return
      }

      armStaleListener(activeQuoteRequestId)
    }, remainingStaleTime)
  }, [quote?.requestId, isQuoteLoading, isWalletPending, staleTimeout])
}
