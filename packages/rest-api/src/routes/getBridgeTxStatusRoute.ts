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
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported destChainId')
      .exists()
      .withMessage('destChainId is required'),
    check('bridgeModule')
      .isString()
      .isIn(VALID_BRIDGE_MODULES)
      .withMessage(
        'Invalid bridge module. Must be one of: ' +
          VALID_BRIDGE_MODULES.join(', ')
      )
      .exists()
      .withMessage('bridgeModule is required'),
    check('synapseTxId')
      .isString()
      .exists()
      .withMessage('synapseTxId is required'),
  ],
  showFirstValidationError,
  getBridgeTxStatusController
)

export default router
