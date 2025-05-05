import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'
import invariant from 'tiny-invariant'

import { ChainProvider, RouterSet } from './routerSet'
import { BridgeTokenType, SynapseRouter } from './synapseRouter'
import { MEDIAN_TIME_BRIDGE, ROUTER_ADDRESS_MAP } from '../constants'
import { BridgeRouteV2, BridgeTokenCandidate } from '../module'
import { logExecutionTime } from '../utils'

/**
 * Wrapper class for interacting with a SynapseRouter contracts deployed on multiple chains.
 */
export class SynapseRouterSet extends RouterSet {
  public readonly moduleName = 'SynapseBridge'
  public readonly allEvents = [
    'DepositEvent',
    'RedeemEvent',
    'WithdrawEvent',
    'MintEvent',
    'DepositAndSwapEvent',
    'MintAndSwapEvent',
    'RedeemAndSwapEvent',
    'RedeemAndRemoveEvent',
    'WithdrawAndRemoveEvent',
    'RedeemV2Event',
  ]
  public readonly isBridgeV2Supported = false

  constructor(chains: ChainProvider[]) {
    super(chains, ROUTER_ADDRESS_MAP, SynapseRouter)
  }

  /**
   * @inheritdoc RouterSet.getOriginAmountOut
   */
  public getEstimatedTime(chainId: number): number {
    const medianTime =
      MEDIAN_TIME_BRIDGE[chainId as keyof typeof MEDIAN_TIME_BRIDGE]
    invariant(medianTime, `No estimated time for chain ${chainId}`)
    return medianTime
  }

  /**
   * @inheritdoc SynapseModuleSet.getGasDropAmount
   */
  public async getGasDropAmount(
    destChainId: number,
    destBridgeToken: string
  ): Promise<BigNumber> {
    const router = this.getSynapseRouter(destChainId)
    // Gas airdrop exists only for minted tokens
    const tokenType = await router.getBridgeTokenType(destBridgeToken)
    if (tokenType !== BridgeTokenType.Redeem) {
      return Zero
    }
    return router.chainGasAmount()
  }

  @logExecutionTime('SynapseRouterSet.getBridgeRoutes')
  public async getBridgeRoutes(
    ...args: Parameters<RouterSet['getBridgeRoutes']>
  ): ReturnType<RouterSet['getBridgeRoutes']> {
    return super.getBridgeRoutes(...args)
  }

  /**
   * Returns the existing SynapseRouter instance for the given chain.
   *
   * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
   */
  public getSynapseRouter(chainId: number): SynapseRouter {
    return this.getExistingModule(chainId) as SynapseRouter
  }

  public async getBridgeTokenCandidates(): Promise<BridgeTokenCandidate[]> {
    return []
  }

  public async getBridgeRouteV2(): Promise<BridgeRouteV2> {
    throw new Error('BridgeRouteV2 is not supported by ' + this.moduleName)
  }
}
