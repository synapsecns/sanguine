import { AnimatedProgressBar } from '../_Transaction/components/AnimatedProgressBar'

export const useUpgradeProgressBar = (startTime: Date, endTime: Date) => {
  const currentDate = new Date()
  const currentTimeInSeconds = currentDate.getTime() / 1000

  /** Testing countdown, remove after testing */
  const startDate = new Date(Date.UTC(2024, 2, 12, 22, 30, 0))
  const endDate = new Date(Date.UTC(2024, 2, 12, 23, 30, 0))
  /** Testing countdown, remove after testing */

  const startTimeInSeconds = startDate.getTime() / 1000
  const endTimeInSeconds = endDate.getTime() / 1000

  const timeRemaining = endDate.getTime() - currentDate.getTime()
  const timeRemainingInMinutes = Math.floor(timeRemaining / (1000 * 60))

  let status: 'idle' | 'pending' | 'complete'

  const isStarted = currentTimeInSeconds > startTimeInSeconds
  const isComplete = timeRemaining <= 0

  if (isComplete) {
    status = 'complete'
  } else if (isStarted) {
    status = 'pending'
  } else {
    status = 'idle'
  }

  return {
    isStarted,
    isComplete,
    UpgradeProgressBar: (
      <>
        <UpgradeProgressBar
          startTime={startTimeInSeconds}
          endTime={endTimeInSeconds}
          status={status}
          timeRemaining={timeRemainingInMinutes}
        />
      </>
    ),
  }
}

/**
 * Start: 15 min prior to Eth Dencun Upgrade Time @ 3/13/24 13:55 UTC
 * End: 30 min after start of Eth Decun Upgrade Time
 */
// const startDate = new Date(Date.UTC(2024, 2, 13, 13, 40, 0))
// const endDate = new Date(Date.UTC(2024, 2, 12, 14, 25, 0))
export const UpgradeProgressBar = ({
  startTime,
  endTime,
  status,
  timeRemaining,
}: {
  startTime: number
  endTime: number
  status: 'idle' | 'pending' | 'complete'
  timeRemaining: number
}) => {
  // const currentDate = new Date()
  // const currentTimeInSeconds = currentDate.getTime() / 1000

  // /** Testing countdown, remove after testing */
  // const startDate = new Date(Date.UTC(2024, 2, 12, 22, 30, 0))
  // const endDate = new Date(Date.UTC(2024, 2, 12, 23, 30, 0))
  // /** Testing countdown, remove after testing */

  // const startTimeInSeconds = startDate.getTime() / 1000
  // const endTimeInSeconds = endDate.getTime() / 1000

  // const timeRemaining = endDate.getTime() - currentDate.getTime()
  // const timeRemainingInMinutes = Math.floor(timeRemaining / (1000 * 60))

  // const isStarted = currentTimeInSeconds > startTimeInSeconds
  // const isComplete = timeRemaining <= 0

  // console.log('timeRemainingInMinutes:', timeRemainingInMinutes)
  // console.log('currentTimeInSeconds:', currentTimeInSeconds)
  // console.log('startTimeInSeconds:', startTimeInSeconds)
  // console.log('isStarted: ', isStarted)

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
          <div>Dencun upgrade in progress</div>
          <div>{timeRemaining}m remaining</div>
        </div>
        <div className="px-1">
          <AnimatedProgressBar
            id="eth-dencun-countdown"
            startTime={startTime}
            estDuration={45 * 1000} // 45 min
            status={status}
          />
        </div>
      </div>
    )
  } else {
    return null
  }
}
