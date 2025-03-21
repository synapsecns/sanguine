import { validationResult } from 'express-validator'

import { Synapse } from '../services/synapseService'
import { logger } from '../middleware/logger'
import {
  DEFAULT_SWAP_SLIPPAGE_BIPS,
  SLIPPAGE_BIPS_DENOMINATOR,
} from '../constants'
import { stringifyTxValue } from '../utils/replaceTxValue'

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

    // Convert percentage slippage to bips
    const slippageBips = slippage ? slippage * 100 : DEFAULT_SWAP_SLIPPAGE_BIPS

    const intentQuotes = await Synapse.intent({
      fromChainId,
      fromToken,
      fromAmount,
      fromSender,
      toChainId,
      toToken,
      toRecipient,
      slippage: {
        numerator: slippageBips,
        denominator: SLIPPAGE_BIPS_DENOMINATOR,
      },
      allowMultipleTxs,
    })

    // Convert all BigNumber values to strings.
    const payload = intentQuotes.map((quote) => ({
      ...quote,
      fromAmount: quote.fromAmount.toString(),
      expectedToAmount: quote.expectedToAmount.toString(),
      minToAmount: quote.minToAmount.toString(),
      steps: quote.steps.map((step) => ({
        ...step,
        fromAmount: step.fromAmount.toString(),
        expectedToAmount: step.expectedToAmount.toString(),
        minToAmount: step.minToAmount.toString(),
        gasDropAmount: step.gasDropAmount.toString(),
        callData: stringifyTxValue({
          tx: step.tx,
          preserveTx: !!fromSender && !!toRecipient,
        }),
        tx: undefined,
      })),
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
