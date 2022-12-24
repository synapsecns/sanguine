import {useQuery} from '@apollo/client'

import {BRIDGE_AMOUNT_STATISTIC} from '@graphql/queries'

export function getBridgeVolume({ chainId = null, duration = "ALL_TIME" }) {
  const { data } = useQuery(BRIDGE_AMOUNT_STATISTIC, {
    variables: {
      type: 'TOTAL_VOLUME_USD',
      duration,
    },
  })

  return data
}

export function getTransactions({ chainId = null, duration = "ALL_TIME" }) {
  const { data } = useQuery(BRIDGE_AMOUNT_STATISTIC, {
    variables: {
      type: 'COUNT_TRANSACTIONS',
      duration,
    },
  })

  return data
}

export function getAddresses({ chainId = null, duration = "ALL_TIME" }) {
  const { data } = useQuery(BRIDGE_AMOUNT_STATISTIC, {
    variables: {
      type: 'COUNT_ADDRESSES',
      duration,
    },
  })

  return data
}
