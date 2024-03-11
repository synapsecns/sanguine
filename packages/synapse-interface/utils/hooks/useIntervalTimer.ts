import { useState, useEffect } from 'react'

import { getTimeMinutesFromNow } from '@/utils/time'

/**
 * Hook for setting an interval based timer
 *
 * @param intervalInMs number, in ms (1000ms = 1s)
 * returns current time in minutes, unix
 */
export const useIntervalTimer = (intervalInMs: number) => {
  const [currentTime, setCurrentTime] = useState<number>(
    getTimeMinutesFromNow(0)
  )

  /** Update time at set intervals */
  useEffect(() => {
    const interval = setInterval(() => {
      const newCurrentTime = getTimeMinutesFromNow(0)
      setCurrentTime(newCurrentTime)
    }, intervalInMs)

    return () => {
      clearInterval(interval) // Clear the interval when the component unmounts
    }
  }, [])

  return currentTime
}
