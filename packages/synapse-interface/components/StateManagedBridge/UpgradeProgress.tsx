import { AnimatedProgressBar } from '../_Transaction/components/AnimatedProgressBar'

export const UpgradeProgress = () => {
  const currentDate = new Date()
  const currentTimeInSeconds = currentDate.getTime() / 1000

  // 15 min prior to Eth Dencun Upgrade Time
  const startDate = new Date(Date.UTC(2024, 2, 13, 13, 40, 0))
  const startTimeInSeconds = startDate.getTime() / 1000

  const timeDifference = startDate.getTime() - currentDate.getTime()
  const isComplete = timeDifference <= 0
  const isStarted = currentTimeInSeconds < startTimeInSeconds

  if (isStarted) {
    return (
      <div
        className={`
          flex p-2 border border-surface rounded-md bg-tint
          text-primary text-xs md:text-base
        `}
      >
        <AnimatedProgressBar
          id="eth-dencun-countdown"
          startTime={startTimeInSeconds}
          estDuration={45 * 1000} // 45 min
          status={isComplete ? 'complete' : 'pending'}
        />
      </div>
    )
  }
}
