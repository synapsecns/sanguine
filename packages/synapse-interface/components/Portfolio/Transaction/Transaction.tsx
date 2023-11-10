import { useCallback } from 'react'
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
  timeRemaining,
  transactionStatus,
  transactionType,
  transactionHash,
  kappa,
  children,
  isCompleted,
}: TransactionProps) => {
  const handleExplorerClick: () => void = useCallback(() => {
    if (kappa && originChain && destinationChain) {
      const explorerLink: string = getTransactionExplorerLink({
        kappa,
        fromChainId: originChain.id,
        toChainId: destinationChain.id,
      })
      window.open(explorerLink, '_blank', 'noopener,noreferrer')
    }
  }, [kappa, originChain, destinationChain, transactionStatus, transactionHash])

  return (
    <div
      data-test-id="transaction"
      className={`
        flex flex-col mb-2 overflow-hidden
        border rounded-lg text-secondary border-surface
        ${transactionType === TransactionType.HISTORICAL && 'bg-background'}
      `}
    >
      <div className={`flex flex-row`}>
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
            isOrigin={true}
          />
        </div>
        <TransactionArrow
          className={`
          ${
            transactionType === TransactionType.PENDING
              ? 'bg-tint fill-surface'
              : 'stroke-surface fill-transparent'
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
              isOrigin={false}
            />
          </div>
          <div className="p-3">
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
                handleExplorerClick={handleExplorerClick}
              />
            )}
          </div>
        </div>
      </div>
      {children}
    </div>
  )
}
