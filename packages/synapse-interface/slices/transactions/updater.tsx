import { useEffect, useMemo } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { useAccount } from 'wagmi'
import {
  useLazyGetUserHistoricalActivityQuery,
  useLazyGetUserPendingTransactionsQuery,
  BridgeTransaction,
} from '../api/generated'
import { useTransactionsState } from './hooks'
import { TransactionsState } from './reducer'
import {
  getTimeMinutesBeforeNow,
  oneMonthInMinutes,
  oneDayInMinutes,
} from '@/utils/time'
import { resetTransactionsState } from './actions'
import {
  updateIsUserHistoricalTransactionsLoading,
  updateUserHistoricalTransactions,
  updateUserPendingTransactions,
} from './actions'

const queryHistoricalTime: number = getTimeMinutesBeforeNow(oneMonthInMinutes)
const queryPendingTime: number = getTimeMinutesBeforeNow(oneDayInMinutes)

export default function Updater(): null {
  const dispatch = useAppDispatch()
  const {
    userPendingTransactions,
    isUserHistoricalTransactionsLoading,
  }: TransactionsState = useTransactionsState()

  const [fetchUserHistoricalActivity, fetchedHistoricalActivity] =
    useLazyGetUserHistoricalActivityQuery({ pollingInterval: 3000 })

  const [fetchUserPendingActivity, fetchedPendingActivity] =
    useLazyGetUserPendingTransactionsQuery({ pollingInterval: 3000 })

  const { address } = useAccount({
    onConnect() {
      fetchUserHistoricalActivity({
        address: address,
        startTime: queryHistoricalTime,
      })
      fetchUserPendingActivity({
        address: address,
        startTime: queryPendingTime,
      })
    },
    onDisconnect() {
      dispatch(resetTransactionsState())
    },
  })

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
    const { isSuccess, data: pendingData } = fetchedPendingActivity

    console.log('pendingData: ', pendingData)

    console.log('fetchedPendingActivity:', fetchedPendingActivity)
    if (address && isSuccess) {
      dispatch(updateUserPendingTransactions(pendingData?.bridgeTransactions))
    }
  }, [fetchedPendingActivity, address])

  // const userPendingActivity: BridgeTransaction[] = useMemo(() => {
  //   const { isSuccess, data: pendingData } = fetchedPendingActivity
  //   return isSuccess ? pendingData?.bridgeTransactions : userPendingTransactions
  // }, [fetchedPendingActivity, address])

  // useEffect(() => {
  //   dispatch(updateUserPendingTransactions(userPendingActivity))
  //   fetchUserHistoricalActivity({
  //     address: address,
  //     startTime: queryHistoricalTime,
  //   })
  // }, [userPendingActivity])

  return null
}
