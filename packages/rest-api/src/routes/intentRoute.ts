import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { intentController } from '../controllers/intentController'
import { CHAINS_ARRAY } from '../constants/chains'
import { INTENTS_SUPPORTED_CHAIN_IDS } from '../constants'
import { checksumAddresses } from '../middleware/checksumAddresses'
import { normalizeNativeTokenAddress } from '../middleware/normalizeNativeTokenAddress'

const router: express.Router = express.Router()

// TODO: openapi, swagger
router.get(
  '/',
  normalizeNativeTokenAddress(['fromToken', 'toToken']),
  checksumAddresses(['fromToken', 'toToken']),
  [
    check('fromChainId')
      .exists()
      .withMessage('fromChainId is required')
      .isInt()
      .withMessage('fromChain must be an integer')
      .toInt()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === value))
      .withMessage('Unsupported fromChainId')
      .custom((value) => INTENTS_SUPPORTED_CHAIN_IDS.includes(value))
      .withMessage('Intents not supported for given chain'),
    check('fromToken')
      .exists()
      .withMessage('fromToken is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid fromToken address'),
    // Don't convert fromAmount to int as it is a BigNumber
    check('fromAmount').exists().withMessage('fromAmount is required').isInt(),
    check('fromSender')
      .optional()
      .custom((value) => isAddress(value))
      .withMessage('Invalid fromSender address'),
    check('toChainId')
      .exists()
      .withMessage('toChainId is required')
      .isInt()
      .withMessage('toChainId must be an integer')
      .toInt()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === value))
      .withMessage('Unsupported toChainId'),
    check('toToken')
      .exists()
      .withMessage('toToken is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid toToken address'),
    check('toRecipient')
      .optional()
      .custom((value) => isAddress(value))
      .withMessage('Invalid toRecipient address'),
    check('slippage')
      .optional()
      .isNumeric()
      .withMessage('slippage must be a number')
      .toFloat()
      .custom((value) => value >= 0 && value <= 100)
      .withMessage('Slippage must be between 0 and 100'),
    check('allowMultipleTxs')
      .optional()
      .isBoolean()
      .withMessage('allowMultipleTxs must be a boolean')
      .toBoolean(),
  ],
  showFirstValidationError,
  intentController
)

export default router
