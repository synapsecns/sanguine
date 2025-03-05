import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { swapV2Controller } from '../controllers/swapV2Controller'
import { CHAINS_ARRAY } from '../constants/chains'
import { isTokenAddress } from '../utils/isTokenAddress'
import { isTokenSupportedOnChain } from '../utils/isTokenSupportedOnChain'
import { checksumAddresses } from '../middleware/checksumAddresses'
import { normalizeNativeTokenAddress } from '../middleware/normalizeNativeTokenAddress'
import { validSwapChain } from '../validations/validSwapChain'
import { validateDecimals } from '../validations/validateDecimals'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'

const router: express.Router = express.Router()

// TODO: openapi, swagger
// TODO: this should support arbitrary to/from tokens
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
    check('amount')
      .exists()
      .withMessage('amount is required')
      .isNumeric()
      .custom((value, { req }) => {
        const fromTokenInfo = tokenAddressToToken(
          req.query.chain,
          req.query.fromToken
        )
        return validateDecimals(value, fromTokenInfo.decimals)
      })
      .withMessage(
        'Amount has too many decimals, beyond the maximum allowed for this token'
      ),
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
