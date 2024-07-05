import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface MaintenanceState {
  pausedChainsData: any[]
  pausedModulesData: any[]
}

const initialState: MaintenanceState = {
  pausedChainsData: null,
  pausedModulesData: null,
}

export const maintenanceSlice = createSlice({
  name: 'maintenance',
  initialState,
  reducers: {
    setPausedChainsData: (state, action: PayloadAction<any[]>) => {
      state.pausedChainsData = action.payload
    },
    setPausedModulesData: (state, action: PayloadAction<any[]>) => {
      state.pausedModulesData = action.payload
    },
  },
})

export const { setPausedChainsData, setPausedModulesData } =
  maintenanceSlice.actions

export default maintenanceSlice.reducer
