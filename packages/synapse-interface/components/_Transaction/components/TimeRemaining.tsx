import { useMemo } from 'react'

/**
 * @param isComplete: Transaction has been marked complete
 * @param remainingTime: Time remaining for Transaction, in seconds
 * @param isDelayed: Elapsed Time is over Estimated Time for Transaction
 * @param delayedTime: Delayed Time, in seconds
 * @returns Remaining time in minutes (in seconds if < 1 min) if not delayed.
 * Otherwise, return text representative of delayed transaction.
 */
export const TimeRemaining = ({
  isComplete,
  remainingTime,
  isDelayed,
  delayedTime,
}: {
  isComplete: boolean
  remainingTime: number
  isDelayed: boolean
  delayedTime: number | null
}) => {
  if (isComplete) {
    return <div className="text-sm text-green-400">Complete!</div>
  }

  if (isDelayed) {
    const delayedTimeInMin = Math.floor(delayedTime / 60)
    const absoluteDelayedTime = Math.abs(delayedTimeInMin)
    const showDelayedTime = delayedTimeInMin < -1
    return (
      <div className="text-sm">
        Waiting... {showDelayedTime ? `(${absoluteDelayedTime}m)` : null}
      </div>
    )
  }

  const estTime = useMemo(() => {
    if (remainingTime > 60) {
      return Math.ceil(remainingTime / 60) + 'm remaining'
    } else {
      return remainingTime + 's remaining'
    }
  }, [remainingTime])

  return <div className="text-sm">{estTime}</div>
}
