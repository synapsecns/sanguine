import { InformationCircleIcon } from '@heroicons/react/outline'

import { fixNumberToPercentageString } from '@/utils/bignumber/format'

import Grid from '@tw/Grid'
import Tooltip from '@tw/Tooltip'

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

export default function ApyTooltip({
  apyData,
  baseApyData = {},
  className,
}: {
  apyData: ApyData
  baseApyData?: BaseApyData
  className?: string
}) {
  const compoundedApy = apyData && apyData.fullCompoundedAPY
  const weeklyApr = apyData && apyData.weeklyAPR
  const dailyApr = weeklyApr && weeklyApr / 7
  const yearlyApr = weeklyApr && weeklyApr * 52

  const baseCompoundedApy = baseApyData.yearlyCompoundedApy ?? 0
  const baseWeeklyApr = (baseApyData.dailyApr ?? 0) * 7
  const baseDailyApr = baseApyData.dailyApr ?? 0
  const baseYearlyApr = baseApyData.yearlyApr ?? 0

  return (
    <Tooltip
      title="Rewards"
      className={className}
      content={
        apyData && (
          <div className="pb-2">
            <Grid
              cols={{ xs: 1, sm: 1 }}
              gap={2}
              className="inline-block font-medium"
            >
              <PercentageRow
                title="Daily APR"
                baseApr={baseDailyApr}
                rewardApr={dailyApr}
              />
              <PercentageRow
                title="Weekly APR"
                baseApr={baseWeeklyApr}
                rewardApr={weeklyApr}
              />
              <PercentageRow
                title="Yearly APR"
                baseApr={baseYearlyApr}
                rewardApr={yearlyApr}
              />
              <PercentageRow
                title="Yearly APY"
                baseApr={baseCompoundedApy}
                rewardApr={compoundedApy}
              />
            </Grid>
          </div>
        )
      }
    >
      <InformationCircleIcon className="w-4 h-4 ml-1 cursor-pointer text-[#252027] fill-bgLighter" />
    </Tooltip>
  )
}

function PercentageRow({ title, rewardApr, baseApr }) {
  const totalApr = baseApr + rewardApr
  return (
    <div>
      <div className="text-sm font-normal text-gray-100 ">
        {title}{' '}
        <span className="inline-block float-right pl-4 font-medium">
          {fixNumberToPercentageString(totalApr)}
        </span>
      </div>
      {baseApr > 0 && (
        <small className="float-left italic font-normal text-gray-300">
          {fixNumberToPercentageString(rewardApr)} reward +{' '}
          {fixNumberToPercentageString(baseApr)} base
        </small>
      )}
    </div>
  )
}
