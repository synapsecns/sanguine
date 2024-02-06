import { createSlice } from '@reduxjs/toolkit'

import { executeBridgeTxn } from './hooks'

export enum BridgeTransactionStatus {
  PENDING = 'pending',
  SUCCESS = 'success',
  IDLE = 'idle',
  FAILED = 'failed',
}

export interface BridgeTransactionState {
  txHash: string
  bridgeModuleName: string
  originAmount: string
  originTokenSymbol: string
  originChainId: number
  destinationChainId: number
  estimatedTime: number
  timestamp: number
  bridgeTxnStatus: string
  error: any
}

const initialState: BridgeTransactionState = {
  txHash: null,
  bridgeModuleName: null,
  originAmount: null,
  originTokenSymbol: null,
  originChainId: null,
  destinationChainId: null,
  estimatedTime: null,
  timestamp: null,
  bridgeTxnStatus: BridgeTransactionStatus.IDLE,
  error: null,
}

export const bridgeTransactionSlice = createSlice({
  name: 'bridgeTransaction',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(executeBridgeTxn.pending, (state) => {
        state.bridgeTxnStatus = BridgeTransactionStatus.PENDING
      })
      .addCase(
        executeBridgeTxn.fulfilled,
        (
          state,
          {
            payload: {
              txHash,
              bridgeModuleName,
              parsedOriginAmount,
              originTokenSymbol,
              originChainId,
              destinationChainId,
              estimatedTime,
              timestamp,
            },
          }
        ) => {
          state.bridgeTxnStatus = BridgeTransactionStatus.SUCCESS
          state.txHash = txHash
          state.bridgeModuleName = bridgeModuleName
          state.originAmount = parsedOriginAmount
          state.originTokenSymbol = originTokenSymbol
          state.originChainId = originChainId
          state.destinationChainId = destinationChainId
          state.estimatedTime = estimatedTime
          state.timestamp = timestamp
        }
      )
      .addCase(executeBridgeTxn.rejected, (state, action) => {
        state.error = action.payload
        state.txHash = null
        state.bridgeModuleName = null
        state.originAmount = null
        state.originTokenSymbol = null
        state.originChainId = null
        state.destinationChainId = null
        state.estimatedTime = null
        state.timestamp = null
        state.bridgeTxnStatus = BridgeTransactionStatus.FAILED
      })
  },
})

export default bridgeTransactionSlice.reducer
