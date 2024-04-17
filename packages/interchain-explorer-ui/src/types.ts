export type InterchainTransactionSent = {
  id: string
  chainId: number
  address: string
  srcSender: string
  dstChainId: number
  dstReceiver: string
  transactionHash: string
  options: string
  timestamp: number
  dbNonce: bigint
  count: number
}

export type InterchainTransactionReceived = {
  id: string
  chainId: number
  address: string
  srcSender: string
  srcChainId: string
  dstReceiver: string
  transactionHash: string
  timestamp: number
  dbNonce: bigint
  count: number
}

export type InterchainTransaction = {
  id: string
  status: string
  interchainTransactionSent: InterchainTransactionSent
  interchainTransactionReceived: InterchainTransactionReceived
}

export type PageInfo = {
  endCursor: string | null
  startCursor: string | null
  hasNextPage: boolean
  hasPreviousPage: boolean
}
