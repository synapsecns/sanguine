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
import Grid from '@tw/Grid'

import { StandardPageContainer } from '@layouts/StandardPageContainer'
import {
  getAddresses,
  getBridgeVolume,
  getTransactions,
} from 'hooks/getHistoricalStats'

export function Home() {
  const bridgeVolume = getBridgeVolume({ days: 30 })
  const transactions = getTransactions({ days: 30 })
  const addresses = getAddresses({ days: 30 })


  const bridgeAllTimeVolume = getBridgeVolume({ days: 3000 })
  const transactionsAllTime = getTransactions({ days: 3000 })
  const addressesAllTime = getAddresses({ days: 3000 })

  const [chartType, setChartType] = useState('BRIDGEVOLUME')
  const [allTime, setAllTime] = useState(0)


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
      {bridgeVolume && transactions && addresses ?
        allTime === 0 ?
          <Stats
            bridgeVolume={bridgeAllTimeVolume.historicalStatistics.total}
            transactions={transactionsAllTime.historicalStatistics.total}
            addresses={addressesAllTime.historicalStatistics.total}
            setChartType={setChartType}
            allTime={true}
          /> :
          <Stats
            bridgeVolume={bridgeVolume.historicalStatistics.total}
            transactions={transactions.historicalStatistics.total}
            addresses={addresses.historicalStatistics.total}
            setChartType={setChartType}
            allTime={false}
          />
        : (
          <ChartLoading />
        )}
      <HorizontalDivider />
      <Grid cols={{ sm: 1, md: 2, lg: 2 }} gap={4} className="my-3">

        <PageLink
          text={'See more'}
          url={ANALYTICS_PATH}
          external={true}
        />
          <div className="mt-2 mb-10 text-right">
        <div  className="text-white text-opacity-50 hover:text-opacity-90 hover:underline cursor-pointer"
          onClick={() => setAllTime(allTime === 0 ? 1 : 0)}
        >{allTime === 0?  "View 30-day" : " View all-time"}</div>
        </div>
      </Grid>
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
