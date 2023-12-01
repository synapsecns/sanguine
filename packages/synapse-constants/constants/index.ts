import * as TOKENS from './tokens'
import * as CHAINS from './chains'
import { Token, Chain } from './types'

// Customizable lists
export const CUSTOM_BRIDGEABLE_TOKENS = [TOKENS.ETH, TOKENS.USDC, TOKENS.USDCe]
export const CUSTOM_CHAINS = [CHAINS.ETH, CHAINS.ARBITRUM, CHAINS.OPTIMISM]

export { Token, Chain, TOKENS, CHAINS }
