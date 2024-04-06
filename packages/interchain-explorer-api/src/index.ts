import { ponder } from '@/generated'

ponder.on(
  'InterchainClientV1:InterchainTransactionSent',
  async ({ event, context }) => {
    const { InterchainTransactionSent } = context.db

    const { name } = event
    const { address, blockNumber, transactionHash } = event.log

    const {
      transactionId,
      dbNonce,
      entryIndex,
      dstChainId,
      srcSender,
      dstReceiver,
      verificationFee,
      executionFee,
      options,
      message,
    } = event.args

    await InterchainTransactionSent.create({
      id: transactionId,
      data: {
        name,
        transactionId,
        dbNonce,
        entryIndex,
        dstChainId,
        srcSender,
        dstReceiver,
        verificationFee,
        executionFee,
        options,
        message,
        address,
        blockNumber,
        transactionHash,
      },
    })
  }
)

ponder.on(
  'InterchainClientV1:InterchainTransactionReceived',
  async ({ event, context }) => {
    const { InterchainTransactionReceived } = context.db

    const { name } = event
    const { address, blockNumber, transactionHash } = event.log
    const {
      transactionId,
      dbNonce,
      entryIndex,
      srcChainId,
      srcSender,
      dstReceiver,
    } = event.args

    await InterchainTransactionReceived.create({
      id: transactionId,
      data: {
        name,
        transactionId,
        dbNonce,
        entryIndex,
        srcChainId,
        srcSender,
        dstReceiver,
        address,
        blockNumber,
        transactionHash,
      },
    })
  }
)
