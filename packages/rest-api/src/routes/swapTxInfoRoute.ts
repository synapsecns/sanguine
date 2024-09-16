import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

import { CHAINS_ARRAY } from '../constants/chains'
import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { swapTxInfoController } from '../controllers/swapTxInfoController'

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
    check('address')
      .exists()
      .withMessage('address is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid Ethereum address'),
  ],
  showFirstValidationError,
  swapTxInfoController
)

export default router
