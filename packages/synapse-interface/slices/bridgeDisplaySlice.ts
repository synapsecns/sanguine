import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface BridgeState {
  showFromTokenListOverlay: boolean
  showToTokenListOverlay: boolean
  showFromChainListOverlay: boolean
  showToChainListOverlay: boolean
  showSettingsSlideOver: boolean
  showDestinationAddress: boolean
}

const initialState: BridgeState = {
  showFromTokenListOverlay: false,
  showToTokenListOverlay: false,
  showFromChainListOverlay: false,
  showToChainListOverlay: false,
  showSettingsSlideOver: false,
  showDestinationAddress: false,
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
    setShowSettingsSlideOver: (state, action: PayloadAction<boolean>) => {
      state.showSettingsSlideOver = action.payload
    },
    setShowDestinationAddress: (state, action: PayloadAction<boolean>) => {
      state.showDestinationAddress = action.payload
    },
  },
})

export const {
  setShowFromChainListOverlay,
  setShowToChainListOverlay,
  setShowFromTokenListOverlay,
  setShowToTokenListOverlay,
  setShowSettingsSlideOver,
  setShowDestinationAddress,
} = bridgeDisplaySlice.actions

export default bridgeDisplaySlice.reducer
