import { useQuery } from '@apollo/client'
import { GET_BRIDGE_TRANSACTIONS_QUERY } from '@graphql/queries'

import Grid from '@components/tailwind/Grid'

import { TransactionCard } from '@components/TransactionCard'

export function SearchResults({ searchField, inputType }) {
  const { error, data } = useQuery(GET_BRIDGE_TRANSACTIONS_QUERY, {
    variables: {
      txnHash: inputType == 'TRANSACTION' ? searchField : undefined,
      address: inputType == 'ADDRESS' ? searchField : undefined,
      chainId: inputType == 'CHAIN' ? parseInt(searchField) : undefined,
    },
    skip: !(inputType && searchField?.length > 0),
  })

  return (
    <Grid cols={{ xs: 1 }} gap={4}>
      {data?.bridgeTransactions.map((txn) => {
        return <TransactionCard txn={txn} key={txn.txnHash} />
      })}
      <div>{error}</div>
    </Grid>
  )
}
