import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'

import { FastBridgeRouterSet } from './rfq'
import {
  SynapseRouterSet,
  SynapseCCTPRouterSet,
  ChainProvider,
  PoolToken,
} from './router'
import * as operations from './operations'
import { ETH_NATIVE_TOKEN_ADDRESS } from './utils/handleNativeToken'
import {
  SynapseModuleSet,
  Query,
  applySlippage,
  applySlippageInBips,
  modifyDeadline,
} from './module'

class SynapseSDK {
  public allModuleSets: SynapseModuleSet[]
  public synapseRouterSet: SynapseRouterSet
  public synapseCCTPRouterSet: SynapseCCTPRouterSet
  public fastBridgeRouterSet: FastBridgeRouterSet
  public providers: { [chainId: number]: Provider }

  /**
   * Constructor for the SynapseSDK class.
   * It sets up the SynapseRouters and SynapseCCTPRouters for the specified chain IDs and providers.
   *
   * @param {number[]} chainIds - The IDs of the chains to initialize routers for.
   * @param {Provider[]} providers - The Ethereum providers for the respective chains.
   */
  constructor(chainIds: number[], providers: Provider[]) {
    invariant(
      chainIds.length === providers.length,
      `Amount of chains and providers does not equal`
    )
    // Zip chainIds and providers into a single object
    const chainProviders: ChainProvider[] = chainIds.map((chainId, index) => ({
      chainId,
      provider: providers[index],
    }))
    // Save chainId => provider mapping
    this.providers = {}
    chainProviders.forEach((chainProvider) => {
      this.providers[chainProvider.chainId] = chainProvider.provider
    })
    // Initialize the Module Sets
    this.synapseRouterSet = new SynapseRouterSet(chainProviders)
    this.synapseCCTPRouterSet = new SynapseCCTPRouterSet(chainProviders)
    this.fastBridgeRouterSet = new FastBridgeRouterSet(chainProviders)
    this.allModuleSets = [
      this.synapseRouterSet,
      this.synapseCCTPRouterSet,
      this.fastBridgeRouterSet,
    ]
  }

  // Define Bridge operations
  public bridge = operations.bridge
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

  // Define Query operations
  public applyBridgeSlippage = operations.applyBridgeSlippage
  public applySlippage = applySlippage
  public applySlippageInBips = applySlippageInBips
  public modifyDeadline = modifyDeadline
}

export { SynapseSDK, ETH_NATIVE_TOKEN_ADDRESS, Query, PoolToken }
