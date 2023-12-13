import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'

import {
  SynapseRouterSet,
  SynapseCCTPRouterSet,
  ChainProvider,
  Query,
  PoolToken,
} from './router'
import * as operations from './operations'
import { ETH_NATIVE_TOKEN_ADDRESS } from './utils/handleNativeToken'

class SynapseSDK {
  public synapseRouterSet: SynapseRouterSet
  public synapseCCTPRouterSet: SynapseCCTPRouterSet
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
    // Initialize SynapseRouterSet and SynapseCCTPRouterSet
    this.synapseRouterSet = new SynapseRouterSet(chainProviders)
    this.synapseCCTPRouterSet = new SynapseCCTPRouterSet(chainProviders)
  }

  // Define Bridge operations
  public bridge = operations.bridge
  public bridgeQuote = operations.bridgeQuote
  public getBridgeModuleName = operations.getBridgeModuleName
  public getEstimatedTime = operations.getEstimatedTime
  public getBridgeID = operations.getBridgeID
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
}

export { SynapseSDK, ETH_NATIVE_TOKEN_ADDRESS, Query, PoolToken }
