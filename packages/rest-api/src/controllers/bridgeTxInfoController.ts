import { validationResult } from 'express-validator'
import { parseUnits } from '@ethersproject/units'

import { Synapse } from '../services/synapseService'

export const bridgeTxInfoController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }

  try {
    const { fromChain, toChain, amount, destAddress } = req.query
    const fromTokenInfo = res.locals.tokenInfo.fromToken
    const toTokenInfo = res.locals.tokenInfo.toToken

    const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)

    const quotes = await Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromTokenInfo.address,
      toTokenInfo.address,
      amountInWei
    )

    const txInfoArray = await Promise.all(
      quotes.map(async (quote) => {
        const txInfo = await Synapse.bridge(
          destAddress,
          quote.routerAddress,
          Number(fromChain),
          Number(toChain),
          fromTokenInfo.address,
          amountInWei,
          quote.originQuery,
          quote.destQuery
        )
        return txInfo
      })
    )
    res.json(txInfoArray)
  } catch (err) {
    res.status(500).json({ error: 'Server error' })
  }
}
