import { ANALYTICS_PATH, TRANSACTIONS_PATH } from '@urls'
import { useState, useEffect } from 'react'
import { Stats } from './Stats'
import { UniversalSearch } from '@components/pages/Home/UniversalSearch'

import { Chart, ChartLoading } from '@components/Chart'
import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { PageLink } from '@components/misc/PageLink'
import Grid from '@components/tailwind/Grid'

import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { BridgeTransactionTable } from '@components/BridgeTransaction/BridgeTransactionTable'
import { useLazyQuery, useQuery } from '@apollo/client'

import { GET_BRIDGE_TRANSACTIONS_QUERY } from '@graphql/queries'
import { useSearchParams } from 'next/navigation'
import HolisticStats from "@components/misc/HolisticStats";
import _ from 'lodash'



export function Home({
  bridgeVolume,
  transactions,
  addresses,
  popularTokens,
  popularChains,
}) {
  const search = useSearchParams()
  const [chartType, setChartType] = useState('VOLUME')
  const [allTime, setAllTime] = useState(1)
  const [pending, setPending] = useState(false)
  const [transactionsArr, setTransactionsArr] = useState([])
  const [completed, setCompleted] = useState(false);

  // var { loading, error, dataTx, refetch } = useQuery(GET_BRIDGE_TRANSACTIONS_QUERY)


  const { loading, pageError, error: error, data: dataTx, stopPolling, startPolling } = useQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY, {
    pollInterval: 10000,
    notifyOnNetworkStatusChange: true,
    onCompleted: () => console.log('called'),

  }
  )

  useEffect(() => {
    // versionRefetch()
    if (!completed) {
      startPolling(10000)
    } else {
      stopPolling()
    }
    return () => {
      stopPolling()
    }
  }, [stopPolling, startPolling, completed])

  // // Get initial data
  // useEffect(() => {
  //   getBridgeTransactions({
  //     variables: {
  //       pending: pending,
  //       page: 1,
  //     },
  //   })
  // }, [])

  // Get data when search params change
  useEffect(() => {
    console.log("SSS", loading, pageError, dataTx)
    if (dataTx) {
      console.log("setting")
      setTransactionsArr(dataTx.bridgeTransactions, {
        variables: {
          pending: pending,
        },
      })
    }

  }, [dataTx, search, pending])

  const handlePending = (arg) => {
    setPending(arg)
    getBridgeTransactions({
      variables: {
        pending: arg,
        page: 1,
      },
    })
  }

  console.log("u:", dataTx?.bridgeTransactions, pageError)
  let data
  if (chartType === 'VOLUME') {
    data = bridgeVolume && bridgeVolume.dailyStatistics.dateResults
  } else if (chartType === 'TRANSACTIONS') {
    data = transactions && transactions.dailyStatistics.dateResults
  } else if (chartType === 'ADDRESSES') {
    data = addresses && addresses.dailyStatistics.dateResults
  }
  let txContent
  let bridgeTransactionsTable = transactionsArr

  bridgeTransactionsTable = _.orderBy(
    bridgeTransactionsTable,
    'fromInfo.time',
    ['desc']
  ).slice(0, 10)

  txContent = <BridgeTransactionTable queryResult={bridgeTransactionsTable} />
  return (

    <StandardPageContainer title={"Synapse Analytics"}>

      <p className='text-white text-2xl font-bold'>All Time Statistics</p>
      <HolisticStats />
      <br />
      <HorizontalDivider />
      <br />

      <p className='text-white text-2xl font-bold'>30-Day Statistics</p>
      <Chart data={data} ttl />
      {bridgeVolume &&
        transactions &&
        addresses ? (

        <Stats
          bridgeVolume={bridgeVolume.dailyStatistics.total}
          transactions={transactions.dailyStatistics.total}
          addresses={addresses.dailyStatistics.total}
          setChartType={setChartType}
          allTime={false}
        />

      ) : (
        <ChartLoading />
      )}
      <HorizontalDivider />
      {/* <Grid cols={{ sm: 1, md: 2, lg: 2 }} gap={4} className="my-3">
        <PageLink text={'See more'} url={ANALYTICS_PATH} external={true} />
        <div className="mt-2 mb-10 text-right">
          <div
            className="text-white text-opacity-50 hover:text-opacity-90 hover:underline cursor-pointer"
            onClick={() => setAllTime(allTime === 0 ? 1 : 0)}
          >
            {allTime === 1 ? 'View 30-day' : ' View all-time'}
          </div>
        </div>
      </Grid> */}
      {/* <UniversalSearch
        placeholder="Search by address or transaction"
        setPending={handlePending}
        pending={pending}
        loading={loading}
      /> */}
      {/* <button onClick={() => setPending(!pending)}>Refetch</button> */}
      <br /> <br />

      <p className='text-white text-2xl font-bold'>Recent all Transactions</p>
        {txContent}
      <div className='text-center text-white my-6'>
        <PageLink text="Explore all transactions" url={TRANSACTIONS_PATH} />

      </div>
      <HorizontalDivider />
    </StandardPageContainer>
  )
}
