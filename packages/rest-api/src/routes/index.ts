import express from 'express'

import indexRoute from './indexRoute'
import swapRoute from './swapRoute'
import swapTxInfoRoute from './swapTxInfoRoute'
import bridgeRoute from './bridgeRoute'
import bridgeTxInfoRoute from './bridgeTxInfoRoute'
import synapseTxIdRoute from './synapseTxIdRoute'
import bridgeTxStatusRoute from './bridgeTxStatusRoute'
import destinationTxRoute from './destinationTxRoute'
import tokenListRoute from './tokenListRoute'
import destinationTokensRoute from './destinationTokensRoute'
import bridgeLimitsRoute from './bridgeLimitsRoute'

const router: express.Router = express.Router()

router.use('/', indexRoute)
router.use('/swap', swapRoute)
router.use('/swapTxInfo', swapTxInfoRoute)
router.use('/bridge', bridgeRoute)
router.use('/bridgeTxInfo', bridgeTxInfoRoute)
router.use('/bridgeLimits', bridgeLimitsRoute)
router.use('/synapseTxId', synapseTxIdRoute)
router.use('/bridgeTxStatus', bridgeTxStatusRoute)
router.use('/destinationTx', destinationTxRoute)
router.use('/tokenList', tokenListRoute)
router.use('/destinationTokens', destinationTokensRoute)

export default router
