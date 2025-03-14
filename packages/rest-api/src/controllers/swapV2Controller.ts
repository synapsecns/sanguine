import { validationResult } from 'express-validator'

import { Synapse } from '../services/synapseService'
import { logger } from '../middleware/logger'
import {
  DEFAULT_SWAP_SLIPPAGE_BIPS,
  SLIPPAGE_BIPS_DENOMINATOR,
} from '../constants'

export const swapV2Controller = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const { chain, amount, fromToken, toToken, address, slippage } =
      req.query as {
        chain: string
        amount: string
        fromToken: string
        toToken: string
        address?: string
        slippage?: string
      }

    // Convert percentage slippage to bips
    const slippageBips = slippage
      ? Number(slippage) * 100
      : DEFAULT_SWAP_SLIPPAGE_BIPS

    const { routerAddress, maxAmountOut, tx } = await Synapse.swapV2({
      chainId: Number(chain),
      tokenIn: fromToken,
      tokenOut: toToken,
      amountIn: amount,
      to: address,
      slippage: {
        numerator: slippageBips,
        denominator: SLIPPAGE_BIPS_DENOMINATOR,
      },
    })

    const callData =
      address && tx
        ? {
            ...tx,
            value: tx.value.toString(),
          }
        : null

    const payload = {
      routerAddress,
      maxAmountOut: maxAmountOut.toString(),
      callData,
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
      error: 'An unexpected error occurred in /swapV2. Please try again later.',
    })
  }
}
