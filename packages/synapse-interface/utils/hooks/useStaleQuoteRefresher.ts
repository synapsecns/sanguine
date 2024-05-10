import { isNull, isNumber } from 'lodash'
import { useEffect } from 'react'

import { BridgeQuote } from '@/utils/types'
import { calculateTimeBetween } from '@/utils/time'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

export const useStaleQuoteRefresher = (
  quote: BridgeQuote,
  isLoadingQuote: boolean,
  refreshQuoteCallback: () => Promise<void>,
  staleTimeout: number = 15000
) => {
  const quoteTime = quote?.timestamp
  const isValidQuote = isNumber(quoteTime) && !isNull(quoteTime)
  const currentTime = useIntervalTimer(staleTimeout, !isValidQuote)

  useEffect(() => {
    if (isValidQuote && !isLoadingQuote) {
      const timeDifference = calculateTimeBetween(currentTime, quoteTime)

      console.log('timeDifference: ', timeDifference)
      console.log('staleTimeout: ', staleTimeout)

      if (timeDifference >= staleTimeout) {
        console.log('refresh quote')
        refreshQuoteCallback()
      }
    }
  }, [refreshQuoteCallback])
}
