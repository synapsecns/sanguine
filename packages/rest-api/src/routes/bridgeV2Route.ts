import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { bridgeV2Controller } from '../controllers/bridgeV2Controller'
import { CHAINS_ARRAY } from '../constants/chains'
import { INTENTS_SUPPORTED_CHAIN_IDS } from '../constants'
import { checksumAddresses } from '../middleware/checksumAddresses'
import { normalizeNativeTokenAddress } from '../middleware/normalizeNativeTokenAddress'
import { isTokenAddress } from '../utils/isTokenAddress'
import { isTokenSupportedOnChain } from '../utils/isTokenSupportedOnChain'

const router: express.Router = express.Router()

/**
 * @openapi
 * /bridge/v2:
 *   get:
 *     summary: Get quotes for bridging tokens between chains
 *     description: Retrieve list of bridge quotes based on origin and destination chains, tokens, and amount. Any origin token can be used, but destination tokens must be supported on the destination chain.
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
 *           type: string
 *         description: The amount of tokens to bridge in the token's native decimals
 *       - in: query
 *         name: fromSender
 *         required: false
 *         schema:
 *           type: string
 *         description: The address of the user on the origin chain (required to generate callData)
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
 *         description: The destination address for the bridge transaction (required to generate callData)
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
 *                     description: The amount of tokens to bridge in the token's native decimals
 *                   toChainId:
 *                     type: integer
 *                     description: The ID of the destination chain
 *                   toToken:
 *                     type: string
 *                     description: The address of the token on the destination chain
 *                   expectedToAmount:
 *                     type: string
 *                     description: The expected amount of tokens that will be received in the token's native decimals
 *                   minToAmount:
 *                     type: string
 *                     description: The minimum amount of tokens that will be received in the token's native decimals (includes slippage)
 *                   routerAddress:
 *                     type: string
 *                     description: The address of the router contract
 *                   estimatedTime:
 *                     type: integer
 *                     description: Estimated time for the bridge in seconds
 *                   moduleName:
 *                     type: string
 *                     description: The name of the bridge module used for this quote
 *                   gasDropAmount:
 *                     type: string
 *                     description: Amount of native token airdropped on destination chain (in native token decimals)
 *                   callData:
 *                     type: object
 *                     nullable: true
 *                     description: Transaction data object, only provided if fromSender and toRecipient parameters were included
 *                     properties:
 *                       to:
 *                         type: string
 *                         description: Contract address to call
 *                       data:
 *                         type: string
 *                         description: Transaction calldata
 *                       value:
 *                         type: string
 *                         description: Amount of native currency to send with transaction (in native token decimals)
 *             example:
 *               - id: "01920c87-7f14-7cdf-90e1-e13b2d4af55f"
 *                 fromChainId: 1
 *                 fromToken: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
 *                 fromAmount: "1000000000000"
 *                 toChainId: 42161
 *                 toToken: "0xff970a61a04b1ca14834a43f5de4533ebddb5cc8"
 *                 expectedToAmount: "999046695719"
 *                 minToAmount: "994051462240"
 *                 routerAddress: "0x512000a034E154908Efb1eC48579F4ffDb000512"
 *                 estimatedTime: 30
 *                 moduleName: "SynapseRFQ"
 *                 gasDropAmount: "0"
 *                 callData: {
 *                   to: "0x512000a034E154908Efb1eC48579F4ffDb000512",
 *                   data: "0x...",
 *                   value: "0"
 *                 }
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
 *                   msg: "Unsupported fromChainId",
 *                   param: "fromChainId",
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
 *               error: "An unexpected error occurred in /bridge/v2. Please try again later."
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
    check('fromAmount').exists().withMessage('amount is required').isInt(),
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
      .custom((value) => isTokenAddress(value))
      .withMessage('Invalid toToken address')
      .custom((value, { req }) =>
        isTokenSupportedOnChain(value, req.query.toChainId as string)
      )
      .withMessage('Token not supported on specified chain'),
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
  bridgeV2Controller
)

export default router
