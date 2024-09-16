import express from 'express'
import { check } from 'express-validator'

import { CHAINS_ARRAY } from '../constants/chains'
import { validateTokens } from '../validations/validateTokens'
import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { getBridgeLimitsController } from '../controllers/getBridgeLimitsController'

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
    validateTokens('fromChain', 'fromToken', 'fromToken'),
    validateTokens('toChain', 'toToken', 'toToken'),
  ],
  showFirstValidationError,
  getBridgeLimitsController
)

export default router
