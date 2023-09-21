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

  // Start fetch when connected address exists
  // Unsubscribe when address is unconnected
  useEffect(() => {
    fetchUserHistoricalActivity({
      address: address,
      startTime: queryHistoricalTime,
    })
    fetchUserPendingActivity({
      address: address,
      startTime: queryPendingTime,
    })
  }, [address])

  useEffect(() => {
    const {
      isLoading,
      isUninitialized,
      isSuccess,
      data: historicalData,
    } = fetchedHistoricalActivity

    if (address && isUserHistoricalTransactionsLoading) {
      !isLoading &&
        !isUninitialized &&
        dispatch(updateIsUserHistoricalTransactionsLoading(false))
    }

    if (address && isSuccess) {
      dispatch(
        updateUserHistoricalTransactions(historicalData?.bridgeTransactions)
      )
    }
  }, [fetchedHistoricalActivity, isUserHistoricalTransactionsLoading, address])

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

  // Remove Recent Bridge Transaction from Bridge State when picked up by indexer
  useEffect(() => {
    const matchingTransactionHashes = new Set(
      pendingBridgeTransactions
        .filter(
          (recentTx) =>
            (userPendingTransactions &&
              userPendingTransactions.some(
                (pendingTx) =>
                  pendingTx.fromInfo.txnHash === recentTx.transactionHash
              )) ||
            (userHistoricalTransactions &&
              userHistoricalTransactions.some(
                (historicalTx) =>
                  historicalTx.fromInfo.txnHash === recentTx.transactionHash
              ))
        )
        .map((matchingTx) => matchingTx.transactionHash)
    )

    if (matchingTransactionHashes.size === 0) {
      return
    } else {
      const updatedRecentBridgeTransactions = pendingBridgeTransactions.filter(
        (recentTx) => !matchingTransactionHashes.has(recentTx.transactionHash)
      )
      dispatch(updatePendingBridgeTransactions(updatedRecentBridgeTransactions))
    }
  }, [
    pendingBridgeTransactions,
    userHistoricalTransactions,
    userPendingTransactions,
  ])

  // Store pending transactions until completed based on subgraph query
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
          const isCompleted: boolean = userHistoricalTransactions.some(
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
          }
        }
      )
    }

    if (hasPendingBridgeTransactions && hasUserHistoricalTransactions) {
      pendingBridgeTransactions.forEach(
        (pendingBridgeTransaction: PendingBridgeTransaction) => {
          const isCompleted: boolean = userHistoricalTransactions.some(
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

  return null
}
