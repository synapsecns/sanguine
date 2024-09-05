import { useRouter } from 'next/router'
import { useEffect } from 'react'
import deepmerge from 'deepmerge'

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

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

export async function getStaticProps({ locale }) {
  const userMessages = (await import(`../../messages/${locale}.json`)).default
  const defaultMessages = (await import(`../../messages/en-US.json`)).default
  const messages = deepmerge(defaultMessages, userMessages)

  return {
    props: {
      messages,
    },
  }
}

const LandingPage = () => {
  const router = useRouter()

  useEffect(() => {
    segmentAnalyticsEvent(`[Landing] arrives`, {
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
