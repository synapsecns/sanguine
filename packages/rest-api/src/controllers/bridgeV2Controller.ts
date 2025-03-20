import { validationResult } from 'express-validator'

import { Synapse } from '../services/synapseService'
import { logger } from '../middleware/logger'
import {
  DEFAULT_SWAP_SLIPPAGE_BIPS,
  SLIPPAGE_BIPS_DENOMINATOR,
} from '../constants'
import { stringifyTxValue } from '../utils/replaceTxValue'

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
      fromSender: string
      toChainId: number
      toToken: string
      toRecipient: string
      slippage: number
    }

    // Convert percentage slippage to bips
    const slippageBips = slippage ? slippage * 100 : DEFAULT_SWAP_SLIPPAGE_BIPS

    const allQuotes = await Synapse.bridgeV2({
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
    })

    // Convert all BigNumber values to strings.
    const payload = allQuotes.map((quote) => {
      return {
        ...quote,
        fromAmount: quote.fromAmount.toString(),
        expectedToAmount: quote.expectedToAmount.toString(),
        minToAmount: quote.minToAmount.toString(),
        gasDropAmount: quote.gasDropAmount.toString(),
        callData: stringifyTxValue({
          tx: quote.tx,
          preserveTx: !!fromSender && !!toRecipient,
        }),
        tx: undefined,
      }
    })

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
