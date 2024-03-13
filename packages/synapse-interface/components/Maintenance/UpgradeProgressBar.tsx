import { useMemo } from 'react'
import { LinearAnimatedProgressBar } from './LinearAnimatedProgressBar'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

export const useUpgradeProgressBar = (
  eventLabel: string,
  startDate: Date,
  endDate: Date
): {
  isPending: boolean
  isComplete: boolean
  UpgradeProgressBar: any
} => {
  useIntervalTimer(60000)
  const currentDate = new Date()
  const currentTimeInSeconds = currentDate.getTime() / 1000
  const startTimeInSeconds = Math.floor(startDate.getTime() / 1000)
  const endTimeInSeconds = Math.floor(endDate.getTime() / 1000)

  const timeRemainingInSeconds = endTimeInSeconds - currentTimeInSeconds
  const timeRemainingInMinutes = Math.ceil(timeRemainingInSeconds / 60)

  const isStarted = currentTimeInSeconds >= startTimeInSeconds
  const isComplete = timeRemainingInSeconds <= 0
  const isPending = isStarted && !isComplete

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
    UpgradeProgressBar: (
      <UpgradeProgressBar
        eventLabel={eventLabel}
        startTime={startTimeInSeconds}
        endTime={endTimeInSeconds}
        status={status}
        timeRemaining={timeRemainingInMinutes}
      />
    ),
  }
}

export const UpgradeProgressBar = ({
  eventLabel,
  startTime,
  endTime,
  status,
  timeRemaining,
}: {
  eventLabel: string
  startTime: number
  endTime: number
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
