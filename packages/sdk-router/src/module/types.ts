import { BigNumber } from '@ethersproject/bignumber'
import { PopulatedTransaction } from 'ethers'

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
  id: string
  feeAmount: BigNumber
  feeConfig: FeeConfig
  routerAddress: string
  maxAmountOut: BigNumber
  originQuery: Query
  destQuery: Query
  estimatedTime: number
  bridgeModuleName: string
  gasDropAmount: BigNumber
  originChainId: number
  destChainId: number
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

/**
 * Quote for a bridge transaction for the supported bridge modules.
 * Returned by SDK to the consumer.
 *
 * @param id - The unique identifier of the bridge quote.
 * @param fromChainId - ID of the origin chain, where funds will be sent from.
 * @param toChainId - ID of the destination chain, where funds will be received.
 * @param expectedToAmount - Expected amount of output tokens on the destination chain.
 * @param routerAddress - Address of the router on the origin chain.
 * @param estimatedTime - Estimated time for the bridge operation to complete.
 * @param moduleName - Name of the module used for the bridge operation.
 * @param gasDropAmount - Amount of gas to be dropped on the destination chain.
 * @param tx - Optional populated transaction for the bridge operation (returned only if `originSender` is provided)
 */
export type BridgeQuoteV2 = {
  id: string
  fromChainId: number
  fromToken: string
  fromAmount: BigNumber
  toChainId: number
  toToken: string
  expectedToAmount: BigNumber
  minToAmount: BigNumber
  routerAddress: string
  estimatedTime: number
  moduleName: string
  gasDropAmount: BigNumber
  tx?: PopulatedTransaction
}

/**
 * Atomic step of an intent.
 *
 * @param fromChainId - ID of the chain, where funds will be sent from.
 * @param fromToken - Address of the token to be sent.
 * @param fromAmount - Amount of tokens to be sent.
 * @param toChainId - ID of the chain, where funds will be received (can be the same as `fromChainId`).
 * @param toToken - Address of the token to be received.
 * @param expectedToAmount - Expected amount of tokens to be received.
 * @param minToAmount - Minimum amount of tokens to be received (slippage protected).
 * @param routerAddress - Address of the router that performs the operation.
 * @param estimatedTime - Estimated time for the operation to complete.
 * @param moduleName - Name of the module used for the operation.
 * @param gasDropAmount - Amount of gas to be dropped after the operation alongside `toToken`.
 * @param tx - Optional populated transaction for the operation (returned only if `fromSender` is provided).
 */
export type IntentStep = {
  fromChainId: number
  fromToken: string
  fromAmount: BigNumber
  toChainId: number
  toToken: string
  expectedToAmount: BigNumber
  minToAmount: BigNumber
  routerAddress: string
  estimatedTime: number
  moduleName: string
  gasDropAmount: BigNumber
  tx?: PopulatedTransaction
}
