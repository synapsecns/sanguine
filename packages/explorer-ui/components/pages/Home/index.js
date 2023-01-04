import { ANALYTICS_PATH, TRANSACTIONS_PATH } from '@urls'
import { useState } from 'react'
import { Stats } from './Stats'
import { UniversalSearch } from '@components/pages/Home/UniversalSearch'

import { Chart, ChartLoading } from '@components/Chart'
import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { PageLink } from '@components/misc/PageLink'
import Grid from '@components/tailwind/Grid'

import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import {BridgeTransactionTable} from "@components/BridgeTransaction/BridgeTransactionTable";
import _ from "lodash";

export function Home({
  bridgeVolume,
  transactions,
  addresses,
  bridgeVolumeAllTime,
  transactionsAllTime,
  addressesAllTime,
  latestBridgeTransactions,
  popularTokens,
  popularChains,
}) {
  const [chartType, setChartType] = useState('BRIDGEVOLUME')
  const [allTime, setAllTime] = useState(1)

  let data

  if (chartType === 'BRIDGEVOLUME') {
    data = bridgeVolume && bridgeVolume.historicalStatistics.dateResults
  } else if (chartType === 'TRANSACTIONS') {
    data = transactions && transactions.historicalStatistics.dateResults
  } else if (chartType === 'ADDRESSES') {
    data = addresses && addresses.historicalStatistics.dateResults
  }

  let { bridgeTransactions: bridgeTransactionsTable } = latestBridgeTransactions

  bridgeTransactionsTable = _.orderBy(
    bridgeTransactionsTable,
    'fromInfo.time',
    ['desc']
  ).slice(0, 10)

  return (
    <StandardPageContainer>
      <Chart data={data} />
      {bridgeVolume &&
      transactions &&
      addresses &&
      bridgeVolumeAllTime &&
      transactionsAllTime &&
      addressesAllTime ? (
        allTime === 1 ? (
          <Stats
            bridgeVolume={bridgeVolumeAllTime.bridgeAmountStatistic.value}
            transactions={transactionsAllTime.bridgeAmountStatistic.value}
            addresses={addressesAllTime.bridgeAmountStatistic.value}
            setChartType={setChartType}
            allTime={true}
          />
        ) : (
          <Stats
            bridgeVolume={bridgeVolume.historicalStatistics.total}
            transactions={transactions.historicalStatistics.total}
            addresses={addresses.historicalStatistics.total}
            setChartType={setChartType}
            allTime={false}
          />
        )
      ) : (
        <ChartLoading />
      )}
      <HorizontalDivider />
      <Grid cols={{ sm: 1, md: 2, lg: 2 }} gap={4} className="my-3">
        <PageLink text={'See more'} url={ANALYTICS_PATH} external={true} />
        <div className="mt-2 mb-10 text-right">
          <div
            className="text-white text-opacity-50 hover:text-opacity-90 hover:underline cursor-pointer"
            onClick={() => setAllTime(allTime === 0 ? 1 : 0)}
          >
            {allTime === 1 ? 'View 30-day' : ' View all-time'}
          </div>
        </div>
      </Grid>
      <UniversalSearch placeholder="Search by address or transaction" />

      <BridgeTransactionTable queryResult={bridgeTransactionsTable} />
      <HorizontalDivider />
      <PageLink text="See all transactions" url={TRANSACTIONS_PATH} />
    </StandardPageContainer>
  )
}
