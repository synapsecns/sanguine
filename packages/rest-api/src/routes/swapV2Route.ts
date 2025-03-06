import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { swapV2Controller } from '../controllers/swapV2Controller'
import { CHAINS_ARRAY } from '../constants/chains'
import { INTENTS_SUPPORTED_CHAIN_IDS } from '../constants'
import { checksumAddresses } from '../middleware/checksumAddresses'
import { normalizeNativeTokenAddress } from '../middleware/normalizeNativeTokenAddress'

const router: express.Router = express.Router()

/**
 * @openapi
 * /swap/v2:
 *   get:
 *     summary: Get swap quote for arbitrary tokens on a specific chain
 *     description: Retrieve detailed swap quote exchanging any token for any other on a specified chain. Note that amounts are expressed in token's native decimals, not in human-readable format.
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
 *           type: string
 *         description: The amount of tokens to swap in the token's native decimals
 *       - in: query
 *         name: address
 *         required: false
 *         schema:
 *           type: string
 *         description: Optional. The address that will perform the swap and receive its proceeds. If provided, returns transaction data.
 *       - in: query
 *         name: slippage
 *         required: false
 *         schema:
 *           type: number
 *         description: Optional. The maximum allowed slippage percentage (0-100). Defaults to 0.5%.
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
 *                   description: The maximum amount of tokens that will be received in the token's native decimals
 *                 callData:
 *                   type: object
 *                   nullable: true
 *                   description: Transaction data object, only provided if address parameter was included
 *             example:
 *               routerAddress: "0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a"
 *               maxAmountOut: "999746386"
 *               callData: {
 *                 to: "0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a",
 *                 data: "0x...",
 *                 value: "0"
 *               }
 *       400:
 *         description: Invalid input
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 errors:
 *                   type: array
 *                   items:
 *                     type: object
 *                     properties:
 *                       value:
 *                         type: string
 *                       msg:
 *                         type: string
 *                       param:
 *                         type: string
 *                       location:
 *                         type: string
 *             example:
 *               errors: [
 *                 {
 *                   value: "111",
 *                   msg: "Unsupported chain",
 *                   param: "chain",
 *                   location: "query"
 *                 }
 *               ]
 *       500:
 *         description: Server error
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 error:
 *                   type: string
 *             example:
 *               error: "An unexpected error occurred in /swap/v2. Please try again later."
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
      .withMessage('Unsupported chain')
      .custom((value) => INTENTS_SUPPORTED_CHAIN_IDS.includes(Number(value)))
      .withMessage('Swap not supported for given chain'),
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
