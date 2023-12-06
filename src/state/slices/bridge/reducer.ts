import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Chain, TokenMetaData } from 'types'

export interface BridgeState {
  originChain: Chain
  destinationChain: Chain
  originToken: TokenMetaData
  destinationToken: TokenMetaData
}

const initialState: BridgeState = {
  originChain: { id: 1, name: 'Ethereum' },
  destinationChain: { id: 42161, name: 'Arbitrum' },
  originToken: {
    tokenAddress: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
    symbol: 'USDC',
    chainId: 1,
    decimals: 6,
  },
  destinationToken: {
    tokenAddress: '0xaf88d065e77c8cc2239327c5edb3a432268e5831',
    symbol: 'USDC',
    chainId: 42161,
    decimals: 6,
  },
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
  },
})

export const {
  setOriginChain,
  setDestinationChain,
  setOriginToken,
  setDestinationToken,
} = bridgeSlice.actions

export default bridgeSlice.reducer
