import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { estimateFeesPerGas } from '@wagmi/core'

import { wagmiConfig } from '@/wagmiConfig'

export interface GasDataState {
  chainId: number
  gasLimit: bigint
  gasData: {
    gasPrice?: bigint
    maxFeePerGas: bigint
    maxPriorityFeePerGas: bigint
    formatted: {
      gasPrice?: string
      maxFeePerGas: string
      maxPriorityFeePerGas: string
    }
  }
  isLoadingGasLimit: boolean
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
    return { gasData, chainId }
  }
)

const initialState: GasDataState = {
  chainId: null,
  gasLimit: null,
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
  isLoadingGasLimit: false,
  isLoadingGasData: true,
}

export const gasDataSlice = createSlice({
  name: 'gasData',
  initialState,
  reducers: {
    setIsLoadingGasLimit: (state: GasDataState, action) => {
      state.isLoadingGasLimit = action.payload
    },
    setGasLimit: (state: GasDataState, action) => {
      state.gasLimit = action.payload
      state.isLoadingGasLimit = false
    },
    resetGasLimit: (state: GasDataState) => {
      state.gasLimit = initialState.gasLimit
      state.isLoadingGasLimit = initialState.isLoadingGasLimit
    },
    resetGasData: (state: GasDataState) => {
      state.chainId = initialState.chainId
      state.gasData = initialState.gasData
    },
  },
  extraReducers: (builder) => {
    builder.addCase(fetchGasData.pending, (state) => {
      state.isLoadingGasData = true
    })
    builder
      .addCase(fetchGasData.fulfilled, (state, action) => {
        const { chainId, gasData } = action.payload
        state.isLoadingGasData = false
        state.chainId = chainId
        state.gasData = gasData
      })
      .addCase(fetchGasData.rejected, (state) => {
        state.isLoadingGasData = false
        console.error('Error fetching gas data')
      })
  },
})

export const {
  setGasLimit,
  resetGasLimit,
  resetGasData,
  setIsLoadingGasLimit,
} = gasDataSlice.actions

export default gasDataSlice.reducer
