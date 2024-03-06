import { useRouter } from 'next/router'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'

const Countdown = () => {
  const router = useRouter()
  const { query, pathname } = router

  const { daysRemaining, hoursRemaining, minutesRemaining, secondsRemaining } =
    calculateTimeUntilTarget()

  console.log(`Days Remaining: ${daysRemaining}`)
  console.log(`Hours Remaining: ${hoursRemaining}`)
  console.log(`Minutes Remaining: ${minutesRemaining}`)
  console.log(`Seconds Remaining: ${secondsRemaining}`)

  return (
    <LandingPageWrapper>
      <section className="flex flex-col items-center justify-center py-24 space-y-8">
        <div className="text-3xl text-white">Countdown to Dencun Upgrade</div>

        <div className="flex space-x-4 text-white">
          <div>
            <div className="text-7xl">{daysRemaining}</div>
            <div className="text-xl">Days</div>
          </div>

          <div>
            <div className="text-7xl">{hoursRemaining}</div>
            <div className="text-xl">Hours</div>
          </div>

          <div>
            <div className="text-7xl">{minutesRemaining}</div>
            <div className="text-xl">Minutes</div>
          </div>

          <div>
            <div className="text-7xl">{secondsRemaining}</div>
            <div className="text-xl">Seconds</div>
          </div>
        </div>
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
