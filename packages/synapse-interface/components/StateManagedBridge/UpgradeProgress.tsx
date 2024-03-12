import { AnimatedProgressBar } from '../_Transaction/components/AnimatedProgressBar'

export const UpgradeProgress = () => {
  const currentDate = new Date()
  const currentTimeInSeconds = currentDate.getTime() / 1000

  // 15 min prior to Eth Dencun Upgrade Time
  // const startDate = new Date(Date.UTC(2024, 2, 13, 13, 40, 0))

  /** Testing countdown, remove after testing */
  const startDate = new Date(Date.UTC(2024, 2, 12, 22, 30, 0))
  /** Testing countdown, remove after testing */

  const startTimeInSeconds = startDate.getTime() / 1000

  const timeDifference = startDate.getTime() - currentDate.getTime()
  const minutesLeft = Math.floor(timeDifference / (1000 * 60))

  const isComplete = timeDifference <= 0
  const isStarted = currentTimeInSeconds > startTimeInSeconds

  console.log('timeDifference:', timeDifference)
  console.log('currentTimeInSeconds:', currentTimeInSeconds)
  console.log('startTimeInSeconds:', startTimeInSeconds)
  console.log('isStarted: ', isStarted)

  if (isStarted) {
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
          <div>{minutesLeft}m remaining</div>
        </div>
        <div className="px-1">
          <AnimatedProgressBar
            id="eth-dencun-countdown"
            startTime={startTimeInSeconds}
            estDuration={45 * 1000} // 45 min
            status={isComplete ? 'complete' : 'pending'}
          />
        </div>
      </div>
    )
  }
}
