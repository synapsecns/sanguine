import { validationResult } from 'express-validator'

import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'
import { formatAndValidateAmount } from '../utils/formatAmounts'
import { logger } from '../middleware/logger'

export const bridgeTxInfoController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }

  try {
    const {
      fromChain,
      toChain,
      amount,
      destAddress,
      fromToken,
      toToken,
      originUserAddress,
    } = req.query

    const fromTokenInfo = tokenAddressToToken(fromChain.toString(), fromToken)

    const amountInWei = formatAndValidateAmount(
      amount.toString(),
      fromTokenInfo.decimals
    )

    const quotes = await Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromToken,
      toToken,
      amountInWei,
      originUserAddress
        ? { originUserAddress: originUserAddress.toString() }
        : {}
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

    logger.info(`Successful bridgeTxInfoController response`, {
      query: req.query,
      txInfoArray,
    })
    res.json(txInfoArray)
  } catch (err) {
    logger.error(`Error in bridgeTxInfoController`, {
      query: req.query,
      error: err.message,
      stack: err.stack,
    })
    res.status(500).json({
      error:
        'An unexpected error occurred in /bridgeTxInfo. Please try again later.',
    })
  }
}
