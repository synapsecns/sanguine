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
      .isNumeric()
      .exists()
      .withMessage('originChainId is required'),
    check('bridgeModule')
      .isString()
      .isIn(VALID_BRIDGE_MODULES)
      .withMessage(
        'Invalid bridge module. Must be one of: ' +
          VALID_BRIDGE_MODULES.join(', ')
      )
      .exists()
      .withMessage('bridgeModule is required'),
    check('txHash').isString().exists().withMessage('txHash is required'),
  ],
  showFirstValidationError,
  getSynapseTxIdController
)

export default router
