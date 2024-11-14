import { validationResult } from 'express-validator'
import { ethers } from 'ethers'

import { Synapse } from '../services/synapseService'
import { getTokenDecimals } from '../utils/getTokenDecimals'
import { logger } from '../middleware/logger'
import { fetchBridgeTransaction } from '../utils/fetchBridgeTransaction'

export const bridgeTxStatusController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }

  try {
    const { destChainId, bridgeModule, synapseTxId } = req.query

    const txIdWith0x = !synapseTxId.startsWith('0x')
      ? `0x${synapseTxId}`
      : synapseTxId

    const status = await Synapse.getBridgeTxStatus(
      Number(destChainId),
      bridgeModule,
      txIdWith0x
    )

    if (status) {
      const txIdWithout0x = synapseTxId.startsWith('0x')
        ? synapseTxId.slice(2)
        : synapseTxId

      const params = { useMv: true, kappa: txIdWithout0x }

      const bridgeTxn = await fetchBridgeTransaction(params)

      const { toInfo = null } = bridgeTxn || {}

      if (toInfo) {
        const { tokenAddress, value, ...restToInfo } = toInfo

        const tokenDecimals = getTokenDecimals(toInfo.chainID, tokenAddress)
        const formattedValue = ethers.utils.formatUnits(value, tokenDecimals)

        const payload = {
          status,
          toInfo: {
            ...restToInfo,
            formattedValue: `${formattedValue}`,
          },
        }

        logger.info(`Successful bridgeTxStatusController response`, {
          query: req.query,
          payload,
        })
        res.json(payload)
      } else {
        const payload = {
          status,
          toInfo: null,
        }

        logger.info(`Successful bridgeTxStatusController response`, {
          query: req.query,
          payload,
        })
        res.json(payload)
      }
    } else {
      const payload = { status }

      logger.info(`Successful bridgeTxStatusController response`, {
        query: req.query,
        payload,
      })
      res.json(payload)
    }
  } catch (err) {
    logger.error(`Error in bridgeTxStatusController`, {
      query: req.query,
      error: err.message,
      stack: err.stack,
    })
    res.status(500).json({
      error:
        'An unexpected error occurred in /bridgeTxStatus. Please try again later.',
    })
  }
}
