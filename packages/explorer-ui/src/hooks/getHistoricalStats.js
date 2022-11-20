import _ from 'lodash'
import { useQuery } from '@apollo/client'

import { GET_HISTORICAL_STATS } from '@graphql/queries'

export function getBridgeVolume({ chainId = null, days = 30 }) {
  const { data } = useQuery(GET_HISTORICAL_STATS, {
    variables: {
      chainId,
      type: 'BRIDGEVOLUME',
      days,
    },
  })

  return data
}

export function getTransactions({ chainId = null, days = 30 }) {
  const { data } = useQuery(GET_HISTORICAL_STATS, {
    variables: {
      chainId,
      type: 'TRANSACTIONS',
      days,
    },
  })

  return data
}

export function getAddresses({ chainId = null, days = 30 }) {
  const { data } = useQuery(GET_HISTORICAL_STATS, {
    variables: {
      chainId,
      type: 'ADDRESSES',
      days,
    },
  })

  return data
}
