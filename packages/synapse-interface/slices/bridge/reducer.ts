import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Address } from 'wagmi'

import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { BridgeQuote, Token } from '@/utils/types'
import {
  getRoutePossibilities,
  getSymbol,
} from '@/utils/routeMaker/generateRoutePossibilities'
import { getFromChainIds } from '@/utils/routeMaker/getFromChainIds'
import { getFromTokens } from '@/utils/routeMaker/getFromTokens'
import { getToChainIds } from '@/utils/routeMaker/getToChainIds'
import { getToTokens } from '@/utils/routeMaker/getToTokens'
import { findTokenByRouteSymbol } from '@/utils/findTokenByRouteSymbol'
import {
  resetFetchedBridgeQuotes,
  resetBridgeInputs,
  updateDebouncedFromValue,
  updateDebouncedToTokensFromValue,
} from './actions'
import { fetchAndStoreBridgeQuotes } from './hooks'
import { BridgeQuoteResponse } from '@/utils/actions/fetchBridgeQuotes'
import { findValidToken } from '@/utils/findValidToken'
import { FetchState } from '../portfolio/actions'

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
  debouncedFromValue: string
  debouncedToTokensFromValue: string
  bridgeQuote: BridgeQuote
  toTokensBridgeQuotes: BridgeQuoteResponse[]
  toTokensBridgeQuotesStatus: FetchState
  isLoading: boolean
  deadlineMinutes: number | null
  destinationAddress: Address | null
}

const {
  fromChainId,
  fromToken,
  toChainId,
  toToken,
  fromChainIds,
  fromTokens,
  toChainIds,
  toTokens,
} = getRoutePossibilities({
  fromChainId: null,
  fromToken: null,
  toChainId: null,
  toToken: null,
})

export const initialState: BridgeState = {
  fromChainId,
  fromToken,
  toChainId,
  toToken,
  fromChainIds,
  toChainIds,
  fromTokens,
  toTokens,

  fromValue: '',
  debouncedFromValue: '',
  debouncedToTokensFromValue: '',
  bridgeQuote: EMPTY_BRIDGE_QUOTE,
  toTokensBridgeQuotes: [],
  toTokensBridgeQuotesStatus: FetchState.IDLE,
  isLoading: false,
  deadlineMinutes: null,
  destinationAddress: null,
}

export const bridgeSlice = createSlice({
  name: 'bridge',
  initialState,
  reducers: {
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload
    },
    setFromChainId: (state, action: PayloadAction<number>) => {
      const incomingFromChainId = action.payload

      const validFromTokens = getFromTokens({
        fromChainId: incomingFromChainId ?? null,
        fromTokenRouteSymbol: state.fromToken?.routeSymbol ?? null,
        toChainId: state.toChainId ?? null,
        toTokenRouteSymbol: null,
      })
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)

      const validToChainIds = getToChainIds({
        fromChainId: incomingFromChainId ?? null,
        fromTokenRouteSymbol: null,
        toChainId: state.toChainId ?? null,
        toTokenRouteSymbol: null,
      })

      const validToTokens = getToTokens({
        fromChainId: incomingFromChainId ?? null,
        fromTokenRouteSymbol: state.fromToken?.routeSymbol ?? null,
        toChainId: null,
        toTokenRouteSymbol: null,
      })
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)

      let validFromToken
      let validToChainId
      let validToToken

      if (
        validFromTokens?.some(
          (token) => token?.routeSymbol === state.fromToken?.routeSymbol
        )
      ) {
        validFromToken = state.fromToken
      } else {
        validFromToken = findValidToken(
          validFromTokens,
          state.toToken?.routeSymbol,
          state.toToken?.swapableType
        )
      }

      if (
        validToChainIds?.includes(state.toChainId) &&
        incomingFromChainId !== state.toChainId
      ) {
        validToChainId = state.toChainId
      } else {
        validToChainId = null
      }

      if (
        validToTokens?.some(
          (token) => token?.routeSymbol === state.toToken?.routeSymbol
        )
      ) {
        validToToken = state.toToken
      } else {
        validToToken = findValidToken(
          validToTokens,
          state.fromToken?.routeSymbol,
          state.fromToken?.swapableType
        )
      }

      const {
        fromChainId,
        fromToken,
        toChainId,
        toToken,
        fromChainIds,
        fromTokens,
        toChainIds,
        toTokens,
      } = getRoutePossibilities({
        fromChainId: incomingFromChainId,
        fromToken: validFromToken,
        toChainId: validToChainId,
        toToken: validToToken,
      })

      state.fromChainId = fromChainId
      state.fromToken = fromToken
      state.toChainId = toChainId
      state.toToken = toToken
      state.fromChainIds = fromChainIds
      state.fromTokens = fromTokens
      state.toChainIds = toChainIds
      state.toTokens = toTokens
    },
    setFromToken: (state, action: PayloadAction<Token>) => {
      const incomingFromToken = action.payload

      const validFromChainIds = getFromChainIds({
        fromChainId: state.fromChainId ?? null,
        fromTokenRouteSymbol: incomingFromToken?.routeSymbol ?? null,
        toChainId: null,
        toTokenRouteSymbol: null,
      })

      const validToChainIds = getToChainIds({
        fromChainId: state.fromChainId ?? null,
        fromTokenRouteSymbol: incomingFromToken?.routeSymbol ?? null,
        toChainId: null,
        toTokenRouteSymbol: null,
      })

      const validToTokens = getToTokens({
        fromChainId: state.fromChainId ?? null,
        fromTokenRouteSymbol: incomingFromToken?.routeSymbol ?? null,
        toChainId: state.toChainId ?? null,
        toTokenRouteSymbol: null,
      })
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)

      let validFromChainId
      let validToChainId
      let validToToken

      if (validFromChainIds?.includes(state.fromChainId)) {
        validFromChainId = state.fromChainId
      } else {
        validFromChainId = null
      }

      if (validToChainIds?.includes(state.toChainId)) {
        validToChainId = state.toChainId
      } else {
        validToChainId = null
      }

      if (
        validToTokens?.some(
          (token) => token?.routeSymbol === state.toToken?.routeSymbol
        )
      ) {
        validToToken = state.toToken
      } else {
        validToToken = findValidToken(
          validToTokens,
          incomingFromToken?.routeSymbol,
          incomingFromToken?.swapableType
        )
      }

      const {
        fromChainId,
        fromToken,
        toChainId,
        toToken,
        fromChainIds,
        fromTokens,
        toChainIds,
        toTokens,
      } = getRoutePossibilities({
        fromChainId: validFromChainId,
        fromToken: incomingFromToken,
        toChainId: validToChainId,
        toToken: validToToken,
      })

      state.fromChainId = fromChainId
      state.fromToken = fromToken
      state.toToken = state.toChainId ? toToken : null
      state.toChainId = toChainId
      state.fromChainIds = fromChainIds
      state.fromTokens = fromTokens
      state.toChainIds = toChainIds
      state.toTokens = toTokens
    },
    setToChainId: (state, action: PayloadAction<number>) => {
      const incomingToChainId = action.payload

      const validFromChainIds = getFromChainIds({
        fromChainId: state.fromChainId ?? null,
        fromTokenRouteSymbol: null,
        toChainId: incomingToChainId ?? null,
        toTokenRouteSymbol: null,
      })

      const validFromTokens = getFromTokens({
        fromChainId: state.fromChainId ?? null,
        fromTokenRouteSymbol: state.fromToken?.routeSymbol ?? null,
        toChainId: incomingToChainId ?? null,
        toTokenRouteSymbol: state.toToken?.routeSymbol ?? null,
      })
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)

      const validToTokens = getToTokens({
        fromChainId: state.fromChainId ?? null,
        fromTokenRouteSymbol: state.fromToken?.routeSymbol ?? null,
        toChainId: incomingToChainId ?? null,
        toTokenRouteSymbol: state.toToken?.routeSymbol ?? null,
      })
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)

      let validFromChainId
      let validFromToken
      let validToToken

      if (
        validFromChainIds?.includes(state.fromChainId) &&
        incomingToChainId !== state.fromChainId
      ) {
        validFromChainId = state.fromChainId
      } else {
        validFromChainId = null
      }

      if (
        validFromTokens?.some(
          (token) => token?.routeSymbol === state.fromToken?.routeSymbol
        )
      ) {
        validFromToken = state.fromToken
      } else {
        validFromToken = findValidToken(
          validFromTokens,
          state.fromToken?.routeSymbol,
          state.fromToken?.swapableType
        )
      }

      if (
        validToTokens?.some(
          (token) => token?.routeSymbol === state.toToken?.routeSymbol
        )
      ) {
        validToToken = state.toToken
      } else {
        validToToken = findValidToken(
          validToTokens,
          state.fromToken?.routeSymbol,
          state.fromToken?.swapableType
        )
      }

      const {
        fromChainId,
        fromToken,
        toChainId,
        toToken,
        fromChainIds,
        fromTokens,
        toChainIds,
        toTokens,
      } = getRoutePossibilities({
        fromChainId: validFromChainId,
        fromToken: validFromToken,
        toChainId: incomingToChainId,
        toToken: validToToken,
      })

      state.fromChainId = fromChainId
      state.fromToken = fromToken
      state.toChainId = toChainId
      state.toToken = toToken
      state.fromChainIds = fromChainIds
      state.fromTokens = fromTokens
      state.toChainIds = toChainIds
      state.toTokens = toTokens
    },
    setToToken: (state, action: PayloadAction<Token>) => {
      const incomingToToken = action.payload

      const validFromChainIds = getFromChainIds({
        fromChainId: state.fromChainId ?? null,
        fromTokenRouteSymbol: null,
        toChainId: state.toChainId ?? null,
        toTokenRouteSymbol: incomingToToken?.routeSymbol ?? null,
      })

      const validFromTokens = getFromTokens({
        fromChainId: state.fromChainId ?? null,
        fromTokenRouteSymbol: state.fromToken?.routeSymbol ?? null,
        toChainId: state.toChainId ?? null,
        toTokenRouteSymbol: incomingToToken?.routeSymbol ?? null,
      })
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)

      const validToChainIds = getToChainIds({
        fromChainId: null,
        fromTokenRouteSymbol: null,
        toChainId: state.toChainId ?? null,
        toTokenRouteSymbol: incomingToToken?.routeSymbol ?? null,
      })

      let validFromChainId
      let validFromToken
      let validToChainId

      if (validFromChainIds?.includes(state.fromChainId)) {
        validFromChainId = state.fromChainId
      } else {
        validFromChainId = null
      }

      if (
        validFromTokens?.some(
          (token) => token?.routeSymbol === state.fromToken?.routeSymbol
        )
      ) {
        validFromToken = state.fromToken
      } else {
        validFromToken = findValidToken(
          validFromTokens,
          incomingToToken?.routeSymbol,
          incomingToToken?.swapableType
        )
      }

      if (validToChainIds?.includes(state.toChainId)) {
        validToChainId = state.toChainId
      } else {
        validToChainId = null
      }

      const {
        fromChainId,
        fromToken,
        toChainId,
        toToken,
        fromChainIds,
        fromTokens,
        toChainIds,
        toTokens,
      } = getRoutePossibilities({
        fromChainId: validFromChainId,
        fromToken: validFromToken,
        toChainId: validToChainId,
        toToken: incomingToToken,
      })

      state.fromChainId = fromChainId
      state.fromToken = fromToken
      state.toChainId = toChainId
      state.toToken = toToken
      state.fromChainIds = fromChainIds
      state.fromTokens = fromTokens
      state.toChainIds = toChainIds
      state.toTokens = toTokens
    },
    setBridgeQuote: (state, action: PayloadAction<BridgeQuote>) => {
      state.bridgeQuote = action.payload
    },
    resetBridgeQuote: (state) => {
      state.bridgeQuote = initialState.bridgeQuote
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
  },
  extraReducers: (builder) => {
    builder
      .addCase(
        updateDebouncedFromValue,
        (state, action: PayloadAction<string>) => {
          state.debouncedFromValue = action.payload
        }
      )
      .addCase(
        updateDebouncedToTokensFromValue,
        (state, action: PayloadAction<string>) => {
          state.debouncedToTokensFromValue = action.payload
        }
      )
      .addCase(resetBridgeInputs, (state) => {
        state.fromChainId = initialState.fromChainId
        state.fromToken = initialState.fromToken
        state.toChainId = initialState.toChainId
        state.toToken = initialState.toToken
        state.fromValue = initialState.fromValue
        state.debouncedFromValue = initialState.debouncedFromValue
      })
      .addCase(fetchAndStoreBridgeQuotes.pending, (state) => {
        state.toTokensBridgeQuotesStatus = FetchState.LOADING
      })
      .addCase(
        fetchAndStoreBridgeQuotes.fulfilled,
        (state, action: PayloadAction<BridgeQuoteResponse[]>) => {
          state.toTokensBridgeQuotes = action.payload
          state.toTokensBridgeQuotesStatus = FetchState.VALID
        }
      )
      .addCase(fetchAndStoreBridgeQuotes.rejected, (state) => {
        state.toTokensBridgeQuotesStatus = FetchState.INVALID
      })
      .addCase(resetFetchedBridgeQuotes, (state) => {
        state.toTokensBridgeQuotes = initialState.toTokensBridgeQuotes
        state.toTokensBridgeQuotesStatus =
          initialState.toTokensBridgeQuotesStatus
      })
  },
})

export const {
  setBridgeQuote,
  resetBridgeQuote,
  setFromChainId,
  setToChainId,
  setFromToken,
  setToToken,
  updateFromValue,
  setDeadlineMinutes,
  setDestinationAddress,
  setIsLoading,
} = bridgeSlice.actions

export default bridgeSlice.reducer
