import {Home} from '@components/pages/Home'
import {ApolloClient, HttpLink, InMemoryCache} from '@apollo/client'
import {
  AMOUNT_STATISTIC,
  COUNT_BY_CHAIN_ID,
  COUNT_BY_TOKEN_ADDRESS,
  GET_BRIDGE_TRANSACTIONS_QUERY,
  GET_DAILY_STATS,
} from '@graphql/queries'
import {API_URL} from '@graphql'

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

function Index({
  bridgeVolume,
  transactions,
  addresses,
  latestBridgeTransactions,
  latestBridgeTransactionsPending,
  popularTokens,
  popularChains,
}) {
  return (
    <Home
      bridgeVolume={bridgeVolume}
      transactions={transactions}
      addresses={addresses}
      popularTokens={popularTokens}
      popularChains={popularChains}
    />
  )
}

export default Index

export async function getServerSideProps() {
  const { data: bridgeVolume } = await client.query({
    query: GET_DAILY_STATS,
    variables: {
      chainId: null,
      type: 'VOLUME',
      platform: 'BRIDGE',
      days: 30,
    },
  })

  const { data: transactions } = await client.query({
    query: GET_DAILY_STATS,
    variables: {
      chainId: null,
      type: 'TRANSACTIONS',
      platform: 'BRIDGE',
      days: 30,
    },
  })

  const { data: addresses } = await client.query({
    query: GET_DAILY_STATS,
    variables: {
      chainId: null,
      type: 'ADDRESSES',
      platform: 'BRIDGE',
      days: 30,
    },
  })

  // const { data: bridgeVolumeAllTime } = await client.query({
  //   query: AMOUNT_STATISTIC,
  //   variables: {
  //     type: 'TOTAL_VOLUME_USD',
  //     duration: 'ALL_TIME',
  //     platform: 'BRIDGE',
  //   },
  // })

  // const { data: transactionsAllTime } = await client.query({
  //   query: AMOUNT_STATISTIC,
  //   variables: {
  //     type: 'COUNT_TRANSACTIONS',
  //     duration: 'ALL_TIME',
  //     platform: 'BRIDGE',
  //   },
  // })

  // const { data: addressesAllTime } = await client.query({
  //   query: AMOUNT_STATISTIC,
  //   variables: {
  //     type: 'COUNT_ADDRESSES',
  //     duration: 'ALL_TIME',
  //     platform: 'BRIDGE',
  //   },
  // })

  const { data: latestBridgeTransactions } = await client.query({
    query: GET_BRIDGE_TRANSACTIONS_QUERY,
    variables: {
      pending: false,
      page: 1,
    },
  })
  const { data: latestBridgeTransactionsPending } = await client.query({
    query: GET_BRIDGE_TRANSACTIONS_QUERY,
    variables: {
      pending: true,
      page: 1,
    },
  })

  const { data: popularTokens } = await client.query({
    query: COUNT_BY_TOKEN_ADDRESS,
    variables: {
      direction: 'IN',
      hours: 720,
    },
  })
  const { countByTokenAddress } = popularTokens ?? {}

  const { data: popularChains } = await client.query({
    query: COUNT_BY_CHAIN_ID,
    variables: {
      direction: 'IN',
      hours: 720,
    },
  })

  const { countByChainId } = popularChains ?? {}

  return {
    props: {
      bridgeVolume: bridgeVolume,
      transactions: transactions,
      addresses: addresses,
      popularTokens: countByTokenAddress,
      popularChains: countByChainId,
    }, // will be passed to the page component as props
  }
}
