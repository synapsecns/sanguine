import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { ETH as ETHEREUM, OPTIMISM } from '@/constants/chains/master'
import { generateRoutePossibilities } from '@/utils/generateRoutePossibilities'

export interface TokenSelectorState {
  fromChainId: number
  fromToken: string
  toToken: string
  toChainId: number
  fromChainIds: number[]
  toChainIds: number[]
  fromTokens: string[]
  toTokens: string[]
  isLoading: boolean
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
} = generateRoutePossibilities({
  fromChainId: ETHEREUM.id,
  fromToken: 'USDC-1',
  toChainId: OPTIMISM.id,
  toToken: 'USDC-10',
})

const initialState: TokenSelectorState = {
  fromChainId,
  fromToken,
  toChainId,
  toToken,
  fromChainIds,
  toChainIds,
  fromTokens,
  toTokens,
  isLoading: false,
}

export const tokenSelectorSlice = createSlice({
  name: 'tokenSelector',
  initialState,
  reducers: {
    resetTokenSelectorState: () => initialState,
    clearTokenSelectorState: (state) => {
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
        fromChainId: null,
        fromToken: null,
        toChainId: null,
        toToken: null,
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
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload
    },
    setSelectFromChainId: (state, action: PayloadAction<number>) => {
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
        fromToken: null,
        toChainId: null,
        toToken: null,
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
    setSelectFromToken: (state, action: PayloadAction<string>) => {
      const incomingFromToken = action.payload
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
        fromToken: incomingFromToken,
        toChainId: state.toChainId,
        toToken: null,
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
    setSelectToChainId: (state, action: PayloadAction<number>) => {
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
        fromToken: state.fromToken,
        toChainId: incomingToChainId,
        toToken: null,
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
    setSelectToToken: (state, action: PayloadAction<string>) => {
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
        fromToken: state.fromToken,
        toChainId: state.toChainId,
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
  },
})

export const {
  setSelectFromChainId,
  setSelectToChainId,
  setSelectFromToken,
  setSelectToToken,
  setIsLoading,
  resetTokenSelectorState,
  clearTokenSelectorState,
} = tokenSelectorSlice.actions

export default tokenSelectorSlice.reducer
