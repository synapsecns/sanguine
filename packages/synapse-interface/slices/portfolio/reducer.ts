import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Address } from 'viem'

import {
  PortfolioTabs,
  FetchState,
  setActiveTab,
  updateSingleTokenAllowance,
  resetPortfolioState,
  typeSearchInput,
} from './actions'
import {
  fetchAndStorePortfolioBalances,
  fetchAndStoreSingleNetworkPortfolioBalances,
  fetchAndStoreSingleTokenAllowance,
  fetchAndStoreSingleTokenBalance,
} from './hooks'
import {
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowances,
} from '@/utils/actions/fetchPortfolioBalances'

export interface PortfolioState {
  activeTab: PortfolioTabs
  balancesAndAllowances: NetworkTokenBalancesAndAllowances
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
      .addCase(fetchAndStorePortfolioBalances.pending, (state) => {
        state.status = FetchState.LOADING
      })
      .addCase(fetchAndStorePortfolioBalances.fulfilled, (state, action) => {
        state.status = FetchState.VALID
        state.balancesAndAllowances = action.payload.balancesAndAllowances
      })
      .addCase(fetchAndStorePortfolioBalances.rejected, (state, action) => {
        state.status = FetchState.INVALID
        state.error = action.error.message
      })
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
      .addCase(typeSearchInput, (state, { payload: { searchInput } }) => {
        state.searchInput = searchInput
      })
      .addCase(resetPortfolioState, (state) => {
        state.activeTab = initialState.activeTab
        state.balancesAndAllowances = initialState.balancesAndAllowances
        state.searchInput = initialState.searchInput
        state.status = initialState.status
        state.error = initialState.error
      })
  },
})

export default portfolioSlice.reducer
