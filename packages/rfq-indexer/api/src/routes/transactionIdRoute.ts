import express from 'express'

import { getTransactionById } from '../controllers/transactionIdController'

const router = express.Router()

/**
 * @openapi
 * /transaction-id/{transactionId}:
 *   get:
 *     summary: Get transaction details by ID
 *     description: Retrieves detailed information about a transaction, including deposit, relay, proof, claim, and refund data if available
 *     parameters:
 *       - in: path
 *         name: transactionId
 *         required: true
 *         schema:
 *           type: string
 *         description: The unique identifier of the transaction
 *     responses:
 *       200:
 *         description: Successful response
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
*                  Bridge:
*                    type: object
*                    description: General transaction fields
*                  BridgeRequest:
*                    type: object
*                    description: Deposit information
*                  BridgeRelay:
*                    type: object
*                    description: Relay information
*                  BridgeRefund:
*                    type: object
*                    description: Refund information
*                  BridgeProof:
*                    type: object
*                    description: Proof information (if available)
*                  BridgeClaim:
*                    type: object
*                    description: Claim information (if available)
*                  BridgeDispute:
*                    type: object
*                    description: Dispute information (if available)
 *       404:
 *         description: Transaction not found
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 message:
 *                   type: string
 *       500:
 *         description: Server error
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 message:
 *                   type: string
 */
router.get('/:transactionId', getTransactionById)

export default router
