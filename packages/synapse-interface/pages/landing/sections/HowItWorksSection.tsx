import { useTranslations } from 'next-intl'

import { SectionContainer } from '@/components/landing/shared'
import { HowSynapseWorks } from '@/components/icons/LandingIcons/HowSynapseWorks'

export default function HowItWorksSection() {
  const t = useTranslations('Landing.HowItWorksSection')

  return (
    <div className="hidden lg:block">
      <SectionContainer
        dataTestId="landing-how-it-works-section"
        styles="bg-[#121114] -mx-4"
      >
        <div className="flex justify-center">
          <div className="max-w-4xl">
            <h2 className="mt-12 mb-3 text-4xl text-center text-white">
              {t('How it all works')}
            </h2>
            <p className="mt-6 text-center text-secondaryTextColor">
              {t('description')}
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
