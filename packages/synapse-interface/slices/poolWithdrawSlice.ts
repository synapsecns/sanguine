import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { Token } from '@/utils/types'
import { ALL } from '@/constants/withdrawTypes'

type WithdrawQuote = {
  priceImpact: bigint
  outputs: Record<
    string,
    {
      value: bigint
      index: number
    }
  >
  allowance: bigint
}

export const DEFAULT_WITHDRAW_QUOTE: WithdrawQuote = {
  priceImpact: 0n,
  outputs: {},
  allowance: undefined,
}

interface PoolWithdrawState {
  withdrawQuote: WithdrawQuote
  isLoading: boolean
  inputValue: string
  pool: Token
  withdrawType: string
}

const initialState: PoolWithdrawState = {
  withdrawQuote: DEFAULT_WITHDRAW_QUOTE,
  isLoading: false,
  inputValue: '',
  pool: null,
  withdrawType: ALL,
}

export const poolWithdrawSlice = createSlice({
  name: 'poolWithdraw',
  initialState,
  reducers: {
    resetPoolWithdraw: () => initialState,
    setWithdrawQuote: (state, action: PayloadAction<WithdrawQuote>) => {
      state.withdrawQuote = action.payload
    },
    setInputValue: (state, action: PayloadAction<string>) => {
      state.inputValue = action.payload
    },
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload
    },
    setWithdrawType: (state, action: PayloadAction<string>) => {
      state.withdrawType = action.payload
    },
  },
})

export const {
  resetPoolWithdraw,
  setInputValue,
  setIsLoading,
  setWithdrawQuote,
  setWithdrawType,
} = poolWithdrawSlice.actions

export default poolWithdrawSlice.reducer
