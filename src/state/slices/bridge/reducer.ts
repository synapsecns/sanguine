import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Chain, TokenMetaData } from 'types'

export interface BridgeState {
  originChain: Chain
  destinationChain: Chain
  originToken: TokenMetaData
  destinationToken: TokenMetaData
  tokens: TokenMetaData[]
}

const initialState: BridgeState = {
  originChain: { id: 42161, name: 'Arbitrum' },
  destinationChain: { id: 137, name: 'Polygon' },
  originToken: {
    tokenAddress: '0xaf88d065e77c8cc2239327c5edb3a432268e5831',
    symbol: 'USDC',
    chainId: 42161,
    decimals: 6,
  },
  destinationToken: {
    tokenAddress: '0x2791bca1f2de4661ed88a30c99a7a9449aa84174',
    symbol: 'USDC',
    chainId: 137,
    decimals: 6,
  },
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
      action: PayloadAction<TokenMetaData>
    ) => {
      state.originToken = action.payload
    },
    setDestinationToken: (
      state: BridgeState,
      action: PayloadAction<TokenMetaData>
    ) => {
      state.destinationToken = action.payload
    },
    setTokens: (state: BridgeState, action: PayloadAction<TokenMetaData[]>) => {
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
