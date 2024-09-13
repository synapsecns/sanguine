import { validationResult } from 'express-validator'

import { Synapse } from '../services/synapseService'

export const getBridgeTxStatusController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }

  try {
    const { destChainId, bridgeModule, synapseTxId } = req.query

    const status = await Synapse.getBridgeTxStatus(
      Number(destChainId),
      bridgeModule,
      synapseTxId
    )

    if (status) {
      const txIdWithout0x = synapseTxId.startsWith('0x')
        ? synapseTxId.slice(2)
        : synapseTxId
      const graphqlEndpoint = 'https://explorer.omnirpc.io/graphql'
      const graphqlQuery = `
          {
            bridgeTransactions(
              useMv: true
              kappa: "${txIdWithout0x}"
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

      res.json({ status, toInfo })
    } else {
      res.json({ status })
    }
  } catch (err) {
    res.status(500).json({ error: 'Server error' })
  }
}
