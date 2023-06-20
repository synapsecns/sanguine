import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'
import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { ETH } from '@/constants/tokens/master'
import { BridgeQuote, Token } from '@/utils/types'

import { ARBITRUM, ETH as ETHEREUM } from '@/constants/chains/master'

export interface BridgeState {
  fromChainId: number
  supportedFromTokens: Token[]
  supportedFromTokenBalances: any
  toChainId: number
  supportedToTokens: Token[]
  fromToken: Token
  toToken: Token
  fromValue: BigNumber
  bridgeQuote: BridgeQuote
  fromChainIds: number[]
  toChainIds: number[]
  isLoading: boolean
  showFromTokenSlideOver: boolean
  showToTokenSlideOver: boolean
  showChainSlideOver: boolean
}

// How do we update query params based on initial state?
// Additionally how do we set query params based on user input updates?
const initialState: BridgeState = {
  fromChainId: ETHEREUM.id,
  supportedFromTokens: [],
  supportedFromTokenBalances: {},
  toChainId: ARBITRUM.id,
  supportedToTokens: [],
  fromToken: ETH,
  toToken: ETH,
  fromValue: Zero,
  bridgeQuote: EMPTY_BRIDGE_QUOTE,
  fromChainIds: [],
  toChainIds: [],
  isLoading: false,
  showFromTokenSlideOver: false,
  showToTokenSlideOver: false,
  showChainSlideOver: false
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
      state.supportedToTokens = action.payload
    },
    setSupportedFromTokenBalances: (state, action: PayloadAction<{}>) => {
      state.supportedFromTokenBalances = action.payload
    },
    setFromChainIds: (state, action: PayloadAction<number[]>) => {
      state.fromChainIds = action.payload
    },
    setToChainIds: (state, action: PayloadAction<number[]>) => {
      state.toChainIds = action.payload
    },
    updateFromValue: (state, action: PayloadAction<BigNumber>) => {
      state.fromValue = action.payload
    },
    setShowFromTokenSlideOver: (state, action: PayloadAction<boolean>) => {
      state.showFromTokenSlideOver = action.payload
    },
    setShowToTokenSlideOver: (state, action: PayloadAction<boolean>) => {
      state.showToTokenSlideOver = action.payload
    },
    setShowChainSlideOver: (state, action: PayloadAction<boolean>) => {
      state.showChainSlideOver = action.payload
    }
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
  setSupportedFromTokenBalances,
  setFromChainIds,
  setToChainIds,
  setIsLoading,
  setShowChainSlideOver,
  setShowFromTokenSlideOver,
  setShowToTokenSlideOver,
} = bridgeSlice.actions

export default bridgeSlice.reducer
