import express from 'express'
import { check } from 'express-validator'

import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { getDestinationTxController } from '../controllers/getDestinationTxController'

const router = express.Router()

router.get(
  '/',
  [
    check('originChainId')
      .isNumeric()
      .exists()
      .withMessage('originChainId is required'),
    check('txHash').isString().exists().withMessage('txHash is required'),
  ],
  showFirstValidationError,
  getDestinationTxController
)

export default router
