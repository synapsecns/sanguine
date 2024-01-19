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

  return <div>{remainingTime} min</div>
}
