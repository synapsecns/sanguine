import { useMemo } from 'react'

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
