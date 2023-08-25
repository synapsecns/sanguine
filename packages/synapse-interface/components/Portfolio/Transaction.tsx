import { useState, useEffect, useMemo } from 'react'
import { Chain, Token } from '@/utils/types'
import {
  TransactionPayloadDetail,
  Completed,
  EstimatedDuration,
} from './Activity'
import { Address } from 'viem'
import { BRIDGE_REQUIRED_CONFIRMATIONS } from '@/constants/bridge'
import { TransactionOptions } from './TransactionOptions'

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
  const estimatedCompletionInSeconds: number = useMemo(() => {
    return originChain
      ? (BRIDGE_REQUIRED_CONFIRMATIONS[originChain.id] *
          originChain.blockTime) /
          1000
      : null
  }, [originChain])

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
          <EstimatedDuration
            estimatedCompletionInSeconds={
              estimatedDuration ?? estimatedCompletionInSeconds
            }
          />
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
  const currentStatus: TransactionStatus = useMemo(() => {
    if (!transactionHash && !isSubmitted)
      return TransactionStatus.PENDING_WALLET_ACTION
    if (transactionHash && !isSubmitted) return TransactionStatus.INITIALIZING
    if (transactionHash && isSubmitted) return TransactionStatus.PENDING
  }, [transactionHash, isSubmitted])

  const estimatedCompletionInSeconds: number = useMemo(() => {
    return originChain
      ? (BRIDGE_REQUIRED_CONFIRMATIONS[originChain.id] *
          originChain.blockTime) /
          1000
      : null
  }, [originChain])

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
        estimatedDuration={estimatedCompletionInSeconds}
      />
      <TransactionStatusDetails transactionStatus={currentStatus} />
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
      <div data-test-id="initializing-status" className="flex justify-between">
        <div>Initializing...</div>
      </div>
    )
  }

  if (transactionStatus === TransactionStatus.PENDING) {
    return (
      <div data-test-id="pending-status" className="flex justify-between">
        <div>Sent</div>
        <TransactionOptions />
      </div>
    )
  }

  if (transactionStatus === TransactionStatus.COMPLETED) {
    return (
      <div data-test-id="completed-status" className="flex justify-between">
        <div>Confirmed on Synapse Explorer</div>
      </div>
    )
  }
}

// const handlePendingTransactionClick: () => void = useCallback(() => {
//   if (transactionHash) {
//     const explorerLink: string = getExplorerTxUrl({
//       chainId: originChain.id,
//       hash: transactionHash,
//     })
//     window.open(explorerLink, '_blank')
//   }
// }, [transactionHash])

// const handleTransactionClick: () => void = useCallback(() => {
//   if (kappa && originChainId && transactionType === ActivityType.RECENT) {
//     const explorerLink: string = getTransactionExplorerLink({
//       kappa,
//       fromChainId: originChainId,
//       toChainId: destinationChainId,
//     })
//     window.open(explorerLink, '_blank')
//   } else {
//     const explorerLink: string = getExplorerTxUrl({
//       chainId: originChainId,
//       hash: originTxnHash,
//     })
//     window.open(explorerLink, '_blank')
//   }
// }, [kappa, originChainId, destinationChainId, transactionType])

// const estimatedCompletionInSeconds: number =
//   (BRIDGE_REQUIRED_CONFIRMATIONS[originChain.id] * originChain.blockTime) /
//     1000 +
//   30 // Add 30 seconds to account for indexing
