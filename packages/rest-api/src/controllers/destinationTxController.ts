import { validationResult } from 'express-validator'
import { ethers } from 'ethers'

import { getTokenDecimals } from '../utils/getTokenDecimals'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'
import { logger } from '../middleware/logger'

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
      const { tokenAddress, value, chainID, ...restToInfo } = toInfo

      const tokenInfo = tokenAddressToToken(chainID.toString(), tokenAddress)
      const tokenDecimals = getTokenDecimals(chainID, tokenAddress)
      const formattedValue = ethers.utils.formatUnits(value, tokenDecimals)

      const payload = {
        status: 'completed',
        toInfo: {
          chainID,
          ...restToInfo,
          tokenSymbol: tokenInfo ? tokenInfo?.symbol : null,
          formattedValue: `${formattedValue}`,
        },
      }

      logger.info(`Successful destinationTxController response`, {
        query: req.query,
        payload,
      })
      res.json(payload)
    } else {
      const payload = {
        status: 'pending',
        toInfo: null,
      }

      logger.info(`Successful destinationTxController response`, {
        query: req.query,
        payload,
      })
      res.json(payload)
    }
  } catch (err) {
    logger.error(`Error in destinationTxController`, {
      query: req.query,
      error: err.message,
      stack: err.stack,
    })
    res.status(500).json({
      error:
        'An unexpected error occurred in /destinationTx. Please try again later.',
      details: err.message,
    })
  }
}
