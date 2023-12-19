import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { executeBridgeTxn } from './hooks'

export enum BridgeTransactionStatus {
  PENDING = 'pending',
  SUCCESS = 'success',
  IDLE = 'idle',
  FAILED = 'failed',
}

export interface BridgeTransactionState {
  txnHash: string
  bridgeTxnStatus: string
  error: any
}

const initialState: BridgeTransactionState = {
  txnHash: null,
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
        (state, action: PayloadAction<string>) => {
          state.txnHash = action.payload
          state.bridgeTxnStatus = BridgeTransactionStatus.SUCCESS
        }
      )
      .addCase(executeBridgeTxn.rejected, (state, action) => {
        state.error = action.payload
        state.txnHash = null
        state.bridgeTxnStatus = BridgeTransactionStatus.FAILED
      })
  },
})

export default bridgeTransactionSlice.reducer
