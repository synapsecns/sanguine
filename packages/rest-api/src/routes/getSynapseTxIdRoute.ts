import express from 'express'
import { check } from 'express-validator'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { getSynapseTxIdController } from '../controllers/getSynapseTxIdController'
import { VALID_BRIDGE_MODULES } from '../constants'

const router = express.Router()

/**
 * @openapi
 * /getSynapseTxId:
 *   get:
 *     summary: Get Synapse Transaction ID
 *     description: Retrieve the Synapse transaction ID for a given origin chain transaction
 *     parameters:
 *       - in: query
 *         name: originChainId
 *         required: true
 *         schema:
 *           type: string
 *         description: The ID of the origin chain where the transaction was initiated
 *       - in: query
 *         name: bridgeModule
 *         required: true
 *         schema:
 *           type: string
 *           enum: [SynapseRFQ, SynapseBridge, SynapseCCTP]
 *         description: The bridge module used for the transaction
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
 *                 synapseTxId:
 *                   type: string
 *                   description: The Synapse transaction ID
 *             example:
 *               synapseTxId: "0x812516c5477aeeb4361ecbdd561abcd10f779a0fce22bad13635b8cae088760a"
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
 *                 value: "SomeOtherModule"
 *                 message: "Invalid bridge module. Must be one of: SynapseRFQ, SynapseBridge, SynapseCCTP"
 *                 field: "bridgeModule"
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
    check('bridgeModule')
      .exists()
      .withMessage('bridgeModule is required')
      .isString()
      .isIn(VALID_BRIDGE_MODULES)
      .withMessage(
        'Invalid bridge module. Must be one of: ' +
          VALID_BRIDGE_MODULES.join(', ')
      ),
    check('txHash').exists().withMessage('txHash is required').isString(),
  ],
  showFirstValidationError,
  getSynapseTxIdController
)

export default router
