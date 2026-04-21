import { isNull, isNumber } from 'lodash'
import { useLayoutEffect, useRef } from 'react'

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
  const isQuoteLoadingRef = useRef(isQuoteLoading)
  const isWalletPendingRef = useRef(isWalletPending)
  const refreshQuoteCallbackRef = useRef(refreshQuoteCallback)
  const staleCycleClosedRef = useRef(false)
  const staleCycleTokenRef = useRef(0)
  const staleTimeoutRef = useRef<ReturnType<
    typeof globalThis.setTimeout
  > | null>(null)

  const clearEventListener = () => {
    if (eventListenerRef.current) {
      document.removeEventListener('mousemove', eventListenerRef.current)
      eventListenerRef.current = null
    }
  }

  const clearStaleTimeout = () => {
    if (staleTimeoutRef.current !== null) {
      globalThis.clearTimeout(staleTimeoutRef.current)
      staleTimeoutRef.current = null
    }
  }

  const resetStaleCycle = () => {
    clearStaleTimeout()
    clearEventListener()
  }

  const armStaleListener = (
    activeQuoteRequestId: number,
    staleCycleToken: number
  ) => {
    if (staleCycleClosedRef.current) {
      return
    }

    const newEventListener = () => {
      if (
        staleCycleClosedRef.current ||
        staleCycleTokenRef.current !== staleCycleToken ||
        activeQuoteRequestIdRef.current !== activeQuoteRequestId ||
        isQuoteLoadingRef.current ||
        isWalletPendingRef.current
      ) {
        return
      }

      staleCycleClosedRef.current = true
      clearEventListener()
      const refreshQuotePromise = refreshQuoteCallbackRef.current()
      refreshQuotePromise.catch(() => undefined)
    }

    eventListenerRef.current = newEventListener
    document.addEventListener('mousemove', newEventListener, {
      once: true,
    })
  }

  useLayoutEffect(() => {
    refreshQuoteCallbackRef.current = refreshQuoteCallback
  }, [refreshQuoteCallback])

  useLayoutEffect(() => {
    const activeQuoteRequestId = quote?.requestId
    const isValidQuote =
      isNumber(activeQuoteRequestId) && !isNull(activeQuoteRequestId)

    activeQuoteRequestIdRef.current = isValidQuote ? activeQuoteRequestId : null
    isQuoteLoadingRef.current = isQuoteLoading
    isWalletPendingRef.current = isWalletPending
    staleCycleClosedRef.current = false

    if (!isValidQuote || isQuoteLoading || isWalletPending) {
      resetStaleCycle()
      return () => {
        staleCycleTokenRef.current += 1
        activeQuoteRequestIdRef.current = null
        staleCycleClosedRef.current = false
        resetStaleCycle()
      }
    }

    const staleCycleToken = staleCycleTokenRef.current
    staleTimeoutRef.current = globalThis.setTimeout(() => {
      staleTimeoutRef.current = null

      if (
        staleCycleTokenRef.current !== staleCycleToken ||
        activeQuoteRequestIdRef.current !== activeQuoteRequestId ||
        staleCycleClosedRef.current ||
        isQuoteLoadingRef.current ||
        isWalletPendingRef.current
      ) {
        return
      }

      armStaleListener(activeQuoteRequestId, staleCycleToken)
    }, staleTimeout)

    return () => {
      staleCycleTokenRef.current += 1
      activeQuoteRequestIdRef.current = null
      staleCycleClosedRef.current = false
      resetStaleCycle()
    }
  }, [quote?.requestId, isQuoteLoading, isWalletPending, staleTimeout])
}
