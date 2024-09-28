import express from 'express'
import { check } from 'express-validator'

import { isTokenAddress } from '../utils/isTokenAddress'
import { CHAINS_ARRAY } from '../constants/chains'
import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { bridgeController } from '../controllers/bridgeController'
import { isTokenSupportedOnChain } from '../utils/isTokenSupportedOnChain'
import { checksumAddresses } from '../middleware/checksumAddresses'
import { normalizeNativeTokenAddress } from '../middleware/normalizeNativeTokenAddress'
import { validateRouteExists } from '../validations/validateRouteExists'

const router = express.Router()

/**
 * @openapi
 * /bridge:
 *   get:
 *     summary: Get quotes for bridging tokens between chains
 *     description: Retrieve list of detailed bridge quotes based on origin and destination chains based on tokens and amount
 *     parameters:
 *       - in: query
 *         name: fromChain
 *         required: true
 *         schema:
 *           type: integer
 *         description: The source chain ID
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
 *           type: number
 *         description: The amount of tokens to bridge
 *       - in: query
 *         name: originUserAddress
 *         required: false
 *         schema:
 *           type: string
 *         description: The address of the user on the origin chain
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
 *                   feeAmount:
 *                     $ref: '#/components/schemas/BigNumber'
 *                   feeConfig:
 *                     type: object
 *                     properties:
 *                       bridgeFee:
 *                         type: integer
 *                       minFee:
 *                         $ref: '#/components/schemas/BigNumber'
 *                       maxFee:
 *                         $ref: '#/components/schemas/BigNumber'
 *                   routerAddress:
 *                     type: string
 *                   maxAmountOut:
 *                     $ref: '#/components/schemas/BigNumber'
 *                   originQuery:
 *                     type: object
 *                   destQuery:
 *                     type: object
 *                   estimatedTime:
 *                     type: integer
 *                   bridgeModuleName:
 *                     type: string
 *                   gasDropAmount:
 *                     $ref: '#/components/schemas/BigNumber'
 *                   originChainId:
 *                     type: integer
 *                   destChainId:
 *                     type: integer
 *                   maxAmountOutStr:
 *                     type: string
 *                   bridgeFeeFormatted:
 *                     type: string
 *             example:
 *               - id: "01920c87-7f14-7cdf-90e1-e13b2d4af55f"
 *                 feeAmount:
 *                   type: "BigNumber"
 *                   hex: "0x17d78400"
 *                 feeConfig:
 *                   bridgeFee: 4000000
 *                   minFee:
 *                     type: "BigNumber"
 *                     hex: "0x3d0900"
 *                   maxFee:
 *                     type: "BigNumber"
 *                     hex: "0x17d78400"
 *                 routerAddress: "0xd5a597d6e7ddf373a92C8f477DAAA673b0902F48"
 *                 maxAmountOut:
 *                   type: "BigNumber"
 *                   hex: "0xe89bd2cb27"
 *                 originQuery:
 *                   routerAdapter: "0x0000000000000000000000000000000000000000"
 *                   tokenOut: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
 *                   minAmountOut:
 *                     type: "BigNumber"
 *                     hex: "0xe8d4a51000"
 *                   deadline:
 *                     type: "BigNumber"
 *                     hex: "0x66ecb04b"
 *                   rawParams: "0x"
 *                 destQuery:
 *                   routerAdapter: "0xd5a597d6e7ddf373a92C8f477DAAA673b0902F48"
 *                   tokenOut: "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9"
 *                   minAmountOut:
 *                     type: "BigNumber"
 *                     hex: "0xe89bd2cb27"
 *                   deadline:
 *                     type: "BigNumber"
 *                     hex: "0x66f5e873"
 *                   rawParams: "0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009a2dea7b81cfe3e0011d44d41c5c5142b8d9abdf00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002"
 *                 estimatedTime: 1020
 *                 bridgeModuleName: "SynapseCCTP"
 *                 gasDropAmount:
 *                   type: "BigNumber"
 *                   hex: "0x0110d9316ec000"
 *                 originChainId: 1
 *                 destChainId: 42161
 *                 maxAmountOutStr: "999046.695719"
 *                 bridgeFeeFormatted: "400"
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
 *                 message: "Unsupported fromChain"
 *                 field: "fromChain"
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
 *                 details:
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
    check('fromChain')
      .exists()
      .withMessage('fromChain is required')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported fromChain'),
    check('toChain')
      .exists()
      .withMessage('toChain is required')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported toChain'),
    check('fromToken')
      .exists()
      .withMessage('fromToken is required')
      .custom((value) => isTokenAddress(value))
      .withMessage('Invalid fromToken address')
      .custom((value, { req }) =>
        isTokenSupportedOnChain(value, req.query.fromChain as string)
      )
      .withMessage('Token not supported on specified chain'),
    check('toToken')
      .exists()
      .withMessage('toToken is required')
      .custom((value) => isTokenAddress(value))
      .withMessage('Invalid toToken address')
      .custom((value, { req }) =>
        isTokenSupportedOnChain(value, req.query.toChain as string)
      )
      .withMessage('Token not supported on specified chain'),
    check('amount').isNumeric().exists().withMessage('amount is required'),
    check()
      .custom((_value, { req }) => {
        const { fromChain, toChain, fromToken, toToken } = req.query

        return validateRouteExists(fromChain, fromToken, toChain, toToken)
      })
      .withMessage('No valid route exists for the chain/token combination'),
  ],
  showFirstValidationError,
  bridgeController
)

export default router
