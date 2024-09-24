import { validationResult } from 'express-validator'

import { BRIDGE_LIMITS_MAP } from '../constants/bridgeLimitsMap'

export const bridgeLimitsController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }
  try {
    const { fromChain, fromToken, toChain, toToken } = req.query

    let maxOriginAmount = null
    let minOriginAmount = null

    if (BRIDGE_LIMITS_MAP[fromChain][fromToken]) {
      minOriginAmount =
        BRIDGE_LIMITS_MAP[fromChain][fromToken].routes[toChain][toToken]
          .minOriginValue
      maxOriginAmount =
        BRIDGE_LIMITS_MAP[fromChain][fromToken].routes[toChain][toToken]
          .maxOriginValue
    }

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
