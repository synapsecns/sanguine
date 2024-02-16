import { useMemo } from 'react'
import { ExclamationCircleIcon } from '@heroicons/react/outline'

/**
 * @param isComplete: Transaction has been marked complete
 * @param isDelayed: Elapsed Time is over Estimated Time for Transaction
 * @param isReverted: Transaction status queried on chain returned reverted status
 * @param remainingTime: Time remaining for Transaction, in seconds
 * @param delayedTime: Delayed Time, in seconds
 * @returns Remaining time in minutes (in seconds if < 1 min) if not delayed.
 * Otherwise, return text representative of delayed transaction.
 */
export const TimeRemaining = ({
  isComplete,
  isDelayed,
  isReverted,
  remainingTime,
  delayedTime,
}: {
  isComplete: boolean
  isDelayed: boolean
  isReverted: boolean
  remainingTime: number
  delayedTime: number | null
}) => {
  if (isComplete) {
    return <div className="text-sm text-green-400">Complete!</div>
  }

  if (isReverted) {
    return (
      <span className="flex items-center text-sm">
        <ExclamationCircleIcon fill="yellow" className="w-4 h-4"/> <span>Reverted</span>
      </span>
    )
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
