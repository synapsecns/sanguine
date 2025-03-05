import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { swapV2Controller } from '../controllers/swapV2Controller'
import { CHAINS_ARRAY } from '../constants/chains'
import { checksumAddresses } from '../middleware/checksumAddresses'
import { normalizeNativeTokenAddress } from '../middleware/normalizeNativeTokenAddress'
import { validSwapChain } from '../validations/validSwapChain'

const router: express.Router = express.Router()

// TODO: openapi, swagger
router.get(
  '/',
  normalizeNativeTokenAddress(['fromToken', 'toToken']),
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
      .custom((value) => isAddress(value))
      .withMessage('Invalid fromToken address'),
    check('toToken')
      .exists()
      .withMessage('toToken is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid toToken address'),
    check('amount').exists().withMessage('amount is required').isInt(),
    check()
      .custom((_value, { req }) => {
        const { chain } = req.query

        return validSwapChain(chain)
      })
      .withMessage('Swap not supported for given chain'),
    check('address')
      .optional()
      .custom((value) => isAddress(value))
      .withMessage('Invalid address'),
    check('slippage')
      .optional()
      .isNumeric()
      .custom((value) => value >= 0 && value <= 100)
      .withMessage('Slippage must be between 0 and 100'),
  ],
  showFirstValidationError,
  swapV2Controller
)

export default router
