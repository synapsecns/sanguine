import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'

import fastBridgeAbi from '../abi/FastBridge.json'
import fastBridgeRouterAbi from '../abi/FastBridgeRouter.json'
import { FastBridgeRouter as FastBridgeRouterContract } from '../typechain/FastBridgeRouter'
import {
  FastBridge as FastBridgeContract,
  IFastBridge,
} from '../typechain/FastBridge'
import { SynapseModule, Query } from '../module'
import { BigintIsh } from '../constants'

// Define type alias
export type BridgeParams = IFastBridge.BridgeParamsStruct

export class FastBridgeRouter implements SynapseModule {
  static fastBridgeInterface = new Interface(fastBridgeAbi)
  static fastBridgeRouterInterface = new Interface(fastBridgeRouterAbi)

  public readonly address: string
  public readonly chainId: number
  public readonly provider: Provider

  private readonly routerContract: FastBridgeRouterContract
  private fastBridgeContractCache: FastBridgeContract | undefined

  // All possible events emitted by the FastBridge contract in the origin transaction (in alphabetical order)
  private readonly originEvents = ['BridgeRequested']

  constructor(chainId: number, provider: Provider, address: string) {
    invariant(chainId, 'CHAIN_ID_UNDEFINED')
    invariant(provider, 'PROVIDER_UNDEFINED')
    invariant(address, 'ADDRESS_UNDEFINED')
    invariant(FastBridgeRouter.fastBridgeRouterInterface, 'INTERFACE_UNDEFINED')
    invariant(FastBridgeRouter.fastBridgeInterface, 'INTERFACE_UNDEFINED')
    this.chainId = chainId
    this.provider = provider
    this.address = address
    this.routerContract = new Contract(
      address,
      FastBridgeRouter.fastBridgeRouterInterface,
      provider
    ) as FastBridgeRouterContract
  }

  /**
   * @inheritdoc SynapseModule.bridge
   */
  public async bridge(
    to: string,
    destChainId: number,
    token: string,
    amount: BigintIsh,
    originQuery: Query,
    destQuery: Query
  ): Promise<PopulatedTransaction> {
    // TODO: implement
    console.log(to, destChainId, token, amount, originQuery, destQuery)
    return null as any
  }

  /**
   * @inheritdoc SynapseModule.getSynapseTxId
   */
  public async getSynapseTxId(txHash: string): Promise<string> {
    // TODO: implement
    console.log(txHash)
    return null as any
  }

  /**
   * @inheritdoc SynapseModule.getBridgeTxStatus
   */
  public async getBridgeTxStatus(synapseTxId: string): Promise<boolean> {
    // TODO: implement
    console.log(synapseTxId)
    return null as any
  }

  public async getFastBridgeContract(): Promise<FastBridgeContract> {
    // Populate the cache if necessary
    if (!this.fastBridgeContractCache) {
      const address = await this.routerContract.fastBridge()
      this.fastBridgeContractCache = new Contract(
        address,
        FastBridgeRouter.fastBridgeInterface,
        this.provider
      ) as FastBridgeContract
    }
    return this.fastBridgeContractCache
  }
}
