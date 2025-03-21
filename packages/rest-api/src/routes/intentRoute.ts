import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { intentController } from '../controllers/intentController'
import { CHAINS_ARRAY } from '../constants/chains'
import { INTENTS_SUPPORTED_CHAIN_IDS } from '../constants'
import { checksumAddresses } from '../middleware/checksumAddresses'
import { normalizeNativeTokenAddress } from '../middleware/normalizeNativeTokenAddress'

const router: express.Router = express.Router()

/**
 * @openapi
 * /intent:
 *   get:
 *     summary: Get intent quotes for transferring tokens
 *     description: Retrieve list of intent quotes for transferring tokens between chains or within a chain. Intents can involve a combination of swaps and bridges to achieve the user's goal of moving from one token to another.
 *     parameters:
 *       - in: query
 *         name: fromChainId
 *         required: true
 *         schema:
 *           type: integer
 *         description: The origin chain ID (must support intents)
 *       - in: query
 *         name: fromToken
 *         required: true
 *         schema:
 *           type: string
 *         description: The address of the token on the origin chain
 *       - in: query
 *         name: fromAmount
 *         required: true
 *         schema:
 *           type: integer
 *         description: The amount of tokens to transfer in the token's native decimals
 *       - in: query
 *         name: fromSender
 *         required: false
 *         schema:
 *           type: string
 *         description: The address of the sender on the origin chain (required to generate callData)
 *       - in: query
 *         name: toChainId
 *         required: true
 *         schema:
 *           type: integer
 *         description: The destination chain ID
 *       - in: query
 *         name: toToken
 *         required: true
 *         schema:
 *           type: string
 *         description: The address of the token on the destination chain
 *       - in: query
 *         name: toRecipient
 *         required: false
 *         schema:
 *           type: string
 *         description: The recipient address on the destination chain (required to generate callData)
 *       - in: query
 *         name: slippage
 *         required: false
 *         schema:
 *           type: number
 *         description: Optional. The maximum allowed slippage percentage (0-100). Defaults to 0.5%.
 *       - in: query
 *         name: allowMultipleTxs
 *         required: false
 *         schema:
 *           type: boolean
 *         description: Optional. Whether to allow intent execution across multiple transactions.
 *     responses:
 *       200:
 *         description: Successful response
 *         content:
 *           application/json:
 *             schema:
 *               type: array
 *               items:
 *                 type: object
 *                 properties:
 *                   id:
 *                     type: string
 *                     description: Unique identifier for the quote (UUIDv7)
 *                   fromChainId:
 *                     type: integer
 *                     description: The ID of the origin chain
 *                   fromToken:
 *                     type: string
 *                     description: The address of the token on the origin chain
 *                   fromAmount:
 *                     type: string
 *                     description: The amount of tokens to transfer (in native token decimals)
 *                   toChainId:
 *                     type: integer
 *                     description: The ID of the destination chain
 *                   toToken:
 *                     type: string
 *                     description: The address of the token on the destination chain
 *                   expectedToAmount:
 *                     type: string
 *                     description: The expected amount of tokens to be received (in native token decimals)
 *                   minToAmount:
 *                     type: string
 *                     description: The minimum amount of tokens that will be received (in native token decimals)
 *                   estimatedTime:
 *                     type: integer
 *                     description: Estimated time for the intent execution in seconds
 *                   steps:
 *                     type: array
 *                     description: The list of steps to execute the intent
 *                     items:
 *                       type: object
 *                       properties:
 *                         fromChainId:
 *                           type: integer
 *                           description: The ID of the origin chain for this step
 *                         fromToken:
 *                           type: string
 *                           description: The address of the token on the origin chain for this step
 *                         fromAmount:
 *                           type: string
 *                           description: The amount of tokens for this step (in native token decimals)
 *                         toChainId:
 *                           type: integer
 *                           description: The ID of the destination chain for this step
 *                         toToken:
 *                           type: string
 *                           description: The address of the token on the destination chain for this step
 *                         expectedToAmount:
 *                           type: string
 *                           description: The expected amount of tokens for this step (in native token decimals)
 *                         minToAmount:
 *                           type: string
 *                           description: The minimum amount of tokens for this step (in native token decimals)
 *                         routerAddress:
 *                           type: string
 *                           description: The address of the router contract for this step
 *                         estimatedTime:
 *                           type: integer
 *                           description: Estimated time for this step in seconds
 *                         moduleNames:
 *                           type: array
 *                           items:
 *                             type: string
 *                           description: The names of the bridge or swap modules used for this step
 *                         gasDropAmount:
 *                           type: string
 *                           description: Amount of native token airdropped on destination chain (in native token decimals)
 *                         callData:
 *                           type: object
 *                           nullable: true
 *                           description: Transaction data object, only provided if addresses were included
 *                           properties:
 *                             to:
 *                               type: string
 *                               description: Contract address to call
 *                             data:
 *                               type: string
 *                               description: Transaction calldata
 *                             value:
 *                               type: string
 *                               description: Amount of native currency to send with transaction (in native token decimals)
 *             example:
 *               - id: "01920c87-7f14-7cdf-90e1-e13b2d4af55f"
 *                 fromChainId: 1
 *                 fromToken: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
 *                 fromAmount: "1000000000"
 *                 toChainId: 42161
 *                 toToken: "0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8"
 *                 expectedToAmount: "999046695719"
 *                 minToAmount: "994046695719"
 *                 estimatedTime: 30
 *                 steps: [
 *                   {
 *                     fromChainId: 1
 *                     fromToken: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
 *                     fromAmount: "1000000000"
 *                     toChainId: 42161
 *                     toToken: "0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8"
 *                     expectedToAmount: "999046695719"
 *                     minToAmount: "994046695719"
 *                     routerAddress: "0x512000a034E154908Efb1eC48579F4ffDb000512"
 *                     estimatedTime: 30
 *                     moduleNames: ["SynapseRFQ"]
 *                     gasDropAmount: "0"
 *                     callData: {
 *                       to: "0x512000a034E154908Efb1eC48579F4ffDb000512",
 *                       data: "0x...",
 *                       value: "0"
 *                     }
 *                   }
 *                 ]
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
 *                   msg: "Unsupported fromChain",
 *                   param: "fromChain",
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
 *               error: "An unexpected error occurred in /intent. Please try again later."
 */
router.get(
  '/',
  normalizeNativeTokenAddress(['fromToken', 'toToken']),
  checksumAddresses(['fromToken', 'toToken']),
  [
    check('fromChainId')
      .exists()
      .withMessage('fromChainId is required')
      .isInt()
      .withMessage('fromChain must be an integer')
      .toInt()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === value))
      .withMessage('Unsupported fromChainId')
      .custom((value) => INTENTS_SUPPORTED_CHAIN_IDS.includes(value))
      .withMessage('Intents not supported for given chain'),
    check('fromToken')
      .exists()
      .withMessage('fromToken is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid fromToken address'),
    // Don't convert fromAmount to int as it is a BigNumber
    check('fromAmount').exists().withMessage('fromAmount is required').isInt(),
    check('fromSender')
      .optional()
      .custom((value) => isAddress(value))
      .withMessage('Invalid fromSender address'),
    check('toChainId')
      .exists()
      .withMessage('toChainId is required')
      .isInt()
      .withMessage('toChainId must be an integer')
      .toInt()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === value))
      .withMessage('Unsupported toChainId'),
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
    check('allowMultipleTxs')
      .optional()
      .isBoolean()
      .withMessage('allowMultipleTxs must be a boolean')
      .toBoolean(),
  ],
  showFirstValidationError,
  intentController
)

export default router
