import { useIntervalTimer } from '@/hooks/useIntervalTimer'
import { getCountdownTimeStatus } from '@/utils/getCountdownTimeStatus'
import { EventCountdownProgressBar } from '@/components/Maintenance/EventCountdownProgressBar'

/**
 * A hook that constructs a progress bar with a custom message and countdown timer.
 *
 * @param eventLabel - A description for the tracked event.
 * @param startDate - The start date of the tracked event.
 * @param endDate - The end date of the tracked event.
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

  const { totalTimeRemainingInMinutes, hoursRemaining, isComplete, isPending } =
    getCountdownTimeStatus(startDate, endDate)

  useIntervalTimer(60000, isComplete || !startDate)

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
