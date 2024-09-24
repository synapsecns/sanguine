import { validationResult } from 'express-validator'
import { BigNumber } from 'ethers'
import { parseUnits } from '@ethersproject/units'

import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'
import { formatBNToString } from '../utils/formatBNToString'
import BRIDGE_LIMITS_MAP from '../constants/limitsMap.ts'

export const bridgeLimitsController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const { fromChain, fromToken, toChain, toToken } = req.query

    const fromTokenInfo = tokenAddressToToken(fromChain, fromToken)
    const toTokenInfo = tokenAddressToToken(toChain, toToken)

    let maxOriginAmount = null
    let minOriginAmount = null

    if (
      BRIDGE_LIMITS_MAP[fromChain][fromTokenInfo.address].routes[toChain][
        toTokenInfo.address
      ]
    ) {
      minOriginAmount =
        BRIDGE_LIMITS_MAP[fromChain][fromTokenInfo.address].routes[toChain][
          toTokenInfo.address
        ].minOriginValue
      maxOriginAmount =
        BRIDGE_LIMITS_MAP[fromChain][fromTokenInfo.address].routes[toChain][
          toTokenInfo.address
        ].maxOriginValue
    }

    // const upperLimitValue = parseUnits('1000000', fromTokenInfo.decimals)
    // const upperLimitBridgeQuotes = await Synapse.allBridgeQuotes(
    //   Number(fromChain),
    //   Number(toChain),
    //   fromTokenInfo.address,
    //   toTokenInfo.address,
    //   upperLimitValue
    // )

    // const lowerLimitValues = ['0.01', '10']
    // let lowerLimitBridgeQuotes = null

    // for (const limit of lowerLimitValues) {
    //   const lowerLimitAmount = parseUnits(limit, fromTokenInfo.decimals)

    //   lowerLimitBridgeQuotes = await Synapse.allBridgeQuotes(
    //     Number(fromChain),
    //     Number(toChain),
    //     fromTokenInfo.address,
    //     toTokenInfo.address,
    //     lowerLimitAmount
    //   )

    //   if (lowerLimitBridgeQuotes && lowerLimitBridgeQuotes.length > 0) {
    //     break
    //   }
    // }

    // const maxBridgeAmountQuote = upperLimitBridgeQuotes.reduce(
    //   (maxQuote, currentQuote) => {
    //     const currentMaxAmount = currentQuote.maxAmountOut
    //     const maxAmount = maxQuote ? maxQuote.maxAmountOut : BigNumber.from(0)

    //     return currentMaxAmount.gt(maxAmount) ? currentQuote : maxQuote
    //   },
    //   null
    // )

    // const minBridgeAmountQuote = lowerLimitBridgeQuotes.reduce(
    //   (minQuote, currentQuote) => {
    //     const currentFeeAmount = currentQuote.feeAmount
    //     const minFeeAmount = minQuote ? minQuote.feeAmount : null

    //     return !minFeeAmount || currentFeeAmount.lt(minFeeAmount)
    //       ? currentQuote
    //       : minQuote
    //   },
    //   null
    // )

    // if (!maxBridgeAmountQuote || !minBridgeAmountQuote) {
    //   return res.json({
    //     maxOriginAmount: null,
    //     minOriginAmount: null,
    //   })
    // }

    // const maxAmountOriginQueryTokenOutInfo = tokenAddressToToken(
    //   toChain,
    //   maxBridgeAmountQuote.destQuery.tokenOut
    // )

    // const minAmountOriginQueryTokenOutInfo = tokenAddressToToken(
    //   fromChain,
    //   minBridgeAmountQuote.originQuery.tokenOut
    // )

    // maxOriginAmount = formatBNToString(
    //   maxBridgeAmountQuote.maxAmountOut,
    //   maxAmountOriginQueryTokenOutInfo.decimals
    // )

    // minOriginAmount = formatBNToString(
    //   minBridgeAmountQuote.feeAmount,
    //   minAmountOriginQueryTokenOutInfo.decimals
    // )

    return res.json({
      maxOriginAmount: processAmount(maxOriginAmount),
      minOriginAmount: processAmount(minOriginAmount),
    })
  } catch (err) {
    res.status(500).json({
      error:
        'An unexpected error occurred in /bridgeLimits. Please try again later.',
      details: err.message,
    })
  }
}

const processAmount = (amount: string): number | null => {
  const value = parseFloat(amount)

  if (isNaN(value)) {
    return null
  }

  if (value < 0.01) {
    return value // If the value is less than 0.01, keep the same value
  } else if (value >= 0.01 && value < 0.1) {
    return 0.1 // If the value is between 0.01 and 0.1, use 0.1
  } else if (value >= 0.1 && value < 0.5) {
    return 0.5 // If the value is between 0.1 and 0.5, use 0.5
  } else if (value >= 0.5 && value <= 1900000) {
    return value // If the value is between 0.5 and 1900000, use that value
  } else {
    return null // If the value is above 1900000, set value to null
  }
}
