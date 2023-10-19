import useWindowFocus from 'use-window-focus'
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
  addFallbackQueryHistoricalTransaction,
  removeFallbackQueryHistoricalTransaction,
  removeFallbackQueryPendingTransaction,
  resetTransactionsState,
  updateIsUserPendingTransactionsLoading,
} from './actions'
import {
  updateIsUserHistoricalTransactionsLoading,
  updateUserHistoricalTransactions,
  updateUserPendingTransactions,
} from './actions'
import { useBridgeState } from '../bridge/hooks'
import { BridgeState } from '../bridge/reducer'
import { PortfolioState } from '../portfolio/reducer'
import { usePortfolioState } from '../portfolio/hooks'
import { PortfolioTabs } from '../portfolio/actions'
import {
  updatePendingBridgeTransactions,
  removePendingBridgeTransaction,
  PendingBridgeTransaction,
} from '../bridge/actions'
import {
  addSeenHistoricalTransaction,
  addPendingAwaitingCompletionTransaction,
  removePendingAwaitingCompletionTransaction,
} from './actions'
import { isValidAddress } from '@/utils/isValidAddress'
import { checkTransactionsExist } from '@/components/Portfolio/Activity'

const queryHistoricalTime: number = getTimeMinutesBeforeNow(oneMonthInMinutes)
const queryPendingTime: number = getTimeMinutesBeforeNow(oneDayInMinutes)

export default function Updater(): null {
  const dispatch = useAppDispatch()
  const isWindowFocused: boolean = useWindowFocus()
  const {
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
    userHistoricalTransactions,
    userPendingTransactions,
    seenHistoricalTransactions,
    pendingAwaitingCompletionTransactions,
    fallbackQueryPendingTransactions,
    fallbackQueryHistoricalTransactions,
  }: TransactionsState = useTransactionsState()
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()
  const {
    activeTab,
    searchInput,
    searchedBalancesAndAllowances,
  }: PortfolioState = usePortfolioState()

  const [fetchUserHistoricalActivity, fetchedHistoricalActivity] =
    useLazyGetUserHistoricalActivityQuery({ pollingInterval: 10000 })

  const [fetchUserPendingActivity, fetchedPendingActivity] =
    useLazyGetUserPendingTransactionsQuery({ pollingInterval: 10000 })

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
   */
  useEffect(() => {
    if (address && isWindowFocused && !masqueradeActive) {
      fetchUserHistoricalActivity({
        address: address,
        startTime: queryHistoricalTime,
      })
      fetchUserPendingActivity({
        address: address,
        startTime: queryPendingTime,
      })
    } else if (
      masqueradeActive &&
      isWindowFocused &&
      searchedBalancesAndAllowances
    ) {
      const queriedAddress: Address = Object.keys(
        searchedBalancesAndAllowances
      )[0] as Address
      fetchUserHistoricalActivity({
        address: queriedAddress,
        startTime: queryHistoricalTime,
      })
      fetchUserPendingActivity({
        address: queriedAddress,
        startTime: queryPendingTime,
      })
    }
  }, [
    address,
    masqueradeActive,
    searchedBalancesAndAllowances,
    isWindowFocused,
  ])

  // Unsubscribe when address is unconnected/disconnected
  // useEffect(() => {
  //   const isLoading: boolean =
  //     isUserHistoricalTransactionsLoading || isUserPendingTransactionsLoading

  //   if ((!isLoading || masqueradeActive) && !isWindowFocused) {
  //     fetchUserHistoricalActivity({
  //       address: null,
  //       startTime: null,
  //     }).unsubscribe()

  //     fetchUserPendingActivity({
  //       address: null,
  //       startTime: null,
  //     }).unsubscribe()
  //   }
  // }, [
  //   isWindowFocused,
  //   masqueradeActive,
  //   isUserHistoricalTransactionsLoading,
  //   isUserPendingTransactionsLoading,
  // ])

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
   * Handles removing recent pending unindexed bridge transactions
   * from Bridge state once Explorer or Fallback query confirms transactions
   */
  useEffect(() => {
    const matchingTransactionHashes = new Set(
      pendingBridgeTransactions
        .filter(
          (recentTx) =>
            (userPendingTransactions &&
              userPendingTransactions.some(
                (pendingTx: BridgeTransaction) =>
                  pendingTx.fromInfo.txnHash === recentTx.transactionHash
              )) ||
            (userHistoricalTransactions &&
              userHistoricalTransactions.some(
                (historicalTx: BridgeTransaction) =>
                  historicalTx.fromInfo.txnHash === recentTx.transactionHash
              )) ||
            (fallbackQueryPendingTransactions &&
              fallbackQueryPendingTransactions.some(
                (pendingTx: BridgeTransaction) =>
                  pendingTx.fromInfo.txnHash === recentTx.transactionHash
              ))
        )
        .map(
          (matchingTx: PendingBridgeTransaction) => matchingTx.transactionHash
        )
    )

    if (matchingTransactionHashes.size === 0) {
      return
    } else {
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
          const isStored: boolean = pendingAwaitingCompletionTransactions.some(
            (storedTransaction: BridgeTransaction) =>
              storedTransaction.kappa === pendingTransaction.kappa
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

    if (
      hasUserHistoricalTransactions &&
      activeTab !== PortfolioTabs.PORTFOLIO
    ) {
      const mostRecentHistoricalTransaction: BridgeTransaction =
        userHistoricalTransactions[0]

      const isTransactionAlreadySeen = seenHistoricalTransactions.some(
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
            userHistoricalTransactions.some(
              (historicalTransaction: BridgeTransaction) => {
                return historicalTransaction.kappa === pendingTransaction.kappa
              }
            ) ||
            fallbackQueryHistoricalTransactions.some(
              (historicalTransaction: BridgeTransaction) => {
                return historicalTransaction.kappa === pendingTransaction.kappa
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

    if (hasPendingBridgeTransactions && hasUserHistoricalTransactions) {
      pendingBridgeTransactions.forEach(
        (pendingBridgeTransaction: PendingBridgeTransaction) => {
          const isCompleted: boolean =
            userHistoricalTransactions.some(
              (historicalTransaction: BridgeTransaction) => {
                return (
                  historicalTransaction.fromInfo.txnHash ===
                  pendingBridgeTransaction.transactionHash
                )
              }
            ) ||
            fallbackQueryHistoricalTransactions.some(
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
          fallbackQueryHistoricalTransactions.some(
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
  }, [fallbackQueryPendingTransactions])

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
            userHistoricalTransactions.some(
              (historicalTransaction: BridgeTransaction) =>
                historicalTransaction.kappa === fallbackTransaction.kappa
            )
          ) {
            dispatch(
              removeFallbackQueryHistoricalTransaction(fallbackTransaction)
            )
          }
        }
      )
    }
  }, [fallbackQueryHistoricalTransactions])

  return null
}
