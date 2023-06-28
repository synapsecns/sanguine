import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'
import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Address } from 'wagmi'

import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { ETH } from '@/constants/tokens/master'
import { ARBITRUM, ETH as ETHEREUM } from '@/constants/chains/master'
import { BridgeQuote, Token } from '@/utils/types'

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
  deadlineMinutes: number | null
  destinationAddress: Address | null
  bridgeTxHashes: string[] | null
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
  deadlineMinutes: null,
  destinationAddress: null,
  bridgeTxHashes: [],
}

export const bridgeSlice = createSlice({
  name: 'bridge',
  initialState,
  reducers: {
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload
    },
    setFromChainId: (state, action: PayloadAction<number>) => {
      if (state.toChainId === action.payload) {
        state.toChainId = state.fromChainId
      }
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
    setDeadlineMinutes: (state, action: PayloadAction<number | null>) => {
      state.deadlineMinutes = action.payload
    },
    setDestinationAddress: (state, action: PayloadAction<Address | null>) => {
      state.destinationAddress = action.payload
    },
    addBridgeTxHash: (state, action: PayloadAction<string>) => {
      state.bridgeTxHashes = [...state.bridgeTxHashes, action.payload]
    },
  },
})

export const tokenDecimalMiddleware =
  ({ getState, dispatch }) =>
  (next) =>
  (action) => {
    // check if the action is setFromToken
    if (action.type === 'bridge/setFromToken') {
      const currentState = getState()

      // if fromValue is 0, no need to adjust it
      if (currentState.bridge.fromValue.isZero()) {
        next(action)
        return
      }

      // get the current fromToken decimal
      const currentDecimal =
        typeof currentState.bridge.fromToken.decimals === 'number'
          ? currentState.bridge.fromToken.decimals
          : currentState.bridge.fromToken.decimals[
              currentState.bridge.fromChainId
            ]

      // get the new token decimal
      const newDecimal =
        typeof action.payload.decimals === 'number'
          ? action.payload.decimals
          : action.payload.decimals[currentState.bridge.fromChainId]

      // calculate the decimal difference
      const decimalDifference = newDecimal - currentDecimal

      if (decimalDifference !== 0) {
        let newFromValue

        if (decimalDifference > 0) {
          // if newDecimal is greater, multiply fromValue by the decimal difference
          newFromValue = currentState.bridge.fromValue.mul(
            BigNumber.from(10).pow(decimalDifference)
          )
        } else {
          // if newDecimal is smaller, divide fromValue by the decimal difference
          newFromValue = currentState.bridge.fromValue.div(
            BigNumber.from(10).pow(Math.abs(decimalDifference))
          )
        }

        // dispatch updateFromValue action to set the new fromValue
        dispatch(updateFromValue(newFromValue))
      }
    }

    // call the next middleware in the line
    next(action)
  }

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
  setDeadlineMinutes,
  setDestinationAddress,
  setIsLoading,
  addBridgeTxHash, // new action
} = bridgeSlice.actions

export default bridgeSlice.reducer
