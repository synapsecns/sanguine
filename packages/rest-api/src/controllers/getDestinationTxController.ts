import { validationResult } from 'express-validator'

export const getDestinationTxController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }

  try {
    const { originChainId, txHash } = req.query

    const graphqlEndpoint = 'https://explorer.omnirpc.io/graphql'
    const graphqlQuery = `
        {
          bridgeTransactions(
            useMv: true
            chainIDFrom: ${originChainId}
            txnHash: "${txHash}"
          ) {
            toInfo {
              chainID
              address
              txnHash
              USDValue
              tokenSymbol
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
    const toInfo = graphqlData.data.bridgeTransactions[0]?.toInfo || null

    if (toInfo === null) {
      res.json({ status: 'pending' })
    } else {
      res.json({ status: 'completed', toInfo })
    }
  } catch (err) {
    res.status(500).json({ error: 'Server error' })
  }
}
