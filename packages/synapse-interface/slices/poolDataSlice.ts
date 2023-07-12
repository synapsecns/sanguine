import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { Token } from '@types'

import { POOL_BY_ROUTER_INDEX } from '@/constants/tokens'
import { getPoolTokenInfoArr, getTokenBalanceInfo } from '@/utils/poolDataFuncs'
import { getPoolApyData } from '@/utils/actions/getPoolApyData'
import { getBalanceData } from '@/utils/actions/getPoolData'
import { getCorePoolData } from '@/utils/actions/getCorePoolData'
import { getAvaxPrice, getEthPrice } from '@/utils/actions/getPrices'

export interface PoolDataState {
  pool: Token | undefined
  // poolData: {
  //   name: string
  //   tokens: Token[]
  //   totalLocked: number
  //   totalLockedUSD: number
  //   virtualPrice: string
  //   swapFee: string
  // }
  poolData: any
  poolAPYData: any
  isLoading: boolean
}

const initialState: PoolDataState = {
  pool: undefined,
  poolData: {
    name: undefined,
    tokens: undefined,
    totalLocked: undefined,
    totalLockedUSD: undefined,
    virtualPrice: undefined,
    swapFee: undefined,
  },
  poolAPYData: {
    fullCompoundedApyStr: undefined,
    weeklyAPR: undefined,
    yearlyAPRUnvested: undefined,
  },
  isLoading: false,
}

export const fetchPoolData = createAsyncThunk(
  'poolData/fetch',
  async ({ poolName }: { poolName: string }) => {
    const pool = POOL_BY_ROUTER_INDEX[poolName]
    const chainId = pool.chainId

    const poolAddress = pool?.swapAddresses[chainId]

    if (!pool || !poolAddress) {
      return null
    }

    const lpTokenAddress = pool?.addresses[chainId]

    const { tokenBalances } = await getBalanceData({
      pool,
      chainId,
      address: poolAddress,
      lpTokenAddress,
    })

    const { swapFee, virtualPrice } = await getCorePoolData(
      poolAddress,
      chainId
    )

    const ethPrice = await getEthPrice()
    const avaxPrice = await getAvaxPrice()

    const {
      tokenBalancesSum,
      tokenBalancesUSD,
    }: { tokenBalancesSum: number; tokenBalancesUSD: number } =
      getTokenBalanceInfo({
        tokenBalances: tokenBalances
          .filter((t) => !t.isLP)
          .map((t) => t.balance),
        prices: {
          ethPrice,
          avaxPrice,
        },
        poolType: pool?.poolType,
      })

    const poolTokensMatured = getPoolTokenInfoArr({
      tokenBalances: tokenBalances.filter((t) => !t.isLP),
      tokenBalancesSum,
    })

    const poolAPYData = await getPoolApyData(chainId, pool)

    // prob should split into multiple thunks to handle loading
    // separately
    return {
      pool,
      poolAPYData,
      poolData: {
        name: pool.name,
        tokens: poolTokensMatured,
        totalLocked: tokenBalancesSum,
        totalLockedUSD: tokenBalancesUSD,
        virtualPrice,
        swapFee,
      },
    }
  }
)

export const poolDataSlice = createSlice({
  name: 'poolData',
  initialState,
  reducers: {
    resetPoolData: () => initialState,
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchPoolData.pending, (state) => {
        state.isLoading = true
      })
      .addCase(fetchPoolData.fulfilled, (state, action) => {
        state.isLoading = false
        state.pool = action.payload.pool
        state.poolData = action.payload.poolData
        state.poolAPYData = action.payload.poolAPYData
      })
      .addCase(fetchPoolData.rejected, (state) => {
        state.isLoading = false
      })
  },
})

export const { resetPoolData } = poolDataSlice.actions
export const selectPool = (state) => state.poolData.pool

export default poolDataSlice.reducer
