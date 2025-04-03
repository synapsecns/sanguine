import { validationResult } from 'express-validator'

import { Synapse } from '../services/synapseService'
import { logger } from '../middleware/logger'
import { DEFAULT_SWAP_SLIPPAGE_PERCENTAGE } from '../constants'
import { formatTransactionData } from '../utils/formatTransactionData'

export const intentController = async (req, res) => {
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
      allowMultipleTxs,
    } = req.query as {
      fromChainId: number
      fromToken: string
      fromAmount: string
      fromSender?: string
      toChainId: number
      toToken: string
      toRecipient?: string
      slippage?: number
      allowMultipleTxs?: boolean
    }

    const intentQuotes = await Synapse.intent({
      fromChainId,
      fromToken,
      fromAmount,
      fromSender,
      toChainId,
      toToken,
      toRecipient,
      slippagePercentage: slippage ?? DEFAULT_SWAP_SLIPPAGE_PERCENTAGE,
      allowMultipleTxs,
    })

    // Include callData only if both fromSender and toRecipient are provided.
    const payload = intentQuotes.map((quote) => ({
      ...quote,
      steps: quote.steps.map((step) =>
        formatTransactionData(step, !!fromSender && !!toRecipient)
      ),
    }))

    logger.info(`Successful intentController response`, {
      query: req.query,
      payload,
    })
    res.json(payload)
  } catch (err) {
    logger.error(`Error in intentController`, {
      query: req.query,
      error: err.message,
      stack: err.stack,
    })
    res.status(500).json({
      error: 'An unexpected error occurred in /intent. Please try again later.',
    })
  }
}
