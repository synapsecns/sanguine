import { useEffect, useMemo, useCallback, useState } from 'react'
import { useAccount, Address } from 'wagmi'
import { useRouter } from 'next/router'
import Link from 'next/link'
import Image from 'next/image'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { PartialInfo, BridgeTransaction } from '@/slices/api/generated'
import {
  convertUnixTimestampToMonthAndDate,
  getTimeMinutesBeforeNow,
  isTimestampToday,
} from '@/utils/time'
import { CHAINS_BY_ID } from '@/constants/chains'
import { Chain, Token } from '@/utils/types'
import {
  BRIDGABLE_TOKENS,
  tokenAddressToToken,
  tokenSymbolToToken,
} from '@/constants/tokens'
import { ANALYTICS_KAPPA, ANALYTICS_PATH } from '@/constants/urls'
import EtherscanIcon from '../icons/EtherscanIcon'
import { TransactionsState } from '@/slices/transactions/reducer'
import { PendingBridgeTransaction } from '@/slices/bridge/actions'
import { BridgeState } from '@/slices/bridge/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { getExplorerTxUrl } from '@/constants/urls'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { PortfolioTabs } from '@/slices/portfolio/actions'
import { shortenAddress } from '@/utils/shortenAddress'
import { BRIDGE_REQUIRED_CONFIRMATIONS } from '@/constants/bridge'
import { Transaction, PendingTransaction, TransactionType } from './Transaction'

export const Activity = ({ visibility }: { visibility: boolean }) => {
  const { address } = useAccount()
  const {
    userHistoricalTransactions,
    userPendingTransactions,
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
  }: TransactionsState = useTransactionsState()
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()

  const hasPendingTransactions: boolean = useMemo(() => {
    if (userPendingTransactions && userPendingTransactions.length > 0)
      return true
    if (pendingBridgeTransactions && pendingBridgeTransactions.length > 0)
      return true
    return false
  }, [userPendingTransactions, pendingBridgeTransactions])

  const hasHistoricalTransactions: boolean = useMemo(
    () => userHistoricalTransactions && userHistoricalTransactions.length > 0,
    [userHistoricalTransactions]
  )

  const historicalTransactionsByTime = useMemo(() => {
    if (!hasHistoricalTransactions) return

    const currentUnixTimestamp = Math.floor(Date.now() / 1000)
    const tenMinutesAgoUnixTimestamp = currentUnixTimestamp - 600

    const transactionsWithinLast10Mins = userHistoricalTransactions.filter(
      (transaction) =>
        transaction.fromInfo?.time &&
        transaction.fromInfo.time >= tenMinutesAgoUnixTimestamp
    )
    const remainingTransactions = userHistoricalTransactions.filter(
      (transaction) =>
        transaction.fromInfo?.time &&
        transaction.fromInfo.time < tenMinutesAgoUnixTimestamp
    )

    return { transactionsWithinLast10Mins, remainingTransactions }
  }, [hasHistoricalTransactions])

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

      {address &&
        !isLoading &&
        (hasPendingTransactions ||
          (historicalTransactionsByTime?.transactionsWithinLast10Mins &&
            historicalTransactionsByTime?.transactionsWithinLast10Mins.length >
              0)) && (
          <ActivitySection title="Pending" twClassName="flex flex-col mb-5">
            <PendingTransactionAwaitingIndexing />
            {userPendingTransactions &&
              userPendingTransactions.map((transaction: BridgeTransaction) => (
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
              ))}
            {historicalTransactionsByTime?.transactionsWithinLast10Mins &&
              historicalTransactionsByTime?.transactionsWithinLast10Mins
                .length > 0 &&
              historicalTransactionsByTime.transactionsWithinLast10Mins.map(
                (transaction: BridgeTransaction) => (
                  <PendingTransaction
                    connectedAddress={address as Address}
                    destinationAddress={
                      transaction?.fromInfo?.address as Address
                    }
                    startedTimestamp={transaction?.fromInfo?.time as number}
                    completedTimestamp={transaction?.toInfo?.time as number}
                    transactionHash={transaction?.fromInfo?.txnHash as string}
                    eventType={transaction?.fromInfo?.eventType as number}
                    kappa={transaction?.kappa}
                    isSubmitted={transaction?.fromInfo?.txnHash ? true : false}
                    isCompleted={true}
                    transactionType={TransactionType.PENDING}
                    originValue={
                      transaction?.fromInfo?.formattedValue as number
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
          </ActivitySection>
        )}

      {address && !isLoading && hasHistoricalTransactions && (
        <ActivitySection title="Recent">
          {historicalTransactionsByTime?.remainingTransactions &&
            historicalTransactionsByTime.remainingTransactions
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
                  destinationValue={
                    transaction?.toInfo?.formattedValue as number
                  }
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
  }: TransactionsState = useTransactionsState()

  const lastPendingBridgeTransaction: PendingBridgeTransaction =
    pendingBridgeTransactions && pendingBridgeTransactions[0]

  const lastPendingTransaction: BridgeTransaction =
    userPendingTransactions && userPendingTransactions[0]

  const currentTime: number = getTimeMinutesBeforeNow(0)
  const tenMinutesInUnix: number = 10 * 60

  const lastHistoricalTransaction: BridgeTransaction =
    userHistoricalTransactions && userHistoricalTransactions[0]
  const isLastHistoricalTransactionRecent: boolean =
    currentTime - lastHistoricalTransaction?.toInfo?.time < tenMinutesInUnix

  let transaction

  if (isUserHistoricalTransactionsLoading || isUserPendingTransactionsLoading) {
    return null
  }

  if (lastPendingBridgeTransaction) {
    transaction = lastPendingBridgeTransaction as PendingBridgeTransaction
    return (
      <PendingTransaction
        connectedAddress={address as Address}
        originChain={transaction.originChain as Chain}
        originToken={transaction.originToken as Token}
        originValue={Number(transaction.originValue)}
        destinationChain={transaction.destinationChain as Chain}
        destinationToken={transaction.destinationToken as Token}
        startedTimestamp={transaction.timestamp as number}
        transactionHash={transaction.transactionHash as string}
        isSubmitted={transaction.isSubmitted as boolean}
        transactionType={TransactionType.PENDING}
      />
    )
  }

  if (lastPendingTransaction) {
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

  if (lastHistoricalTransaction && isLastHistoricalTransactionRecent) {
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
  estimatedCompletionInSeconds,
}: {
  startTime: number
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
  const timeRemaining: number = estimatedMinutes - elapsedTime

  return (
    <div className="text-[#C2C2D6] text-sm">
      {timeRemaining >= 0 ? (
        <div>
          {timeRemaining} - {timeRemaining + 1} min
        </div>
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
}: {
  transactionCompletedTime: number
  connectedAddress?: Address | string
  destinationAddress: string
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
    >
      {!destinationIsSender && (
        <div>to {shortenAddress(destinationAddress, 3)} </div>
      )}
      {isToday ? (
        <div className="text-[#3BDD77]">Today</div>
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
