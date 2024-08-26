import { USDC } from '@constants/tokens/bridgeable'
import * as CHAINS from '@constants/chains/master'

export const QUOTE_POLLING_INTERVAL = 10000

export const EMPTY_BRIDGE_QUOTE = {
  inputAmountForQuote: '',
  originTokenForQuote: null,
  destTokenForQuote: null,
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
  timestamp: null,
  originChainId: null,
  destChainId: null,
  requestId: null,
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

export const DEFAULT_FROM_TOKEN_SYMBOL = 'USDC'
export const DEFAULT_TO_TOKEN_SYMBOL = 'USDC'
export const DEFAULT_FROM_TOKEN = USDC
export const DEFAULT_TO_TOKEN = USDC
export const DEFAULT_SWAPABLE_TYPE = 'USD'
export const DEFAULT_FROM_CHAIN = CHAINS.ETH.id
export const DEFAULT_TO_CHAIN = CHAINS.ARBITRUM.id

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
