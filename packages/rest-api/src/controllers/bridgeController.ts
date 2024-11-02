import { validationResult } from 'express-validator'
import { parseUnits } from '@ethersproject/units'
import { defaultAbiCoder, Interface } from '@ethersproject/abi'

import { formatBNToString } from '../utils/formatBNToString'
import { Synapse } from '../services/synapseService'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'
import { logger } from '../middleware/logger'

const BRIDGE_INTERFACE = new Interface([
  'function bridge(address to, uint256 chainId, address token, uint256 amount, tuple(address routerAdapter, address tokenOut, uint256 minAmountOut, uint256 deadline, bytes rawParams) originQuery, tuple(address routerAdapter, address tokenOut, uint256 minAmountOut, uint256 deadline, bytes rawParams) destQuery) external payable',
])

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
    } = req.query

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

    const payload = resp.map((quote) => {
      const originQueryTokenOutInfo = tokenAddressToToken(
        fromChain.toString(),
        quote.originQuery.tokenOut
      )

      const txData = BRIDGE_INTERFACE.encodeFunctionData('bridge', [
        originUserAddress,
        Number(toChain),
        fromToken,
        amountInWei,
        [
          quote.originQuery.routerAdapter,
          quote.originQuery.tokenOut,
          quote.originQuery.minAmountOut,
          quote.originQuery.deadline,
          quote.originQuery.rawParams
        ],
        [
          quote.destQuery.routerAdapter,
          quote.destQuery.tokenOut,
          quote.destQuery.minAmountOut,
          quote.destQuery.deadline,
          quote.destQuery.rawParams
        ]
      ])

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
        txData,
      }
    })

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
