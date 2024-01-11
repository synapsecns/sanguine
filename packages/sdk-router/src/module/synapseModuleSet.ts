import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'

import { BigintIsh } from '../constants'
import { BridgeQuote, BridgeRoute, FeeConfig } from './types'
import { SynapseModule } from './synapseModule'
import { applyOptionalDeadline } from '../utils/deadlines'
import { Query } from './query'

export abstract class SynapseModuleSet {
  abstract readonly bridgeModuleName: string
  abstract readonly allEvents: string[]

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
   * This method find all possible routes for a bridge transaction between two chains.
   *
   * @param originChainId - The ID of the original chain.
   * @param destChainId - The ID of the destination chain.
   * @param tokenIn - The input token.
   * @param tokenOut - The output token.
   * @param amountIn - The amount of input token.
   *
   * @returns - A list of BridgeRoute objects with the found routes.
   */
  abstract getBridgeRoutes(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh
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
   * Returns the default deadline periods for this bridge module.
   *
   * @returns The default deadline periods.
   */
  abstract getDefaultPeriods(): {
    originPeriod: number
    destPeriod: number
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
      bridgeRoute.bridgeModuleName === this.bridgeModuleName,
      'Invalid bridge module name'
    )
    const { originQuery, destQuery } = bridgeRoute
    const { originPeriod, destPeriod } = this.getDefaultPeriods()
    originQuery.deadline = applyOptionalDeadline(originDeadline, originPeriod)
    destQuery.deadline = applyOptionalDeadline(destDeadline, destPeriod)
    const { feeAmount, feeConfig } = await this.getFeeData(bridgeRoute)
    return {
      feeAmount,
      feeConfig,
      routerAddress: originModule.address,
      maxAmountOut: destQuery.minAmountOut,
      originQuery,
      destQuery,
      estimatedTime: this.getEstimatedTime(bridgeRoute.originChainId),
      bridgeModuleName: bridgeRoute.bridgeModuleName,
    }
  }
}
