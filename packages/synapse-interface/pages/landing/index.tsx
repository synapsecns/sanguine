import Grid from '@/components/ui/tailwind/Grid'
import Card from '@/components/ui/tailwind/Card'
import Button from '@/components/ui/tailwind/Button'

import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'

function LandingPageContainer({ children }: { children: React.ReactNode }) {
  return (
    <div
      data-test-id="landing-page-container"
      className="relative px-4 md:px-24"
    >
      {children}
    </div>
  )
}

function SectionContainer({ children }: { children: React.ReactNode }) {
  return <div className="py-6 md:py-12 space-y-[1rem]">{children}</div>
}

function HeroSection() {
  return (
    <SectionContainer>
      <div
        className={`
          text-left text-4xl font-medium text-white
          md:text-[46px] md:text-center pb-4
        `}
      >
        <p>Secure cross-chain</p>
        <p>communication</p>
      </div>
      <div className="hidden text-left text-secondaryTextColor md:text-center lg:text-center md:block lg:block">
        <p>Synapse is the most widely used, extensible, secure cross-</p>
        <p>chain communications network. Build truly cross-chain</p>
        <p>applications using the Synapse Protocol.</p>
      </div>
      <div className="text-left text-secondaryTextColor md:text-center lg:text-center xs:block sm:block md:hidden lg:hidden">
        <p>
          Synapse is the most widely used, extensible, secure cross-chain
          communications network. Build truly cross-chain applications using the
          Synapse Protocol.
        </p>
      </div>
    </SectionContainer>
  )
}
const LandingPage = () => {
  return (
    <LandingPageWrapper>
      <LandingPageContainer>
        <HeroSection />
      </LandingPageContainer>
    </LandingPageWrapper>
  )
}

export default LandingPage
