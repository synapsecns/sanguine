import { validationResult } from 'express-validator'

import { Synapse } from '../services/synapseService'
import { logger } from '../middleware/logger'
import {
  DEFAULT_SWAP_SLIPPAGE_BIPS,
  SLIPPAGE_BIPS_DENOMINATOR,
} from '../constants'

export const bridgeV2Controller = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const {
      fromChain,
      toChain,
      amount,
      fromToken,
      toToken,
      originUserAddress,
      destAddress,
      slippage,
    } = req.query as {
      fromChain: string
      toChain: string
      amount: string
      fromToken: string
      toToken: string
      originUserAddress?: string
      destAddress?: string
      slippage?: string
    }

    // Convert percentage slippage to bips
    const slippageBips = slippage
      ? Number(slippage) * 100
      : DEFAULT_SWAP_SLIPPAGE_BIPS

    const allQuotes = await Synapse.bridgeV2({
      originChainId: Number(fromChain),
      destChainId: Number(toChain),
      tokenIn: fromToken,
      tokenOut: toToken,
      amountIn: amount,
      originSender: originUserAddress,
      destRecipient: destAddress,
      slippage: {
        numerator: slippageBips,
        denominator: SLIPPAGE_BIPS_DENOMINATOR,
      },
    })

    const payload = allQuotes.map((quote) => {
      const callData =
        destAddress && originUserAddress && quote.tx
          ? {
              ...quote.tx,
              value: quote.tx.value.toString(),
            }
          : null
      return {
        ...quote,
        callData,
        maxAmountOutStr: quote.maxAmountOut.toString(),
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
