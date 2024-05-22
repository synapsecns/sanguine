import { useState, useEffect } from 'react'

import { getTimeMinutesFromNow } from '@/utils/getTimeMinutesFromNow'

/**
 * A hook for setting an interval based timer.
 *
 * @param intervalInMs - The duration between intervals, in ms.
 * @param isDisabled - A boolean that determines if timer runs.
 */
export const useIntervalTimer = (
  intervalInMs: number,
  isDisabled?: boolean
) => {
  const [currentTime, setCurrentTime] = useState<number>(
    getTimeMinutesFromNow(0)
  )

  useEffect(() => {
    if (!isDisabled) {
      const interval = setInterval(() => {
        const newCurrentTime = getTimeMinutesFromNow(0)
        setCurrentTime(newCurrentTime)
      }, intervalInMs)

      return () => {
        clearInterval(interval)
      }
    }
  }, [isDisabled])

  return currentTime
}
