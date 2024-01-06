import { useState, useEffect, useMemo } from 'react'
import { useAccount, Address } from 'wagmi'
import { Chain, Token } from '@/utils/types'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { TransactionsState } from '@/slices/transactions/reducer'
import { PendingBridgeTransaction } from '@/slices/transactions/actions'
import { BridgeTransaction } from '@/slices/api/generated'
import { getTimeMinutesBeforeNow } from '@/utils/time'
import { TransactionType } from './Transaction'
import { tokenAddressToToken } from '@/constants/tokens'
import { CHAINS_BY_ID } from '@/constants/chains'
import { PendingTransaction } from './PendingTransaction'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'

export const MostRecentTransaction = () => {
  const { address } = useAccount()
  const {
    userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
    seenHistoricalTransactions,
    pendingAwaitingCompletionTransactions,
    fallbackQueryHistoricalTransactions,
    fallbackQueryPendingTransactions,
    pendingBridgeTransactions,
  }: TransactionsState = useTransactionsState()
  const { searchInput, searchedBalancesAndAllowances }: PortfolioState =
    usePortfolioState()

  const [currentTime, setCurrentTime] = useState<number>(
    getTimeMinutesBeforeNow(0)
  )

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentTime(getTimeMinutesBeforeNow(0))
    }, 30000)

    return () => clearInterval(interval)
  }, [])

  const masqueradeActive: boolean = useMemo(() => {
    return Object.keys(searchedBalancesAndAllowances).length > 0
  }, [searchedBalancesAndAllowances])

  const pendingAwaitingCompletionTransactionsWithFallback: BridgeTransaction[] =
    useMemo(() => {
      let transactions: BridgeTransaction[] = []

      if (checkTransactionsExist(pendingAwaitingCompletionTransactions)) {
        transactions = [...pendingAwaitingCompletionTransactions]
      }

      if (checkTransactionsExist(fallbackQueryPendingTransactions)) {
        const mergedTransactions = [
          ...transactions,
          ...fallbackQueryPendingTransactions,
        ]

        const uniqueMergedTransactions = Array.from(
          new Set(mergedTransactions.map((transaction) => transaction?.kappa))
        ).map((kappa) =>
          mergedTransactions.find((item) => item?.kappa === kappa)
        )
        return uniqueMergedTransactions
      }

      return transactions
    }, [
      pendingAwaitingCompletionTransactions,
      fallbackQueryPendingTransactions,
    ])

  const userHistoricalTransactionsWithFallback: BridgeTransaction[] =
    useMemo(() => {
      let transactions: BridgeTransaction[] = []

      if (checkTransactionsExist(userHistoricalTransactions)) {
        transactions = [...userHistoricalTransactions]
      }

      if (checkTransactionsExist(fallbackQueryHistoricalTransactions)) {
        const mergedTransactions = [
          ...fallbackQueryHistoricalTransactions,
          ...transactions,
        ]

        const uniqueMergedTransactions = Array.from(
          new Set(mergedTransactions.map((transaction) => transaction?.kappa))
        ).map((kappa) =>
          mergedTransactions.find((item) => item?.kappa === kappa)
        )
        return uniqueMergedTransactions
      }

      return transactions
    }, [userHistoricalTransactions, fallbackQueryHistoricalTransactions])

  const lastPendingBridgeTransaction: PendingBridgeTransaction = useMemo(() => {
    return pendingBridgeTransactions?.[0]
  }, [pendingBridgeTransactions])

  const lastPendingTransaction: BridgeTransaction = useMemo(() => {
    return pendingAwaitingCompletionTransactionsWithFallback?.[0]
  }, [pendingAwaitingCompletionTransactionsWithFallback])

  const lastHistoricalTransaction: BridgeTransaction = useMemo(() => {
    return userHistoricalTransactionsWithFallback?.[0]
  }, [userHistoricalTransactionsWithFallback])

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
    } else {
      return seenHistoricalTransactions?.some(
        (transaction: BridgeTransaction) =>
          transaction?.kappa ===
          (lastHistoricalTransaction?.kappa as BridgeTransaction)
      )
    }
  }, [seenHistoricalTransactions, lastHistoricalTransaction])

  let transaction

  return useMemo(() => {
    if (
      isUserHistoricalTransactionsLoading ||
      isUserPendingTransactionsLoading
    ) {
      return null
    }

    if (!masqueradeActive && lastPendingBridgeTransaction) {
      transaction = lastPendingBridgeTransaction as PendingBridgeTransaction
      return (
        <div
          data-test-id="most-recent-transaction-bridge-pending"
          className="mt-6"
        >
          <PendingTransaction
            connectedAddress={address as Address}
            originChain={transaction.originChain as Chain}
            originToken={transaction.originToken as Token}
            originValue={Number(transaction.originValue)}
            destinationChain={transaction.destinationChain as Chain}
            destinationToken={transaction.destinationToken as Token}
            estimatedDuration={transaction?.estimatedTime}
            bridgeModuleName={transaction?.bridgeModuleName}
            startedTimestamp={
              transaction.timestamp ? transaction.timestamp : transaction.id
            }
            transactionHash={transaction.transactionHash}
            isSubmitted={transaction.isSubmitted}
            transactionType={TransactionType.PENDING}
          />
        </div>
      )
    }

    if (!masqueradeActive && lastPendingTransaction) {
      transaction = lastPendingTransaction as BridgeTransaction
      return (
        <div data-test-id="most-recent-transaction-pending" className="mt-6">
          <PendingTransaction
            connectedAddress={address as Address}
            startedTimestamp={transaction?.fromInfo?.time}
            transactionHash={transaction?.fromInfo?.txnHash}
            transactionType={TransactionType.PENDING}
            originValue={transaction?.fromInfo?.value}
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
            formattedEventType={transaction?.fromInfo?.formattedEventType}
            estimatedDuration={transaction?.estimatedTime}
            destinationAddress={transaction?.fromInfo?.address as Address}
            isSubmitted={transaction?.fromInfo?.txnHash ? true : false}
            isCompleted={transaction?.toInfo?.time ? true : false}
            kappa={transaction?.kappa}
          />
        </div>
      )
    }

    if (
      !masqueradeActive &&
      lastHistoricalTransaction &&
      isLastHistoricalTransactionRecent &&
      !seenLastHistoricalTransaction
    ) {
      transaction = lastHistoricalTransaction as BridgeTransaction
      return (
        <div data-test-id="most-recent-transaction-historical" className="mt-6">
          <PendingTransaction
            connectedAddress={address as Address}
            destinationAddress={transaction?.fromInfo?.address as Address}
            startedTimestamp={transaction?.fromInfo?.time}
            completedTimestamp={transaction?.toInfo?.time}
            transactionHash={transaction?.fromInfo?.txnHash}
            kappa={transaction?.kappa}
            transactionType={TransactionType.PENDING}
            originValue={transaction?.fromInfo?.value}
            destinationValue={transaction?.toInfo?.value}
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
            formattedEventType={transaction?.fromInfo?.formattedEventType}
            isSubmitted={transaction?.fromInfo?.txnHash ? true : false}
            isCompleted={true}
          />
        </div>
      )
    }
  }, [
    currentTime,
    lastPendingBridgeTransaction,
    lastHistoricalTransaction,
    lastPendingTransaction,
    userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
    seenHistoricalTransactions,
    pendingAwaitingCompletionTransactions,
    fallbackQueryHistoricalTransactions,
    fallbackQueryPendingTransactions,
    pendingBridgeTransactions,
    masqueradeActive,
    seenLastHistoricalTransaction,
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
  ])
}
