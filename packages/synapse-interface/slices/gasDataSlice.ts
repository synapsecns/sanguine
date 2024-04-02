import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'

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
