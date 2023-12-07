import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Chain, BridgeableToken } from 'types'
import { USDC } from '@/constants/bridgeable'

export interface BridgeState {
  originChain: Chain
  destinationChain: Chain
  originToken: BridgeableToken
  destinationToken: BridgeableToken
  tokens: BridgeableToken[]
}

const initialState: BridgeState = {
  originChain: { id: 42161, name: 'Arbitrum' },
  destinationChain: { id: 137, name: 'Polygon' },
  originToken: USDC,
  destinationToken: USDC,
  tokens: [],
}

export const bridgeSlice = createSlice({
  name: 'bridge',
  initialState,
  reducers: {
    setOriginChain: (state: BridgeState, action: PayloadAction<Chain>) => {
      if (action.payload.id === state.destinationChain.id) {
        state.destinationChain = state.originChain
      }
      state.originChain = action.payload
    },
    setDestinationChain: (state: BridgeState, action: PayloadAction<Chain>) => {
      if (action.payload.id === state.originChain.id) {
        state.originChain = state.destinationChain
      }
      state.destinationChain = action.payload
    },
    setOriginToken: (
      state: BridgeState,
      action: PayloadAction<BridgeableToken>
    ) => {
      state.originToken = action.payload
    },
    setDestinationToken: (
      state: BridgeState,
      action: PayloadAction<BridgeableToken>
    ) => {
      state.destinationToken = action.payload
    },
    setTokens: (
      state: BridgeState,
      action: PayloadAction<BridgeableToken[]>
    ) => {
      state.tokens = action.payload
    },
  },
})

export const {
  setOriginChain,
  setDestinationChain,
  setOriginToken,
  setDestinationToken,
  setTokens,
} = bridgeSlice.actions

export default bridgeSlice.reducer
