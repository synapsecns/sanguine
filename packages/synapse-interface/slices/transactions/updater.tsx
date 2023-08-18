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
import { updateRecentBridgeTransactions } from '../bridge/actions'

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
  const { recentBridgeTransactions }: BridgeState = useBridgeState()

  const [fetchUserHistoricalActivity, fetchedHistoricalActivity] =
    useLazyGetUserHistoricalActivityQuery({ pollingInterval: 3000 })

  const [fetchUserPendingActivity, fetchedPendingActivity] =
    useLazyGetUserPendingTransactionsQuery({ pollingInterval: 3000 })

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
    const matchingTransactionHashes = recentBridgeTransactions
      .filter(
        (recentTx) =>
          userPendingTransactions.some(
            (pendingTx) =>
              pendingTx.fromInfo.txnHash === recentTx.transactionHash
          ) ||
          userHistoricalTransactions.some(
            (historicalTx) =>
              historicalTx.fromInfo.txnHash === recentTx.transactionHash
          )
      )
      .map((matchingTx) => matchingTx.transactionHash)

    if (matchingTransactionHashes.length > 0) {
      const updatedRecentBridgeTransactions = recentBridgeTransactions.filter(
        (recentTx) =>
          !matchingTransactionHashes.includes(recentTx.transactionHash)
      )
      dispatch(updateRecentBridgeTransactions(updatedRecentBridgeTransactions))
    }
  }, [
    recentBridgeTransactions,
    userHistoricalTransactions,
    userPendingTransactions,
  ])

  return null
}
