import { SectionContainer } from '../../../components/landing/shared'
import { HowSynapseWorks } from '@/components/icons/LandingIcons/HowSynapseWorks'

export default function HowItWorksSection() {
  return (
    <div className="hidden lg:block">
      <SectionContainer
        dataTestId="landing-how-it-works-section"
        styles="bg-bgBase/10 -mx-4"
      >
        <div className="flex justify-center">
          <div className="max-w-4xl">
            <h2 className="mt-12 mb-3 text-4xl text-center text-white">
              How it all works
            </h2>
            <p className="mt-6 text-center text-secondaryTextColor">
              Smart contracts from one chain use the Synapse Messaging Router to
              send the message to the destination chain, where a corresponding
              Messaging Router sends it to the destination contract. Messages
              are optimistically verified to ensure security and trustlessness.
            </p>
          </div>
        </div>
        <div className="flex justify-center">
          <HowSynapseWorks />
        </div>
      </SectionContainer>
    </div>
  )
}
