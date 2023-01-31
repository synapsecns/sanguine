import { useState, useEffect } from 'react'
import { useQuery } from '@apollo/client'
import { AMOUNT_STATISTIC } from '@graphql/queries'
import Card from '@components/tailwind/Card'
import Grid from '@components/tailwind/Grid'
import numeral from 'numeral'



export default function HolisticStats() {
  const [volume, setVolume] = useState("--")
  const [revenue, setRevenue] = useState("--")
  const [addresses, setAddresses] = useState("--")
  const [txs, setTxs] = useState("--")
  const [platform, setPlatform] = useState("ALL");
  const unSelectStyle = "border-l-0 border-gray-700 border-opacity-30 text-gray-500 bg-gray-700 bg-opacity-30"
  const selectStyle = "text-white border-[#BE78FF] bg-synapse-radial"

  const { loading: loadingVolume, error: errorVolume, data: dataVolume } = useQuery(
    AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
    variables: {
      platform: platform,
      duration: "ALL_TIME",
      type: "TOTAL_VOLUME_USD",
    },
    pollInterval: 10000,
  }
  )
  const { loading: loadingRevenue, error: errorRevenue, data: dataRevenue } = useQuery(
    AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
    variables: {
      platform: platform,
      duration: "ALL_TIME",
      type: "TOTAL_FEE_USD",
    },
    pollInterval: 10000,
  }
  )
  const { loading: loadingAddresses, error: errorAddresses, data: dataAddresses } = useQuery(
    AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
    variables: {
      platform: platform,
      duration: "ALL_TIME",
      type: "COUNT_ADDRESSES",
    },
    pollInterval: 10000,
    notifyOnNetworkStatusChange: true,

  }
  )
  const { loading: loadingTxs, error: errorTxs, data: dataTxs, startPolling: stopPollingTx, stopPolling: startPollingTx } = useQuery(
    AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
    variables: {
      platform: platform,
      duration: "ALL_TIME",
      type: "COUNT_TRANSACTIONS",
    },
    pollInterval: 10000,
    notifyOnNetworkStatusChange: true,

  }
  )

  // useEffect(() => {
  //   // versionRefetch()
  //   startPollingTx(10000)

  //   return () => {
  //     stopPollingTx()
  //   }
  // }, [stopPollingTx, startPollingTx])
  // Get initial data
  useEffect(() => {
    setVolume("--")
    setRevenue("--")
    setAddresses("--")
    setTxs("--")

  }, [platform])

  // Get data when search params change
  useEffect(() => {
    if (dataVolume) {
      console.log("setting", dataVolume)
      setVolume(dataVolume.amountStatistic.value)
    }
  }, [dataVolume])
  useEffect(() => {
    if (dataRevenue) {
      console.log("setting", dataRevenue)
      setRevenue(dataRevenue.amountStatistic.value)
    }
  }, [dataRevenue])
  useEffect(() => {
    if (dataAddresses) {
      console.log("setting", dataAddresses)
      setAddresses(dataAddresses.amountStatistic?.value)
    }
  }, [dataAddresses])
  useEffect(() => {
    if (dataTxs) {
      console.log("setting", dataTxs)
      setTxs(dataTxs.amountStatistic?.value)
    }
  }, [dataTxs])

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
    <button onClick={() => setPlatform("ALL")} className={"font-medium rounded-l-md px-4 py-2 border  " + (platform === "ALL" ? selectStyle : unSelectStyle) + (loadingVolume ? " pointer-events-none" : "")}>
      All
    </button>
    <button onClick={() => setPlatform("BRIDGE")} className={"font-medium  px-4 py-2 border  " + (platform === "BRIDGE" ? selectStyle : unSelectStyle) + (loadingVolume ? " pointer-events-none" : "")}>
      Bridge
    </button>
    <button onClick={() => setPlatform("SWAP")} className={"font-medium  px-4 py-2 border  " + (platform === "SWAP" ? selectStyle : unSelectStyle) + (loadingVolume ? " pointer-events-none" : "")}>
      Swap
    </button>
    <button onClick={() => setPlatform("MESSAGING")} className={"font-medium rounded-r-md px-4 py-2 border mr-5 " + (platform === "MESSAGING" ? selectStyle : unSelectStyle) + (loadingVolume ? " pointer-events-none" : "")}>
      Messaging
    </button>
  </div>
    <Grid cols={{ sm: 1, md: 4, lg: 4 }} gap={4} className="my-4">
      <AllTimeStatCard
        title="Volume"
      >
        <div className="text-4xl font-bold text-white">
          {numeral(volume / 1000000000).format('$0.000')}B
        </div>
      </AllTimeStatCard>
      <AllTimeStatCard
        title="Revenue"
      >
        <div className="text-4xl font-bold text-white">
          {numeral(revenue/ 1000000).format('$0.000')}M
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

    </Grid></>
  )
}


export function AllTimeStatCard({
  title,
  children,
  duration = 'All-Time',
}) {
  return (
    <Card
      className={`px-0 pb-2 space-y-3 text-white bg-transparent cursor-pointer`}
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
