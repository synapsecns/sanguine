import { ponder } from '@/generated'

ponder.on(
  'InterchainClientV1:InterchainTransactionSent',
  async ({ event, context }) => {
    const {
      db: { InterchainTransactionSent, InterchainTransaction },
      network: { chainId },
    } = context

    const {
      name,
      log: { address, blockNumber, transactionHash },
      block: { timestamp },
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
        chainId,
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
        timestamp,
      },
    })

    await InterchainTransaction.create({
      id: transactionId,
      data: {
        sentAt: timestamp,
        // TODO: ms or seconds to match timestamp?
        createdAt: BigInt(Date.now()),
        interchainTransactionSentId: record.id,
      },
    })
  }
)

ponder.on(
  'InterchainClientV1:InterchainTransactionReceived',
  async ({ event, context }) => {
    const {
      db: { InterchainTransactionReceived, InterchainTransaction },
      network: { chainId },
    } = context

    const {
      name,
      log: { address, blockNumber, transactionHash },
      block: { timestamp },
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
        chainId,
        transactionId,
        dbNonce,
        entryIndex,
        srcChainId,
        srcSender,
        dstReceiver,
        address,
        blockNumber,
        transactionHash,
        timestamp,
      },
    })

    // TODO: Investigate update vs upsert
    await InterchainTransaction.update({
      id: transactionId,
      data: {
        receivedAt: timestamp,
        // TODO: ms or seconds to match timestamp?
        updatedAt: BigInt(Date.now()),
        interchainTransactionReceivedId: record.id,
      },
    })
  }
)
