import { Chain, Token } from '@/utils/types'
import {
  TransactionPayloadDetail,
  Completed,
  EstimatedDuration,
} from './Activity'
import { Address } from 'viem'

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
  connectedAddress: Address
  destinationAddress: Address
  originChain: Chain
  originToken: Token
  originValue: number
  destinationChain: Chain
  destinationToken: Token
  destinationValue?: number
  startedTimestamp: number
  completedTimestamp?: number
  estimatedDuration?: number
  transactionStatus: TransactionStatus
  transactionType: TransactionType
}

export const Transaction = ({
  connectedAddress,
  destinationAddress,
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
}: TransactionProps) => {
  return (
    <div data-test-id="transaction" className="flex flex-row">
      <div className="flex">
        <TransactionPayloadDetail
          chain={originChain}
          token={originToken}
          tokenSymbol={originToken?.symbol}
          tokenAmount={originValue}
        />
      </div>
      <div>
        <TransactionPayloadDetail
          chain={destinationChain}
          token={destinationToken}
          tokenSymbol={destinationToken?.symbol}
          tokenAmount={destinationValue}
        />
      </div>
      <div>
        {transactionType === TransactionType.PENDING ? (
          <EstimatedDuration estimatedCompletionInSeconds={estimatedDuration} />
        ) : (
          <Completed
            transactionCompletedTime={completedTimestamp}
            connectedAddress={connectedAddress}
            destinationAddress={destinationAddress}
          />
        )}
      </div>
    </div>
  )
}
