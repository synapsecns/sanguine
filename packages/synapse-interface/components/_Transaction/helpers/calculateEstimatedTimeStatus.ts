/**
 * Calculates remaining time based on given initial, current, and estimated times.
 *
 * @param currentTime - The current time, as a unix timestamp.
 * @param initialTime - The initial time, as a unix timestamp.
 * @param estimatedTime - The estimated duration to calculate, in seconds.
 */
export const calculateEstimatedTimeStatus = (
  currentTime: number,
  initialTime: number,
  estimatedTime: number
) => {
  const elapsedTime = currentTime - initialTime
  const nonNegativeElapsedTime = 0 > elapsedTime ? 0 : elapsedTime
  const remainingTime = estimatedTime - nonNegativeElapsedTime
  const targetTime = initialTime + estimatedTime

  const oneMinuteInSeconds = 60
  const fourHoursInSeconds = 14400

  const isEstimatedTimeReached = remainingTime < 0
  const isCheckTxStatus = remainingTime < oneMinuteInSeconds
  const isCheckTxForRevert = elapsedTime > 30
  const isCheckTxForRefund = elapsedTime > fourHoursInSeconds

  const delayedTime = isEstimatedTimeReached ? remainingTime : null
  const delayedTimeInMin = remainingTime ? Math.floor(remainingTime / 60) : null

  return {
    targetTime,
    elapsedTime: nonNegativeElapsedTime,
    remainingTime,
    delayedTime,
    delayedTimeInMin,
    isEstimatedTimeReached,
    isCheckTxStatus,
    isCheckTxForRevert,
    isCheckTxForRefund,
  }
}
