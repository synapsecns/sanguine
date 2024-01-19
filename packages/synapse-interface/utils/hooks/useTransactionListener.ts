/* eslint-disable */
import { useEffect, useMemo } from 'react'
import { useAccount, Address } from 'wagmi'

import {
  resetTransactionsState,
  updateIsUserHistoricalTransactionsLoading,
  updateUserHistoricalTransactions,
} from '@slices/transactions/actions'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { getValidAddress } from '@/utils/isValidAddress'
import { useLazyGetUserHistoricalActivityQuery } from '@/slices/api/generated'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { TransactionsState } from '@/slices/transactions/reducer'
import { useAppDispatch } from '@/store/hooks'
import { getTimeMinutesBeforeNow, oneMonthInMinutes } from '@/utils/time'

const queryHistoricalTime: number = getTimeMinutesBeforeNow(oneMonthInMinutes)
// const queryPendingTime: number = getTimeMinutesBeforeNow(oneDayInMinutes)

const POLLING_INTERVAL: number = 300000 // 5 minutes in ms

export const useTransactionListener = () => {
  const dispatch = useAppDispatch()
  const { isUserHistoricalTransactionsLoading }: TransactionsState =
    useTransactionsState()
  const { activeTab, searchedBalances }: PortfolioState = usePortfolioState()

  const [fetchUserHistoricalActivity, fetchedHistoricalActivity] =
    useLazyGetUserHistoricalActivityQuery({
      pollingInterval: POLLING_INTERVAL,
    })

  const { address } = useAccount({
    onDisconnect() {
      dispatch(resetTransactionsState())
    },
  })

  const masqueradeActive: boolean = useMemo(() => {
    return Object.keys(searchedBalances).length > 0
  }, [searchedBalances])

  /**
   * Handle fetching for historical activity by polling Explorer endpoint
   * Will retrigger fetching for Masquerade Mode address when active
   * Will unsubscribe when no valid address provided
   */
  useEffect(() => {
    if (address && !masqueradeActive) {
      fetchUserHistoricalActivity({
        address,
        startTime: queryHistoricalTime,
      })
    } else if (masqueradeActive && searchedBalances) {
      const queriedAddress: Address = Object.keys(
        searchedBalances
      )[0] as Address
      fetchUserHistoricalActivity({
        address: getValidAddress(queriedAddress),
        startTime: queryHistoricalTime,
      })
    } else {
      fetchUserHistoricalActivity({
        address: null,
        startTime: null,
      }).unsubscribe()
    }
  }, [address, masqueradeActive, searchedBalances])

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

  return null
}
