import { useState, useEffect } from 'react'
import { useQuery } from '@apollo/client'
import { AMOUNT_STATISTIC } from '@graphql/queries'
import Card from '@components/tailwind/Card'
import Grid from '@components/tailwind/Grid'
import numeral from 'numeral'
import { formatUSD } from '@utils/formatUSD'

interface HolisticStatsProps {
  platform: string;
  setPlatform: (platform: string) => void;
  loading: boolean;
  chainID?: string;
  baseVariables?: any;
  noMessaging?: boolean;
}

export default function HolisticStats({
                                        platform: parentPlatform,
                                        setPlatform: parentSetPlatform,
                                        loading,
                                        chainID,
                                        baseVariables,
                                        noMessaging,
                                      }: HolisticStatsProps) {
  const [volume, setVolume] = useState<string>('--')
  const [fee, setFee] = useState<string>('--')
  const [addresses, setAddresses] = useState<string>('--')
  const [txs, setTxs] = useState<string>('--')
  const [useCache, setUseCache] = useState<boolean>(true)
  const [skip, setSkip] = useState<boolean>(false)
  const [variables, setVariables] = useState<any>({})

  const [platform, setPlatform] = useState<any>(true)

  useEffect(() => {
    setVariables(baseVariables)
  }, [baseVariables])

  const unSelectStyle =
    'transition ease-out border-l-0 border-gray-700 border-opacity-30 text-gray-500 bg-gray-700 bg-opacity-30 hover:bg-opacity-20 hover:text-white'
  const selectStyle = 'text-white border-[#BE78FF] bg-synapse-radial'

  const handleVariable = (type: string) => {
    const queryVariables = JSON.parse(JSON.stringify(variables))
    queryVariables['type'] = type
    return queryVariables
  }

  const {
    loading: loadingVolume,
    error: errorVolume,
    data: dataVolume,
  } = useQuery(AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
    variables: baseVariables
      ? handleVariable('TOTAL_VOLUME_USD')
      : {
        platform,
        duration: 'ALL_TIME',
        type: 'TOTAL_VOLUME_USD',
        useCache: true,
        useMv: true,
      },
    onCompleted: (data: any) => {
      setVolume(data.amountStatistic.value)
    },
  })

  const {
    loading: loadingFee,
    error: errorFee,
    data: dataFee,
  } = useQuery(AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
    variables: baseVariables
      ? handleVariable('TOTAL_FEE_USD')
      : {
        platform,
        duration: 'ALL_TIME',
        type: 'TOTAL_FEE_USD',
        useCache: true,
      },
    onCompleted: (data: any) => {
      setFee(data.amountStatistic.value)
    },
  })

  const {
    loading: loadingAddresses,
    error: errorAddresses,
    data: dataAddresses,
  } = useQuery(AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
    variables: baseVariables
      ? handleVariable('COUNT_ADDRESSES')
      : {
        platform,
        duration: 'ALL_TIME',
        type: 'COUNT_ADDRESSES',
        useCache,
      },
    onCompleted: (data: any) => {
      setAddresses(data.amountStatistic.value)
    },
  })

  const {
    loading: loadingTxs,
    error: errorTxs,
    data: dataTxs,
  } = useQuery(AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
    variables: baseVariables
      ? handleVariable('COUNT_TRANSACTIONS')
      : {
        platform,
        duration: 'ALL_TIME',
        type: 'COUNT_TRANSACTIONS',
        useCache,
      },
    onCompleted: (data: any) => {
      setTxs(data.amountStatistic.value)
    },
  })

  useEffect(() => {
    setPlatform(parentPlatform)
  }, [parentPlatform])

  useEffect(() => {
    setVolume('--')
    setFee('--')
    setAddresses('--')
    setTxs('--')
  }, [platform])

  const stats = [
    parentPlatform === 'MESSAGE_BUS'
      ? null
      : {
        title: 'Volume',
        usd: true,
        loading: loadingVolume,
        value: formatUSD(volume),
      },
    {
      title: 'Transactions',
      loading: false,
      usd: false,
      value: formatUSD(txs),
    },
    parentPlatform === 'MESSAGE_BUS'
      ? null
      : {
        title: 'Addresses',
        loading: false,
        usd: false,
        value: formatUSD(addresses),
      },
    parentPlatform === 'MESSAGE_BUS'
      ? null
      : {
        title: 'Fees',
        loading: loadingFee,
        usd: true,
        value: formatUSD(fee),
      },
  ]

  return (
    <>
      <div className="my-2 mt-8">
        <button
          onClick={() => parentSetPlatform('ALL')}
          className={
            'font-medium rounded-l-md px-4 py-2 border  ' +
            (platform === 'ALL' ? selectStyle : unSelectStyle) +
            (loadingVolume ? ' pointer-events-none' : '')
          }
        >
          All
        </button>
        <button
          onClick={() => parentSetPlatform('BRIDGE')}
          className={
            'font-medium  px-4 py-2 border  ' +
            (platform === 'BRIDGE' ? selectStyle : unSelectStyle) +
            (loadingVolume ? ' pointer-events-none' : '')
          }
        >
          Bridge
        </button>
        <button
          onClick={() => parentSetPlatform('SWAP')}
          className={
            'font-medium  px-4 py-2 border  ' +
            (platform === 'SWAP' ? selectStyle : unSelectStyle) +
            (loadingVolume ? ' pointer-events-none' : '') +
            (noMessaging ? ' rounded-r-md' : '')
          }
        >
          Swap
        </button>
        {noMessaging ? null : (
          <button
            onClick={() => parentSetPlatform('MESSAGE_BUS')}
            className={
              'font-medium rounded-r-md px-4 py-2 border ' +
              (platform === 'MESSAGE_BUS' ? selectStyle : unSelectStyle) +
              (loadingVolume ? ' pointer-events-none' : '')
            }
          >
            Messaging
          </button>
        )}
      </div>
      <div className="flex flex-wrap flex-row min-h-[90px]">
        {stats.map((stat, i) => {
          return stat && stat.value !== '--' ? (
            <Card
              key={i}
              className={`px-0 pb-2 space-y-3 text-white bg-transparent mr-[10%] min-w-[10%]`}
            >
              <div className="text-xl opacity-80">{stat.title}</div>
              <div className="text-4xl font-bold text-white">
                {stat.loading ? (
                  <div className="h-9 w-full mt-4 bg-slate-700 rounded animate-pulse"></div>
                ) : stat.usd ? (
                  '$' + stat.value
                ) : (
                  stat.value
                )}
              </div>
            </Card>
          ) : parentPlatform !== 'MESSAGE_BUS' ? (
            <Card
              key={i}
              className={`px-0 pb-2 space-y-3 text-white bg-transparent mr-[10%] min-w-[10%]`}
            >
              <div className="h-9 w-full mt-4 bg-slate-700 rounded animate-pulse"></div>
            </Card>
          ) : null
        })}
      </div>
      <div className="flex space-x-2 text-sm pt-3 font-medium">
        <div className="text-transparent bg-clip-text bg-gradient-to-r from-purple-500 to-purple-400">
          All Time
        </div>
      </div>
    </>
  )
}
