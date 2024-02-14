import React from 'react'
import { Chain, Token } from '@/utils/types'
import { Address } from 'viem'
import { getTransactionExplorerLink } from './components/TransactionExplorerLink'
import { TransactionPayloadDetail } from './components/TransactionPayloadDetail'
import { Completed } from './components/Completed'
import { EstimatedDuration } from './components/EstimatedDuration'
import TransactionArrow from '../../icons/TransactionArrow'

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
  originValue: string | number
  destinationChain: Chain
  destinationToken: Token
  destinationValue?: string | number
  startedTimestamp: number
  completedTimestamp?: number
  estimatedDuration?: number
  timeRemaining?: number
  transactionStatus?: TransactionStatus
  transactionType: TransactionType
  transactionHash?: string
  kappa?: string
  children?: React.ReactNode
  isCompleted?: boolean
}

export const Transaction = React.memo(
  ({
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
    timeRemaining,
    transactionStatus,
    transactionType,
    transactionHash,
    kappa,
    children,
    isCompleted,
  }: TransactionProps) => {

    const explorerLink: string = getTransactionExplorerLink({
      kappa,
      fromChainId: originChain.id,
      toChainId: destinationChain.id,
    })

    return (
      <div
        data-test-id="transaction"
        className={`
          flex my-2 rounded-md text-secondary border border-white/20
          ${transactionType === TransactionType.HISTORICAL && 'bg-bgBase/10'}
        `}
      >
        <TransactionPayloadDetail
          chain={originChain}
          token={originToken}
          tokenAmount={originValue}
          isOrigin={true}
          className="p-2 bg-bgBase/10 rounded-l"
        />
        <TransactionArrow
          className={
            transactionType === TransactionType.PENDING
              ? 'bg-bgBase/20 fill-bgBase/10'
              : 'stroke-white/20 fill-bgBase/10 '
          }
        />
        <TransactionPayloadDetail
          chain={destinationChain}
          token={destinationToken}
          tokenAmount={destinationValue}
          isOrigin={false}
          className="p-2"
        />
        <div className="ml-auto h-full py-2.5 -my-px px-4">
          {!isCompleted && transactionType === TransactionType.PENDING ? (
            <EstimatedDuration
              timeRemaining={timeRemaining}
              transactionStatus={transactionStatus}
            />
          ) : (
            <Completed
              transactionCompletedTime={completedTimestamp}
              connectedAddress={connectedAddress}
              destinationAddress={destinationAddress}
              explorerLink={
                (kappa && originChain && destinationChain)
                  ? explorerLink
                  : '#'
              }
            />
          )}
        </div>
        {children}
      </div>
    )
  }
)
