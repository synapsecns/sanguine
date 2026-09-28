import { useEffect, useLayoutEffect, useRef, useState } from 'react'

import { BridgeQuote } from '@/utils/types'

export const useStaleQuoteUpdater = (
  quote: BridgeQuote,
  refreshQuoteCallback: () => Promise<void>,
  enabled: boolean,
  staleTimeout: number = 15000,
  autoRefreshDuration: number = 30000
) => {
  const [isStale, setIsStale] = useState(false)
  const refreshRef = useRef(refreshQuoteCallback)
  const enabledRef = useRef(enabled)
  const refreshInFlightRef = useRef(false)
  const cycleRef = useRef(0)
  const autoRefreshStartTimeRef = useRef<number | null>(null)
  const mouseMovedRef = useRef(false)

  useLayoutEffect(() => {
    refreshRef.current = refreshQuoteCallback
    enabledRef.current = enabled
  }, [refreshQuoteCallback, enabled])

  useEffect(() => {
    const onMove = () => {
      mouseMovedRef.current = true
    }
    document.addEventListener('mousemove', onMove)
    return () => document.removeEventListener('mousemove', onMove)
  }, [])

  useEffect(() => {
    if (mouseMovedRef.current && autoRefreshStartTimeRef.current !== null) {
      autoRefreshStartTimeRef.current = null
      mouseMovedRef.current = false
    }
  }, [quote])

  useEffect(() => {
    const cycle = ++cycleRef.current
    let listener: (() => void) | undefined
    let timer: ReturnType<typeof setTimeout> | undefined
    setIsStale(false)

    const removeListener = () => {
      if (listener) {
        document.removeEventListener('mousemove', listener)
        listener = undefined
      }
    }

    const refresh = () => {
      if (
        !enabledRef.current ||
        cycle !== cycleRef.current ||
        refreshInFlightRef.current
      ) {
        return
      }
      removeListener()
      setIsStale(false)
      refreshInFlightRef.current = true
      Promise.resolve()
        .then(() => {
          if (enabledRef.current && cycle === cycleRef.current) {
            return refreshRef.current()
          }
        })
        .catch(() => undefined)
        .finally(() => {
          refreshInFlightRef.current = false
          if (enabledRef.current && cycle === cycleRef.current) {
            scheduleRefresh()
          }
        })
    }

    const scheduleRefresh = () => {
      const autoRefresh =
        Date.now() - autoRefreshStartTimeRef.current < autoRefreshDuration
      timer = setTimeout(() => {
        if (!enabledRef.current || cycle !== cycleRef.current) return
        if (refreshInFlightRef.current) {
          scheduleRefresh()
        } else if (autoRefresh) {
          refresh()
        } else {
          setIsStale(true)
          listener = refresh
          document.addEventListener('mousemove', listener)
        }
      }, staleTimeout)
    }

    if (enabled) {
      if (autoRefreshStartTimeRef.current === null) {
        autoRefreshStartTimeRef.current = Date.now()
      }
      scheduleRefresh()
    }

    return () => {
      cycleRef.current += 1
      if (timer !== undefined) clearTimeout(timer)
      removeListener()
    }
  }, [quote, enabled, staleTimeout, autoRefreshDuration])

  return isStale
}
