import { Chain } from '@components/pages/Chain'

import {
  GET_HISTORICAL_STATS,
  GET_BRIDGE_TRANSACTIONS_QUERY,
} from '@graphql/queries'
import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client'
import { API_URL } from '@graphql'

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
  return (
    <Chain
      bridgeVolume={bridgeVolume}
      transactionsData={transactionsData}
      addresses={addresses}
      txQueryResult={latestBridgeTransactions}
    />
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
