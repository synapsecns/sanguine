import { useCallback } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { createAsyncThunk } from '@reduxjs/toolkit'
import { getAccount, Address } from '@wagmi/core'

import { AppDispatch, RootState } from '@/store/store'
import { FetchState, typeSearchInput, resetSearchState } from './actions'
import { useAppDispatch, useAppSelector } from '@/store/hooks'
import {
  fetchPortfolioBalances,
  NetworkTokenBalancesAndAllowances,
  getTokenBalances,
  TokenAndBalance,
} from '@/utils/actions/fetchPortfolioBalances'
import { getTokenAllowance } from './../../utils/actions/getTokenAllowance'
import { Token } from '@/utils/types'
import { initialState } from './reducer'

export const usePortfolioState = (): RootState['portfolio'] => {
  return useAppSelector((state) => state.portfolio)
}

export const usePortfolioBalances = (): NetworkTokenBalancesAndAllowances => {
  return useAppSelector((state) => state.portfolio.balancesAndAllowances)
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

export const fetchAndStoreSingleTokenAllowance = createAsyncThunk(
  'portfolio/fetchAndStoreSingleTokenAllowance',
  async ({
    routerAddress,
    tokenAddress,
    address,
    chainId,
  }: {
    routerAddress: Address
    tokenAddress: Address
    address: Address
    chainId: number
  }) => {
    const allowance = await getTokenAllowance(
      routerAddress,
      tokenAddress,
      address,
      chainId
    )
    return { routerAddress, chainId, tokenAddress, allowance }
  }
)

export const fetchAndStoreSingleTokenBalance = createAsyncThunk(
  'portfolio/fetchAndStoreSingleTokenBalance',
  async ({
    token,
    routerAddress,
    address,
    chainId,
  }: {
    token: Token
    routerAddress: Address
    address: Address
    chainId: number
  }) => {
    const data: TokenAndBalance[] = await getTokenBalances(
      address,
      [token],
      chainId
    )
    const { balance, parsedBalance }: TokenAndBalance = data[0]
    const tokenAddress = token.addresses[chainId] as Address
    const allowance = await getTokenAllowance(
      routerAddress,
      tokenAddress,
      address,
      chainId
    )
    return {
      routerAddress,
      chainId,
      tokenAddress,
      allowance,
      balance,
      parsedBalance,
    }
  }
)
