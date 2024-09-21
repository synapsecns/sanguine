import express from 'express'

import { conflictingProofsController } from '../controllers/conflictingProofsController'

const router = express.Router()

/**
 * @openapi
 * /conflicting-proofs:
 *   get:
 *     summary: Get conflicting proofs
 *     description: Retrieves a list of transactions where the relayer in the proof differs from the relayer in the relay event
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
 *         description: No conflicting proofs found
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
router.get('/', conflictingProofsController)

export default router
