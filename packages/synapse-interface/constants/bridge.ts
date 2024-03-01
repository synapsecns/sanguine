import { USDC } from '@constants/tokens/bridgeable'
import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'
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
  outputAmount: 0n,
  outputAmountString: '0',
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
  [CHAINS.BNB.id]: 33,
  [CHAINS.POLYGON.id]: 128,
  [CHAINS.FANTOM.id]: 80,
  [CHAINS.BOBA.id]: 33,
  [CHAINS.OPTIMISM.id]: 750,
  [CHAINS.MOONBEAM.id]: 33,
  [CHAINS.MOONRIVER.id]: 33,
  [CHAINS.ARBITRUM.id]: 200,
  [CHAINS.AVALANCHE.id]: 80,
  [CHAINS.DFK.id]: 33,
  [CHAINS.HARMONY.id]: 33,
  [CHAINS.AURORA.id]: 33,
  [CHAINS.CRONOS.id]: 33,
  [CHAINS.METIS.id]: 33,
  [CHAINS.DOGE.id]: 33,
  [CHAINS.CANTO.id]: 20,
  [CHAINS.BASE.id]: 750,
  [CHAINS.KLAYTN.id]: 20,
  [CHAINS.BLAST.id]: 750,
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

export const ROLE_EVENTS = ['RoleGranted', 'RoleRevoked', 'RoleAdminChanged']

export const INCLUDED_BRIDGE_EVENTS = [
  'TokenMintAndSwap',
  'TokenRedeem',
  'TokenRedeemV2',
  'TokenRedeemAndSwap',
  'TokenRedeemAndRemove',
  'TokenDepositAndSwap',
  'TokenDeposit',
  'TokenMint',
  'TokenWithdraw',
  'TokenWithdrawAndRemove',
]
