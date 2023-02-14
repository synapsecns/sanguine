import { useState, useEffect } from 'react'
import { useQuery } from '@apollo/client'
import { AMOUNT_STATISTIC } from '@graphql/queries'
import Card from '@components/tailwind/Card'
import Grid from '@components/tailwind/Grid'
import numeral from 'numeral'
import { formatUSD } from '@utils/formatUSD'



export default function HolisticStats({platform: parentPlatform, setPlatform:  parentSetPlatform, loading, chainID, tokenAddress}) {
  const [volume, setVolume] = useState("--")
  const [revenue, setRevenue] = useState("--")
  const [addresses, setAddresses] = useState("--")
  const [txs, setTxs] = useState("--")
  const [useCache, setUseCache] = useState(true)
  const [skip, setSkip] = useState(false)
  const [platform, setPlatform] = useState(true)

  const unSelectStyle =
    'transition ease-out border-l-0 border-gray-700 border-opacity-30 text-gray-500 bg-gray-700 bg-opacity-30 hover:bg-opacity-20 '
  const selectStyle = 'text-white border-[#BE78FF] bg-synapse-radial'


  const { loading: loadingVolume, error: errorVolume, data: dataVolume } = useQuery(
    AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
    variables: chainID? {
      platform: platform,
      duration: "ALL_TIME",
      type: "TOTAL_VOLUME_USD",
      useCache: true,
      chainID: chainID
    }: {
      platform: platform,
      duration: "ALL_TIME",
      type: "TOTAL_VOLUME_USD",
      useCache: true
    },
    onCompleted: (data) => { setVolume(data.amountStatistic.value) }

  }
  )
  const { loading: loadingRevenue, error: errorRevenue, data: dataRevenue } = useQuery(
    AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
    variables: chainID ? {
      platform: platform,
      duration: "ALL_TIME",
      type: "TOTAL_FEE_USD",
      useCache: true,
      chainID: chainID
    }: {
      platform: platform,
      duration: "ALL_TIME",
      type: "TOTAL_FEE_USD",
      useCache: true
    },
    onCompleted: (data) => { setRevenue(data.amountStatistic.value) }

  }
  )
  const { loading: loadingAddresses, error: errorAddresses, data: dataAddresses } = useQuery(
    AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
    variables: chainID ? {
      platform: platform,
      duration: "ALL_TIME",
      type: "COUNT_ADDRESSES",
      useCache: useCache,
      chainID: chainID,
    }: {
      platform: platform,
      duration: "ALL_TIME",
      type: "COUNT_ADDRESSES",
      useCache: useCache,
    },
    pollInterval: 5000,
    skip: skip,
    onCompleted: (data) => {
      if (useCache) {
        setUseCache(true)
      }
      setAddresses(data.amountStatistic.value)

    }
  }
  )
  const { loading: loadingTxs, error: errorTxs, data: dataTxs } = useQuery(
    AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
    variables: chainID ? {
      platform: platform,
      duration: "ALL_TIME",
      type: "COUNT_TRANSACTIONS",
      useCache: useCache,
      chainID: chainID
    } : {
      platform: platform,
      duration: "ALL_TIME",
      type: "COUNT_TRANSACTIONS",
      useCache: useCache,
    },
    pollInterval: 5000,
    skip: skip,
    onCompleted: (data) => {
      if (useCache) {
        setUseCache(true)
      }
      setTxs(data.amountStatistic.value)
    }

  }
  )

  useEffect(() => {
    setPlatform(parentPlatform)
  }, [parentPlatform])


  useEffect(() => {
    setVolume("--")
    setRevenue("--")
    setAddresses("--")
    setTxs("--")

  }, [platform])


  const handlePlatform = (arg) => {
    setPlatform(arg)
    getVolume({
      variables: {
        platform: arg,
        duration: "ALL_TIME",
        type: "TOTAL_VOLUME_USD",
      },
    })
    getRevenue({
      variables: {
        platform: arg,
        duration: "ALL_TIME",
        type: "TOTAL_FEE_USD",
      },
    })
    getAddresses({
      variables: {
        platform: arg,
        duration: "ALL_TIME",
        type: "COUNT_ADDRESSES",
      },
    })
    getTxs({
      variables: {
        platform: arg,
        duration: "ALL_TIME",
        type: "COUNT_TRANSACTIONS",
      },
    })
  }

  return (
    <>
      <div className="my-2 mt-8">
        <button onClick={() => parentSetPlatform("ALL")} className={"font-medium rounded-l-md px-4 py-2 border  " + (platform === "ALL" ? selectStyle : unSelectStyle) + (loadingVolume ? " pointer-events-none" : "")}>
          All
        </button>
        <button onClick={() => parentSetPlatform("BRIDGE")} className={"font-medium  px-4 py-2 border  " + (platform === "BRIDGE" ? selectStyle : unSelectStyle) + (loadingVolume ? " pointer-events-none" : "")}>
          Bridge
        </button>
        <button onClick={() => parentSetPlatform("SWAP")} className={"font-medium  px-4 py-2 border  " + (platform === "SWAP" ? selectStyle : unSelectStyle) + (loadingVolume ? " pointer-events-none" : "")}>
          Swap
        </button>
        <button onClick={() => parentSetPlatform("MESSAGE_BUS")} className={"font-medium rounded-r-md px-4 py-2 border mr-5 " + (platform === "MESSAGE_BUS" ? selectStyle : unSelectStyle) + (loadingVolume ? " pointer-events-none" : "")}>
          Messaging
        </button>
      </div>
      <div className="flex flex-wrap flex-row">
        <Card
          className={`px-0 pb-2 space-y-3 text-white bg-transparent mr-[10%] min-w-[10%]`}
        >
          <div className="text-xl opacity-80">Volume</div>
          <div className="text-4xl font-bold text-white">
            {loadingVolume ? (<div className="h-9 w-full mt-4 bg-slate-700 rounded animate-pulse"></div>) :
              formatUSD(volume)}
          </div>
          <div className="flex space-x-2 text-sm font-medium">
            <div className="text-transparent bg-clip-text bg-gradient-to-r from-purple-500 to-purple-400">
              All Time
            </div>
          </div>
        </Card>
        <Card
          className={`px-0 pb-2 space-y-3 text-white bg-transparent mr-[10%] min-w-[10%]`}
        >
          <div className="text-xl opacity-80">Fees</div>
          <div className="text-4xl font-bold text-white">
            {loadingRevenue ? (<div className="h-9 w-full mt-4 bg-slate-700 rounded animate-pulse"></div>) :
              formatUSD(revenue)}
          </div>
          <div className="flex space-x-2 text-sm font-medium">

          </div>
        </Card>
        <Card
          className={`px-0 pb-2 space-y-3 text-white bg-transparent mr-[10%] min-w-[10%]`}
        >
          <div className="text-xl opacity-80">Transactions</div>
          <div className="text-4xl font-bold text-white">
            {numeral(txs).format('0,0')}
          </div>
          <div className="flex space-x-2 text-sm font-medium">

          </div>
        </Card>
        <Card
          className={`px-0 pb-2 space-y-3 text-white bg-transparent min-w-[10%]`}
        >
          <div className="text-xl opacity-80">Addresses</div>
          <div className="text-4xl font-bold text-white">
            {numeral(addresses).format('0,0')}
          </div>
          <div className="flex space-x-2 text-sm font-medium">

          </div>
        </Card>

      </div>

      {/*
      <Grid cols={{ sm: 1, md: 4, lg: 4 }} gap={4} className="my-4">
        <AllTimeStatCard
          title="Volume"
        >

        </AllTimeStatCard>
        <AllTimeStatCard
          title="Fees"
        >
          <div className="text-4xl font-bold text-white">
            {numeral(revenue / 1000000).format('$0.0')}M
          </div>
        </AllTimeStatCard>
        <AllTimeStatCard
          title="Transactions"
        >
          <div className="text-4xl font-bold text-white">
            {numeral(txs).format('0,0')}
          </div>
        </AllTimeStatCard>
        <AllTimeStatCard
          title="Addresses"
        >
          <div className="text-4xl font-bold text-white">
            {numeral(addresses).format('0,0')}
          </div>
        </AllTimeStatCard>

      </Grid>*/}
    </>
  )
}


export function AllTimeStatCard({
  title,
  children,
  duration = 'All Time',
}) {
  return (
    <Card
      className={`px-0 pb-2 space-y-3 text-white bg-transparent`}
    >
      <div className="text-xl">{title}</div>
      {children}
      <div className="flex space-x-2 text-sm font-medium">
        <div className="text-transparent bg-clip-text bg-gradient-to-r from-purple-500 to-purple-400">
          {duration}
        </div>
      </div>
    </Card>
  )
}
