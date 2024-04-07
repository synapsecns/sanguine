import { ponder } from '@/generated'

ponder.on(
  'InterchainClientV1:InterchainTransactionSent',
  async ({ event, context }) => {
    const { InterchainTransactionSent, InterchainTransaction } = context.db

    const {
      name,
      log: { address, blockNumber, transactionHash },
      args: {
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
      },
    } = event

    const record = await InterchainTransactionSent.create({
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

    await InterchainTransaction.create({
      id: transactionId,
      data: {
        interchainTransactionSentId: record.id,
        interchainTransactionReceivedId: undefined,
      },
    })
  }
)

ponder.on(
  'InterchainClientV1:InterchainTransactionReceived',
  async ({ event, context }) => {
    const { InterchainTransactionReceived, InterchainTransaction } = context.db

    const {
      name,
      log: { address, blockNumber, transactionHash },
      args: {
        transactionId,
        dbNonce,
        entryIndex,
        srcChainId,
        srcSender,
        dstReceiver,
      },
    } = event

    const record = await InterchainTransactionReceived.create({
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

    await InterchainTransaction.update({
      id: transactionId,
      data: {
        interchainTransactionReceivedId: record.id,
      },
    })
  }
)
