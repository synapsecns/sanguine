import { useTranslations } from 'next-intl'

import Grid from '@tw/Grid'

interface ApyData {
  fullCompoundedAPY: number
  weeklyAPR: number
  dailyApr: number
  yearlyApr: number
}

interface BaseApyData {
  yearlyCompoundedApy?: number
  dailyApr?: number
  yearlyApr?: number
}

export function ApyTooltipData({
  apyData,
  baseApyData = {},
}: {
  apyData: ApyData
  baseApyData?: BaseApyData
}) {
  const compoundedApy: number = apyData && apyData.fullCompoundedAPY
  const weeklyApr: number = apyData && apyData.weeklyAPR
  const dailyApr: number = weeklyApr && weeklyApr / 7
  const yearlyApr: number = weeklyApr && weeklyApr * 52

  const baseCompoundedApy: number = baseApyData.yearlyCompoundedApy ?? 0
  const baseWeeklyApr: number = (baseApyData.dailyApr ?? 0) * 7
  const baseDailyApr: number = baseApyData.dailyApr ?? 0
  const baseYearlyApr: number = baseApyData.yearlyApr ?? 0

  const t = useTranslations('Pools')

  return (
    apyData && (
      <div className="w-56 pb-2">
        <Grid
          cols={{ xs: 1, sm: 1 }}
          gap={2}
          className="inline-block font-medium"
        >
          <PercentageRow
            title={t('Daily APR')}
            baseApr={baseDailyApr}
            rewardApr={dailyApr}
          />
          <PercentageRow
            title={t('Weekly APR')}
            baseApr={baseWeeklyApr}
            rewardApr={weeklyApr}
          />
          <PercentageRow
            title={t('Yearly APR')}
            baseApr={baseYearlyApr}
            rewardApr={yearlyApr}
          />
          <PercentageRow
            title={t('Yearly APY')}
            baseApr={baseCompoundedApy}
            rewardApr={compoundedApy}
          />
        </Grid>
      </div>
    )
  )
}

const PercentageRow = ({
  title,
  rewardApr,
  baseApr,
}: {
  title: string
  rewardApr: number
  baseApr: number
}) => {
  const totalApr = baseApr + rewardApr

  const t = useTranslations('Pools')

  return (
    <div>
      <div className="text-sm font-normal text-gray-100 ">
        {title}{' '}
        <span className="inline-block float-right pl-4 font-medium">
          {totalApr.toFixed(2)} %
        </span>
      </div>
      {baseApr > 0 && (
        <small className="float-left italic font-normal text-gray-300">
          {rewardApr.toFixed(2)} {t('reward')} + {baseApr.toFixed(2)}{' '}
          {t('base')}
        </small>
      )}
    </div>
  )
}
