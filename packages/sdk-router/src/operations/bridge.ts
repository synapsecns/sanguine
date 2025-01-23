import invariant from 'tiny-invariant'
import { BigNumber, PopulatedTransaction } from 'ethers'

import { BigintIsh } from '../constants'
import { SynapseSDK } from '../sdk'
import { handleNativeToken } from '../utils/handleNativeToken'
import {
  BridgeQuote,
  SynapseModuleSet,
  Query,
  applyDeadlineToQuery,
} from '../module'
import { logger } from '../utils/logger'

/**
 * Creates a populated bridge transaction ready for signing and submission to the origin chain.
 * The method selects the appropriate router based on the origin router address:
 * - `SynapseRouter` is used for SynapseBridge module
 * - `SynapseCCTPRouter` is used for SynapseCCTP module
 * - `FastBridgeRouter` is used for SynapseRFQ module
 *
 * @param to - Recipient address for the bridged tokens on the destination chain.
 * @param originRouterAddress - Address of the router on the origin chain.
 * @param originChainId - ID of the origin chain.
 * @param destChainId - ID of the destination chain.
 * @param token - Address of the token to be bridged.
 * @param amount - Amount of tokens to bridge.
 * @param originQuery - Query for the origin chain, obtained from `allBridgeQuotes()` or `bridgeQuote()`.
 * @param destQuery - Query for the destination chain, obtained from `allBridgeQuotes()` or `bridgeQuote()`.
 *
 * @returns A Promise resolving to a populated transaction object, ready for sending.
 *
 * @throws Error if any issues arise during the bridge operation.
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
 * Options for the bridgeQuote and allBridgeQuotes functions.
 *
 * @param deadline - Optional transaction deadline on the origin chain.
 * @param excludedModules - Optional array of module names to exclude from the quote.
 * @param originUserAddress - Optional address of the user on the origin chain. This parameter must be
 * specified if a smart contract will initiate the bridge operation on behalf of the user.
 */
interface BridgeQuoteOptions {
  deadline?: BigNumber
  excludedModules?: string[]
  originUserAddress?: string
}

/**
 * Retrieves the best quote from all available bridge modules (SynapseBridge, SynapseCCTP, and SynapseRFQ).
 * Users can customize the query by specifying a deadline, excluding certain modules, and providing the user's address on the origin chain.
 *
 * Important: The originUserAddress MUST be provided if a smart contract will initiate the bridge operation on the user's behalf.
 * This applies to smart wallets (e.g., Safe) and third-party integrations (such as bridge aggregator smart contracts).
 *
 * The returned quote will not have any slippage settings applied. To add slippage to the quote, use the `applyBridgeSlippage` function.
 * The returned quote will use the origin chain deadline provided in the options. If no deadline is provided, the module's default origin deadline is used.
 * The returned quote will use the module's default destination deadline.
 *
 * @param originChainId - ID of the origin chain.
 * @param destChainId - ID of the destination chain.
 * @param tokenIn - Address of the token to be bridged from the origin chain.
 * @param tokenOut - Address of the token to be received on the destination chain.
 * @param amountIn - Amount of input tokens on the origin chain.
 * @param options - Optional parameters including origin deadline, excludedModules, and originUserAddress.
 *
 * @returns A promise resolving to the best available bridge quote.
 *
 * @throws An error if no bridge route is found.
 */
export async function bridgeQuote(
  this: SynapseSDK,
  originChainId: number,
  destChainId: number,
  tokenIn: string,
  tokenOut: string,
  amountIn: BigintIsh,
  options: BridgeQuoteOptions = {}
): Promise<BridgeQuote> {
  // Get the quotes sorted by maxAmountOut
  const allQuotes = await allBridgeQuotes.call(
    this,
    originChainId,
    destChainId,
    tokenIn,
    tokenOut,
    amountIn,
    options
  )
  const bestQuote = allQuotes[0]
  if (!bestQuote) {
    throw new Error('No route found')
  }
  return bestQuote
}

/**
 * Fetches all available quotes from the supported bridge modules (SynapseBridge, SynapseCCTP, and SynapseRFQ).
 * Users can customize the query by specifying a deadline, excluding certain modules, and providing the user's address on the origin chain.
 *
 * Important: The originUserAddress MUST be provided if a smart contract will initiate the bridge operation on the user's behalf.
 * This applies to smart wallets (e.g., Safe) and third-party integrations (such as bridge aggregator smart contracts).
 *
 * The returned quotes will not have any slippage settings applied. To add slippage to the quotes, use the `applyBridgeSlippage` function.
 * The returned quotes will use the origin chain deadline provided in the options. If no deadline is provided, the module's default origin deadline is used.
 * The returned quotes will use the module's default destination deadline.
 *
 * @param originChainId - ID of the origin chain.
 * @param destChainId - ID of the destination chain.
 * @param tokenIn - Address of the token to be bridged from the origin chain.
 * @param tokenOut - Address of the token to be received on the destination chain.
 * @param amountIn - Amount of input tokens on the origin chain.
 * @param options - Optional parameters including origin deadline, excludedModules, and originUserAddress.
 *
 * @returns A promise that resolves to an array of bridge quotes.
 * The returned array is sorted by maxAmountOut in descending order, with all quotes having a non-zero amountOut.
 */
export async function allBridgeQuotes(
  this: SynapseSDK,
  originChainId: number,
  destChainId: number,
  tokenIn: string,
  tokenOut: string,
  amountIn: BigintIsh,
  options: BridgeQuoteOptions = {}
): Promise<BridgeQuote[]> {
  invariant(
    originChainId !== destChainId,
    'Origin chainId cannot be equal to destination chainId'
  )
  tokenOut = handleNativeToken(tokenOut)
  tokenIn = handleNativeToken(tokenIn)
  const allQuotes: BridgeQuote[][] = await Promise.all(
    this.allModuleSets.map(async (moduleSet) => {
      // No-op if the module is explicitly excluded
      if (options.excludedModules?.includes(moduleSet.bridgeModuleName)) {
        return []
      }
      const routes = await moduleSet.getBridgeRoutes(
        originChainId,
        destChainId,
        tokenIn,
        tokenOut,
        amountIn,
        options.originUserAddress
      )
      logger.info(
        { routes },
        `Found ${routes.length} routes for ${moduleSet.bridgeModuleName}`
      )
      // Filter out routes with zero minAmountOut and finalize the rest
      return Promise.all(
        routes
          .filter((route) => route.destQuery.minAmountOut.gt(0))
          .map((route) =>
            moduleSet.finalizeBridgeRoute(route, options.deadline)
          )
      )
    })
  )
  // Flatten the result and sort by maxAmountOut in descending order
  return allQuotes
    .flat()
    .sort((a, b) => (a.maxAmountOut.gt(b.maxAmountOut) ? -1 : 1))
}

/**
 * Applies the deadlines to the given bridge queries, according to bridge module's default deadline settings.
 *
 * @param bridgeModuleName - The name of the bridge module.
 * @param originQueryInitial - The query for the origin chain
 * @param destQueryInitial - The query for the destination chain
 * @param originDeadline - The deadline to use on the origin chain (optional, default depends on the module).
 * @param destDeadline - The deadline to use on the destination chain (optional, default depends on the module).
 * @returns The origin and destination queries with the deadlines applied.
 */
export function applyBridgeDeadline(
  this: SynapseSDK,
  bridgeModuleName: string,
  originQueryInitial: Query,
  destQueryInitial: Query,
  originDeadline?: BigNumber,
  destDeadline?: BigNumber
): { originQuery: Query; destQuery: Query } {
  const moduleSet = getModuleSet.call(this, bridgeModuleName)
  const { originModuleDeadline, destModuleDeadline } =
    moduleSet.getModuleDeadlines(originDeadline, destDeadline)
  return {
    originQuery: applyDeadlineToQuery(originQueryInitial, originModuleDeadline),
    destQuery: applyDeadlineToQuery(destQueryInitial, destModuleDeadline),
  }
}

/**
 * Applies slippage to the given bridge queries, according to bridge module's slippage tolerance.
 * Note: default slippage is 10 bips (0.1%).
 *
 * @param bridgeModuleName - The name of the bridge module.
 * @param originQueryInitial - The query for the origin chain, coming from `allBridgeQuotes()`.
 * @param destQueryInitial - The query for the destination chain, coming from `allBridgeQuotes()`.
 * @param slipNumerator - The numerator of the slippage tolerance, defaults to 10.
 * @param slipDenominator - The denominator of the slippage tolerance, defaults to 10000.
 * @returns - The origin and destination queries with slippage applied.
 */
export function applyBridgeSlippage(
  this: SynapseSDK,
  bridgeModuleName: string,
  originQueryInitial: Query,
  destQueryInitial: Query,
  slipNumerator: number = 10,
  slipDenominator: number = 10000
): { originQuery: Query; destQuery: Query } {
  const moduleSet = getModuleSet.call(this, bridgeModuleName)
  return moduleSet.applySlippage(
    originQueryInitial,
    destQueryInitial,
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
