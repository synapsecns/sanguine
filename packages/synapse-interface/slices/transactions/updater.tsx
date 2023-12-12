import { useEffect, useMemo } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { useAccount, Address } from 'wagmi'
import {
  useLazyGetUserHistoricalActivityQuery,
  useLazyGetUserPendingTransactionsQuery,
  BridgeTransaction,
  api,
} from '../api/generated'
import { useTransactionsState } from './hooks'
import { TransactionsState } from './reducer'
import {
  getTimeMinutesBeforeNow,
  oneMonthInMinutes,
  oneDayInMinutes,
} from '@/utils/time'
import {
  updatePendingBridgeTransactions,
  removePendingBridgeTransaction,
  PendingBridgeTransaction,
  addFallbackQueryHistoricalTransaction,
  removeFallbackQueryHistoricalTransaction,
  removeFallbackQueryPendingTransaction,
  resetTransactionsState,
  updateIsUserPendingTransactionsLoading,
  updateIsUserHistoricalTransactionsLoading,
  updateUserHistoricalTransactions,
  updateUserPendingTransactions,
  addSeenHistoricalTransaction,
  addPendingAwaitingCompletionTransaction,
  removePendingAwaitingCompletionTransaction,
} from './actions'
import { PortfolioState } from '../portfolio/reducer'
import { usePortfolioState } from '../portfolio/hooks'
import { PortfolioTabs } from '../portfolio/actions'
import { getValidAddress } from '@/utils/isValidAddress'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'

const queryHistoricalTime: number = getTimeMinutesBeforeNow(oneMonthInMinutes)
const queryPendingTime: number = getTimeMinutesBeforeNow(oneDayInMinutes)

export default function Updater(): null {
  const dispatch = useAppDispatch()
  const {
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
    userHistoricalTransactions,
    userPendingTransactions,
    seenHistoricalTransactions,
    pendingAwaitingCompletionTransactions,
    fallbackQueryPendingTransactions,
    fallbackQueryHistoricalTransactions,
    pendingBridgeTransactions,
  }: TransactionsState = useTransactionsState()
  const {
    activeTab,
    searchInput,
    searchedBalancesAndAllowances,
  }: PortfolioState = usePortfolioState()

  const [fetchUserHistoricalActivity, fetchedHistoricalActivity] =
    useLazyGetUserHistoricalActivityQuery({ pollingInterval: 3000000 })

  const [fetchUserPendingActivity, fetchedPendingActivity] =
    useLazyGetUserPendingTransactionsQuery({ pollingInterval: 3000000 })

  const { address } = useAccount({
    onDisconnect() {
      dispatch(resetTransactionsState())
    },
  })

  const masqueradeActive: boolean = useMemo(() => {
    return Object.keys(searchedBalancesAndAllowances).length > 0
  }, [searchedBalancesAndAllowances])

  /**
   * Handle fetching for historical and pending activity by polling Explorer endpoint
   * Will retrigger fetching for Masquerade Mode address when active
   * Will unsubscribe when no valid address provided
   */
  useEffect(() => {
    if (address && !masqueradeActive) {
      fetchUserHistoricalActivity({
        address: address,
        startTime: queryHistoricalTime,
      })
      fetchUserPendingActivity({
        address: address,
        startTime: queryPendingTime,
      })
    } else if (masqueradeActive && searchedBalancesAndAllowances) {
      const queriedAddress: Address = Object.keys(
        searchedBalancesAndAllowances
      )[0] as Address
      fetchUserHistoricalActivity({
        address: getValidAddress(queriedAddress),
        startTime: queryHistoricalTime,
      })
      fetchUserPendingActivity({
        address: getValidAddress(queriedAddress),
        startTime: queryPendingTime,
      })
    } else {
      fetchUserHistoricalActivity({
        address: null,
        startTime: null,
      }).unsubscribe()

      fetchUserPendingActivity({
        address: null,
        startTime: null,
      }).unsubscribe()
    }
  }, [address, masqueradeActive, searchedBalancesAndAllowances])

  // Load fetched historical transactions into state along with fetch status
  useEffect(() => {
    const {
      isLoading,
      isUninitialized,
      isSuccess,
      data: historicalData,
    } = fetchedHistoricalActivity

    if ((masqueradeActive || address) && isUserHistoricalTransactionsLoading) {
      !isLoading &&
        !isUninitialized &&
        dispatch(updateIsUserHistoricalTransactionsLoading(false))
    }

    if ((masqueradeActive || address) && isSuccess) {
      dispatch(
        updateUserHistoricalTransactions(historicalData?.bridgeTransactions)
      )
    }
  }, [
    fetchedHistoricalActivity,
    isUserHistoricalTransactionsLoading,
    address,
    masqueradeActive,
  ])

  // Load fetched pending transactions into state along with fetch status
  useEffect(() => {
    const {
      isLoading,
      isUninitialized,
      isSuccess,
      data: pendingData,
    } = fetchedPendingActivity

    if (address && isUserPendingTransactionsLoading) {
      !isLoading &&
        !isUninitialized &&
        dispatch(updateIsUserPendingTransactionsLoading(false))
    }

    if (address && isSuccess) {
      dispatch(updateUserPendingTransactions(pendingData?.bridgeTransactions))
    }
  }, [fetchedPendingActivity, isUserPendingTransactionsLoading, address])

  /**
   * Handles removing recent pending unindexed bridge transactions + stale unsubmitted pending bridge transactions
   * from Bridge state once Explorer or Fallback query confirms transactions
   */
  useEffect(() => {
    const matchingTransactionHashes = new Set(
      pendingBridgeTransactions
        ?.filter(
          (recentTx) =>
            (userPendingTransactions &&
              userPendingTransactions?.some(
                (pendingTx: BridgeTransaction) =>
                  pendingTx.fromInfo.txnHash === recentTx.transactionHash
              )) ||
            (userHistoricalTransactions &&
              userHistoricalTransactions?.some(
                (historicalTx: BridgeTransaction) =>
                  historicalTx.fromInfo.txnHash === recentTx.transactionHash
              )) ||
            (fallbackQueryPendingTransactions &&
              fallbackQueryPendingTransactions?.some(
                (pendingTx: BridgeTransaction) =>
                  pendingTx.fromInfo.txnHash === recentTx.transactionHash
              ))
        )
        .map(
          (matchingTx: PendingBridgeTransaction) => matchingTx.transactionHash
        )
    )

    if (matchingTransactionHashes.size !== 0) {
      const updatedRecentBridgeTransactions = pendingBridgeTransactions.filter(
        (recentTx: PendingBridgeTransaction) =>
          !matchingTransactionHashes.has(recentTx.transactionHash)
      )
      dispatch(updatePendingBridgeTransactions(updatedRecentBridgeTransactions))
    }
  }, [
    pendingBridgeTransactions,
    userHistoricalTransactions,
    userPendingTransactions,
    fallbackQueryPendingTransactions,
  ])

  // Store pending transactions until completed based on Explorer query
  useEffect(() => {
    const hasUserPendingTransactions: boolean =
      Array.isArray(userPendingTransactions) &&
      !isUserPendingTransactionsLoading

    if (hasUserPendingTransactions) {
      userPendingTransactions.forEach(
        (pendingTransaction: BridgeTransaction) => {
          const isStored: boolean = pendingAwaitingCompletionTransactions?.some(
            (storedTransaction: BridgeTransaction) =>
              storedTransaction?.kappa === pendingTransaction?.kappa
          )

          if (!isStored) {
            dispatch(
              addPendingAwaitingCompletionTransaction(pendingTransaction)
            )
          }
        }
      )
    }
  }, [userPendingTransactions])

  // Handle updating stored pending transactions state throughout progress
  useEffect(() => {
    const hasUserHistoricalTransactions: boolean =
      Array.isArray(userHistoricalTransactions) &&
      !isUserHistoricalTransactionsLoading

    const hasPendingBridgeTransactions: boolean =
      Array.isArray(pendingBridgeTransactions) &&
      pendingBridgeTransactions.length > 0

    if (hasUserHistoricalTransactions && activeTab === PortfolioTabs.ACTIVITY) {
      const mostRecentHistoricalTransaction: BridgeTransaction =
        userHistoricalTransactions[0]

      const isTransactionAlreadySeen = seenHistoricalTransactions?.some(
        (transaction: BridgeTransaction) =>
          transaction === (mostRecentHistoricalTransaction as BridgeTransaction)
      )

      if (!isTransactionAlreadySeen) {
        dispatch(addSeenHistoricalTransaction(mostRecentHistoricalTransaction))
      }
    }

    if (hasUserHistoricalTransactions) {
      pendingAwaitingCompletionTransactions.forEach(
        (pendingTransaction: BridgeTransaction) => {
          const isCompleted: boolean =
            userHistoricalTransactions?.some(
              (historicalTransaction: BridgeTransaction) => {
                return (
                  historicalTransaction?.kappa === pendingTransaction?.kappa
                )
              }
            ) ||
            fallbackQueryHistoricalTransactions?.some(
              (historicalTransaction: BridgeTransaction) => {
                return (
                  historicalTransaction?.kappa === pendingTransaction?.kappa
                )
              }
            )

          if (isCompleted) {
            dispatch(
              removePendingAwaitingCompletionTransaction(
                pendingTransaction.kappa
              )
            )
            dispatch(
              removeFallbackQueryPendingTransaction(pendingTransaction.kappa)
            )
          }
        }
      )
    }

    // Handle updating initial bridge transaction (unindexed) if completed
    if (hasPendingBridgeTransactions && hasUserHistoricalTransactions) {
      pendingBridgeTransactions.forEach(
        (pendingBridgeTransaction: PendingBridgeTransaction) => {
          const isCompleted: boolean =
            userHistoricalTransactions?.some(
              (historicalTransaction: BridgeTransaction) => {
                return (
                  historicalTransaction.fromInfo.txnHash ===
                  pendingBridgeTransaction.transactionHash
                )
              }
            ) ||
            fallbackQueryHistoricalTransactions?.some(
              (historicalTransaction: BridgeTransaction) => {
                return (
                  historicalTransaction.fromInfo.txnHash ===
                  pendingBridgeTransaction.transactionHash
                )
              }
            )

          if (isCompleted) {
            dispatch(
              removePendingBridgeTransaction(pendingBridgeTransaction.timestamp)
            )
          }
        }
      )
    }
  }, [userHistoricalTransactions, activeTab])

  // Handle adding completed fallback historical transaction to seen list
  useEffect(() => {
    const hasFallbackQueryHistoricalTransactions: boolean =
      checkTransactionsExist(fallbackQueryHistoricalTransactions)

    if (
      hasFallbackQueryHistoricalTransactions &&
      activeTab === PortfolioTabs.ACTIVITY
    ) {
      const mostRecentFallbackHistoricalTransaction: BridgeTransaction =
        fallbackQueryHistoricalTransactions[0]

      const isTransactionAlreadySeen = seenHistoricalTransactions?.some(
        (transaction: BridgeTransaction) =>
          transaction?.kappa ===
          (mostRecentFallbackHistoricalTransaction?.kappa as BridgeTransaction)
      )

      if (!isTransactionAlreadySeen) {
        dispatch(
          addSeenHistoricalTransaction(mostRecentFallbackHistoricalTransaction)
        )
      }
    }
  }, [fallbackQueryHistoricalTransactions, activeTab])

  /**
   * Handle fallback query returned transactions
   * If transaction is finalized (require destination Info and kappa),
   * move transaction into historical state to display in Activity
   */
  useEffect(() => {
    fallbackQueryPendingTransactions.forEach(
      (transaction: BridgeTransaction) => {
        const { fromInfo, toInfo, kappa } = transaction

        const alreadyMovedToHistorical: boolean =
          fallbackQueryHistoricalTransactions?.some(
            (historicalTransaction: BridgeTransaction) =>
              historicalTransaction !== transaction
          ) ||
          userHistoricalTransactions?.some(
            (historicalTransaction: BridgeTransaction) =>
              historicalTransaction !== transaction
          )

        if (fromInfo && toInfo && kappa && !alreadyMovedToHistorical) {
          dispatch(addFallbackQueryHistoricalTransaction(transaction))
          dispatch(removeFallbackQueryPendingTransaction(kappa))
          dispatch(removePendingAwaitingCompletionTransaction(kappa))
        }
      }
    )
  }, [fallbackQueryPendingTransactions, userHistoricalTransactions])

  /**
   * Handle removing fallback historical transaction from state
   * when identical transaction gets picked up by Explorer
   */
  useEffect(() => {
    const hasUserHistoricalTransactions: boolean =
      Array.isArray(userHistoricalTransactions) &&
      !isUserHistoricalTransactionsLoading

    if (
      hasUserHistoricalTransactions &&
      checkTransactionsExist(fallbackQueryHistoricalTransactions)
    ) {
      fallbackQueryHistoricalTransactions.forEach(
        (fallbackTransaction: BridgeTransaction) => {
          if (
            userHistoricalTransactions?.some(
              (historicalTransaction: BridgeTransaction) =>
                historicalTransaction?.kappa === fallbackTransaction?.kappa
            )
          ) {
            console.log('removed this transaction: ', fallbackTransaction)
            dispatch(
              removeFallbackQueryHistoricalTransaction(fallbackTransaction)
            )
          }
        }
      )
    }
  }, [fallbackQueryHistoricalTransactions, userHistoricalTransactions])

  /**
   * Handle removing fallback pending transaction from state
   * when completed fallback transaction appears in Historical Fallback Transactions
   */
  useEffect(() => {
    if (checkTransactionsExist(fallbackQueryHistoricalTransactions)) {
      fallbackQueryHistoricalTransactions.forEach(
        (fallbackHistoricalTransaction: BridgeTransaction) => {
          const matched: boolean = fallbackQueryPendingTransactions?.some(
            (pendingTransaction: BridgeTransaction) =>
              pendingTransaction?.kappa === fallbackHistoricalTransaction?.kappa
          )

          if (matched) {
            dispatch(
              removeFallbackQueryPendingTransaction(
                fallbackHistoricalTransaction.kappa
              )
            )
          }
        }
      )
    }
    if (checkTransactionsExist(fallbackQueryPendingTransactions)) {
      fallbackQueryPendingTransactions.forEach(
        (fallbackPendingTransaction: BridgeTransaction) => {
          const matched: boolean = userHistoricalTransactions?.some(
            (pendingTransaction: BridgeTransaction) =>
              pendingTransaction?.kappa === fallbackPendingTransaction?.kappa
          )

          if (matched) {
            dispatch(
              removeFallbackQueryPendingTransaction(
                fallbackPendingTransaction.kappa
              )
            )
          }
        }
      )
    }
  }, [fallbackQueryPendingTransactions, fallbackQueryHistoricalTransactions])

  return null
}
