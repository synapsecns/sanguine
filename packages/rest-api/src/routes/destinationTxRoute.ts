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
 *           type: integer
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
 *                   type: string
 *                   description: The status of the transaction
 *                 fromInfo:
 *                   type: object
 *                   properties:
 *                     chainID:
 *                       type: integer
 *                       description: The origin chain ID
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
 *               status: "completed"
 *               fromInfo:
 *                 chainID: 8453
 *                 address: "0x6eA4207627aAf2Ef86642eD8B331579b606471c3"
 *                 txnHash: "0x13486d9eaefd68de6a20b704d70deb8436effbac1f77fddfc0c7ef14f08e96c3"
 *                 USDValue: 11660.93019,
 *                 tokenSymbol: "USDC"
 *                 blockNumber: 19857812,
 *                 formattedTime: "2024-09-16 16:42:51 +0000 UTC"
 *                 formattedValue: "11637.654884"
 *               toInfo:
 *                 chainID: 42161
 *                 address: "0xfC8f27Bcf34FfD52869ffa4A5A6B9b0A872281Ad"
 *                 txnHash: "0xe26be8f4296c14dc8da6ef92d39c1d20577a43704bfb0b2cea5ee2f516be0f4e"
 *                 USDValue: 11660.92558
 *                 tokenSymbol: "USDC"
 *                 blockNumber: 254173724
 *                 formattedTime: "2024-09-16 16:42:55 +0000 UTC"
 *                 formattedValue: "11637.650281"
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
 *       404:
 *         description: Not found
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 status:
 *                   type: string
 *                 fromInfo:
 *                   type: null
 *                 toInfo:
 *                   type: null
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
