import { validationResult } from 'express-validator'
import { formatUnits, parseUnits } from '@ethersproject/units'
import { BigNumber } from '@ethersproject/bignumber'

import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'
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

    const fromTokenInfo = tokenAddressToToken(chain.toString(), fromToken)
    const toTokenInfo = tokenAddressToToken(chain.toString(), toToken)

    const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)
    // Convert percentage slippage to bips
    const slippageBips = slippage
      ? Number(slippage) * 100
      : DEFAULT_SWAP_SLIPPAGE_BIPS

    const { routerAddress, maxAmountOut, tx } = await Synapse.swapV2(
      Number(chain),
      fromToken,
      toToken,
      amountInWei,
      {
        originUserAddress: address,
        to: address,
        slippage: {
          numerator: slippageBips,
          denominator: SLIPPAGE_BIPS_DENOMINATOR,
        },
      }
    )

    const formattedMaxAmountOut = formatUnits(
      BigNumber.from(maxAmountOut),
      toTokenInfo.decimals
    )

    const callData = address ? tx : null

    const payload = {
      routerAddress,
      maxAmountOut: formattedMaxAmountOut,
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
