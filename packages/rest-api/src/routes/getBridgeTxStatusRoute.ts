import express from 'express'
import { check } from 'express-validator'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { getBridgeTxStatusController } from '../controllers/getBridgeTxStatusController'
import { CHAINS_ARRAY } from '../constants/chains'
import { VALID_BRIDGE_MODULES } from '../constants'

const router = express.Router()

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
      .isString(),
  ],
  showFirstValidationError,
  getBridgeTxStatusController
)

export default router
