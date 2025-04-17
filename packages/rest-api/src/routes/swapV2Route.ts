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
 *         name: chainId
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
 *         name: fromAmount
 *         required: true
 *         schema:
 *           type: string
 *         description: The amount of tokens to swap in the token's native decimals
 *       - in: query
 *         name: toToken
 *         required: true
 *         schema:
 *           type: string
 *         description: The address of the token to swap to
 *       - in: query
 *         name: toRecipient
 *         required: false
 *         schema:
 *           type: string
 *         description: Optional. The address that will receive the swapped tokens. If provided, returns transaction data.
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
 *                 id:
 *                   type: string
 *                   description: Unique identifier for the quote
 *                 chainId:
 *                   type: integer
 *                   description: The chain ID where the swap occurs
 *                 fromToken:
 *                   type: string
 *                   description: The address of the token being swapped from
 *                 fromAmount:
 *                   type: string
 *                   description: The amount of tokens to swap in the token's native decimals
 *                 toToken:
 *                   type: string
 *                   description: The address of the token being swapped to
 *                 expectedToAmount:
 *                   type: string
 *                   description: The expected amount of tokens that will be received in the token's native decimals
 *                 minToAmount:
 *                   type: string
 *                   description: The minimum amount of tokens that will be received in the token's native decimals (includes slippage)
 *                 routerAddress:
 *                   type: string
 *                   description: The address of the router contract
 *                 moduleNames:
 *                   type: array
 *                   items:
 *                     type: string
 *                   description: The names of the swap modules used for this quote
 *                 callData:
 *                   type: object
 *                   nullable: true
 *                   description: Transaction data object, only provided if toRecipient parameter was included
 *                   properties:
 *                     to:
 *                       type: string
 *                       description: Contract address to call
 *                     data:
 *                       type: string
 *                       description: Transaction calldata
 *                     value:
 *                       type: string
 *                       description: Amount of native currency to send with transaction (in native token decimals)
 *             example:
 *               id: "01920c87-7f14-7cdf-90e1-e13b2d4af55f"
 *               chainId: 1
 *               fromToken: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
 *               fromAmount: "1000000"
 *               toToken: "0xdac17f958d2ee523a2206206994597c13d831ec7"
 *               expectedToAmount: "999746386"
 *               minToAmount: "994747654"
 *               routerAddress: "0x512000a034E154908Efb1eC48579F4ffDb000512"
 *               moduleNames: ["DefaultPools"]
 *               callData: {
 *                 to: "0x512000a034E154908Efb1eC48579F4ffDb000512",
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
 *                   param: "chainId",
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
    check('chainId')
      .exists()
      .withMessage('chainId is required')
      .isInt()
      .withMessage('chainId must be an integer')
      .toInt()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === value))
      .withMessage('Unsupported chain')
      .custom((value) => INTENTS_SUPPORTED_CHAIN_IDS.includes(value))
      .withMessage('Swap not supported for given chain'),
    check('fromToken')
      .exists()
      .withMessage('fromToken is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid fromToken address'),
    // Don't convert fromAmount to int as it is a BigNumber
    check('fromAmount').exists().withMessage('fromAmount is required').isInt(),
    check('toToken')
      .exists()
      .withMessage('toToken is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid toToken address'),
    check('toRecipient')
      .optional()
      .custom((value) => isAddress(value))
      .withMessage('Invalid toRecipient address'),
    check('slippage')
      .optional()
      .isNumeric()
      .withMessage('slippage must be a number')
      .toFloat()
      .custom((value) => value >= 0 && value <= 100)
      .withMessage('Slippage must be between 0 and 100'),
  ],
  showFirstValidationError,
  swapV2Controller
)

export default router
