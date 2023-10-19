import { useMemo, useEffect, useState } from 'react'
import Image from 'next/image'
import { waitForTransaction, Address } from '@wagmi/core'
import { useAppDispatch } from '@/store/hooks'
import { updatePendingBridgeTransaction } from '@/slices/bridge/actions'
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
import { BridgeType } from '@/slices/api/generated'

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

  const currentTime: number = Math.floor(Date.now() / 1000)
  console.log('currentTime:', currentTime)

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

  // console.log('elapsedTime:', elapsedTime)
  // console.log('currentTime:', currentTime)
  // console.log('estimatedCompletionInSeconds: ', estimatedCompletionInSeconds)

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
    if (!startedTimestamp || !elapsedTime) {
      return estimatedMinutes
    } else {
      return estimatedMinutes - elapsedTime
    }
  }, [estimatedMinutes, elapsedTime, startedTimestamp])

  // console.log('timeRemaining: ', timeRemaining)

  const isDelayed: boolean = useMemo(() => timeRemaining < 0, [timeRemaining])

  // // testing origin fallback query
  const originFallback = useFallbackBridgeOriginQuery({
    useFallback: true,
    chainId: originChain?.id,
    txnHash: transactionHash,
    bridgeType: BridgeType.Bridge,
  })

  // //testing dest fallback query
  // const destinationFallback = useFallbackBridgeDestinationQuery({
  //   chainId: destinationChain?.id,
  //   address: destinationAddress,
  //   kappa: kappa,
  //   timestamp: startedTimestamp,
  //   bridgeType: BridgeType.Bridge,
  //   useFallback: true,
  // })

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
        timeRemaining={timeRemaining}
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
                src={originChain.explorerImg}
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
