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
 *         name: fromChain
 *         required: true
 *         schema:
 *           type: integer
 *         description: The source chain ID (must support intents)
 *       - in: query
 *         name: toChain
 *         required: true
 *         schema:
 *           type: integer
 *         description: The destination chain ID
 *       - in: query
 *         name: fromToken
 *         required: true
 *         schema:
 *           type: string
 *         description: The address of the token on the source chain
 *       - in: query
 *         name: toToken
 *         required: true
 *         schema:
 *           type: string
 *         description: The address of the token on the destination chain
 *       - in: query
 *         name: amount
 *         required: true
 *         schema:
 *           type: integer
 *         description: The amount of tokens to bridge in the token's native decimals
 *       - in: query
 *         name: originUserAddress
 *         required: false
 *         schema:
 *           type: string
 *         description: The address of the user on the origin chain (required to generate callData)
 *       - in: query
 *         name: destAddress
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
 *                   routerAddress:
 *                     type: string
 *                     description: The address of the router contract
 *                   maxAmountOut:
 *                     $ref: '#/components/schemas/BigNumber'
 *                     description: The maximum amount of tokens that will be received (BigNumber format)
 *                   maxAmountOutStr:
 *                     type: string
 *                     description: String representation of maxAmountOut in native token decimals
 *                   bridgeModule:
 *                     type: string
 *                     description: The name of the bridge module used for this quote
 *                   estimatedTime:
 *                     type: integer
 *                     description: Estimated time for the bridge in seconds
 *                   originChainId:
 *                     type: integer
 *                     description: The ID of the origin chain
 *                   destChainId:
 *                     type: integer
 *                     description: The ID of the destination chain
 *                   gasDropAmount:
 *                     $ref: '#/components/schemas/BigNumber'
 *                     description: Amount of native token airdropped on destination chain (BigNumber format)
 *                   callData:
 *                     type: object
 *                     nullable: true
 *                     description: Transaction data object, only provided if addresses were included
 *                     properties:
 *                       to:
 *                         type: string
 *                         description: Contract address to call
 *                       data:
 *                         type: string
 *                         description: Transaction calldata
 *                       value:
 *                         type: string
 *                         description: Amount of native currency to send with transaction
 *             example:
 *               - id: "01920c87-7f14-7cdf-90e1-e13b2d4af55f"
 *                 routerAddress: "0x512000a034E154908Efb1eC48579F4ffDb000512"
 *                 maxAmountOut: {
 *                   type: "BigNumber",
 *                   hex: "0xe89bd2cb27"
 *                 }
 *                 maxAmountOutStr: "999046695719"
 *                 bridgeModule: "SynapseRFQ"
 *                 estimatedTime: 30
 *                 originChainId: 1
 *                 destChainId: 42161
 *                 gasDropAmount: {
 *                   type: "BigNumber",
 *                   hex: "0x00"
 *                 }
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
 *               error: "An unexpected error occurred in /bridge/v2. Please try again later."
 */
router.get(
  '/',
  normalizeNativeTokenAddress(['fromToken', 'toToken']),
  checksumAddresses(['fromToken', 'toToken']),
  [
    check('fromChain')
      .exists()
      .withMessage('fromChain is required')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported fromChain')
      .custom((value) => INTENTS_SUPPORTED_CHAIN_IDS.includes(Number(value)))
      .withMessage('Intents not supported for given chain'),
    check('toChain')
      .exists()
      .withMessage('toChain is required')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported toChain'),
    check('fromToken')
      .exists()
      .withMessage('fromToken is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid fromToken address'),
    check('toToken')
      .exists()
      .withMessage('toToken is required')
      .custom((value) => isTokenAddress(value))
      .withMessage('Invalid toToken address')
      .custom((value, { req }) =>
        isTokenSupportedOnChain(value, req.query.toChain as string)
      )
      .withMessage('Token not supported on specified chain'),
    check('amount').exists().withMessage('amount is required').isInt(),
    check('originUserAddress')
      .optional()
      .custom((value) => isAddress(value))
      .withMessage('Invalid originUserAddress address'),
    check('destAddress')
      .optional()
      .custom((value) => isAddress(value))
      .withMessage('Invalid destAddress'),
    check('slippage')
      .optional()
      .isNumeric()
      .custom((value) => value >= 0 && value <= 100)
      .withMessage('Slippage must be between 0 and 100'),
  ],
  showFirstValidationError,
  bridgeV2Controller
)

export default router
