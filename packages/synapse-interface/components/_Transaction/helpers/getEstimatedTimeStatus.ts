/**
 * Provide additional data around elapsed time from an initial start time
 *
 * @param currentTime in seconds, unix
 * @param initialTime in seconds, unix
 * @param estimatedTime in seconds, unix
 */
export const getEstimatedTimeStatus = (
  currentTime: number,
  initialTime: number,
  estimatedTime: number
) => {
  const oneMinuteInSeconds = 60

  const elapsedTimeInSeconds = currentTime - initialTime
  const remainingTimeInSeconds = estimatedTime - elapsedTimeInSeconds

  const remainingTimeInMinutes = Math.ceil(
    remainingTimeInSeconds / oneMinuteInSeconds
  )
  const elapsedTimeInMinutes = Math.ceil(
    remainingTimeInSeconds / oneMinuteInSeconds
  )

  const isEstimatedTimeReached = remainingTimeInSeconds < 0
  const isStartCheckingTimeReached = remainingTimeInSeconds < oneMinuteInSeconds

  return {
    elapsedTimeInSeconds,
    elapsedTimeInMinutes,
    remainingTimeInSeconds,
    remainingTimeInMinutes,
    isEstimatedTimeReached,
    isStartCheckingTimeReached,
  }
}
