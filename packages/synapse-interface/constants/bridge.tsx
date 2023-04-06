import { USDC } from '@constants/tokens/master'
import { Zero } from '@ethersproject/constants'
import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'
import * as CHAINS from '@constants/chains/master'

export const QUOTE_POLLING_INTERVAL = 10000

export const EMPTY_BRIDGE_QUOTE = {
  outputAmount: Zero,
  outputAmountString: '',
  routerAddress: '',
  allowance: Zero,
  exchangeRate: Zero,
  feeAmount: Zero,
  delta: Zero,
  quotes: { originQuery: null, destQuery: null },
}

export const EMPTY_BRIDGE_QUOTE_ZERO = {
  outputAmount: Zero,
  outputAmountString: '0',
  routerAddress: '',
  allowance: Zero,
  exchangeRate: Zero,
  feeAmount: Zero,
  delta: Zero,
  quotes: { originQuery: null, destQuery: null },
}
/**
 * ETH Only Bridge Config used to calculate swap fees
 *
 * abi specified by {@link `@abis/bridgeConfig.json`}
 */
export const BRIDGE_CONFIG_ADDRESSES = {
  [CHAINS.ETH.id]: '0x5217c83ca75559B1f8a8803824E5b7ac233A12a1',
  [CHAINS.POLYGON.id]: '0xd69229f223a8fc84998e1361ae7b4ff724cf4a49', // TESTING ADDRESS
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
  // [ChainId.TERRA]: 1,
  [CHAINS.DOGE.id]: 20,
  [CHAINS.CANTO.id]: 20,
}

export const DEFAULT_FROM_TOKEN_SYMBOL = 'USDC'
export const DEFAULT_TO_TOKEN_SYMBOL = 'USDC'
export const DEFAULT_FROM_TOKEN = USDC
export const DEFAULT_TO_TOKEN = USDC
export const DEFAULT_SWAPABLE_TYPE = 'USD'
export const DEFAULT_FROM_CHAIN = CHAINS.ETH.id
export const DEFAULT_TO_CHAIN = CHAINS.ARBITRUM.id

export const TRANSITIONS_PROPS = {
  ...COIN_SLIDE_OVER_PROPS,
  className: `
    origin-bottom absolute
    w-full h-full
    md:w-[95%] md:h-[95%]
    -ml-0 md:-ml-3
    md:mt-3
    bg-bgBase
    z-20 rounded-3xl
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
    z-20 rounded-3xl
  `,
}
