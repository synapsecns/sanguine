import invariant from 'tiny-invariant'
import { BigNumber, PopulatedTransaction } from 'ethers'

import { BigintIsh } from '../constants'
import { SynapseSDK } from '../sdk'
import { handleNativeToken } from '../utils/handleNativeToken'
import { BridgeQuote, SynapseModuleSet, Query } from '../module'

/**
 * Executes a bridge operation between two different chains. Depending on the origin router address, the operation
 * will use either a SynapseRouter or a SynapseCCTPRouter. This function creates a populated transaction ready
 * to be signed and sent to the origin chain.
 *
 * @param to - The recipient address of the bridged tokens.
 * @param originRouterAddress - The address of the origin router.
 * @param originChainId - The ID of the origin chain.
 * @param destChainId - The ID of the destination chain.
 * @param token - The token to bridge.
 * @param amount - The amount of token to bridge.
 * @param originQuery - The query for the origin chain.
 * @param destQuery - The query for the destination chain.
 *
 * @returns A promise that resolves to a populated transaction object which can be used to send the transaction.
 *
 * @throws Will throw an error if there's an issue with the bridge operation.
 */
export async function bridge(
  this: SynapseSDK,
  to: string,
  originRouterAddress: string,
  originChainId: number,
  destChainId: number,
  token: string,
  amount: BigintIsh,
  originQuery: Query,
  destQuery: Query
): Promise<PopulatedTransaction> {
  invariant(
    originChainId !== destChainId,
    'Origin chainId cannot be equal to destination chainId'
  )
  token = handleNativeToken(token)
  // Find the module that is using the given router address
  const module = this.allModuleSets
    .map((set) => set.getModuleWithAddress(originChainId, originRouterAddress))
    .find(Boolean)
  if (!module) {
    throw new Error('Invalid router address')
  }
  return module.bridge(to, destChainId, token, amount, originQuery, destQuery)
}

/**
 * This method tries to fetch the best quote from either the Synapse Router or SynapseCCTP Router.
 * It first handles the native token, then fetches the best quote for both types of routers.
 * If the router addresses are valid for CCTP, it will fetch the quote from the CCTP routers, otherwise it will resolve to undefined.
 * It waits for both types of quotes, then determines the best one by comparing the maximum output amount.
 * If no best quote can be found, it will throw an error.
 *
 * @param originChainId - The ID of the original chain.
 * @param destChainId - The ID of the destination chain.
 * @param tokenIn - The input token.
 * @param tokenOut - The output token.
 * @param amountIn - The amount of input token.
 * @param deadline - The transaction deadline, optional.
 * @param excludeCCTP - Flag to exclude CCTP quotes from the result, optional and defaults to False.
 *
 * @returns - A promise that resolves to the best bridge quote.
 *
 * @throws - Will throw an error if no best quote could be determined.
 */
export async function bridgeQuote(
  this: SynapseSDK,
  originChainId: number,
  destChainId: number,
  tokenIn: string,
  tokenOut: string,
  amountIn: BigintIsh,
  deadline?: BigNumber,
  excludeCCTP: boolean = false
): Promise<BridgeQuote> {
  // Get the quotes sorted by maxAmountOut
  const allQuotes = await allBridgeQuotes.call(
    this,
    originChainId,
    destChainId,
    tokenIn,
    tokenOut,
    amountIn,
    deadline
  )
  // Get the first quote that is not excluded
  const bestQuote = allQuotes.find(
    (quote) =>
      !excludeCCTP ||
      quote.bridgeModuleName !== this.synapseCCTPRouterSet.bridgeModuleName
  )
  if (!bestQuote) {
    throw new Error('No route found')
  }
  return bestQuote
}

/**
 * This method tries to fetch all available quotes from all available bridge modules.
 *
 * @param originChainId - The ID of the original chain.
 * @param destChainId - The ID of the destination chain.
 * @param tokenIn - The input token.
 * @param tokenOut - The output token.
 * @param amountIn - The amount of input token.
 * @param deadline - The transaction deadline, optional.
 * @returns - A promise that resolves to an array of bridge quotes.
 */
export async function allBridgeQuotes(
  this: SynapseSDK,
  originChainId: number,
  destChainId: number,
  tokenIn: string,
  tokenOut: string,
  amountIn: BigintIsh,
  deadline?: BigNumber
): Promise<BridgeQuote[]> {
  invariant(
    originChainId !== destChainId,
    'Origin chainId cannot be equal to destination chainId'
  )
  tokenOut = handleNativeToken(tokenOut)
  tokenIn = handleNativeToken(tokenIn)
  const allQuotes: BridgeQuote[][] = await Promise.all(
    this.allModuleSets.map(async (moduleSet) => {
      const routes = await moduleSet.getBridgeRoutes(
        originChainId,
        destChainId,
        tokenIn,
        tokenOut,
        amountIn
      )
      // Filter out routes with zero minAmountOut and finalize the rest
      return Promise.all(
        routes
          .filter((route) => route.destQuery.minAmountOut.gt(0))
          .map((route) => moduleSet.finalizeBridgeRoute(route, deadline))
      )
    })
  )
  // Flatten the result and sort by maxAmountOut in descending order
  return allQuotes
    .flat()
    .sort((a, b) => (a.maxAmountOut.gt(b.maxAmountOut) ? -1 : 1))
}

/**
 * Applies slippage to the given bridge queries, according to bridge module's slippage tolerance.
 *
 * @param bridgeModuleName - The name of the bridge module.
 * @param originQueryPrecise - The query for the origin chain, coming from `allBridgeQuotes()`.
 * @param destQueryPrecise - The query for the destination chain, coming from `allBridgeQuotes()`.
 * @param slipNumerator - The numerator of the slippage tolerance.
 * @param slipDenominator - The denominator of the slippage tolerance, defaults to 10000.
 * @returns - The origin and destination queries with slippage applied.
 */
export function applyBridgeSlippage(
  this: SynapseSDK,
  bridgeModuleName: string,
  originQueryPrecise: Query,
  destQueryPrecise: Query,
  slipNumerator: number,
  slipDenominator: number = 10000
): { originQuery: Query; destQuery: Query } {
  const moduleSet = getModuleSet.call(this, bridgeModuleName)
  return moduleSet.applySlippage(
    originQueryPrecise,
    destQueryPrecise,
    slipNumerator,
    slipDenominator
  )
}

/**
 * Gets the unique Synapse txId for a bridge operation that happened within a given transaction.
 * Synapse txId is known as "kappa" for SynapseBridge contract and "requestID" for SynapseCCTP contract.
 * This function is meant to abstract away the differences between the two bridge modules.
 *
 * @param originChainId - The ID of the origin chain.
 * @param bridgeModuleName - The name of the bridge module.
 * @param txHash - The transaction hash of the bridge operation on the origin chain.
 * @returns A promise that resolves to the unique Synapse txId of the bridge operation.
 */
export async function getSynapseTxId(
  this: SynapseSDK,
  originChainId: number,
  bridgeModuleName: string,
  txHash: string
): Promise<string> {
  return getModuleSet
    .call(this, bridgeModuleName)
    .getSynapseTxId(originChainId, txHash)
}

/**
 * Checks whether a bridge operation has been completed on the destination chain.
 *
 * @param destChainId - The ID of the destination chain.
 * @param bridgeModuleName - The name of the bridge module.
 * @param synapseTxId - The unique Synapse txId of the bridge operation.
 * @returns A promise that resolves to a boolean indicating whether the bridge operation has been completed.
 */
export async function getBridgeTxStatus(
  this: SynapseSDK,
  destChainId: number,
  bridgeModuleName: string,
  synapseTxId: string
): Promise<boolean> {
  return getModuleSet
    .call(this, bridgeModuleName)
    .getBridgeTxStatus(destChainId, synapseTxId)
}

/**
 * Returns the name of the bridge module that emits the given event.
 * This will be either SynapseBridge or SynapseCCTP.
 *
 * @param eventName - The name of the event.
 * @returns - The name of the bridge module.
 */
export function getBridgeModuleName(
  this: SynapseSDK,
  eventName: string
): string {
  const moduleSet = this.allModuleSets.find((set) =>
    set.allEvents.includes(eventName)
  )
  if (!moduleSet) {
    throw new Error('Unknown event')
  }
  return moduleSet.bridgeModuleName
}

/**
 * Returns the estimated time for a bridge operation from a given origin chain using a given bridge module.
 * This will be the estimated time for the bridge operation to be completed regardless of the destination chain,
 * or the bridge token.
 *
 * @param originChainId - The ID of the origin chain.
 * @param bridgeModuleName - The name of the bridge module.
 * @returns - The estimated time for a bridge operation, in seconds.
 * @throws - Will throw an error if the bridge module is unknown for the given chain.
 */
export function getEstimatedTime(
  this: SynapseSDK,
  originChainId: number,
  bridgeModuleName: string
): number {
  return getModuleSet
    .call(this, bridgeModuleName)
    .getEstimatedTime(originChainId)
}

/**
 * Gets the chain gas amount for the Synapse bridge.
 *
 * @param chainId The chain ID
 * @returns The chain gas amount
 * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
 */
export async function getBridgeGas(
  this: SynapseSDK,
  chainId: number
): Promise<BigNumber> {
  return this.synapseRouterSet.getSynapseRouter(chainId).chainGasAmount()
}

/**
 * Extracts the SynapseModuleSet from the SynapseSDK based on the given bridge module name.
 *
 * @param bridgeModuleName - The name of the bridge module, SynapseBridge or SynapseCCTP.
 * @returns The corresponding SynapseModuleSet.
 * @throws Will throw an error if the bridge module is unknown.
 */
export function getModuleSet(
  this: SynapseSDK,
  bridgeModuleName: string
): SynapseModuleSet {
  const moduleSet = this.allModuleSets.find(
    (set) => set.bridgeModuleName === bridgeModuleName
  )
  if (!moduleSet) {
    throw new Error('Unknown bridge module')
  }
  return moduleSet
}
