import express from 'express'
import { check } from 'express-validator'
import { isAddress } from 'ethers/lib/utils'

import { CHAINS_ARRAY } from '../constants/chains'
import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { swapTxInfoController } from '../controllers/swapTxInfoController'
import { isTokenAddress } from '../utils/isTokenAddress'
import { isTokenSupportedOnChain } from '../utils/isTokenSupportedOnChain'
import { checksumAddresses } from '../middleware/checksumAddresses'

const router = express.Router()

/**
 * @openapi
 * /swapTxInfo:
 *   get:
 *     summary: Get swap transaction information
 *     description: Retrieve transaction information for swapping tokens on a specific chain
 *     parameters:
 *       - in: query
 *         name: chain
 *         required: true
 *         schema:
 *           type: string
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
 *         description: The amount of tokens to swap
 *       - in: query
 *         name: address
 *         required: true
 *         schema:
 *           type: string
 *         description: The Ethereum address of the user performing the swap
 *     responses:
 *       200:
 *         description: Successful response
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 data:
 *                   type: string
 *                   description: Encoded transaction data
 *                 to:
 *                   type: string
 *                   description: The address of the contract to interact with
 *                 value:
 *                   type: object
 *                   properties:
 *                     type:
 *                       type: string
 *                       enum: [BigNumber]
 *                     hex:
 *                       type: string
 *                   description: The amount of native currency to send with the transaction
 *             example:
 *               data: "0xb5d1cdd4000000000000000000000000abb4f79430002534df3f62e964d62659a010ef3c000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000000000000000000000000000000000003b9aca0000000000000000000000000000000000000000000000000000000000000000800000000000000000000000007e7a0e201fd38d3adaa9523da6c109a07118c96a000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000000000000000000000000000000000003b96eaed0000000000000000000000000000000000000000000000000000000066ecbb7c00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001116898dda4015ed8ddefb84b6e8bc24528af2d800000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"
 *               to: "0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a"
 *               value:
 *                 type: "BigNumber"
 *                 hex: "0x00"
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
 *                 details:
 *                   type: string
 */
router.get(
  '/',
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
    check('amount').isNumeric().exists().withMessage('amount is required'),
    check('address')
      .exists()
      .withMessage('address is required')
      .custom((value) => isAddress(value))
      .withMessage('Invalid Ethereum address'),
  ],
  showFirstValidationError,
  swapTxInfoController
)

export default router
