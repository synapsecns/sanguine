import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import ExternalLinkIcon from '@/components/icons/ExternalLinkIcon'

const Countdown = () => {
  useIntervalTimer(1000)

  const {
    daysRemaining,
    hoursRemaining,
    minutesRemaining,
    secondsRemaining,
    isComplete,
  } = calculateTimeUntilTarget()

  return (
    <LandingPageWrapper>
      <section className="flex flex-col items-center justify-center pt-24 pb-32 space-y-16 text-center">
        <div className="text-5xl text-white">Countdown to Dencun Upgrade</div>

        <div className="flex space-x-3 text-center text-synapsePurple">
          <div className="inline-block m-auto">
            <div className="w-19 text-7xl">
              {isComplete ? '0' : daysRemaining}
            </div>
            <div className="text-xl">Days</div>
          </div>

          <div className="inline-block m-auto">
            <div className="w-24 text-7xl">
              {isComplete ? '00' : hoursRemaining}
            </div>
            <div className="text-xl">Hours</div>
          </div>

          <div className="inline-block m-auto">
            <div className="w-[5.5rem] text-7xl">
              {isComplete ? '00' : minutesRemaining}
            </div>
            <div className="text-xl">Minutes</div>
          </div>

          <div className="inline-block mt-auto">
            <div className="w-[5.5rem] text-6xl">
              {isComplete ? '00' : secondsRemaining}
            </div>
            <div className="text-xl">Seconds</div>
          </div>
        </div>

        <div className="space-y-8 max-w-[56rem]">
          <div className="flex flex-wrap items-center justify-center space-x-2 text-xl text-center text-white">
            Upgrade occurring at Beacon Slot 8626176.
          </div>

          <a
            href="https://synapse.mirror.xyz/N1dwTpAATINNsCqkXUrJlgYH5szUMBdi-6m8nApDf3I"
            target="_blank"
            className="flex flex-wrap items-center justify-center text-xl text-center text-white hover:text-synapsePurple"
          >
            <div>
              4844 will birth a thousand rollups,
              <div className="inline-block ml-1 text-synapsePurple">
                Synapse Interchain Network
              </div>
            </div>
            <div className="flex items-center justify-center ml-1">
              will bring them together.
              <ExternalLinkIcon
                height="20px"
                width="20px"
                className="mt-px ml-1"
              />
            </div>
          </a>
        </div>
      </section>
    </LandingPageWrapper>
  )
}

export default Countdown

const calculateTimeUntilTarget = () => {
  const currentDate = new Date()
  const currentDay = currentDate.getDate()
  const currentHour = currentDate.getHours()

  let targetDate: Date

  /**
   * Shift target time to actual time after daylight savings
   * Daylight Savings time occurs on March 10th, 2024 @ 2AM on PST, CST, EST
   */
  if (currentDay >= 10 && currentHour >= 2) {
    targetDate = new Date(Date.UTC(2024, 2, 13, 13, 55, 0))
  } else {
    targetDate = new Date(Date.UTC(2024, 2, 13, 12, 55, 0))
  }

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
