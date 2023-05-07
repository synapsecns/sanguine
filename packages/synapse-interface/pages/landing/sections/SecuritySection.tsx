import Grid from '@/components/ui/tailwind/Grid'
import {
  SectionContainer,
  SupportCard,
} from '../../../components/landing/shared'
import {
  SynapseCircuit,
  SynapseCircuitSmall,
} from '@/components/icons/LandingIcons/SynapseCircuit'
import { SecureIcon } from '@/components/icons/LandingIcons/SecureIcon'
import { ExtensibleIcon } from '@/components/icons/LandingIcons/ExtensibleIcon'
import { GeneralizedIcon } from '@/components/icons/LandingIcons/GeneralizedIcon'

export default function SecuritySection() {
  return (
    <SectionContainer
      styles={`
        flex-wrap items-center
        md:justify-center lg:flex
      `}
      dataTestId="landing-security-section"
    >
      <div className="hidden lg:block">
        <SynapseCircuit />
      </div>
      <div className="flex justify-center pb-6 lg:hidden ">
        <SynapseCircuitSmall />
      </div>

      <div className="max-w-sm md:ml-12 no-mt">
        <div
          className={`
            mb-4 text-4xl font-medium
            text-left text-white
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
        gap={6}
        className="mx-auto mb-4 sm:py-6 md:pt-12 md:pb-0 2xl:w-3/4"
      >
        <SupportCard header="Extensible" image={<ExtensibleIcon />}>
          Synapse’s cross-chain messaging contracts can be deployed across any
          blockchain
        </SupportCard>
        <SupportCard header="Secure" image={<SecureIcon />}>
          Synapse employs an Optimistic security model to ensure integrity of
          cross-chain messages
        </SupportCard>
        <SupportCard header="Generalized" image={<GeneralizedIcon />}>
          Any arbitrary data can be sent across chains including contract calls,
          NFTs, snapshots, and more
        </SupportCard>
      </Grid>
    </SectionContainer>
  )
}
