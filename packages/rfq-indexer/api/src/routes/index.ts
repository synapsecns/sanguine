import express from 'express'

import pendingTransactionsRoute from './pendingTransactionsRoute'
import refundedAndRelayedRoute from './refundedAndRelayedRoute'
import invalidRelaysRoute from './invalidRelaysRoute'
import conflictingProofsRoute from './conflictingProofsRoute'
import transactionIdRoute from './transactionIdRoute'

const router = express.Router()

router.use('/pending-transactions', pendingTransactionsRoute)
router.use('/refunded-and-relayed', refundedAndRelayedRoute)
router.use('/invalid-relays', invalidRelaysRoute)
router.use('/conflicting-proofs', conflictingProofsRoute)
router.use('/transaction-id', transactionIdRoute)

export default router