import { TRANSACTIONS_PATH } from '@urls'
import { useState, useEffect } from 'react'
import { TableHeader } from '@components/TransactionTable/TableHeader'
import { ChainInfo } from '@components/misc/ChainInfo'
import { OverviewChart } from '@components/ChainChart'

import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { formatUSD } from '@utils/formatUSD'
import { formatDate } from '@utils/formatDate'

import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { BridgeTransactionTable } from '@components/BridgeTransaction/BridgeTransactionTable'
import { useLazyQuery, useQuery } from '@apollo/client'
import { SynapseLogoSvg } from "@components/layouts/MainLayout/SynapseLogoSvg";
import { CHAIN_ID_NAMES_REVERSE } from '@constants/networks'

import {
  GET_BRIDGE_TRANSACTIONS_QUERY,
  DAILY_STATISTICS_BY_CHAIN,
} from '@graphql/queries'
import HolisticStats from '@components/misc/HolisticStats'
import _ from 'lodash'

const titles = {
  VOLUME: 'Volume',
  FEE: 'Fees',
  ADDRESSES: 'Addrs',
  TRANSACTIONS: 'TXs',
}
const platformTitles = {
  BRIDGE: 'Bridge',
  SWAP: 'Swap',
  MESSAGE_BUS: 'Message Bus',
}
const formatCurrency = new Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'USD',
})

export function Home() {
  const [currentTooltipIndex, setCurrentTooltipIndex] = useState(0)
  const [platform, setPlatform] = useState("ALL");
  const [transactionsArr, setTransactionsArr] = useState([])
  const [dailyDataArr, setDailyDataArr] = useState([])

  const [completed, setCompleted] = useState(false)
  const [dailyStatisticType, setDailyStatisticType] = useState('VOLUME')
  const [dailyStatisticDuration, SetDailyStatisticDuration] =
    useState('PAST_MONTH')
  const [dailyStatisticCumulative, SetDailyStatisticCumulative] =
    useState(false)
  const unSelectStyle =
    'transition ease-out border-l-0 border-gray-700 border-opacity-30 text-gray-500 bg-gray-700 bg-opacity-30 hover:bg-opacity-20 '
  const selectStyle = 'text-white border-[#BE78FF] bg-synapse-radial'
  const returnChainData = () => {
    var items = Object.keys(dailyDataArr?.[currentTooltipIndex]).map((key) => { return [key, dailyDataArr?.[currentTooltipIndex][key]] })
    items.sort((first, second) => { return second[1] - first[1] })
    var keys = items.map((e) => { return e[0] })
    return keys
  }
  const {
    loading,
    error,
    data: dataTx,
    stopPolling,
    startPolling,
  } = useQuery(GET_BRIDGE_TRANSACTIONS_QUERY, {
    pollInterval: 10000,
    fetchPolicy: 'network-only',
    onCompleted: (data) => {
      let bridgeTransactionsTable = data.bridgeTransactions

      bridgeTransactionsTable = _.orderBy(
        bridgeTransactionsTable,
        'fromInfo.time',
        ['desc']
      ).slice(0, 10)
      setTransactionsArr(bridgeTransactionsTable)

    },
  })

  const [
    getDailyStatisticsByChain,
    { loading: loadingDailyData, error: errorDailyData, data: dailyData },
  ] = useLazyQuery(DAILY_STATISTICS_BY_CHAIN, {
    onCompleted: (data) => {
      setDailyDataArr(data.dailyStatisticsByChain);
      setCurrentTooltipIndex(data.dailyStatisticsByChain.length - 1);
    }
  })


  // update chart
  useEffect(() => {
    let type = dailyStatisticType
    if (platform === "MESSAGE_BUS" && dailyStatisticType === "VOLUME") {
      type = "FEE"
      setDailyStatisticType("FEE")
    }
    getDailyStatisticsByChain({
      variables: {
        type: type,
        duration: dailyStatisticDuration,
        platform: platform,
        useCache: true,
      },
    })

  }, [dailyStatisticDuration, dailyStatisticType, dailyStatisticCumulative, platform])



  // Get initial data
  useEffect(() => {
    getDailyStatisticsByChain({
      variables: {
        type: dailyStatisticType,
        duration: dailyStatisticDuration,
        useCache: true,
      },
    })
  }, [])
  let chartData = dailyDataArr
  if (dailyStatisticCumulative) {
    chartData = JSON.parse(JSON.stringify(dailyDataArr))
    for (let i = 1; i < chartData.length; i++) {
      for (let key in dailyDataArr[i]) {
        if (key !== 'date' && key !== '__typename') {
          chartData[i][key] += (chartData[i - 1]?.[key] ? chartData[i - 1][key] : 0)
        }

      }
    }
  }

  useEffect(() => {
    if (!completed) {
      startPolling(10000)
    } else {
      stopPolling()
    }
    return () => {
      stopPolling()
    }
  }, [stopPolling, startPolling, completed])


  const totalChainVolume = () => {
    if (dailyStatisticCumulative) {
      return chartData[chartData.length - 1]["total"]
    }
    let totalRankedChainVolume = 0
    for (let i = 0; i < chartData.length; i++) {
      totalRankedChainVolume += chartData[i]["total"]

    }
    return totalRankedChainVolume

  }

  return (
    <StandardPageContainer title={'Synapse Analytics'}>
      <HolisticStats
        platform={platform}
        loading={false}
        setPlatform={setPlatform}
      />
      <br />
      <HorizontalDivider />
      <div className="grid grid-cols-3 gap-4">
        <div className="col-span-1">
          <div className="my-5">
            {currentTooltipIndex >= 0 && chartData?.[currentTooltipIndex] ? (
              <p
                className="text-2xl font-medium text-default
              font-bold
              text-white pl-2"
              >
                {formatDate(chartData[0].date)} to{' '}
                {formatDate(chartData[chartData.length - 1].date)}
              </p>
            ) : null}
            {chartData?.length > 0 ?
              <p className="pl-2 text-md font-medium text-default mt-2 text-white">{' '} {' '} Total {platform !== "ALL" ? platformTitles[platform] + " " : ""}{titles[dailyStatisticType]}: {' '}{formatUSD(totalChainVolume())}
              </p> : <div className="h-3 w-[50%] mt-4 bg-slate-700 rounded animate-pulse"></div>}
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
                  ((loadingDailyData || platform === "MESSAGE_BUS") ? ' pointer-events-none' : '') +
                  (platform === "MESSAGE_BUS" ? ' opacity-[0.6]' : '')
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
        <div className="col-span-1 w-[100%]">
          <p className="text-lg font-bold text-white pl-2 pt-4 ">{titles[dailyStatisticType]} for {formatDate(chartData?.[currentTooltipIndex]?.date)}</p>
          <table className='min-w-full'>

            <TableHeader headers={['Chain', titles[dailyStatisticType]]} />
            {loadingDailyData ? <tbody> {Object.values(CHAIN_ID_NAMES_REVERSE).map((i) =>
              <tr
                key={i}

              ><td className='w-[70%]'> <div className="h-3 w-full mt-4 bg-slate-700 rounded animate-pulse"></div></td><td className='w-[30%]'><div className="h-3 w-full mt-4 bg-slate-700 rounded animate-pulse"></div></td></tr>)}</tbody> :
              (<tbody>

                {currentTooltipIndex >= 0 && chartData?.[currentTooltipIndex] ? returnChainData().map((key, i) => {
                  return chartData[currentTooltipIndex][key] > 0 ? (<tr
                    key={i}
                    className=" rounded-md w-[100%]"
                  >
                    <td className='w-[70%]'>
                      {key === "total" ? <p className="pl-2 whitespace-nowrap text-sm text-white">All Chains</p> :
                        <ChainInfo
                          useExplorerLink={true}
                          chainId={CHAIN_ID_NAMES_REVERSE[key]}
                          imgClassName="w-4 h-4 ml-2"
                          textClassName="whitespace-nowrap px-2  text-sm  text-white"
                        />}
                    </td>
                    <td className='w-fit '>
                      <div className="ml-1 mr-2 self-center">
                        <p className='whitespace-nowrap px-2  text-sm  text-white'>{formatUSD(chartData[currentTooltipIndex][key])}</p>
                      </div>
                    </td>
                  </tr>) : null
                }) : null}
              </tbody>)}
          </table>
        </div>
        <div className="col-span-3 ">
          {/* { loadingDailyData ?  <div className={"flex justify-center align-center w-full animate-spin mt-[" + (Object.values(CHAIN_ID_NAMES_REVERSE).length * 10).toString() + "px]"}><SynapseLogoSvg /></div> : */}
          <OverviewChart
            setCurrentTooltipIndex={setCurrentTooltipIndex}
            loading={loadingDailyData}
            height={Object.keys(CHAIN_ID_NAMES_REVERSE).length * 31}
            chartData={chartData}
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
            currentIndex={currentTooltipIndex}
          />

        </div>
      </div>
      <br /> <br />
      <HorizontalDivider />
      <br /> <br />
      <p className="text-white text-2xl font-bold">Recent Transactions</p>
      {loading ? <div className="flex justify-center align-center w-full my-[100px] animate-spin"><SynapseLogoSvg /></div> : <BridgeTransactionTable queryResult={transactionsArr} />}


      <br />
      <div className="text-center text-white my-6 ">
        <div className="mt-2 mb-14 ">

          <a
            className="text-white rounded-md px-5 py-3 text-opacity-100 transition-all ease-in hover:bg-synapse-radial border-l-0 border-gray-700 border-opacity-30 bg-gray-700 bg-opacity-30 hover:border-[#BE78FF] cursor-pointer"
            href={TRANSACTIONS_PATH}
          >
            {'Explore all transactions'}
          </a>
        </div>
      </div>
      <HorizontalDivider />
    </StandardPageContainer>
  )
}
