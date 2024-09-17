import { validationResult } from 'express-validator'
import axios from 'axios'
import { BigNumber } from 'ethers'
import { parseUnits } from '@ethersproject/units'
import { getAddress } from '@ethersproject/address'

import { Synapse } from '../services/synapseService'

export const getBridgeLimitsController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const { fromChain, toChain } = req.query

    const fromTokenInfo = res.locals.tokenInfo.fromToken
    const toTokenInfo = res.locals.tokenInfo.toToken

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
            getAddress(quote.origin_token_addr) ===
              getAddress(fromTokenInfo.address) &&
            Number(quote.dest_chain_id) === Number(toChain) &&
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

    if (
      !Array.isArray(upperLimitBridgeQuotes) ||
      upperLimitBridgeQuotes.length === 0
    ) {
      return res.status(404).json({ error: 'No bridge quotes found' })
    }

    const lowerLimitAmount = parseUnits('100', fromTokenInfo.decimals)
    const lowerLimitBridgeQuotes = await Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromTokenInfo.address,
      toTokenInfo.address,
      lowerLimitAmount
    )

    const bestSDKQuote = upperLimitBridgeQuotes[0]

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
      maxOriginQuote = bestRfqQuoteMaxAmountBN.gt(bestSDKQuote.maxAmountOut)
        ? { source: 'RFQ', amount: bestRfqQuoteMaxAmountBN }
        : {
            source: bestSDKQuote.bridgeModuleName,
            amount: bestSDKQuote.maxAmountOut,
          }
    } else {
      // If no RFQ quote, simply use the SDK quote
      maxOriginQuote = {
        source: bestSDKQuote.bridgeModuleName,
        amount: bestSDKQuote.maxAmountOut,
      }
    }

    return res.json({
      rfqBestQuote: bestRfqQuote,
      sdkBestQuote: bestSDKQuote,
      maxOriginQuote,
      minBridgeFeeQuote,
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
