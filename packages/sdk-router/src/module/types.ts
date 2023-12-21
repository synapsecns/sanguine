import { BigNumber } from '@ethersproject/bignumber'

import { Query } from './query'

/**
 * Matches BridgeToken returned by SynapseRouter (V1) and SynapseCCTPRouter.
 */
export type BridgeToken = {
  symbol: string
  token: string
}

/**
 * Reduces the object to contain only the keys that are present in the BridgeToken type.
 */
export const reduceToBridgeToken = (bridgeToken: BridgeToken): BridgeToken => {
  return {
    symbol: bridgeToken.symbol,
    token: bridgeToken.token,
  }
}

/**
 * Fee configuration for a bridge token used in SynapseRouter (V1) and SynapseCCTPRouter.
 */
export type FeeConfig = {
  bridgeFee: number
  minFee: BigNumber
  maxFee: BigNumber
}

/**
 * Reduces the object to contain only the keys that are present in the FeeConfig type.
 */
export const reduceToFeeConfig = (feeConfig: FeeConfig): FeeConfig => {
  return {
    bridgeFee: feeConfig.bridgeFee,
    minFee: feeConfig.minFee,
    maxFee: feeConfig.maxFee,
  }
}

/**
 * Quote for a swap transaction for SynapseRouter (V1).
 * Returned by SDK to the consumer.
 */
export type SwapQuote = {
  routerAddress: string
  maxAmountOut: BigNumber
  query: Query
}

/**
 * Quote for a bridge transaction for SynapseRouter (V1) and SynapseCCTPRouter.
 * Returned by SDK to the consumer.
 */
export type BridgeQuote = {
  feeAmount: BigNumber
  feeConfig: FeeConfig
  routerAddress: string
  maxAmountOut: BigNumber
  originQuery: Query
  destQuery: Query
  estimatedTime: number
  bridgeModuleName: string
}

/**
 * Internal representation of a found bridge route for SynapseRouter (V1) and SynapseCCTPRouter.
 */
export type BridgeRoute = {
  originChainId: number
  destChainId: number
  originQuery: Query
  destQuery: Query
  bridgeToken: BridgeToken
  bridgeModuleName: string
}

/**
 * Finds the best route: the one with the maximum amount out in the destination query.
 */
export const findBestRoute = (bridgeRoutes: BridgeRoute[]): BridgeRoute => {
  return bridgeRoutes.reduce((best, current) => {
    return current.destQuery.minAmountOut.gt(best.destQuery.minAmountOut)
      ? current
      : best
  })
}
