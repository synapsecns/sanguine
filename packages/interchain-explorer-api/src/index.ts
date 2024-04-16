import { ponder } from '@/generated'

let sentCount = 0
let receivedCount = 0

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

    // Manually counting until we can do aggregations
    sentCount = sentCount + 1

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
        count: sentCount,
      },
    })

    await InterchainTransaction.upsert({
      id: transactionId,
      update: {
        sentAt: timestamp,
        updatedAt: BigInt(Math.trunc(Date.now() / 1000)),
        interchainTransactionSentId: record.id,
      },
      create: {
        sentAt: timestamp,
        createdAt: BigInt(Math.trunc(Date.now() / 1000)),
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

    // Manually counting until we can do aggregations
    receivedCount = receivedCount + 1

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
        count: receivedCount,
      },
    })

    await InterchainTransaction.upsert({
      id: transactionId,
      create: {
        receivedAt: timestamp,
        createdAt: BigInt(Math.trunc(Date.now() / 1000)),
        interchainTransactionReceivedId: record.id,
      },
      update: {
        receivedAt: timestamp,
        updatedAt: BigInt(Math.trunc(Date.now() / 1000)),
        interchainTransactionReceivedId: record.id,
      },
    })
  }
)

ponder.on(
  'InterchainDB:InterchainBatchVerificationRequested',
  async ({ event, context }) => {
    const {
      db: { InterchainTransactionSent },
    } = context

    console.log(event.args.dbNonce)

    const entry = await InterchainTransactionSent.findMany({
      where: { dstChainId: event.args.dstChainId, dbNonce: event.args.dbNonce },
    })

    console.log('==========')
    console.log('InterchainDB:InterchainBatchVerificationRequested')
    console.log(`event.args.dbNonce`, event.args.dbNonce)
    console.log(`event`, event)
    console.log(`entry`, entry)
    console.log('==========')
  }
)

ponder.on(
  'InterchainDB:InterchainBatchVerified',
  async ({ event, context }) => {
    const {
      db: { InterchainTransactionSent },
    } = context

    const entry = await InterchainTransactionSent.findMany({
      where: { dbNonce: event.args.dbNonce },
    })

    console.log('==========')
    console.log('InterchainDB:InterchainBatchVerified')
    console.log(`event.args.dbNonce`, event.args.dbNonce)
    console.log(`event`, event)
    console.log(`entry`, entry)
    console.log('==========')
  }
)

ponder.on('SynapseModule:BatchVerificationRequested', async ({ event }) => {
  console.log('==========')
  console.log('SynapseModule:BatchVerificationRequested')
  console.log(`event`, event)
  console.log('==========')
})

ponder.on('SynapseModule:BatchVerified', async ({ event }) => {
  console.log('==========')
  console.log('SynapseModule:BatchVerified')
  console.log(`event`, event)
  console.log('==========')
})
