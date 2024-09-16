import { validationResult } from 'express-validator'
import { parseUnits } from '@ethersproject/units'
import { isAddress } from 'ethers/lib/utils'

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

    if (!isAddress(fromToken) || !isAddress(toToken)) {
      return res.status(400).json({ error: 'Invalid token address' })
    }

    const fromTokenInfo = tokenAddressToToken(fromChain.toString(), fromToken)
    const toTokenInfo = tokenAddressToToken(toChain.toString(), toToken)

    if (!fromTokenInfo || !toTokenInfo) {
      return res
        .status(400)
        .json({ error: 'Token not supported on specified chain' })
    }

    const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)

    const resp = await Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromToken,
      toToken,
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
