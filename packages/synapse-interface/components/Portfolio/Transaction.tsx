import { Chain, Token } from '@/utils/types'

export enum TransactionType {
  PENDING,
  HISTORICAL,
}

export enum TransactionStatus {
  PENDING_WALLET_ACTION,
  INITIALIZING,
  PENDING,
  COMPLETED,
}

export interface TransactionProps {
  originChain: Chain
  originToken: Token
  originValue: number
  destinationChain: Chain
  destinationToken: Token
  destinationValue?: number
  startedTimestamp: number
  completedTimestamp?: number
  estimatedDuration: number
  transactionStatus: TransactionStatus
  transactionType: TransactionType
}

export const Transaction = ({
  originChain,
  originToken,
  originValue,
  destinationChain,
  destinationToken,
  destinationValue,
  startedTimestamp,
  completedTimestamp,
  estimatedDuration,
  transactionStatus,
  transactionType,
}: TransactionProps) => {}
