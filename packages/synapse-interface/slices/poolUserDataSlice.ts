import { PoolUserData, Token } from '@types'
import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { type Address } from 'viem'

import { getBalanceData } from '@/utils/actions/getPoolData'
import { getStakedBalance } from '@/utils/actions/getStakedBalance'

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
    stakedBalance: {
      amount: 0n,
      reward: 0n,
    },
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

    const stakedBalance = await getStakedBalance(
      address,
      pool.chainId,
      pool.poolId[pool.chainId],
      pool
    )

    return {
      name: pool.name,
      tokens,
      lpTokenBalance,
      stakedBalance,
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
