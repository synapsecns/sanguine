import { isNull } from 'lodash'

export const getCountdownTimeStatus = (
  startDate: Date,
  endDate: Date | null
) => {
  const currentDate = new Date()

  const currentTimeInSeconds = Math.floor(currentDate.getTime() / 1000)
  const startTimeInSeconds = Math.floor(startDate.getTime() / 1000)

  const isStarted = currentTimeInSeconds >= startTimeInSeconds
  const isIndefinite = isNull(endDate)

  if (isIndefinite) {
    return {
      currentDate,
      currentTimeInSeconds,
      startTimeInSeconds,
      endTimeInSeconds: null,
      totalTimeInSeconds: null,
      totalTimeElapsedInSeconds: null,
      totalTimeRemainingInSeconds: null,
      totalTimeRemainingInMinutes: null,
      daysRemaining: null,
      hoursRemaining: null,
      minutesRemaining: null,
      secondsRemaining: null,
      isStarted,
      isComplete: false,
      isPending: isStarted,
    }
  }

  const { daysRemaining, hoursRemaining, minutesRemaining, secondsRemaining } =
    calculateTimeUntilTarget(endDate)

  const endTimeInSeconds = Math.floor(endDate.getTime() / 1000)
  const totalTimeInSeconds = endTimeInSeconds - startTimeInSeconds

  const totalTimeElapsedInSeconds = currentTimeInSeconds - startTimeInSeconds
  const totalTimeRemainingInSeconds = endTimeInSeconds - currentTimeInSeconds
  const totalTimeRemainingInMinutes = Math.ceil(
    totalTimeRemainingInSeconds / 60
  )

  const isComplete = totalTimeRemainingInSeconds <= 0
  const isPending = isStarted && !isComplete

  return {
    currentDate,
    currentTimeInSeconds,
    startTimeInSeconds,
    endTimeInSeconds,
    totalTimeInSeconds,
    totalTimeElapsedInSeconds,
    totalTimeRemainingInSeconds,
    totalTimeRemainingInMinutes,
    daysRemaining,
    hoursRemaining,
    minutesRemaining,
    secondsRemaining,
    isStarted,
    isComplete,
    isPending,
  }
}
const calculateTimeUntilTarget = (targetDate: Date) => {
  const currentDate = new Date()

  const timeDifference = targetDate.getTime() - currentDate.getTime()

  const isComplete = timeDifference <= 0

  const daysRemaining = Math.floor(timeDifference / (1000 * 60 * 60 * 24))
  const hoursRemaining = Math.ceil(
    (timeDifference % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60)
  )
  const minutesRemaining = Math.floor(
    (timeDifference % (1000 * 60 * 60)) / (1000 * 60)
  )
  const secondsRemaining = Math.floor((timeDifference % (1000 * 60)) / 1000)

  return {
    daysRemaining,
    hoursRemaining,
    minutesRemaining,
    secondsRemaining,
    isComplete,
  }
}
