import express from 'express'
import { check } from 'express-validator'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { bridgeTxStatusController } from '../controllers/bridgeTxStatusController'
import { CHAINS_ARRAY } from '../constants/chains'
import { VALID_BRIDGE_MODULES } from '../constants'
import { validateKappa } from '../validations/validateKappa'

const router: express.Router = express.Router()

/**
 * @openapi
 * /bridgeTxStatus:
 *   get:
 *     summary: Get Bridge Transaction Status
 *     description: Used to get the status of a bridge transaction, and the destination transaction information if the transaction is finalized
 *     parameters:
 *       - in: query
 *         name: destChainId
 *         required: true
 *         schema:
 *           type: integer
 *         description: The ID of the destination chain
 *       - in: query
 *         name: bridgeModule
 *         required: true
 *         schema:
 *           type: string
 *           enum: [SynapseRFQ, SynapseBridge, SynapseCCTP]
 *         description: The bridge module used for the transaction
 *       - in: query
 *         name: synapseTxId
 *         required: true
 *         schema:
 *           type: string
 *         description: The Synapse transaction ID
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
 *                 chainID: 10
 *                 address: "0xRL3Bab0e4c09Ff447863f507E16090A9F22792d2"
 *                 txnHash: "0x4eff784e85df5265dcc8e3c30b9df4b5c8a0c940300f6d8ad7ed737e9beb6fab"
 *                 USDValue: 1.79848
 *                 tokenSymbol: "USDC"
 *                 formattedTime: "2024-09-01 17:10:41 +0000 UTC"
 *                 formattedValue: "1.797684"
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
 *                 message: "Unsupported destChainId"
 *                 field: "destChainId"
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
  [
    check('destChainId')
      .exists()
      .withMessage('destChainId is required')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported destChainId'),
    check('bridgeModule')
      .exists()
      .withMessage('bridgeModule is required')
      .isString()
      .isIn(VALID_BRIDGE_MODULES)
      .withMessage(
        'Invalid bridge module. Must be one of: ' +
          VALID_BRIDGE_MODULES.join(', ')
      ),
    check('synapseTxId')
      .exists()
      .withMessage('synapseTxId is required')
      .isString()
      .withMessage('synapseTxId must be a string')
      .custom((value) => validateKappa(value))
      .withMessage('synapseTxId must be valid hex string'),
  ],
  showFirstValidationError,
  bridgeTxStatusController
)

export default router
