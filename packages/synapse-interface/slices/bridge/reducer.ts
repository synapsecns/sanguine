import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Address } from 'wagmi'

import { NUSD, USDT } from '@/constants/tokens/master'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { ARBITRUM, ETH as ETHEREUM } from '@/constants/chains/master'
import { BridgeQuote, Token } from '@/utils/types'
import { getRoutePossibilities } from '@/utils/routeMaker/generateRoutePossibilities'

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
  fromChainId: ETHEREUM.id,
  fromToken: NUSD,
  toChainId: ARBITRUM.id,
  toToken: USDT,
})

// How do we update query params based on initial state?
// Additionally how do we set query params based on user input updates?
const initialState: BridgeState = {
  fromChainId,
  fromToken,
  toChainId,
  toToken,
  fromChainIds,
  toChainIds,
  fromTokens,
  toTokens,

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


If the state.toToken exists and it's not in the possible lists of toToken set it to null and the user can select

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
      } = getRoutePossibilities({
        fromChainId: incomingFromChainId,
        fromToken: state.fromToken,
        toChainId: state.toChainId,
        toToken: state.toToken,
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
        fromChainId: state.fromChainId,
        fromToken: incomingFromToken,
        toChainId: state.toChainId,
        toToken: state.toToken,
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
      } = getRoutePossibilities({
        fromChainId: state.fromChainId,
        fromToken: state.fromToken,
        toChainId: incomingToChainId,
        toToken: state.toToken,
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
