import invariant from 'tiny-invariant'
import { BigNumber, PopulatedTransaction } from 'ethers'

import { BigintIsh } from '../constants'
import { SynapseSDK } from '../sdk'
import { handleNativeToken } from '../utils/handleNativeToken'
import { BridgeQuote, Query, findBestRoute } from '../module'
import { RouterSet } from '../router'

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
  // Get Router instance for given chain and address
  const router =
    this.synapseRouterSet.getModuleWithAddress(
      originChainId,
      originRouterAddress
    ) ??
    this.synapseCCTPRouterSet.getModuleWithAddress(
      originChainId,
      originRouterAddress
    )
  // Throw if Router is not found
  invariant(router, 'Invalid router address')
  // Ask the Router to populate the bridge transaction
  return router.bridge(to, destChainId, token, amount, originQuery, destQuery)
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
  invariant(
    originChainId !== destChainId,
    'Origin chainId cannot be equal to destination chainId'
  )
  tokenOut = handleNativeToken(tokenOut)
  tokenIn = handleNativeToken(tokenIn)
  // Construct objects for both types of routers
  const allSets: { set: RouterSet; exclude: boolean }[] = [
    { set: this.synapseRouterSet, exclude: false },
    { set: this.synapseCCTPRouterSet, exclude: excludeCCTP },
  ]
  // Fetch bridge routes from both types of routers
  const allRoutesPromises = allSets.map(({ set, exclude }) =>
    exclude
      ? Promise.resolve([])
      : set.getBridgeRoutes(
          originChainId,
          destChainId,
          tokenIn,
          tokenOut,
          amountIn
        )
  )
  // Wait for all quotes to resolve and flatten the result
  const allRoutes = (await Promise.all(allRoutesPromises)).flat()
  invariant(allRoutes.length > 0, 'No route found')
  const bestRoute = findBestRoute(allRoutes)
  // Find the Router Set that yielded the best route
  const bestSet: RouterSet = getRouterSet.call(this, bestRoute.bridgeModuleName)
  // Finalize the Bridge Route
  return bestSet.finalizeBridgeRoute(bestRoute, deadline)
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
  return getRouterSet
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
  return getRouterSet
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
  if (this.synapseRouterSet.allEvents.includes(eventName)) {
    return this.synapseRouterSet.bridgeModuleName
  }
  if (this.synapseCCTPRouterSet.allEvents.includes(eventName)) {
    return this.synapseCCTPRouterSet.bridgeModuleName
  }
  throw new Error('Unknown event')
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
  return getRouterSet
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
 * Extracts the RouterSet from the SynapseSDK based on the given bridge module name.
 *
 * @param bridgeModuleName - The name of the bridge module, SynapseBridge or SynapseCCTP.
 * @returns The corresponding RouterSet.
 * @throws Will throw an error if the bridge module is unknown.
 */
export function getRouterSet(
  this: SynapseSDK,
  bridgeModuleName: string
): RouterSet {
  if (this.synapseRouterSet.bridgeModuleName === bridgeModuleName) {
    return this.synapseRouterSet
  }
  if (this.synapseCCTPRouterSet.bridgeModuleName === bridgeModuleName) {
    return this.synapseCCTPRouterSet
  }
  throw new Error('Unknown bridge module')
}
