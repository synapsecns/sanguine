import { createSchema } from '@ponder/core'

export default createSchema((p) => ({
  AppConfigV1: p.createTable({
    id: p.string(),
    requiredResponses: p.int(),
    optimisticPeriod: p.int(),
    modules: p.string().list(),
  }),

  InterchainBatch: p.createTable(
    {
      id: p.string(),
      batchRoot: p.string(),
      srcDbNonce: p.bigint(),
      srcChainId: p.int().optional(),
      dstDbNonce: p.bigint().optional(),
      dstChainId: p.int().optional(),
      interchainTransactions: p.many('InterchainTransaction.interchainBatchId'),
      status: p.string(),
      verifiedAt: p.bigint().optional(),
      appConfigV1Id: p.string().references('AppConfigV1.id').optional(),
      appConfigV1: p.one('appConfigV1Id'),
    },
    {
      batchRootIndex: p.index('batchRoot'),
      srcDbNonceIndex: p.index('srcDbNonce'),
      srcChainIdIndex: p.index('srcChainId'),
      dstDbNonceIndex: p.index('dstDbNonce'),
      dstChainIdIndex: p.index('dstChainId'),
      statusIndex: p.index('status'),
      appConfigV1IdIndex: p.index('appConfigV1Id'),
    }
  ),

  InterchainTransactionSent: p.createTable(
    {
      id: p.string(),
      srcChainId: p.int(),
      name: p.string(),
      transactionId: p.string(),
      dbNonce: p.bigint(),
      entryIndex: p.bigint(),
      dstChainId: p.int(),
      srcSender: p.string(),
      dstReceiver: p.string(),
      verificationFee: p.bigint(),
      executionFee: p.bigint(),
      options: p.string(),
      message: p.string(),
      address: p.string(),
      blockNumber: p.bigint(),
      transactionHash: p.string(),
      timestamp: p.bigint(),
    },
    {
      srcChainIdIndex: p.index('srcChainId'),
      nameIndex: p.index('name'),
      transactionIdIndex: p.index('transactionId'),
      dbNonceIndex: p.index('dbNonce'),
      dstChainIdIndex: p.index('dstChainId'),
      srcSenderIndex: p.index('srcSender'),
      dstReceiverIndex: p.index('dstReceiver'),
      addressIndex: p.index('address'),
      transactionHashIndex: p.index('transactionHash'),
    }
  ),

  InterchainTransactionReceived: p.createTable(
    {
      id: p.string(),
      dstChainId: p.int(),
      name: p.string(),
      transactionId: p.string(),
      dbNonce: p.bigint(),
      entryIndex: p.bigint(),
      srcChainId: p.int(),
      srcSender: p.string(),
      dstReceiver: p.string(),
      address: p.string(),
      blockNumber: p.bigint(),
      transactionHash: p.string(),
      timestamp: p.bigint(),
    },
    {
      dstChainIdIndex: p.index('dstChainId'),
      nameIndex: p.index('name'),
      transactionIdIndex: p.index('transactionId'),
      dbNonceIndex: p.index('dbNonce'),
      srcChainIdIndex: p.index('srcChainId'),
      srcSenderIndex: p.index('srcSender'),
      dstReceiverIndex: p.index('dstReceiver'),
      addressIndex: p.index('address'),
      transactionHashIndex: p.index('transactionHash'),
    }
  ),

  InterchainTransaction: p.createTable(
    {
      id: p.string(),
      srcChainId: p.int(),
      dstChainId: p.int(),
      srcSender: p.string(),
      dstReceiver: p.string(),
      sentAt: p.bigint().optional(),
      receivedAt: p.bigint().optional(),
      createdAt: p.bigint().optional(),
      updatedAt: p.bigint().optional(),
      interchainTransactionSentId: p
        .string()
        .references('InterchainTransactionSent.id')
        .optional(),
      interchainTransactionSent: p.one('interchainTransactionSentId'),
      interchainTransactionReceivedId: p
        .string()
        .references('InterchainTransactionReceived.id')
        .optional(),
      interchainTransactionReceived: p.one('interchainTransactionReceivedId'),
      status: p.string().optional(),
      interchainBatchId: p.string().references('InterchainBatch.id').optional(),
      interchainBatch: p.one('interchainBatchId'),
    },
    {
      srcChainIdIndex: p.index('srcChainId'),
      dstChainIdIndex: p.index('dstChainId'),
      srcSenderIndex: p.index('srcSender'),
      dstReceiverIndex: p.index('dstReceiver'),
      interchainTransactionSentIdIndex: p.index('interchainTransactionSentId'),
      interchainTransactionReceivedIdIndex: p.index(
        'interchainTransactionReceivedId'
      ),
      statusIndex: p.index('status'),
      interchainBatchIdIndex: p.index('interchainBatchId'),
    }
  ),
}))
