import { ponder } from '@/generated'

let sentCount = 0
let receivedCount = 0

ponder.on(
  'InterchainClientV1:InterchainTransactionSent',
  async ({ event, context }) => {
    const {
      db: {
        InterchainTransactionSent,
        InterchainTransaction,
        InterchainBatch,
        AppConfigV1,
      },
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

    // Manually counting until we can do aggregations
    sentCount = sentCount + 1

    const record = await InterchainTransactionSent.create({
      id: transactionId,
      data: {
        name,
        srcChainId: chainId,
        transactionId,
        dbNonce,
        entryIndex,
        dstChainId: Number(dstChainId),
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
        count: sentCount,
      },
    })

    const appConfig =
      (await AppConfigV1.findUnique({ id: dstReceiver })) ??
      (await AppConfigV1.create({
        id: dstReceiver,
        data: {
          requiredResponses: 1,
          optimisticPeriod: 30,
          modules: [],
        },
      }))

    const batch = await InterchainBatch.findMany({
      where: {
        srcDbNonce: dbNonce,
        dstChainId: Number(dstChainId),
      },
    })

    batch.items.forEach(async (b) => {
      await InterchainBatch.update({
        id: b.id,
        data: {
          appConfigId: appConfig.id,
        },
      })
      await InterchainTransaction.upsert({
        id: transactionId,
        update: {
          sentAt: timestamp,
          updatedAt: BigInt(Math.trunc(Date.now() / 1000)),
          interchainTransactionSentId: record.id,
          interchainBatchId: b.id,
          status: 'Sent',
          srcChainId: chainId,
          dstChainId: Number(dstChainId),
        },
        create: {
          sentAt: timestamp,
          createdAt: BigInt(Math.trunc(Date.now() / 1000)),
          interchainTransactionSentId: record.id,
          interchainBatchId: b.id,
          status: 'Sent',
          srcChainId: chainId,
          dstChainId: Number(dstChainId),
        },
      })
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

    // Manually counting until we can do aggregations
    receivedCount = receivedCount + 1

    const record = await InterchainTransactionReceived.create({
      id: transactionId,
      data: {
        name,
        dstChainId: chainId,
        transactionId,
        dbNonce,
        entryIndex,
        srcChainId: Number(srcChainId),
        srcSender,
        dstReceiver,
        address,
        blockNumber,
        transactionHash,
        timestamp,
        count: receivedCount,
      },
    })

    await InterchainTransaction.upsert({
      id: transactionId,
      create: {
        receivedAt: timestamp,
        createdAt: BigInt(Math.trunc(Date.now() / 1000)),
        interchainTransactionReceivedId: record.id,
        status: 'Received',
        srcChainId: Number(srcChainId),
        dstChainId: chainId,
      },
      update: {
        receivedAt: timestamp,
        updatedAt: BigInt(Math.trunc(Date.now() / 1000)),
        interchainTransactionReceivedId: record.id,
        status: 'Received',
        srcChainId: Number(srcChainId),
        dstChainId: chainId,
      },
    })
  }
)

ponder.on(
  'InterchainDB:InterchainBatchFinalized',
  async ({ event, context }) => {
    const {
      db: { InterchainBatch },
    } = context

    const { batchRoot, dbNonce } = event.args

    await InterchainBatch.create({
      id: batchRoot,
      data: {
        batchRoot,
        srcDbNonce: dbNonce,
        status: 'InterchainBatchFinalized',
      },
    })
  }
)

ponder.on(
  'InterchainDB:InterchainBatchVerificationRequested',
  async ({ event, context }) => {
    const {
      db: { InterchainBatch },
    } = context

    const { batchRoot, dstChainId } = event.args

    await InterchainBatch.update({
      id: batchRoot,
      data: {
        dstChainId: Number(dstChainId),
        status: 'InterchainBatchVerificationRequested',
      },
    })
  }
)

ponder.on(
  'InterchainDB:InterchainBatchVerified',
  async ({ event, context }) => {
    const {
      db: { InterchainBatch },
    } = context

    const { srcChainId, dbNonce, batchRoot } = event.args

    const { timestamp } = event.block

    await InterchainBatch.update({
      id: batchRoot,
      data: {
        srcChainId: Number(srcChainId),
        dstDbNonce: dbNonce,
        verifiedAt: timestamp,
        status: 'InterchainBatchVerified',
      },
    })
  }
)
