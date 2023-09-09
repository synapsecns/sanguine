import { BigNumber } from '@ethersproject/bignumber'

import { BigintIsh } from '../constants'

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
 * Matches DestRequest passed to SynapseRouter (V1) and SynapseCCTPRouter.
 */
export type DestRequest = {
  symbol: string
  amountIn: BigintIsh
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
