import _ from 'lodash'
import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Address } from 'wagmi'
import * as ALL_TOKENS from '@constants/tokens/master'

import { USDC } from '@/constants/tokens/master'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { ARBITRUM, ETH as ETHEREUM } from '@/constants/chains/master'
import { BridgeQuote, Token } from '@/utils/types'
import {
  extractFirstChainIdBySymbol,
  generateRoutePossibilities,
  getPossibleFromChainIds,
  getPossibleFromTokensByFromChainId,
  getPossibleToTokensByFromTokenAndToChainId,
} from '@/utils/generateRoutePossibilities'

const fromTokenDefaults = extractFirstChainIdBySymbol()

const getToken = (tokenAndChainId: string) => {
  if (tokenAndChainId) {
    const symbol = tokenAndChainId.split('-')[0]
    return ALL_TOKENS[symbol]
  } else {
    return null
  }
}

const getSymbol = (tokenAndChainId: string): string => {
  return tokenAndChainId.split('-')[0]
}

export interface BridgeState {
  fromChainId: number
  fromToken: Token
  toChainId: number
  toToken: Token
  fromChainIds: number[]
  toChainIds: number[]
  fromTokens: Token[]
  toTokens: Token[]

  fromValue: string
  bridgeQuote: BridgeQuote
  isLoading: boolean
  deadlineMinutes: number | null
  destinationAddress: Address | null
  bridgeTxHashes: string[] | null
}

const initialFromTokens = _.uniq(
  getPossibleFromTokensByFromChainId(ETHEREUM.id).map(getSymbol)
).map((symbol) => ALL_TOKENS[symbol])

const initialToTokens = _.uniq(
  getPossibleToTokensByFromTokenAndToChainId('USDC-1', ARBITRUM.id).map(
    getSymbol
  )
).map((symbol) => ALL_TOKENS[symbol])

const initialFromChainIds = getPossibleFromChainIds()
const initialToChainIds = getPossibleFromChainIds()

// How do we update query params based on initial state?
// Additionally how do we set query params based on user input updates?
const initialState: BridgeState = {
  fromChainId: ETHEREUM.id,
  fromToken: USDC,
  toChainId: ARBITRUM.id,
  toToken: USDC,
  fromChainIds: initialFromChainIds,
  toChainIds: initialToChainIds,
  fromTokens: initialFromTokens,
  toTokens: initialToTokens,

  fromValue: '',
  bridgeQuote: EMPTY_BRIDGE_QUOTE,
  isLoading: false,
  deadlineMinutes: null,
  destinationAddress: null,
  bridgeTxHashes: [],
}

/*

Notes on default sorting

fromTokens: separateAndSortTokensWithBalances, sortTokensByBalanceDescending
toTokens: sortTokensByPriorityRankAndAlpha
fromChainIds: 
toChainIds: 

handling default toChain, toToken

are swapExceptions still vaild?


*/

export const bridgeSlice = createSlice({
  name: 'bridge',
  initialState,
  reducers: {
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload
    },
    setFromChainId: (state, action: PayloadAction<number>) => {
      const incomingFromChainId = action.payload
      const {
        fromChainId,
        fromToken,
        toChainId,
        toToken,
        fromChainIds,
        fromTokens,
        toChainIds,
        toTokens,
      } = generateRoutePossibilities({
        fromChainId: incomingFromChainId,
        fromToken:
          incomingFromChainId &&
          state.fromToken &&
          `${state.fromToken.routeSymbol}-${incomingFromChainId}`,
        toChainId: incomingFromChainId && state.toChainId,
        toToken:
          incomingFromChainId &&
          state.toToken &&
          state.toChainId &&
          `${state.toToken.routeSymbol}-${state.toChainId}`,
      })

      state.fromChainId = fromChainId
      state.fromToken = getToken(fromToken)
      state.toChainId = toChainId
      state.toToken = getToken(toToken)
      state.fromChainIds = fromChainIds
      state.fromTokens = _.uniq(fromTokens.map(getSymbol)).map(
        (symbol: string) => ALL_TOKENS[symbol]
      )
      state.toChainIds = toChainIds
      state.toTokens = _.uniq(toTokens.map(getSymbol)).map(
        (symbol: string) => ALL_TOKENS[symbol]
      )
    },
    setFromToken: (state, action: PayloadAction<Token>) => {
      const incomingFromToken = action.payload

      const stringifiedFromToken = state.fromChainId
        ? `${incomingFromToken.routeSymbol}-${state.fromChainId}`
        : `${incomingFromToken.routeSymbol}-${
            fromTokenDefaults[incomingFromToken.routeSymbol]
          }`

      const {
        fromChainId,
        fromToken,
        toChainId,
        toToken,
        fromChainIds,
        fromTokens,
        toChainIds,
        toTokens,
      } = generateRoutePossibilities({
        fromChainId: state.fromChainId,
        fromToken: stringifiedFromToken,
        toChainId: state.toChainId,
        toToken: null,
      })

      state.fromChainId = fromChainId
      state.fromToken = getToken(fromToken)
      state.toChainId = toChainId
      state.toToken = getToken(toToken)
      state.fromChainIds = fromChainIds
      state.fromTokens = _.uniq(fromTokens.map(getSymbol)).map(
        (symbol: string) => ALL_TOKENS[symbol]
      )
      state.toChainIds = toChainIds
      state.toTokens = _.uniq(toTokens.map(getSymbol)).map(
        (symbol: string) => ALL_TOKENS[symbol]
      )
    },
    setToChainId: (state, action: PayloadAction<number>) => {
      const incomingToChainId = action.payload

      const {
        fromChainId,
        fromToken,
        toChainId,
        toToken,
        fromChainIds,
        fromTokens,
        toChainIds,
        toTokens,
      } = generateRoutePossibilities({
        fromChainId: state.fromChainId,
        fromToken:
          state.fromToken &&
          `${state.fromToken.routeSymbol}-${state.fromChainId}`,
        toChainId: incomingToChainId,
        toToken: null,
      })

      state.fromChainId = fromChainId
      state.fromToken = getToken(fromToken)
      state.toChainId = toChainId
      state.toToken = getToken(toToken)
      state.fromChainIds = fromChainIds
      state.fromTokens = _.uniq(fromTokens.map(getSymbol)).map(
        (symbol: string) => ALL_TOKENS[symbol]
      )
      state.toChainIds = toChainIds
      state.toTokens = _.uniq(toTokens.map(getSymbol)).map(
        (symbol: string) => ALL_TOKENS[symbol]
      )
    },
    setToToken: (state, action: PayloadAction<Token>) => {
      const incomingToToken = action.payload
      const {
        fromChainId,
        fromToken,
        toChainId,
        toToken,
        fromChainIds,
        fromTokens,
        toChainIds,
        toTokens,
      } = generateRoutePossibilities({
        fromChainId: state.fromChainId,
        fromToken:
          state.fromToken &&
          `${state.fromToken.routeSymbol}-${state.fromChainId}`,
        toChainId: state.toChainId,
        toToken: `${incomingToToken.routeSymbol}-${state.toChainId}`,
      })

      state.fromChainId = fromChainId
      state.fromToken = getToken(fromToken)
      state.toChainId = toChainId
      state.toToken = getToken(toToken)
      state.fromChainIds = fromChainIds
      state.fromTokens = _.uniq(fromTokens.map(getSymbol)).map(
        (symbol: string) => ALL_TOKENS[symbol]
      )
      state.toChainIds = toChainIds
      state.toTokens = _.uniq(toTokens.map(getSymbol)).map(
        (symbol: string) => ALL_TOKENS[symbol]
      )
    },
    setBridgeQuote: (state, action: PayloadAction<BridgeQuote>) => {
      state.bridgeQuote = action.payload
    },
    updateFromValue: (state, action: PayloadAction<string>) => {
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

export const {
  setBridgeQuote,
  setFromChainId,
  setToChainId,
  setFromToken,
  setToToken,
  updateFromValue,
  setDeadlineMinutes,
  setDestinationAddress,
  setIsLoading,
  addBridgeTxHash,
} = bridgeSlice.actions

export default bridgeSlice.reducer
