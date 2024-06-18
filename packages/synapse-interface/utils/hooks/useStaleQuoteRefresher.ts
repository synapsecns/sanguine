import { isNull, isNumber } from 'lodash'
import { useEffect } from 'react'

import { BridgeQuote } from '@/utils/types'
import { calculateTimeBetween } from '@/utils/time'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

/**
 * Refreshes quotes based on selected stale timeout duration.
 * Will refresh quote when browser is active and wallet prompt is not pending.
 */
export const useStaleQuoteRefresher = (
  quote: BridgeQuote,
  refreshQuoteCallback: () => Promise<void>,
  isQuoteLoading: boolean,
  isWalletPending: boolean,
  staleTimeout: number = 15000 // 15_000ms or 15s
) => {
  const quoteTime = quote?.timestamp
  const isValidQuote = isNumber(quoteTime) && !isNull(quoteTime)
  const currentTime = useIntervalTimer(staleTimeout, !isValidQuote)

  useEffect(() => {
    if (isValidQuote && !isQuoteLoading && !isWalletPending) {
      const timeDifference = calculateTimeBetween(currentTime, quoteTime)
      const isStaleQuote = timeDifference >= staleTimeout
      if (isStaleQuote) {
        document.addEventListener('mousemove', refreshQuoteCallback, {
          once: true,
        })
      }
    }
  }, [currentTime, staleTimeout])
}
