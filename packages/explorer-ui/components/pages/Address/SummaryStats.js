import {
  PresentationChartLineIcon,
  CurrencyDollarIcon,
  FireIcon,
  ScaleIcon,
  VariableIcon,
} from '@heroicons/react/outline'

import Grid from '@components/tailwind/Grid'

import { ContainerCard } from '@components/ContainerCard'

import { BridgeBehavior } from './BridgeBehavior'
import { StatisticBlock } from './StatisticBlock'

export function SummaryStats({
  address,
  topOriginChainId,
  topDestinationChainId,
}) {
  const iconClassName = 'w-5 h-5 text-purple-500 inline -mt-1'

  return (
    <Grid gap={4} cols={{ sm: 1, md: 1, lg: 3 }} className="my-5">
      <ContainerCard
        className="text-gray-500 sm:col-span-1 md:col-span-1 lg:md:col-span-2 dark:bg-gray-900 dark:text-gray-200 hover:border-purple-500"
        title="Summary Statistics"
        icon={<PresentationChartLineIcon className="w-5 h-5 text-purple-500" />}
      >
        <Grid gap={4} cols={{ sm: 1, md: 2 }} className="pt-4">
          <StatisticBlock
            title="Total Bridge Transactions"
            logo={<FireIcon className={iconClassName} />}
            address={address}
            type="COUNT_TRANSACTIONS"
          />
          <StatisticBlock
            title="Total Bridge Volume"
            logo={<CurrencyDollarIcon className={iconClassName} />}
            address={address}
            type="TOTAL_VOLUME_USD"
            prefix="$"
          />
          <StatisticBlock
            title="Mean Value"
            logo={<ScaleIcon className={iconClassName} />}
            address={address}
            type="MEAN_VOLUME_USD"
            prefix="$"
          />
          <StatisticBlock
            title="Median Value"
            logo={<VariableIcon className={iconClassName} />}
            address={address}
            type="MEDIAN_VOLUME_USD"
            prefix="$"
          />
        </Grid>
      </ContainerCard>
      <BridgeBehavior
        topDestinationChainId={topDestinationChainId}
        topOriginChainId={topOriginChainId}
        address={address}
      />
    </Grid>
  )
}
