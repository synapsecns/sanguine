import * as BRIDGEABLE from './tokens/bridgeable'
import * as CHAINS from './chains/master'

export const BRIDGEABLE_TOKENS = BRIDGEABLE
export const ALL_CHAIN = CHAINS

// Customizable lists
export const CUSTOM_BRIDGEABLE_TOKENS = [
  BRIDGEABLE.ETH,
  BRIDGEABLE.USDC,
  BRIDGEABLE.USDCe,
]
export const CUSTOM_CHAINS = [CHAINS.ETH, CHAINS.ARBITRUM, CHAINS.OPTIMISM]
