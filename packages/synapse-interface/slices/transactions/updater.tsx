import { useEffect, useMemo } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { useAccount } from 'wagmi'
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
import { updatePendingBridgeTransactions } from '../bridge/actions'
import { addSeenHistoricalTransaction } from './actions'

const queryHistoricalTime: number = getTimeMinutesBeforeNow(oneMonthInMinutes)
const queryPendingTime: number = getTimeMinutesBeforeNow(oneDayInMinutes)

export default function Updater(): null {
  const dispatch = useAppDispatch()
  const {
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
    userHistoricalTransactions,
    userPendingTransactions,
  }: TransactionsState = useTransactionsState()
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()
  const { activeTab }: PortfolioState = usePortfolioState()

  useEffect(() => {
    if (
      userHistoricalTransactions &&
      userHistoricalTransactions.length > 0 &&
      activeTab !== PortfolioTabs.PORTFOLIO
    ) {
      dispatch(addSeenHistoricalTransaction(userHistoricalTransactions[0]))
    }
  }, [userHistoricalTransactions, activeTab])

  const [fetchUserHistoricalActivity, fetchedHistoricalActivity] =
    useLazyGetUserHistoricalActivityQuery({ pollingInterval: 5000 })

  const [fetchUserPendingActivity, fetchedPendingActivity] =
    useLazyGetUserPendingTransactionsQuery({ pollingInterval: 5000 })

  const { address } = useAccount({
    onDisconnect() {
      dispatch(resetTransactionsState())
    },
  })

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

  return null
}
