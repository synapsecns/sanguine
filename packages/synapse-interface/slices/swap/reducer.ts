import _ from 'lodash'
import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { DAI, USDC } from '@/constants/tokens/bridgeable'
import { EMPTY_SWAP_QUOTE } from '@/constants/swap'
import { ETH as ETHEREUM } from '@/constants/chains/master'
import { getSwapPossibilities } from '@/utils/swapFinder/generateSwapPossibilities'
import { SwapQuote, Token } from '@/utils/types'
import { getSwapFromTokens } from '@/utils/swapFinder/getSwapFromTokens'
import { getSymbol } from '@/utils/getSymbol'
import { findTokenByRouteSymbol } from '@/utils/findTokenByRouteSymbol'
import { getSwapToTokens } from '@/utils/swapFinder/getSwapToTokens'
import { getSwapFromChainIds } from '@/utils/swapFinder/getSwapFromChainIds'
import { findValidToken } from '@/utils/findValidToken'
import { flattenPausedTokens } from '@/utils/flattenPausedTokens'

export interface SwapState {
  swapChainId: number
  swapFromToken: Token
  swapToToken: Token
  swapFromChainIds: number[]
  swapFromTokens: Token[]
  swapToTokens: Token[]

  swapFromValue: string
  swapQuote: SwapQuote
  isLoading: boolean
}

const { fromChainId, fromToken, toToken, fromChainIds, fromTokens, toTokens } =
  getSwapPossibilities({
    fromChainId: ETHEREUM.id,
    fromToken: USDC,
    toChainId: ETHEREUM.id,
    toToken: DAI,
  })

export const initialState: SwapState = {
  swapChainId: fromChainId,
  swapFromToken: fromToken,
  swapToToken: toToken,
  swapFromChainIds: fromChainIds,
  swapFromTokens: fromTokens,
  swapToTokens: toTokens,

  swapFromValue: '',
  swapQuote: EMPTY_SWAP_QUOTE,
  isLoading: false,
}

export const swapSlice = createSlice({
  name: 'swap',
  initialState,
  reducers: {
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload
    },
    setSwapChainId: (state, action: PayloadAction<number>) => {
      const incomingFromChainId = action.payload

      const validFromTokens = _(
        getSwapFromTokens({
          fromChainId: incomingFromChainId ?? null,
          fromTokenRouteSymbol: state.swapFromToken?.routeSymbol ?? null,
          toChainId: incomingFromChainId ?? null,
          toTokenRouteSymbol: null,
        })
      )
        .difference(flattenPausedTokens())
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)
        .value()

      const validToTokens = _(
        getSwapToTokens({
          fromChainId: incomingFromChainId ?? null,
          fromTokenRouteSymbol: state.swapFromToken?.routeSymbol ?? null,
          toChainId: incomingFromChainId ?? null,
          toTokenRouteSymbol: null,
        })
      )
        .difference(flattenPausedTokens())
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)
        .value()

      let validFromToken
      let validToToken

      if (
        validFromTokens?.some(
          (token) => token?.routeSymbol === state.swapFromToken?.routeSymbol
        )
      ) {
        validFromToken = state.swapFromToken
      } else {
        validFromToken = findValidToken(
          validFromTokens,
          state.swapToToken?.routeSymbol,
          state.swapToToken?.swapableType
        )
      }

      if (
        validToTokens?.some(
          (token) => token?.routeSymbol === state.swapToToken?.routeSymbol
        )
      ) {
        validToToken = state.swapToToken
      } else {
        validToToken = findValidToken(
          validToTokens,
          state.swapFromToken?.routeSymbol,
          state.swapFromToken?.swapableType
        )
      }

      const {
        fromChainId,
        fromToken,
        toToken,
        fromChainIds,
        fromTokens,
        toTokens,
      } = getSwapPossibilities({
        fromChainId: incomingFromChainId,
        fromToken: validFromToken,
        toChainId: incomingFromChainId,
        toToken: validToToken,
      })

      state.swapChainId = fromChainId
      state.swapFromToken = fromToken
      state.swapToToken = toToken
      state.swapFromChainIds = fromChainIds
      state.swapFromTokens = fromTokens
      state.swapToTokens = toTokens
    },
    setSwapFromToken: (state, action: PayloadAction<Token>) => {
      const incomingFromToken = action.payload

      const validFromChainIds = getSwapFromChainIds({
        fromChainId: state.swapChainId ?? null,
        fromTokenRouteSymbol: incomingFromToken?.routeSymbol ?? null,
        toChainId: null,
        toTokenRouteSymbol: null,
      })

      const validToTokens = _(
        getSwapToTokens({
          fromChainId: state.swapChainId ?? null,
          fromTokenRouteSymbol: incomingFromToken?.routeSymbol ?? null,
          toChainId: state.swapChainId ?? null,
          toTokenRouteSymbol: null,
        })
      )
        .difference(flattenPausedTokens())
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)
        .value()

      let validFromChainId
      let validToToken

      if (validFromChainIds?.includes(state.swapChainId)) {
        validFromChainId = state.swapChainId
      } else {
        validFromChainId = null
      }

      if (
        validToTokens?.some(
          (token) => token?.routeSymbol === state.swapToToken?.routeSymbol
        )
      ) {
        validToToken = state.swapToToken
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
        toToken,
        fromChainIds,
        fromTokens,
        toTokens,
      } = getSwapPossibilities({
        fromChainId: validFromChainId,
        fromToken: incomingFromToken,
        toChainId: validFromChainId,
        toToken: validToToken,
      })

      state.swapChainId = fromChainId
      state.swapFromToken = fromToken
      state.swapToToken = toToken
      state.swapFromChainIds = fromChainIds
      state.swapFromTokens = fromTokens
      state.swapToTokens = toTokens
    },
    setSwapToToken: (state, action: PayloadAction<Token>) => {
      const incomingToToken = action.payload

      const validFromChainIds = getSwapFromChainIds({
        fromChainId: state.swapChainId ?? null,
        fromTokenRouteSymbol: null,
        toChainId: state.swapChainId ?? null,
        toTokenRouteSymbol: incomingToToken?.routeSymbol ?? null,
      })

      const validFromTokens = _(
        getSwapFromTokens({
          fromChainId: state.swapChainId ?? null,
          fromTokenRouteSymbol: state.swapFromToken?.routeSymbol ?? null,
          toChainId: state.swapChainId ?? null,
          toTokenRouteSymbol: incomingToToken?.routeSymbol ?? null,
        })
      )
        .difference(flattenPausedTokens())
        ?.map(getSymbol)
        .map((s) => findTokenByRouteSymbol(s))
        .filter(Boolean)
        .value()

      let validFromChainId
      let validFromToken

      if (validFromChainIds?.includes(state.swapChainId)) {
        validFromChainId = state.swapChainId
      } else {
        validFromChainId = null
      }

      if (
        validFromTokens?.some(
          (token) => token?.routeSymbol === state.swapFromToken?.routeSymbol
        )
      ) {
        validFromToken = state.swapFromToken
      } else {
        validFromToken = findValidToken(
          validFromTokens,
          incomingToToken?.routeSymbol,
          incomingToToken?.swapableType
        )
      }

      const {
        fromChainId,
        fromToken,
        toToken,
        fromChainIds,
        fromTokens,
        toTokens,
      } = getSwapPossibilities({
        fromChainId: validFromChainId,
        fromToken: validFromToken,
        toChainId: validFromChainId,
        toToken: incomingToToken,
      })

      state.swapChainId = fromChainId
      state.swapFromToken = fromToken
      state.swapToToken = toToken
      state.swapFromChainIds = fromChainIds
      state.swapFromTokens = fromTokens
      state.swapToTokens = toTokens
    },
    setSwapQuote: (state, action: PayloadAction<SwapQuote>) => {
      state.swapQuote = action.payload
    },
    updateSwapFromValue: (state, action: PayloadAction<string>) => {
      state.swapFromValue = action.payload
    },
  },
})

export const {
  setSwapChainId,
  setSwapFromToken,
  setSwapToToken,
  updateSwapFromValue,
  setSwapQuote,
  setIsLoading,
} = swapSlice.actions

export default swapSlice.reducer
