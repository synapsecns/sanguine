import { useMemo, useEffect, useState } from 'react'
import Image from 'next/image'
import { waitForTransaction, Address } from '@wagmi/core'
import { useAppDispatch } from '@/store/hooks'
import {
  removePendingBridgeTransaction,
  updatePendingBridgeTransaction,
} from '@/slices/transactions/actions'
import { BridgeType } from '@/slices/api/generated'
import { getTimeMinutesFromNow } from '@/utils/time'
import { ARBITRUM, ETH } from '@/constants/chains/master'
import { USDC } from '@/constants/tokens/bridgeable'
import {
  Transaction,
  TransactionProps,
  TransactionType,
  TransactionStatus,
} from './Transaction'
import { BRIDGE_REQUIRED_CONFIRMATIONS } from '@/constants/bridge'
import { TransactionOptions } from './TransactionOptions'
import { getExplorerTxUrl, getExplorerAddressUrl } from '@/constants/urls'
import { getTransactionExplorerLink } from './components/TransactionExplorerLink'
import { Chain } from '@/utils/types'
import { useFallbackBridgeOriginQuery } from '@/utils/hooks/useFallbackBridgeOriginQuery'
import { useFallbackBridgeDestinationQuery } from '@/utils/hooks/useFallbackBridgeDestinationQuery'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'

interface PendingTransactionProps extends TransactionProps {
  eventType?: number
  formattedEventType?: string
  bridgeModuleName?: string
  isSubmitted: boolean
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
  estimatedDuration,
  startedTimestamp,
  completedTimestamp,
  transactionHash,
  eventType,
  formattedEventType,
  bridgeModuleName,
  kappa,
  isSubmitted,
  isCompleted = false,
  transactionType = TransactionType.PENDING,
}: PendingTransactionProps) => {
  const { synapseSDK } = useSynapseContext()
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
    // Fallback last resort estimated duration calculation
    // Remove this when fallback origin queries return eventType
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
  }, [originChain, eventType, originToken, bridgeModuleName, transactionHash])

  const currentTime: number = Math.floor(Date.now() / 1000)

  // Tracks initial elapsed minutes when transaction mounts to populate updatedElapsedTime
  const initialElapsedMinutes: number = useMemo(() => {
    if (!isSubmitted || currentTime < startedTimestamp) {
      return 0
    } else if (startedTimestamp < currentTime) {
      return Math.floor((currentTime - startedTimestamp) / 60)
    } else {
      return 0
    }
  }, [startedTimestamp, currentTime, isSubmitted, transactionHash])

  // Holds most updated value for elapsed time to calculate timeRemaining
  const [updatedElapsedTime, setUpdatedElapsedTime] = useState<number>(
    initialElapsedMinutes
  )

  // Ensures we reset elapsed time so unique transactions track elapsed time accurately
  useEffect(() => {
    if (!initialElapsedMinutes && updatedElapsedTime > initialElapsedMinutes) {
      setUpdatedElapsedTime(0)
    }
  }, [initialElapsedMinutes, updatedElapsedTime])

  // Update elapsed time in set intervals for countdown
  useEffect(() => {
    const interval = setInterval(() => {
      const currentTime: number = Math.floor(Date.now() / 1000)
      const elapsedMinutes: number = Math.floor(
        (currentTime - startedTimestamp) / 60
      )
      if (isSubmitted) {
        setUpdatedElapsedTime(elapsedMinutes)
      }
    }, 30000)

    return () => {
      clearInterval(interval)
    }
  }, [startedTimestamp, isSubmitted, transactionHash])

  const estimatedCompletionInMinutes: number = Math.floor(
    estimatedCompletionInSeconds / 60
  )

  const timeRemaining: number = useMemo(() => {
    if (!startedTimestamp || !updatedElapsedTime) {
      return estimatedCompletionInMinutes
    } else {
      return estimatedCompletionInMinutes - updatedElapsedTime
    }
  }, [
    estimatedCompletionInMinutes,
    initialElapsedMinutes,
    updatedElapsedTime,
    startedTimestamp,
  ])

  const isDelayed: boolean = useMemo(() => timeRemaining < 0, [timeRemaining])

  // Set fallback period to extend 5 mins past estimated duration
  const useFallback: boolean = useMemo(
    () => timeRemaining >= -5 && timeRemaining <= 1 && !isCompleted,
    [timeRemaining, isCompleted]
  )

  const bridgeType: BridgeType = useMemo(() => {
    if (synapseSDK && formattedEventType) {
      const moduleName: string =
        synapseSDK.getBridgeModuleName(formattedEventType)

      if (moduleName === 'SynapseBridge') return BridgeType.Bridge
      if (moduleName === 'SynapseCCTP') return BridgeType.Cctp
    }
    if (synapseSDK && bridgeModuleName) {
      if (bridgeModuleName === 'SynapseBridge') return BridgeType.Bridge
      if (bridgeModuleName === 'SynapseCCTP') return BridgeType.Cctp
    }
    return BridgeType.Bridge
  }, [synapseSDK, bridgeModuleName, formattedEventType])

  const originFallback = useFallbackBridgeOriginQuery({
    useFallback: isDelayed && useFallback,
    chainId: originChain?.id,
    txnHash: transactionHash,
    bridgeType: bridgeType,
  })

  const destinationFallback = useFallbackBridgeDestinationQuery({
    useFallback: isDelayed && useFallback,
    chainId: destinationChain?.id,
    address: destinationAddress,
    kappa: kappa,
    timestamp: startedTimestamp,
    bridgeType: bridgeType,
  })

  useEffect(() => {
    if (!isSubmitted && transactionHash) {
      const maxRetries = 3
      const retryDelay = 5000
      let attempt = 0

      const updateResolvedTransaction = async () => {
        try {
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

            console.log('resolved transaction:', resolvedTransaction)
            dispatch(updatePendingBridgeTransaction(updatedTransaction))
          }
        } catch (error) {
          console.error('resolving transaction failed: ', error)
          if (attempt < maxRetries) {
            attempt++
            console.log(`Retrying (${attempt}/${maxRetries})...`)
            setTimeout(updateResolvedTransaction, retryDelay)
          }
        }
      }

      updateResolvedTransaction()
    }
  }, [
    startedTimestamp,
    isSubmitted,
    transactionHash,
    dispatch,
    updatePendingBridgeTransaction,
    getTimeMinutesFromNow,
    updatedElapsedTime,
  ])

  useEffect(() => {
    const currentTimestamp: number = getTimeMinutesFromNow(0)
    const isStale: boolean =
      !transactionHash &&
      !isSubmitted &&
      currentTimestamp - startedTimestamp > 100

    if (!isSubmitted && isStale) {
      dispatch(removePendingBridgeTransaction(startedTimestamp))
    }
  }, [timeRemaining, isSubmitted, startedTimestamp])

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
        timeRemaining={
          isSubmitted ? timeRemaining : estimatedCompletionInMinutes
        }
        transactionStatus={transactionStatus}
        isCompleted={isCompleted}
        kappa={kappa}
      >
        <TransactionStatusDetails
          connectedAddress={connectedAddress}
          originChain={originChain}
          destinationChain={destinationChain}
          transactionHash={transactionHash}
          transactionStatus={transactionStatus}
          isDelayed={isDelayed}
          kappa={kappa}
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
        className={`${sharedClass} py-3 px-3 justify-between`}
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
        className={`${sharedClass} py-2 px-3 justify-between`}
      >
        <div>Initiating...</div>
        <TransactionOptions
          connectedAddress={connectedAddress as Address}
          originChain={originChain}
          destinationChain={destinationChain}
          kappa={kappa}
          transactionHash={transactionHash}
          transactionStatus={transactionStatus}
          isDelayed={false}
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
            <div className="flex items-center p-1 ml-1 rounded-sm cursor-default">
              <div className="text-[#FFDD33]">Taking longer than expected.</div>
            </div>
            <TransactionOptions
              connectedAddress={connectedAddress as Address}
              originChain={originChain}
              destinationChain={destinationChain}
              kappa={kappa}
              transactionHash={transactionHash}
              transactionStatus={transactionStatus}
              isDelayed={isDelayed}
            />
          </>
        ) : (
          <>
            <div
              className="flex cursor-pointer hover:bg-[#101018] rounded-sm hover:text-[#99E6FF] hover:underline py-1 px-1 items-center"
              onClick={handleOriginExplorerClick}
            >
              <Image
                className="w-4 h-4 mx-1 ml-1 mr-1.5 rounded-full"
                src={originChain?.explorerImg}
                alt={`${originChain.explorerName} logo`}
              />
              <div>Confirmed on {originChain.explorerName}.</div>
            </div>
            <div
              onClick={handleDestinationExplorerClick}
              className="mr-auto cursor-pointer hover:bg-[#101018] rounded-sm hover:text-[#99E6FF] hover:underline py-1 px-1"
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
              isDelayed={isDelayed}
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
          className="flex cursor-pointer hover:bg-[#101018] rounded-sm hover:text-[#99E6FF] hover:underline py-1 px-1"
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
          isDelayed={false}
        />
      </div>
    )
  }
}
