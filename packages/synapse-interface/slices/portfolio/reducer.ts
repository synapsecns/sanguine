import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Address } from 'viem'

import { Token } from '@/utils/types'
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
    updateSingleTokenAllowance: (
      state,
      action: PayloadAction<{
        chainId: number
        allowance: bigint
        spender: Address
        token: Token
      }>
    ) => {
      const { chainId, allowance, spender, token } = action.payload

      state.balancesAndAllowances[chainId].forEach(
        (t: TokenWithBalanceAndAllowances) => {
          if (t.tokenAddress === token.addresses[chainId]) {
            t.allowances[spender] = allowance
          }
        }
      )
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
  },
})

export default portfolioSlice.reducer
