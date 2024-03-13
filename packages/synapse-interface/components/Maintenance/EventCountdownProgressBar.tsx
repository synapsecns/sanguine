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

  const { timeRemainingInMinutes, isComplete, isPending } =
    getCountdownTimeStatus(startDate, endDate)

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
        status={status}
        timeRemaining={timeRemainingInMinutes}
      />
    ),
  }
}

export const EventCountdownProgressBar = ({
  eventLabel,
  startDate,
  endDate,
  status,
  timeRemaining,
}: {
  eventLabel: string
  startDate: Date
  endDate: Date
  status: 'idle' | 'pending' | 'complete'
  timeRemaining: number
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

export const getCountdownTimeStatus = (startDate: Date, endDate: Date) => {
  const currentDate = new Date()
  const currentTimeInSeconds = Math.floor(currentDate.getTime() / 1000)

  const startTimeInSeconds = Math.floor(startDate.getTime() / 1000)
  const endTimeInSeconds = Math.floor(endDate.getTime() / 1000)
  const totalTimeInSeconds = endTimeInSeconds - startTimeInSeconds

  const timeElapsedInSeconds = currentTimeInSeconds - startTimeInSeconds
  const timeRemainingInSeconds = endTimeInSeconds - currentTimeInSeconds
  const timeRemainingInMinutes = Math.ceil(timeRemainingInSeconds / 60)

  const isStarted = currentTimeInSeconds >= startTimeInSeconds
  const isComplete = timeRemainingInSeconds <= 0
  const isPending = isStarted && !isComplete

  return {
    currentTimeInSeconds,
    startTimeInSeconds,
    endTimeInSeconds,
    totalTimeInSeconds,
    timeElapsedInSeconds,
    timeRemainingInSeconds,
    timeRemainingInMinutes,
    isStarted,
    isComplete,
    isPending,
  }
}
