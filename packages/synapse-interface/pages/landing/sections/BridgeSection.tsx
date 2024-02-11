import Grid from '@tw/Grid'
import {
  SectionContainer,
  SupportCard,
} from '@/components/landing/shared'
import {
  BridgeImage,
  BridgeImageSmall,
} from '@/components/icons/LandingIcons/BridgeImage'
import { DeepLiquidityIcon } from '@/components/icons/LandingIcons/DeepLiquidityIcon'
import { WideSupportIcon } from '@/components/icons/LandingIcons/WideSupportIcon'
import { DeveloperIcon } from '@/components/icons/LandingIcons/DeveloperIcon'

export default function BridgeSection() {
  return (
    <SectionContainer
      styles="flex flex-wrap gap-8 justify-center max-w-4xl m-auto"
      dataTestId="landing-bridge-section"
    >
      <div className="flex flex-col lg:flex-row gap-8 w-full items-center">
        <div className="flex items-center w-full max-w-lg">
          <div>
            <h2 className="mb-4 text-4xl font-medium text-left text-white">
              Powering the most popular bridge
            </h2>
            <p className="font-normal leading-7 text-left text-secondaryTextColor">
              <strong className="font-medium text-white">Synapse Bridge</strong> is
              built on top of the cross-chain infrastructure enabling users to
              seamlessly transfer assets across all blockchains. The Bridge has
              become the most widely-used method to move assets cross-chain,
              offering low cost, fast, and secure bridging.
            </p>
          </div>
        </div>
        <div className="hidden xs:block">
          <BridgeImage />
        </div>
        <div className="block xs:hidden">
          <BridgeImageSmall />
        </div>
      </div>

      <Grid
        cols={{ sm: 1, md: 3 }}
        gap={6}
      >
        <SupportCard header="Deep Liquidity" image={<DeepLiquidityIcon />}>
          Swap native assets using our cross-chain AMM liquidity pools
        </SupportCard>
        <SupportCard header="Wide Support" image={<WideSupportIcon />}>
          Access over 16 different EVM and non-EVM blockchains with more
          integrations coming soon
        </SupportCard>
        <SupportCard header="Dev Friendly" image={<DeveloperIcon />}>
          Easily integrate cross-chain token bridging natively into your
          decentralized application
        </SupportCard>
      </Grid>
    </SectionContainer>
  )
}
