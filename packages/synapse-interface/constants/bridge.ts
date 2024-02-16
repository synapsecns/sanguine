import { USDC } from '@constants/tokens/bridgeable'
import * as CHAINS from '@constants/chains/master'

export const QUOTE_POLLING_INTERVAL = 10000

export const EMPTY_BRIDGE_QUOTE = {
  outputAmount: 0n,
  outputAmountString: '',
  routerAddress: '',
  allowance: 0n,
  exchangeRate: 0n,
  feeAmount: 0n,
  delta: 0n,
  originQuery: null,
  destQuery: null,
  estimatedTime: null,
  bridgeModuleName: null,
  gasDropAmount: 0n,
}

export const EMPTY_BRIDGE_QUOTE_ZERO = {
  ...EMPTY_BRIDGE_QUOTE,
  outputAmountString: '0',
}




export const DEFAULT_FROM_TOKEN_SYMBOL = 'USDC'
export const DEFAULT_TO_TOKEN_SYMBOL = 'USDC'
export const DEFAULT_FROM_TOKEN = USDC
export const DEFAULT_TO_TOKEN = USDC
export const DEFAULT_SWAPABLE_TYPE = 'USD'
export const DEFAULT_FROM_CHAIN = CHAINS.ETH.id
export const DEFAULT_TO_CHAIN = CHAINS.ARBITRUM.id
