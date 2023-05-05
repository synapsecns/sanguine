import Grid from '@/components/ui/tailwind/Grid'
import {
  SectionContainer,
  SupportCard,
} from '../../../components/landing/shared'
import {
  BridgeImage,
  BridgeImageSmall,
} from '@/components/icons/LandingIcons/BridgeImage'
import { DeepLiquidityIcon } from '@/components/icons/LandingIcons/DeepLiquidityIcon'
import { WideSupportIcon } from '@/components/icons/LandingIcons/WideSupportIcon'
import { DeveloperIcon } from '@/components/icons/LandingIcons/DeveloperIcon'

export default function BridgeSection() {
  return (
    <SectionContainer dataTestId="landing-bridge-section">
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
        gap={6}
        className="mx-auto mb-4 sm:py-6 md:pt-12 md:pb-0 2xl:w-3/4"
      >
        <SupportCard header="Deep Liquidity" image={<DeepLiquidityIcon />}>
          Swap native assets using our cross-chain AMM liquidity pools
        </SupportCard>
        <SupportCard header="Wide Support" image={<WideSupportIcon />}>
          Access over 16 different EVM and non-EVM blockchains with more
          integrations coming soon
        </SupportCard>
        <SupportCard header="Developer Friendly" image={<DeveloperIcon />}>
          Easily integrate cross-chain token bridging natively into your
          decentralized application
        </SupportCard>
      </Grid>
    </SectionContainer>
  )
}
