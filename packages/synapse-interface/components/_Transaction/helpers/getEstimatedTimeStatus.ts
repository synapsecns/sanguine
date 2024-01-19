const oneMinuteInSeconds = 60

export const getEstimatedTimeStatus = (
  currentTime: number,
  initialTime: number,
  estimatedTime: number
) => {
  const elapsedTime = currentTime - initialTime // in seconds
  const remainingTime = estimatedTime - elapsedTime
  const endTime = currentTime + estimatedTime

  const remainingTimeInMinutes = Math.ceil(remainingTime / oneMinuteInSeconds) // add additional min for buffer
  const elapsedTimeInMinutes = Math.ceil(elapsedTime / oneMinuteInSeconds)

  const isEstimatedTimeReached = currentTime > endTime

  const startCheckingTimeReached =
    currentTime > currentTime + estimatedTime - oneMinuteInSeconds

  return {
    elapsedTime,
    elapsedTimeInMinutes,
    remainingTime,
    remainingTimeInMinutes,
    isEstimatedTimeReached,
    startCheckingTimeReached,
  }
}
