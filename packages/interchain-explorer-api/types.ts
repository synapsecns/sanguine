export interface InterchainBatch {
  id: string
  batchRoot: string
  srcDbNonce: bigint
  srcChainId?: number
  dstDbNonce?: bigint
  dstChainId?: number
  status: string
  verifiedAt?: bigint
  appConfigId?: string
}

export interface InterchainTransaction {
  id: string
  srcChainId: number
  dstChainId: number
  srcSender: string
  dstReceiver: string
  sentAt?: bigint
  receivedAt?: bigint
  createdAt?: bigint
  updatedAt?: bigint
  interchainTransactionSentId?: string
  interchainTransactionReceivedId?: string
  interchainBatchId?: string
}
