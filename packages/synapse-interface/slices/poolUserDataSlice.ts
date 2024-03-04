import type { PoolUserData, Token } from '@types'
import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import type { Address } from '@wagmi/core'

import { getBalanceData } from '@/utils/actions/getPoolData'

export interface PoolDataState {
  poolUserData: PoolUserData
  isLoading: boolean
}

const initialState: PoolDataState = {
  poolUserData: {
    name: undefined,
    tokens: undefined,
    lpTokenBalance: undefined,
    nativeTokens: undefined,
  },
  isLoading: false,
}

export const fetchPoolUserData = createAsyncThunk(
  'poolUserData/fetch',
  async ({ pool, address }: { pool: Token; address: Address }) => {
    const chainId = pool.chainId
    const poolAddress = pool?.swapAddresses[chainId]

    if (!poolAddress || !pool || !address) {
      return null
    }

    const lpTokenAddress = pool?.addresses[chainId]

    const { tokenBalances, lpTokenBalance } = await getBalanceData({
      pool,
      chainId,
      address,
      lpTokenAddress,
    })

    const tokens = tokenBalances.filter((token) => !token.isLP)

    return {
      name: pool.name,
      tokens,
      lpTokenBalance,
      nativeTokens: pool.nativeTokens,
    }
  }
)

export const poolUserDataSlice = createSlice({
  name: 'poolUserData',
  initialState,
  reducers: {
    resetPoolUserData: () => initialState,
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchPoolUserData.pending, (state) => {
        state.isLoading = true
      })
      .addCase(fetchPoolUserData.fulfilled, (state, action) => {
        state.isLoading = false
        state.poolUserData = action.payload
      })
      .addCase(fetchPoolUserData.rejected, (state) => {
        state.isLoading = false
      })
  },
})

export const { resetPoolUserData } = poolUserDataSlice.actions
export default poolUserDataSlice.reducer
