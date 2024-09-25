import express from 'express'

import { refundedAndRelayedTransactionsController } from '../controllers/refundedAndRelayedController'

const router = express.Router()

/**
 * @openapi
 * /refunded-and-relayed:
 *   get:
 *     summary: Get refunded and relayed transactions
 *     description: Retrieves a list of transactions that have been both refunded and relayed
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
 *                   Bridge:
 *                     type: object
 *                     description: General transaction fields
 *                   BridgeRequest:
 *                     type: object
 *                     description: Deposit information
 *                   BridgeRelay:
 *                     type: object
 *                     description: Relay information
 *                   BridgeRefund:
 *                     type: object
 *                     description: Refund information
 *                   BridgeProof:
 *                     type: object
 *                     description: Proof information (if available)
 *                   BridgeClaim:
 *                     type: object
 *                     description: Claim information (if available)
 *                   BridgeDispute:
 *                     type: object
 *                     description: Dispute information (if available)
 *       404:
 *         description: No refunded and relayed transactions found
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
router.get('/', refundedAndRelayedTransactionsController)

export default router
