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

  return graphqlData
}
