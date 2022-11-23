import numeral from 'numeral'
import { useQuery } from '@apollo/client'
import { CurrencyDollarIcon, TrendingUpIcon } from '@heroicons/react/outline'
import Grid from '@tw/Grid'

import { BRIDGE_AMOUNT_STATISTIC } from '@graphql/queries'

import { ContainerCard } from '@components/ContainerCard'
import { InfoDisplay } from './InfoDisplay'
import { InfoLoader } from './InfoLoader'
import { infoBlockIconClassName } from '@constants'

export function SummaryMetricsCard({ chainId }) {
  return (
    <ContainerCard
      title="Summary Metrics"
      icon={<CurrencyDollarIcon className="w-5 h-5 text-purple-500" />}
      subtitle="All Time"
    >
      <SummaryMetrics chainId={chainId} />
    </ContainerCard>
  )
}

function SummaryMetrics({ chainId }) {
  const { data: volume } = useQuery(BRIDGE_AMOUNT_STATISTIC, {
    variables: {
      type: 'TOTAL',
      chainId: Number(chainId),
      duration: 'ALL_TIME',
    },
  })

  const { data: count } = useQuery(BRIDGE_AMOUNT_STATISTIC, {
    variables: {
      type: 'COUNT',
      chainId: Number(chainId),
      duration: 'ALL_TIME',
    },
  })

  const totalVolume = volume?.bridgeAmountStatistic?.USDValue
  const totalCount = count?.bridgeAmountStatistic?.USDValue

  return (
    <Grid gap={4} cols={{ sm: 1 }}>
      <InfoDisplay
        arr={[
          {
            title: 'Total Bridge Volume',
            content:
              totalVolume && totalCount ? (
                <div className="my-2 text-4xl text-center text-slate-200">
                  {numeral(totalVolume).format('$0,0')}
                </div>
              ) : (
                <InfoLoader />
              ),
            logo: <CurrencyDollarIcon className={infoBlockIconClassName} />,
          },
          {
            title: 'Total Transaction Count',
            content:
              totalVolume && totalCount ? (
                <div className="my-2 text-4xl text-center text-slate-200">
                  {numeral(totalCount).format('0,0')}
                </div>
              ) : (
                <InfoLoader />
              ),
            logo: <TrendingUpIcon className={infoBlockIconClassName} />,
          },
        ]}
      />
    </Grid>
  )
}
