import { validationResult } from 'express-validator'

import { tokenAddressToToken } from '../utils/tokenAddressToToken'
import { BRIDGE_ROUTE_MAPPING } from '../utils/bridgeRouteMapping'
import { logger } from '../middleware/logger'

export const destinationTokensController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }

  try {
    const { fromChain, fromToken } = req.query

    const fromTokenInfo = tokenAddressToToken(fromChain.toString(), fromToken)

    const constructedKey = `${fromTokenInfo.symbol}-${fromChain}`

    const payload = BRIDGE_ROUTE_MAPPING[constructedKey]

    logger.info(`Successful destinationTokensController response`, { payload })

    res.json(payload)
  } catch (err) {
    logger.error(`Error in destinationTokensController`, {
      error: err.message,
      stack: err.stack,
    })
    res.status(500).json({
      error:
        'An unexpected error occurred in /destinationTokens. Please try again later.',
      details: err.message,
    })
  }
}
