import { validationResult } from 'express-validator'
import { ethers } from 'ethers'

import { Synapse } from '../services/synapseService'
import { getTokenDecimals } from '../utils/getTokenDecimals'

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
          bridgeTransactions(useMv: true, kappa: "${txIdWithout0x}") {
            toInfo {
              chainID
              address
              txnHash
              value
              USDValue
              tokenSymbol
              tokenAddress
              formattedTime
            }
          }
        }
      `

      const graphqlResponse = await fetch(graphqlEndpoint, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ query: graphqlQuery }),
      })

      const graphqlData = await graphqlResponse.json()
      const toInfo = graphqlData.data.bridgeTransactions[0]?.toInfo

      if (toInfo) {
        const tokenDecimals = getTokenDecimals(
          toInfo.chainID,
          toInfo.tokenAddress
        )
        const formattedValue = ethers.utils.formatUnits(
          toInfo.value,
          tokenDecimals
        )
        // the below line is to deconstruct the toInfo object
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        const { tokenAddress, value, ...restToInfo } = toInfo
        res.json({
          status,
          toInfo: {
            ...restToInfo,
            formattedValue: `${formattedValue}`,
          },
        })
      } else {
        res.json({ status, toInfo: null })
      }
    } else {
      res.json({ status })
    }
  } catch (err) {
    res.status(500).json({
      error:
        'An unexpected error occurred in /getBridgeTxStatus. Please try again later.',
      details: err.message,
    })
  }
}
