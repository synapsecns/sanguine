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
  kappa?: string
  children?: React.ReactNode
  isCompleted?: boolean
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
  kappa,
  children,
  isCompleted,
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
      className="flex flex-col my-2 text-[#C2C2D6]
      border border-[#222235] rounded-lg overflow-hidden"
    >
      <div className="flex flex-row">
        <div
          className={`
          flex border-r border-r-[#252537] p-3 min-w-[125px]
          ${transactionType === TransactionType.PENDING && 'bg-[#27273B]'}
          `}
        >
          <TransactionPayloadDetail
            chain={originChain}
            token={originToken}
            tokenAmount={originValue}
          />
        </div>
        <div
          className={`
          flex flex-row justify-between flex-1
          ${transactionType === TransactionType.PENDING && 'bg-[#1B1B29]'}
          `}
        >
          <div className="p-3">
            <TransactionPayloadDetail
              chain={destinationChain}
              token={destinationToken}
              tokenAmount={destinationValue}
            />
          </div>
          <div className="p-3">
            {!isCompleted && transactionType === TransactionType.PENDING ? (
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
      </div>
      {children}
    </div>
  )
}

interface PendingTransactionProps extends TransactionProps {
  isSubmitted?: boolean
  isCompleted?: boolean
}

export const PendingTransaction = ({
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
  transactionHash,
  kappa,
  isSubmitted,
  isCompleted = false,
  transactionType = TransactionType.PENDING,
}: PendingTransactionProps) => {
  const currentStatus: TransactionStatus = useMemo(() => {
    if (!transactionHash && !isSubmitted) {
      return TransactionStatus.PENDING_WALLET_ACTION
    }
    if (transactionHash && !isSubmitted) {
      return TransactionStatus.INITIALIZING
    }
    if (transactionHash && isSubmitted && !isCompleted) {
      return TransactionStatus.PENDING
    }
    if (isCompleted) {
      return TransactionStatus.COMPLETED
    }
  }, [transactionHash, isSubmitted, isCompleted])

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
        destinationAddress={destinationAddress}
        originChain={originChain}
        originToken={originToken}
        originValue={originValue}
        destinationChain={destinationChain}
        destinationToken={destinationToken}
        destinationValue={destinationValue}
        startedTimestamp={startedTimestamp}
        completedTimestamp={completedTimestamp}
        transactionType={TransactionType.PENDING}
        estimatedDuration={estimatedCompletionInSeconds}
        isCompleted={isCompleted}
      >
        <TransactionStatusDetails
          originChain={originChain}
          destinationChain={destinationChain}
          kappa={kappa}
          transactionStatus={currentStatus}
        />
      </Transaction>
    </div>
  )
}

const TransactionStatusDetails = ({
  originChain,
  destinationChain,
  kappa,
  transactionStatus,
}: {
  originChain: Chain
  destinationChain: Chain
  kappa?: string
  transactionStatus: TransactionStatus
}) => {
  const sharedClass: string =
    'flex justify-between bg-[#1B1B29] border-t border-[#252537] px-3 py-2 text-sm items-center'
  if (transactionStatus === TransactionStatus.PENDING_WALLET_ACTION) {
    return (
      <div data-test-id="pending-wallet-action-status" className={sharedClass}>
        <div>Wallet signature required</div>
        <div>Check wallet</div>
      </div>
    )
  }

  if (transactionStatus === TransactionStatus.INITIALIZING) {
    return (
      <div data-test-id="initializing-status" className={sharedClass}>
        <div>Initializing...</div>
      </div>
    )
  }

  if (transactionStatus === TransactionStatus.PENDING) {
    return (
      <div data-test-id="pending-status" className={sharedClass}>
        <div>Sent: {originChain.explorerName} </div>
        <TransactionOptions
          originChain={originChain}
          transactionStatus={transactionStatus}
        />
      </div>
    )
  }

  if (transactionStatus === TransactionStatus.COMPLETED) {
    return (
      <div data-test-id="completed-status" className={sharedClass}>
        <div>Confirmed on Synapse Explorer</div>
        <TransactionOptions
          originChain={originChain}
          transactionStatus={transactionStatus}
        />
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
