const oneMinuteInSeconds = 60

export const getEstimatedTimeStatus = (
  currentTime: number,
  initialTime: number,
  estimatedTime: number
) => {
  const elapsedTime = currentTime - initialTime
  const remainingTime = estimatedTime - elapsedTime

  const remainingTimeInMinutes = Math.ceil(remainingTime / oneMinuteInSeconds)
  const elapsedTimeInMinutes = Math.ceil(elapsedTime / oneMinuteInSeconds)

  const isEstimatedTimeReached = remainingTime < 0
  const isStartCheckingTimeReached = remainingTime < oneMinuteInSeconds

  return {
    elapsedTime,
    elapsedTimeInMinutes,
    remainingTime,
    remainingTimeInMinutes,
    isEstimatedTimeReached,
    isStartCheckingTimeReached,
  }
}
