import { useMemo } from 'react'
import ExclamationIcon from '@components/icons/ExclamationIcon'

/**
 * @param isDelayed: Elapsed Time is over Estimated Time for Transaction
 * @param remainingTime: Time remaining for Transaction, in seconds
 * @param delayedTime: Delayed Time, in seconds
 * @param status: Transaction status
 * @returns Remaining time in minutes (in seconds if < 1 min) if not delayed.
 * Otherwise, return text representative of delayed transaction.
 */
export const TimeRemaining = ({
  isDelayed,
  remainingTime,
  delayedTime,
  status,
}: {
  isDelayed: boolean
  remainingTime: number
  delayedTime: number | null
  status: string
}) => {
  const isComplete = status === 'completed'
  const isReverted = status === 'reverted'

  if (isComplete) {
    return <div className="text-sm text-green-400">Complete!</div>
  }

  if (isReverted) {
    return (
      <span className="flex items-center space-x-1 text-sm">
        <ExclamationIcon className="w-4 h-4" /> <span>Reverted</span>
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
