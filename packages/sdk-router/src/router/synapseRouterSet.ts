import invariant from 'tiny-invariant'

import { SynapseRouter } from './synapseRouter'
import { ChainProvider, RouterSet } from './routerSet'
import { MEDIAN_TIME_BRIDGE, ROUTER_ADDRESS_MAP } from '../constants'

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
   * @inheritdoc RouterSet.getBridgeID
   */
  public async getBridgeID(
    originChainId: number,
    txHash: string
  ): Promise<string> {
    return this.getSynapseRouter(originChainId).getBridgeID(txHash)
  }

  /**
   * @inheritdoc RouterSet.getBridgeTxStatus
   */
  public async getBridgeTxStatus(
    destChainId: number,
    bridgeID: string
  ): Promise<boolean> {
    return this.getSynapseRouter(destChainId).getBridgeTxStatus(bridgeID)
  }

  /**
   * Returns the existing SynapseRouter instance for the given chain.
   *
   * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
   */
  public getSynapseRouter(chainId: number): SynapseRouter {
    invariant(
      this.routers[chainId],
      `No SynapseRouter deployed on chain ${chainId}`
    )
    return this.routers[chainId] as SynapseRouter
  }
}
