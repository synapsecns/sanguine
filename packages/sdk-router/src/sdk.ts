import { Provider } from '@ethersproject/abstract-provider'
import { JsonRpcProvider } from '@ethersproject/providers'
import invariant from 'tiny-invariant'

import { GasZipModuleSet } from './gaszip'
import { SynapseModuleSet } from './module'
import * as operations from './operations'
import { FastBridgeRouterSet } from './rfq'
import { SynapseRouterSet, SynapseCCTPRouterSet, ChainProvider } from './router'
import { SynapseIntentRouterSet } from './sir/synapseIntentRouterSet'
import { SwapEngineSet } from './swap/swapEngineSet'
import { TokenMetadataFetcher } from './utils'

class SynapseSDK {
  public allModuleSets: SynapseModuleSet[]
  public synapseRouterSet: SynapseRouterSet
  public synapseCCTPRouterSet: SynapseCCTPRouterSet
  public fastBridgeRouterSet: FastBridgeRouterSet
  public gasZipModuleSet: GasZipModuleSet

  public sirSet: SynapseIntentRouterSet
  public swapEngineSet: SwapEngineSet
  public tokenMetadataFetcher: TokenMetadataFetcher
  public providers: { [chainId: number]: Provider }

  /**
   * Constructor for the SynapseSDK class.
   * It sets up the SynapseRouters and SynapseCCTPRouters for the specified chain IDs and providers.
   *
   * @param {number[]} chainIds - The IDs of the chains to initialize routers for.
   * @param {(Provider | string)[]} providersOrUrls - The Ethereum providers for the respective chains or URLs for providers.
   */
  constructor(chainIds: number[], providersOrUrls: (Provider | string)[]) {
    invariant(
      chainIds.length === providersOrUrls.length,
      `Amount of chains and providers does not equal`
    )
    // Zip chainIds and providers into a single object
    const chainProviders: ChainProvider[] = chainIds.map((chainId, index) => ({
      chainId,
      provider:
        typeof providersOrUrls[index] === 'string'
          ? new JsonRpcProvider(providersOrUrls[index] as string)
          : (providersOrUrls[index] as Provider),
    }))
    // Save chainId => provider mapping
    this.providers = {}
    chainProviders.forEach((chainProvider) => {
      this.providers[chainProvider.chainId] = chainProvider.provider
    })
    // Initialize the utility classes
    this.tokenMetadataFetcher = new TokenMetadataFetcher(this.providers)

    // Initialize the Module Sets
    this.synapseRouterSet = new SynapseRouterSet(chainProviders)
    this.synapseCCTPRouterSet = new SynapseCCTPRouterSet(chainProviders)
    this.fastBridgeRouterSet = new FastBridgeRouterSet(chainProviders)
    this.gasZipModuleSet = new GasZipModuleSet(chainProviders)
    this.allModuleSets = [
      this.synapseRouterSet,
      this.synapseCCTPRouterSet,
      this.fastBridgeRouterSet,
      this.gasZipModuleSet,
    ]
    this.sirSet = new SynapseIntentRouterSet(chainProviders)
    this.swapEngineSet = new SwapEngineSet(
      chainProviders,
      this.tokenMetadataFetcher
    )
  }

  public intent = operations.intent

  // Define Bridge operations
  public bridge = operations.bridge
  public bridgeV2 = operations.bridgeV2
  public bridgeQuote = operations.bridgeQuote
  public allBridgeQuotes = operations.allBridgeQuotes
  public getBridgeModuleName = operations.getBridgeModuleName
  public getEstimatedTime = operations.getEstimatedTime
  public getSynapseTxId = operations.getSynapseTxId
  public getBridgeTxStatus = operations.getBridgeTxStatus

  public getBridgeGas = operations.getBridgeGas

  // Define Pool operations
  public getPoolTokens = operations.getPoolTokens
  public getPoolInfo = operations.getPoolInfo
  public getAllPools = operations.getAllPools

  public calculateAddLiquidity = operations.calculateAddLiquidity
  public calculateRemoveLiquidity = operations.calculateRemoveLiquidity
  public calculateRemoveLiquidityOne = operations.calculateRemoveLiquidityOne

  // Define Swap operations
  public swap = operations.swap
  public swapQuote = operations.swapQuote
  public swapV2 = operations.swapV2

  // Define Query operations
  public applyBridgeDeadline = operations.applyBridgeDeadline
  public applyBridgeSlippage = operations.applyBridgeSlippage
  public applySwapDeadline = operations.applySwapDeadline
  public applySwapSlippage = operations.applySwapSlippage
}

export { SynapseSDK }
