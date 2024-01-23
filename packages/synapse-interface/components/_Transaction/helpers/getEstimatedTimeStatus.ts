/**
 * Provide additional data around elapsed time from an initial start time
 *
 * @param currentTime in seconds, unix
 * @param initialTime in seconds, unix
 * @param estimatedTime in seconds, unix
 * @returns elapsedTime and remainingTime (in seconds)
 */
export const getEstimatedTimeStatus = (
  currentTime: number,
  initialTime: number,
  estimatedTime: number
) => {
  const elapsedTime = currentTime - initialTime
  const remainingTime = estimatedTime - elapsedTime

  const oneMinuteInSeconds = 60

  const isEstimatedTimeReached = remainingTime < 0
  const isStartCheckingTimeReached = remainingTime < oneMinuteInSeconds

  return {
    elapsedTime,
    remainingTime,
    isEstimatedTimeReached,
    isStartCheckingTimeReached,
  }
}
