import { useRouter } from 'next/router'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'

const Countdown = () => {
  const router = useRouter()
  const { query, pathname } = router

  return (
    <LandingPageWrapper>
      <section className="flex justify-center py-24">
        <div className="text-3xl text-white">Countdown to Dencun Upgrade</div>
      </section>
    </LandingPageWrapper>
  )
}

export default Countdown
