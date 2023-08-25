import { useState, useEffect, useMemo, useCallback } from 'react'
import Image from 'next/image'
import { Chain, Token } from '@/utils/types'
import {
  TransactionPayloadDetail,
  Completed,
  EstimatedDuration,
} from './Activity'
import { Address } from 'viem'
import { BRIDGE_REQUIRED_CONFIRMATIONS } from '@/constants/bridge'
import { TransactionOptions } from './TransactionOptions'
import { getTransactionExplorerLink } from './Activity'
import { getExplorerTxUrl } from '@/constants/urls'

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
          flex border-r border-r-[#252537] p-2
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
          <div className="p-2">
            <TransactionPayloadDetail
              chain={destinationChain}
              token={destinationToken}
              tokenAmount={destinationValue}
            />
          </div>
          <div className="p-3">
            {!isCompleted && transactionType === TransactionType.PENDING ? (
              <EstimatedDuration
                startTime={startedTimestamp}
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
          transactionHash={transactionHash}
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
  transactionHash,
  transactionStatus,
}: {
  originChain: Chain
  destinationChain: Chain
  kappa?: string
  transactionHash?: string
  transactionStatus: TransactionStatus
}) => {
  const sharedClass: string =
    'flex justify-between bg-[#1B1B29] border-t border-[#252537] p-2 text-sm items-center'

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
    const handleExplorerClick = () => {
      const explorerLink: string = getExplorerTxUrl({
        chainId: originChain.id,
        hash: transactionHash,
      })
      window.open(explorerLink, '_blank', 'noopener,noreferrer')
    }
    return (
      <div data-test-id="pending-status" className={sharedClass}>
        <div
          className="flex cursor-pointer hover:bg-[#101018] p-1 rounded-md"
          onClick={handleExplorerClick}
        >
          <Image
            className="w-4 h-4 my-auto mr-1.5 rounded-full"
            src={originChain.chainImg}
            alt={`${originChain.explorerName} logo`}
          />
          <div>Sent: {originChain.explorerName}</div>
        </div>
        <TransactionOptions
          originChain={originChain}
          destinationChain={destinationChain}
          kappa={kappa}
          transactionHash={transactionHash}
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
          destinationChain={destinationChain}
          kappa={kappa}
          transactionHash={transactionHash}
          transactionStatus={transactionStatus}
        />
      </div>
    )
  }
}
