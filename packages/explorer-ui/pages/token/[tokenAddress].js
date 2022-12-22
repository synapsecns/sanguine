import { TokenAddress } from '@components/pages/TokenAddress'

import {
  GET_LATEST_BRIDGE_TRANSACTIONS_QUERY,
  BRIDGE_AMOUNT_STATISTIC,
  GET_HISTORICAL_STATS,
  COUNT_BY_TOKEN_ADDRESS,
  COUNT_BY_CHAIN_ID,
  GET_BRIDGE_TRANSACTIONS_QUERY,
} from '@graphql/queries'
import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client'
import { API_URL } from '@graphql'
import numeral from 'numeral'

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

export default function tokenAddressRoute({
  allTimeBridgeVolume,
  allTimeTransactionCount,
  allTimeAddresses,
  bridgeTransactions,
}) {
  return (
    <TokenAddress
      allTimeBridgeVolume={allTimeBridgeVolume}
      allTimeTransactionCount={allTimeTransactionCount}
      allTimeAddresses={allTimeAddresses}
      txQueryResult={bridgeTransactions}
    />
  )
}

export async function getServerSideProps(context) {
  const { tokenAddress, chainId } = context.query
  const { data: allTimeBridgeVolume } = await client.query({
    query: BRIDGE_AMOUNT_STATISTIC,
    variables: {
      chainId: chainId,
      duration: 'ALL_TIME',
      tokenAddress: tokenAddress,
      type: 'TOTAL_VOLUME_USD',
    },
  })

  const { data: allTimeTransactionCount } = await client.query({
    query: BRIDGE_AMOUNT_STATISTIC,
    variables: {
      chainId: chainId,
      duration: 'ALL_TIME',
      tokenAddress: tokenAddress,
      type: 'COUNT_TRANSACTIONS',
    },
  })

  const { data: allTimeAddresses } = await client.query({
    query: BRIDGE_AMOUNT_STATISTIC,
    variables: {
      chainId: chainId,
      duration: 'ALL_TIME',
      tokenAddress: tokenAddress,
      type: 'COUNT_ADDRESSES',
    },
  })

  const { data: bridgeTransactions } = await client.query({
    query: GET_BRIDGE_TRANSACTIONS_QUERY,
    variables: {
      chainId: chainId,
      tokenAddress: tokenAddress,
      page: 1,
    },
  })

  return {
    props: {
      allTimeBridgeVolume: normalizeValue(
        allTimeBridgeVolume?.bridgeAmountStatistic?.value
      ),
      allTimeTransactionCount: normalizeValue(
        allTimeTransactionCount?.bridgeAmountStatistic?.value
      ),
      allTimeAddresses: normalizeValue(
        allTimeAddresses?.bridgeAmountStatistic?.value
      ),
      bridgeTransactions: bridgeTransactions,
    },
  }
}

function normalizeValue(value) {
  if (value >= 1000000000) {
    return (
      numeral(value / 1000000000)
        .format('0.00')
        .toString() + 'B'
    )
  } else if (value >= 1000000) {
    return (
      numeral(value / 1000000)
        .format('0.00')
        .toString() + 'M'
    )
  }
  return numeral(value).format('0,0').toString()
}
