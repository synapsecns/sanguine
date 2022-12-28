import {Home} from '@components/pages/Home'
import {ApolloClient, HttpLink, InMemoryCache} from '@apollo/client'
import {
  BRIDGE_AMOUNT_STATISTIC,
  COUNT_BY_CHAIN_ID,
  COUNT_BY_TOKEN_ADDRESS,
  GET_BRIDGE_TRANSACTIONS_QUERY,
  GET_HISTORICAL_STATS,
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
  bridgeVolumeAllTime,
  transactionsAllTime,
  addressesAllTime,
  latestBridgeTransactions,
  popularTokens,
  popularChains,
}) {
  return (
    <Home
      bridgeVolume={bridgeVolume}
      transactions={transactions}
      addresses={addresses}
      bridgeVolumeAllTime={bridgeVolumeAllTime}
      transactionsAllTime={transactionsAllTime}
      addressesAllTime={addressesAllTime}
      latestBridgeTransactions={latestBridgeTransactions}
      popularTokens={popularTokens}
      popularChains={popularChains}
    />
  )
}

export default Index

export async function getServerSideProps() {
  const { data: bridgeVolume } = await client.query({
    query: GET_HISTORICAL_STATS,
    variables: {
      chainId: null,
      type: 'BRIDGEVOLUME',
      days: 30,
    },
  })

  const { data: transactions } = await client.query({
    query: GET_HISTORICAL_STATS,
    variables: {
      chainId: null,
      type: 'TRANSACTIONS',
      days: 30,
    },
  })

  const { data: addresses } = await client.query({
    query: GET_HISTORICAL_STATS,
    variables: {
      chainId: null,
      type: 'ADDRESSES',
      days: 30,
    },
  })

  const { data: bridgeVolumeAllTime } = await client.query({
    query: BRIDGE_AMOUNT_STATISTIC,
    variables: {
      type: 'TOTAL_VOLUME_USD',
      duration: 'ALL_TIME',
    },
  })

  const { data: transactionsAllTime } = await client.query({
    query: BRIDGE_AMOUNT_STATISTIC,
    variables: {
      type: 'COUNT_TRANSACTIONS',
      duration: 'ALL_TIME',
    },
  })

  const { data: addressesAllTime } = await client.query({
    query: BRIDGE_AMOUNT_STATISTIC,
    variables: {
      type: 'COUNT_ADDRESSES',
      duration: 'ALL_TIME',
    },
  })

  const { data: latestBridgeTransactions } = await client.query({
    query: GET_BRIDGE_TRANSACTIONS_QUERY,
    variables: {
      includePending: false,
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
      bridgeVolumeAllTime: bridgeVolumeAllTime,
      transactionsAllTime: transactionsAllTime,
      addressesAllTime: addressesAllTime,
      latestBridgeTransactions: latestBridgeTransactions,
      popularTokens: countByTokenAddress,
      popularChains: countByChainId,
    }, // will be passed to the page component as props
  }
}
