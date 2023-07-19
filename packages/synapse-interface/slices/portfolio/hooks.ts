import { useDispatch, useSelector } from 'react-redux'
import { createAsyncThunk } from '@reduxjs/toolkit'
import { getAccount } from '@wagmi/core'

import { AppDispatch, RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'
import { FetchState } from './reducer'
import {
  fetchPortfolioBalances,
  NetworkTokenBalancesAndAllowances,
} from '@/utils/hooks/usePortfolioBalances'

export const usePortfolioState = (): RootState['portfolio'] => {
  return useAppSelector((state) => state.portfolio)
}

export const usePortfolioBalances = (): NetworkTokenBalancesAndAllowances => {
  return useAppSelector((state) => state.portfolio.balancesAndAllowances)
}

export const fetchAndStorePortfolioBalances = createAsyncThunk(
  'portfolio/fetchAndStorePortfolioBalances',
  async (address: string) => {
    console.log('got hit')
    const portfolioData = await fetchPortfolioBalances(address)
    return portfolioData
  }
)

export const useFetchPortfolioBalances = (): {
  balancesAndAllowances: NetworkTokenBalancesAndAllowances
  fetchPortfolioBalances: () => void
  status: FetchState
  error: string
} => {
  const dispatch: AppDispatch = useDispatch()
  const { address } = getAccount()
  const { balancesAndAllowances, status, error } = useSelector(
    (state: RootState) => state.portfolio
  )

  const fetch = () => {
    if (address) {
      dispatch(fetchAndStorePortfolioBalances(address))
    }
  }

  return { balancesAndAllowances, fetchPortfolioBalances: fetch, status, error }
}
