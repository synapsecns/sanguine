import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface BridgeState {
  showFromTokenListOverlay: boolean
  showToTokenListOverlay: boolean
  showFromChainListOverlay: boolean
  showToChainListOverlay: boolean
  showDestinationWarning: boolean
}

const initialState: BridgeState = {
  showFromTokenListOverlay: false,
  showToTokenListOverlay: false,
  showFromChainListOverlay: false,
  showToChainListOverlay: false,
  showDestinationWarning: true,
}

export const bridgeDisplaySlice = createSlice({
  name: 'bridgeDisplay',
  initialState,
  reducers: {
    setShowFromTokenListOverlay: (state, action: PayloadAction<boolean>) => {
      state.showFromTokenListOverlay = action.payload
    },
    setShowToTokenListOverlay: (state, action: PayloadAction<boolean>) => {
      state.showToTokenListOverlay = action.payload
    },
    setShowFromChainListOverlay: (state, action: PayloadAction<boolean>) => {
      state.showFromChainListOverlay = action.payload
    },
    setShowToChainListOverlay: (state, action: PayloadAction<boolean>) => {
      state.showToChainListOverlay = action.payload
    },
    setShowDestinationWarning: (state, action: PayloadAction<boolean>) => {
      state.showDestinationWarning = action.payload
    },
  },
})

export const {
  setShowFromChainListOverlay,
  setShowToChainListOverlay,
  setShowFromTokenListOverlay,
  setShowToTokenListOverlay,
  setShowDestinationWarning,
} = bridgeDisplaySlice.actions

export default bridgeDisplaySlice.reducer
