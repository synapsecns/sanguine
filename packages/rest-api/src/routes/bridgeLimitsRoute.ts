import express from 'express'
import { check } from 'express-validator'
import { isAddress } from '@ethersproject/address'

import { CHAINS_ARRAY } from '../constants/chains'
import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { getBridgeLimitsController } from '../controllers/bridgeLimitsController'

const router = express.Router()

router.get(
  '/',
  [
    check('fromChain')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported fromChain')
      .exists()
      .withMessage('originChainId is required'),
    check('toChain')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported toChain')
      .exists()
      .withMessage('toChain is required'),
    check('fromToken')
      .exists()
      .withMessage('fromToken is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid fromToken address'),
    check('toToken')
      .exists()
      .withMessage('toToken is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid toToken address'),
  ],
  showFirstValidationError,
  getBridgeLimitsController
)

export default router
