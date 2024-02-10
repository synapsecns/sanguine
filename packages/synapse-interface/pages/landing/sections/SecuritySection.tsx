import Grid from '@tw/Grid'
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
      styles="flex flex-wrap gap-8 justify-center max-w-4xl m-auto"
      dataTestId="landing-security-section"
    >
      <div className="flex flex-col lg:flex-row gap-8 items-center">
        <SynapseCircuit />
        <div className="max-w-md text-white text-center lg:text-left">
          <h2 className="mb-4 text-4xl font-medium">
            Securely connect every blockchain
          </h2>
          <div className="text-secondaryTextColor">
            Synapse is comprised of a{' '}
            <strong className="font-medium text-white">
              cross-chain messaging framework
            </strong>{' '}
            and an{' '}
            <strong className="font-medium text-white">
              economically secure method
            </strong>{' '}
            to reach consensus on the validity of cross-chain transactions,
            enabling developers to build truly native cross-chain apps.
          </div>{' '}
        </div>
      </div>

      <Grid
        cols={{ sm: 1, md: 3 }}
        gap={8}
        className="md:p-4"
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
