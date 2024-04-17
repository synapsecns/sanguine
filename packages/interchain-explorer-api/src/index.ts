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
        status: 'Sent',
      },
      create: {
        sentAt: timestamp,
        createdAt: BigInt(Math.trunc(Date.now() / 1000)),
        interchainTransactionSentId: record.id,
        status: 'Sent',
      },
    })

    console.log('==========')
    console.log(`on: ${context.network.name}`)
    console.log(`InterchainClientV1:InterchainTransactionSent`)
    console.log(`event.args`, event.args)
    console.log(`event.log.blockHash`, event.log.blockHash)
    console.log('==========')
  }
)

ponder.on(
  'InterchainClientV1:InterchainTransactionReceived',
  async ({ event, context }) => {
    const {
      db: {
        InterchainTransactionSent,
        InterchainTransactionReceived,
        InterchainTransaction,
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
        status: 'Complete',
      },
      update: {
        receivedAt: timestamp,
        updatedAt: BigInt(Math.trunc(Date.now() / 1000)),
        interchainTransactionReceivedId: record.id,
        status: 'Complete',
      },
    })

    console.log('==========')
    console.log(`on: ${context.network.name}`)
    console.log(`InterchainClientV1:InterchainTransactionReceived`)
    console.log(`event.args`, event.args)
    console.log(`event.log.blockHash`, event.log.blockHash)
    console.log('==========')
  }
)

ponder.on('InterchainDB:InterchainEntryWritten', async ({ event, context }) => {
  const {
    db: { InterchainTransactionSent, InterchainTransactionReceived },
  } = context

  console.log('==========')
  console.log(`on: ${context.network.name}`)
  console.log('InterchainDB:InterchainEntryWritten')
  console.log(`event.args`, event.args)
  console.log(`event.log.blockHash`, event.log.blockHash)
  console.log('==========')
})

ponder.on(
  'InterchainClientV1:ExecutionProofWritten',
  async ({ event, context }) => {
    const {
      db: { InterchainTransactionSent, InterchainTransactionReceived },
    } = context

    console.log('==========')
    console.log(`on: ${context.network.name}`)
    console.log('InterchainClientV1:ExecutionProofWritten')
    console.log(`event.args`, event.args)
    console.log(`event.log.blockHash`, event.log.blockHash)
    console.log('==========')
  }
)

ponder.on(
  'InterchainDB:InterchainBatchFinalized',
  async ({ event, context }) => {
    const {
      db: { InterchainTransactionSent, InterchainTransactionReceived },
    } = context

    console.log('==========')
    console.log(`on: ${context.network.name}`)
    console.log('InterchainDB:InterchainBatchFinalized')
    console.log(`event.args`, event.args)
    console.log(`event.log.blockHash`, event.log.blockHash)
    console.log('==========')
  }
)

ponder.on(
  'InterchainDB:InterchainBatchVerificationRequested',
  async ({ event, context }) => {
    const {
      db: { InterchainTransactionSent, InterchainTransactionReceived },
    } = context

    console.log('==========')
    console.log(`on: ${context.network.name}`)
    console.log('InterchainDB:InterchainBatchVerificationRequested')
    console.log(`event.args`, event.args)
    console.log(`event.log.blockHash`, event.log.blockHash)
    console.log('==========')
  }
)

ponder.on(
  'InterchainDB:InterchainBatchVerified',
  async ({ event, context }) => {
    const {
      db: { InterchainTransactionSent, InterchainTransaction },
    } = context

    const entry = await InterchainTransactionSent.findMany({
      where: {
        dbNonce: event.args.dbNonce,
        chainId: Number(event.args.srcChainId),
      },
    })

    console.log('==========')
    console.log(`on: ${context.network.name}`)
    console.log('InterchainDB:InterchainBatchVerified')
    console.log(`event.args`, event.args)
    console.log(`event.log.blockHash`, event.log.blockHash)
    console.log('==========')

    entry.items.forEach(async (item) => {
      await InterchainTransaction.update({
        id: item.id,
        data: {
          status: 'InterchainDB:InterchainBatchVerified',
        },
      })
    })
  }
)

ponder.on(
  'SynapseModule:BatchVerificationRequested',
  async ({ event, context }) => {
    console.log('==========')
    console.log(`on: ${context.network.name}`)
    console.log('SynapseModule:BatchVerificationRequested')
    console.log(`event.args`, event.args)
    console.log(`event.log.blockHash`, event.log.blockHash)
    console.log('==========')
  }
)

ponder.on('SynapseModule:BatchVerified', async ({ event, context }) => {
  console.log('==========')
  console.log(`on: ${context.network.name}`)
  console.log('SynapseModule:BatchVerified')
  console.log(`event.args`, event.args)
  console.log(`event.log.blockHash`, event.log.blockHash)
  console.log('==========')
})
