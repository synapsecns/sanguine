import invariant from 'tiny-invariant'
import { BigNumber } from '@ethersproject/bignumber'

import { SynapseRouter } from './synapseRouter'
import { ChainProvider, RouterSet } from './routerSet'
import { MEDIAN_TIME_BRIDGE, ROUTER_ADDRESS_MAP } from '../constants'
import { BridgeRoute } from '../module'

/**
 * Wrapper class for interacting with a SynapseRouter contracts deployed on multiple chains.
 */
export class SynapseRouterSet extends RouterSet {
  public readonly bridgeModuleName = 'SynapseBridge'
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
  getGasDropAmount(bridgeRoute: BridgeRoute): Promise<BigNumber> {
    return this.getSynapseRouter(bridgeRoute.destChainId).chainGasAmount()
  }

  /**
   * Returns the existing SynapseRouter instance for the given chain.
   *
   * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
   */
  public getSynapseRouter(chainId: number): SynapseRouter {
    return this.getExistingModule(chainId) as SynapseRouter
  }
}
