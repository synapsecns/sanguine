import { BigNumber } from '@ethersproject/bignumber'
import { BigNumberish } from 'ethers'
import invariant from 'tiny-invariant'
import { uuidv7 } from 'uuidv7'

import { SynapseModule } from './synapseModule'
import {
  BridgeRoute,
  BridgeRouteV2,
  BridgeTokenCandidate,
  FeeConfig,
} from './types'
import { applyOptionalDeadline } from '../utils'
import { Query } from './query'
import { Slippage } from '../swap'
import { BridgeQuote, BridgeQuoteV2 } from '../types'

/**
 * Parameters for `getBridgeTokenCandidates` function.
 *
 * @param originChainId - The ID of the origin chain.
 * @param destChainId - The ID of the destination chain.
 * @param tokenIn - The input token.
 * @param tokenOut - The output token, optional.
 */
export type GetBridgeTokenCandidatesParameters = {
  fromChainId: number
  toChainId: number
  fromToken: string
  toToken?: string
}

/**
 * Parameters for `getBridgeRouteV2` function.
 *
 * @param originAmountIn - The amount of the bridge token on the origin chain.
 * @param bridgeToken - The bridge token to be used for the bridge.
 * @param destTokenOut - The output token on the destination chain that needs to be received.
 * @param originSender - The address of the user on the origin chain.
 * @param destRecipient - The address of the user on the destination chain.
 * @param slippage - The slippage to be used for the swap.
 * @param allowMultipleTxs - Whether to allow multiple transactions for the bridge, in which case the returned BridgeRouteV2
 * might have a fallback quote to the `bridgeToken` instead of `destTokenOut`.
 */
export type GetBridgeRouteV2Parameters = {
  fromAmount: BigNumberish
  bridgeToken: BridgeTokenCandidate
  toToken: string
  fromSender?: string
  toRecipient?: string
  slippage?: Slippage
  allowMultipleTxs?: boolean
}

export abstract class SynapseModuleSet {
  abstract readonly moduleName: string
  abstract readonly allEvents: string[]
  abstract readonly isBridgeV2Supported: boolean

  /**
   * Returns the estimated time for a bridge transaction to be completed,
   * when the transaction is sent from the given chain.
   *
   * @param originChainId - The ID of the origin chain.
   * @returns The estimated time in seconds.
   * @throws Will throw an error if the chain ID is not supported.
   */
  abstract getEstimatedTime(originChainId: number): number

  /**
   * Returns the Synapse transaction ID for a given transaction hash on a given chain.
   * This is used to track the status of a bridge transaction.
   *
   * @param originChainId - The ID of the origin chain.
   * @param txHash - The transaction hash of the bridge transaction.
   * @returns A promise that resolves to the Synapse transaction ID.
   */
  getSynapseTxId(originChainId: number, txHash: string): Promise<string> {
    return this.getExistingModule(originChainId).getSynapseTxId(txHash)
  }

  /**
   * Checks whether a bridge transaction has been completed on the destination chain.
   *
   * @param destChainId - The ID of the destination chain.
   * @param synapseTxId - The unique Synapse txId of the bridge transaction.
   * @returns A promise that resolves to a boolean indicating whether the bridge transaction has been completed.
   */
  getBridgeTxStatus(
    destChainId: number,
    synapseTxId: string
  ): Promise<boolean> {
    return this.getExistingModule(destChainId).getBridgeTxStatus(synapseTxId)
  }

  /**
   * Returns the existing Module instance on the given chain.
   * Returns undefined if a Module instance does not exist on the given chain.
   *
   * @param chainId - The ID of the chain.
   * @returns The Module instance, or undefined if it does not exist.
   */
  abstract getModule(chainId: number): SynapseModule | undefined

  /**
   * Returns the existing Module instance for the given address on the given chain.
   * If the module address is not valid, it will return undefined.
   *
   * @param chainId - The ID of the chain.
   * @param moduleAddress - The address of the module.
   * @returns The Module instance, or undefined if the module address is not valid.
   */
  getModuleWithAddress(
    chainId: number,
    moduleAddress: string
  ): SynapseModule | undefined {
    const module = this.getModule(chainId)
    if (module?.address.toLowerCase() === moduleAddress.toLowerCase()) {
      return module
    }
    return undefined
  }

  /**
   * Returns the existing Module instance for the given chain.
   *
   * @param chainId - The ID of the chain.
   * @returns The Module instance.
   * @throws Will throw an error if the module does not exist.
   */
  getExistingModule(chainId: number): SynapseModule {
    const module = this.getModule(chainId)
    if (!module) {
      throw new Error(`No module found for chain ${chainId}`)
    }
    return module
  }

  /**
   * Returns the list of bridge token candidates that can facilitate a given intent from `tokenIn` to `tokenOut`.
   * If `tokenOut` is not provided, all bridge tokens that can facilitate the intent from `tokenIn` to any token will be returned.
   *
   * @param params - The parameters for the bridge token candidates.
   * @returns A promise that resolves to a list of bridge token candidates.
   */
  abstract getBridgeTokenCandidates(
    params: GetBridgeTokenCandidatesParameters
  ): Promise<BridgeTokenCandidate[]>

  /**
   * Returns the bridge route with a non-zero quote for a given path.
   * If `allowMultipleTxs` is true, the returned BridgeRouteV2 might have a fallback quote to the `bridgeToken` instead of `destTokenOut`,
   * should the direct path to `destTokenOut` not be available through this module.
   *
   * @param params - The parameters for the bridge route.
   * @returns A promise that resolves to the bridge route with a non-zero quote, or undefined if no route is found.
   */
  abstract getBridgeRouteV2(
    params: GetBridgeRouteV2Parameters
  ): Promise<BridgeRouteV2 | undefined>

  /**
   * This method find all possible routes for a bridge transaction between two chains.
   *
   * @param originChainId - The ID of the original chain.
   * @param destChainId - The ID of the destination chain.
   * @param tokenIn - The input token.
   * @param tokenOut - The output token.
   * @param amountIn - The amount of input token.
   * @param originUserAddress - The address of the user on the origin chain.
   *
   * @returns - A list of BridgeRoute objects with the found routes.
   */
  abstract getBridgeRoutes(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigNumberish,
    originUserAddress?: string
  ): Promise<BridgeRoute[]>

  /**
   * Retrieves the fee data for a given bridge route.
   *
   * @param bridgeRoute - The bridge route to get fee data for.
   * @returns A promise that resolves to the fee data.
   */
  abstract getFeeData(bridgeRoute: BridgeRoute): Promise<{
    feeAmount: BigNumber
    feeConfig: FeeConfig
  }>

  /**
   * Retrieves the gas drop amount for a given bridge route.
   * User will receive this amount of gas tokens on the destination chain,
   * when the module transaction is completed.
   *
   * @param destChainId - The ID of the destination chain.
   * @param destBridgeToken - The destination bridge token.
   * @returns A promise that resolves to the gas drop amount.
   */
  abstract getGasDropAmount(
    destChainId: number,
    destBridgeToken: string
  ): Promise<BigNumber>

  /**
   * Returns the default deadline periods for this bridge module.
   *
   * @returns The default deadline periods.
   */
  abstract getDefaultPeriods(): {
    originPeriod: number
    destPeriod: number
  }

  /**
   * Returns the deadlines to use for the given module transaction.
   *
   * @param originDeadline - The deadline to use on the origin chain (default depends on the module).
   * @param destDeadline - The deadline to use on the destination chain (default depends on the module).
   * @returns The deadlines to use.
   */
  public getModuleDeadlines(
    originDeadline?: BigNumber,
    destDeadline?: BigNumber
  ): {
    originModuleDeadline: BigNumber
    destModuleDeadline: BigNumber
  } {
    const { originPeriod, destPeriod } = this.getDefaultPeriods()
    return {
      originModuleDeadline: applyOptionalDeadline(originDeadline, originPeriod),
      destModuleDeadline: applyOptionalDeadline(destDeadline, destPeriod),
    }
  }

  /**
   * Applies the specified slippage to the given queries by modifying the minAmountOut.
   * Note: the original queries are preserved unchanged.
   *
   * @param originQueryPrecise - The query for the origin chain with the precise minAmountOut.
   * @param destQueryPrecise - The query for the destination chain with the precise minAmountOut.
   * @param slipNumerator - The numerator of the slippage.
   * @param slipDenominator - The denominator of the slippage.
   * @returns The modified queries with the reduced minAmountOut.
   */
  abstract applySlippage(
    originQueryPrecise: Query,
    destQueryPrecise: Query,
    slipNumerator: number,
    slipDenominator: number
  ): { originQuery: Query; destQuery: Query }

  /**
   * Finalizes the bridge route by getting fee data and setting default deadlines.
   *
   * @param destChainId - The ID of the destination chain.
   * @param bridgeRoute - Bridge route to finalize.
   * @param originDeadline - The deadline to use on the origin chain (default depends on the module).
   * @param destDeadline - The deadline to use on the destination chain (default depends on the module).
   * @returns The finalized quote with fee data and deadlines.
   */
  async finalizeBridgeRoute(
    bridgeRoute: BridgeRoute,
    originDeadline?: BigNumber,
    destDeadline?: BigNumber
  ): Promise<BridgeQuote> {
    // Check that route is supported on both chains
    const originModule = this.getExistingModule(bridgeRoute.originChainId)
    this.getExistingModule(bridgeRoute.destChainId)
    invariant(
      bridgeRoute.bridgeModuleName === this.moduleName,
      'Invalid bridge module name'
    )
    const uuid = uuidv7()
    const { originQuery, destQuery } = bridgeRoute
    const { originModuleDeadline, destModuleDeadline } =
      this.getModuleDeadlines(originDeadline, destDeadline)
    originQuery.deadline = originModuleDeadline
    destQuery.deadline = destModuleDeadline
    const { feeAmount, feeConfig } = await this.getFeeData(bridgeRoute)
    return {
      id: uuid,
      feeAmount,
      feeConfig,
      routerAddress: originModule.address,
      maxAmountOut: destQuery.minAmountOut,
      originQuery,
      destQuery,
      estimatedTime: this.getEstimatedTime(bridgeRoute.originChainId),
      bridgeModuleName: bridgeRoute.bridgeModuleName,
      gasDropAmount: await this.getGasDropAmount(
        bridgeRoute.destChainId,
        bridgeRoute.bridgeToken.token
      ),
      originChainId: bridgeRoute.originChainId,
      destChainId: bridgeRoute.destChainId,
    }
  }

  async finalizeBridgeQuoteV2(
    bridgeToken: BridgeTokenCandidate,
    bridgeQuote: BridgeQuoteV2
  ): Promise<BridgeQuoteV2> {
    const gasDropAmount = await this.getGasDropAmount(
      bridgeQuote.toChainId,
      bridgeToken.destToken
    )
    return {
      ...bridgeQuote,
      estimatedTime: this.getEstimatedTime(bridgeQuote.fromChainId),
      moduleNames: [...bridgeQuote.moduleNames, this.moduleName],
      gasDropAmount: gasDropAmount.toString(),
    }
  }
}
