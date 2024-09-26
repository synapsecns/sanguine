import { validationResult } from 'express-validator'
import { parseUnits } from '@ethersproject/units'

import { formatBNToString } from '../utils/formatBNToString'
import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'

export const bridgeController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const { fromChain, toChain, amount, fromToken, toToken } = req.query

    const fromTokenInfo = tokenAddressToToken(fromChain.toString(), fromToken)
    const toTokenInfo = tokenAddressToToken(toChain.toString(), toToken)

    const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)

    const resp = await Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromToken,
      toToken,
      amountInWei
    )

    const payload = resp.map((quote) => {
      const originQueryTokenOutInfo = tokenAddressToToken(
        fromChain.toString(),
        quote.originQuery.tokenOut
      )
      return {
        ...quote,
        maxAmountOutStr: formatBNToString(
          quote.maxAmountOut,
          toTokenInfo.decimals
        ),
        bridgeFeeFormatted: formatBNToString(
          quote.feeAmount,
          originQueryTokenOutInfo.decimals
        ),
      }
    })
    res.json(payload)
  } catch (err) {
    res.status(500).json({
      error: 'An unexpected error occurred in /bridge. Please try again later.',
      details: err.message,
    })
  }
}
