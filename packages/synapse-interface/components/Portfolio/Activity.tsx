import React, { useEffect, useMemo, useState } from 'react'
import { useAccount, Address } from 'wagmi'
import Link from 'next/link'
import Image from 'next/image'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { BridgeTransaction } from '@/slices/api/generated'
import {
  convertUnixTimestampToMonthAndDate,
  getTimeMinutesBeforeNow,
  isTimestampToday,
} from '@/utils/time'
import { CHAINS_BY_ID } from '@/constants/chains'
import { Chain, Token } from '@/utils/types'
import { tokenAddressToToken } from '@/constants/tokens'
import { ANALYTICS_KAPPA, ANALYTICS_PATH } from '@/constants/urls'
import { TransactionsState } from '@/slices/transactions/reducer'
import { PendingBridgeTransaction } from '@/slices/bridge/actions'
import { BridgeState } from '@/slices/bridge/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { shortenAddress } from '@/utils/shortenAddress'
import {
  Transaction,
  PendingTransaction,
  TransactionType,
  TransactionStatus,
} from './Transaction'
import ProcessingIcon from '../icons/ProcessingIcon'

export const Activity = ({ visibility }: { visibility: boolean }) => {
  const { address } = useAccount()
  const {
    userHistoricalTransactions,
    userPendingTransactions,
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
    pendingAwaitingCompletionTransactions,
  }: TransactionsState = useTransactionsState()
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()

  const hasPendingTransactions: boolean = useMemo(() => {
    if (
      pendingAwaitingCompletionTransactions &&
      pendingAwaitingCompletionTransactions.length > 0
    ) {
      return true
    }
    if (pendingBridgeTransactions && pendingBridgeTransactions.length > 0) {
      return true
    }
    return false
  }, [pendingBridgeTransactions, pendingAwaitingCompletionTransactions])

  const hasHistoricalTransactions: boolean = useMemo(
    () => userHistoricalTransactions && userHistoricalTransactions.length > 0,
    [userHistoricalTransactions]
  )

  const hasNoTransactions: boolean = useMemo(() => {
    return !hasPendingTransactions && !hasHistoricalTransactions
  }, [hasPendingTransactions, hasHistoricalTransactions, address])

  const isLoading: boolean =
    isUserHistoricalTransactionsLoading && isUserPendingTransactionsLoading

  return (
    <div
      data-test-id="activity"
      className={`${visibility ? 'block' : 'hidden'}`}
    >
      {!address && (
        <div className="text-[#C2C2D6]">
          Your pending and recent transactions will appear here.
        </div>
      )}

      {address && isLoading && (
        <div className="text-[#C2C2D6]">Loading activity...</div>
      )}

      {address && !isLoading && hasNoTransactions && (
        <div className="text-[#C2C2D6]">
          Your pending and recent transactions will appear here.
          <ExplorerLink connectedAddress={address} />
        </div>
      )}

      {address && !isLoading && hasPendingTransactions && (
        <ActivitySection title="Pending" twClassName="flex flex-col mb-5">
          {pendingAwaitingCompletionTransactions &&
            pendingAwaitingCompletionTransactions.map(
              (transaction: BridgeTransaction) => (
                <PendingTransaction
                  connectedAddress={address as Address}
                  destinationAddress={transaction?.fromInfo?.address as Address}
                  startedTimestamp={transaction?.fromInfo?.time as number}
                  transactionHash={transaction?.fromInfo?.txnHash as string}
                  eventType={transaction?.fromInfo?.eventType as number}
                  isSubmitted={transaction?.fromInfo?.txnHash ? true : false}
                  isCompleted={transaction?.toInfo?.time ? true : false}
                  transactionType={TransactionType.PENDING}
                  originValue={transaction?.fromInfo?.formattedValue as number}
                  destinationValue={
                    transaction?.toInfo?.formattedValue as number
                  }
                  originChain={
                    CHAINS_BY_ID[transaction?.fromInfo?.chainID] as Chain
                  }
                  destinationChain={
                    CHAINS_BY_ID[
                      transaction?.fromInfo?.destinationChainID
                    ] as Chain
                  }
                  originToken={
                    tokenAddressToToken(
                      transaction?.fromInfo?.chainID,
                      transaction?.fromInfo?.tokenAddress
                    ) as Token
                  }
                  destinationToken={
                    tokenAddressToToken(
                      transaction?.toInfo?.chainID,
                      transaction?.toInfo?.tokenAddress
                    ) as Token
                  }
                />
              )
            )}
          <PendingTransactionAwaitingIndexing />
        </ActivitySection>
      )}

      {address && !isLoading && hasHistoricalTransactions && (
        <ActivitySection title="Recent">
          {userHistoricalTransactions
            .slice(0, 6) //temporarily only show recent 6
            .map((transaction: BridgeTransaction) => (
              <Transaction
                key={transaction.kappa}
                connectedAddress={address as Address}
                destinationAddress={transaction?.fromInfo?.address as Address}
                startedTimestamp={transaction?.fromInfo?.time as number}
                completedTimestamp={transaction?.toInfo?.time as number}
                transactionHash={transaction?.fromInfo?.txnHash as string}
                kappa={transaction?.kappa}
                isCompleted={true}
                transactionType={TransactionType.HISTORICAL}
                originValue={transaction?.fromInfo?.formattedValue as number}
                destinationValue={transaction?.toInfo?.formattedValue as number}
                originChain={
                  CHAINS_BY_ID[transaction?.fromInfo?.chainID] as Chain
                }
                originToken={
                  tokenAddressToToken(
                    transaction?.fromInfo?.chainID,
                    transaction?.fromInfo?.tokenAddress
                  ) as Token
                }
                destinationChain={
                  CHAINS_BY_ID[
                    transaction?.fromInfo?.destinationChainID
                  ] as Chain
                }
                destinationToken={
                  tokenAddressToToken(
                    transaction?.toInfo?.chainID,
                    transaction?.toInfo?.tokenAddress
                  ) as Token
                }
              />
            ))}
          <ExplorerLink connectedAddress={address} />
        </ActivitySection>
      )}
    </div>
  )
}

export const MostRecentTransaction = () => {
  const { address } = useAccount()
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()
  const {
    userPendingTransactions,
    userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
    seenHistoricalTransactions,
    pendingAwaitingCompletionTransactions,
  }: TransactionsState = useTransactionsState()
  const { activeTab }: PortfolioState = usePortfolioState()

  const [currentTime, setCurrentTime] = useState<number>(
    getTimeMinutesBeforeNow(0)
  )

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentTime(getTimeMinutesBeforeNow(0))
    }, 60000)

    return () => clearInterval(interval)
  }, [])

  const lastPendingBridgeTransaction: PendingBridgeTransaction = useMemo(() => {
    return pendingBridgeTransactions && pendingBridgeTransactions[0]
  }, [pendingBridgeTransactions])

  const lastPendingTransaction: BridgeTransaction = useMemo(() => {
    return (
      pendingAwaitingCompletionTransactions &&
      pendingAwaitingCompletionTransactions[0]
    )
  }, [pendingAwaitingCompletionTransactions])

  const lastHistoricalTransaction: BridgeTransaction = useMemo(() => {
    return userHistoricalTransactions && userHistoricalTransactions[0]
  }, [userHistoricalTransactions])

  const recentMinutesInUnix: number = 15 * 60

  const isLastHistoricalTransactionRecent: boolean = useMemo(
    () =>
      currentTime - lastHistoricalTransaction?.toInfo?.time <
      recentMinutesInUnix,
    [currentTime]
  )

  const seenLastHistoricalTransaction: boolean = useMemo(() => {
    if (!seenHistoricalTransactions || !userHistoricalTransactions) {
      return false
    }
    return seenHistoricalTransactions.some(
      (transaction: BridgeTransaction) =>
        transaction === (lastHistoricalTransaction as BridgeTransaction)
    )
  }, [seenHistoricalTransactions, lastHistoricalTransaction])

  let transaction

  if (isUserHistoricalTransactionsLoading || isUserPendingTransactionsLoading) {
    return null
  }

  if (lastPendingBridgeTransaction) {
    console.log('a1')
    transaction = lastPendingBridgeTransaction as PendingBridgeTransaction
    return (
      <PendingTransaction
        connectedAddress={address as Address}
        originChain={transaction.originChain as Chain}
        originToken={transaction.originToken as Token}
        originValue={Number(transaction.originValue)}
        destinationChain={transaction.destinationChain as Chain}
        destinationToken={transaction.destinationToken as Token}
        startedTimestamp={transaction.id ?? transaction.startedTimestamp}
        transactionHash={transaction.transactionHash as string}
        isSubmitted={transaction.isSubmitted as boolean}
        transactionType={TransactionType.PENDING}
      />
    )
  }

  if (lastPendingTransaction) {
    console.log('a2')
    transaction = lastPendingTransaction as BridgeTransaction
    return (
      <PendingTransaction
        connectedAddress={address as Address}
        startedTimestamp={transaction?.fromInfo?.time as number}
        transactionHash={transaction?.fromInfo?.txnHash as string}
        transactionType={TransactionType.PENDING}
        originValue={transaction?.fromInfo?.formattedValue as number}
        originChain={CHAINS_BY_ID[transaction?.fromInfo?.chainID] as Chain}
        destinationChain={
          CHAINS_BY_ID[transaction?.fromInfo?.destinationChainID] as Chain
        }
        originToken={
          tokenAddressToToken(
            transaction?.fromInfo?.chainID,
            transaction?.fromInfo?.tokenAddress
          ) as Token
        }
        destinationToken={
          tokenAddressToToken(
            transaction?.toInfo?.chainID,
            transaction?.toInfo?.tokenAddress
          ) as Token
        }
        destinationAddress={transaction?.fromInfo?.address as Address}
        isSubmitted={transaction?.fromInfo?.txnHash ? true : false}
        isCompleted={transaction?.toInfo?.time ? true : false}
      />
    )
  }

  if (
    lastHistoricalTransaction &&
    isLastHistoricalTransactionRecent &&
    !seenLastHistoricalTransaction
  ) {
    console.log('a3')
    transaction = lastHistoricalTransaction as BridgeTransaction
    return (
      <PendingTransaction
        connectedAddress={address as Address}
        destinationAddress={transaction?.fromInfo?.address as Address}
        startedTimestamp={transaction?.fromInfo?.time as number}
        completedTimestamp={transaction?.toInfo?.time as number}
        transactionHash={transaction?.fromInfo?.txnHash as string}
        kappa={transaction?.kappa as string}
        transactionType={TransactionType.PENDING}
        originValue={transaction?.fromInfo?.formattedValue as number}
        destinationValue={transaction?.toInfo?.formattedValue as number}
        originChain={CHAINS_BY_ID[transaction?.fromInfo?.chainID] as Chain}
        destinationChain={
          CHAINS_BY_ID[transaction?.fromInfo?.destinationChainID] as Chain
        }
        originToken={
          tokenAddressToToken(
            transaction?.fromInfo?.chainID,
            transaction?.fromInfo?.tokenAddress
          ) as Token
        }
        destinationToken={
          tokenAddressToToken(
            transaction?.toInfo?.chainID,
            transaction?.toInfo?.tokenAddress
          ) as Token
        }
        isSubmitted={transaction?.fromInfo?.txnHash ? true : false}
        isCompleted={true}
      />
    )
  }
}

export const PendingTransactionAwaitingIndexing = () => {
  const { address } = useAccount()
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()
  return (
    <>
      {pendingBridgeTransactions.map(
        (transaction: PendingBridgeTransaction) => (
          <PendingTransaction
            connectedAddress={address as Address}
            originChain={transaction.originChain as Chain}
            originToken={transaction.originToken as Token}
            originValue={Number(transaction.originValue)}
            destinationChain={transaction.destinationChain as Chain}
            destinationToken={transaction.destinationToken as Token}
            transactionHash={transaction.transactionHash}
            isSubmitted={transaction.isSubmitted as boolean}
            startedTimestamp={transaction.timestamp as number}
            transactionType={TransactionType.PENDING}
          />
        )
      )}
    </>
  )
}

export const ExplorerLink = ({
  connectedAddress,
}: {
  connectedAddress?: Address | string
}) => {
  const explorerLink: string = connectedAddress
    ? `${ANALYTICS_PATH}address/${connectedAddress}`
    : ANALYTICS_PATH
  return (
    <div data-test-id="explorer-link" className="text-[#99E6FF] my-3">
      <Link href={explorerLink} target="_blank">
        <span className="hover:underline">Explorer</span> →
      </Link>
    </div>
  )
}

export const ActivitySection = ({
  title,
  children,
  twClassName,
}: {
  title: string
  children?: React.ReactNode
  twClassName?: string
}) => {
  return (
    <div data-test-id="activity-section" className={twClassName}>
      <h3 className="mb-2 text-xl text-white">{title}</h3>
      {children}
    </div>
  )
}

export const getTransactionExplorerLink = ({
  kappa,
  fromChainId,
  toChainId,
}: {
  kappa: string
  fromChainId: number
  toChainId: number
}): string => {
  return `${ANALYTICS_KAPPA}${kappa}?chainIdFrom=${fromChainId}&chainIdTo=${toChainId}`
}

export const EstimatedDuration = ({
  startTime,
  transactionStatus,
  estimatedCompletionInSeconds,
}: {
  startTime: number
  transactionStatus: TransactionStatus
  estimatedCompletionInSeconds: number
}) => {
  const [elapsedTime, setElapsedTime] = useState<number>(0)

  useEffect(() => {
    const currentTime: number = Math.floor(Date.now() / 1000)
    const elapsedMinutes: number = Math.floor((currentTime - startTime) / 60)
    setElapsedTime(elapsedMinutes)
  }, [startTime])

  useEffect(() => {
    const interval = setInterval(() => {
      const currentTime: number = Math.floor(Date.now() / 1000)
      const elapsedMinutes: number = Math.floor((currentTime - startTime) / 60)
      setElapsedTime(elapsedMinutes)
    }, 60000)

    return () => {
      clearInterval(interval)
    }
  }, [startTime])

  const estimatedMinutes: number = Math.floor(estimatedCompletionInSeconds / 60)
  const timeRemaining: number = useMemo(() => {
    if (!startTime || !elapsedTime) return estimatedMinutes
    return estimatedMinutes - elapsedTime
  }, [estimatedMinutes, elapsedTime, startTime])

  return (
    <div className="text-[#C2C2D6] text-sm flex flex-col">
      {timeRemaining >= 0 ? (
        <React.Fragment>
          <div>
            {timeRemaining} - {timeRemaining + 1} min
          </div>
          {transactionStatus !== TransactionStatus.PENDING_WALLET_ACTION && (
            <ProcessingIcon className="fill-[#343036] mt-0.5" />
          )}
        </React.Fragment>
      ) : (
        <div>Waiting... </div>
      )}
    </div>
  )
}

export const Completed = ({
  transactionCompletedTime,
  connectedAddress,
  destinationAddress,
  handleExplorerClick,
}: {
  transactionCompletedTime: number
  connectedAddress?: Address | string
  destinationAddress: string
  handleExplorerClick: () => void
}) => {
  const formattedTime: string =
    transactionCompletedTime &&
    convertUnixTimestampToMonthAndDate(transactionCompletedTime)

  const isToday: boolean = isTimestampToday(transactionCompletedTime)

  const destinationIsSender: boolean =
    String(connectedAddress) === String(destinationAddress)

  return (
    <div
      data-test-id="completed"
      className="flex flex-col text-right text-[#C2C2D6] gap-1 text-sm"
      onClick={handleExplorerClick}
    >
      {!destinationIsSender && (
        <div>to {shortenAddress(destinationAddress, 3)} </div>
      )}
      {isToday ? (
        <div className="text-[#3BDD77] hover:underline cursor-pointer">
          Today
        </div>
      ) : (
        <div>{formattedTime}</div>
      )}
    </div>
  )
}

export const TransactionPayloadDetail = ({
  chain,
  token,
  tokenAmount,
}: {
  chain?: Chain
  token?: Token
  tokenAmount?: number
}) => {
  return (
    <div
      data-test-id="transaction-payload-detail"
      className="flex flex-col p-1 space-y-1"
    >
      {chain && (
        <div
          data-test-id="transaction-payload-network"
          className="flex flex-row items-center"
        >
          <Image
            src={chain.chainImg}
            className="w-4 h-4 mr-1.5 rounded-full"
            alt={`${chain.name} icon`}
          />
          <div>{chain.name}</div>
        </div>
      )}

      {token && tokenAmount && (
        <div
          data-test-id="transaction-payload-token"
          className="flex flex-row items-center"
        >
          <Image
            src={token?.icon}
            className="items-center w-4 h-4 mr-1.5 rounded-full"
            alt={`${token?.name} icon`}
          />
          {typeof tokenAmount === 'number' ? (
            <div className="mr-1">{tokenAmount}</div>
          ) : (
            <div className="mr-1">...</div>
          )}
          <div className="mt-0.5 text-sm">{token?.symbol}</div>
        </div>
      )}
    </div>
  )
}
