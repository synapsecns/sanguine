import { useEffect, useMemo } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { useAccount, Address } from 'wagmi'
import { useLazyGetUserHistoricalActivityQuery } from '../api/generated'
import { useTransactionsState } from './hooks'
import { TransactionsState } from './reducer'
import { getTimeMinutesBeforeNow, oneMonthInMinutes } from '@/utils/time'
import {
  resetTransactionsState,
  updateIsUserHistoricalTransactionsLoading,
  updateUserHistoricalTransactions,
} from './actions'
import { PortfolioState } from '../portfolio/reducer'
import { usePortfolioState } from '../portfolio/hooks'
import { getValidAddress } from '@/utils/isValidAddress'

const queryHistoricalTime: number = getTimeMinutesBeforeNow(oneMonthInMinutes)

const POLLING_INTERVAL: number = 300000 // 5 minutes in ms

export default function Updater(): null {
  const dispatch = useAppDispatch()
  const { isUserHistoricalTransactionsLoading }: TransactionsState =
    useTransactionsState()
  const { searchedBalances }: PortfolioState = usePortfolioState()

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
        address: address,
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
