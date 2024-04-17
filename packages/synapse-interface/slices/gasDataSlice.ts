import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { estimateFeesPerGas } from '@wagmi/core'

import { wagmiConfig } from '@/wagmiConfig'

export interface GasDataState {
  gasData: {
    gasPrice: bigint
    maxFeePerGas: bigint
    maxPriorityFeePerGas: bigint
    formatted: {
      gasPrice: string
      maxFeePerGas: string
      maxPriorityFeePerGas: string
    }
  }
  isLoadingGasData: boolean
}

const getGasData = async (chainId: number) => {
  const feeData = await estimateFeesPerGas(wagmiConfig, {
    chainId,
    formatUnits: 'gwei',
  })

  return feeData
}

export const fetchGasData = createAsyncThunk(
  'gasData/fetchGasData',
  async (chainId: number) => {
    const gasData = await getGasData(chainId)
    return gasData
  }
)

const initialState: GasDataState = {
  gasData: {
    gasPrice: null,
    maxFeePerGas: null,
    maxPriorityFeePerGas: null,
    formatted: {
      gasPrice: null,
      maxFeePerGas: null,
      maxPriorityFeePerGas: null,
    },
  },
  isLoadingGasData: true,
}

export const gasDataSlice = createSlice({
  name: 'gasData',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchGasData.pending, (state) => {
      state.isLoadingGasData = true
    })
    builder
      .addCase(fetchGasData.fulfilled, (state, action) => {
        state.isLoadingGasData = false
        state.gasData = action.payload
      })
      .addCase(fetchGasData.rejected, (state) => {
        state.isLoadingGasData = false
        console.error('Error fetching gas data')
      })
  },
})

export default gasDataSlice.reducer
