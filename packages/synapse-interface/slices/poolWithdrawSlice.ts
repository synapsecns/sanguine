import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Token } from '@/utils/types'
import { formatBNToString } from '@/utils/bigint/format'
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
  routerAddress: string
}

type InputValue = {
  bi: bigint
  str: string
}

export const DEFAULT_WITHDRAW_QUOTE: WithdrawQuote = {
  priceImpact: 0n,
  outputs: {},
  allowance: undefined,
  routerAddress: '',
}

const DEFAULT_INPUT_VALUE = { bi: 0n, str: '' }

export interface PoolDepositState {
  withdrawQuote: WithdrawQuote
  isLoading: boolean
  inputValue: InputValue
  pool: Token
  withdrawType: string
}

const initialState: PoolDepositState = {
  withdrawQuote: DEFAULT_WITHDRAW_QUOTE,
  isLoading: false,
  inputValue: DEFAULT_INPUT_VALUE,
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
    setInputValue: (state, action: PayloadAction<InputValue>) => {
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
