import Link from 'next/link'
import { useTranslations } from 'next-intl'

import Grid from '@/components/ui/tailwind/Grid'
import Button from '@/components/ui/tailwind/Button'
import Card from '@/components/ui/tailwind/Card'
import { SectionContainer } from '@/components/landing/shared'
import { EXPLORER_PATH } from '@/constants/urls'
import {
  getTotalBridgeVolume,
  getTotalTxCount,
  getTotalValueLocked,
  ExplorerQueryStatsResponse,
} from '@/utils/hooks/useExplorerStats'

export default function ExplorerSection() {
  const totalBridgeVolume = getTotalBridgeVolume()
  const totalTxCount = getTotalTxCount()
  const totalValueLocked = getTotalValueLocked()

  const t = useTranslations('Landing.ExplorerSection')

  return (
    <SectionContainer dataTestId="landing-explorer-section" styles="-mx-4">
      <div className="w-full bg-[#1F1D22] pb-6">
        <Grid
          cols={{ sm: 1, md: 2 }}
          className="flex items-center max-w-4xl p-6 m-auto"
        >
          <div>
            <h2 className="mb-3 text-3xl font-medium text-center text-white md:text-left">
              {t('Battle tested')}
            </h2>
            <p className="text-center text-secondaryTextColor md:text-left">
              {t('Synapse has processed')}
            </p>
          </div>
          <div className="hidden col-span-1 text-right md:block">
            <Link href={EXPLORER_PATH} target="_blank">
              <Button
                className={`
                    border border-[#AC8FFF] text-sm
                    px-4 py-3 hover:opacity-75
                  `}
                style={{
                  background:
                    'linear-gradient(310.65deg, rgba(255, 0, 255, 0.2) -17.9%, rgba(172, 143, 255, 0.2) 86.48%)',
                  borderRadius: '10px',
                }}
              >
                {t('Go to Explorer')}
              </Button>
            </Link>
          </div>
        </Grid>

        <Grid
          cols={{ sm: 1, md: 3 }}
          gap={4}
          className="justify-center max-w-4xl m-auto"
        >
          <StatisticsCard
            title={t('Total Value Locked')}
            value={totalValueLocked}
          />
          <StatisticsCard
            title={t('Total Bridge Volume')}
            value={totalBridgeVolume}
          />
          <StatisticsCard title={t('Total TX Count')} value={totalTxCount} />
        </Grid>
      </div>
    </SectionContainer>
  )
}

function StatisticsCard({
  title,
  value,
}: {
  title: string
  value: ExplorerQueryStatsResponse
}) {
  return (
    <Card
      title={title}
      titleClassName="text-white opacity-75"
      className="justify-center p-4 bg-transparent"
      divider={false}
    >
      {value ? (
        <div className={`flex text-3xl font-medium text-white justify-left`}>
          {value}
        </div>
      ) : (
        <div
          className="w-full h-8 bg-slate-700 animate-pulse"
          style={{ maxWidth: '200px' }}
        />
      )}
    </Card>
  )
}
