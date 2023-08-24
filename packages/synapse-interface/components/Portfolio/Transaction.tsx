import { useState, useEffect } from 'react'
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
  destinationAddress?: Address
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
  transactionHash?: string
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
  transactionHash,
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
          tokenAmount={originValue}
        />
      </div>
      <div className="px-2">
        <TransactionPayloadDetail
          chain={destinationChain}
          token={destinationToken}
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

interface PendingTransactionProps extends TransactionProps {
  isSubmitted?: boolean
  isCompleted?: boolean
}

export const PendingTransaction = ({
  connectedAddress,
  originChain,
  originToken,
  originValue,
  destinationChain,
  destinationToken,
  startedTimestamp,
  transactionHash,
  isSubmitted,
  isCompleted = false,
  transactionType = TransactionType.PENDING,
}: PendingTransactionProps) => {
  const [status, setStatus] = useState<TransactionStatus>(
    transactionHash ? TransactionStatus.PENDING : TransactionStatus.INITIALIZING
  )

  const isPendingWalletAction: boolean = transactionHash ? true : false
  const isInitializing: boolean = isSubmitted ? false : true

  useEffect(() => {
    if (isPendingWalletAction)
      setStatus(TransactionStatus.PENDING_WALLET_ACTION)
    else if (isInitializing) setStatus(TransactionStatus.INITIALIZING)
    else if (transactionHash) setStatus(TransactionStatus.PENDING)
    else if (isCompleted) setStatus(TransactionStatus.COMPLETED)
  }, [isPendingWalletAction, isInitializing, transactionHash, isCompleted])

  return (
    <div data-test-id="pending-transaction" className="flex flex-col">
      <Transaction
        connectedAddress={connectedAddress}
        originChain={originChain}
        originToken={originToken}
        originValue={originValue}
        destinationChain={destinationChain}
        destinationToken={destinationToken}
        startedTimestamp={startedTimestamp}
        transactionType={TransactionType.PENDING}
      />
      <TransactionStatusDetails transactionStatus={status} />
    </div>
  )
}

const TransactionStatusDetails = ({
  transactionStatus,
}: {
  transactionStatus: TransactionStatus
}) => {
  if (transactionStatus === TransactionStatus.PENDING_WALLET_ACTION) {
    return (
      <div
        data-test-id="pending-wallet-action-status"
        className="flex justify-between"
      >
        <div>Wallet signature required</div>
        <button>Open wallet</button>
      </div>
    )
  }

  if (transactionStatus === TransactionStatus.INITIALIZING) {
    return (
      <div data-test-id="initializing-status">
        <div>Initializing...</div>
      </div>
    )
  }

  if (transactionStatus === TransactionStatus.PENDING) {
    return (
      <div data-test-id="pending-status">
        <div>Sent</div>
      </div>
    )
  }

  if (transactionStatus === TransactionStatus.COMPLETED) {
    return (
      <div data-test-id="completed-status">
        <div>Confirmed on Synapse Explorer</div>
      </div>
    )
  }
}
