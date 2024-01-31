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
    return
  }

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
