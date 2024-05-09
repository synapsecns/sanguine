export interface InterchainBatch {
  id: string
  batchRoot: string
  srcDbNonce: bigint
  srcChainId?: number
  dstDbNonce?: bigint
  dstChainId?: number
  status: string
  verifiedAt?: bigint
  appConfigV1Id?: string
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

export interface InterchainBatchQueryFilter {
  srcChainId?: number
  dstChainId?: number
  status?: string
}

export interface InterchainTransactionQueryFilter {
  srcChainId?: number
  dstChainId?: number
  status?: string
  dstReceiver?: string
  id?: string
}
