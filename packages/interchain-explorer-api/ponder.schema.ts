import { createSchema } from '@ponder/core'

export default createSchema((p) => ({
  InterchainTransactionSent: p.createTable({
    id: p.string(),
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
  }),

  InterchainTransactionReceived: p.createTable({
    id: p.string(),
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
  }),

  InterchainTransaction: p.createTable({
    id: p.string(),
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
