import { validationResult } from 'express-validator'
import { parseUnits } from '@ethersproject/units'

import { formatBNToString } from '../utils/formatBNToString'
import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'
import { logger } from '../middleware/logger'

export const bridgeController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const {
      fromChain,
      toChain,
      amount,
      fromToken,
      toToken,
      originUserAddress,
      destAddress, // Optional parameter
    } = req.query as {
      fromChain: string
      toChain: string
      amount: string
      fromToken: string
      toToken: string
      originUserAddress?: string
      destAddress?: string
    }

    const fromTokenInfo = tokenAddressToToken(fromChain.toString(), fromToken)
    const toTokenInfo = tokenAddressToToken(toChain.toString(), toToken)

    const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)

    const resp = await Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromToken,
      toToken,
      amountInWei,
      originUserAddress
        ? { originUserAddress: originUserAddress.toString() }
        : {}
    )

    // Check if no bridge quotes were found
    if (!resp || resp.length === 0) {
      logger.info(`No bridge routes found`, {
        query: req.query,
      })
      return res.status(404).json({ error: 'No bridge routes found' })
    }

    const payload = await Promise.all(
      resp.map(async (quote) => {
        const originQueryTokenOutInfo = tokenAddressToToken(
          fromChain.toString(),
          quote.originQuery.tokenOut
        )

        const { originQuery, destQuery } = Synapse.applyBridgeSlippage(
          quote.bridgeModuleName,
          quote.originQuery,
          quote.destQuery
        )

        const callData =
          destAddress && originUserAddress
            ? await Synapse.bridge(
                destAddress,
                quote.routerAddress,
                Number(fromChain),
                Number(toChain),
                fromToken,
                amountInWei,
                originQuery,
                destQuery
              )
            : null

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
          callData,
        }
      })
    )

    logger.info(`Successful bridgeController response`, {
      payload,
      query: req.query,
    })
    res.json(payload)
  } catch (err) {
    logger.error(`Error in bridgeController`, {
      query: req.query,
      error: err.message,
      stack: err.stack,
    })
    res.status(500).json({
      error: 'An unexpected error occurred in /bridge. Please try again later.',
    })
  }
}
