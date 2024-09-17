import { validationResult } from 'express-validator'
import { parseUnits } from '@ethersproject/units'

import { formatBNToString } from '../utils/formatBNToString'
import { Synapse } from '../services/synapseService'

export const bridgeController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const { fromChain, toChain, amount } = req.query
    const fromTokenInfo = res.locals.tokenInfo.fromToken
    const toTokenInfo = res.locals.tokenInfo.toToken

    const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)

    const resp = await Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromTokenInfo.address,
      toTokenInfo.address,
      amountInWei
    )
    const payload = resp.map((quote) => ({
      ...quote,
      maxAmountOutStr: formatBNToString(
        quote.maxAmountOut,
        toTokenInfo.decimals
      ),
      bridgeFeeFormatted: formatBNToString(
        quote.feeAmount,
        toTokenInfo.decimals
      ),
    }))
    res.json(payload)
  } catch (err) {
    res.status(500).json({
      error: 'An unexpected error occurred in /bridge. Please try again later.',
      details: err.message,
    })
  }
}
