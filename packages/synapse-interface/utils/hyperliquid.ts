import { ARBITRUM, HYPERLIQUID } from '@/constants/chains/master'

// The existing Hyperliquid destination uses an Arbitrum deposit only for USDC.
// SYN reaches HyperCore directly through the OFT composer.
export const isHyperliquidUsdcDeposit = (
  toChainId: number,
  toToken?: { routeSymbol?: string }
) => toChainId === HYPERLIQUID.id && toToken?.routeSymbol === 'USDC'

export const getBridgeDestinationChainId = (
  toChainId: number,
  toToken?: { routeSymbol?: string }
) => (isHyperliquidUsdcDeposit(toChainId, toToken) ? ARBITRUM.id : toChainId)
