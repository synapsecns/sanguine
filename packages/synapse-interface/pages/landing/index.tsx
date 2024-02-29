import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { useAccount } from 'wagmi'

import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { LandingPageContainer } from '@/components/landing/shared'

import HeroSection from '@/components/landing/sections/HeroSection'
import SecuritySection from '@/components/landing/sections/SecuritySection'
import BridgeSection from '@/components/landing/sections/BridgeSection'
import ExplorerSection from '@/components/landing/sections/ExplorerSection'
import IntegrationSection from '@/components/landing/sections/IntegrationSection'
import HowItWorksSection from '@/components/landing/sections/HowItWorksSection'
import UseCasesSection from '@/components/landing/sections/UseCasesSection'
import ResourcesSection from '@/components/landing/sections/ResourcesSection'

import { segmentAnalyticsEvent } from '@/contexts/segmentAnalyticsEvent'

const LandingPage = () => {
  const { address: currentAddress } = useAccount()
  const router = useRouter()

  useEffect(() => {
    segmentAnalyticsEvent(`[Landing] arrives`, {
      address: currentAddress,
      query: router.query,
      pathname: router.pathname,
    })
  }, [])

  return (
    <LandingPageWrapper>
      <LandingPageContainer>
        <HeroSection />
        <SecuritySection />
        <BridgeSection />
        <ExplorerSection />
        <IntegrationSection />
        <HowItWorksSection />
        <UseCasesSection />
        <ResourcesSection />
      </LandingPageContainer>
    </LandingPageWrapper>
  )
}

export default LandingPage
