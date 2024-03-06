import { useRouter } from 'next/router'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import ExternalLinkIcon from '@/components/icons/ExternalLinkIcon'

const Countdown = () => {
  useIntervalTimer(1000)

  const { daysRemaining, hoursRemaining, minutesRemaining, secondsRemaining } =
    calculateTimeUntilTarget()

  /** Remove logs after testing */
  // console.log(`Days Remaining: ${daysRemaining}`)
  // console.log(`Hours Remaining: ${hoursRemaining}`)
  // console.log(`Minutes Remaining: ${minutesRemaining}`)
  // console.log(`Seconds Remaining: ${secondsRemaining}`)

  return (
    <LandingPageWrapper>
      <section className="flex flex-col items-center justify-center py-16 space-y-16 text-center">
        <div className="text-5xl text-white">Countdown to Dencun Upgrade</div>

        <div className="flex space-x-3 text-center text-synapsePurple">
          <div className="inline-block m-auto">
            <div className="w-19 text-7xl">{daysRemaining}</div>
            <div className="text-xl">Days</div>
          </div>

          <div className="inline-block m-auto">
            <div className="w-24 text-7xl">{hoursRemaining}</div>
            <div className="text-xl">Hours</div>
          </div>

          <div className="inline-block m-auto">
            <div className="w-[5.5rem] text-7xl">{minutesRemaining}</div>
            <div className="text-xl">Minutes</div>
          </div>

          <div className="inline-block mt-auto">
            <div className="w-[5.5rem] text-6xl">{secondsRemaining}</div>
            <div className="text-xl">Seconds</div>
          </div>
        </div>

        <a
          href="https://www.ethernow.xyz/"
          target="_blank"
          className="flex flex-wrap items-center justify-center space-x-2 text-xl text-center text-white"
        >
          <div>Upgrade occurring at Beacon Slot 8626176.</div>

          <div className="flex items-center">
            Watch live on Ethernow.
            <ExternalLinkIcon height="20px" width="20px" className="ml-1" />
          </div>
        </a>
      </section>
    </LandingPageWrapper>
  )
}

export default Countdown

const calculateTimeUntilTarget = () => {
  const currentDate = new Date()
  const targetDate = new Date(Date.UTC(2024, 2, 13, 13, 55, 0))

  const timeDifference = targetDate.getTime() - currentDate.getTime()

  const daysRemaining = Math.floor(timeDifference / (1000 * 60 * 60 * 24))
  const hoursRemaining = Math.floor(
    (timeDifference % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60)
  )
    .toString()
    .padStart(2, '0')
  const minutesRemaining = Math.floor(
    (timeDifference % (1000 * 60 * 60)) / (1000 * 60)
  )
    .toString()
    .padStart(2, '0')
  const secondsRemaining = Math.floor((timeDifference % (1000 * 60)) / 1000)
    .toString()
    .padStart(2, '0')

  return {
    daysRemaining,
    hoursRemaining,
    minutesRemaining,
    secondsRemaining,
  }
}
