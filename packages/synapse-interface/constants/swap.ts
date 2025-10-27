import { USDC, DAI } from '@constants/tokens/bridgeable'
import * as CHAINS from '@constants/chains/master'
import { SwapQuote } from '@types'

export const QUOTE_POLLING_INTERVAL = 10000

export const EMPTY_SWAP_QUOTE: SwapQuote = {
  outputAmount: 0n,
  outputAmountString: '',
  routerAddress: '',
  allowance: 0n,
  exchangeRate: 0n,
  delta: 0n,
}

export const EMPTY_SWAP_QUOTE_ZERO: SwapQuote = {
  outputAmount: 0n,
  outputAmountString: '0',
  routerAddress: '',
  allowance: 0n,
  exchangeRate: 0n,
  delta: 0n,
}

export const DEFAULT_FROM_TOKEN_SYMBOL = 'USDC'
export const DEFAULT_TO_TOKEN_SYMBOL = 'DAI'
export const DEFAULT_FROM_TOKEN = USDC
export const DEFAULT_TO_TOKEN = DAI
export const DEFAULT_SWAPABLE_TYPE = 'USD'
export const DEFAULT_FROM_CHAIN = CHAINS.ETH.id
