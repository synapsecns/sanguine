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

const LandingPage = () => {
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
