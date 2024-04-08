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
  }),

  InterchainTransaction: p.createTable({
    id: p.string(),
    sentAt: p.bigint(),
    receivedAt: p.bigint().optional(),
    createdAt: p.bigint(),
    updatedAt: p.bigint().optional(),
    interchainTransactionSentId: p
      .string()
      .references('InterchainTransactionSent.id'),
    interchainTransactionSent: p.one('interchainTransactionSentId'),
    interchainTransactionReceivedId: p
      .string()
      .references('InterchainTransactionReceived.id')
      .optional(),
    interchainTransactionReceived: p.one('interchainTransactionReceivedId'),
  }),
}))
