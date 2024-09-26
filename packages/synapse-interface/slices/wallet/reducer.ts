import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface WalletState {
  isWalletPending: boolean
}

export const initialState: WalletState = {
  isWalletPending: false,
}

export const walletSlice = createSlice({
  name: 'wallet',
  initialState,
  reducers: {
    setIsWalletPending: (state, action: PayloadAction<boolean>) => {
      state.isWalletPending = action.payload
    },
  },
})

export const { setIsWalletPending } = walletSlice.actions

export default walletSlice.reducer
