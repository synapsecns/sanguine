import { validationResult } from 'express-validator'

import { tokenAddressToToken } from '../utils/tokenAddressToToken'
import { BRIDGE_ROUTE_MAPPING } from '../utils/bridgeRouteMapping'

export const destinationTokensController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }

  try {
    const { fromChain, fromToken } = req.query

    const fromTokenInfo = tokenAddressToToken(fromChain.toString(), fromToken)

    const constructedKey = `${fromTokenInfo.symbol}-${fromChain}`

    const options = BRIDGE_ROUTE_MAPPING[constructedKey]

    res.json(options)
  } catch (err) {
    res.status(500).json({
      error:
        'An unexpected error occurred in /destinationTokens. Please try again later.',
      details: err.message,
    })
  }
}
