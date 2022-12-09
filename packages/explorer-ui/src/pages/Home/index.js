import _ from 'lodash'
import { ANALYTICS_PATH, TRANSACTIONS_PATH } from '@urls'
import { useState } from 'react'

import { PopularChains, PopularTokens } from './Popularity'
import { Stats } from './Stats'
import { UniversalSearch } from './UniversalSearch'
import { LatestBridgeTransactions } from './LatestBridgeTransactions'

import { Chart, ChartLoading } from '@components/Chart'
import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { PageLink } from '@components/misc/PageLink'

import { StandardPageContainer } from '@layouts/StandardPageContainer'
import {
  getAddresses,
  getBridgeVolume,
  getTransactions,
} from 'hooks/getHistoricalStats'

export function Home() {
  const bridgeVolume = getBridgeVolume({ days: 3000 })
  const transactions = getTransactions({ days: 3000 })
  const addresses = getAddresses({ days: 3000 })

  const [chartType, setChartType] = useState('BRIDGEVOLUME')

  let data

  if (chartType === 'BRIDGEVOLUME') {
    data = bridgeVolume && bridgeVolume.historicalStatistics.dateResults
  } else if (chartType === 'TRANSACTIONS') {
    data = transactions && transactions.historicalStatistics.dateResults
  } else if (chartType === 'ADDRESSES') {
    data = addresses && addresses.historicalStatistics.dateResults
  }

  return (
    <StandardPageContainer>
      <Chart data={data} />
      {bridgeVolume && transactions && addresses ? (
        <Stats
          bridgeVolume={bridgeVolume.historicalStatistics.total}
          transactions={transactions.historicalStatistics.total}
          addresses={addresses.historicalStatistics.total}
          setChartType={setChartType}
        />
      ) : (
        <ChartLoading />
      )}
      <HorizontalDivider />
      <PageLink
        text={'View all analytics'}
        url={ANALYTICS_PATH}
        external={true}
      />

      <UniversalSearch />

      <LatestBridgeTransactions />
      <HorizontalDivider />
      <PageLink text="See all transactions" url={TRANSACTIONS_PATH} />

      <PopularTokens />
      <HorizontalDivider />
      <PageLink text="View all tokens" url={TRANSACTIONS_PATH} />

      <PopularChains />
      <HorizontalDivider />
      <PageLink text="View all chains" url={TRANSACTIONS_PATH} />
    </StandardPageContainer>
  )
}
