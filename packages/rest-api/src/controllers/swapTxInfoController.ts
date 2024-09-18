import { validationResult } from 'express-validator'
import { parseUnits } from '@ethersproject/units'

import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'

export const swapTxInfoController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }

  try {
    const { chain, amount, address, fromToken, toToken } = req.query

    const fromTokenInfo = tokenAddressToToken(chain.toString(), fromToken)

    const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)

    const quote = await Synapse.swapQuote(
      Number(chain),
      fromToken,
      toToken,
      amountInWei
    )

    const txInfo = await Synapse.swap(
      Number(chain),
      address,
      fromToken,
      amountInWei,
      quote.query
    )

    res.json(txInfo)
  } catch (err) {
    res.status(500).json({
      error:
        'An unexpected error occurred in /swapTxInfo. Please try again later.',
      details: err.message,
    })
  }
}
