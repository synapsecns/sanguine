import { useTranslations } from 'next-intl'
import Grid from '@/components/ui/tailwind/Grid'
import { SectionContainer, SupportCard } from '@/components/landing/shared'
import SynapseCircuit from '@/components/icons/LandingIcons/SynapseCircuit'
import { SecureIcon } from '@/components/icons/LandingIcons/SecureIcon'
import { ExtensibleIcon } from '@/components/icons/LandingIcons/ExtensibleIcon'
import { GeneralizedIcon } from '@/components/icons/LandingIcons/GeneralizedIcon'

export default function SecuritySection() {
  const t = useTranslations('Landing.SecuritySection')

  return (
    <SectionContainer
      styles="flex flex-wrap gap-8 justify-center max-w-4xl m-auto"
      dataTestId="landing-security-section"
    >
      <div className="flex flex-col items-center gap-8 dark lg:flex-row">
        <SynapseCircuit />
        <div className="max-w-md text-center text-white lg:text-left">
          <h2 className="mb-4 text-4xl font-medium">{t('title')}</h2>
          <div className="text-secondaryTextColor">
            {t('description.part1')}{' '}
            <strong className="font-medium text-white">
              {t('description.strong1')}
            </strong>{' '}
            {t('description.part2')}{' '}
            <strong className="font-medium text-white">
              {t('description.strong2')}
            </strong>{' '}
            {t('description.part3')}
          </div>
        </div>
      </div>
      <Grid cols={{ sm: 1, md: 3 }} gap={8} className="md:p-4">
        <SupportCard header={t('extensibleHeader')} image={<ExtensibleIcon />}>
          {t('extensibleDescription')}
        </SupportCard>
        <SupportCard header={t('secureHeader')} image={<SecureIcon />}>
          {t('secureDescription')}
        </SupportCard>
        <SupportCard
          header={t('generalizedHeader')}
          image={<GeneralizedIcon />}
        >
          {t('generalizedDescription')}
        </SupportCard>
      </Grid>
    </SectionContainer>
  )
}
