import { InformationCircleIcon } from '@heroicons/react/outline'

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
  const compoundedApy: number = apyData && apyData.fullCompoundedAPY
  const weeklyApr: number = apyData && apyData.weeklyAPR
  const dailyApr: number = weeklyApr && weeklyApr / 7
  const yearlyApr: number = weeklyApr && weeklyApr * 52

  const baseCompoundedApy: number = baseApyData.yearlyCompoundedApy ?? 0
  const baseWeeklyApr: number = (baseApyData.dailyApr ?? 0) * 7
  const baseDailyApr: number = baseApyData.dailyApr ?? 0
  const baseYearlyApr: number = baseApyData.yearlyApr ?? 0

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
          {rewardApr.toFixed(2)} reward + {baseApr.toFixed(2)} base
        </small>
      )}
    </div>
  )
}
