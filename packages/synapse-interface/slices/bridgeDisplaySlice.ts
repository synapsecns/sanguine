import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface BridgeState {
  showFromTokenSlideOver: boolean
  showToTokenSlideOver: boolean
  showFromChainSlideOver: boolean
  showToChainSlideOver: boolean
  showSettingsSlideOver: boolean
  showDestinationAddress: boolean
}

const initialState: BridgeState = {
  showFromTokenSlideOver: false,
  showToTokenSlideOver: false,
  showFromChainSlideOver: false,
  showToChainSlideOver: false,
  showSettingsSlideOver: false,
  showDestinationAddress: false,
}

export const bridgeDisplaySlice = createSlice({
  name: 'bridgeDisplay',
  initialState,
  reducers: {
    setShowFromTokenSlideOver: (state, action: PayloadAction<boolean>) => {
      state.showFromTokenSlideOver = action.payload
    },
    setShowToTokenSlideOver: (state, action: PayloadAction<boolean>) => {
      state.showToTokenSlideOver = action.payload
    },
    setShowFromChainSlideOver: (state, action: PayloadAction<boolean>) => {
      state.showFromChainSlideOver = action.payload
    },
    setShowToChainSlideOver: (state, action: PayloadAction<boolean>) => {
      state.showToChainSlideOver = action.payload
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
  setShowFromChainSlideOver,
  setShowToChainSlideOver,
  setShowFromTokenSlideOver,
  setShowToTokenSlideOver,
  setShowSettingsSlideOver,
  setShowDestinationAddress,
} = bridgeDisplaySlice.actions

export default bridgeDisplaySlice.reducer
