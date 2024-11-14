import express from 'express'

import { recentInvalidRelaysController } from '../controllers/invalidRelaysController'

const router = express.Router()

/**
 * @openapi
 * /invalid-relays:
 *   get:
 *     summary: Get recent invalid relays
 *     description: Retrieves a list of recent invalid relay events from the past 2 weeks
 *     responses:
 *       200:
 *         description: Successful response (may be an empty array)
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
router.get('/', recentInvalidRelaysController)

export default router
