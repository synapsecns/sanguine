import { validationResult } from 'express-validator'

import { Synapse } from '../services/synapseService'
import { logger } from '../middleware/logger'
import { DEFAULT_SWAP_SLIPPAGE_PERCENTAGE } from '../constants'
import { formatTransactionData } from '../utils/formatTransactionData'

export const bridgeV2Controller = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const {
      fromChainId,
      fromToken,
      fromAmount,
      fromSender,
      toChainId,
      toToken,
      toRecipient,
      slippage,
    } = req.query as {
      fromChainId: number
      fromToken: string
      fromAmount: string
      fromSender?: string
      toChainId: number
      toToken: string
      toRecipient?: string
      slippage?: number
    }

    const allQuotes = await Synapse.bridgeV2({
      fromChainId,
      fromToken,
      fromAmount,
      fromSender,
      toChainId,
      toToken,
      toRecipient,
      slippagePercentage: slippage ?? DEFAULT_SWAP_SLIPPAGE_PERCENTAGE,
    })

    // Check if no bridge quotes were found
    if (!allQuotes || allQuotes.length === 0) {
      logger.info(`No bridge routes found`, {
        query: req.query,
      })
      return res.status(404).json({ error: 'No bridge routes found' })
    }

    // Include callData only if both fromSender and toRecipient are provided.
    const payload = allQuotes.map((quote) =>
      formatTransactionData(quote, !!fromSender && !!toRecipient)
    )

    logger.info(`Successful bridgeV2Controller response`, {
      payload,
      query: req.query,
    })
    res.json(payload)
  } catch (err) {
    logger.error(`Error in bridgeV2Controller`, {
      query: req.query,
      error: err.message,
      stack: err.stack,
    })
    res.status(500).json({
      error:
        'An unexpected error occurred in /bridge/v2. Please try again later.',
    })
  }
}
