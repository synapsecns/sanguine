import { LinearAnimatedProgressBar } from './LinearAnimatedProgressBar'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

export const useEventCountdownProgressBar = (
  eventLabel: string,
  startDate: Date,
  endDate: Date
): {
  isPending: boolean
  isComplete: boolean
  EventCountdownProgressBar: JSX.Element
} => {
  useIntervalTimer(60000)

  const { totalTimeRemainingInMinutes, hoursRemaining, isComplete, isPending } =
    getCountdownTimeStatus(startDate, endDate)

  const timeRemaining: string =
    totalTimeRemainingInMinutes > 90
      ? `${hoursRemaining}h`
      : `${totalTimeRemainingInMinutes}min`

  let status: 'idle' | 'pending' | 'complete'

  if (isComplete) {
    status = 'complete'
  } else if (isPending) {
    status = 'pending'
  } else {
    status = 'idle'
  }

  return {
    isPending,
    isComplete,
    EventCountdownProgressBar: (
      <EventCountdownProgressBar
        eventLabel={eventLabel}
        startDate={startDate}
        endDate={endDate}
        timeRemaining={timeRemaining}
        status={status}
      />
    ),
  }
}

export const EventCountdownProgressBar = ({
  eventLabel,
  startDate,
  endDate,
  timeRemaining,
  status,
}: {
  eventLabel: string
  startDate: Date
  endDate: Date
  timeRemaining: string
  status: 'idle' | 'pending' | 'complete'
}) => {
  if (status === 'pending') {
    return (
      <div
        className={`
          flex flex-col bg-bgLighter mb-3
          border border-surface rounded-md
           text-primary text-xs md:text-base
        `}
      >
        <div className="flex justify-between px-3 py-2">
          <div>{eventLabel}</div>
          <div>{timeRemaining}m remaining</div>
        </div>
        <div className="px-1">
          <LinearAnimatedProgressBar
            id="linear-animated-progress-bar"
            startTime={startTime}
            endTime={endTime}
            status={status}
          />
        </div>
      </div>
    )
  } else {
    return null
  }
}

export const getCountdownTimeStatus = (startDate: Date, endDate: Date) => {
  const currentDate = new Date()

  const { daysRemaining, hoursRemaining, minutesRemaining, secondsRemaining } =
    calculateTimeUntilTarget(endDate)

  const currentTimeInSeconds = Math.floor(currentDate.getTime() / 1000)

  const startTimeInSeconds = Math.floor(startDate.getTime() / 1000)
  const endTimeInSeconds = Math.floor(endDate.getTime() / 1000)
  const totalTimeInSeconds = endTimeInSeconds - startTimeInSeconds

  const totalTimeElapsedInSeconds = currentTimeInSeconds - startTimeInSeconds
  const totalTimeRemainingInSeconds = endTimeInSeconds - currentTimeInSeconds
  const totalTimeRemainingInMinutes = Math.ceil(
    totalTimeRemainingInSeconds / 60
  )

  const isStarted = currentTimeInSeconds >= startTimeInSeconds
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
  const hoursRemaining = Math.floor(
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

