import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Address } from 'wagmi'

import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { ETH } from '@/constants/tokens/master'
import { ARBITRUM, ETH as ETHEREUM } from '@/constants/chains/master'
import { BridgeQuote, Token } from '@/utils/types'
import { TokenWithBalanceAndAllowances } from '@/utils/actions/fetchPortfolioBalances'
import {
  PendingBridgeTransaction,
  addPendingBridgeTransaction,
  removePendingBridgeTransaction,
  updatePendingBridgeTransaction,
  updatePendingBridgeTransactions,
} from './actions'

export interface BridgeState {
  fromChainId: number
  supportedFromTokens: TokenWithBalanceAndAllowances[]
  toChainId: number
  supportedToTokens: Token[]
  fromToken: Token
  toToken: Token
  fromValue: string
  bridgeQuote: BridgeQuote
  fromChainIds: number[]
  toChainIds: number[]
  isLoading: boolean
  deadlineMinutes: number | null
  destinationAddress: Address | null
  bridgeTxHashes: string[] | null
  pendingBridgeTransactions: PendingBridgeTransaction[]
}

// How do we update query params based on initial state?
// Additionally how do we set query params based on user input updates?
export const initialState: BridgeState = {
  fromChainId: ETHEREUM.id,
  supportedFromTokens: [],
  toChainId: ARBITRUM.id,
  supportedToTokens: [],
  fromToken: ETH,
  toToken: ETH,
  fromValue: '',
  bridgeQuote: EMPTY_BRIDGE_QUOTE,
  fromChainIds: [],
  toChainIds: [],
  isLoading: false,
  deadlineMinutes: null,
  destinationAddress: null,
  bridgeTxHashes: [],
  pendingBridgeTransactions: [],
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
    setSupportedFromTokens: (
      state,
      action: PayloadAction<TokenWithBalanceAndAllowances[]>
    ) => {
      state.supportedFromTokens = action.payload
    },
    setSupportedToTokens: (state, action: PayloadAction<Token[]>) => {
      state.supportedToTokens = action.payload
    },
    setFromChainIds: (state, action: PayloadAction<number[]>) => {
      state.fromChainIds = action.payload
    },
    setToChainIds: (state, action: PayloadAction<number[]>) => {
      state.toChainIds = action.payload
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
  extraReducers: (builder) => {
    builder
      .addCase(
        addPendingBridgeTransaction,
        (state, action: PayloadAction<PendingBridgeTransaction>) => {
          state.pendingBridgeTransactions = [
            ...state.pendingBridgeTransactions,
            action.payload,
          ]
        }
      )
      .addCase(
        updatePendingBridgeTransaction,
        (
          state,
          action: PayloadAction<{
            timestamp: number
            transactionHash: string
            isSubmitted: boolean
          }>
        ) => {
          const { timestamp, transactionHash, isSubmitted } = action.payload
          const transactionIndex = state.pendingBridgeTransactions.findIndex(
            (transaction) => transaction.timestamp === timestamp
          )

          if (transactionIndex !== -1) {
            state.pendingBridgeTransactions =
              state.pendingBridgeTransactions.map((transaction, index) =>
                index === transactionIndex
                  ? { ...transaction, transactionHash, isSubmitted }
                  : transaction
              )
          }
        }
      )
      .addCase(
        removePendingBridgeTransaction,
        (state, action: PayloadAction<number>) => {
          const timestampToRemove = action.payload
          state.pendingBridgeTransactions =
            state.pendingBridgeTransactions.filter(
              (transaction) => transaction.timestamp !== timestampToRemove
            )
        }
      )
      .addCase(
        updatePendingBridgeTransactions,
        (state, action: PayloadAction<PendingBridgeTransaction[]>) => {
          state.pendingBridgeTransactions = action.payload
        }
      )
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
  setFromChainIds,
  setToChainIds,
  setDeadlineMinutes,
  setDestinationAddress,
  setIsLoading,
  addBridgeTxHash,
} = bridgeSlice.actions

export default bridgeSlice.reducer
