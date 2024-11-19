import { validationResult } from 'express-validator'
import { formatUnits, parseUnits } from '@ethersproject/units'
import { BigNumber } from '@ethersproject/bignumber'

import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'
import { logger } from '../middleware/logger'

export const swapController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const { chain, amount, fromToken, toToken, address } = req.query

    const fromTokenInfo = tokenAddressToToken(chain.toString(), fromToken)
    const toTokenInfo = tokenAddressToToken(chain.toString(), toToken)

    const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)
    const quote = await Synapse.swapQuote(
      Number(chain),
      fromToken,
      toToken,
      amountInWei
    )

    const formattedMaxAmountOut = formatUnits(
      BigNumber.from(quote.maxAmountOut),
      toTokenInfo.decimals
    )

    const callData = address
      ? await Synapse.swap(
          Number(chain),
          address,
          fromToken,
          amountInWei,
          quote.query
        )
      : null

    const payload = {
      ...quote,
      maxAmountOut: formattedMaxAmountOut,
      callData,
    }

    logger.info(`Successful swapController response`, {
      query: req.query,
      payload,
    })
    res.json(payload)
  } catch (err) {
    logger.error(`Error in swapController`, {
      query: req.query,
      error: err.message,
      stack: err.stack,
    })
    res.status(500).json({
      error: 'An unexpected error occurred in /swap. Please try again later.',
    })
  }
}
