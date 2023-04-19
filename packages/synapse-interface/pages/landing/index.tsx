import Link from 'next/link'
import Grid from '@/components/ui/tailwind/Grid'
import Card from '@/components/ui/tailwind/Card'
import Button from '@/components/ui/tailwind/Button'

import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import {
  SynapseCircuit,
  SynapseCircuitSmall,
} from '@/components/icons/LandingIcons/SynapseCircuit'
import {
  BridgeImage,
  BridgeImageSmall,
} from '@/components/icons/LandingIcons/BridgeImage'
import { UniversalMoneyMarketsIcon } from '@/components/icons/LandingIcons/UniversalMoneyMarketsIcon'
import { CrossChainExchangeIcon } from '@/components/icons/LandingIcons/CrossChainExchangeIcon'
import { MultiChainGamingIcon } from '@/components/icons/LandingIcons/MultiChainGamingIcon'
import { HowSynapseWorks } from '@/components/icons/LandingIcons/HowSynapseWorks'
import { ChainId, ORDERED_CHAINS_BY_ID, CHAINS_BY_ID } from '@/constants/chains'
import { Chain } from '@/utils/types'
import {
  DOCS_URL,
  BRIDGE_PATH,
  ANALYTICS_PATH,
  GITHUB_URL,
  SYNAPSE_DOCS_URL,
  MEDIUM_URL,
} from '@/constants/urls'
import {
  getTotalBridgeVolume,
  getTotalPoolVolume,
  getTotalValueLocked,
  ExplorerQueryStatsResponse,
} from '@/utils/hooks/useExplorerStats'
import { getNetworkButtonBorderHover } from '@/utils/styles/networks'

import { LandingPageContainer, SectionContainer, SupportCard } from './shared'
import HeroSection from './sections/HeroSection'
import SecuritySection from './sections/SecuritySection'
import BridgeSection from './sections/BridgeSection'
import ExplorerSection from './sections/ExplorerSection'
import IntegrationSection from './sections/IntegrationSection'

function HowItWorksSection() {
  return (
    <SectionContainer>
      <div className="flex justify-center">
        <div className="w-2/3">
          <div className="mt-12 mb-3 text-4xl text-center text-white">
            How it all works
          </div>
          <div className="mt-6 -mb-12 text-center text-secondaryTextColor">
            Smart contracts from one chain use the Synapse Messaging Router to
            send the message to the destination chain, where a corresponding
            Messaging Router sends it to the destination contract. Messages are
            optimistically verified to ensure security and trustlessness.
          </div>
        </div>
      </div>
      <div className="flex justify-center mx-50">
        <HowSynapseWorks />
      </div>
    </SectionContainer>
  )
}

interface useCaseProp {
  title: string
  image: JSX.Element
  description: string
}

const useCases: useCaseProp[] = [
  {
    title: 'Cross-chain exchange',
    image: <CrossChainExchangeIcon />,
    description: 'Swap any asset on any blockchain using Synapse’s token swaps',
  },
  {
    title: 'Universal money markets',
    image: <UniversalMoneyMarketsIcon />,
    description: 'Swap any asset on any blockchain using Synapse’s token swaps',
  },
  {
    title: 'Multi-chain gaming',
    image: <MultiChainGamingIcon />,
    description: 'Swap any asset on any blockchain using Synapse’s token swaps',
  },
]

function UseCasesSection() {
  return (
    <SectionContainer>
      <div className="flex-col items-center py-6 mt-0 justify-left md:mt-0 lg:mt-0 md:flex md:justify-center lg:flex lg:justify-center ">
        <div className="flex items-center mb-4">
          <span className="mr-6 text-4xl text-white">Use cases</span>
        </div>
        <div className="mt-2 text-left text-secondaryTextColor md:text-center lg:text-center md:mt-0 lg:mt-0">
          Here’s a preview of what you can do using Synapse.
        </div>
      </div>

      <Grid
        cols={{ xs: 1, sm: 1, md: 1, lg: 3 }}
        gap={4}
        className="py-6 pt-6 pb-24 mx-auto place-items-center 2xl:w-3/4"
      >
        {useCases.map((useCase: useCaseProp) => (
          <UseCaseCard
            image={useCase.image}
            title={useCase.title}
            description={useCase.description}
          />
        ))}
      </Grid>
    </SectionContainer>
  )
}

function UseCaseCard({ image, title, description }) {
  return (
    <Card
      className={`
        border border-white border-opacity-10
        bg-opacity-70 bg-[#2F2F2F] px-4 py-4
        md:py-0 w-full md:w-[300px]
      `}
      divider={false}
    >
      <div className="pb-4">
        <div className="flex justify-center mb-2">{image}</div>
        <div className="px-2">
          <div className="text-lg font-medium text-left text-white">
            {title}
          </div>
          <div className="mt-1 text-sm leading-6 text-left text-opacity-75 text-secondaryTextColor">
            {description}
          </div>
        </div>
      </div>
    </Card>
  )
}

function ResourceCard({ title, description, buttonText, linkUrl }) {
  return (
    <Card
      className={`
        text-center rounded-lg border
        border-white border-opacity-10
        bg-[#2F2F2F] bg-opacity-70 py-6 px-6
      `}
      divider={false}
    >
      <div className="text-lg font-medium text-left text-white">{title}</div>
      <div
        className={`
          mt-1 mb-4 text-sm text-left
          text-opacity-75 text-secondaryTextColor
        `}
      >
        {description}
      </div>
      <div className="float-left">
        <Link href={linkUrl} target="_blank">
          <Button
            className={`
            bg-white hover:opacity-75
            text-sm text-[#18171B] font-medium
            px-4 py-3 border rounded-md
            `}
            onClick={() => null}
          >
            {buttonText}
          </Button>
        </Link>
      </div>
    </Card>
  )
}

function ResourcesSection() {
  return (
    <SectionContainer>
      <div
        className={`
          mt-8 mb-4 text-4xl font-medium text-left
          text-white lg:text-center md:text-center
        `}
      >
        Get started now
      </div>
      <div
        className={`
          mb-8 text-left text-secondaryTextColor
          md:text-center lg:text-center
        `}
      >
        Find the resources you need to create integrations with Synapse.
      </div>

      <Grid
        cols={{ sm: 1, md: 1, lg: 3 }}
        gap={6}
        className="py-4 mx-auto lg:px-12 2xl:w-3/4"
      >
        <ResourceCard
          title="References"
          description="Find the resources you need to create integrations with Synapse."
          buttonText="See references"
          linkUrl={GITHUB_URL}
        />
        <ResourceCard
          title="Documentation"
          description="Read a detailed breakdown of our APIs and smart contracts."
          buttonText="Read the docs"
          linkUrl={DOCS_URL}
        />
        <ResourceCard
          title="Tutorials"
          description="Watch interactive tutorials to learn how Synapse works."
          buttonText="Go to tutorials"
          linkUrl={MEDIUM_URL}
        />
      </Grid>
    </SectionContainer>
  )
}

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
