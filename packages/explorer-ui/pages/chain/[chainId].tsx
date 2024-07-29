import _ from 'lodash'
import { useState, useEffect } from 'react'
import { useRouter } from 'next/router'
import { CHAINS } from 'synapse-constants'
import { useLazyQuery, useQuery } from '@apollo/client'
import {
  GET_BRIDGE_TRANSACTIONS_QUERY,
  DAILY_STATISTICS_BY_CHAIN,
} from '@graphql/queries'
import { ChainInfo } from '@components/misc/ChainInfo'
import { OverviewChart } from '@components/ChainChart'
import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { BridgeTransactionTable } from '@components/BridgeTransaction/BridgeTransactionTable'
import { SynapseLogoSvg } from '@components/layouts/MainLayout/SynapseLogoSvg'
import { HolisticStats } from '@components/misc/HolisticStats'
import { TRANSACTIONS_PATH } from '@urls'

const CHAIN_ID_NAMES_REVERSE = CHAINS.CHAIN_ID_NAMES_REVERSE

interface variablesType {
  chainIDFrom?: any
  chainIDTo?: any
  useMv?: boolean
}

export const ChainSummary = () => {
  const router = useRouter()
  const { chainId: chainIdRouter } = router.query
  const [platform, setPlatform] = useState('ALL')
  const [transactionsArr, setTransactionsArr] = useState([])
  const [dailyDataArr, setDailyDataArr] = useState([])
  const [variables, setVariables] = useState<variablesType>({})
  // eslint-disable-next-line @typescript-eslint/no-shadow
  const [chainId, setChainId] = useState<any>(0)
  const [completed] = useState(false)
  const [dailyStatisticType, setDailyStatisticType] = useState('VOLUME')
  const [dailyStatisticDuration, SetDailyStatisticDuration] =
    useState('PAST_6_MONTHS')
  const [dailyStatisticCumulative, SetDailyStatisticCumulative] = useState(true)
  const unSelectStyle =
    'transition ease-out border-l-0 border-gray-700 border-opacity-30 text-gray-500 bg-gray-700 bg-opacity-30 hover:bg-opacity-20 '
  const selectStyle = 'text-white border-[#BE78FF] bg-synapse-radial'

  const { loading, stopPolling, startPolling } = useQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY,
    {
      pollInterval: 5000,
      variables,
      fetchPolicy: 'network-only',
      onCompleted: (data) => {
        let bridgeTransactionsTable = data.bridgeTransactions

        bridgeTransactionsTable = _.orderBy(
          bridgeTransactionsTable,
          'fromInfo.time',
          ['desc']
        ).slice(0, 25)
        setTransactionsArr(bridgeTransactionsTable)
      },
    }
  )

  const [getDailyStatisticsByChain, { loading: loadingDailyData }] =
    useLazyQuery(DAILY_STATISTICS_BY_CHAIN, {
      onCompleted: (data) => {
        let chartData = data.dailyStatisticsByChain
        if (dailyStatisticCumulative) {
          chartData = JSON.parse(JSON.stringify(data.dailyStatisticsByChain))
          for (let i = 1; i < chartData.length; i++) {
            for (const key in data.dailyStatisticsByChain[i]) {
              if (key !== 'date' && key !== '__typename') {
                chartData[i][key] += chartData[i - 1]?.[key]
                  ? chartData[i - 1][key]
                  : 0
              }
            }
          }
        }
        setDailyDataArr(chartData)
      },
    })

  // update chart
  useEffect(() => {
    let type = dailyStatisticType
    if (platform === 'MESSAGE_BUS' && dailyStatisticType === 'VOLUME') {
      type = 'FEE'
      setDailyStatisticType('FEE')
    }
    getDailyStatisticsByChain({
      variables: {
        type,
        duration: dailyStatisticDuration,
        platform,
        useCache: false,
        chainID: chainId,
        useMv: true,
      },
    })
  }, [
    dailyStatisticDuration,
    dailyStatisticType,
    dailyStatisticCumulative,
    platform,
  ])

  // Get initial data
  useEffect(() => {
    getDailyStatisticsByChain({
      variables: {
        type: dailyStatisticType,
        duration: dailyStatisticDuration,
        useCache: false,
        chainID: chainIdRouter,
        useMv: true,
      },
    })
    setVariables({ chainIDFrom: chainIdRouter, useMv: true })
    setChainId(chainIdRouter)
  }, [chainIdRouter])

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

  return (
    <StandardPageContainer title={''}>
      <div className="flex items-center mt-10 mb-2">
        <h3 className="text-2xl font-semibold text-white">
          <ChainInfo
            chainId={chainId}
            imgClassName="w-10 h-10"
            textClassName="pl-1 whitespace-nowrap text-6xl text-white"
            linkClassName="float-right text-white transition ease-out hover:text-[#8FEBFF] px-2 ml-4 mt-4 rounded-md ease-in-out bg-[#191919]"
          />
        </h3>
      </div>
      <HolisticStats
        noMessaging={true}
        platform={platform}
        loading={false}
        setPlatform={setPlatform}
        baseVariables={{
          platform,
          duration: 'ALL_TIME',
          useCache: false,
          chainID: chainId,
          useMv: true,
        }}
      />
      <br />
      <HorizontalDivider />
      <div className="grid grid-cols-4 gap-4">
        <div className="col-span-1">
          <div className="flex w-full h-full bg-center bg-no-repeat z-1 bg-synapse-logo">
            <div id="tooltip-sidebar" className="w-full " />
          </div>
        </div>
        <div className="flex flex-col justify-end col-span-3 my-6 ">
          <div className="flex flex-wrap justify-end ">
            <div className="flex items-center h-full mr-4">
              {platform === 'MESSAGE_BUS' ? null : (
                <button
                  onClick={() => setDailyStatisticType('VOLUME')}
                  className={
                    'font-medium rounded-l-md px-4 py-2 border h-fit  ' +
                    (dailyStatisticType === 'VOLUME'
                      ? selectStyle
                      : unSelectStyle) +
                    (loadingDailyData || platform === 'MESSAGE_BUS'
                      ? ' pointer-events-none'
                      : '')
                  }
                >
                  Vol
                </button>
              )}
              <button
                onClick={() => setDailyStatisticType('FEE')}
                className={
                  'font-medium px-4 py-2 border  h-fit ' +
                  (dailyStatisticType === 'FEE' ? selectStyle : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '') +
                  (platform === 'MESSAGE_BUS' ? ' rounded-l-md' : '')
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
            <div className="flex items-center h-full mr-4">
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
                1mo
              </button>
              <button
                onClick={() => SetDailyStatisticDuration('PAST_3_MONTHS')}
                className={
                  'font-medium  px-4 py-2 border  h-fit   ' +
                  (dailyStatisticDuration === 'PAST_3_MONTHS'
                    ? selectStyle
                    : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '')
                }
              >
                3mo
              </button>
              <button
                onClick={() => SetDailyStatisticDuration('PAST_6_MONTHS')}
                className={
                  'font-medium rounded-r-md px-4 py-2 border  h-fit ' +
                  (dailyStatisticDuration === 'PAST_6_MONTHS'
                    ? selectStyle
                    : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '')
                }
              >
                6mo
              </button>
            </div>
            <div className="flex items-center h-full">
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

          <OverviewChart
            singleChain={true}
            loading={loadingDailyData}
            height={Object.keys(CHAIN_ID_NAMES_REVERSE).length * 31}
            chartData={dailyDataArr}
            dailyStatisticType={dailyStatisticType}
            isUSD={
              dailyStatisticType === 'TRANSACTIONS' ||
              dailyStatisticType === 'ADDRESSES'
                ? false
                : true
            }
            showAggregated={false}
            platform={platform}
            noTooltipLink={true}
          />
        </div>
      </div>
      <br /> <br />
      <HorizontalDivider />
      <br /> <br />
      <p className="text-2xl font-bold text-white">Recent Transactions</p>
      <div className="flex items-center h-full mt-4">
        <button
          onClick={() => setVariables({ chainIDFrom: chainId, useMv: true })}
          className={
            'font-medium rounded-l-md px-4 py-2 border  h-fit  ' +
            (variables?.chainIDFrom ? selectStyle : unSelectStyle) +
            (loadingDailyData ? ' pointer-events-none' : '')
          }
        >
          Outgoing
        </button>
        <button
          onClick={() => setVariables({ chainIDTo: chainId, useMv: true })}
          className={
            'font-medium rounded-r-md px-4 py-2 border  h-fit ' +
            (variables?.chainIDTo ? selectStyle : unSelectStyle) +
            (loadingDailyData ? ' pointer-events-none' : '')
          }
        >
          Incoming
        </button>
      </div>
      {loading ? (
        <div className="flex justify-center align-center w-full my-[100px]">
          <div className="w-[39px] animate-spin">
            <SynapseLogoSvg />
          </div>
        </div>
      ) : (
        <BridgeTransactionTable queryResult={transactionsArr} />
      )}
      <br />
      <div className="my-6 text-center text-white ">
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

const ChainPage = () => {
  return <ChainSummary />
}

export default ChainPage
