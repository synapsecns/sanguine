import { ColumnType } from 'kysely'

// Define types for each table in the database
export interface BridgeRequestEvents {
  id: ColumnType<string>
  transactionId: ColumnType<string>
  blockNumber: ColumnType<bigint>
  blockTimestamp: ColumnType<number>
  transactionHash: ColumnType<string>
  originChainId: ColumnType<number>
  originChain: ColumnType<string>
  sender: ColumnType<string>
  request: ColumnType<string>
  originToken: ColumnType<string>
  destToken: ColumnType<string>
  originAmount: ColumnType<bigint>
  originAmountFormatted: ColumnType<string>
  destAmount: ColumnType<bigint>
  destAmountFormatted: ColumnType<string>
  destChainId: ColumnType<number>
  destChain: ColumnType<string>
  sendChainGas: ColumnType<boolean>
  deadline: ColumnType<number>
  exclusivityRelayer: ColumnType<string>
  exclusivityEndTime: ColumnType<number>
  zapNative: ColumnType<bigint>
  zapData: ColumnType<string>
}

export interface BridgeRelayedEvents {
  id: ColumnType<string>
  transactionId: ColumnType<string>
  blockNumber: ColumnType<bigint>
  blockTimestamp: ColumnType<number>
  transactionHash: ColumnType<string>
  originChainId: ColumnType<number>
  originChain: ColumnType<string>
  relayer: ColumnType<string>
  to: ColumnType<string>
  originToken: ColumnType<string>
  destToken: ColumnType<string>
  originAmount: ColumnType<bigint>
  originAmountFormatted: ColumnType<string>
  destAmount: ColumnType<bigint>
  destAmountFormatted: ColumnType<string>
  destChainId: ColumnType<number>
  destChain: ColumnType<string>
}

export interface BridgeProofProvidedEvents {
  id: ColumnType<string>
  transactionId: ColumnType<string>
  blockNumber: ColumnType<bigint>
  blockTimestamp: ColumnType<number>
  transactionHash: ColumnType<string>
  originChainId: ColumnType<number>
  originChain: ColumnType<string>
  relayer: ColumnType<string>
}

export interface BridgeDepositRefundedEvents {
  id: ColumnType<string>
  transactionId: ColumnType<string>
  blockNumber: ColumnType<bigint>
  blockTimestamp: ColumnType<number>
  transactionHash: ColumnType<string>
  originChainId: ColumnType<number>
  originChain: ColumnType<string>
  to: ColumnType<string>
  token: ColumnType<string>
  amount: ColumnType<bigint>
  amountFormatted: ColumnType<string>
}

export interface BridgeDepositClaimedEvents {
  id: ColumnType<string>
  transactionId: ColumnType<string>
  blockNumber: ColumnType<bigint>
  blockTimestamp: ColumnType<number>
  transactionHash: ColumnType<string>
  originChainId: ColumnType<number>
  originChain: ColumnType<string>
  relayer: ColumnType<string>
  to: ColumnType<string>
  token: ColumnType<string>
  amount: ColumnType<bigint>
  amountFormatted: ColumnType<string>
}

export interface BridgeProofDisputedEvents {
  id: ColumnType<string>
  transactionId: ColumnType<string>
  blockNumber: ColumnType<bigint>
  blockTimestamp: ColumnType<number>
  transactionHash: ColumnType<string>
  chainId: ColumnType<number>
  chain: ColumnType<string>
  originChainId: ColumnType<number>
  originChain: ColumnType<string>
}
// Add any other shared types used across the API
export type EventType =
  | 'REQUEST'
  | 'RELAYED'
  | 'PROOF_PROVIDED'
  | 'DEPOSIT_REFUNDED'
  | 'DEPOSIT_CLAIMED'
  | 'DISPUTE'
export interface EventFilter {
  type?: EventType
  transactionId?: string
  originChainId?: number
  destChainId?: number
}
