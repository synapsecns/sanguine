import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { executeApproveTxn } from './hooks'

export enum ApproveTransactionStatus {
  PENDING = 'pending',
  SUCCESS = 'success',
  IDLE = 'idle',
  FAILED = 'failed',
}

type ApprovalReceipt = {
  blockHash: string
  blockNumber: number
  contractAddress: string
  from: string
  hash: string
  index: number
  logsBloom: string
  to: string
  type: number
}

export interface BridgeTransactionState {
  approveTxnStatus: string
  receipt: ApprovalReceipt
  error: any
}

const initialState: BridgeTransactionState = {
  approveTxnStatus: ApproveTransactionStatus.IDLE,
  receipt: null,
  error: null,
}

export const approveTransactionSlice = createSlice({
  name: 'approveTransaction',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(executeApproveTxn.pending, (state) => {
        state.approveTxnStatus = ApproveTransactionStatus.PENDING
      })
      .addCase(
        executeApproveTxn.fulfilled,
        (state, action: PayloadAction<ApprovalReceipt>) => {
          state.receipt = action.payload
          state.approveTxnStatus = ApproveTransactionStatus.SUCCESS
        }
      )
      .addCase(executeApproveTxn.rejected, (state, action) => {
        state.error = action.payload
        state.approveTxnStatus = ApproveTransactionStatus.FAILED
      })
  },
})

export default approveTransactionSlice.reducer
