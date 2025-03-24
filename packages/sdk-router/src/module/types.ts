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

export type BridgeTokenCandidate = {
  originChainId: number
  destChainId: number
  originToken: string
  destToken: string
}

export type BridgeRouteV2 = {
  bridgeToken: BridgeTokenCandidate
  minFromAmount: BigNumber
  toToken: string
  expectedToAmount: BigNumber
  minToAmount: BigNumber
  zapData?: string
}
