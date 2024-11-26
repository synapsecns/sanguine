import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

import { CHAINS_ARRAY } from '../constants/chains'
import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { bridgeTxInfoController } from '../controllers/bridgeTxInfoController'
import { isTokenAddress } from '../utils/isTokenAddress'
import { isTokenSupportedOnChain } from '../utils/isTokenSupportedOnChain'
import { checksumAddresses } from '../middleware/checksumAddresses'
import { normalizeNativeTokenAddress } from '../middleware/normalizeNativeTokenAddress'
import { validateRouteExists } from '../validations/validateRouteExists'
import { validateDecimals } from '../validations/validateDecimals'
import { tokenAddressToToken } from '../utils/tokenAddressToToken'

const router: express.Router = express.Router()

/**
 * @openapi
 * /bridgeTxInfo:
 *   get:
 *     summary: "[Deprecated] in favor of using the /bridge endpoint, which now returns call data"
 *     description: "[Deprecated] Originally used to get Bridge transaction information"
 *     parameters:
 *       - in: query
 *         name: fromChain
 *         required: true
 *         schema:
 *           type: integer
 *         description: The source chain ID
 *       - in: query
 *         name: fromToken
 *         required: true
 *         schema:
 *           type: string
 *         description: The address of the token on the source chain
 *       - in: query
 *         name: toChain
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
 *         name: amount
 *         required: true
 *         schema:
 *           type: number
 *         description: The amount of tokens to bridge
 *       - in: query
 *         name: destAddress
 *         required: true
 *         schema:
 *           type: string
 *         description: The destination address for the bridged tokens
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
 *                   data:
 *                     type: string
 *                     description: Encoded transaction data
 *                   to:
 *                     type: string
 *                     description: The address of the contract to interact with
 *                   value:
 *                     type: object
 *                     properties:
 *                       type:
 *                         type: string
 *                         enum: [BigNumber]
 *                       hex:
 *                         type: string
 *                     description: The amount of native currency to send with the transaction
 *             example:
 *               - data: "0xc2288147000000000000000000000000abb4f79430002534df3f62e964d62659a010ef3c000000000000000000000000000000000000000000000000000000000000a4b1000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000000000000000000000000000000000174876e80000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000000000000000000000000000000000174876e8000000000000000000000000000000000000000000000000000000000066ecbadf00000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d5a597d6e7ddf373a92c8f477daaa673b0902f48000000000000000000000000fd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb90000000000000000000000000000000000000000000000000000001744e380400000000000000000000000000000000000000000000000000000000066f5f30700000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009a2dea7b81cfe3e0011d44d41c5c5142b8d9abdf00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004"
 *                 to: "0xd5a597d6e7ddf373a92C8f477DAAA673b0902F48"
 *                 value:
 *                   type: "BigNumber"
 *                   hex: "0x00"
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
    check('amount')
      .exists()
      .withMessage('amount is required')
      .isNumeric()
      .custom((value, { req }) => {
        const fromTokenInfo = tokenAddressToToken(
          req.query.fromChain,
          req.query.fromToken
        )
        return validateDecimals(value, fromTokenInfo.decimals)
      })
      .withMessage(
        'Amount has too many decimals, beyond the maximum allowed for this token'
      ),
    check('destAddress')
      .exists()
      .withMessage('destAddress is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid destination address'),
    check()
      .custom((_value, { req }) => {
        const { fromChain, toChain, fromToken, toToken } = req.query

        return validateRouteExists(fromChain, fromToken, toChain, toToken)
      })
      .withMessage('No valid route exists for the chain/token combination'),
    check('originUserAddress')
      .optional()
      .custom((value) => isAddress(value))
      .withMessage('Invalid originUserAddress address'),
  ],
  showFirstValidationError,
  bridgeTxInfoController
)

export default router
