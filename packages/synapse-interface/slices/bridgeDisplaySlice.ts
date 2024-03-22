import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface BridgeState {
  showSettingsSlideOver: boolean
  showDestinationAddress: boolean
}

const initialState: BridgeState = {
  showSettingsSlideOver: false,
  showDestinationAddress: false,
}

export const bridgeDisplaySlice = createSlice({
  name: 'bridgeDisplay',
  initialState,
  reducers: {
    setShowSettingsSlideOver: (state, action: PayloadAction<boolean>) => {
      state.showSettingsSlideOver = action.payload
    },
    setShowDestinationAddress: (state, action: PayloadAction<boolean>) => {
      state.showDestinationAddress = action.payload
    },
  },
})

export const { setShowSettingsSlideOver, setShowDestinationAddress } =
  bridgeDisplaySlice.actions

export default bridgeDisplaySlice.reducer
