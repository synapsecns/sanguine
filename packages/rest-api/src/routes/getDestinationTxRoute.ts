import express from 'express'
import { check } from 'express-validator'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { getDestinationTxController } from '../controllers/getDestinationTxController'

const router = express.Router()

router.get(
  '/',
  [
    check('originChainId')
      .exists()
      .withMessage('originChainId is required')
      .isNumeric(),
    check('txHash').exists().withMessage('txHash is required').isString(),
  ],
  showFirstValidationError,
  getDestinationTxController
)

export default router
