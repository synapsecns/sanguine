import { validationResult } from 'express-validator'
import { parseUnits } from '@ethersproject/units'
import { isAddress } from 'ethers/lib/utils'

import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'

export const bridgeTxInfoController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }

  try {
    const { fromChain, toChain, amount, destAddress, fromToken, toToken } =
      req.query

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

    const quotes = await Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromToken,
      toToken,
      amountInWei
    )

    const txInfoArray = await Promise.all(
      quotes.map(async (quote) => {
        const txInfo = await Synapse.bridge(
          destAddress,
          quote.routerAddress,
          Number(fromChain),
          Number(toChain),
          fromToken,
          amountInWei,
          quote.originQuery,
          quote.destQuery
        )
        return txInfo
      })
    )
    res.json(txInfoArray)
  } catch (err) {
    res.status(500).json({
      error:
        'An unexpected error occurred in /bridgeTxInfo. Please try again later.',
      details: err.message,
    })
  }
}
