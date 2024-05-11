import { type Address, trim } from 'viem'

import { ponder } from '@/generated'
import { networkDetails } from '@/ponder.config'

/* ORIGIN CHAIN EVENTS */

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
      },
    })

    const batch = await InterchainBatch.findMany({
      where: {
        srcDbNonce: dbNonce,
        dstChainId: Number(dstChainId),
      },
    })

    batch.items.forEach(async (b) => {
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
  'InterchainDB:InterchainBatchFinalized',
  async ({ event, context }) => {
    const {
      db: { InterchainBatch },
      network: { chainId },
    } = context

    const { batchRoot, dbNonce } = event.args

    await InterchainBatch.create({
      id: batchRoot,
      data: {
        batchRoot,
        srcChainId: chainId,
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

/* DESTINATION CHAIN EVENTS */

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
      client,
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

    if (batch) {
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

      /* TODO */
      /* Currently we can't use client to read contract from a different chain Id */
      /* This needs to be fixed so we generating appConfig table entry when txn is originally sent */
      for (const txn of sentTxns) {
        const dstReceiver = txn?.dstReceiver
        const dstChainId = txn?.dstChainId
        if (dstChainId && dstReceiver && networkDetails[dstChainId]) {
          const { InterchainClientV1 } = networkDetails[dstChainId]
          const response = await client.readContract({
            abi: InterchainClientV1.abi,
            address: InterchainClientV1.address,
            functionName: 'getAppReceivingConfigV1',
            args: [trim(dstReceiver as Address)],
          })

          const [{ requiredResponses, optimisticPeriod }, modules] = response

          const appConfig =
            (await AppConfigV1.findUnique({ id: dstReceiver })) ??
            (await AppConfigV1.create({
              id: dstReceiver,
              data: {
                requiredResponses: Number(requiredResponses),
                optimisticPeriod: Number(optimisticPeriod),
                modules,
              },
            }))

          await InterchainBatch.update({
            id: batch.id,
            data: {
              appConfigV1Id: appConfig.id,
            },
          })
        }
      }

      // checks srcChainId, dbNonce match values from the InterchainTransactionSent event
      const allValidTransactions = sentTxns.every(
        (t) => t && t.srcChainId === Number(srcChainId) && t.dbNonce === dbNonce
      )
      const currentBatch = await InterchainBatch.findUnique({ id: batchRoot })
      const appConfigV1 =
        currentBatch &&
        (await AppConfigV1.findUnique({
          id: currentBatch.appConfigV1Id as string,
        }))

      // check that module is in the list of modules that dst receiver trusts
      if (
        allValidTransactions &&
        appConfigV1 &&
        appConfigV1.modules?.includes(module)
      ) {
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
