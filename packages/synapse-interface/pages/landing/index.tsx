import Grid from '@/components/ui/tailwind/Grid'
import Card from '@/components/ui/tailwind/Card'
import Button from '@/components/ui/tailwind/Button'

import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import {
  SynapseCircuit,
  SynapseCircuitSmall,
} from '@/components/icons/LandingIcons/SynapseCircuit'

import Link from 'next/link'

import { DOCS_URL, BRIDGE_PATH } from '@/constants/urls'

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

function SectionContainer({
  children,
  styles,
}: {
  children: React.ReactNode
  styles?: string
}) {
  return (
    <div
      className={`
        py-6 md:py-12 space-y-[1rem]
        ${styles}
      `}
    >
      {children}
    </div>
  )
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
      <div
        className={`
          text-left text-secondaryTextColor md:text-center
          lg:text-center xs:block sm:block md:hidden lg:hidden
        `}
      >
        <p>
          Synapse is the most widely used, extensible, secure cross-chain
          communications network. Build truly cross-chain applications using the
          Synapse Protocol.
        </p>
      </div>

      <div className="flex justify-center py-4 space-x-2 ">
        <Link
          href={DOCS_URL}
          target="_blank"
          className={`
            h-12 mr-2 px-4 py-3
            flex items-center text-base
            border border-white hover:opacity-75 bg-[#2f2f2f] hover:bg-[#2f2f2f]
            rounded-lg text-center transform-gpu transition-all duration-75
          `}
        >
          <Button
            className="flex items-center justify-center font-medium"
            onClick={() => null}
          >
            Build on Synapse
          </Button>
        </Link>
        <Link
          href={BRIDGE_PATH}
          className={`
            h-12 border-[#AC8FFF] flex items-center border
            text-base px-4 py-3 hover:opacity-75 rounded-lg text-center
            transform-gpu transition-all duration-75
          `}
          style={{
            background:
              'linear-gradient(310.65deg, rgba(255, 0, 255, 0.2) -17.9%, rgba(172, 143, 255, 0.2) 86.48%)',
            borderRadius: '10px',
          }}
        >
          <Button className="font-medium" onClick={() => null}>
            Enter Bridge
          </Button>
        </Link>
      </div>
    </SectionContainer>
  )
}

function SecuritySection() {
  return (
    <SectionContainer
      styles={`
        flex-wrap items-center
        md:justify-center lg:flex
      `}
    >
      <div className="hidden lg:block">
        <SynapseCircuit />
      </div>
      <div className="flex justify-center lg:hidden ">
        <SynapseCircuitSmall />
      </div>

      <div className="max-w-sm md:ml-12">
        <div className="mt-4 mb-4 text-4xl font-medium text-left text-white ">
          Securely connect every blockchain
        </div>
        <div className="font-normal text-left text-secondaryTextColor">
          Synapse is comprised of a{' '}
          <span className="font-medium text-white">
            cross-chain messaging framework
          </span>{' '}
          and an{' '}
          <span className="font-medium text-white">
            economically secure method
          </span>{' '}
          to reach consensus on the validity of cross-chain transactions,
          enabling developers to build truly native cross-chain apps.
        </div>{' '}
      </div>
    </SectionContainer>
  )
}

const LandingPage = () => {
  return (
    <LandingPageWrapper>
      <LandingPageContainer>
        <HeroSection />
        <SecuritySection />
      </LandingPageContainer>
    </LandingPageWrapper>
  )
}

export default LandingPage
