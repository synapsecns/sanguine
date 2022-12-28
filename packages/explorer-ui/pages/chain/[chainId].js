import {GET_BRIDGE_TRANSACTIONS_QUERY, GET_HISTORICAL_STATS,} from '@graphql/queries'
import {ApolloClient, HttpLink, InMemoryCache} from '@apollo/client'
import {API_URL} from '@graphql'

import _ from 'lodash'
import { useState, useEffect } from 'react'
import { useQuery, useLazyQuery } from '@apollo/client'

import { Pagination } from '@components/Pagination'
import { ChainInfo } from '@components/misc/ChainInfo'

import { StandardPageContainer } from '@layouts/StandardPageContainer'

import { Stats } from '@components/pages/Home/Stats'
import { Chart, ChartLoading } from '@components/Chart'
import {BridgeTransactionTable} from "@components/BridgeTransaction/BridgeTransactionTable";
import {useRouter} from "next/router";
import {useSearchParams} from 'next/navigation'



const link = new HttpLink({
  uri: API_URL,
  useGETForQueries: true,
})

const client = new ApolloClient({
  link: link,
  cache: new InMemoryCache(),
  fetchPolicy: 'network-only',
  fetchOptions: {
    mode: 'no-cors',
  },
})

export default function chainId({
  bridgeVolume,
  transactionsData,
  addresses,
  latestBridgeTransactions,
}) {
  const router = useRouter()
  const { chainId } = router.query

  const [chartType, setChartType] = useState('BRIDGEVOLUME')

  let chartData

  if (chartType === 'BRIDGEVOLUME') {
    chartData = bridgeVolume && bridgeVolume.historicalStatistics.dateResults
  } else if (chartType === 'TRANSACTIONS') {
    chartData =
      transactionsData && transactionsData.historicalStatistics.dateResults
  } else if (chartType === 'ADDRESSES') {
    chartData = addresses && addresses.historicalStatistics.dateResults
  }

  const search = useSearchParams()
  const p = Number(search.get('page')) || 1

  const [page, setPage] = useState(p)
  const [transactions, setTransactions] = useState([])

  const [getBridgeTransactions, { error: pageError, data }] = useLazyQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY
  )

  useEffect(() => {
    if (data) {
      setTransactions(data.bridgeTransactions, {
        variables: {
          chainId: Number(chainId),
        },
      })
    }

    const num = Number(search.get('page'))

    if (num === 0) {
      setPage(1)
      getBridgeTransactions({
        variables: {
          chainId: Number(chainId),
          page: 1,
        },
      })
    } else {
      setPage(num)
      getBridgeTransactions({
        variables: {
          chainId: Number(chainId),
          page: num,
        },
      })
    }
  }, [data, search, chainId])

  const nextPage = () => {
    let newPage = page + 1
    setPage(newPage)
    // setSearch({ page: newPage })

    getBridgeTransactions({
      variables: { chainId: Number(chainId), page: newPage },
    })
  }

  const prevPage = () => {
    if (page > 1) {
      let newPage = page - 1
      setPage(newPage)
      // setSearch({ page: newPage })
      getBridgeTransactions({
        variables: { chainId: Number(chainId), page: newPage },
      })
    }
  }

  const resetPage = () => {
    setPage(1)
    // setSearch({ page: 1 })
    getBridgeTransactions({
      variables: { chainId: Number(chainId), page: 1 },
    })
  }

  let content


    let bridgeTransactions = transactions

    bridgeTransactions = _.orderBy(bridgeTransactions, 'fromInfo.time', [
      'desc',
    ]).slice(0, 10)

  content=<BridgeTransactionTable queryResult={bridgeTransactions} />

  let title = <ChainInfo chainId={chainId} imgClassName="w-7 h-7" />

  return (
    <StandardPageContainer title={title}>
      <Chart data={chartData} />
      {bridgeVolume && transactionsData && addresses ? (
        <Stats
          bridgeVolume={bridgeVolume.historicalStatistics.total}
          transactions={transactionsData.historicalStatistics.total}
          addresses={addresses.historicalStatistics.total}
          setChartType={setChartType}
        />
      ) : (
        <ChartLoading />
      )}
      <div className="mb-10" />
      {content}
      <Pagination
        page={page}
        resetPage={resetPage}
        prevPage={prevPage}
        nextPage={nextPage}
      />
    </StandardPageContainer>
  )
}

export async function getServerSideProps(context) {
  const { data: bridgeVolume } = await client.query({
    query: GET_HISTORICAL_STATS,
    variables: {
      type: 'BRIDGEVOLUME',
      days: 30,
      chainId: Number(context.params.chainId),
    },
  })

  const { data: transactionsData } = await client.query({
    query: GET_HISTORICAL_STATS,
    variables: {
      type: 'TRANSACTIONS',
      days: 30,
      chainId: Number(context.params.chainId),
    },
  })

  const { data: addresses } = await client.query({
    query: GET_HISTORICAL_STATS,
    variables: {
      type: 'ADDRESSES',
      days: 30,
      chainId: Number(context.params.chainId),
    },
  })

  const { data: latestBridgeTransactions } = await client.query({
    query: GET_BRIDGE_TRANSACTIONS_QUERY,
    variables: {
      chainId: Number(context.params.chainId),
      page: 1,
    },
  })

  return {
    props: {
      bridgeVolume: bridgeVolume,
      transactionsData: transactionsData,
      addresses: addresses,
      latestBridgeTransactions: latestBridgeTransactions,
    },
  }
}
