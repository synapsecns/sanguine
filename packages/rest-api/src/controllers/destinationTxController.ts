import { validationResult } from 'express-validator'
import { ethers } from 'ethers'

import { getTokenDecimals } from '../utils/getTokenDecimals'

export const destinationTxController = async (req, res) => {
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
    const toInfo = graphqlData.data.bridgeTransactions[0]?.toInfo || null

    if (toInfo) {
      const { tokenAddress, value, ...restToInfo } = toInfo

      const tokenDecimals = getTokenDecimals(toInfo.chainID, tokenAddress)
      const formattedValue = ethers.utils.formatUnits(value, tokenDecimals)

      res.json({
        status: 'completed',
        toInfo: {
          ...restToInfo,
          formattedValue: `${formattedValue}`,
        },
      })
    } else {
      res.json({ status: 'pending', toInfo: null })
    }
  } catch (err) {
    res.status(500).json({
      error:
        'An unexpected error occurred in /destinationTx. Please try again later.',
      details: err.message,
    })
  }
}
