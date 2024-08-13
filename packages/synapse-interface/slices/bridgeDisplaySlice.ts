import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface BridgeState {
  showDestinationWarning: boolean
  isDestinationWarningAccepted: boolean
  showSettingsSlideOver: boolean
  showDestinationAddress: boolean
}

const initialState: BridgeState = {
  showDestinationWarning: false,
  isDestinationWarningAccepted: false,
  showSettingsSlideOver: false,
  showDestinationAddress: false,
}

export const bridgeDisplaySlice = createSlice({
  name: 'bridgeDisplay',
  initialState,
  reducers: {
    setShowDestinationWarning: (state, action: PayloadAction<boolean>) => {
      state.showDestinationWarning = action.payload
    },
    setIsDestinationWarningAccepted: (
      state,
      action: PayloadAction<boolean>
    ) => {
      state.isDestinationWarningAccepted = action.payload
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
  setShowDestinationWarning,
  setIsDestinationWarningAccepted,
  setShowSettingsSlideOver,
  setShowDestinationAddress,
} = bridgeDisplaySlice.actions

export default bridgeDisplaySlice.reducer
