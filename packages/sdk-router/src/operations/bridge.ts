import { BigNumber, BigNumberish, PopulatedTransaction, utils } from 'ethers'
import invariant from 'tiny-invariant'

import { areIntentsSupported, isChainIdSupported } from '../constants/chainIds'
import {
  SynapseModuleSet,
  Query,
  applyDeadlineToQuery,
  isSwapQuery,
} from '../module'
import { SynapseSDK } from '../sdk'
import {
  EngineID,
  RecipientEntity,
  RouteInput,
  slippageFromPercentage,
  SwapEngineRoute,
} from '../swap'
import { BridgeQuote, BridgeQuoteV2, BridgeV2Parameters } from '../types'
import {
  handleNativeToken,
  handleParams,
  isSameAddress,
  Prettify,
  stringifyPopulatedTransaction,
} from '../utils'

type BridgeV2InternalParameters = Prettify<
  BridgeV2Parameters & { allowMultipleTxs?: boolean }
>

export async function bridgeV2(
  this: SynapseSDK,
  params: BridgeV2Parameters
): Promise<BridgeQuoteV2[]> {
  params = handleParams(params)
  // Don't allow multiple transactions for exported bridgeV2 function.
  return _bridgeV2Internal.call(this, { ...params, allowMultipleTxs: false })
}

export async function _bridgeV2Internal(
  this: SynapseSDK,
  params: BridgeV2InternalParameters
): Promise<BridgeQuoteV2[]> {
  const bridgeV2Modules = this.allModuleSets.filter(
    (set) => set.isBridgeV2Supported
  )
  const [bridgeV1Quotes, bridgeV2Quotes] = await Promise.all([
    _collectV1Quotes.call(this, params, bridgeV2Modules),
    _collectV2Quotes.call(this, params, bridgeV2Modules),
  ])
  // Combine the quotes and sort by expectedToAmount in descending order
  return [...bridgeV1Quotes, ...bridgeV2Quotes].sort((a, b) =>
    BigNumber.from(a.expectedToAmount).gte(b.expectedToAmount) ? -1 : 1
  )
}

async function _collectV1Quotes(
  this: SynapseSDK,
  params: BridgeV2Parameters,
  bridgeV2Modules: SynapseModuleSet[]
): Promise<BridgeQuoteV2[]> {
  if (
    !isChainIdSupported(params.fromChainId) ||
    !isChainIdSupported(params.toChainId)
  ) {
    return []
  }
  const slippage = slippageFromPercentage(params.slippagePercentage)
  const deadlineBN = params.deadline
    ? BigNumber.from(params.deadline)
    : undefined
  const bridgeV1Quotes = await allBridgeQuotes.call(
    this,
    params.fromChainId,
    params.toChainId,
    params.fromToken,
    params.toToken,
    params.fromAmount,
    {
      deadline: deadlineBN,
      originUserAddress: params.fromSender,
      excludedModules: bridgeV2Modules.map((set) => set.moduleName),
    }
  )
  return Promise.all(
    bridgeV1Quotes.map(async (quote: BridgeQuote) => {
      // Apply slippage
      const { originQuery: originQuerySlippage, destQuery: destQuerySlippage } =
        applyBridgeSlippage.call(
          this,
          quote.bridgeModuleName,
          quote.originQuery,
          quote.destQuery,
          slippage?.numerator,
          slippage?.denominator
        )
      // Apply deadline
      const { originQuery, destQuery } = applyBridgeDeadline.call(
        this,
        quote.bridgeModuleName,
        originQuerySlippage,
        destQuerySlippage,
        deadlineBN
      )
      const swapModuleNames = isSwapQuery(originQuery)
        ? [EngineID[EngineID.DefaultPools]]
        : []
      const moduleNames = [...swapModuleNames, quote.bridgeModuleName]
      // Generate the transaction calldata
      const tx = params.fromSender
        ? await bridge.call(
            this,
            params.toRecipient || params.fromSender,
            quote.routerAddress,
            params.fromChainId,
            params.toChainId,
            params.fromToken,
            params.fromAmount,
            originQuery,
            destQuery
          )
        : undefined
      const bridgeQuoteV2: BridgeQuoteV2 = {
        id: quote.id,
        fromChainId: params.fromChainId,
        fromToken: params.fromToken,
        fromAmount: params.fromAmount,
        toChainId: params.toChainId,
        toToken: params.toToken,
        expectedToAmount: quote.maxAmountOut.toString(),
        minToAmount: destQuery.minAmountOut.toString(),
        routerAddress: quote.routerAddress,
        estimatedTime: quote.estimatedTime,
        moduleNames,
        gasDropAmount: quote.gasDropAmount.toString(),
        tx: stringifyPopulatedTransaction(tx),
      }
      return bridgeQuoteV2
    })
  )
}

async function _collectV2Quotes(
  this: SynapseSDK,
  params: BridgeV2InternalParameters,
  bridgeV2Modules: SynapseModuleSet[]
): Promise<BridgeQuoteV2[]> {
  // Intents need to be supported on the `from` chain, while `to` chain needs to be supported in general
  if (
    !areIntentsSupported(params.fromChainId) ||
    !isChainIdSupported(params.toChainId)
  ) {
    return []
  }
  const slippage = slippageFromPercentage(params.slippagePercentage)
  const candidates = await Promise.all(
    bridgeV2Modules.map(async (set) =>
      set.getBridgeTokenCandidates({
        fromChainId: params.fromChainId,
        toChainId: params.toChainId,
        fromToken: params.fromToken,
        toToken: params.allowMultipleTxs ? undefined : params.toToken,
      })
    )
  )
  // Flatten and remove duplicates
  const originCandidates = candidates
    .flat()
    .map((c) => utils.getAddress(c.originToken))
    .filter((c, index, self) => self.indexOf(c) === index)
  const tokenZap = this.swapEngineSet.getTokenZap(params.fromChainId)
  const originRoutes = (
    await Promise.all(
      originCandidates.map(async (originBridgeToken) => {
        const input: RouteInput = {
          chainId: params.fromChainId,
          fromToken: params.fromToken,
          fromAmount: params.fromAmount,
          swapper: tokenZap,
          toToken: originBridgeToken,
          toRecipient: {
            entity: RecipientEntity.Self,
            address: tokenZap,
          },
          restrictComplexity: false,
        }
        const quote = await this.swapEngineSet.getBestQuote(input, {
          allowMultiStep: true,
        })
        if (!quote) {
          return
        }
        const route = await this.swapEngineSet.generateRoute(input, quote, {
          allowMultiStep: true,
          slippage,
        })
        return route
      })
    )
  ).filter((route): route is SwapEngineRoute => route !== undefined)
  const bridgeQuotesV2 = await Promise.all(
    bridgeV2Modules.map(async (moduleSet, index) =>
      Promise.all(
        candidates[index].map(async (bridgeToken) => {
          const originSwapRoute = originRoutes.find((swapRoute) =>
            isSameAddress(bridgeToken.originToken, swapRoute.toToken)
          )
          if (!originSwapRoute) {
            return
          }
          const bridgeRoute = await moduleSet.getBridgeRouteV2({
            originSwapRoute,
            bridgeToken,
            toToken: params.toToken,
            fromSender: params.fromSender,
            toRecipient: params.toRecipient || params.fromSender,
            slippage,
            allowMultipleTxs: params.allowMultipleTxs,
          })
          // Check that a route was found.
          if (!bridgeRoute) {
            return
          }
          // For single-tx params we need to check that the route matches the destination token.
          if (
            !params.allowMultipleTxs &&
            !isSameAddress(bridgeRoute.toToken, params.toToken)
          ) {
            return
          }
          const bridgeQuoteV2 = await this.sirSet.finalizeBridgeRouteV2(
            params.fromToken,
            params.fromAmount,
            originSwapRoute,
            bridgeRoute,
            params.deadline
          )
          return moduleSet.finalizeBridgeQuoteV2(bridgeToken, bridgeQuoteV2)
        })
      )
    )
  )
  return bridgeQuotesV2
    .flat()
    .filter((quote): quote is BridgeQuoteV2 => quote !== undefined)
}

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
  amount: BigNumberish,
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
  amountIn: BigNumberish,
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
  amountIn: BigNumberish,
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
      if (options.excludedModules?.includes(moduleSet.moduleName)) {
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
 * @param moduleName - The name of the bridge module.
 * @param originQueryInitial - The query for the origin chain
 * @param destQueryInitial - The query for the destination chain
 * @param originDeadline - The deadline to use on the origin chain (optional, default depends on the module).
 * @param destDeadline - The deadline to use on the destination chain (optional, default depends on the module).
 * @returns The origin and destination queries with the deadlines applied.
 */
export function applyBridgeDeadline(
  this: SynapseSDK,
  moduleName: string,
  originQueryInitial: Query,
  destQueryInitial: Query,
  originDeadline?: BigNumber,
  destDeadline?: BigNumber
): { originQuery: Query; destQuery: Query } {
  const moduleSet = getModuleSet.call(this, moduleName)
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
 * @param moduleName - The name of the bridge module.
 * @param originQueryInitial - The query for the origin chain, coming from `allBridgeQuotes()`.
 * @param destQueryInitial - The query for the destination chain, coming from `allBridgeQuotes()`.
 * @param slipNumerator - The numerator of the slippage tolerance, defaults to 10.
 * @param slipDenominator - The denominator of the slippage tolerance, defaults to 10000.
 * @returns - The origin and destination queries with slippage applied.
 */
export function applyBridgeSlippage(
  this: SynapseSDK,
  moduleName: string,
  originQueryInitial: Query,
  destQueryInitial: Query,
  slipNumerator: number = 10,
  slipDenominator: number = 10000
): { originQuery: Query; destQuery: Query } {
  const moduleSet = getModuleSet.call(this, moduleName)
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
 * @param moduleName - The name of the bridge module.
 * @param txHash - The transaction hash of the bridge operation on the origin chain.
 * @returns A promise that resolves to the unique Synapse txId of the bridge operation.
 */
export async function getSynapseTxId(
  this: SynapseSDK,
  originChainId: number,
  moduleName: string,
  txHash: string
): Promise<string> {
  return getModuleSet
    .call(this, moduleName)
    .getSynapseTxId(originChainId, txHash)
}

/**
 * Checks whether a bridge operation has been completed on the destination chain.
 *
 * @param destChainId - The ID of the destination chain.
 * @param moduleName - The name of the bridge module.
 * @param synapseTxId - The unique Synapse txId of the bridge operation.
 * @returns A promise that resolves to a boolean indicating whether the bridge operation has been completed.
 */
export async function getBridgeTxStatus(
  this: SynapseSDK,
  destChainId: number,
  moduleName: string,
  synapseTxId: string
): Promise<boolean> {
  return getModuleSet
    .call(this, moduleName)
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
  return moduleSet.moduleName
}

/**
 * Returns the estimated time for a bridge operation from a given origin chain using a given bridge module.
 * This will be the estimated time for the bridge operation to be completed regardless of the destination chain,
 * or the bridge token.
 *
 * @param originChainId - The ID of the origin chain.
 * @param moduleName - The name of the bridge module.
 * @returns - The estimated time for a bridge operation, in seconds.
 * @throws - Will throw an error if the bridge module is unknown for the given chain.
 */
export function getEstimatedTime(
  this: SynapseSDK,
  originChainId: number,
  moduleName: string
): number {
  return getModuleSet.call(this, moduleName).getEstimatedTime(originChainId)
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
 * @param moduleName - The name of the bridge module, SynapseBridge or SynapseCCTP.
 * @returns The corresponding SynapseModuleSet.
 * @throws Will throw an error if the bridge module is unknown.
 */
export function getModuleSet(
  this: SynapseSDK,
  moduleName: string
): SynapseModuleSet {
  const moduleSet = this.allModuleSets.find(
    (set) => set.moduleName === moduleName
  )
  if (!moduleSet) {
    throw new Error('Unknown bridge module')
  }
  return moduleSet
}
