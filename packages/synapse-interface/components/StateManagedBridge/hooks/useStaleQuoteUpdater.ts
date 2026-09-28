import { useEffect, useRef, useState } from 'react'

import { BridgeQuote } from '@/utils/types'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

export const useStaleQuoteUpdater = (
  quote: BridgeQuote,
  refreshQuoteCallback: () => Promise<void>,
  enabled: boolean,
  staleTimeout: number = 15000, // in ms
  autoRefreshDuration: number = 30000 // in ms
) => {
  const [isStale, setIsStale] = useState<boolean>(false)
  const autoRefreshIntervalRef = useRef<null | NodeJS.Timeout>(null)
  const autoRefreshStartTimeRef = useRef<null | number>(null)
  const mouseMoveListenerRef = useRef<null | (() => void)>(null)
  const manualRefreshRef = useRef<null | NodeJS.Timeout>(null)

  useIntervalTimer(staleTimeout, !enabled)

  const [mouseMoved, resetMouseMove] = useTrackMouseMove()

  const clearManualRefreshTimeout = () => {
    if (manualRefreshRef.current) {
      clearTimeout(manualRefreshRef.current)
    }
  }

  const clearAutoRefreshInterval = () => {
    if (autoRefreshIntervalRef.current) {
      clearInterval(autoRefreshIntervalRef.current)
    }
  }

  const clearMouseMoveListener = () => {
    if (mouseMoveListenerRef.current) {
      mouseMoveListenerRef.current = null
    }
  }

  useEffect(() => {
    if (mouseMoved && autoRefreshStartTimeRef.current) {
      autoRefreshStartTimeRef.current = null
      resetMouseMove()
    }
  }, [quote])

  // Start auto-refresh logic for ${autoRefreshDuration}ms seconds
  useEffect(() => {
    if (enabled) {
      // If auto-refresh has not started yet, initialize the start time
      if (autoRefreshStartTimeRef.current === null) {
        autoRefreshStartTimeRef.current = Date.now()
      }

      const elapsedTime = Date.now() - autoRefreshStartTimeRef.current

      // If ${autoRefreshDuration}ms hasn't passed, keep auto-refreshing
      if (elapsedTime < autoRefreshDuration) {
        clearManualRefreshTimeout()
        clearAutoRefreshInterval()

        autoRefreshIntervalRef.current = setInterval(() => {
          refreshQuoteCallback()
        }, staleTimeout)
      } else {
        // If more than ${autoRefreshDuration}ms have passed, stop auto-refreshing and switch to mousemove logic
        clearAutoRefreshInterval()

        manualRefreshRef.current = setTimeout(() => {
          clearMouseMoveListener()
          setIsStale(true)

          const handleMouseMove = () => {
            refreshQuoteCallback()
            clearMouseMoveListener()
            setIsStale(false)
          }

          document.addEventListener('mousemove', handleMouseMove, {
            once: true,
          })

          mouseMoveListenerRef.current = handleMouseMove
        }, staleTimeout)
      }
    }

    return () => {
      clearManualRefreshTimeout()
      clearAutoRefreshInterval()
      setIsStale(false)
    }
  }, [quote, enabled])

  return isStale
}

export const useTrackMouseMove = (): [boolean, () => void] => {
  const [moved, setMoved] = useState<boolean>(false)

  const onMove = () => setMoved(true)
  const onReset = () => setMoved(false)

  useEffect(() => {
    document.addEventListener('mousemove', onMove)

    return () => {
      document.removeEventListener('mousemove', onMove)
    }
  }, [])

  return [moved, onReset]
}
