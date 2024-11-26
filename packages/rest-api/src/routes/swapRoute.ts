import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { swapController } from '../controllers/swapController'
import { CHAINS_ARRAY } from '../constants/chains'
import { isTokenAddress } from '../utils/isTokenAddress'
import { isTokenSupportedOnChain } from '../utils/isTokenSupportedOnChain'
import { checksumAddresses } from '../middleware/checksumAddresses'
import { normalizeNativeTokenAddress } from '../middleware/normalizeNativeTokenAddress'
import { validSwapTokens } from '../validations/validSwapTokens'
import { validSwapChain } from '../validations/validSwapChain'
import { validateDecimals } from '../validations/validateDecimals'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'

const router: express.Router = express.Router()

/**
 * @openapi
 * /swap:
 *   get:
 *     summary: Get swap quote for tokens on a specific chain
 *     description: Retrieve detailed swap quote for exchanging one token for another on a specified chain
 *     parameters:
 *       - in: query
 *         name: chain
 *         required: true
 *         schema:
 *           type: integer
 *         description: The chain ID where the swap will occur
 *       - in: query
 *         name: fromToken
 *         required: true
 *         schema:
 *           type: string
 *         description: The address of the token to swap from
 *       - in: query
 *         name: toToken
 *         required: true
 *         schema:
 *           type: string
 *         description: The address of the token to swap to
 *       - in: query
 *         name: amount
 *         required: true
 *         schema:
 *           type: number
 *         description: The amount of tokens to swap
 *       - in: query
 *         name: address
 *         required: false
 *         schema:
 *           type: string
 *         description: Optional. The address that will perform the swap. If provided, returns transaction data.
 *     responses:
 *       200:
 *         description: Successful response
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 routerAddress:
 *                   type: string
 *                   description: The address of the router contract
 *                 maxAmountOut:
 *                   type: string
 *                   description: The maximum amount of tokens that will be received
 *                 query:
 *                   type: object
 *                   properties:
 *                     swapAdapter:
 *                       type: string
 *                       description: The address of the swap adapter
 *                     tokenOut:
 *                       type: string
 *                       description: The address of the token being received
 *                     minAmountOut:
 *                       $ref: '#/components/schemas/BigNumber'
 *                     deadline:
 *                       $ref: '#/components/schemas/BigNumber'
 *                     rawParams:
 *                       type: string
 *                       description: Raw parameters for the swap
 *                     callData:
 *                       type: object
 *                       nullable: true
 *                       properties:
 *                         to:
 *                           type: string
 *                         data:
 *                           type: string
 *                         value:
 *                           type: string
 *             example:
 *               routerAddress: "0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a"
 *               maxAmountOut: "999.746386"
 *               query:
 *                 swapAdapter: "0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a"
 *                 tokenOut: "0xdAC17F958D2ee523a2206206994597C13D831ec7"
 *                 minAmountOut:
 *                   type: "BigNumber"
 *                   hex: "0x3b96eb52"
 *                 deadline:
 *                   type: "BigNumber"
 *                   hex: "0x66ecb470"
 *                 rawParams: "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000001116898dda4015ed8ddefb84b6e8bc24528af2d800000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"
 *       400:
 *         description: Invalid input
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 error:
 *                   type: object
 *                   properties:
 *                     value:
 *                       type: string
 *                     message:
 *                       type: string
 *                     field:
 *                       type: string
 *                     location:
 *                       type: string
 *             example:
 *               error:
 *                 value: "999"
 *                 message: "Unsupported chain"
 *                 field: "chain"
 *                 location: "query"
 *       500:
 *         description: Server error
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 error:
 *                   type: string
 *
 * components:
 *   schemas:
 *     BigNumber:
 *       type: object
 *       properties:
 *         type:
 *           type: string
 *           enum: [BigNumber]
 *         hex:
 *           type: string
 */
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
    check()
      .custom((_value, { req }) => {
        const { chain, fromToken, toToken } = req.query

        return validSwapTokens(chain, fromToken, toToken)
      })
      .withMessage('Swap not supported for given tokens'),
    check('address')
      .optional()
      .custom((value) => isAddress(value))
      .withMessage('Invalid address'),
  ],
  showFirstValidationError,
  swapController
)

export default router
