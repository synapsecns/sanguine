import { TRANSACTIONS_PATH } from '@urls'
import { useState, useEffect } from 'react'
import { Stats } from './Stats'
import { UniversalSearch } from '@components/pages/Home/UniversalSearch'
import { TableHeader } from '@components/TransactionTable/TableHeader'
import { TableBody } from '@components/TransactionTable/TableBody'
import { ChainInfo } from '@components/misc/ChainInfo'

import { Chart, ChartLoading } from '@components/Chart'
import { OverviewChart } from '@components/ChainChart'

import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { PageLink } from '@components/misc/PageLink'
import Grid from '@components/tailwind/Grid'

import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { BridgeTransactionTable } from '@components/BridgeTransaction/BridgeTransactionTable'
import { useLazyQuery, useQuery } from '@apollo/client'

import {
  GET_BRIDGE_TRANSACTIONS_QUERY,
  DAILY_STATISTICS_BY_CHAIN,
  RANKED_CHAINIDS_BY_VOLUME,
} from '@graphql/queries'
import { useSearchParams } from 'next/navigation'
import HolisticStats from '@components/misc/HolisticStats'
import _ from 'lodash'
import { isCompositeType } from 'graphql'

const formatCurrency = new Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'USD',
})

export function Home({ }) {
  const search = useSearchParams()
  const [currentChainID, setCurrentChainID] = useState(0)
  const [pending, setPending] = useState(false)
  const [transactionsArr, setTransactionsArr] = useState([])
  const [dailyDataArr, setDailyDataArr] = useState([])
  const [completed, setCompleted] = useState(false)
  const [dailyStatisticType, setDailyStatisticType] = useState('VOLUME')
  const [dailyStatisticDuration, SetDailyStatisticDuration] =
    useState('PAST_MONTH')
  const [dailyStatisticCumulative, SetDailyStatisticCumulative] =
    useState(false)
  const [rankedChainIDs, setRankedChainIDs] = useState([])

  const unSelectStyle =
    'border-l-0 border-gray-700 border-opacity-30 text-gray-500 bg-gray-700 bg-opacity-30'
  const selectStyle = 'text-white border-[#BE78FF] bg-synapse-radial'
  const formatDate = (date) => {
    const d = new Date(date)
    const month = d.getMonth() + 1
    const day = d.getDate()
    const year = d.getFullYear()
    return `${month}/${day}/${year}`
  }
  // var { loading, error, dataTx, refetch } = useQuery(GET_BRIDGE_TRANSACTIONS_QUERY)

  const {
    loading,
    error,
    data: dataTx,
    stopPolling,
    startPolling,
  } = useQuery(GET_BRIDGE_TRANSACTIONS_QUERY, {
    pollInterval: 10000000,
    notifyOnNetworkStatusChange: true,
    onCompleted: () => console.log('called'),
  })

  const [
    getDailyStatisticsByChain,
    { loading: loadingDailyData, error: errorDailyData, data: dailyData },
  ] = useLazyQuery(DAILY_STATISTICS_BY_CHAIN)

  const [
    getChainIDsRankedByVolume,
    {
      loading: loadingRankedChains,
      error: errorRankedChains,
      data: rankedChainsData,
    },
  ] = useLazyQuery(RANKED_CHAINIDS_BY_VOLUME)

  useEffect(() => {
    if (dailyData) {
      setDailyDataArr(dailyData.dailyStatisticsByChain, {
        variables: {
          type: dailyStatisticType,
          duration: dailyStatisticDuration,
        },
      })
    }
  }, [dailyData, loadingDailyData])

  useEffect(() => {
    if (rankedChainsData) {
      setRankedChainIDs(rankedChainsData.rankedChainIDsByVolume, {
        variables: {
          duration: dailyStatisticDuration,
        },
      })
    }
  }, [rankedChainsData, loadingRankedChains])
  // Get initial data
  useEffect(() => {
    getDailyStatisticsByChain({
      variables: {
        type: dailyStatisticType,
        duration: dailyStatisticDuration,
      },
    })
    getChainIDsRankedByVolume({
      variables: {
        duration: dailyStatisticDuration,
      },
    })
  }, [])

  useEffect(() => {
    getDailyStatisticsByChain({
      variables: {
        type: dailyStatisticType,
        duration: dailyStatisticDuration,
      },
    })
  }, [dailyStatisticDuration, dailyStatisticType, dailyStatisticCumulative])

  useEffect(() => {
    getChainIDsRankedByVolume({
      variables: {
        duration: dailyStatisticDuration,
      },
    })
  }, [dailyStatisticDuration])

  useEffect(() => {
    // versionRefetch()
    if (!completed) {
      startPolling(10000000)
    } else {
      stopPolling()
    }
    return () => {
      stopPolling()
    }
  }, [stopPolling, startPolling, completed])

  // Get data when search params change
  useEffect(() => {
    if (dataTx) {
      setTransactionsArr(dataTx.bridgeTransactions, {
        variables: {
          pending: pending,
        },
      })
    }
  }, [dataTx, search, pending])

  // let data
  // if (chartType === 'VOLUME') {
  //   data = bridgeVolume && bridgeVolume.dailyStatistics.dateResults
  // } else if (chartType === 'TRANSACTIONS') {
  //   data = transactions && transactions.dailyStatistics.dateResults
  // } else if (chartType === 'ADDRESSES') {
  //   data = addresses && addresses.dailyStatistics.dateResults
  // }
  let txContent
  let bridgeTransactionsTable = transactionsArr

  bridgeTransactionsTable = _.orderBy(
    bridgeTransactionsTable,
    'fromInfo.time',
    ['desc']
  ).slice(0, 10)

  txContent = <BridgeTransactionTable queryResult={bridgeTransactionsTable} />

  let totalRankedChainVolume = 0
  for (let i = 0; i < rankedChainIDs.length; i++) {
    totalRankedChainVolume += rankedChainIDs[i].total
  }

  return (
    <StandardPageContainer title={'Synapse Analytics'}>
      <HolisticStats />
      <br />
      <HorizontalDivider />
      <div className="grid grid-cols-3 gap-4">
        <div className="col-span-1">
          <div className="my-5">
            {currentChainID === 0 ? (
              <p
                className="text-4xl font-medium text-default
              font-bold
              text-white"
              >
                All Chains
              </p>
            ) : (
              <ChainInfo
                chainId={currentChainID}
                imgClassName="w-8 h-8"
                linkClassName="bg-gray-700 p-1 rounded-md ml-1 mt-2"
                textClassName="text-4xl font-medium text-default
            font-bold
            text-white"
              />
            )}
            {dailyDataArr?.length > 0 ?
              <p className="text-md font-medium text-default mt-2 text-white">
                {formatDate(dailyDataArr[0].date)} to{' '}
                {formatDate(dailyDataArr[dailyDataArr.length - 1].date)}
              </p> : null}
          </div>
        </div>
        <div className="col-span-2 flex justify-end">
          <div className="flex flex-wrap">
            <div className="h-full flex items-center mr-4">
              <button
                onClick={() => setDailyStatisticType('VOLUME')}
                className={
                  'font-medium rounded-l-md px-4 py-2 border h-fit  ' +
                  (dailyStatisticType === 'VOLUME'
                    ? selectStyle
                    : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '')
                }
              >
                Vol
              </button>
              <button
                onClick={() => setDailyStatisticType('FEE')}
                className={
                  'font-medium px-4 py-2 border  h-fit ' +
                  (dailyStatisticType === 'FEE' ? selectStyle : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '')
                }
              >
                Fees
              </button>
              <button
                onClick={() => setDailyStatisticType('TRANSACTIONS')}
                className={
                  'font-medium  px-4 py-2 border  h-fit ' +
                  (dailyStatisticType === 'TRANSACTIONS'
                    ? selectStyle
                    : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '')
                }
              >
                TXs
              </button>
              <button
                onClick={() => setDailyStatisticType('ADDRESSES')}
                className={
                  'font-medium rounded-r-md px-4 py-2 border h-fit  ' +
                  (dailyStatisticType === 'ADDRESSES'
                    ? selectStyle
                    : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '')
                }
              >
                Addr
              </button>
            </div>
            <div className="h-full flex items-center mr-4">
              <button
                onClick={() => SetDailyStatisticDuration('PAST_MONTH')}
                className={
                  'font-medium rounded-l-md px-4 py-2 border  h-fit  ' +
                  (dailyStatisticDuration === 'PAST_MONTH'
                    ? selectStyle
                    : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '')
                }
              >
                30d
              </button>
              <button
                onClick={() => SetDailyStatisticDuration('PAST_YEAR')}
                className={
                  'font-medium  px-4 py-2 border  h-fit   ' +
                  (dailyStatisticDuration === 'PAST_YEAR'
                    ? selectStyle
                    : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '')
                }
              >
                365d
              </button>
              <button
                onClick={() => SetDailyStatisticDuration('ALL_TIME')}
                className={
                  'font-medium rounded-r-md px-4 py-2 border  h-fit ' +
                  (dailyStatisticDuration === 'ALL_TIME'
                    ? selectStyle
                    : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '')
                }
              >
                All Time
              </button>
            </div>
            <div className="h-full flex items-center">
              <button
                onClick={() => SetDailyStatisticCumulative(false)}
                className={
                  'font-medium rounded-l-md px-4 py-2 border  h-fit  ' +
                  (!dailyStatisticCumulative ? selectStyle : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '')
                }
              >
                Daily
              </button>
              <button
                onClick={() => SetDailyStatisticCumulative(true)}
                className={
                  'font-medium rounded-r-md px-4 py-2 border  h-fit ' +
                  (dailyStatisticCumulative ? selectStyle : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '')
                }
              >
                Cumulative
              </button>
            </div>
          </div>
        </div>
      </div>
      <HorizontalDivider />
      <div className="grid grid-cols-4 gap-4">
        <div className="col-span-1">
          <div className="pb-2 px-4 sm:px-6 lg:px-8">
            <div className="mt-8 flex flex-col">
              <div className="-my-2 -mx-4 overflow-x-auto sm:-mx-6 lg:-mx-8">
                <div className="inline-block min-w-full py-2 align-middle">
                  <div className="overflow-hidden shadow-sm ring-1 ring-black ring-opacity-5">
                    <table className="min-w-full">
                      <TableHeader headers={['Chain', 'Volume']} />
                      <tbody>
                      <tr
                            key={0}
                            className="hover:bg-synapse-radial rounded-md cursor-pointer "
                            onClick={() => setCurrentChainID(0)}
                          >
                            <td>
                              <p
                                className="text-1xl font-medium text-default text-white ml-2"
                              >All Chains</p>
                            </td>
                            <td>
                              <div className="ml-1 mr-2 text-white self-center">
                                {formatCurrency.format(totalRankedChainVolume)}
                              </div>
                            </td>
                          </tr>
                        {rankedChainIDs.map((row, i) => (

                          <tr
                            key={i}
                            className="hover:bg-synapse-radial rounded-md cursor-pointer "
                            onClick={() => setCurrentChainID(row.chainID)}
                          >
                            <td>
                              <ChainInfo
                                noLink={true}
                                chainId={row.chainID}
                                imgClassName="w-4 h-4 ml-2"
                                textClassName="text-1xl font-medium text-default text-white"
                              />
                            </td>
                            <td>
                              <div className="ml-1 mr-2 text-white self-center">
                                {formatCurrency.format(row.total)}
                              </div>
                            </td>
                          </tr>
                        ))}
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="col-span-3 ">
          <br />
          <OverviewChart
            data={dailyDataArr}
            isCumulativeData={dailyStatisticCumulative}
            isUSD={
              dailyStatisticType === 'TRANSACTIONS' ||
                dailyStatisticType === 'ADDRESSES'
                ? false
                : true
            }
            showAggregated={false}
            monthlyData={false}
            currency
          />
        </div>{' '}
      </div>
      <HorizontalDivider />
      <br /> <br />
      <p className="text-white text-2xl font-bold">Recent Transactions</p>
      {txContent}
      <br />
      <div className="text-center text-white my-6 ">
        <div className="mt-2 mb-14 ">
          <a
            className="text-white rounded-md px-5 py-3 border text-opacity-100 transition-all ease-in hover:bg-synapse-radial "
            href={TRANSACTIONS_PATH}
            target="_blank"
            rel="noreferrer"
          >
            {'Explore all transactions'}
          </a>
        </div>
      </div>
      <HorizontalDivider />
    </StandardPageContainer>
  )
}
