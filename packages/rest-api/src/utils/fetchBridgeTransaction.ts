import { EXPLORER_GRAPHQL_URL } from '../constants'
import { BridgeTransaction } from '../types'
import { constructBridgeTransactionsQuery } from './constructBridgeTransactionsQuery'

export const fetchBridgeTransaction = async ({
  originChainId,
  txnHash,
  kappa,
}: {
  originChainId?: number | string
  txnHash?: string | null
  kappa?: string | null
}) => {
  const params = { useMv: true, originChainId, txnHash, kappa }

  const query = constructBridgeTransactionsQuery(params)

  const graphqlResponse = await fetch(EXPLORER_GRAPHQL_URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ query }),
  })

  const graphqlData = await graphqlResponse.json()

  if (
    graphqlData.data &&
    graphqlData.data.bridgeTransactions &&
    graphqlData.data.bridgeTransactions.length > 0
  ) {
    return graphqlData.data.bridgeTransactions[0] as BridgeTransaction
  }

  return null
}
