import { isNull } from 'lodash'

import { LinearAnimatedProgressBar } from './LinearAnimatedProgressBar'

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
        <div className="flex justify-between px-3 py-2">
          <div>{eventLabel}</div>
          {isIndefinite ? null : <div>{timeRemaining} remaining</div>}
        </div>
        <div className="flex px-1">
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
