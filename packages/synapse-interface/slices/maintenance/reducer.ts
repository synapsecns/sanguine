import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import {
  type ChainPause,
  type BridgeModulePause,
} from '@/components/Maintenance/Maintenance'

export interface MaintenanceState {
  pausedChainsData: ChainPause[]
  pausedModulesData: BridgeModulePause[]
  isFetching: boolean
}

const initialState: MaintenanceState = {
  pausedChainsData: null,
  pausedModulesData: null,
  isFetching: false,
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
    setIsFetching: (state, action: PayloadAction<boolean>) => {
      state.isFetching = action.payload
    },
  },
})

export const { setPausedChainsData, setPausedModulesData, setIsFetching } =
  maintenanceSlice.actions

export default maintenanceSlice.reducer
