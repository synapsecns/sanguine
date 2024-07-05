import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import {
  type ChainPause,
  type BridgeModulePause,
} from '@/components/Maintenance/Maintenance'

export interface MaintenanceState {
  pausedChainsData: ChainPause[]
  pausedModulesData: BridgeModulePause[]
}

const initialState: MaintenanceState = {
  pausedChainsData: null,
  pausedModulesData: null,
}

export const maintenanceSlice = createSlice({
  name: 'maintenance',
  initialState,
  reducers: {
    setPausedChainsData: (state, action: PayloadAction<ChainPause[]>) => {
      state.pausedChainsData = action.payload
    },
    setPausedModulesData: (
      state,
      action: PayloadAction<BridgeModulePause[]>
    ) => {
      state.pausedModulesData = action.payload
    },
  },
})

export const { setPausedChainsData, setPausedModulesData } =
  maintenanceSlice.actions

export default maintenanceSlice.reducer
