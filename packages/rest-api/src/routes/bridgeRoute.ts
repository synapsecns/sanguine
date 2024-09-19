import express from 'express'
import { check } from 'express-validator'

import { isTokenAddress } from '../utils/isTokenAddress'
import { CHAINS_ARRAY } from '../constants/chains'
import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { bridgeController } from '../controllers/bridgeController'
import { isTokenSupportedOnChain } from '../utils/isTokenSupportedOnChain'
import { checksumAddresses } from '../middleware/checksumAddresses'

const router = express.Router()

router.get(
  '/',
  checksumAddresses(['fromToken', 'toToken']),
  [
    check('fromChain')
      .exists()
      .withMessage('fromChain is required')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported fromChain'),
    check('toChain')
      .exists()
      .withMessage('toChain is required')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported toChain'),
    check('fromToken')
      .exists()
      .withMessage('fromToken is required')
      .custom((value) => isTokenAddress(value))
      .withMessage('Invalid fromToken address')
      .custom((value, { req }) =>
        isTokenSupportedOnChain(value, req.query.fromChain as string)
      )
      .withMessage('Token not supported on specified chain'),
    check('toToken')
      .exists()
      .withMessage('toToken is required')
      .custom((value) => isTokenAddress(value))
      .withMessage('Invalid toToken address')
      .custom((value, { req }) =>
        isTokenSupportedOnChain(value, req.query.toChain as string)
      )
      .withMessage('Token not supported on specified chain'),
    check('amount').isNumeric().exists().withMessage('amount is required'),
  ],
  showFirstValidationError,
  bridgeController
)

export default router
