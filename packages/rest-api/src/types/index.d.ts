export interface BridgeableToken {
  addresses: {}
  decimals: {}
  symbol: string
  name: string
  swapableType: string
  color: string
  priorityRank: number
  routeSymbol: string
  imgUrl: string
}

export type CustomRpcs = {
  [chainId: number]: string
}

export interface Chain {
  id: number
  name: string
  rpcUrls: {
    primary: string
    fallback: string
  }
  explorerUrl: string
  explorerName: string
  blockTime: number
  imgUrl: string
  networkName: string
  networkUrl: string
  nativeCurrency: {
    name: string
    symbol: string
    decimals: number
  }
}

export interface BridgeTransaction {
  kappa: string
  fromInfo: BridgeTransactionInfo
  toInfo: BridgeTransactionInfo
}

export interface BridgeTransactionInfo {
  chainID: number
  address: string
  txnHash: string
  value: string
  USDValue: string
  tokenSymbol: string
  tokenAddress: string
  blockNumber: number
  formattedTime: string
}
