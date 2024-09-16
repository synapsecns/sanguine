import express from 'express'
import axios from 'axios'
import { isString } from 'lodash'
import { BigNumber } from 'ethers'
import { parseUnits } from '@ethersproject/units'
import { getAddress } from '@ethersproject/address'

import { Synapse } from '../services/synapseService'
import { findTokenInfo } from '../utils/findTokenInfo'

const router = express.Router()

router.get('/get-quote-max', async (req, res) => {
  try {
    const {
      originChainId,
      originTokenAddress,
      originTokenSymbol,
      destChainId,
      destTokenAddress,
      destTokenSymbol,
    } = req.query

    if (
      !originChainId ||
      !originTokenSymbol ||
      !originTokenAddress ||
      !destChainId ||
      !destTokenSymbol ||
      !destTokenAddress
    ) {
      return res.status(400).json({ error: 'Missing required parameters' })
    }

    if (
      !isString(originChainId) ||
      !isString(originTokenAddress) ||
      !isString(destChainId) ||
      !isString(originTokenSymbol) ||
      !isString(destTokenSymbol) ||
      !isString(destTokenAddress)
    ) {
      return res.status(400).json({ error: 'Invalid parameters type' })
    }

    const originTokenInfo = findTokenInfo(originChainId, originTokenSymbol)
    const destTokenInfo = findTokenInfo(destChainId, destTokenSymbol)

    if (!originTokenInfo || !destTokenInfo) {
      return res
        .status(404)
        .json({ error: 'Invalid token symbols or chainids' })
    }

    const { decimals: originTokenDecimals } = originTokenInfo

    const rfqResponse = await axios.get('https://rfq-api.omnirpc.io/quotes', {
      params: {
        originChainId,
        originTokenAddress,
        destChainId,
        destTokenAddress,
      },
    })

    const rfqQuotes = rfqResponse.data

    let bestRfqQuote = null

    if (Array.isArray(rfqQuotes) && rfqQuotes.length > 0) {
      const filteredQuotes = rfqQuotes
        .slice(0, 20)
        .filter(
          (quote) =>
            Number(quote.origin_chain_id) === Number(originChainId) &&
            getAddress(quote.origin_token_addr) === originTokenAddress &&
            Number(quote.dest_chain_id) === Number(destChainId) &&
            getAddress(quote.dest_token_addr) === destTokenAddress
        )

      bestRfqQuote = filteredQuotes.reduce((maxQuote, currentQuote) => {
        const currentAmount = Number(currentQuote.max_origin_amount)
        const maxAmount = maxQuote ? Number(maxQuote.max_origin_amount) : 0
        return currentAmount > maxAmount ? currentQuote : maxQuote
      }, null)
    }

    const amount = parseUnits('1000000', originTokenDecimals)

    const bridgeQuotes = await Synapse.allBridgeQuotes(
      Number(originChainId),
      Number(destChainId),
      originTokenAddress,
      destTokenAddress,
      amount
    )

    if (!Array.isArray(bridgeQuotes) || bridgeQuotes.length === 0) {
      return res.status(404).json({ error: 'No bridge quotes found' })
    }

    const bestSDKQuote = bridgeQuotes[0]

    let maxOriginQuote
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
    })
  } catch (error) {
    console.error('Error fetching quotes:', error)
    return res.status(500).json({ error: 'Failed to fetch quotes' })
  }
})
