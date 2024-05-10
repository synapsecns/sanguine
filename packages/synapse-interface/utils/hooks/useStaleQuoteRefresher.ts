import { isNull, isNumber } from 'lodash'
import { useEffect } from 'react'

import { BridgeQuote } from '@/utils/types'
import { calculateTimeBetween } from '@/utils/time'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

export const useStaleQuoteRefresher = (
  quote: BridgeQuote,
  refreshQuoteCallback: () => Promise<void>,
  staleTimeout: number = 15000
) => {
  const quoteTime = quote?.timestamp
  const isValidQuote = isNumber(quoteTime) && !isNull(quoteTime)
  const currentTime = useIntervalTimer(staleTimeout, isValidQuote)

  useEffect(() => {
    if (isValidQuote) {
      const timeDifference = calculateTimeBetween(currentTime, quoteTime)

      if (timeDifference >= staleTimeout) {
        console.log('refresh quote')
        refreshQuoteCallback()
      }
    }
  }, [quote, refreshQuoteCallback, currentTime])
}
