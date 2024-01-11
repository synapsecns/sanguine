import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Address } from 'viem'

import {
  PortfolioTabs,
  FetchState,
  setActiveTab,
  resetPortfolioState,
  typeSearchInput,
  resetSearchState,
} from './actions'
import {
  fetchAndStorePortfolioBalances,
  fetchAndStoreSingleNetworkPortfolioBalances,
  fetchAndStoreSingleTokenBalance,
  fetchAndStoreSearchInputPortfolioBalances,
} from './hooks'
import {
  NetworkTokenBalances,
  TokenAndBalance,
} from '@/utils/actions/fetchPortfolioBalances'

export interface PortfolioState {
  activeTab: PortfolioTabs
  balances: NetworkTokenBalances
  poolTokenBalances: NetworkTokenBalances
  status: FetchState
  error?: string
  searchInput: string
  searchedBalances: {
    [index: Address]: NetworkTokenBalances
  }
  searchStatus: FetchState
}

export const initialState: PortfolioState = {
  activeTab: PortfolioTabs.PORTFOLIO,
  balances: {},
  poolTokenBalances: {},
  status: FetchState.IDLE,
  error: null,
  searchInput: '',
  searchedBalances: {},
  searchStatus: FetchState.IDLE,
}

export const portfolioSlice = createSlice({
  name: 'portfolio',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(setActiveTab, (state, action: PayloadAction<PortfolioTabs>) => {
        state.activeTab = action.payload
      })
      .addCase(typeSearchInput, (state, { payload: { searchInput } }) => {
        state.searchInput = searchInput
      })
      .addCase(fetchAndStoreSingleTokenBalance.fulfilled, (state, action) => {
        const { chainId, tokenAddress, balance, parsedBalance } = action.payload

        state.balances[chainId].forEach((token: TokenAndBalance) => {
          if (token.tokenAddress === tokenAddress) {
            token.balance = balance
            token.parsedBalance = parsedBalance
          }
        })
      })
      .addCase(fetchAndStorePortfolioBalances.pending, (state) => {
        state.status = FetchState.LOADING
      })
      .addCase(fetchAndStorePortfolioBalances.fulfilled, (state, action) => {
        state.status = FetchState.VALID
        state.balances = action.payload.balances
        state.poolTokenBalances = action.payload.poolTokenBalances
      })
      .addCase(fetchAndStorePortfolioBalances.rejected, (state, action) => {
        state.status = FetchState.INVALID
        state.error = action.error.message
      })
      .addCase(fetchAndStoreSearchInputPortfolioBalances.pending, (state) => {
        state.searchStatus = FetchState.LOADING
      })
      .addCase(
        fetchAndStoreSearchInputPortfolioBalances.fulfilled,
        (state, action) => {
          const { balances, address } = action.payload
          state.searchStatus = FetchState.VALID
          state.searchedBalances[address] = balances
        }
      )
      .addCase(
        fetchAndStoreSearchInputPortfolioBalances.rejected,
        (state, action) => {
          state.searchStatus = FetchState.INVALID
          state.error = action.error.message
        }
      )
      .addCase(
        fetchAndStoreSingleNetworkPortfolioBalances.fulfilled,
        (state, action) => {
          const { balances } = action.payload

          Object.entries(balances).forEach(
            ([chainId, mergedBalancesAndAllowances]) => {
              state.balances[chainId] = [...mergedBalancesAndAllowances]
            }
          )
        }
      )
      .addCase(resetSearchState, (state) => {
        state.searchedBalances = initialState.searchedBalances
        state.searchStatus = initialState.searchStatus
      })
      .addCase(resetPortfolioState, (state) => {
        state.activeTab = initialState.activeTab
        state.balances = initialState.balances
        state.status = initialState.status
        state.error = initialState.error
        state.searchInput = initialState.searchInput
        state.searchedBalances = initialState.searchedBalances
        state.searchStatus = initialState.searchStatus
        state.poolTokenBalances = initialState.poolTokenBalances
      })
  },
})

export default portfolioSlice.reducer
