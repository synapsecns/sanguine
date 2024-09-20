import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

import { CHAINS_ARRAY } from '../constants/chains'
import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { swapTxInfoController } from '../controllers/swapTxInfoController'
import { isTokenAddress } from '../utils/isTokenAddress'
import { isTokenSupportedOnChain } from '../utils/isTokenSupportedOnChain'
import { checksumAddresses } from '../middleware/checksumAddresses'

const router = express.Router()

router.get(
  '/',
  checksumAddresses(['fromToken', 'toToken']),
  [
    check('chain')
      .exists()
      .withMessage('chain is required')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported chain'),
    check('fromToken')
      .exists()
      .withMessage('fromToken is required')
      .custom((value) => isTokenAddress(value))
      .withMessage('Invalid fromToken address')
      .custom((value, { req }) =>
        isTokenSupportedOnChain(value, req.query.chain as string)
      )
      .withMessage('Token not supported on specified chain'),
    check('toToken')
      .exists()
      .withMessage('toToken is required')
      .custom((value) => isTokenAddress(value))
      .withMessage('Invalid toToken address')
      .custom((value, { req }) =>
        isTokenSupportedOnChain(value, req.query.chain as string)
      )
      .withMessage('Token not supported on specified chain'),
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
