import { validationResult } from 'express-validator'
import { BigNumber } from 'ethers'
import { parseUnits } from '@ethersproject/units'

import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'
import { formatBNToString } from '../utils/formatBNToString'

export const bridgeLimitsController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const { fromChain, fromToken, toChain, toToken } = req.query

    const fromTokenInfo = tokenAddressToToken(fromChain, fromToken)
    const toTokenInfo = tokenAddressToToken(toChain, toToken)

    const upperLimitValue = parseUnits('1000000', fromTokenInfo.decimals)
    const upperLimitBridgeQuotes = await Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromTokenInfo.address,
      toTokenInfo.address,
      upperLimitValue
    )

    const lowerLimitValues = ['0.01', '10']
    let lowerLimitBridgeQuotes = null

    for (const limit of lowerLimitValues) {
      const lowerLimitAmount = parseUnits(limit, fromTokenInfo.decimals)

      lowerLimitBridgeQuotes = await Synapse.allBridgeQuotes(
        Number(fromChain),
        Number(toChain),
        fromTokenInfo.address,
        toTokenInfo.address,
        lowerLimitAmount
      )

      if (lowerLimitBridgeQuotes && lowerLimitBridgeQuotes.length > 0) {
        break
      }
    }

    const maxBridgeAmountQuote = upperLimitBridgeQuotes.reduce(
      (maxQuote, currentQuote) => {
        const currentMaxAmount = currentQuote.maxAmountOut
        const maxAmount = maxQuote ? maxQuote.maxAmountOut : BigNumber.from(0)

        return currentMaxAmount.gt(maxAmount) ? currentQuote : maxQuote
      },
      null
    )

    const minBridgeAmountQuote = lowerLimitBridgeQuotes.reduce(
      (minQuote, currentQuote) => {
        const currentFeeAmount = currentQuote.feeAmount
        const minFeeAmount = minQuote ? minQuote.feeAmount : null

        return !minFeeAmount || currentFeeAmount.lt(minFeeAmount)
          ? currentQuote
          : minQuote
      },
      null
    )

    if (!maxBridgeAmountQuote || !minBridgeAmountQuote) {
      return res.json({
        maxOriginAmount: null,
        minOriginAmount: null,
      })
    }

    const maxAmountOriginQueryTokenOutInfo = tokenAddressToToken(
      toChain,
      maxBridgeAmountQuote.destQuery.tokenOut
    )

    const minAmountOriginQueryTokenOutInfo = tokenAddressToToken(
      fromChain,
      minBridgeAmountQuote.originQuery.tokenOut
    )

    const maxOriginAmount = formatBNToString(
      maxBridgeAmountQuote.maxAmountOut,
      maxAmountOriginQueryTokenOutInfo.decimals
    )

    const minOriginAmount = formatBNToString(
      minBridgeAmountQuote.feeAmount,
      minAmountOriginQueryTokenOutInfo.decimals
    )

    return res.json({
      maxOriginAmount,
      minOriginAmount,
    })
  } catch (err) {
    res.status(500).json({
      error:
        'An unexpected error occurred in /bridgeLimits. Please try again later.',
      details: err.message,
    })
  }
}
