import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { BridgeableToken } from 'types'
import { getFromTokens } from '@/utils/routeMaker/getFromTokens'
import { findTokenByRouteSymbol } from '@/utils/findTokenByRouteSymbol'

import {
  getRoutePossibilities,
  getSymbol,
} from '@/utils/routeMaker/generateRoutePossibilities'

import { getToChainIds } from '@/utils/routeMaker/getToChainIds'
import { getToTokens } from '@/utils/routeMaker/getToTokens'
import { findValidToken } from '@/utils/findValidTokens'
import { getFromChainIds } from '@/utils/routeMaker/getFromChainIds'

export interface BridgeState {
  debouncedInputAmount: string
  originChainId: number
  originToken: BridgeableToken
  destinationChainId: number
  destinationToken: BridgeableToken
  originChainIds: number[]
  originTokens: BridgeableToken[]
  destinationChainIds: number[]
  destinationTokens: BridgeableToken[]
  tokens: BridgeableToken[]
}

const initialState: BridgeState = {
  debouncedInputAmount: '',
  originChainId: 42161,
  originToken: null,
  destinationChainId: null,
  destinationToken: null,
  originChainIds: [],
  originTokens: [],
  destinationChainIds: [],
  destinationTokens: [],
  tokens: [],
}

export const bridgeSlice = createSlice({
  name: 'bridge',
  initialState,
  reducers: {
    setDebouncedInputAmount: (
      state: BridgeState,
      action: PayloadAction<string>
    ) => {
      state.debouncedInputAmount = action.payload
    },
    setOriginChainId: (state: BridgeState, action: PayloadAction<number>) => {
      const incomingFromChainId = action.payload

      const validFromTokens = getFromTokens({
        fromChainId: incomingFromChainId ?? null,
        fromTokenRouteSymbol: state.originToken?.routeSymbol ?? null,
        toChainId: state.destinationChainId ?? null,
        toTokenRouteSymbol: null,
      })
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)

      const validToChainIds = getToChainIds({
        fromChainId: incomingFromChainId ?? null,
        fromTokenRouteSymbol: null,
        toChainId: state.destinationChainId ?? null,
        toTokenRouteSymbol: null,
      })

      const validToTokens = getToTokens({
        fromChainId: incomingFromChainId ?? null,
        fromTokenRouteSymbol: state.originToken?.routeSymbol ?? null,
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
          (token) => token?.routeSymbol === state.originToken?.routeSymbol
        )
      ) {
        validFromToken = state.originToken
      } else {
        validFromToken = findValidToken(
          validFromTokens,
          state.destinationToken?.routeSymbol,
          state.destinationToken?.swapableType
        )
      }

      if (
        validToChainIds?.includes(state.destinationChainId) &&
        incomingFromChainId !== state.destinationChainId
      ) {
        validToChainId = state.destinationChainId
      } else {
        validToChainId = null
      }

      if (
        validToTokens?.some(
          (token) => token?.routeSymbol === state.destinationToken?.routeSymbol
        )
      ) {
        validToToken = state.destinationToken
      } else {
        validToToken = findValidToken(
          validToTokens,
          state.originToken?.routeSymbol,
          state.originToken?.swapableType
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

      state.originChainId = fromChainId
      state.originToken = fromToken
      state.destinationChainId = toChainId
      state.destinationToken = toToken
      state.originChainIds = fromChainIds
      state.originTokens = fromTokens
      state.destinationChainIds = toChainIds
      state.destinationTokens = toTokens
    },
    setOriginToken: (
      state: BridgeState,
      action: PayloadAction<BridgeableToken>
    ) => {
      const incomingFromToken = action.payload

      const validFromChainIds = getFromChainIds({
        fromChainId: state.originChainId ?? null,
        fromTokenRouteSymbol: incomingFromToken?.routeSymbol ?? null,
        toChainId: null,
        toTokenRouteSymbol: null,
      })

      const validToChainIds = getToChainIds({
        fromChainId: state.originChainId ?? null,
        fromTokenRouteSymbol: incomingFromToken?.routeSymbol ?? null,
        toChainId: null,
        toTokenRouteSymbol: null,
      })

      const validToTokens = getToTokens({
        fromChainId: state.originChainId ?? null,
        fromTokenRouteSymbol: incomingFromToken?.routeSymbol ?? null,
        toChainId: state.destinationChainId ?? null,
        toTokenRouteSymbol: null,
      })
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)

      let validFromChainId
      let validToChainId
      let validToToken

      if (validFromChainIds?.includes(state.originChainId)) {
        validFromChainId = state.originChainId
      } else {
        validFromChainId = null
      }

      if (validToChainIds?.includes(state.destinationChainId)) {
        validToChainId = state.destinationChainId
      } else {
        validToChainId = null
      }

      if (
        validToTokens?.some(
          (token) => token?.routeSymbol === state.destinationToken?.routeSymbol
        )
      ) {
        validToToken = state.destinationToken
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

      state.originChainId = fromChainId
      state.originToken = fromToken
      state.destinationToken = state.destinationChainId ? toToken : null
      state.destinationChainId = toChainId
      state.originChainIds = fromChainIds
      state.originTokens = fromTokens
      state.destinationChainIds = toChainIds
      state.destinationTokens = toTokens
    },
    setDestinationChainId: (
      state: BridgeState,
      action: PayloadAction<number>
    ) => {
      const incomingToChainId = action.payload

      const validFromChainIds = getFromChainIds({
        fromChainId: state.originChainId ?? null,
        fromTokenRouteSymbol: null,
        toChainId: incomingToChainId ?? null,
        toTokenRouteSymbol: null,
      })

      const validFromTokens = getFromTokens({
        fromChainId: state.originChainId ?? null,
        fromTokenRouteSymbol: state.originToken?.routeSymbol ?? null,
        toChainId: incomingToChainId ?? null,
        toTokenRouteSymbol: state.destinationToken?.routeSymbol ?? null,
      })
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)

      const validToTokens = getToTokens({
        fromChainId: state.originChainId ?? null,
        fromTokenRouteSymbol: state.originToken?.routeSymbol ?? null,
        toChainId: incomingToChainId ?? null,
        toTokenRouteSymbol: state.destinationToken?.routeSymbol ?? null,
      })
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)

      let validFromChainId
      let validFromToken
      let validToToken

      if (
        validFromChainIds?.includes(state.originChainId) &&
        incomingToChainId !== state.originChainId
      ) {
        validFromChainId = state.originChainId
      } else {
        validFromChainId = null
      }

      if (
        validFromTokens?.some(
          (token) => token?.routeSymbol === state.originToken?.routeSymbol
        )
      ) {
        validFromToken = state.originToken
      } else {
        validFromToken = findValidToken(
          validFromTokens,
          state.originToken?.routeSymbol,
          state.originToken?.swapableType
        )
      }

      if (
        validToTokens?.some(
          (token) => token?.routeSymbol === state.destinationToken?.routeSymbol
        )
      ) {
        validToToken = state.destinationToken
      } else {
        validToToken = findValidToken(
          validToTokens,
          state.originToken?.routeSymbol,
          state.originToken?.swapableType
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

      state.originChainId = fromChainId
      state.originToken = fromToken
      state.destinationChainId = toChainId
      state.destinationToken = toToken
      state.originChainIds = fromChainIds
      state.originTokens = fromTokens
      state.destinationChainIds = toChainIds
      state.destinationTokens = toTokens
    },
    setDestinationToken: (
      state: BridgeState,
      action: PayloadAction<BridgeableToken>
    ) => {
      const incomingToToken = action.payload

      const validFromChainIds = getFromChainIds({
        fromChainId: state.originChainId ?? null,
        fromTokenRouteSymbol: null,
        toChainId: state.destinationChainId ?? null,
        toTokenRouteSymbol: incomingToToken?.routeSymbol ?? null,
      })

      const validFromTokens = getFromTokens({
        fromChainId: state.originChainId ?? null,
        fromTokenRouteSymbol: state.originToken?.routeSymbol ?? null,
        toChainId: state.destinationChainId ?? null,
        toTokenRouteSymbol: incomingToToken?.routeSymbol ?? null,
      })
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)

      const validToChainIds = getToChainIds({
        fromChainId: null,
        fromTokenRouteSymbol: null,
        toChainId: state.destinationChainId ?? null,
        toTokenRouteSymbol: incomingToToken?.routeSymbol ?? null,
      })

      let validFromChainId
      let validFromToken
      let validToChainId

      if (validFromChainIds?.includes(state.originChainId)) {
        validFromChainId = state.originChainId
      } else {
        validFromChainId = null
      }

      if (
        validFromTokens?.some(
          (token) => token?.routeSymbol === state.originToken?.routeSymbol
        )
      ) {
        validFromToken = state.originToken
      } else {
        validFromToken = findValidToken(
          validFromTokens,
          incomingToToken?.routeSymbol,
          incomingToToken?.swapableType
        )
      }

      if (validToChainIds?.includes(state.destinationChainId)) {
        validToChainId = state.destinationChainId
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

      state.originChainId = fromChainId
      state.originToken = fromToken
      state.destinationChainId = toChainId
      state.destinationToken = toToken
      state.originChainIds = fromChainIds
      state.originTokens = fromTokens
      state.destinationChainIds = toChainIds
      state.destinationTokens = toTokens
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
  setDebouncedInputAmount,
  setOriginChainId,
  setDestinationChainId,
  setOriginToken,
  setDestinationToken,
  setTokens,
} = bridgeSlice.actions

export default bridgeSlice.reducer
