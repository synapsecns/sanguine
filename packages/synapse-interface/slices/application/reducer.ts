import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Address } from 'viem'

import { updateLastConnectedAddress, updateLastConnectedTime } from './actions'

export interface ApplicationState {
  lastConnectedTimestamp: number
  lastConnectedAddress: Address
}
const initialState: ApplicationState = {
  lastConnectedTimestamp: undefined,
  lastConnectedAddress: undefined,
}

export const applicationSlice = createSlice({
  name: 'application',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(
        updateLastConnectedAddress,
        (state, action: PayloadAction<Address>) => {
          state.lastConnectedAddress = action.payload
        }
      )
      .addCase(
        updateLastConnectedTime,
        (state, action: PayloadAction<number>) => {
          state.lastConnectedTimestamp = action.payload
        }
      )
  },
})

export default applicationSlice.reducer
