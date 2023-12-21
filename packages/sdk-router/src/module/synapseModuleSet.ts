import { BigNumber } from '@ethersproject/bignumber'

import { BigintIsh } from '../constants'
import { BridgeQuote, BridgeRoute } from '../router/types'
import { SynapseModule } from './synapseModule'

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
   * Finalizes the bridge route by getting fee data and setting default deadlines.
   *
   * @param destChainId - The ID of the destination chain.
   * @param bridgeRoute - Bridge route to finalize.
   * @param deadline - The deadline to use on the origin chain (default 10 mins).
   * @returns The finalized quote with fee data and deadlines.
   */
  abstract finalizeBridgeRoute(
    bridgeRoute: BridgeRoute,
    deadline?: BigNumber
  ): Promise<BridgeQuote>
}
