import { BridgeTransaction } from '../types'

export const fetchBridgeTransaction = async (
  originChainId: number | string,
  txHash: string
) => {
  const graphqlEndpoint = 'https://explorer.omnirpc.io/graphql'
  const graphqlQuery = `
      {
        bridgeTransactions(
          useMv: true
          chainIDFrom: ${originChainId}
          txnHash: "${txHash}"
        ) {
          kappa
          fromInfo {
            chainID
            address
            txnHash
            value
            USDValue
            tokenSymbol
            tokenAddress
            blockNumber
            formattedTime
          }
          toInfo {
            chainID
            address
            txnHash
            value
            USDValue
            tokenSymbol
            tokenAddress
            blockNumber
            formattedTime
          }
        }
      }
    `

  const graphqlResponse = await fetch(graphqlEndpoint, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ query: graphqlQuery }),
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
