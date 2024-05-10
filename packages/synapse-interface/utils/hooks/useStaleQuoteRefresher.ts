import { useEffect } from 'react'

import { BridgeQuote } from '@/utils/types'
import { calculateTimeBetween } from '@/utils/time'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

export const useStaleQuoteRefresher = (
  quote: BridgeQuote,
  refreshQuoteCallback: () => Promise<void>,
  staleTimeout: number = 15000
) => {
  const currentTime = useIntervalTimer(staleTimeout, false)

  useEffect(() => {
    if (quote) {
      const timeDifference = calculateTimeBetween(currentTime, quote?.timestamp)

      if (timeDifference >= staleTimeout) {
        console.log('refresh quote')
        refreshQuoteCallback()
      }
    }
  }, [quote, refreshQuoteCallback, currentTime])
}
