import { isNull } from 'lodash'

import { LinearAnimatedProgressBar } from './LinearAnimatedProgressBar'
import { useIntervalTimer } from '@/hooks/useIntervalTimer'

/**
 * Hook to construct an Animated Progress Bar that displays
 * the remaining time left based on given start / end time.
 * Hook also provides status updates on whether Event is pending or complete.
 * If end date is null, progress bar will display an indefinite status.
 *
 * @param eventLabel Message to display with animated progress bar
 * @param startDate Start time of event to track
 * @param endDate End time of event to track
 */
export const useEventCountdownProgressBar = (
  eventLabel: string,
  startDate: Date,
  endDate: Date | null
): {
  isPending: boolean
  isComplete: boolean
  EventCountdownProgressBar: JSX.Element
} => {
  let status: 'idle' | 'pending' | 'complete'

  const { totalTimeRemainingInMinutes, hoursRemaining, isComplete, isPending } =
    getCountdownTimeStatus(startDate, endDate)

  useIntervalTimer(60000, isComplete)

  const timeRemaining: string =
    totalTimeRemainingInMinutes > 90
      ? `${hoursRemaining}h`
      : `${totalTimeRemainingInMinutes}m`

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
  endDate: Date | null
  timeRemaining: string
  status: 'idle' | 'pending' | 'complete'
}) => {
  const isIndefinite = isNull(endDate)

  if (status === 'pending') {
    return (
      <div
        className={`
          flex flex-col bg-[--synapse-surface]
          border border-[--synapse-border] rounded-md
          text-[--synapse-text] text-xs md:text-base
        `}
      >
        <div className="flex justify-between px-3 pt-2">
          <div className="text-sm">{eventLabel}</div>
          {isIndefinite ? null : <div>{timeRemaining} remaining</div>}
        </div>
        <div className="px-1">
          <LinearAnimatedProgressBar
            id="event-countdown-progress-bar"
            startDate={startDate}
            endDate={endDate}
          />
        </div>
      </div>
    )
  } else {
    return null
  }
}

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
