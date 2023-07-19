import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'

import {
  fetchPortfolioBalances,
  NetworkTokenBalancesAndAllowances,
} from '@/utils/hooks/usePortfolioBalances'

export const fetchAndStorePortfolioBalances = createAsyncThunk(
  'portfolio/fetchAndStorePortfolioBalances',
  async (address: string) => {
    console.log('this got hit 2')
    const portfolioData = await fetchPortfolioBalances(address)
    console.log('portfolioData: ', portfolioData)
    return portfolioData
  }
)

export interface PortfolioState {
  balancesAndAllowances: NetworkTokenBalancesAndAllowances
}

export const portfolioSlice = createSlice({
  name: 'portfolio',
  initialState: {
    balancesAndAllowances: {},
    status: 'idle',
    error: null,
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchAndStorePortfolioBalances.pending, (state) => {
        state.status = 'loading'
      })
      .addCase(fetchAndStorePortfolioBalances.fulfilled, (state, action) => {
        state.status = 'succeeded'
        state.balancesAndAllowances = action.payload.balancesAndAllowances
      })
      .addCase(fetchAndStorePortfolioBalances.rejected, (state, action) => {
        state.status = 'failed'
        state.error = action.error.message
      })
  },
})

export default portfolioSlice.reducer
