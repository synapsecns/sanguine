import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

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
    check('amount').isNumeric().exists().withMessage('amount is required'),
  ],
  showFirstValidationError,
  swapController
)

export default router
