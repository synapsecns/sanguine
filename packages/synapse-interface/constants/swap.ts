import { USDC, DAI } from '@constants/tokens/bridgeable'
import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'
import * as CHAINS from '@constants/chains/master'
import { SwapQuote } from '@types'
export const QUOTE_POLLING_INTERVAL = 10000
export const EMPTY_SWAP_QUOTE: SwapQuote = {
  rawOutputAmount: 0n,
  parsedOutputAmount: '',
  routerAddress: '',
  allowance: 0n,
  exchangeRate: 0n,
  delta: 0n,
  quote: null,
}
export const EMPTY_SWAP_QUOTE_ZERO: SwapQuote = {
  rawOutputAmount: 0n,
  parsedOutputAmount: '0',
  routerAddress: '',
  allowance: 0n,
  exchangeRate: 0n,
  delta: 0n,
  quote: null,
}
/**
 * number of required confirmations from bridge
 */
export const BRIDGE_REQUIRED_CONFIRMATIONS = {
  [CHAINS.ETH.id]: 33,
  [CHAINS.BNB.id]: 14,
  [CHAINS.POLYGON.id]: 128,
  [CHAINS.FANTOM.id]: 5,
  [CHAINS.BOBA.id]: 1, // rewrite
  [CHAINS.OPTIMISM.id]: 1, // rewrite
  [CHAINS.MOONBEAM.id]: 21,
  [CHAINS.MOONRIVER.id]: 21, // 5,
  [CHAINS.ARBITRUM.id]: 40,
  [CHAINS.AVALANCHE.id]: 5,
  [CHAINS.DFK.id]: 6,
  [CHAINS.HARMONY.id]: 1, // rewrite
  [CHAINS.AURORA.id]: 5,
  [CHAINS.CRONOS.id]: 6,
  [CHAINS.METIS.id]: 6,
  [CHAINS.DOGE.id]: 20,
  [CHAINS.CANTO.id]: 20,
}

export const DEFAULT_FROM_TOKEN_SYMBOL = 'USDC'
export const DEFAULT_TO_TOKEN_SYMBOL = 'DAI'
export const DEFAULT_FROM_TOKEN = USDC
export const DEFAULT_TO_TOKEN = DAI
export const DEFAULT_SWAPABLE_TYPE = 'USD'
export const DEFAULT_FROM_CHAIN = CHAINS.ETH.id

export const TRANSITIONS_PROPS = {
  ...COIN_SLIDE_OVER_PROPS,
  className: `
    origin-bottom absolute
    w-full h-full
    md:w-[95%] md:h-[95%]
    -ml-0 md:-ml-3
    md:mt-3
    bg-bgBase
    z-20 rounded-lg
  `,
}

export const SETTINGS_TRANSITIONS_PROPS = {
  ...COIN_SLIDE_OVER_PROPS,
  className: `
    origin-bottom absolute
    w-full h-full
    md:w-[95%]
    -ml-0 md:-ml-3
    md:-mt-3
    bg-bgBase
    z-20 rounded-lg
  `,
}
