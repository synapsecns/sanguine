import { validationResult } from 'express-validator'
import { formatUnits, parseUnits } from '@ethersproject/units'
import { isAddress } from 'ethers/lib/utils'

import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'

export const swapController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const { chain, amount, fromToken, toToken } = req.query

    if (!isAddress(fromToken) || !isAddress(toToken)) {
      return res.status(400).json({ error: 'Invalid token address' })
    }

    const fromTokenInfo = tokenAddressToToken(chain.toString(), fromToken)
    const toTokenInfo = tokenAddressToToken(chain.toString(), toToken)

    if (!fromTokenInfo || !toTokenInfo) {
      return res
        .status(400)
        .json({ error: 'Token not supported on specified chain' })
    }

    const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)
    const quote = await Synapse.swapQuote(
      Number(chain),
      fromToken,
      toToken,
      amountInWei
    )
    res.json({
      maxAmountOut: formatUnits(quote.maxAmountOut, toTokenInfo.decimals),
      ...quote,
    })
  } catch (err) {
    res.status(500).json({
      error: 'An unexpected error occurred in /swap. Please try again later.',
      details: err.message,
    })
  }
}
