import express from 'express'
import { check } from 'express-validator'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { destinationTxController } from '../controllers/destinationTxController'

const router = express.Router()

/**
 * @openapi
 * /destinationTx:
 *   get:
 *     summary: Get Destination Transaction Information
 *     description: Used to get the status of a bridge transaction, and the destination transaction information if the transaction is finalized
 *     parameters:
 *       - in: query
 *         name: originChainId
 *         required: true
 *         schema:
 *           type: string
 *         description: The ID of the origin chain where the transaction was initiated
 *       - in: query
 *         name: txHash
 *         required: true
 *         schema:
 *           type: string
 *         description: The transaction hash on the origin chain
 *     responses:
 *       200:
 *         description: Successful response
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 status:
 *                   type: boolean
 *                   description: The status of the transaction
 *                 toInfo:
 *                   type: object
 *                   properties:
 *                     chainID:
 *                       type: integer
 *                       description: The destination chain ID
 *                     address:
 *                       type: string
 *                       description: The recipient address
 *                     txnHash:
 *                       type: string
 *                       description: The transaction hash on the destination chain
 *                     USDValue:
 *                       type: number
 *                       description: The USD value of the transaction
 *                     tokenSymbol:
 *                       type: string
 *                       description: The symbol of the token transferred
 *                     formattedTime:
 *                       type: string
 *                       description: The formatted time of the transaction
 *                     formattedValue:
 *                       type: string
 *                       description: The formatted value of the transaction
 *             example:
 *               status: true
 *               toInfo:
 *                 chainID: 8453
 *                 address: "0xABb4F79430002534df3F62E964D62659A010Ef3C"
 *                 txnHash: "0xc9284b2de9ba74ab618573884930e51575c1a3511216d9949da2955efb69afa8"
 *                 USDValue: 5999.98657
 *                 tokenSymbol: "USDC"
 *                 formattedTime: "2024-09-19 23:32:29 +0000 UTC"
 *                 formattedValue: "5999.986575"
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
 *                 message: "originChainId is required"
 *                 field: "originChainId"
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
  [
    check('originChainId')
      .exists()
      .withMessage('originChainId is required')
      .isNumeric(),
    check('txHash').exists().withMessage('txHash is required').isString(),
  ],
  showFirstValidationError,
  destinationTxController
)

export default router
