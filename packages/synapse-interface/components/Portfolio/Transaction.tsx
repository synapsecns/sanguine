import { useState, useEffect, useMemo, useCallback } from 'react'
import Image from 'next/image'
import { waitForTransaction } from '@wagmi/core'
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
import { useAppDispatch } from '@/store/hooks'
import { updatePendingBridgeTransaction } from '@/slices/bridge/actions'
import { ARBITRUM, ETH } from '@/constants/chains/master'
import { USDC } from '@/constants/tokens/master'
import { getTimeMinutesFromNow } from '@/utils/time'
import TransactionArrow from '../icons/TransactionArrow'

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
  const handleExplorerClick: () => void = useCallback(() => {
    if (
      kappa &&
      originChain &&
      transactionType === TransactionType.HISTORICAL
    ) {
      const explorerLink: string = getTransactionExplorerLink({
        kappa,
        fromChainId: originChain.id,
        toChainId: destinationChain.id,
      })
      window.open(explorerLink, '_blank', 'noopener,noreferrer')
    }
  }, [kappa, originChain, destinationChain, transactionStatus, transactionHash])

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
      className="flex flex-col my-2 overflow-hidden border rounded-lg text-secondary border-surface"
    >
      <div
        onClick={handleExplorerClick}
        className={`flex flex-row ${
          transactionType === TransactionType.HISTORICAL && 'cursor-pointer'
        }`}
      >
        <div
          className={`
          flex items-center p-2
          ${transactionType === TransactionType.PENDING && 'bg-surface'}
          `}
        >
          <TransactionPayloadDetail
            chain={originChain}
            token={originToken}
            tokenAmount={originValue}
          />
        </div>
        <TransactionArrow
          className={`
          ${
            transactionType === TransactionType.PENDING
              ? 'bg-tint fill-surface'
              : 'stroke-surface'
          }
          `}
        />
        <div
          className={`
          flex flex-row justify-between flex-1
          ${transactionType === TransactionType.PENDING && 'bg-tint'}
          `}
        >
          <div className="flex items-center p-2">
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
                transactionStatus={transactionStatus}
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
  eventType?: number
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
  eventType,
  kappa,
  isSubmitted,
  isCompleted = false,
  transactionType = TransactionType.PENDING,
}: PendingTransactionProps) => {
  const dispatch = useAppDispatch()

  const transactionStatus: TransactionStatus = useMemo(() => {
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
    // CCTP Classification
    if (originChain.id === ARBITRUM.id || originChain.id === ETH.id) {
      const isCCTP: boolean =
        originToken.addresses[originChain.id] === USDC.addresses[originChain.id]
      if (eventType === 10 || eventType === 11 || isCCTP) {
        console.log('in here')
        const attestationTime: number = 13 * 60
        return (
          (BRIDGE_REQUIRED_CONFIRMATIONS[originChain.id] *
            originChain.blockTime) /
            1000 +
          attestationTime
        )
      }
    }
    // All other transactions
    console.log('in here 1')
    return originChain
      ? (BRIDGE_REQUIRED_CONFIRMATIONS[originChain.id] *
          originChain.blockTime) /
          1000
      : null
  }, [originChain, eventType, originToken])

  console.log('estimatedCompletionInSeconds:', estimatedCompletionInSeconds)

  useEffect(() => {
    if (!isSubmitted && transactionHash) {
      const updateResolvedTransaction = async () => {
        const resolvedTransaction = await waitForTransaction({
          hash: transactionHash as Address,
        })
        if (resolvedTransaction) {
          const currentTimestamp: number = getTimeMinutesFromNow(0)
          const updatedTransaction = {
            id: startedTimestamp,
            timestamp: currentTimestamp,
            transactionHash: transactionHash,
            isSubmitted: true,
          }

          await dispatch(updatePendingBridgeTransaction(updatedTransaction))
        }
      }
      updateResolvedTransaction()
    }
  }, [startedTimestamp, isSubmitted, transactionHash])

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
        transactionStatus={transactionStatus}
        isCompleted={isCompleted}
      >
        <TransactionStatusDetails
          originChain={originChain}
          destinationChain={destinationChain}
          kappa={kappa}
          transactionHash={transactionHash}
          transactionStatus={transactionStatus}
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
    'flex justify-between bg-tint border-t border-surface text-sm items-center'

  if (transactionStatus === TransactionStatus.PENDING_WALLET_ACTION) {
    return (
      <div
        data-test-id="pending-wallet-action-status"
        className={`${sharedClass} p-3`}
      >
        <div>Wallet signature required</div>
        <div>Check wallet</div>
      </div>
    )
  }

  if (transactionStatus === TransactionStatus.INITIALIZING) {
    return (
      <div data-test-id="initializing-status" className={`${sharedClass} p-3`}>
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
      <div data-test-id="pending-status" className={`${sharedClass} p-2`}>
        <div
          className="flex cursor-pointer hover:bg-[#101018] rounded-md"
          onClick={handleExplorerClick}
        >
          <Image
            className="w-4 h-4 my-auto mr-1.5 rounded-full"
            src={originChain.explorerImg}
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
    const handleExplorerClick = () => {
      const explorerLink: string = getTransactionExplorerLink({
        kappa,
        fromChainId: originChain.id,
        toChainId: destinationChain.id,
      })
      window.open(explorerLink, '_blank', 'noopener,noreferrer')
    }
    return (
      <div data-test-id="completed-status" className={`${sharedClass} p-2`}>
        <div
          className="flex cursor-pointer hover:bg-[#101018] rounded-md p-1"
          onClick={handleExplorerClick}
        >
          <div>Confirmed on Synapse Explorer</div>
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
}
