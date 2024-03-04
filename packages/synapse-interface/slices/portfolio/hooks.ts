import { useCallback } from 'react'
import { createAsyncThunk } from '@reduxjs/toolkit'
import { getAccount } from '@wagmi/core'

import { RootState } from '@/store/store'
import { FetchState, typeSearchInput, resetSearchState } from './actions'
import { useAppDispatch, useAppSelector } from '@/store/hooks'
import {
  fetchPortfolioBalances,
  NetworkTokenBalances,
} from '@/utils/actions/fetchPortfolioBalances'
import { initialState } from './reducer'

export const usePortfolioState = (): RootState['portfolio'] => {
  return useAppSelector((state: RootState) => state.portfolio)
}

export const usePortfolioBalances = (): NetworkTokenBalances => {
  return useAppSelector((state: RootState) => state.portfolio.balances)
}

export const usePortfolioActionHandlers = (): {
  onSearchInput: (searchInput: string) => void
  clearSearchInput: () => void
  clearSearchResults: () => void
} => {
  const dispatch = useAppDispatch()

  const onSearchInput = useCallback(
    (searchInput: string) => {
      dispatch(typeSearchInput({ searchInput }))
    },
    [dispatch]
  )

  const clearSearchInput = useCallback(() => {
    dispatch(typeSearchInput({ searchInput: initialState.searchInput }))
  }, [dispatch])

  const clearSearchResults = useCallback(() => {
    dispatch(resetSearchState())
  }, [dispatch])

  return {
    onSearchInput,
    clearSearchInput,
    clearSearchResults,
  }
}

export const fetchAndStorePortfolioBalances = createAsyncThunk(
  'portfolio/fetchAndStorePortfolioBalances',
  async (address: string) => {
    const portfolioData = await fetchPortfolioBalances(address)
    return portfolioData
  }
)

export const fetchAndStoreSearchInputPortfolioBalances = createAsyncThunk(
  'portfolio/fetchAndStoreSearchInputPortfolioBalances',
  async (address: string) => {
    const portfolioData = await fetchPortfolioBalances(address)
    return { ...portfolioData, address }
  }
)

export const fetchAndStoreSingleNetworkPortfolioBalances = createAsyncThunk(
  'portfolio/fetchAndStoreSingleNetworkPortfolioBalances',
  async ({ address, chainId }: { address: string; chainId: number }) => {
    const portfolioData = await fetchPortfolioBalances(address, chainId)
    return portfolioData
  }
)

export const useFetchPortfolioBalances = (): {
  balances: NetworkTokenBalances
  fetchPortfolioBalances: () => void
  status: FetchState
  error: string
} => {
  const dispatch = useAppDispatch()
  const { address } = getAccount()
  const { balances, status, error } = usePortfolioState()

  const fetchPortfolioBalances = () => {
    if (address) {
      dispatch(fetchAndStorePortfolioBalances(address))
    }
  }

  return { balances, fetchPortfolioBalances, status, error }
}
