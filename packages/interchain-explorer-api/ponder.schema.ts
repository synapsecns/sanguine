import { createSchema } from '@ponder/core'

export default createSchema((p) => ({
  InterchainTransactionSent: p.createTable({
    id: p.string(),
    chainId: p.int(),
    name: p.string(),
    transactionId: p.string(),
    dbNonce: p.bigint(),
    entryIndex: p.bigint(),
    dstChainId: p.bigint(),
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
    count: p.int(),
  }),

  InterchainTransactionReceived: p.createTable({
    id: p.string(),
    chainId: p.int(),
    name: p.string(),
    transactionId: p.string(),
    dbNonce: p.bigint(),
    entryIndex: p.bigint(),
    srcChainId: p.bigint(),
    srcSender: p.string(),
    dstReceiver: p.string(),
    address: p.string(),
    blockNumber: p.bigint(),
    transactionHash: p.string(),
    timestamp: p.bigint(),
    count: p.int(),
  }),

  InterchainTransaction: p.createTable({
    id: p.string(),
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
  }),
}))
