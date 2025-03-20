import { validationResult } from 'express-validator'

import { Synapse } from '../services/synapseService'
import { logger } from '../middleware/logger'
import {
  DEFAULT_SWAP_SLIPPAGE_BIPS,
  SLIPPAGE_BIPS_DENOMINATOR,
} from '../constants'
import { stringifyTxValue } from '../utils/replaceTxValue'

export const swapV2Controller = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const { chainId, fromToken, toToken, fromAmount, toRecipient, slippage } =
      req.query as {
        chainId: number
        fromToken: string
        toToken: string
        fromAmount: string
        toRecipient?: string
        slippage?: number
      }

    // Convert percentage slippage to bips
    const slippageBips = slippage ? slippage * 100 : DEFAULT_SWAP_SLIPPAGE_BIPS

    const quote = await Synapse.swapV2({
      chainId,
      fromToken,
      toToken,
      fromAmount,
      toRecipient,
      slippage: {
        numerator: slippageBips,
        denominator: SLIPPAGE_BIPS_DENOMINATOR,
      },
    })

    // Convert all BigNumber values to strings.
    const payload = {
      ...quote,
      fromAmount: quote.fromAmount.toString(),
      expectedToAmount: quote.expectedToAmount.toString(),
      minToAmount: quote.minToAmount.toString(),
      callData: stringifyTxValue({
        tx: quote.tx,
        preserveTx: !!toRecipient,
      }),
      tx: undefined,
    }

    logger.info(`Successful swapV2Controller response`, {
      query: req.query,
      payload,
    })
    res.json(payload)
  } catch (err) {
    logger.error(`Error in swapV2Controller`, {
      query: req.query,
      error: err.message,
      stack: err.stack,
    })
    res.status(500).json({
      error:
        'An unexpected error occurred in /swap/v2. Please try again later.',
    })
  }
}
