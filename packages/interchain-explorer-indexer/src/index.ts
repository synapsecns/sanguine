import { type Address, trim } from 'viem'

import { ponder } from '@/generated'

let sentCount = 0
let receivedCount = 0

interface TempAppConfig {
  [key: string]: {
    [key: number]: {
      requiredResponses: number
      optimisticPeriod: number
      modules: Address[]
    }
  }
}

const TEMP_APP_CONFIG: TempAppConfig = {
  '0x521931f62298605de22485bb72a86d599f43f823': {
    11155111: {
      requiredResponses: 1,
      optimisticPeriod: 30,
      modules: ['0x95f2e2fAFE38f2aAdC9F9cBef98785809cc4bb6B'],
    },
  },
  '0x67829fee24ae01e2a7f9f09f36bf96fcf7771738': {
    421614: {
      requiredResponses: 1,
      optimisticPeriod: 30,
      modules: ['0xC13e2b478f6531Ef096FF05733Ed65E3bc7fC5AF'],
    },
  },
}

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

    /* Currently we can't use client to read contract from dst chain */
    // const response = await client.readContract({
    //   abi: InterchainClientV1.abi,
    //   address: InterchainClientV1.address as Address,
    //   functionName: 'getAppReceivingConfigV1',
    //   args: [trim(dstReceiver)],
    // })

    const { requiredResponses, optimisticPeriod, modules } = TEMP_APP_CONFIG[
      trim(dstReceiver)
    ]?.[Number(dstChainId)] ?? {
      requiredResponses: 1,
      optimisticPeriod: 30,
      modules: [],
    }

    const appConfig =
      (await AppConfigV1.findUnique({ id: dstReceiver })) ??
      (await AppConfigV1.create({
        id: dstReceiver,
        data: {
          requiredResponses,
          optimisticPeriod,
          modules,
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
          srcSender,
          dstReceiver,
        },
        create: {
          sentAt: timestamp,
          createdAt: BigInt(Math.trunc(Date.now() / 1000)),
          interchainTransactionSentId: record.id,
          interchainBatchId: b.id,
          status: 'Sent',
          srcChainId: chainId,
          dstChainId: Number(dstChainId),
          srcSender,
          dstReceiver,
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
        srcSender,
        dstReceiver,
      },
      update: {
        receivedAt: timestamp,
        updatedAt: BigInt(Math.trunc(Date.now() / 1000)),
        interchainTransactionReceivedId: record.id,
        status: 'Received',
        srcChainId: Number(srcChainId),
        dstChainId: chainId,
        srcSender,
        dstReceiver,
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
      db: {
        InterchainBatch,
        InterchainTransaction,
        InterchainTransactionSent,
        AppConfigV1,
      },
    } = context

    const { module, srcChainId, dbNonce, batchRoot } = event.args

    const { timestamp } = event.block

    /*
      1. batchRoot matches the value from InterchainBatchFinalized event

      2. module is in the list of modules that dst receiver trusts

      3. srcChainId, dbNonce match values from the InterchainTransactionSent event
    */

    // check if batchRoot matches the value from InterchainBatchFinalized event
    const batch = await InterchainBatch.findUnique({
      id: batchRoot,
    })

    if (batch && batch.appConfigId) {
      const appConfig = await AppConfigV1.findUnique({ id: batch.appConfigId })

      const modules = appConfig?.modules

      const txns = await InterchainTransaction.findMany({
        where: {
          interchainBatchId: batch.batchRoot,
        },
      })

      const sentTxns = await Promise.all(
        txns.items.map(async (t) =>
          InterchainTransactionSent.findUnique({
            id: t.interchainTransactionSentId as string,
          })
        )
      )

      // checks srcChainId, dbNonce match values from the InterchainTransactionSent event
      const allValidTransactions = sentTxns.every(
        (t) => t && t.srcChainId === Number(srcChainId) && t.dbNonce === dbNonce
      )

      // check that module is in the list of modules that dst receiver trusts
      if (allValidTransactions && modules?.includes(module)) {
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
    }
  }
)
