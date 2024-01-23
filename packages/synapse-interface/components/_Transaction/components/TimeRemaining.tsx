import { useMemo } from 'react'

/**
 * @param isComplete: Transaction has been marked complete
 * @param remainingTime: Time remaining for Transaction, in seconds
 * @param isDelayed: Elapsed Time is over Estimated Time for Transaction
 * @returns Remaining time in minutes (in seconds if < 1 min) if not delayed.
 * Otherwise, return text representative of delayed transaction.
 */
export const TimeRemaining = ({
  isComplete,
  remainingTime,
  isDelayed,
}: {
  isComplete: boolean
  remainingTime: number
  isDelayed: boolean
}) => {
  if (isComplete) return

  if (isDelayed) {
    return <div>Waiting...</div>
  }

  const estTime = useMemo(() => {
    if (remainingTime > 60) {
      return Math.ceil(remainingTime / 60) + ' minutes'
    } else {
      return remainingTime + ' seconds'
    }
  }, [remainingTime])

  return <div>{estTime}</div>
}
