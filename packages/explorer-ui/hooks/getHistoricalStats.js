import {useQuery} from '@apollo/client'

import {GET_HISTORICAL_STATS} from '@graphql/queries'

export function getHistoricalBridgeVolume({ chainId = null, days = 30 }) {
  const { data } = useQuery(GET_HISTORICAL_STATS, {
    variables: {
      chainId,
      type: 'BRIDGEVOLUME',
      days,
    },
  })

  return data
}

export function getHistoricalTransactions({ chainId = null, days = 30 }) {
  const { data } = useQuery(GET_HISTORICAL_STATS, {
    variables: {
      chainId,
      type: 'TRANSACTIONS',
      days,
    },
  })

  return data
}

export function getHistoricalAddresses({ chainId = null, days = 30 }) {
  const { data } = useQuery(GET_HISTORICAL_STATS, {
    variables: {
      chainId,
      type: 'ADDRESSES',
      days,
    },
  })

  return data
}
