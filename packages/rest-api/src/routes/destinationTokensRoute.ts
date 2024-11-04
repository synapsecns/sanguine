import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

import { CHAINS_ARRAY } from '../constants/chains'
import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { destinationTokensController } from '../controllers/destinationTokensController'
import { isTokenAddress } from '../utils/isTokenAddress'
import { isTokenSupportedOnChain } from '../utils/isTokenSupportedOnChain'
import { checksumAddresses } from '../middleware/checksumAddresses'
import { normalizeNativeTokenAddress } from '../middleware/normalizeNativeTokenAddress'

const router: express.Router = express.Router()

/**
 * @openapi
 * /destinationTokens:
 *   get:
 *     summary: Get possible destination tokens for a bridge
 *     description: Retrieve possible destination tokens for a given source chain ID and token address
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
 *                   symbol:
 *                     type: string
 *                     description: The token symbol
 *                   chainId:
 *                     type: string
 *                     description: The chain ID where the token is available
 *                   address:
 *                     type: string
 *                     description: The token contract address
 *             example:
 *               - symbol: "USDC"
 *                 chainId: "1"
 *                 address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
 *               - symbol: "USDT"
 *                 chainId: "42161"
 *                 address: "0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9"
 *               - symbol: "crvUSD"
 *                 chainId: "8453"
 *                 address: "0x417Ac0e078398C154EdFadD9Ef675d30Be60Af93"
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
 *                 value: "100"
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
  normalizeNativeTokenAddress(['fromToken']),
  checksumAddresses(['fromToken']),
  [
    check('fromChain')
      .exists()
      .withMessage('fromChain is required')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported fromChain'),
    check('fromToken')
      .exists()
      .withMessage('fromToken is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid fromToken address')
      .custom((value) => isTokenAddress(value))
      .withMessage('Unsupported fromToken address')
      .custom((value, { req }) =>
        isTokenSupportedOnChain(value, req.query.fromChain as string)
      )
      .withMessage('Token not supported on specified chain'),
  ],
  showFirstValidationError,
  destinationTokensController
)

export default router
