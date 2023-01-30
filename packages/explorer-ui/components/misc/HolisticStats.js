import { useState, useEffect } from 'react'
import { useLazyQuery } from '@apollo/client'
import { AMOUNT_STATISTIC } from '@graphql/queries'
import Card from '@components/tailwind/Card'
import Grid from '@components/tailwind/Grid'
import numeral from 'numeral'



export default function HolisticStats() {
  const [platform, setPlatform] = useState("ALL")
  const [volume, setVolume] = useState("--")
  const [revenue, setRevenue] = useState("--")
  const [addresses, setAddresses] = useState("--")
  const [txs, setTxs] = useState("--")


  const [getVolume, { loading: loadingVolume, error: errorVolume, data: dataVolume }] = useLazyQuery(
    AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
  }
  )
  const [getRevenue, { loading: loadingRevenue, error: errorRevenue, data: dataRevenue }] = useLazyQuery(
    AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
  }
  )
  const [getAddresses, { loading: loadingAddresses, error: errorAddresses, data: dataAddresses }] = useLazyQuery(
    AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
  }
  )
  const [getTxs, { loading: loadingTxs, error: errorTxs, data: dataTxs }] = useLazyQuery(
    AMOUNT_STATISTIC, {
    fetchPolicy: 'network-only',
  }
  )

  // Get initial data
  useEffect(() => {
    getVolume({
      variables: {
        platform: platform,
        duration: "ALL_TIME",
        type: "TOTAL_VOLUME_USD",
      },
    })
    getRevenue({
      variables: {
        platform: platform,
        duration: "ALL_TIME",
        type: "TOTAL_FEE_USD",
      },
    })
    getAddresses({
      variables: {
        platform: platform,
        duration: "ALL_TIME",
        type: "COUNT_ADDRESSES",
      },
    })
    getTxs({
      variables: {
        platform: platform,
        duration: "ALL_TIME",
        type: "COUNT_TRANSACTIONS",
      },
    })
  }, [])

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
      setAddresses(dataAddresses.amountStatistic.value)
    }
  }, [dataAddresses])
  useEffect(() => {
    if (dataTxs) {
      console.log("setting", dataTxs)
      setTxs(dataTxs.amountStatistic.value)
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
    <Grid cols={{ sm: 1, md: 4, lg: 4 }} gap={4} className="my-4">
      <AllTimeStatCard
        title="Bridge Volume"
      >
        <div className="text-4xl font-bold text-white">
          {numeral(volume / 1000000000).format('$0.000')}B
        </div>
      </AllTimeStatCard>
      <AllTimeStatCard
        title="Revenue"
      >
        <div className="text-4xl font-bold text-white">
          {numeral(revenue).format('0,0')}
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

    </Grid>
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
