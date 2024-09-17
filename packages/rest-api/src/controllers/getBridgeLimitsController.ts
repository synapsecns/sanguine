import axios from 'axios'
import { validationResult } from 'express-validator'
import { BigNumber } from 'ethers'
import { parseUnits } from '@ethersproject/units'
import { getAddress } from '@ethersproject/address'

import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'

export const getBridgeLimitsController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const { fromChain, fromToken, toChain, toToken } = req.query

    const fromTokenInfo = tokenAddressToToken(fromChain, fromToken)
    const toTokenInfo = tokenAddressToToken(toChain, toToken)

    const rfqResponse = await axios.get('https://rfq-api.omnirpc.io/quotes', {
      params: {
        originChainId: fromChain,
        originTokenAddress: fromTokenInfo.address,
        destChainId: toChain,
        destTokenAddress: toTokenInfo.address,
      },
    })

    const rfqQuotes = rfqResponse.data

    let bestRfqQuote = null

    if (Array.isArray(rfqQuotes) && rfqQuotes.length > 0) {
      const filteredQuotes = rfqQuotes
        .slice(0, 20)
        .filter(
          (quote) =>
            Number(quote.origin_chain_id) === Number(fromChain) &&
            Number(quote.dest_chain_id) === Number(toChain) &&
            getAddress(quote.origin_token_addr) ===
              getAddress(fromTokenInfo.address) &&
            getAddress(quote.dest_token_addr) ===
              getAddress(toTokenInfo.address)
        )

      bestRfqQuote = filteredQuotes.reduce((maxQuote, currentQuote) => {
        const currentAmount = Number(currentQuote.max_origin_amount)
        const maxAmount = maxQuote ? Number(maxQuote.max_origin_amount) : 0
        return currentAmount > maxAmount ? currentQuote : maxQuote
      }, null)
    }

    const upperLimitAmount = parseUnits('1000000', fromTokenInfo.decimals)
    const upperLimitBridgeQuotes = await Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromTokenInfo.address,
      toTokenInfo.address,
      upperLimitAmount
    )

    const lowerLimitAmount = parseUnits('100', fromTokenInfo.decimals)
    const lowerLimitBridgeQuotes = await Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromTokenInfo.address,
      toTokenInfo.address,
      lowerLimitAmount
    )

    const bestUpperLimitSDKQuote = upperLimitBridgeQuotes[0]

    let maxOriginQuote

    const minBridgeFeeQuote = lowerLimitBridgeQuotes.reduce(
      (minQuote, currentQuote) => {
        const currentFeeAmount = currentQuote.feeAmount
        const minFeeAmount = minQuote ? minQuote.feeAmount : null

        return !minFeeAmount || currentFeeAmount.lt(minFeeAmount)
          ? currentQuote
          : minQuote
      },
      null
    )

    if (bestRfqQuote) {
      const bestRfqQuoteMaxAmountBN = BigNumber.from(
        bestRfqQuote.max_origin_amount
      )
      maxOriginQuote = bestRfqQuoteMaxAmountBN.gt(
        bestUpperLimitSDKQuote.maxAmountOut
      )
        ? { source: 'RFQ', amount: bestRfqQuoteMaxAmountBN }
        : {
            source: bestUpperLimitSDKQuote.bridgeModuleName,
            amount: bestUpperLimitSDKQuote.maxAmountOut,
          }
    } else {
      maxOriginQuote = {
        source: bestUpperLimitSDKQuote.bridgeModuleName,
        amount: bestUpperLimitSDKQuote.maxAmountOut,
      }
    }

    return res.json({
      maxOriginAmount: maxOriginQuote.amount,
      minOriginAmount: minBridgeFeeQuote.feeAmount,
    })
  } catch (err) {
    res.status(500).json({
      error:
        'An unexpected error occurred in /getBridgeLimits. Please try again later.',
      details: err.message,
    })
  }
}
