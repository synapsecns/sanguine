import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Address } from 'viem'

import {
  PortfolioTabs,
  FetchState,
  setActiveTab,
  updateSingleTokenAllowance,
  resetPortfolioState,
  typeSearchInput,
  resetSearchState,
} from './actions'
import {
  fetchAndStorePortfolioBalances,
  fetchAndStoreSingleNetworkPortfolioBalances,
  fetchAndStoreSingleTokenAllowance,
  fetchAndStoreSingleTokenBalance,
  fetchAndStoreSearchInputPortfolioBalances,
} from './hooks'
import {
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowances,
} from '@/utils/actions/fetchPortfolioBalances'

export interface PortfolioState {
  activeTab: PortfolioTabs
  balancesAndAllowances: NetworkTokenBalancesAndAllowances
  poolTokenBalances: NetworkTokenBalancesAndAllowances
  status: FetchState
  error?: string
  searchInput: string
  searchedBalancesAndAllowances: {
    [index: Address]: NetworkTokenBalancesAndAllowances
  }
  searchStatus: FetchState
}

export const initialState: PortfolioState = {
  activeTab: PortfolioTabs.PORTFOLIO,
  balancesAndAllowances: {},
  poolTokenBalances: {},
  status: FetchState.IDLE,
  error: null,
  searchInput: '',
  searchedBalancesAndAllowances: {},
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
      .addCase(updateSingleTokenAllowance, (state, action) => {
        const { chainId, allowance, spender, token } = action.payload

        state.balancesAndAllowances[chainId].forEach(
          (t: TokenWithBalanceAndAllowances) => {
            if (t.tokenAddress === token.addresses[chainId]) {
              t.allowances[spender] = allowance
            }
          }
        )
      })
      .addCase(fetchAndStoreSingleTokenAllowance.fulfilled, (state, action) => {
        const { routerAddress, chainId, tokenAddress, allowance } =
          action.payload

        state.balancesAndAllowances[chainId].forEach(
          (token: TokenWithBalanceAndAllowances) => {
            if (token.tokenAddress === tokenAddress) {
              token.allowances[routerAddress] = allowance
            }
          }
        )
      })
      .addCase(fetchAndStoreSingleTokenBalance.fulfilled, (state, action) => {
        const {
          routerAddress,
          chainId,
          tokenAddress,
          allowance,
          balance,
          parsedBalance,
        } = action.payload

        state.balancesAndAllowances[chainId].forEach(
          (token: TokenWithBalanceAndAllowances) => {
            if (token.tokenAddress === tokenAddress) {
              token.allowances[routerAddress] = allowance
              token.balance = balance
              token.parsedBalance = parsedBalance
            }
          }
        )
      })
      .addCase(fetchAndStorePortfolioBalances.pending, (state) => {
        state.status = FetchState.LOADING
      })
      .addCase(fetchAndStorePortfolioBalances.fulfilled, (state, action) => {
        state.status = FetchState.VALID
        state.balancesAndAllowances = action.payload.balancesAndAllowances
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
          const { balancesAndAllowances, address } = action.payload
          state.searchStatus = FetchState.VALID
          state.searchedBalancesAndAllowances[address] = balancesAndAllowances
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
          const { balancesAndAllowances } = action.payload

          Object.entries(balancesAndAllowances).forEach(
            ([chainId, mergedBalancesAndAllowances]) => {
              state.balancesAndAllowances[chainId] = [
                ...mergedBalancesAndAllowances,
              ]
            }
          )
        }
      )
      .addCase(resetSearchState, (state) => {
        state.searchedBalancesAndAllowances =
          initialState.searchedBalancesAndAllowances
        state.searchStatus = initialState.searchStatus
      })
      .addCase(resetPortfolioState, (state) => {
        state.activeTab = initialState.activeTab
        state.balancesAndAllowances = initialState.balancesAndAllowances
        state.status = initialState.status
        state.error = initialState.error
        state.searchInput = initialState.searchInput
        state.searchedBalancesAndAllowances =
          initialState.searchedBalancesAndAllowances
        state.searchStatus = initialState.searchStatus
        state.poolTokenBalances = initialState.poolTokenBalances
      })
  },
})

export default portfolioSlice.reducer
