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
  transactionStatus?: TransactionStatus
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
    <div
      data-test-id="transaction"
      className="flex flex-row p-3 text-[#C2C2D6]"
    >
      <div className="flex border-r border-r-[#252537] px-2">
        <TransactionPayloadDetail
          chain={originChain}
          token={originToken}
          tokenSymbol={originToken?.symbol}
          tokenAmount={originValue}
        />
      </div>
      <div className="px-2">
        <TransactionPayloadDetail
          chain={destinationChain}
          token={destinationToken}
          tokenSymbol={destinationToken?.symbol}
          tokenAmount={destinationValue}
        />
      </div>
      <div className="ml-auto">
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
