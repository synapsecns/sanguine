import { validationResult } from 'express-validator'
import { parseUnits } from '@ethersproject/units'

import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'
import { formatBNToString } from '../utils/formatBNToString'

export const getBridgeLimitsController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const { fromChain, fromToken, toChain, toToken } = req.query

    const fromTokenInfo = tokenAddressToToken(fromChain, fromToken)
    const toTokenInfo = tokenAddressToToken(toChain, toToken)

    const testValues = [
      '0.01',
      '0.1',
      '1',
      '10',
      '100',
      '1000',
      '10000',
      '100000',
      '1000000',
    ]
    let smallestBridgeQuotes = null

    for (const value of testValues) {
      const amount = parseUnits(value, fromTokenInfo.decimals)

      smallestBridgeQuotes = await Synapse.allBridgeQuotes(
        Number(fromChain),
        Number(toChain),
        fromTokenInfo.address,
        toTokenInfo.address,
        amount
      )

      if (smallestBridgeQuotes && smallestBridgeQuotes.length > 0) {
        break
      }
    }

    const minAmountQuote = smallestBridgeQuotes.reduce(
      (minQuote, currentQuote) => {
        const currentFeeAmount = currentQuote.feeAmount
        const minFeeAmount = minQuote ? minQuote.feeAmount : null

        return !minFeeAmount || currentFeeAmount.lt(minFeeAmount)
          ? currentQuote
          : minQuote
      },
      null
    )

    if (!minAmountQuote) {
      return res.status(400).json({ errors: 'Route does not exist' })
    }

    const minAmountQuoteOriginQueryTokenOutInfo = tokenAddressToToken(
      fromChain,
      minAmountQuote.originQuery.tokenOut
    )

    const minOriginAmount = formatBNToString(
      minAmountQuote.feeAmount,
      minAmountQuoteOriginQueryTokenOutInfo.decimals
    )

    return res.json({
      maxOriginAmount: '1000000',
      minOriginAmount,
    })
  } catch (err) {
    res.status(500).json({
      error:
        'An unexpected error occurred in /getBridgeLimits. Please try again later.',
      details: err.message,
    })
  }
}
