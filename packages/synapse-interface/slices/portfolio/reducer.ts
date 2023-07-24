import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { PortfolioTabs } from './actions'
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

export enum FetchState {
  IDLE = 'idle',
  LOADING = 'loading',
  VALID = 'valid',
  INVALID = 'invalid',
}

export interface PortfolioState {
  activeTab: PortfolioTabs
  balancesAndAllowances: NetworkTokenBalancesAndAllowances
  status: FetchState
  error?: string
}

const initialState: PortfolioState = {
  activeTab: PortfolioTabs.HOME,
  balancesAndAllowances: {},
  status: FetchState.IDLE,
  error: null,
}

export const portfolioSlice = createSlice({
  name: 'portfolio',
  initialState,
  reducers: {
    setActiveTab: (state, action: PayloadAction<PortfolioTabs>) => {
      state.activeTab = action.payload
    },
  },
  extraReducers: (builder) => {
    builder
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

          // Update the existing balancesAndAllowances object
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
  },
})

export default portfolioSlice.reducer
