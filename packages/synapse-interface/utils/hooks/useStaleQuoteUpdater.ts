import { isNull, isNumber } from 'lodash'
import { useEffect, useRef, useState } from 'react'

import { BridgeQuote } from '@/utils/types'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import { convertUuidToUnix } from '@/utils/convertUuidToUnix'

export const useStaleQuoteUpdater = (
  quote: BridgeQuote,
  refreshQuoteCallback: () => Promise<void>,
  isQuoteLoading: boolean,
  isWalletPending: boolean,
  staleTimeout: number = 15000, // in ms
  autoRefreshDuration: number = 60000 // in ms
) => {
  const [isStale, setIsStale] = useState<boolean>(false)
  const eventListenerRef = useRef<null | (() => void)>(null)
  const timeoutRef = useRef<null | NodeJS.Timeout>(null)
  const autoRefreshIntervalRef = useRef<null | NodeJS.Timeout>(null)
  const autoRefreshStartTimeRef = useRef<null | number>(null)
  const autoRefreshEndTimeRef = useRef<null | NodeJS.Timeout>(null)

  const quoteTime = quote?.id ? convertUuidToUnix(quote?.id) : null
  const isValid = isNumber(quoteTime) && !isNull(quoteTime)

  useIntervalTimer(staleTimeout, !isValid)

  const [moved, reset] = useTrackMouseMove()

  useEffect(() => {
    if (moved && autoRefreshStartTimeRef.current) {
      autoRefreshStartTimeRef.current = null
      reset()
      console.log('reset autorefresh')
    }
  }, [quote])

  // Start auto-refresh logic for 60 seconds
  useEffect(() => {
    if (isValid && !isQuoteLoading && !isWalletPending) {
      // If auto-refresh has not started yet, initialize the start time
      if (autoRefreshStartTimeRef.current === null) {
        autoRefreshStartTimeRef.current = Date.now()
      }

      const elapsedTime = Date.now() - autoRefreshStartTimeRef.current

      console.log('elapsedTime: ', elapsedTime)
      // If autoRefreshDuration hasn't passed, keep auto-refreshing
      if (elapsedTime < autoRefreshDuration) {
        if (timeoutRef.current) clearTimeout(timeoutRef.current)
        if (autoRefreshIntervalRef.current)
          clearInterval(autoRefreshIntervalRef.current)

        autoRefreshIntervalRef.current = setInterval(() => {
          refreshQuoteCallback()
        }, staleTimeout)
      } else {
        // If more than autoRefreshDuration have passed, stop auto-refreshing and switch to mousemove logic
        clearInterval(autoRefreshIntervalRef.current)

        timeoutRef.current = setTimeout(() => {
          eventListenerRef.current = null
          setIsStale(true)

          const newEventListener = () => {
            refreshQuoteCallback()
            eventListenerRef.current = null
            setIsStale(false)
          }

          document.addEventListener('mousemove', newEventListener, {
            once: true,
          })

          eventListenerRef.current = newEventListener
        }, staleTimeout)
      }
    }

    return () => {
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current)
      }
      if (autoRefreshIntervalRef.current) {
        clearInterval(autoRefreshIntervalRef.current)
      }
      if (autoRefreshEndTimeRef.current) {
        clearTimeout(autoRefreshEndTimeRef.current)
      }
      setIsStale(false)
    }
  }, [quote, isQuoteLoading, isWalletPending])

  return isStale
}

export const useTrackMouseMove = (): [boolean, () => void] => {
  const [moved, setMoved] = useState<boolean>(false)

  const onMove = () => setMoved(true)
  const reset = () => setMoved(false)

  useEffect(() => {
    document.addEventListener('mousemove', onMove)

    return () => {
      document.removeEventListener('mousemove', onMove)
    }
  }, [])

  return [moved, reset]
}
