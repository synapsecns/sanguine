import { useMemo } from 'react'

export const TimeRemaining = ({
  isComplete,
  remainingTime,
  isDelayed,
}: {
  isComplete: boolean
  remainingTime: number
  isDelayed: boolean
}) => {
  if (isComplete) {
    return <div>Complete!</div>
  }

  if (isDelayed) {
    const delayedTime = Math.floor(remainingTime / 60)
    const absoluteDelayedTime = Math.abs(delayedTime)
    const showDelayedTime = delayedTime < -1
    return (
      <div>
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

  return <div>{estTime}</div>
}
