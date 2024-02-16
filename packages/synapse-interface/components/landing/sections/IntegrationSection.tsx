import Link from 'next/link'
import Grid from '@tw/Grid'
import Card from '@tw/Card'
import { SectionContainer } from '@/components/landing/shared'
import { ORDERED_CHAINS_BY_ID, ChainId, CHAINS_BY_ID } from '@/constants/chains'
import { Chain } from '@/utils/types'
import { getNetworkButtonBorderHover } from '@/styles/chains'
import { useAppDispatch } from '@/store/hooks'
import { setFromChainId } from '@/slices/bridge/reducer'
import { NAVIGATION } from '@/constants/routes'

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
        <h2
          className={`
            mr-6 pr-6 text-3xl text-white
            border-r-0 md:border-r md:border-b-1 md:border-white
          `}
        >
          Widely integrated
        </h2>
        <p
          className={`
            mt-2 text-left md:mt-0
            text-secondaryTextColor
            max-w-lg
          `}
        >
          Synapse is widely integrated across the most-used Layer 1 and{' '}
          <br className="hidden md:block" />
          Layer 2 networks for a seamless cross-chain experience.
        </p>
      </div>

      <Grid
        cols={{ xs: 2, sm: 2, md: 3, lg: 5 }}
        gap={4}
        className="py-8 m-auto max-w-6xl"
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

function ChainImg({ src }: { src: string }) {
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
  const dispatch = useAppDispatch()
  const chain = CHAINS_BY_ID[chainId]

  const handleSelection = () => {
    dispatch(setFromChainId(chainId))
  }

  return (
    <Link href={NAVIGATION.Bridge.path} onClick={handleSelection}>
      <Card
        className={`
          text-center
          border border-white/10 hover:border-white/20
          bg-opacity-0 bg-bgBase/10
          px-0
          py-3 md:py-5
          transform-gpu hover:transition-all duration-75
          ${getNetworkButtonBorderHover(chain?.color)}
        `}
        divider={false}
      >
        <div className="flex justify-center mt-2 mb-2">
          <ChainImg src={chainImg} />
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
