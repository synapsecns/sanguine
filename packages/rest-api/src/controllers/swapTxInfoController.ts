import { validationResult } from 'express-validator'
import { parseUnits } from '@ethersproject/units'

import { Synapse } from '../services/synapseService'

export const swapTxInfoController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }

  try {
    const { chain, amount } = req.query
    const fromTokenInfo = res.locals.tokenInfo.fromToken
    const toTokenInfo = res.locals.tokenInfo.toToken

    if (!fromTokenInfo || !toTokenInfo) {
      return res.status(400).json({ error: 'Invalid token symbol' })
    }

    const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)

    const quote = await Synapse.swapQuote(
      Number(chain),
      fromTokenInfo.address,
      toTokenInfo.address,
      amountInWei
    )

    const txInfo = await Synapse.swap(
      Number(chain),
      fromTokenInfo.address,
      toTokenInfo.address,
      amountInWei,
      quote.query
    )

    res.json(txInfo)
  } catch (err) {
    res.status(500).json({ error: 'Server error' })
  }
}
