import { validationResult } from 'express-validator'
import { ethers } from 'ethers'

import { getTokenDecimals } from '../utils/getTokenDecimals'

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

    if (toInfo === null) {
      res.json({ status: 'pending' })
    } else {
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
        status: 'completed',
        toInfo: {
          ...restToInfo,
          formattedValue: `${formattedValue}`,
        },
      })
    }
  } catch (err) {
    res.status(500).json({
      error: 'An unexpected error occurred in /getDestinationTx. Please try again later.',
      details: err.message,
    })
  }
}
