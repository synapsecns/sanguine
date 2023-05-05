import { SectionContainer } from '../../../components/landing/shared'
import { HowSynapseWorks } from '@/components/icons/LandingIcons/HowSynapseWorks'

export default function HowItWorksSection() {
  return (
    <div className="hidden lg:block">
      <SectionContainer
        dataTestId="landing-how-it-works-section"
        styles="absolute-darkened-bg relative"
      >
        <div className="relative flex justify-center">
          <div className="w-2/3">
            <div className="mt-12 mb-3 text-4xl text-center text-white">
              How it all works
            </div>
            <div className="mt-6 text-center text-secondaryTextColor">
              Smart contracts from one chain use the Synapse Messaging Router to
              send the message to the destination chain, where a corresponding
              Messaging Router sends it to the destination contract. Messages
              are optimistically verified to ensure security and trustlessness.
            </div>
          </div>
        </div>
        <div className="relative flex justify-center mx-50">
          <HowSynapseWorks />
        </div>
      </SectionContainer>
    </div>
  )
}
