import { useRouter } from 'next/router'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'

const Countdown = () => {
  const router = useRouter()
  const { query, pathname } = router

  return <LandingPageWrapper>Countdown to Dencun Upgrade</LandingPageWrapper>
}

export default Countdown
