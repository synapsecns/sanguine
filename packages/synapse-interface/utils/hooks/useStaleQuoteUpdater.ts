import { isEmpty, isNull, isNumber, isUndefined } from 'lodash'
import { useEffect, useRef } from 'react'

import { BridgeQuote } from '@/utils/types'
import { calculateTimeBetween } from '@/utils/time'
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
  isWalletPending: boolean,
  staleTimeout: number = 15000 // 15_000ms or 15s
) => {
  const eventListenerRef = useRef<null | (() => void)>(null)

  const quoteTime = quote?.id ? convertUuidToUnix(quote?.id) : null
  const isValidQuote = isNumber(quoteTime) && !isNull(quoteTime)

  const currentTime = useIntervalTimer(staleTimeout, !isValidQuote)

  useEffect(() => {
    if (isValidQuote && !isQuoteLoading && !isWalletPending) {
      const timeDifference = calculateTimeBetween(currentTime, quoteTime)
      const isStaleQuote = timeDifference >= staleTimeout

      if (isStaleQuote) {
        if (eventListenerRef.current) {
          document.removeEventListener('mousemove', eventListenerRef.current)
        }

        const newEventListener = () => {
          refreshQuoteCallback()
          eventListenerRef.current = null
        }

        document.addEventListener('mousemove', newEventListener, {
          once: true,
        })

        eventListenerRef.current = newEventListener
      }
    }
  }, [currentTime, staleTimeout])
}
