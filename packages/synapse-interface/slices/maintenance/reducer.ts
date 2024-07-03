import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface MaintenanceState {
  pausedChainData: any
  pausedModuleData: any
}

const initialState: MaintenanceState = {
  pausedChainData: null,
  pausedModuleData: null,
}

export const maintenanceSlice = createSlice({
  name: 'maintenance',
  initialState,
  reducers: {
    setPausedChainData: (state, action: PayloadAction<any>) => {
      state.pausedChainData = action.payload
    },
    setPausedModuleData: (state, action: PayloadAction<any>) => {
      state.pausedModuleData = action.payload
    },
  },
})

export const { setPausedChainData, setPausedModuleData } =
  maintenanceSlice.actions

export default maintenanceSlice.reducer
