import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { useAccount } from 'wagmi'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import Hero from './Hero'
import ValueProps from './ValueProps'

import Wrapper from '@/components/WipWrapperComponents/Wrapper'

const LandingPage = () => {
  const router = useRouter()

  useEffect(() => {
    segmentAnalyticsEvent(`[Teaser] arrives`, {
      query: router.query,
      pathname: router.pathname,
    })
  }, [])

  return (
    <Wrapper>
      <Hero />
      <ValueProps />
    </Wrapper>
  )
}

export default LandingPage
