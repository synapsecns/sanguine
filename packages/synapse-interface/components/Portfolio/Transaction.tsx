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
import { getExplorerTxUrl, getExplorerAddressUrl } from '@/constants/urls'
import { useAppDispatch } from '@/store/hooks'
import { updatePendingBridgeTransaction } from '@/slices/bridge/actions'
import { ARBITRUM, ETH } from '@/constants/chains/master'
import { USDC } from '@/constants/tokens/bridgeable'
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
  const [elapsedTime, setElapsedTime] = useState<number>(0)

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

  const estimatedCompletionInSeconds: number = useMemo(() => {
    return originChain
      ? (BRIDGE_REQUIRED_CONFIRMATIONS[originChain.id] *
          originChain.blockTime) /
          1000
      : null
  }, [originChain])

  useEffect(() => {
    const currentTime: number = Math.floor(Date.now() / 1000)
    const elapsedMinutes: number = Math.floor(
      (currentTime - startedTimestamp) / 60
    )
    setElapsedTime(elapsedMinutes)
  }, [startedTimestamp])

  useEffect(() => {
    const interval = setInterval(() => {
      const currentTime: number = Math.floor(Date.now() / 1000)
      const elapsedMinutes: number = Math.floor(
        (currentTime - startedTimestamp) / 60
      )
      setElapsedTime(elapsedMinutes)
    }, 60000)

    return () => {
      clearInterval(interval)
    }
  }, [startedTimestamp])

  const estimatedMinutes: number = Math.floor(estimatedCompletionInSeconds / 60)
  const timeRemaining: number = useMemo(() => {
    if (!startedTimestamp || !elapsedTime) return estimatedMinutes
    return estimatedMinutes - elapsedTime
  }, [estimatedMinutes, elapsedTime, startedTimestamp])

  return (
    <div
      data-test-id="transaction"
      className={`
        flex flex-col my-2 overflow-hidden
        border rounded-lg text-secondary border-surface
        ${transactionType === TransactionType.HISTORICAL && 'bg-[#18151A]'}
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
      if ((eventType === 10 || eventType === 11) && isCCTP) {
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
    return originChain
      ? (BRIDGE_REQUIRED_CONFIRMATIONS[originChain.id] *
          originChain.blockTime) /
          1000
      : null
  }, [originChain, eventType, originToken])

  const [elapsedTime, setElapsedTime] = useState<number>(0)

  useEffect(() => {
    const currentTime: number = Math.floor(Date.now() / 1000)
    const elapsedMinutes: number = Math.floor(
      (currentTime - startedTimestamp) / 60
    )
    setElapsedTime(elapsedMinutes)
  }, [startedTimestamp])

  useEffect(() => {
    const interval = setInterval(() => {
      const currentTime: number = Math.floor(Date.now() / 1000)
      const elapsedMinutes: number = Math.floor(
        (currentTime - startedTimestamp) / 60
      )
      setElapsedTime(elapsedMinutes)
    }, 60000)

    return () => {
      clearInterval(interval)
    }
  }, [startedTimestamp])

  const estimatedMinutes: number = Math.floor(estimatedCompletionInSeconds / 60)
  const timeRemaining: number = useMemo(() => {
    if (!startedTimestamp || !elapsedTime) return estimatedMinutes
    return estimatedMinutes - elapsedTime
  }, [estimatedMinutes, elapsedTime, startedTimestamp])

  const isDelayed: boolean = useMemo(() => timeRemaining < 0, [timeRemaining])

  useEffect(() => {
    if (!isSubmitted && transactionHash) {
      const updateResolvedTransaction = async () => {
        const resolvedTransaction = await waitForTransaction({
          hash: transactionHash as Address,
        })

        console.log('resolvedTransaction:', resolvedTransaction)

        if (resolvedTransaction) {
          const currentTimestamp: number = getTimeMinutesFromNow(0)
          const updatedTransaction = {
            id: startedTimestamp,
            timestamp: currentTimestamp,
            transactionHash: transactionHash,
            isSubmitted: true,
          }

          dispatch(updatePendingBridgeTransaction(updatedTransaction))
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
        kappa={kappa}
      >
        <TransactionStatusDetails
          connectedAddress={connectedAddress}
          originChain={originChain}
          destinationChain={destinationChain}
          kappa={kappa}
          transactionHash={transactionHash}
          transactionStatus={transactionStatus}
          isDelayed={isDelayed}
        />
      </Transaction>
    </div>
  )
}

const TransactionStatusDetails = ({
  connectedAddress,
  originChain,
  destinationChain,
  kappa,
  transactionHash,
  transactionStatus,
  isDelayed,
}: {
  connectedAddress: Address
  originChain: Chain
  destinationChain: Chain
  kappa?: string
  transactionHash?: string
  transactionStatus: TransactionStatus
  isDelayed: boolean
}) => {
  const sharedClass: string =
    'flex bg-tint border-t border-surface text-sm items-center'

  if (transactionStatus === TransactionStatus.PENDING_WALLET_ACTION) {
    return (
      <div
        data-test-id="pending-wallet-action-status"
        className={`${sharedClass} p-3 justify-between`}
      >
        <div>Wallet signature required</div>
        <div>Check wallet</div>
      </div>
    )
  }

  if (transactionStatus === TransactionStatus.INITIALIZING) {
    return (
      <div
        data-test-id="initializing-status"
        className={`${sharedClass} p-2 justify-between`}
      >
        <div>Initiating...</div>
        <TransactionOptions
          connectedAddress={connectedAddress as Address}
          originChain={originChain}
          destinationChain={destinationChain}
          kappa={kappa}
          transactionHash={transactionHash}
          transactionStatus={transactionStatus}
        />
      </div>
    )
  }

  if (transactionStatus === TransactionStatus.PENDING) {
    const handleOriginExplorerClick = () => {
      const explorerLink: string = getExplorerTxUrl({
        chainId: originChain.id,
        hash: transactionHash,
      })
      window.open(explorerLink, '_blank', 'noopener,noreferrer')
    }

    const handleDestinationExplorerClick = () => {
      const explorerLink: string = getExplorerAddressUrl({
        chainId: destinationChain.id,
        address: connectedAddress as Address,
      })
      window.open(explorerLink, '_blank', 'noopener,noreferrer')
    }

    return (
      <div
        data-test-id="pending-status"
        className={`${sharedClass} p-2 flex justify-between`}
      >
        {isDelayed ? (
          <>
            <div
              className="flex cursor-pointer hover:bg-[#101018] rounded-sm hover:text-[#FFDD33] hover:underline p-1 items-center"
              onClick={handleDestinationExplorerClick}
            >
              <Image
                className="w-4 h-4 my-auto mr-1.5 rounded-full"
                src={destinationChain.explorerImg}
                alt={`${destinationChain.explorerName} logo`}
              />
              <div className="text-[#FFDD33]">Taking longer than expected.</div>
            </div>
            <TransactionOptions
              connectedAddress={connectedAddress as Address}
              originChain={originChain}
              destinationChain={destinationChain}
              kappa={kappa}
              transactionHash={transactionHash}
              transactionStatus={transactionStatus}
            />
          </>
        ) : (
          <>
            <div
              className="flex cursor-pointer hover:bg-[#101018] rounded-sm hover:text-[#99E6FF] hover:underline p-1 items-center"
              onClick={handleOriginExplorerClick}
            >
              <Image
                className="w-4 h-4 my-auto mr-1.5 rounded-full"
                src={originChain.explorerImg}
                alt={`${originChain.explorerName} logo`}
              />
              <div>Confirmed on {originChain.explorerName}.</div>
            </div>
            <div
              onClick={handleDestinationExplorerClick}
              className="mr-auto cursor-pointer hover:bg-[#101018] rounded-sm hover:text-[#99E6FF] hover:underline p-1"
            >
              Bridging to {destinationChain.name}.
            </div>
            <TransactionOptions
              connectedAddress={connectedAddress as Address}
              originChain={originChain}
              destinationChain={destinationChain}
              kappa={kappa}
              transactionHash={transactionHash}
              transactionStatus={transactionStatus}
            />
          </>
        )}
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
      <div
        data-test-id="completed-status"
        className={`${sharedClass} p-2 justify-between`}
      >
        <div
          className="flex cursor-pointer hover:bg-[#101018] rounded-md hover:text-[#99E6FF] hover:underline p-1"
          onClick={handleExplorerClick}
        >
          <div>Confirmed on Synapse Explorer</div>
        </div>
        <TransactionOptions
          connectedAddress={connectedAddress as Address}
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
