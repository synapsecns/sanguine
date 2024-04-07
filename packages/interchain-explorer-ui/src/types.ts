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
}

export type InterchainTransaction = {
  id: string
  interchainTransactionSent: InterchainTransactionSent
  interchainTransactionReceived: InterchainTransactionReceived
}
