import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { fetchFeeData } from '@wagmi/core'

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
}

const getFeeData = async (chainId: number) => {
  const feeData = await fetchFeeData({
    chainId,
  })

  return feeData
}

export const fetchGasData = createAsyncThunk(
  'gasData/fetchGasData',
  async (chainId: number) => {
    const gasData = await getFeeData(chainId)
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
}
