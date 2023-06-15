import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { LandingPageContainer } from '../../components/landing/shared'

import HeroSection from './sections/HeroSection'
import SecuritySection from './sections/SecuritySection'
import BridgeSection from './sections/BridgeSection'
import ExplorerSection from './sections/ExplorerSection'
import IntegrationSection from './sections/IntegrationSection'
import HowItWorksSection from './sections/HowItWorksSection'
import UseCasesSection from './sections/UseCasesSection'
import ResourcesSection from './sections/ResourcesSection'

import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { useAccount } from 'wagmi'

import { useAnalytics } from '@/contexts/AnalyticsProvider'
import { shortenAddress } from '@/utils/shortenAddress'

const LandingPage = () => {
  const { address: currentAddress } = useAccount()
  const router = useRouter()

  const analytics = useAnalytics()

  useEffect(() => {
    analytics.track(
      `[Landing] ${shortenAddress(currentAddress)}arrives`,
      {
        address: currentAddress,
        query: router.query,
        pathname: router.pathname,
      },
      { context: { ip: '0.0.0.0' } }
    )
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
