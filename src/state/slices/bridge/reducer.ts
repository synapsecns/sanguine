import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface BridgeState {
  originChainId: number
  destinationChainId: number
}

const initialState: BridgeState = {
  originChainId: 1,
  destinationChainId: 42161,
}

export const bridgeSlice = createSlice({
  name: 'bridge',
  initialState,
  reducers: {},
})

export default bridgeSlice.reducer
