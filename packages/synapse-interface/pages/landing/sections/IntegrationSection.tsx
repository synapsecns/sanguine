import Link from 'next/link'
import Grid from '@/components/ui/tailwind/Grid'
import Card from '@/components/ui/tailwind/Card'
import { SectionContainer } from '../../../components/landing/shared'
import { ORDERED_CHAINS_BY_ID, ChainId, CHAINS_BY_ID } from '@/constants/chains'
import { Chain } from '@/utils/types'
import { getNetworkButtonBorderHover } from '@/styles/chains'

export default function IntegrationSection() {
  const OrderedSupportedNetworks: Chain[] = ORDERED_CHAINS_BY_ID.filter(
    (chainId) => Number(chainId) !== ChainId.TERRA
  ).map((chainId) => {
    return CHAINS_BY_ID[chainId]
  })

  return (
    <SectionContainer dataTestId="landing-integration-section">
      <div
        className={`
          flex flex-col md:flex-row
          items-center justify-center
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
        {OrderedSupportedNetworks.map((network: Chain, index: number) => (
          <NetworkCard
            key={index}
            chainId={network.id}
            chainName={network.name}
            chainImg={network.chainImg.src}
            layer={network.layer}
          />
        ))}
      </Grid>
    </SectionContainer>
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
  const chain = CHAINS_BY_ID[chainId]
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
          ${getNetworkButtonBorderHover(chain?.color)}
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
