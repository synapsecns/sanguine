import _ from 'lodash'
import { useQuery } from '@apollo/client'

import {
  TransactionCard,
  TransactionCardLoader,
} from '@components/TransactionCard'
import { GET_LATEST_BRIDGE_TRANSACTIONS_QUERY } from '@graphql/queries'

export function LatestBridgeTransactions() {
  const { error, data } = useQuery(GET_LATEST_BRIDGE_TRANSACTIONS_QUERY, {
    variables: { includePending: false, page: 1 },
  })

  let content

  if (!data) {
    content = [...Array(5).keys()].map((i) => (
      <TransactionCardLoader key={i} ordinal={i} />
    ))
  } else if (error) {
    content = 'Error'
  } else {
    let { latestBridgeTransactions } = data

    latestBridgeTransactions = _.orderBy(
      latestBridgeTransactions,
      'fromInfo.time',
      ['desc']
    ).slice(0, 5)

    content = latestBridgeTransactions.map((txn, i) => (
      <TransactionCard txn={txn} key={i} ordinal={i} />
    ))
  }

  return <>{content}</>
}
