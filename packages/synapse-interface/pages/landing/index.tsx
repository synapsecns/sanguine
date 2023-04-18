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
import { DOCS_URL, BRIDGE_PATH, ANALYTICS_PATH } from '@/constants/urls'
import {
  getTotalBridgeVolume,
  getTotalPoolVolume,
  getTotalValueLocked,
  ExplorerQueryStatsResponse,
} from '@/utils/hooks/useExplorerStats'
import { getNetworkButtonBorderHover } from '@/utils/styles/networks'

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
      <div
        className={`
          hidden text-left text-secondaryTextColor
          md:text-center lg:text-center md:block lg:block
        `}
      >
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
            h-12 mr-2 px-4 py-3 flex items-center
            text-base border border-white hover:opacity-75
            bg-[#2f2f2f] hover:bg-[#2f2f2f] rounded-lg
            text-center transform-gpu transition-all duration-75
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
            text-base px-4 py-3 hover:opacity-75 rounded-lg
            text-center transform-gpu transition-all duration-75
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
        <div
          className={`
            mt-4 mb-4 text-4xl
            font-medium text-left text-white
          `}
        >
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

      <Grid
        cols={{ sm: 1, md: 3 }}
        gap={12}
        className="py-6 mx-auto mb-12 md:py-12 2xl:w-3/4"
      >
        <SupportCard header="Extensible">
          Synapse’s cross-chain messaging contracts can be deployed across any
          blockchain
        </SupportCard>
        <SupportCard header="Secure">
          Synapse employs an Optimistic security model to ensure integrity of
          cross-chain messages
        </SupportCard>
        <SupportCard header="Generalized">
          Any arbitrary data can be sent across chains including contract calls,
          NFTs, snapshots, and more
        </SupportCard>
      </Grid>
    </SectionContainer>
  )
}

function BridgeSection() {
  return (
    <SectionContainer>
      <Grid cols={{ sm: 1, md: 2 }} gap={10} className="py-6 mx-auto 2xl:w-3/4">
        <div
          className={`
            absolute hidden w-screen -mt-12
            bg-black lg:block bg-opacity-20
            -left-12 -z-10
          `}
          style={{ height: '50rem' }}
        />
        <div className="flex items-center max-w-md">
          <div>
            <div className="mb-4 text-4xl font-medium text-left text-white ">
              Powering the most popular bridge
            </div>
            <div className="font-normal leading-7 text-left text-secondaryTextColor">
              <span className="font-medium text-white">Synapse Bridge</span> is
              built on top of the cross-chain infrastructure enabling users to
              seamlessly transfer assets across all blockchains. The Bridge has
              become the most widely-used method to move assets cross-chain,
              offering low cost, fast, and secure bridging.
            </div>
          </div>
        </div>
        <div className="justify-center hidden md:block">
          <BridgeImage />
        </div>
      </Grid>

      <div className="flex justify-center md:hidden">
        <BridgeImageSmall />
      </div>

      <Grid
        cols={{ sm: 1, md: 3 }}
        gap={12}
        className="py-6 mx-auto mb-6 md:py-12 2xl:w-3/4"
      >
        <SupportCard header="Deep Liquidity">
          Swap native assets using our cross-chain AMM liquidity pools
        </SupportCard>
        <SupportCard header="Wide Support">
          Access over 16 different EVM and non-EVM blockchains with more
          integrations coming soon
        </SupportCard>
        <SupportCard header="Developer Friendly">
          Easily integrate cross-chain token bridging natively into your
          decentralized application
        </SupportCard>
      </Grid>
    </SectionContainer>
  )
}

function ExplorerSection() {
  const totalBridgeVolume = getTotalBridgeVolume()
  const totalPoolVolume = getTotalPoolVolume()
  const totalValueLocked = getTotalValueLocked()

  return (
    <SectionContainer>
      <Grid
        cols={{ sm: 1, md: 2 }}
        gap={4}
        className="flex items-center px-8 py-6 mx-auto md:px-12"
      >
        <div className="max-w-sm mx-auto mt-12 text-left">
          <div className="mb-3 text-3xl font-medium text-white">
            Battle-tested infrastructure
          </div>
          <div className="text-secondaryTextColor ">
            Synapse has processed millions of transactions and tens of billions
            in bridged assets.
          </div>
        </div>
        <div className="hidden col-span-1 text-center md:block">
          <Link href={ANALYTICS_PATH} target="_blank">
            <Button
              className={`
                  border-[#AC8FFF] border text-sm
                  px-4 py-3 hover:opacity-75
                `}
              style={{
                background:
                  'linear-gradient(310.65deg, rgba(255, 0, 255, 0.2) -17.9%, rgba(172, 143, 255, 0.2) 86.48%)',
                borderRadius: '10px',
              }}
            >
              Go to Explorer
            </Button>
          </Link>
        </div>
      </Grid>

      <Grid
        cols={{ sm: 1, md: 2, lg: 3 }}
        gap={4}
        className="max-w-4xl pb-12 mx-auto space-x-0 "
      >
        <StatisticsCard title="Total Value Locked" value={totalValueLocked} />
        <StatisticsCard title="Total Bridge Volume" value={totalBridgeVolume} />
        <StatisticsCard title="Total Pool Volume" value={totalPoolVolume} />
      </Grid>
    </SectionContainer>
  )
}

function IntegrationSection() {
  const OrderedSupportedNetworks: Chain[] = ORDERED_CHAINS_BY_ID.filter(
    (chainId) => Number(chainId) !== ChainId.TERRA
  ).map((chainId) => {
    return CHAINS_BY_ID[chainId]
  })

  return (
    <SectionContainer>
      <div
        className={`
          flex flex-col md:flex-row
          items-center justify-center
          py-6 md:mt-6 lg:mt-6
          lg:flex lg:justify-center
        `}
      >
        <div
          className={`
            mr-6 pr-6 text-3xl text-white
            border-r-0 md:border-r md:border-b-1 md:border-white
          `}
        >
          Widely integrated
        </div>
        <div
          className={`
            mt-2 text-left md:mt-0
            text-secondaryTextColor
            max-w-lg
          `}
        >
          Synapse is widely integrated across the most-used Layer 1 and{' '}
          <br className="hidden md:block" />
          Layer 2 networks for a seamless cross-chain experience.
        </div>
      </div>

      <Grid
        cols={{ xs: 2, sm: 2, md: 3, lg: 5 }}
        gap={4}
        className="py-6 mx-auto md:py-12 lg:py-12 2xl:w-3/4"
      >
        {OrderedSupportedNetworks.map((network: Chain) => (
          <NetworkCard
            chainId={network.id}
            chainName={network.chainName}
            chainImg={network.chainImg.src}
            layer={network.layer}
          />
        ))}
      </Grid>
    </SectionContainer>
  )
}

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

function SupportCard({
  header,
  children,
}: {
  header: string
  children: React.ReactNode
}) {
  return (
    <Card
      title={header}
      titleClassName="text-[1.69rem] font-medium text-white"
      className="px-0 bg-transparent text-secondaryTextColor"
      divider={false}
    >
      {children}
    </Card>
  )
}

function StatisticsCard({
  title,
  value,
}: {
  title: string
  value: ExplorerQueryStatsResponse
}) {
  return (
    <Card
      title={title}
      titleClassName="text-white opacity-75"
      className="justify-center px-2 py-12 bg-transparent"
      divider={false}
    >
      {value ? (
        <div className="flex text-3xl font-medium text-white justify-left">
          {value}
        </div>
      ) : (
        <div
          className="w-full h-8 bg-slate-700 animate-pulse"
          style={{ maxWidth: '200px' }}
        />
      )}
    </Card>
  )
}

function generateNetworkCardHref(chainId) {
  let inputCurrency
  let outputCurrency

  switch (chainId) {
    case ChainId.DOGECHAIN:
      inputCurrency = 'ETH'
      outputCurrency = 'WETH'
      break
    case ChainId.MOONBEAM:
      inputCurrency = 'SYN'
      outputCurrency = 'SYN'
      break
    case ChainId.MOONRIVER:
      inputCurrency = 'SYN'
      outputCurrency = 'SYN'
      break
    default:
      inputCurrency = 'USDC'
      outputCurrency = 'USDC'
  }

  return `/?inputCurrency=${inputCurrency}&outputCurrency=${outputCurrency}&outputChain=${chainId}`
}

function ChainLogo({ src }: { src: string }) {
  return (
    <div className="overflow-visible sm:px-1 md:px-2 lg:px-4">
      <img src={src} className="w-12 overflow-visible rounded-full" />
    </div>
  )
}

interface NetworkCardProps {
  chainId: number
  chainName: string
  layer: number
  chainImg: any
}

function NetworkCard({
  chainId,
  chainName,
  layer,
  chainImg,
}: NetworkCardProps) {
  const href = generateNetworkCardHref(chainId)
  return (
    <Link href={href}>
      <Card
        className={`
          text-center
          border border-[#2F2F2F]
          bg-opacity-0 bg-[#2F2F2F]
          px-0
          py-3 md:py-5
          transform-gpu hover:transition-all duration-75
          ${getNetworkButtonBorderHover(chainId)}
        `}
        divider={false}
      >
        <div className="flex justify-center mt-2 mb-2">
          <ChainLogo src={chainImg} />
        </div>
        <div className="inline-block ">
          <div className="text-lg font-medium text-white">{chainName}</div>
          <div className="mt-1 text-sm text-opacity-75 text-secondaryTextColor">
            Layer {layer}
          </div>
        </div>
      </Card>
    </Link>
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
      </LandingPageContainer>
    </LandingPageWrapper>
  )
}

export default LandingPage
