import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface SwapDisplayState {
  showSwapChainListOverlay: boolean
  showSwapFromTokenListOverlay: boolean
  showSwapToTokenListOverlay: boolean
}

const initialState: SwapDisplayState = {
  showSwapChainListOverlay: false,
  showSwapFromTokenListOverlay: false,
  showSwapToTokenListOverlay: false,
}

export const swapDisplaySlice = createSlice({
  name: 'swapDisplay',
  initialState,
  reducers: {
    setShowSwapFromTokenListOverlay: (
      state,
      action: PayloadAction<boolean>
    ) => {
      state.showSwapFromTokenListOverlay = action.payload
    },
    setShowSwapToTokenListOverlay: (state, action: PayloadAction<boolean>) => {
      state.showSwapToTokenListOverlay = action.payload
    },
    setShowSwapChainListOverlay: (state, action: PayloadAction<boolean>) => {
      state.showSwapChainListOverlay = action.payload
    },
  },
})

export const {
  setShowSwapChainListOverlay,
  setShowSwapFromTokenListOverlay,
  setShowSwapToTokenListOverlay,
} = swapDisplaySlice.actions

export default swapDisplaySlice.reducer
