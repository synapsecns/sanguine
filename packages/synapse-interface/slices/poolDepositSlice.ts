import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { Token } from '@/utils/types'
import { formatBigIntToString } from '@/utils/bigint/format'

type DepositQuote = {
  priceImpact: any
  allowances: {}
  routerAddress: string
}

type InputValue = {
  bi: Record<string, bigint>
  str: Record<string, string>
}

const DEFAULT_DEPOSIT_QUOTE = {
  priceImpact: 0n,
  allowances: {},
  routerAddress: '',
}

const DEFAULT_INPUT_VALUE = { bi: {}, str: {} }

interface PoolDepositState {
  depositQuote: DepositQuote
  isLoading: boolean
  inputValue: InputValue
  filteredInputValue: InputValue
  inputSum: any
  pool: Token
}

const initialState: PoolDepositState = {
  depositQuote: DEFAULT_DEPOSIT_QUOTE,
  isLoading: false,
  inputValue: DEFAULT_INPUT_VALUE,
  filteredInputValue: DEFAULT_INPUT_VALUE,
  inputSum: 0,
  pool: null,
}

export const poolDepositSlice = createSlice({
  name: 'poolDeposit',
  initialState,
  reducers: {
    resetPoolDeposit: () => initialState,
    setDepositQuote: (state, action: PayloadAction<DepositQuote>) => {
      state.depositQuote = action.payload
    },
    setInputValue: (state, action: PayloadAction<InputValue>) => {
      state.inputValue = action.payload

      const filteredVal = filterAndSerializeInputValues(
        action.payload,
        state.pool
      )
      state.filteredInputValue = filteredVal
      state.inputSum = sumBigIntegers(state.pool, filteredVal)
    },
    setPool: (state, action: PayloadAction<Token>) => {
      state.pool = action.payload
    },
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload
    },
  },
})

const filterInputValues = (inputValues, pool) => {
  const chainId = pool.chainId
  const poolTokens = pool.nativeTokens ?? pool.poolTokens

  const poolTokenAddresses = []

  poolTokens.map((nativeToken) => {
    poolTokenAddresses.push(nativeToken.addresses[chainId])
  })

  const filteredObj = poolTokenAddresses.reduce((obj, key) => {
    obj[key] = inputValues.bi.hasOwnProperty(key) ? inputValues.bi[key] : 0n
    return obj
  }, {})

  return filteredObj
}

const filterAndSerializeInputValues = (inputValues, pool) => {
  const chainId = pool.chainId
  const filteredInputValues = filterInputValues(inputValues, pool)
  const showTokens = pool.nativeTokens ?? pool.poolTokens

  const keys = Object.keys(filteredInputValues)

  const serializedValues = { bi: {}, str: {} }

  keys.map((key) => {
    const token = showTokens.find((token) => token.addresses[chainId] === key)

    serializedValues['bi'][key] = filteredInputValues[key]
    serializedValues['str'][key] = formatBigIntToString(
      filteredInputValues[key],
      token.decimals[chainId],
      8
    )
  })

  return serializedValues
}

const sumBigIntegers = (pool, filteredInputValue) => {
  if (!pool) {
    return 0n
  }

  const chainId = pool.chainId

  const showTokens = pool.nativeTokens ?? pool.poolTokens

  return showTokens.reduce((sum, token) => {
    if (Object.keys(filteredInputValue.bi).length === 0) {
      return 0n
    }
    const scalarFactor = pow10BigInt(
      BigInt(18) - BigInt(token.decimals[chainId])
    )

    const valueToAdd =
      BigInt(filteredInputValue.bi[token.addresses[chainId]]) * scalarFactor

    return sum + valueToAdd
  }, 0n)
}

const pow10BigInt = (n) => {
  let result = 1n
  for (let i = 0n; i < n; i++) {
    result *= 10n
  }
  return result
}

export const {
  resetPoolDeposit,
  setDepositQuote,
  setInputValue,
  setIsLoading,
  setPool,
} = poolDepositSlice.actions

export default poolDepositSlice.reducer
