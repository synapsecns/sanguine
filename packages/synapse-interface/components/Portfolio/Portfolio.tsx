import React, { useEffect, useMemo } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { useAppDispatch } from '@/store/hooks'
import { setFromChainId } from '@/slices/bridge/reducer'
import { PortfolioTabManager } from './PortfolioTabManager'
import {
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowance,
} from '@/utils/actions/fetchPortfolioBalances'
import { PortfolioContent, HomeContent } from './PortfolioContent'
import {
  useFetchPortfolioBalances,
  fetchAndStorePortfolioBalances,
  usePortfolioState,
} from '@/slices/portfolio/hooks'
import {
  PortfolioTabs,
  resetPortfolioState,
  setActiveTab,
} from '@/slices/portfolio/actions'
import { resetTransactionsState } from '@/slices/transactions/actions'
import { Activity } from './Activity'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import {
  updateUserHistoricalTransactions,
  updateIsUserHistoricalTransactionsLoading,
  updateUserPendingTransactions,
} from '@/slices/transactions/actions'
import {
  useLazyGetUserHistoricalActivityQuery,
  useLazyGetUserPendingTransactionsQuery,
  BridgeTransaction,
} from '@/slices/api/generated'
import {
  getTimeMinutesBeforeNow,
  oneMonthInMinutes,
  oneDayInMinutes,
} from '@/utils/time'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { TransactionsState } from '@/slices/transactions/reducer'

// const queryHistoricalTime: number = getTimeMinutesBeforeNow(oneMonthInMinutes)
// const queryPendingTime: number = getTimeMinutesBeforeNow(oneDayInMinutes)

export const Portfolio = () => {
  const dispatch = useAppDispatch()
  const { fromChainId }: BridgeState = useBridgeState()
  const { activeTab }: PortfolioState = usePortfolioState()
  // const {
  //   userPendingTransactions,
  //   isUserHistoricalTransactionsLoading,
  // }: TransactionsState = useTransactionsState()

  // const [fetchUserHistoricalActivity, fetchedHistoricalActivity] =
  //   useLazyGetUserHistoricalActivityQuery()

  // const [fetchUserPendingActivity, fetchedPendingActivity] =
  //   useLazyGetUserPendingTransactionsQuery({ pollingInterval: 3000 })

  const { chain } = useNetwork()
  const { address } = useAccount({
    onConnect() {
      dispatch(setActiveTab(PortfolioTabs.PORTFOLIO))
      // fetchUserHistoricalActivity({
      //   address: address,
      //   startTime: queryHistoricalTime,
      // })
      // fetchUserPendingActivity({
      //   address: address,
      //   startTime: queryPendingTime,
      // })
    },
    onDisconnect() {
      dispatch(resetPortfolioState())
      dispatch(resetTransactionsState())
    },
  })

  const { balancesAndAllowances: portfolioData, status: fetchState } =
    useFetchPortfolioBalances()

  const filteredPortfolioDataForBalances: NetworkTokenBalancesAndAllowances =
    filterPortfolioBalancesWithBalances(portfolioData)

  // useEffect(() => {
  //   const {
  //     isLoading,
  //     isUninitialized,
  //     isSuccess,
  //     data: historicalData,
  //   } = fetchedHistoricalActivity

  //   if (address && isUserHistoricalTransactionsLoading) {
  //     !isLoading &&
  //       !isUninitialized &&
  //       dispatch(updateIsUserHistoricalTransactionsLoading(false))
  //   }

  //   if (address && isSuccess) {
  //     dispatch(
  //       updateUserHistoricalTransactions(historicalData?.bridgeTransactions)
  //     )
  //   }
  // }, [fetchedHistoricalActivity, isUserHistoricalTransactionsLoading, address])

  // const userPendingActivity: BridgeTransaction[] = useMemo(() => {
  //   const { isSuccess, data: pendingData } = fetchedPendingActivity
  //   // console.log('fetchedPendingActivity useMemo:', fetchedPendingActivity)
  //   // console.log('pendingData:', pendingData)
  //   return isSuccess ? pendingData?.bridgeTransactions : userPendingTransactions
  // }, [fetchedPendingActivity, address])

  // useEffect(() => {
  //   dispatch(updateUserPendingTransactions(userPendingActivity))
  //   fetchUserHistoricalActivity({
  //     address: address,
  //     startTime: queryHistoricalTime,
  //   })
  // }, [userPendingActivity])

  useEffect(() => {
    ;(async () => {
      if (address && chain?.id) {
        await dispatch(setFromChainId(chain.id))
        await dispatch(fetchAndStorePortfolioBalances(address))
      }
    })()
  }, [chain, address])

  return (
    <div
      data-test-id="portfolio"
      className="flex flex-col w-full max-w-lg mx-auto lg:mx-0"
    >
      <PortfolioTabManager />
      <div className="mt-4">
        {activeTab === PortfolioTabs.HOME && <HomeContent />}
        {activeTab === PortfolioTabs.PORTFOLIO && (
          <PortfolioContent
            connectedAddress={address}
            connectedChainId={chain?.id}
            selectedFromChainId={fromChainId}
            networkPortfolioWithBalances={filteredPortfolioDataForBalances}
            fetchState={fetchState}
          />
        )}
        {activeTab === PortfolioTabs.ACTIVITY && <Activity />}
      </div>
    </div>
  )
}

function filterPortfolioBalancesWithBalances(
  balancesAndAllowances: NetworkTokenBalancesAndAllowances
): NetworkTokenBalancesAndAllowances {
  return Object.entries(balancesAndAllowances).reduce(
    (
      filteredBalances: NetworkTokenBalancesAndAllowances,
      [key, tokenWithBalances]
    ) => {
      const filtered = tokenWithBalances.filter(
        (token: TokenWithBalanceAndAllowance) => token.balance > 0n
      )
      if (filtered.length > 0) {
        filteredBalances[key] = filtered
      }
      return filteredBalances
    },
    {}
  )
}
