import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'
import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { ETH } from '@/constants/tokens/master'
import { BridgeQuote, Token } from '@/utils/types'

export interface BridgeState {
  fromChainId: number
  supportedFromTokens: Token[]
  toChainId: number
  supportedToTokens: Token[]
  fromToken: Token
  toToken: Token
  fromValue: BigNumber
  bridgeQuote: BridgeQuote
  isLoading: boolean
}

const initialState: BridgeState = {
  fromChainId: 1,
  supportedFromTokens: [],
  toChainId: 10,
  supportedToTokens: [],
  fromToken: ETH,
  toToken: ETH,
  fromValue: Zero,
  bridgeQuote: EMPTY_BRIDGE_QUOTE,
  isLoading: false,
}

export const bridgeSlice = createSlice({
  name: 'bridge',
  initialState,
  reducers: {
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload
    },
    setFromChainId: (state, action: PayloadAction<number>) => {
      state.fromChainId = action.payload
    },
    setToChainId: (state, action: PayloadAction<number>) => {
      state.toChainId = action.payload
    },
    setFromToken: (state, action: PayloadAction<Token>) => {
      state.fromToken = action.payload
    },
    setToToken: (state, action: PayloadAction<Token>) => {
      state.toToken = action.payload
    },
    setBridgeQuote: (state, action: PayloadAction<BridgeQuote>) => {
      state.bridgeQuote = action.payload
    },
    setSupportedFromTokens: (state, action: PayloadAction<Token[]>) => {
      state.supportedFromTokens = action.payload
    },
    setSupportedToTokens: (state, action: PayloadAction<Token[]>) => {
      state.supportedFromTokens = action.payload
    },
    updateFromValue: (state, action: PayloadAction<BigNumber>) => {
      state.fromValue = action.payload
    },
  },
})

export const {
  setBridgeQuote,
  setFromChainId,
  setToChainId,
  setFromToken,
  setToToken,
  updateFromValue,
  setSupportedFromTokens,
  setSupportedToTokens,
  setIsLoading,
} = bridgeSlice.actions

export default bridgeSlice.reducer
