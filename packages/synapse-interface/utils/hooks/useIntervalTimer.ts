import { useState, useEffect } from 'react'

import { getTimeMinutesFromNow } from '@/utils/time'

/**
 * Hook for setting an interval based timer
 *
 * @param intervalInMs number, in ms (1000ms = 1s)
 * @param isDisabled boolean, determines if we update at intervals
 * returns current time in minutes, unix
 */
export const useIntervalTimer = (
  intervalInMs: number,
  isDisabled?: boolean
) => {
  const [currentTime, setCurrentTime] = useState<number>(
    getTimeMinutesFromNow(0)
  )

  /** Update time at set intervals if not disabled */
  useEffect(() => {
    if (!isDisabled) {
      const interval = setInterval(() => {
        const newCurrentTime = getTimeMinutesFromNow(0)
        setCurrentTime(newCurrentTime)
      }, intervalInMs)

      return () => {
        clearInterval(interval) // Clear the interval when the component unmounts
      }
    }
  }, [isDisabled])

  return currentTime
}
