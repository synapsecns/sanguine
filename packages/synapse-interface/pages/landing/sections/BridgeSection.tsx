import { useTranslations } from 'next-intl'

import Grid from '@/components/ui/tailwind/Grid'
import { SectionContainer, SupportCard } from '@/components/landing/shared'
import FauxBridge from '@/components/icons/LandingIcons/FauxBridge'
import { DeepLiquidityIcon } from '@/components/icons/LandingIcons/DeepLiquidityIcon'
import { WideSupportIcon } from '@/components/icons/LandingIcons/WideSupportIcon'
import { DeveloperIcon } from '@/components/icons/LandingIcons/DeveloperIcon'

export default function BridgeSection() {
  const t = useTranslations('Landing.BridgeSection')

  return (
    <SectionContainer
      styles="flex flex-wrap gap-8 justify-center max-w-4xl m-auto"
      dataTestId="landing-bridge-section"
    >
      <div className="flex flex-col items-center w-full gap-8 dark lg:flex-row">
        <div className="flex items-center w-full max-w-lg">
          <div>
            <h2 className="mb-4 text-4xl font-medium text-left text-white">
              {t('Powering')}
            </h2>
            <p className="font-normal leading-7 text-left text-secondaryTextColor">
              <strong className="font-medium text-white">
                {t('Synapse Bridge')}
              </strong>{' '}
              {t('Built on top of')}
            </p>
          </div>
        </div>
        <FauxBridge />
      </div>

      <Grid cols={{ sm: 1, md: 3 }} gap={6}>
        <SupportCard header={t('Deep Liquidity')} image={<DeepLiquidityIcon />}>
          {t('Swap native assets')}
        </SupportCard>
        <SupportCard header={t('Wide Support')} image={<WideSupportIcon />}>
          {t('Access')}
        </SupportCard>
        <SupportCard header={t('Developer Friendly')} image={<DeveloperIcon />}>
          {t('Integrate')}
        </SupportCard>
      </Grid>
    </SectionContainer>
  )
}
