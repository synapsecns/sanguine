import { TRANSACTIONS_PATH } from '@urls'
import { useState, useEffect } from 'react'
import { OverviewChart } from '@components/ChainChart'
import TextField from '@mui/material/TextField'
import { inputStyle } from '@utils/styles/muiStyles'
import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { BridgeTransactionTable } from '@components/BridgeTransaction/BridgeTransactionTable'
import { useLazyQuery, useQuery } from '@apollo/client'
import { SynapseLogoSvg } from '@components/layouts/MainLayout/SynapseLogoSvg'
import { CHAINS } from 'synapse-constants'
import {
  GET_BRIDGE_TRANSACTIONS_QUERY,
  DAILY_STATISTICS_BY_CHAIN,
} from '@graphql/queries'
import { HolisticStats } from '@components/misc/HolisticStats'
import _ from 'lodash'

const CHAIN_ID_NAMES_REVERSE = CHAINS.CHAIN_ID_NAMES_REVERSE

export const Home = () => {
  const [platform, setPlatform] = useState('ALL')
  const [transactionsArr, setTransactionsArr] = useState([])
  const [dailyDataArr, setDailyDataArr] = useState([])
  const [kappa, setKappa] = useState('')
  const [pending, setPending] = useState(false)

  const [completed] = useState(false)
  const [dailyStatisticType, setDailyStatisticType] = useState('VOLUME')
  const [dailyStatisticDuration, SetDailyStatisticDuration] =
    useState('PAST_6_MONTHS')
  const [dailyStatisticCumulative, SetDailyStatisticCumulative] = useState(true)
  const unSelectStyle =
    'transition ease-out border-l-0 border-gray-700 border-opacity-30 text-gray-500 bg-gray-700 bg-opacity-30 hover:bg-opacity-20 hover:text-white'
  const selectStyle = 'text-white border-[#BE78FF] bg-synapse-radial'
  const { loading, stopPolling, startPolling } = useQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY,
    {
      pollInterval: 10000,
      fetchPolicy: 'network-only',
      variables: {
        pending,
        useMv: true,
      },
      onCompleted: (data) => {
        let bridgeTransactionsTable = data.bridgeTransactions

        bridgeTransactionsTable = _.orderBy(
          bridgeTransactionsTable,
          'fromInfo.time',
          ['desc']
        ).slice(0, 15)
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
    if (platform === 'MESSAGE_BUS' && dailyStatisticType !== 'TRANSACTIONS') {
      type = 'TRANSACTIONS'
      setDailyStatisticType('TRANSACTIONS')
    }
    getDailyStatisticsByChain({
      variables: {
        type,
        duration: dailyStatisticDuration,
        platform,
        useCache: true,
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
        useCache: true,
        useMv: true,
      },
    })
  }, [])

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
    <StandardPageContainer title={'Synapse Analytics'}>
      <HolisticStats
        platform={platform}
        loading={false}
        setPlatform={setPlatform}
      />
      <br />
      <HorizontalDivider className="hidden sm:block" />
      <HorizontalDivider className="" />
      <div className="grid hidden grid-cols-4 gap-4 sm:grid">
        <div className="col-span-1">
          <div className="flex w-full h-full bg-center bg-no-repeat z-1 bg-synapse-logo">
            <div id="tooltip-sidebar" className="w-full " />
          </div>
        </div>
        <div className="flex flex-col justify-end col-span-3 my-6">
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
                onClick={() => setDailyStatisticType('TRANSACTIONS')}
                className={
                  'font-medium  px-4 py-2 border  h-fit ' +
                  (dailyStatisticType === 'TRANSACTIONS'
                    ? selectStyle
                    : unSelectStyle) +
                  (loadingDailyData ? ' pointer-events-none' : '') +
                  (platform === 'MESSAGE_BUS'
                    ? ' rounded-l-md rounded-r-md'
                    : '')
                }
              >
                TXs
              </button>
              {platform === 'MESSAGE_BUS' ? null : (
                <>
                  <button
                    onClick={() => setDailyStatisticType('ADDRESSES')}
                    className={
                      'font-medium  px-4 py-2 border h-fit  ' +
                      (dailyStatisticType === 'ADDRESSES'
                        ? selectStyle
                        : unSelectStyle) +
                      (loadingDailyData ? ' pointer-events-none' : '')
                    }
                  >
                    Addr
                  </button>
                  <button
                    onClick={() => setDailyStatisticType('FEE')}
                    className={
                      'font-medium px-4 py-2 border  h-fit rounded-r-md ' +
                      (dailyStatisticType === 'FEE'
                        ? selectStyle
                        : unSelectStyle) +
                      (loadingDailyData ? ' pointer-events-none' : '')
                    }
                  >
                    Fees
                  </button>
                </>
              )}
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
          />
        </div>
      </div>
      <br className="hidden sm:block" /> <br className="hidden sm:block" />
      <HorizontalDivider className="hidden sm:block" />
      <br /> <br />
      <p className="text-2xl font-bold text-center text-white sm:text-left">
        Recent Transactions
      </p>
      <div className="flex flex-col items-center justify-center py-6 space-y-3 sm:flex-row gap-x-4 sm:space-y-0">
        <div className="flex flex-row items-center w-full space-x-2">
          <TextField
            size="small"
            value={kappa}
            onChange={(e) => {
              setKappa(e.target.value)
            }}
            onKeyDown={(e) => {
              if (e.key === 'Enter') {
                window.location.href =
                  TRANSACTIONS_PATH + (kappa ? '?hash=' + kappa : '')
              }
            }}
            id="outlined-basic"
            label="Search by TXID / TXHash"
            variant="outlined"
            sx={inputStyle}
          />
          <a
            href={TRANSACTIONS_PATH + (kappa ? '?hash=' + kappa : '')}
            className={
              'font-medium rounded-md border border-l-0 border-gray-700 text-white bg-gray-700 px-4 py-2 hover:bg-opacity-70 ease-in-out duration-200 pointer-cursor z-10' +
              (loading ? ' pointer-events-none opacity-[0.4]' : '')
            }
            style={{ flexShrink: 0 }}
          >
            Search
          </a>
        </div>
        <div className="flex w-full sm:w-auto">
          <button
            disabled={loading}
            onClick={() => setPending(false)}
            className={
              'font-medium rounded-l-md px-4 py-2 border ' +
              (pending ? unSelectStyle : selectStyle) +
              (loading ? ' pointer-events-none' : '')
            }
          >
            Confirmed
          </button>
          <button
            disabled={loading}
            onClick={() => setPending(true)}
            className={
              'font-medium rounded-r-md px-4 py-2 border ' +
              (pending ? selectStyle : unSelectStyle) +
              (loading ? ' pointer-events-none' : '')
            }
          >
            Pending
          </button>
        </div>
      </div>
      {loading ? (
        <div className="flex justify-center align-center w-full my-[100px] ">
          <div className="mx-[1.5px] animate-spin">
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
