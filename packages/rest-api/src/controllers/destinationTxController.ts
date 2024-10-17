import { validationResult } from 'express-validator'

import { logger } from '../middleware/logger'
import { getBridgeStatus } from '../utils/getBridgeStatus'
import { BridgeStatus } from '../constants/enums'
import { fetchBridgeTransaction } from '../utils/fetchBridgeTransaction'
import { serializeBridgeInfo } from '../serializers/serializeBridgeInfo'

export const destinationTxController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }

  let payload = {}

  try {
    const { originChainId, txHash } = req.query

    const bridgeTxn = await fetchBridgeTransaction(originChainId, txHash)

    const { fromInfo = null, toInfo = null, kappa = null } = bridgeTxn || {}

    if (!bridgeTxn) {
      payload = {
        status: 'not found',
        fromInfo,
        toInfo,
      }

      logger.info(`Successful destinationTxController response`, {
        query: req.query,
        payload,
      })
      return res.status(404).json(payload)
    }

    if (fromInfo && !toInfo) {
      const status = await getBridgeStatus(originChainId, kappa)

      if (status === BridgeStatus.REFUNDED) {
        payload = {
          status: 'refunded',
          fromInfo: serializeBridgeInfo(fromInfo),
          toInfo,
        }
      } else {
        payload = {
          status: 'pending',
          fromInfo: serializeBridgeInfo(fromInfo),
          toInfo,
        }
      }
    }

    if (fromInfo && toInfo) {
      payload = {
        status: 'completed',
        fromInfo: serializeBridgeInfo(fromInfo),
        toInfo: serializeBridgeInfo(toInfo),
      }
    }

    logger.info(`Successful destinationTxController response`, {
      query: req.query,
      payload,
    })

    return res.json(payload)
  } catch (err) {
    logger.error(`Error in destinationTxController`, {
      query: req.query,
      error: err.message,
      stack: err.stack,
    })
    return res.status(500).json({
      error:
        'An unexpected error occurred in /destinationTx. Please try again later.',
    })
  }
}
