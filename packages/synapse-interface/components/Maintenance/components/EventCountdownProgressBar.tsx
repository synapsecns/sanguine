import { isNull } from 'lodash'
import { LinearAnimatedProgressBar } from './LinearAnimatedProgressBar'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

/**
 * Hook that 1) constructs an time-tracking animated progress bar using defined start <> end dates,
 * and 2) tracks whether tracked event is started or finished.
 *
 * @param {string} eventLabel - The message to display in the progress bar.
 * @param {Date} startDate - The start date that initiates rendering the progress bar.
 * @param {Date | null} endDate - The end date that removes the progress bar. If null, the progress bar will render indefinitely.
 */
export const useEventCountdownProgressBar = (
  eventLabel: string,
  startDate: Date,
  endDate: Date | null,
  hideProgress?: boolean
): {
  isPending: boolean
  isComplete: boolean
  EventCountdownProgressBar: JSX.Element
} => {
  let status: 'idle' | 'pending' | 'complete'

  const {
    totalTimeRemainingInMinutes,
    daysRemaining,
    hoursRemaining,
    isComplete,
    isPending,
  } = getCountdownTimeStatus(startDate, endDate)

  useIntervalTimer(60000, isComplete)

  const timeRemaining: string =
    daysRemaining > 0
      ? `${daysRemaining}d`
      : totalTimeRemainingInMinutes > 90
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
    EventCountdownProgressBar: !hideProgress && (
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
          flex flex-col bg-bgLighter
          border border-surface rounded-md
           text-primary text-xs md:text-base
        `}
      >
        <div className="flex justify-between px-3 py-2 text-sm">
          <div>{eventLabel}</div>
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
  if (!startDate && !endDate) {
    return {
      currentDate: null,
      currentTimeInSeconds: null,
      startTimeInSeconds: null,
      endTimeInSeconds: null,
      totalTimeInSeconds: null,
      totalTimeElapsedInSeconds: null,
      totalTimeRemainingInSeconds: null,
      totalTimeRemainingInMinutes: null,
      daysRemaining: null,
      hoursRemaining: null,
      minutesRemaining: null,
      secondsRemaining: null,
      isStarted: false,
      isComplete: false,
      isPending: false,
    }
  }

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
