import express from 'express'

import indexRoute from './indexRoute'
import swapRoute from './swapRoute'
import swapTxInfoRoute from './swapTxInfoRoute'
import bridgeRoute from './bridgeRoute'
import bridgeTxInfoRoute from './bridgeTxInfoRoute'
import getBridgeLimitsRoute from './getBridgeLimitsRoute'
import getSynapseTxIdRoute from './getSynapseTxIdRoute'
import getBridgeTxStatusRoute from './getBridgeTxStatusRoute'
import getDestinationTxRoute from './getDestinationTxRoute'
import tokenListRoute from './tokenListRoute'
import destinationTokensRoute from './destinationTokensRoute'

const router = express.Router()

router.use('/', indexRoute)
router.use('/swap', swapRoute)
router.use('/swapTxInfo', swapTxInfoRoute)
router.use('/bridge', bridgeRoute)
router.use('/bridgeTxInfo', bridgeTxInfoRoute)
router.use('/getBridgeLimits', getBridgeLimitsRoute)
router.use('/getSynapseTxId', getSynapseTxIdRoute)
router.use('/getBridgeTxStatus', getBridgeTxStatusRoute)
router.use('/getDestinationTx', getDestinationTxRoute)
router.use('/tokenList', tokenListRoute)
router.use('/destinationTokens', destinationTokensRoute)

export default router
