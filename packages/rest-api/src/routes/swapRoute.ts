import express from 'express'
import { check } from 'express-validator'

import { validateTokens } from '../validations/validateTokens'
import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { swapController } from '../controllers/swapController'
import { CHAINS_ARRAY } from '../constants/chains'

const router = express.Router()

router.get(
  '/',
  [
    check('chain')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported chain')
      .exists()
      .withMessage('chain is required'),
    validateTokens('chain', 'fromToken', 'fromToken'),
    validateTokens('chain', 'toToken', 'toToken'),
    check('amount').isNumeric().exists().withMessage('amount is required'),
  ],
  showFirstValidationError,
  swapController
)

export default router
