import express from 'express'
import { check } from 'express-validator'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { getSynapseTxIdController } from '../controllers/getSynapseTxIdController'
import { VALID_BRIDGE_MODULES } from '../constants'

const router = express.Router()

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
