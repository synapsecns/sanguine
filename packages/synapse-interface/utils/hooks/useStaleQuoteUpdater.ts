import { isNull, isNumber } from 'lodash'
import { useEffect, useRef } from 'react'

import { BridgeQuote } from '@/utils/types'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import { convertUuidToUnix } from '@/utils/convertUuidToUnix'

/**
 * Refreshes quotes based on selected stale timeout duration.
 * Will refresh quote when browser is active and wallet prompt is not pending.
 */
export const useStaleQuoteUpdater = (
  quote: BridgeQuote,
  refreshQuoteCallback: () => Promise<void>,
  isQuoteLoading: boolean,
  staleTimeout: number = 15000 // Default 15_000ms or 15s
) => {
  const eventListenerRef = useRef<null | (() => void)>(null)
  const timeoutRef = useRef<null | NodeJS.Timeout>(null)
  const quoteTime = quote?.id ? convertUuidToUnix(quote?.id) : null
  const isValidQuote = isNumber(quoteTime) && !isNull(quoteTime)

  useIntervalTimer(staleTimeout, !isValidQuote)

  useEffect(() => {
    if (isValidQuote && !isQuoteLoading) {
      timeoutRef.current = setTimeout(() => {
        eventListenerRef.current = null

        const newEventListener = () => {
          refreshQuoteCallback()
          eventListenerRef.current = null
        }

        document.addEventListener('mousemove', newEventListener, {
          once: true,
        })

        eventListenerRef.current = newEventListener
      }, staleTimeout)
    }

    return () => {
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current)
      }
    }
  }, [quote, isQuoteLoading])
}
